(function(){"use strict";self.onmessage=function(t){let e=t.data,s=0;e.forEach(a=>{a.key=`${s}`,s++}),self.postMessage(e)}})();
