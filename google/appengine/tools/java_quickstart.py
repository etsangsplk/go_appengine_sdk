#!/usr/bin/env python
#
# Copyright 2007 Google Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
"""Handles the generation of quickstart-web.xml based on servlet annotations."""

from __future__ import with_statement

import cStringIO
import itertools
import os
import subprocess
import xml.etree

from google.appengine.tools import java_utils

_JAVA_VMRUNTIME_PATH = os.path.join(
    'google', 'appengine', 'tools', 'java', 'lib',
    'java-managed-vm', 'appengine-java-vmruntime')
_QUICKSTART_JAR_PATH = os.path.join(
    _JAVA_VMRUNTIME_PATH, 'quickstartgenerator.jar')
_WEBDEFAULT_XML_PATH = os.path.join(
    _JAVA_VMRUNTIME_PATH, 'etc', 'webdefault.xml')

_OLD_NAMESPACE_PREFIX = '{http://java.sun.com/xml/ns/javaee}'
_NEW_NAMESPACE_PREFIX = '{http://xmlns.jcp.org/xml/ns/javaee}'

_SDK_ROOT = os.path.dirname(os.path.dirname(
    os.path.dirname(os.path.dirname(java_utils.__file__))))


def quickstart_generator(war_path, sdk_root=None):
  """Run the quickstart-web.xml generator on the given Java app.

  If the generator succeeds in creating quickstart-web.xml, this method returns
  its contents. Otherwise, it raises RuntimeError. If there was already a
  quickstart-web.xml when this method started, it is removed before generation
  is attempted.

  Args:
    war_path: a string that is the path to a Java app. It should name a
      directory that contains a WEB-INF subdirectory.
    sdk_root: a string that is the path to an App Engine SDK with Java support.

  Returns:
    a tuple (xml, path), where xml is a string that is the contents of the
    generated quickstart-web.xml and path is a string that is the path to the
    quickstart-web.xml file itself.

  Raises:
    CalledProcessError: if the quickstart generation fails.
    IOError: if the quickstart generation apparently succeeds but the
      quickstart-web.xml file is not created. (This should not happen.)
  """
  if not sdk_root:
    sdk_root = _SDK_ROOT

  quickstart_xml_path = os.path.join(war_path, 'WEB-INF', 'quickstart-web.xml')
  if os.path.exists(quickstart_xml_path):
    os.remove(quickstart_xml_path)

  java_home, exec_suffix = java_utils.JavaHomeAndSuffix()
  java_command = os.path.join(java_home, 'bin', 'java') + exec_suffix

  quickstartgenerator_jar = os.path.join(sdk_root, _QUICKSTART_JAR_PATH)
  command = [java_command, '-jar', quickstartgenerator_jar, war_path]
  subprocess.check_call(command)

  with open(quickstart_xml_path) as f:
    return f.read(), quickstart_xml_path


def get_webdefault_xml(sdk_root=None):
  """Returns the contents of webdefault.xml in the SDK.

  webdefault.xml defines the servlets and filters that the Java web server will
  apply to every web app.

  Args:
    sdk_root: a str that is the path to the root of the SDK, or None to use a
      path based on the location of this code.

  Returns:
    the contents of webdefault.xml.
  """
  if not sdk_root:
    sdk_root = _SDK_ROOT
  with open(os.path.join(sdk_root, _WEBDEFAULT_XML_PATH)) as f:
    return f.read()


def remove_mappings(quickstart_xml, webdefault_xml):
  """Removes mappings from quickstart-web.xml that come from webdefault.xml.

  The quickstart-web.xml generated by the quickstart_generator function includes
  the contents of the user's web.xml, entries derived from Java annotations,
  and entries derived from the contents of webdefault.xml. All of those are
  appropriate for the Java web server. But when generating an app.yaml for
  appcfg or dev_appserver, the webdefault.xml entries are not appropriate, since
  app.yaml should only reflect what is specific to the user's app. So this
  function returns a modified quickstart-web.xml from which the webdefault.xml
  entries have been removed. Specifically, we look at the <url-pattern> inside
  every <servlet-mapping> or <filter-mapping> element in webdefault.xml to
  determine default patterns; then we look at those elements inside
  quickstart-web.xml and remove any whose <url-pattern> is one of the default
  patterns.

  Args:
    quickstart_xml: a str that is the contents of a quickstart-web.xml file.
    webdefault_xml: a str that is the contents of a webdefault.xml file.

  Returns:
    a str that is the updated contents for a quickstart-web.xml file,
    appropriate for translation into app.yaml.
  """
  tags_to_examine = ['filter-mapping', 'servlet-mapping']

  default_urls = set()
  default_root = xml.etree.ElementTree.fromstring(webdefault_xml)
  for child in _children_with_tags(default_root, tags_to_examine):
    for grandchild in _children_with_tag(child, 'url-pattern'):
      url = grandchild.text.strip()
      if url.startswith('/'):

        default_urls.add(url)

  to_remove = []
  quickstart_root = xml.etree.ElementTree.fromstring(quickstart_xml)
  for child in _children_with_tags(quickstart_root, tags_to_examine):
    for grandchild in _children_with_tag(child, 'url-pattern'):
      if grandchild.text.strip() in default_urls:
        to_remove.append(child)

  for child in to_remove:
    quickstart_root.remove(child)

  output = cStringIO.StringIO()
  xml.etree.ElementTree.ElementTree(quickstart_root).write(output)
  return output.getvalue()


def _children_with_tags(element, tags):
  """Returns child elements of the given element whose tag is in a given list.

  Args:
    element: an ElementTree.Element.
    tags: a list of strings that are the tags to look for in child elements.

  Returns:
    an iterable of ElementTree.Element instances, which are the children of
    the input element whose tags matched one of the elements of the list.
  """
  return itertools.chain(*(_children_with_tag(element, tag) for tag in tags))


def _children_with_tag(element, tag):
  """Returns child elements of the given element with the given tag.

  Args:
    element: an ElementTree.Element.
    tag: a str that is the tag to look for in child elements.

  Returns:
    an iterable of ElementTree.Element instances, which are the children of
    the input element whose tags matched the given tag.
  """
  return itertools.chain(element.iterfind(_OLD_NAMESPACE_PREFIX + tag),
                         element.iterfind(_NEW_NAMESPACE_PREFIX + tag))
