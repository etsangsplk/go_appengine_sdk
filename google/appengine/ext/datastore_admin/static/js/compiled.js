var g=document,l=Array,m=Error,aa=parseInt,n=String;function ba(a,b){return a.keyCode=b}function p(a,b){return a.currentTarget=b}function q(a,b){return a.disabled=b}
var r="shift",t="exec",u="replace",ca="preventDefault",v="keyCode",w="toString",da="propertyIsEnumerable",ea="checked",x="split",y="style",z="push",fa="slice",ga="listener",A="value",B="indexOf",C="type",ha="name",D="length",E="prototype",ia="target",F="call",G,H=this,I=function(a){var b=typeof a;if("object"==b)if(a){if(a instanceof l)return"array";if(a instanceof Object)return b;var c=Object[E][w][F](a);if("[object Window]"==c)return"object";if("[object Array]"==c||"number"==typeof a[D]&&"undefined"!=
typeof a.splice&&"undefined"!=typeof a[da]&&!a[da]("splice"))return"array";if("[object Function]"==c||"undefined"!=typeof a[F]&&"undefined"!=typeof a[da]&&!a[da]("call"))return"function"}else return"null";else if("function"==b&&"undefined"==typeof a[F])return"object";return b},ja=function(a){var b=I(a);return"array"==b||"object"==b&&"number"==typeof a[D]},J=function(a){return"string"==typeof a},ka=function(a){var b=typeof a;return"object"==b&&null!=a||"function"==b},la=function(a,b){function c(){}
c.prototype=b[E];a.D=b[E];a.prototype=new c;a.G=function(a,c,f){for(var h=l(arguments[D]-2),k=2;k<arguments[D];k++)h[k-2]=arguments[k];return b[E][c].apply(a,h)}};var K=function(a){if(m.captureStackTrace)m.captureStackTrace(this,K);else{var b=m().stack;b&&(this.stack=b)}a&&(this.message=n(a))};la(K,m);K[E].name="CustomError";var ma=function(a,b){for(var c=a[x]("%s"),d="",e=l[E][fa][F](arguments,1);e[D]&&1<c[D];)d+=c[r]()+e[r]();return d+c.join("%s")},na=n[E].trim?function(a){return a.trim()}:function(a){return a[u](/^[\s\xa0]+|[\s\xa0]+$/g,"")},va=function(a,b){if(b)a=a[u](oa,"&amp;")[u](pa,"&lt;")[u](qa,"&gt;")[u](ra,"&quot;")[u](sa,"&#39;")[u](ta,"&#0;");else{if(!ua.test(a))return a;-1!=a[B]("&")&&(a=a[u](oa,"&amp;"));-1!=a[B]("<")&&(a=a[u](pa,"&lt;"));-1!=a[B](">")&&(a=a[u](qa,"&gt;"));-1!=a[B]('"')&&(a=a[u](ra,"&quot;"));
-1!=a[B]("'")&&(a=a[u](sa,"&#39;"));-1!=a[B]("\x00")&&(a=a[u](ta,"&#0;"))}return a},oa=/&/g,pa=/</g,qa=/>/g,ra=/"/g,sa=/'/g,ta=/\x00/g,ua=/[\x00&<>"']/,wa=function(a,b){return a<b?-1:a>b?1:0},xa=function(a){return n(a)[u](/\-([a-z])/g,function(a,c){return c.toUpperCase()})},ya=function(a,b){var c=J(b)?n(b)[u](/([-()\[\]{}+?*.$\^|,:#<!\\])/g,"\\$1")[u](/\x08/g,"\\x08"):"\\s";return a[u](new RegExp("(^"+(c?"|["+c+"]+":"")+")([a-z])","g"),function(a,b,c){return b+c.toUpperCase()})};var za=function(a,b){b.unshift(a);K[F](this,ma.apply(null,b));b[r]()};la(za,K);za[E].name="AssertionError";var L=function(a,b,c){if(!a){var d="Assertion failed";if(b)var d=d+(": "+b),e=l[E][fa][F](arguments,2);throw new za(""+d,e||[]);}return a};var M=l[E],Aa=M[B]?function(a,b,c){L(null!=a[D]);return M[B][F](a,b,c)}:function(a,b,c){c=null==c?0:0>c?Math.max(0,a[D]+c):c;if(J(a))return J(b)&&1==b[D]?a[B](b,c):-1;for(;c<a[D];c++)if(c in a&&a[c]===b)return c;return-1},Ba=M.forEach?function(a,b,c){L(null!=a[D]);M.forEach[F](a,b,c)}:function(a,b,c){for(var d=a[D],e=J(a)?a[x](""):a,f=0;f<d;f++)f in e&&b[F](c,e[f],f,a)},Ca=function(a){var b=a[D];if(0<b){for(var c=l(b),d=0;d<b;d++)c[d]=a[d];return c}return[]};var Da=function(a,b,c){for(var d in a)b[F](c,a[d],d,a)},Ea="constructor hasOwnProperty isPrototypeOf propertyIsEnumerable toLocaleString toString valueOf".split(" "),Fa=function(a,b){for(var c,d,e=1;e<arguments[D];e++){d=arguments[e];for(c in d)a[c]=d[c];for(var f=0;f<Ea[D];f++)c=Ea[f],Object[E].hasOwnProperty[F](d,c)&&(a[c]=d[c])}};var N;a:{var Ga=H.navigator;if(Ga){var Ha=Ga.userAgent;if(Ha){N=Ha;break a}}N=""};var O=function(){return-1!=N[B]("Edge")};var Ia=-1!=N[B]("Opera")||-1!=N[B]("OPR"),P=-1!=N[B]("Edge")||-1!=N[B]("Trident")||-1!=N[B]("MSIE"),Q=-1!=N[B]("Gecko")&&!(-1!=N.toLowerCase()[B]("webkit")&&!O())&&!(-1!=N[B]("Trident")||-1!=N[B]("MSIE"))&&!O(),R=-1!=N.toLowerCase()[B]("webkit")&&!O(),Ja=function(){var a=N;if(Q)return/rv\:([^\);]+)(\)|;)/[t](a);if(P&&O())return/Edge\/([\d\.]+)/[t](a);if(P)return/\b(?:MSIE|rv)[: ]([^\);]+)(\)|;)/[t](a);if(R)return/WebKit\/(\S+)/[t](a)},Ka=function(){var a=H.document;return a?a.documentMode:void 0},
La=function(){if(Ia&&H.opera){var a=H.opera.version;return"function"==I(a)?a():a}var a="",b=Ja();b&&(a=b?b[1]:"");return P&&!O()&&(b=Ka(),b>parseFloat(a))?n(b):a}(),Ma={},S=function(a){var b;if(!(b=Ma[a])){b=0;for(var c=na(n(La))[x]("."),d=na(n(a))[x]("."),e=Math.max(c[D],d[D]),f=0;0==b&&f<e;f++){var h=c[f]||"",k=d[f]||"",W=RegExp("(\\d*)(\\D*)","g"),lb=RegExp("(\\d*)(\\D*)","g");do{var X=W[t](h)||["","",""],Y=lb[t](k)||["","",""];if(0==X[0][D]&&0==Y[0][D])break;b=wa(0==X[1][D]?0:aa(X[1],10),0==Y[1][D]?
0:aa(Y[1],10))||wa(0==X[2][D],0==Y[2][D])||wa(X[2],Y[2])}while(0==b)}b=Ma[a]=0<=b}return b},Na=H.document,Oa=Ka(),Pa=!Na||!P||!Oa&&O()?void 0:Oa||("CSS1Compat"==Na.compatMode?aa(La,10):5);var Qa=!P||P&&(O()||9<=Pa);!Q&&!P||P&&P&&(O()||9<=Pa)||Q&&S("1.9.1");P&&S("9");var T=function(a,b){return J(b)?a.getElementById(b):b},Ra=function(a,b,c,d){a=d||a;var e=b&&"*"!=b?b.toUpperCase():"";if(a.querySelectorAll&&a.querySelector&&(e||c))return a.querySelectorAll(e+(c?"."+c:""));if(c&&a.getElementsByClassName){b=a.getElementsByClassName(c);if(e){a={};for(var f=d=0,h;h=b[f];f++)e==h.nodeName&&(a[d++]=h);a.length=d;return a}return b}b=a.getElementsByTagName(e||"*");if(c){a={};for(f=d=0;h=b[f];f++){var e=h.className,k;if(k="function"==typeof e[x])k=0<=Aa(e[x](/\s+/),c);k&&
(a[d++]=h)}a.length=d;return a}return b},Ta=function(a,b){Da(b,function(b,d){"style"==d?a[y].cssText=b:"class"==d?a.className=b:"for"==d?a.htmlFor=b:d in Sa?a.setAttribute(Sa[d],b):0==d.lastIndexOf("aria-",0)||0==d.lastIndexOf("data-",0)?a.setAttribute(d,b):a[d]=b})},Sa={cellpadding:"cellPadding",cellspacing:"cellSpacing",colspan:"colSpan",frameborder:"frameBorder",height:"height",maxlength:"maxLength",role:"role",rowspan:"rowSpan",type:"type",usemap:"useMap",valign:"vAlign",width:"width"},Va=function(a,
b,c){var d=arguments,e=d[0],f=d[1];if(!Qa&&f&&(f[ha]||f[C])){e=["<",e];f[ha]&&e[z](' name="',va(f[ha]),'"');if(f[C]){e[z](' type="',va(f[C]),'"');var h={};Fa(h,f);delete h[C];f=h}e[z](">");e=e.join("")}e=g.createElement(e);f&&(J(f)?e.className=f:"array"==I(f)?e.className=f.join(" "):Ta(e,f));2<d[D]&&Ua(g,e,d,2);return e},Ua=function(a,b,c,d){function e(c){c&&b.appendChild(J(c)?a.createTextNode(c):c)}for(;d<c[D];d++){var f=c[d];if(!ja(f)||ka(f)&&0<f.nodeType)e(f);else{var h;a:{if(f&&"number"==typeof f[D]){if(ka(f)){h=
"function"==typeof f.item||"string"==typeof f.item;break a}if("function"==I(f)){h="function"==typeof f.item;break a}}h=!1}Ba(h?Ca(f):f,e)}}};var Wa=function(a){var b=a[C];if(void 0===b)return null;switch(b.toLowerCase()){case "checkbox":case "radio":return a[ea]?a[A]:null;case "select-one":return b=a.selectedIndex,0<=b?a.options[b][A]:null;case "select-multiple":for(var b=[],c,d=0;c=a.options[d];d++)c.selected&&b[z](c[A]);return b[D]?b:null;default:return void 0!==a[A]?a[A]:null}};var Xa=function(a){Xa[" "](a);return a};Xa[" "]=function(){};var Ya=!P||P&&(O()||9<=Pa),Za=P&&!S("9");!R||S("528");Q&&S("1.9b")||P&&S("8")||Ia&&S("9.5")||R&&S("528");Q&&!S("8")||P&&S("9");var $a=function(a,b){this.type=a;this.target=b;p(this,this[ia]);this.defaultPrevented=this.w=!1};$a[E].preventDefault=function(){this.defaultPrevented=!0};var U=function(a,b){$a[F](this,a?a[C]:"");this.target=null;p(this,null);this.relatedTarget=null;this.button=this.screenY=this.screenX=this.clientY=this.clientX=this.offsetY=this.offsetX=0;ba(this,0);this.charCode=0;this.metaKey=this.shiftKey=this.altKey=this.ctrlKey=!1;this.u=this.state=null;a&&this.B(a,b)};la(U,$a);
U[E].B=function(a,b){var c=this.type=a[C];this.target=a[ia]||a.srcElement;p(this,b);var d=a.relatedTarget;if(d){if(Q){var e;a:{try{Xa(d.nodeName);e=!0;break a}catch(f){}e=!1}e||(d=null)}}else"mouseover"==c?d=a.fromElement:"mouseout"==c&&(d=a.toElement);this.relatedTarget=d;this.offsetX=R||void 0!==a.offsetX?a.offsetX:a.layerX;this.offsetY=R||void 0!==a.offsetY?a.offsetY:a.layerY;this.clientX=void 0!==a.clientX?a.clientX:a.pageX;this.clientY=void 0!==a.clientY?a.clientY:a.pageY;this.screenX=a.screenX||
0;this.screenY=a.screenY||0;this.button=a.button;ba(this,a[v]||0);this.charCode=a.charCode||("keypress"==c?a[v]:0);this.ctrlKey=a.ctrlKey;this.altKey=a.altKey;this.shiftKey=a.shiftKey;this.metaKey=a.metaKey;this.state=a.state;this.u=a;a.defaultPrevented&&this[ca]()};U[E].preventDefault=function(){U.D[ca][F](this);var a=this.u;if(a[ca])a[ca]();else if(a.returnValue=!1,Za)try{(a.ctrlKey||112<=a[v]&&123>=a[v])&&ba(a,-1)}catch(b){}};var ab="closure_listenable_"+(1E6*Math.random()|0),bb=0;var cb=function(a,b,c,d,e,f){this.listener=a;this.i=b;this.src=c;this.type=d;this.m=!!e;this.o=f;this.key=++bb;this.g=this.l=!1};cb[E].v=function(){this.g=!0;this.o=this.src=this.i=this.listener=null};var db=function(a){this.src=a;this.b={};this.s=0};db[E].add=function(a,b,c,d,e){var f=a[w]();a=this.b[f];a||(a=this.b[f]=[],this.s++);var h;a:{for(h=0;h<a[D];++h){var k=a[h];if(!k.g&&k[ga]==b&&k.m==!!d&&k.o==e)break a}h=-1}-1<h?(b=a[h],c||(b.l=!1)):(b=new cb(b,null,this.src,f,!!d,e),b.l=c,a[z](b));return b};db[E].C=function(a){var b=a[C];if(!(b in this.b))return!1;var c=this.b[b],d=Aa(c,a),e;if(e=0<=d)L(null!=c[D]),M.splice[F](c,d,1);e&&(a.v(),0==this.b[b][D]&&(delete this.b[b],this.s--));return e};var eb="closure_lm_"+(1E6*Math.random()|0),fb={},gb=0,hb=function(a,b,c,d,e){if("array"==I(b)){for(var f=0;f<b[D];f++)hb(a,b[f],c,d,e);return null}c=ib(c);if(a&&a[ab])a=a.H(b,c,d,e);else{if(!b)throw m("Invalid event type");var f=!!d,h=jb(a);h||(a[eb]=h=new db(a));c=h.add(b,c,!1,d,e);c.i||(d=kb(),c.i=d,d.src=a,d.listener=c,a.addEventListener?a.addEventListener(b[w](),d,f):a.attachEvent(mb(b[w]()),d),gb++);a=c}return a},kb=function(){var a=nb,b=Ya?function(c){return a[F](b.src,b[ga],c)}:function(c){c=
a[F](b.src,b[ga],c);if(!c)return c};return b},mb=function(a){return a in fb?fb[a]:fb[a]="on"+a},pb=function(a,b,c,d){var e=!0;if(a=jb(a))if(b=a.b[b[w]()])for(b=b.concat(),a=0;a<b[D];a++){var f=b[a];f&&f.m==c&&!f.g&&(f=ob(f,d),e=e&&!1!==f)}return e},ob=function(a,b){var c=a[ga],d=a.o||a.src;if(a.l&&"number"!=typeof a&&a&&!a.g){var e=a.src;if(e&&e[ab])e.I(a);else{var f=a[C],h=a.i;e.removeEventListener?e.removeEventListener(f,h,a.m):e.detachEvent&&e.detachEvent(mb(f),h);gb--;(f=jb(e))?(f.C(a),0==f.s&&
(f.src=null,e[eb]=null)):a.v()}}return c[F](d,b)},nb=function(a,b){if(a.g)return!0;if(!Ya){var c;if(!(c=b))a:{c=["window","event"];for(var d=H,e;e=c[r]();)if(null!=d[e])d=d[e];else{c=null;break a}c=d}e=c;c=new U(e,this);d=!0;if(!(0>e[v]||void 0!=e.returnValue)){a:{var f=!1;if(0==e[v])try{ba(e,-1);break a}catch(h){f=!0}if(f||void 0==e.returnValue)e.returnValue=!0}e=[];for(f=c.currentTarget;f;f=f.parentNode)e[z](f);for(var f=a[C],k=e[D]-1;!c.w&&0<=k;k--){p(c,e[k]);var W=pb(e[k],f,!0,c),d=d&&W}for(k=
0;!c.w&&k<e[D];k++)p(c,e[k]),W=pb(e[k],f,!1,c),d=d&&W}return d}return ob(a,new U(b,this))},jb=function(a){a=a[eb];return a instanceof db?a:null},qb="__closure_events_fn_"+(1E9*Math.random()>>>0),ib=function(a){L(a,"Listener can not be null.");if("function"==I(a))return a;L(a.handleEvent,"An object listener must have handleEvent method.");a[qb]||(a[qb]=function(b){return a.handleEvent(b)});return a[qb]};var sb=function(a,b,c){var d=rb[c];if(!d){var e=xa(c),d=e;void 0===a[y][e]&&(e=(R?"Webkit":Q?"Moz":P?"ms":Ia?"O":null)+ya(e),void 0!==a[y][e]&&(d=e));rb[c]=d}(c=d)&&(a[y][c]=b)},rb={};P&&S(12);var tb=function(a,b){var c=[];1<arguments[D]&&(c=l[E][fa][F](arguments)[fa](1));var d=Ra(g,"th","tct-selectall",a);if(0!=d[D]){var d=d[0],e=0,f=Ra(g,"tbody",null,a);f[D]&&(e=f[0].rows[D]);this.c=Va("input",{type:"checkbox"});d.appendChild(this.c);e?hb(this.c,"click",this.F,!1,this):q(this.c,!0);this.f=[];this.h=[];this.j=[];d=Ra(g,"input",null,a);for(e=0;f=d[e];e++)"checkbox"==f[C]&&f!=this.c?(this.f[z](f),hb(f,"click",this.A,!1,this)):"action"==f[ha]&&(0<=c[B](f[A])?this.j[z](f):this.h[z](f),q(f,
!0))}};G=tb[E];G.f=null;G.a=0;G.c=null;G.h=null;G.j=null;G.F=function(a){for(var b=a[ia][ea],c=a=0,d;d=this.f[c];c++)d.checked=b,a+=1;this.a=b?this.f[D]:0;for(c=0;b=this.h[c];c++)q(b,!this.a);for(c=0;b=this.j[c];c++)q(b,1!=a?!0:!1)};G.A=function(a){this.a+=a[ia][ea]?1:-1;this.c.checked=this.a==this.f[D];a=0;for(var b;b=this.h[a];a++)q(b,!this.a);for(a=0;b=this.j[a];a++)q(b,1!=this.a?!0:!1)};var ub=function(){var a=T(g,"kinds");a&&new tb(a);(a=T(g,"pending_backups"))&&new tb(a);(a=T(g,"backups"))&&new tb(a,"Restore");var b=T(g,"ae-datastore-admin-filesystem");b&&hb(b,"change",function(){var a="gs"==Wa(b);T(g,"gs_bucket_tr")[y].display=a?"":"none"});if(a=T(g,"confirm_delete_form")){var c=T(g,"confirm_readonly_delete");c&&(a.onsubmit=function(){var a=T(g,"confirm_message");if(J("color"))sb(a,"red","color");else for(var b in"color")sb(a,"color"[b],b);return c[ea]})}},V=["ae","Datastore",
"Admin","init"],Z=H;V[0]in Z||!Z.execScript||Z.execScript("var "+V[0]);for(var vb;V[D]&&(vb=V[r]());)V[D]||void 0===ub?Z=Z[vb]?Z[vb]:Z[vb]={}:Z[vb]=ub;
