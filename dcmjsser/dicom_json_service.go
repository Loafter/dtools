package main

import "io/ioutil"
import "net/http"
import "strconv"
import "log"
import "encoding/json"
import "errors"
import "time"
import "os"

import "encoding/base64"

const htmlData = "PCFkb2N0eXBlIGh0bWw+DQo8aHRtbD4NCg0KPGhlYWQ+DQoJPHRpdGxlPmR0b29scyBVSTwvdGl0bGU+DQoJPG1ldGEgbmFtZT0idmlld3BvcnQiIGNvbnRlbnQ9IndpZHRoPWRldmljZS13aWR0aCI+DQoJPGxpbmsgcmVsPSJzdHlsZXNoZWV0IiBocmVmPSJodHRwczovL25ldGRuYS5ib290c3RyYXBjZG4uY29tL2Jvb3Rzd2F0Y2gvMy4wLjAvc2xhdGUvYm9vdHN0cmFwLm1pbi5jc3MiPg0KCTxzY3JpcHQgdHlwZT0idGV4dC9qYXZhc2NyaXB0IiBzcmM9Imh0dHBzOi8vYWpheC5nb29nbGVhcGlzLmNvbS9hamF4L2xpYnMvanF1ZXJ5LzIuMC4zL2pxdWVyeS5taW4uanMiPjwvc2NyaXB0Pg0KCTxzY3JpcHQgdHlwZT0idGV4dC9qYXZhc2NyaXB0IiBzcmM9Imh0dHBzOi8vbmV0ZG5hLmJvb3RzdHJhcGNkbi5jb20vYm9vdHN0cmFwLzMuMS4xL2pzL2Jvb3RzdHJhcC5taW4uanMiPjwvc2NyaXB0Pg0KCTxzdHlsZSB0eXBlPSJ0ZXh0L2NzcyI+DQoJCWJvZHkgew0KCQkJcGFkZGluZy10b3A6IDIwcHg7DQoJCX0NCgkJLmZvb3RlciB7DQoJCQlib3JkZXItdG9wOiAxcHggc29saWQgI2VlZTsNCgkJCW1hcmdpbi10b3A6IDQwcHg7DQoJCQlwYWRkaW5nLXRvcDogNDBweDsNCgkJCXBhZGRpbmctYm90dG9tOiA0MHB4Ow0KCQl9DQoJCS8qIE1haW4gbWFya2V0aW5nIG1lc3NhZ2UgYW5kIHNpZ24gdXAgYnV0dG9uICovDQoJCQ0KCQkuanVtYm90cm9uIHsNCgkJCXRleHQtYWxpZ246IGNlbnRlcjsNCgkJCWJhY2tncm91bmQtY29sb3I6IHRyYW5zcGFyZW50Ow0KCQl9DQoJCS5qdW1ib3Ryb24gLmJ0biB7DQoJCQlmb250LXNpemU6IDIxcHg7DQoJCQlwYWRkaW5nOiAxNHB4IDI0cHg7DQoJCX0NCgkJLyogQ3VzdG9taXplIHRoZSBuYXYtanVzdGlmaWVkIGxpbmtzIHRvIGJlIGZpbGwgdGhlIGVudGlyZSBzcGFjZSBvZiB0aGUgLm5hdmJhciAqLw0KCQkNCgkJLm5hdi1qdXN0aWZpZWQgew0KCQkJYmFja2dyb3VuZC1jb2xvcjogI2VlZTsNCgkJCWJvcmRlci1yYWRpdXM6IDVweDsNCgkJCWJvcmRlcjogMXB4IHNvbGlkICNjY2M7DQoJCX0NCgkJLm5hdi1qdXN0aWZpZWQgPiBsaSA+IGEgew0KCQkJcGFkZGluZy10b3A6IDE1cHg7DQoJCQlwYWRkaW5nLWJvdHRvbTogMTVweDsNCgkJCWNvbG9yOiAjNzc3Ow0KCQkJZm9udC13ZWlnaHQ6IGJvbGQ7DQoJCQl0ZXh0LWFsaWduOiBjZW50ZXI7DQoJCQlib3JkZXItYm90dG9tOiAxcHggc29saWQgI2Q1ZDVkNTsNCgkJCWJhY2tncm91bmQtY29sb3I6ICNlNWU1ZTU7DQoJCQkvKiBPbGQgYnJvd3NlcnMgKi8NCgkJCQ0KCQkJYmFja2dyb3VuZC1yZXBlYXQ6IHJlcGVhdC14Ow0KCQkJLyogUmVwZWF0IHRoZSBncmFkaWVudCAqLw0KCQkJDQoJCQliYWNrZ3JvdW5kLWltYWdlOiAtbW96LWxpbmVhci1ncmFkaWVudCh0b3AsICNmNWY1ZjUgMCUsICNlNWU1ZTUgMTAwJSk7DQoJCQkvKiBGRjMuNisgKi8NCgkJCQ0KCQkJYmFja2dyb3VuZC1pbWFnZTogLXdlYmtpdC1ncmFkaWVudChsaW5lYXIsIGxlZnQgdG9wLCBsZWZ0IGJvdHRvbSwgY29sb3Itc3RvcCgwJSwgI2Y1ZjVmNSksIGNvbG9yLXN0b3AoMTAwJSwgI2U1ZTVlNSkpOw0KCQkJLyogQ2hyb21lLFNhZmFyaTQrICovDQoJCQkNCgkJCWJhY2tncm91bmQtaW1hZ2U6IC13ZWJraXQtbGluZWFyLWdyYWRpZW50KHRvcCwgI2Y1ZjVmNSAwJSwgI2U1ZTVlNSAxMDAlKTsNCgkJCS8qIENocm9tZSAxMCssU2FmYXJpIDUuMSsgKi8NCgkJCQ0KCQkJYmFja2dyb3VuZC1pbWFnZTogLW1zLWxpbmVhci1ncmFkaWVudCh0b3AsICNmNWY1ZjUgMCUsICNlNWU1ZTUgMTAwJSk7DQoJCQkvKiBJRTEwKyAqLw0KCQkJDQoJCQliYWNrZ3JvdW5kLWltYWdlOiAtby1saW5lYXItZ3JhZGllbnQodG9wLCAjZjVmNWY1IDAlLCAjZTVlNWU1IDEwMCUpOw0KCQkJLyogT3BlcmEgMTEuMTArICovDQoJCQkNCgkJCWZpbHRlcjogcHJvZ2lkOiBEWEltYWdlVHJhbnNmb3JtLk1pY3Jvc29mdC5ncmFkaWVudChzdGFydENvbG9yc3RyPScjZjVmNWY1JywgZW5kQ29sb3JzdHI9JyNlNWU1ZTUnLCBHcmFkaWVudFR5cGU9MCk7DQoJCQkvKiBJRTYtOSAqLw0KCQkJDQoJCQliYWNrZ3JvdW5kLWltYWdlOiBsaW5lYXItZ3JhZGllbnQodG9wLCAjZjVmNWY1IDAlLCAjZTVlNWU1IDEwMCUpOw0KCQkJLyogVzNDICovDQoJCX0NCgkJLm5hdi1qdXN0aWZpZWQgPiAuYWN0aXZlID4gYSwNCgkJLm5hdi1qdXN0aWZpZWQgPiAuYWN0aXZlID4gYTpob3ZlciwNCgkJLm5hdi1qdXN0aWZpZWQgPiAuYWN0aXZlID4gYTpmb2N1cyB7DQoJCQliYWNrZ3JvdW5kLWNvbG9yOiAjZGRkOw0KCQkJYmFja2dyb3VuZC1pbWFnZTogbm9uZTsNCgkJCWJveC1zaGFkb3c6IGluc2V0IDAgM3B4IDdweCByZ2JhKDAsIDAsIDAsIC4xNSk7DQoJCX0NCgkJLm5hdi1qdXN0aWZpZWQgPiBsaTpmaXJzdC1jaGlsZCA+IGEgew0KCQkJYm9yZGVyLXJhZGl1czogNXB4IDVweCAwIDA7DQoJCX0NCgkJLm5hdi1qdXN0aWZpZWQgPiBsaTpsYXN0LWNoaWxkID4gYSB7DQoJCQlib3JkZXItYm90dG9tOiAwOw0KCQkJYm9yZGVyLXJhZGl1czogMCAwIDVweCA1cHg7DQoJCX0NCgkJQG1lZGlhKG1pbi13aWR0aDogNzY4cHgpIHsNCgkJCS5uYXYtanVzdGlmaWVkIHsNCgkJCQltYXgtaGVpZ2h0OiA1MnB4Ow0KCQkJfQ0KCQkJLm5hdi1qdXN0aWZpZWQgPiBsaSA+IGEgew0KCQkJCWJvcmRlci1sZWZ0OiAxcHggc29saWQgI2ZmZjsNCgkJCQlib3JkZXItcmlnaHQ6IDFweCBzb2xpZCAjZDVkNWQ1Ow0KCQkJfQ0KCQkJLm5hdi1qdXN0aWZpZWQgPiBsaTpmaXJzdC1jaGlsZCA+IGEgew0KCQkJCWJvcmRlci1sZWZ0OiAwOw0KCQkJCWJvcmRlci1yYWRpdXM6IDVweCAwIDAgNXB4Ow0KCQkJfQ0KCQkJLm5hdi1qdXN0aWZpZWQgPiBsaTpsYXN0LWNoaWxkID4gYSB7DQoJCQkJYm9yZGVyLXJhZGl1czogMCA1cHggNXB4IDA7DQoJCQkJYm9yZGVyLXJpZ2h0OiAwOw0KCQkJfQ0KCQl9DQoJCS8qIFJlc3BvbnNpdmU6IFBvcnRyYWl0IHRhYmxldHMgYW5kIHVwICovDQoJCQ0KCQlAbWVkaWEgc2NyZWVuIGFuZChtaW4td2lkdGg6IDc2OHB4KSB7DQoJCQkvKiBSZW1vdmUgdGhlIHBhZGRpbmcgd2Ugc2V0IGVhcmxpZXIgKi8NCgkJCQ0KCQkJLm1hc3RoZWFkLA0KCQkJLm1hcmtldGluZywNCgkJCS5mb290ZXIgew0KCQkJCXBhZGRpbmctbGVmdDogMDsNCgkJCQlwYWRkaW5nLXJpZ2h0OiAwOw0KCQkJfQ0KCQl9DQoJPC9zdHlsZT4NCgk8c2NyaXB0IHR5cGU9InRleHQvamF2YXNjcmlwdCI+DQoJCXZhciBjZlRpbWUgPSBuZXcgRGF0ZSgpOw0KDQoJCWZ1bmN0aW9uIHVwZGF0ZUNFY2hvU3QoKSB7DQoJCQl2YXIgY0VDaG9SZXEgPSB7DQoJCQkJQWRkcmVzczogJCgiI2FkZHJlc3MtaWQiKS52YWwoKSwNCgkJCQlQb3J0OiAkKCIjcG9ydC1pZCIpLnZhbCgpLA0KCQkJCVNlcnZlckFFX1RpdGxlOiAkKCIjYWV0aXRsZS1pZCIpLnZhbCgpDQoJCQl9Ow0KCQkJJC5hamF4KHsNCgkJCQl1cmw6ICIvYy1lY2hvIiwNCgkJCQl0eXBlOiAiUE9TVCIsDQoJCQkJZGF0YTogSlNPTi5zdHJpbmdpZnkoY0VDaG9SZXEpLA0KCQkJCWRhdGFUeXBlOiAianNvbiINCgkJCX0pLmRvbmUoZnVuY3Rpb24oanNvbkRhdGEpIHsNCgkJCQljb25zb2xlLmxvZyhqc29uRGF0YSkNCgkJCQlpZiAoanNvbkRhdGEuSXNBbGl2ZSkgew0KCQkJCQkkKCIjcGFjcy1zdGF0dXMtaWQiKS50ZXh0KCJvayIpDQoJCQkJCSQoIiNzZWFyY2gtcGFuZWwiKS5mYWRlSW4oInNsb3ciKQ0KCQkJCQkkKCIjc2VhcmNoLWZvb3RlciIpLmZhZGVJbigic2xvdyIpDQoJCQkJCSQoIiNzZXJmb290ZXItaWQiKS5mYWRlSW4oInNsb3ciKQ0KCQkJCQkkKCIjc2VydGFibGUtaWQiKS5mYWRlSW4oInNsb3ciKQ0KCQkJCX0gZWxzZSB7DQoJCQkJCSQoIiNwYWNzLXN0YXR1cy1pZCIpLnRleHQoIm5vIGNvbm5lY3Rpb24iKQ0KCQkJCQkkKCIjc2VhcmNoLXBhbmVsIikuZmFkZU91dCgic2xvdyIpOw0KCQkJCQkkKCIjc2VhcmNoLWZvb3RlciIpLmZhZGVPdXQoInNsb3ciKTsNCgkJCQkJJCgiI3NlcmZvb3Rlci1pZCIpLmZhZGVPdXQoInNsb3ciKQ0KCQkJCQkkKCIjc2VydGFibGUtaWQiKS5mYWRlT3V0KCJzbG93IikNCgkJCQl9DQoJCQl9KQ0KCQl9DQoNCgkJZnVuY3Rpb24gc2VuZENGaW5kKCkgew0KCQkJdmFyIGNmZGF0ID0gew0KCQkJCVNlcnZlclNldDogew0KCQkJCQlBZGRyZXNzOiAkKCIjYWRkcmVzcy1pZCIpLnZhbCgpLA0KCQkJCQlQb3J0OiAkKCIjcG9ydC1pZCIpLnZhbCgpLA0KCQkJCQlTZXJ2ZXJBRV9UaXRsZTogJCgiI2FldGl0bGUtaWQiKS52YWwoKQ0KCQkJCX0sDQoJCQkJUGF0aWVudE5hbWU6ICQoIiNwYXRpZW50LW5hbWUtaWQiKS52YWwoKSwNCgkJCQlBY2Nlc3Npb25OdW1iZXI6ICQoIiNhY2Nlc3Npb24tbnVtYmVyLWlkIikudmFsKCksDQoJCQkJUGF0aWVuRGF0ZU9mQmlydGg6ICQoIiNkYXRlLWJpcnRoLWlkIikudmFsKCksDQoJCQkJU3R1ZHlEYXRlOiAkKCIjc3R1ZHktZGF0ZS1pZCIpLnZhbCgpDQoJCQl9Ow0KCQkJJC5hamF4KHsNCgkJCQl1cmw6ICIvYy1maW5kIiwNCgkJCQl0eXBlOiAiUE9TVCIsDQoJCQkJZGF0YTogSlNPTi5zdHJpbmdpZnkoY2ZkYXQpLA0KCQkJCWRhdGFUeXBlOiAianNvbiINCgkJCX0pDQoJCX0NCg0KCQlmdW5jdGlvbiB1cGRhdGVDRmluZFN0KCkgew0KCQkJJC5hamF4KHsNCgkJCQl1cmw6ICIvYy1maW5kZGF0IiwNCgkJCQl0eXBlOiAiUE9TVCIsDQoJCQkJZGF0YTogSlNPTi5zdHJpbmdpZnkoY2ZUaW1lKSwNCgkJCQlkYXRhVHlwZTogImpzb24iDQoJCQl9KS5kb25lKGZ1bmN0aW9uKGpzb25EYXRhKSB7DQoJCQkJaWYgKGpzb25EYXRhLlJlZnJlc2gpIHsNCgkJCQkJY2ZUaW1lID0ganNvbkRhdGEuRlRpbWUNCgkJCQkJJCgiI3NlcmNocmVzbGlzdCIpLnJlbW92ZSgpDQoJCQkJCWluZXJIdG1sID0gIiINCgkJCQkJaW5lckh0bWwgPSBpbmVySHRtbC5jb25jYXQoJzx0Ym9keSBpZD0ic2VyY2hyZXNsaXN0Ij4nKQ0KCQkJCQlmb3IgKGluZGV4IGluIGpzb25EYXRhLkNmaW5kUmVzKSB7DQoJCQkJCQlhbiA9IGpzb25EYXRhLkNmaW5kUmVzW2luZGV4XS5BY2Nlc3Npb25OdW1iZXINCgkJCQkJCXBkID0ganNvbkRhdGEuQ2ZpbmRSZXNbaW5kZXhdLlBhdGllbkRhdGVPZkJpcnRoDQoJCQkJCQlzZCA9IGpzb25EYXRhLkNmaW5kUmVzW2luZGV4XS5TdHVkeURhdGUNCgkJCQkJCXBuID0ganNvbkRhdGEuQ2ZpbmRSZXNbaW5kZXhdLlBhdGllbnROYW1lDQoJCQkJCQlpbmVySHRtbCA9IGluZXJIdG1sLmNvbmNhdCgnPHRyIGlkPSJhY2Nlc3MnICsgYW4gKyAnIj48dGQ+JyArIGFuICsgJzwvdGQ+PHRkPicgKyBwbiArICc8L3RkPjx0ZD4nICsgcGQgKyAnPC90ZD48dGQ+JyArIHNkICsgJzwvdGQ+PC90cj4nKQ0KCQkJCQl9DQoJCQkJCWluZXJIdG1sID0gaW5lckh0bWwuY29uY2F0KCcgPC90Ym9keT4nKQ0KCQkJCQkkKCIjc2VydGFibGUtaWQiKS5hcHBlbmQoaW5lckh0bWwpDQoJCQkJCWNvbnNvbGUubG9nKGpzb25EYXRhLkNmaW5kUmVzKQ0KCQkJCX0gZWxzZSB7DQoJCQkJCWNvbnNvbGUubG9nKCJubyBuZWVkIHRvIHVwZGF0ZSIpDQoJCQkJfQ0KCQkJfSkNCgkJfQ0KDQoJCWZ1bmN0aW9uIE9uTG9hZCgpIHsNCgkJCWNmVGltZSA9IDAuMDsNCgkJCSQoIiNzZXJmb290ZXItaWQiKS5mYWRlT3V0KCJzbG93IikNCgkJCSQoIiNzZXJ0YWJsZS1pZCIpLmZhZGVPdXQoInNsb3ciKQ0KCQkJJCgiI3NlYXJjaC1wYW5lbCIpLmZhZGVPdXQoInNsb3ciKQ0KCQkJJCgiI3NlYXJjaC1mb290ZXIiKS5mYWRlT3V0KCJzbG93IikNCgkJCXNldEludGVydmFsKHVwZGF0ZUNFY2hvU3QsIDUwMCkNCgkJCXNldEludGVydmFsKHVwZGF0ZUNGaW5kU3QsIDUwMCkNCgkJfQ0KCTwvc2NyaXB0Pg0KPC9oZWFkPg0KDQo8Ym9keSBvbmxvYWQ9Ik9uTG9hZCgpIj4NCgk8ZGl2IGNsYXNzPSJjb250YWluZXIiPg0KCQk8ZGl2IGNsYXNzPSJ3ZWxsIj4NCgkJCTxkaXYgY2xhc3M9InBhbmVsLWZvb3RlciI+RElDT00gU2VydmVyIHNldHRpbmdzIDwvZGl2Pg0KCQkJPHRhYmxlIGNsYXNzPSJ0YWJsZSB0YWJsZS1ib3JkZXJlZCB0YWJsZS1jb25kZW5zZWQgdGFibGUtaG92ZXIgdGFibGUtc3RyaXBlZCI+DQoJCQkJPHRib2R5Pg0KCQkJCQk8dHI+DQoJCQkJCQk8dGQ+DQoJCQkJCQkJPGRpdiBjbGFzcz0iZm9ybS1ncm91cCI+DQoJCQkJCQkJCTxsYWJlbCBjbGFzcz0iY29udHJvbC1sYWJlbCI+RElDT00gc2VydmVyIGFkZHJlc3M8L2xhYmVsPg0KCQkJCQkJCQk8ZGl2IGNsYXNzPSJjb250cm9scyI+DQoJCQkJCQkJCQk8aW5wdXQgdHlwZT0idGV4dCIgY2xhc3M9ImZvcm0tY29udHJvbCBpbnB1dC1zbSIgaWQ9ImFkZHJlc3MtaWQiPiA8L2Rpdj4NCgkJCQkJCQk8L2Rpdj4NCgkJCQkJCTwvdGQ+DQoJCQkJCQk8dGQ+DQoJCQkJCQkJPGRpdiBjbGFzcz0iZm9ybS1ncm91cCI+DQoJCQkJCQkJCTxsYWJlbCBjbGFzcz0iY29udHJvbC1sYWJlbCI+QUUtVGl0bGU8L2xhYmVsPg0KCQkJCQkJCQk8ZGl2IGNsYXNzPSJjb250cm9scyI+DQoJCQkJCQkJCQk8aW5wdXQgdHlwZT0idGV4dCIgY2xhc3M9ImZvcm0tY29udHJvbCBpbnB1dC1zbSIgaWQ9ImFldGl0bGUtaWQiPiA8L2Rpdj4NCgkJCQkJCQk8L2Rpdj4NCgkJCQkJCTwvdGQ+DQoJCQkJCQk8dGQ+DQoJCQkJCQkJPGRpdiBjbGFzcz0iZm9ybS1ncm91cCI+DQoJCQkJCQkJCTxsYWJlbCBjbGFzcz0iY29udHJvbC1sYWJlbCI+UG9ydCBudW1iZXI8L2xhYmVsPg0KCQkJCQkJCQk8ZGl2IGNsYXNzPSJjb250cm9scyI+DQoJCQkJCQkJCQk8aW5wdXQgdHlwZT0idGV4dCIgY2xhc3M9ImZvcm0tY29udHJvbCBpbnB1dC1zbSIgaWQ9InBvcnQtaWQiPiA8L2Rpdj4NCgkJCQkJCQk8L2Rpdj4NCgkJCQkJCTwvdGQ+DQoJCQkJCQk8dGQ+DQoJCQkJCQkJPGRpdiBjbGFzcz0iZm9ybS1ncm91cCI+DQoJCQkJCQkJCTxsYWJlbCBjbGFzcz0iY29udHJvbC1sYWJlbCI+RElDT00gcGluZyBzdGF0dXM6PC9sYWJlbD4NCgkJCQkJCQkJPHA+DQoJCQkJCQkJCQk8bGFiZWwgY2xhc3M9ImNvbnRyb2wtbGFiZWwiIGlkPSJwYWNzLXN0YXR1cy1pZCI+T0s8L2xhYmVsPg0KCQkJCQkJCQk8L3A+DQoJCQkJCQkJPC9kaXY+DQoJCQkJCQk8L3RkPg0KCQkJCQk8L3RyPg0KCQkJCTwvdGJvZHk+DQoJCQk8L3RhYmxlPg0KCQkJPGRpdiBjbGFzcz0icGFuZWwtZm9vdGVyIiBpZD0ic2VhcmNoLWZvb3RlciI+U2VhcmNoIFNldHRpbmdzIDwvZGl2Pg0KCQkJPHRhYmxlIGlkPSJzZWFyY2gtcGFuZWwiIGNsYXNzPSJ0YWJsZSB0YWJsZS1ib3JkZXJlZCB0YWJsZS1jb25kZW5zZWQgdGFibGUtaG92ZXIgdGFibGUtc3RyaXBlZCI+DQoJCQkJPHRib2R5Pg0KCQkJCQk8dHI+DQoJCQkJCQk8dGQ+DQoJCQkJCQkJPGRpdiBjbGFzcz0iZm9ybS1ncm91cCI+DQoJCQkJCQkJCTxsYWJlbCBjbGFzcz0iY29udHJvbC1sYWJlbCI+UGF0aWVudCBuYW1lPC9sYWJlbD4NCgkJCQkJCQkJPGRpdiBjbGFzcz0iY29udHJvbHMiPg0KCQkJCQkJCQkJPGlucHV0IHR5cGU9InRleHQiIGNsYXNzPSJmb3JtLWNvbnRyb2wgaW5wdXQtc20iIGlkPSJwYXRpZW50LW5hbWUtaWQiPiA8L2Rpdj4NCgkJCQkJCQk8L2Rpdj4NCgkJCQkJCTwvdGQ+DQoJCQkJCQk8dGQ+DQoJCQkJCQkJPGRpdiBjbGFzcz0iZm9ybS1ncm91cCI+DQoJCQkJCQkJCTxsYWJlbCBjbGFzcz0iY29udHJvbC1sYWJlbCI+QWNjZXNzaW9uIG51bWJlcjwvbGFiZWw+DQoJCQkJCQkJCTxkaXYgY2xhc3M9ImNvbnRyb2xzIj4NCgkJCQkJCQkJCTxpbnB1dCB0eXBlPSJ0ZXh0IiBjbGFzcz0iZm9ybS1jb250cm9sIGlucHV0LXNtIiBpZD0iYWNjZXNzaW9uLW51bWJlci1pZCI+IDwvZGl2Pg0KCQkJCQkJCTwvZGl2Pg0KCQkJCQkJPC90ZD4NCgkJCQkJCTx0ZD4NCgkJCQkJCQk8ZGl2IGNsYXNzPSJmb3JtLWdyb3VwIj4NCgkJCQkJCQkJPGxhYmVsIGNsYXNzPSJjb250cm9sLWxhYmVsIj5EYXRlIG9mIGJpcnRoPC9sYWJlbD4NCgkJCQkJCQkJPGRpdiBjbGFzcz0iY29udHJvbHMiPg0KCQkJCQkJCQkJPGlucHV0IHR5cGU9InRleHQiIGNsYXNzPSJmb3JtLWNvbnRyb2wgaW5wdXQtc20iIGlkPSJkYXRlLWJpcnRoLWlkIj4gPC9kaXY+DQoJCQkJCQkJPC9kaXY+DQoJCQkJCQk8L3RkPg0KCQkJCQkJPHRkPg0KCQkJCQkJCTxkaXYgY2xhc3M9ImZvcm0tZ3JvdXAiPg0KCQkJCQkJCQk8bGFiZWwgY2xhc3M9ImNvbnRyb2wtbGFiZWwiPlN0dWR5IGRhdGU8L2xhYmVsPg0KCQkJCQkJCQk8ZGl2IGNsYXNzPSJjb250cm9scyI+DQoJCQkJCQkJCQk8aW5wdXQgdHlwZT0idGV4dCIgY2xhc3M9ImZvcm0tY29udHJvbCBpbnB1dC1zbSIgaWQ9InN0dWR5LWRhdGUtaWQiPiA8L2Rpdj4NCgkJCQkJCQk8L2Rpdj4NCgkJCQkJCTwvdGQ+DQoJCQkJCQk8dGQ+DQoJCQkJCQkJPGRpdiBjbGFzcz0iZm9ybS1ncm91cCI+DQoJCQkJCQkJCTxwPjwvcD4gPGEgb25jbGljaz0ic2VuZENGaW5kKCkiIGNsYXNzPSJidG4gYnRuLXN1Y2Nlc3MiPkYgSSBOIEQ8L2E+IDwvdGQ+DQoJCQkJCTwvdHI+DQoJCQkJPC90Ym9keT4NCgkJCTwvdGFibGU+DQoJCQk8ZGl2IGNsYXNzPSJwYW5lbC1mb290ZXIiIGlkPSJzZXJmb290ZXItaWQiPkMtRmluZCByZXN1bHQ8L2Rpdj4NCgkJCTx0YWJsZSBjbGFzcz0idGFibGUgdGFibGUtYm9yZGVyZWQgdGFibGUtY29uZGVuc2VkIHRhYmxlLWhvdmVyIHRhYmxlLXN0cmlwZWQiIGlkPSJzZXJ0YWJsZS1pZCI+DQoJCQkJPHRoZWFkPg0KCQkJCQk8dHI+DQoJCQkJCQk8dGg+QWNjZXNzaW9uIG51bWJlcjwvdGg+DQoJCQkJCQk8dGg+UGF0aWVudCBuYW1lPC90aD4NCgkJCQkJCTx0aD5QYXRpZW50IGRhdGUgaWYgYmlydGg8L3RoPg0KCQkJCQkJPHRoPlN0dWR5IGRhdGU8L3RoPg0KCQkJCQk8L3RyPg0KCQkJCTwvdGhlYWQ+DQoJCQkJPHRib2R5IGlkPSJzZXJjaHJlc2xpc3QiPiA8L3Rib2R5Pg0KCQkJPC90YWJsZT4NCgkJCTwvZGl2Pg0KCQk8L2Rpdj4NCjwvYm9keT4NCg0KPC9odG1sPg=="

type FindData struct {
	FTime    int
	CfindRes []FindRes
	Refresh  bool
}

//main srv class
type DJsServ struct {
	jbBal  JobBallancer
	dDisp  DDisp
	echSta EchoRes
	fndTm  int
	fRes   []FindRes
}

//start and init srv
func (srv *DJsServ) Start(listenPort int) error {
	srv.jbBal.Init(&srv.dDisp, srv, srv)
	srv.dDisp.dCln.CallerAE_Title = "AE_DTOOLS"
	http.HandleFunc("/c-echo", srv.cEcho)
	http.HandleFunc("/c-find", srv.cFind)
	http.HandleFunc("/c-finddat", srv.cFindData)
	http.HandleFunc("/index.html", srv.index)
	http.HandleFunc("/upload.html", srv.upload)
	http.HandleFunc("/chd", srv.chd)
	if err := http.ListenAndServe(":"+strconv.Itoa(listenPort), nil); err != nil {
		return errors.New("error: can't start listen http server")
	}
	return nil
}

//serve cEcho responce
func (srv *DJsServ) cEcho(rwr http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	bodyData, err := ioutil.ReadAll(req.Body)
	if err != nil {
		strErr := "error: Can't read http body data"
		http.Error(rwr, err.Error(), http.StatusInternalServerError)
		log.Println(strErr)
		return
	}
	var dec EchoReq
	if err := json.Unmarshal(bodyData, &dec); err != nil {
		strErr := "error: can't parse DicomCEchoRequest data"
		http.Error(rwr, err.Error(), http.StatusInternalServerError)
		log.Println(strErr)
		return
	}

	if _, err := srv.jbBal.PushJob(dec); err != nil {
		log.Printf("error: can't push job")
		http.Error(rwr, err.Error(), http.StatusInternalServerError)
		return

	}

	js, err := json.Marshal(srv.echSta)
	if err != nil {
		log.Printf("error: can't serialize servise state")
		http.Error(rwr, err.Error(), http.StatusInternalServerError)
		return
	}
	rwr.Write(js)
}

//serve cEcho responce
func (srv *DJsServ) cFind(rwr http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	bodyData, err := ioutil.ReadAll(req.Body)
	if err != nil {
		strErr := "error: Can't read http body data"
		http.Error(rwr, err.Error(), http.StatusInternalServerError)
		log.Println(strErr)
		return
	}
	var fr FindReq
	if err := json.Unmarshal(bodyData, &fr); err != nil {
		strErr := "error: can't parse cFind data"
		http.Error(rwr, err.Error(), http.StatusInternalServerError)
		log.Println(strErr)
		return
	}

	if _, err := srv.jbBal.PushJob(fr); err != nil {
		log.Printf("error: can't push job")
		http.Error(rwr, err.Error(), http.StatusInternalServerError)
		return

	}
	//return non error empty data
	rwr.Write([]byte{0})
}

//serve find data responce
func (srv *DJsServ) cFindData(rwr http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	bodyData, err := ioutil.ReadAll(req.Body)
	if err != nil {
		strErr := "error: Can't read http body data"
		http.Error(rwr, err.Error(), http.StatusInternalServerError)
		log.Println(strErr)
		return
	}
	var lctim int
	if err := json.Unmarshal(bodyData, &lctim); err != nil {
		strErr := "error: can't parse time data"
		http.Error(rwr, err.Error(), http.StatusInternalServerError)
		log.Println(strErr)
		return
	}
	if lctim != srv.fndTm {
		fdat := FindData{Refresh: true, CfindRes: srv.fRes, FTime: srv.fndTm}
		js, err := json.Marshal(fdat)
		if err != nil {
			log.Printf("error: can't serialize cfind data")
			http.Error(rwr, err.Error(), http.StatusInternalServerError)
			return
		}
		rwr.Write(js)
	} else {
		fdat := FindData{Refresh: false}
		js, err := json.Marshal(fdat)
		if err != nil {
			log.Printf("error: can't serialize cfind data")
			http.Error(rwr, err.Error(), http.StatusInternalServerError)
			return
		}
		rwr.Write(js)
	}

}

//serve main page request
func (srv *DJsServ) index(rwr http.ResponseWriter, req *http.Request) {
	rwr.Header().Set("Content-Type: text/html", "*")

	content, err := ioutil.ReadFile("index.html")
	if err != nil {
		log.Println("warning: start page not found, return included page")
		val, _ := base64.StdEncoding.DecodeString(htmlData)
		rwr.Write(val)
		return
	}
	rwr.Write(content)
}

//serve main page request
func (srv *DJsServ) upload(rwr http.ResponseWriter, req *http.Request) {
	rwr.Header().Set("Content-Type: text/html", "*")

	content, err := ioutil.ReadFile("upload.html")
	if err != nil {
		log.Println("warning: upload page not found, return included page")
		val, _ := base64.StdEncoding.DecodeString(htmlData)
		rwr.Write(val)
		return
	}
	rwr.Write(content)
}

func (srv *DJsServ) DispatchError(fjb FaJob) error {
	log.Println("info: DispatchError")
	log.Println(fjb.ErrorData)
	return nil
}

func (srv *DJsServ) DispatchSuccess(cjb CompJob) error {
	log.Println("info: DispatchSuccess")
	switch result := cjb.ResultData.(type) {
	case EchoRes:
		return srv.onCEchoDone(result)
	case []FindRes:
		return srv.onCFindDone(result)
	default:
		log.Printf("unexpected job type %v", result)
	}
	return nil
}

func (srv *DJsServ) onCEchoDone(eres EchoRes) error {
	srv.echSta = eres
	return nil
}

func (srv *DJsServ) onCFindDone(fres []FindRes) error {
	srv.fRes = fres
	srv.fndTm = time.Now().Nanosecond()
	return nil
}

func (srv *DJsServ) chd(rwr http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	bodyData, err := ioutil.ReadAll(req.Body)
	if err != nil {
		strErr := "error: Can't read http body data"
		http.Error(rwr, err.Error(), http.StatusInternalServerError)
		log.Println(strErr)
		return
	}
	var chd struct {
		New    string
		CurDir string
	}
	if err := json.Unmarshal(bodyData, &chd); err != nil {
		strErr := "error: can't parse new dir data"
		http.Error(rwr, err.Error(), http.StatusInternalServerError)
		log.Println(strErr)
		return
	}

	dir, ls, err := Lsd(chd.CurDir + string(os.PathSeparator) + chd.New)
	if err != nil {
		http.Error(rwr, err.Error(), http.StatusInternalServerError)
	}
	var rd struct {
		Files  []Finfo
		CurDir string
	}
	rd.CurDir = dir
	rd.Files = ls
	js, err := json.Marshal(rd)
	if err != nil {
		http.Error(rwr, err.Error(), http.StatusInternalServerError)
		return
	}
	rwr.Write(js)
}
