package main

import "io/ioutil"
import "net/http"
import "strconv"
import "log"
import "encoding/json"
import "errors"
import "time"

import "encoding/base64"

const htmlData = "PCFkb2N0eXBlIGh0bWw+DQo8aHRtbD4NCg0KPGhlYWQ+DQoJPHRpdGxlPkp1c3RpZmllZCBOYXY8L3RpdGxlPg0KCTxtZXRhIG5hbWU9InZpZXdwb3J0IiBjb250ZW50PSJ3aWR0aD1kZXZpY2Utd2lkdGgiPg0KCTxsaW5rIHJlbD0ic3R5bGVzaGVldCIgaHJlZj0iaHR0cHM6Ly9uZXRkbmEuYm9vdHN0cmFwY2RuLmNvbS9ib290c3dhdGNoLzMuMC4wL3NsYXRlL2Jvb3RzdHJhcC5taW4uY3NzIj4NCgk8c2NyaXB0IHR5cGU9InRleHQvamF2YXNjcmlwdCIgc3JjPSJodHRwczovL2FqYXguZ29vZ2xlYXBpcy5jb20vYWpheC9saWJzL2pxdWVyeS8yLjAuMy9qcXVlcnkubWluLmpzIj48L3NjcmlwdD4NCgk8c2NyaXB0IHR5cGU9InRleHQvamF2YXNjcmlwdCIgc3JjPSJodHRwczovL25ldGRuYS5ib290c3RyYXBjZG4uY29tL2Jvb3RzdHJhcC8zLjEuMS9qcy9ib290c3RyYXAubWluLmpzIj48L3NjcmlwdD4NCgk8c3R5bGUgdHlwZT0idGV4dC9jc3MiPg0KCQlib2R5IHsNCgkJCXBhZGRpbmctdG9wOiAyMHB4Ow0KCQl9DQoJCS5mb290ZXIgew0KCQkJYm9yZGVyLXRvcDogMXB4IHNvbGlkICNlZWU7DQoJCQltYXJnaW4tdG9wOiA0MHB4Ow0KCQkJcGFkZGluZy10b3A6IDQwcHg7DQoJCQlwYWRkaW5nLWJvdHRvbTogNDBweDsNCgkJfQ0KCQkvKiBNYWluIG1hcmtldGluZyBtZXNzYWdlIGFuZCBzaWduIHVwIGJ1dHRvbiAqLw0KCQkNCgkJLmp1bWJvdHJvbiB7DQoJCQl0ZXh0LWFsaWduOiBjZW50ZXI7DQoJCQliYWNrZ3JvdW5kLWNvbG9yOiB0cmFuc3BhcmVudDsNCgkJfQ0KCQkuanVtYm90cm9uIC5idG4gew0KCQkJZm9udC1zaXplOiAyMXB4Ow0KCQkJcGFkZGluZzogMTRweCAyNHB4Ow0KCQl9DQoJCS8qIEN1c3RvbWl6ZSB0aGUgbmF2LWp1c3RpZmllZCBsaW5rcyB0byBiZSBmaWxsIHRoZSBlbnRpcmUgc3BhY2Ugb2YgdGhlIC5uYXZiYXIgKi8NCgkJDQoJCS5uYXYtanVzdGlmaWVkIHsNCgkJCWJhY2tncm91bmQtY29sb3I6ICNlZWU7DQoJCQlib3JkZXItcmFkaXVzOiA1cHg7DQoJCQlib3JkZXI6IDFweCBzb2xpZCAjY2NjOw0KCQl9DQoJCS5uYXYtanVzdGlmaWVkID4gbGkgPiBhIHsNCgkJCXBhZGRpbmctdG9wOiAxNXB4Ow0KCQkJcGFkZGluZy1ib3R0b206IDE1cHg7DQoJCQljb2xvcjogIzc3NzsNCgkJCWZvbnQtd2VpZ2h0OiBib2xkOw0KCQkJdGV4dC1hbGlnbjogY2VudGVyOw0KCQkJYm9yZGVyLWJvdHRvbTogMXB4IHNvbGlkICNkNWQ1ZDU7DQoJCQliYWNrZ3JvdW5kLWNvbG9yOiAjZTVlNWU1Ow0KCQkJLyogT2xkIGJyb3dzZXJzICovDQoJCQkNCgkJCWJhY2tncm91bmQtcmVwZWF0OiByZXBlYXQteDsNCgkJCS8qIFJlcGVhdCB0aGUgZ3JhZGllbnQgKi8NCgkJCQ0KCQkJYmFja2dyb3VuZC1pbWFnZTogLW1vei1saW5lYXItZ3JhZGllbnQodG9wLCAjZjVmNWY1IDAlLCAjZTVlNWU1IDEwMCUpOw0KCQkJLyogRkYzLjYrICovDQoJCQkNCgkJCWJhY2tncm91bmQtaW1hZ2U6IC13ZWJraXQtZ3JhZGllbnQobGluZWFyLCBsZWZ0IHRvcCwgbGVmdCBib3R0b20sIGNvbG9yLXN0b3AoMCUsICNmNWY1ZjUpLCBjb2xvci1zdG9wKDEwMCUsICNlNWU1ZTUpKTsNCgkJCS8qIENocm9tZSxTYWZhcmk0KyAqLw0KCQkJDQoJCQliYWNrZ3JvdW5kLWltYWdlOiAtd2Via2l0LWxpbmVhci1ncmFkaWVudCh0b3AsICNmNWY1ZjUgMCUsICNlNWU1ZTUgMTAwJSk7DQoJCQkvKiBDaHJvbWUgMTArLFNhZmFyaSA1LjErICovDQoJCQkNCgkJCWJhY2tncm91bmQtaW1hZ2U6IC1tcy1saW5lYXItZ3JhZGllbnQodG9wLCAjZjVmNWY1IDAlLCAjZTVlNWU1IDEwMCUpOw0KCQkJLyogSUUxMCsgKi8NCgkJCQ0KCQkJYmFja2dyb3VuZC1pbWFnZTogLW8tbGluZWFyLWdyYWRpZW50KHRvcCwgI2Y1ZjVmNSAwJSwgI2U1ZTVlNSAxMDAlKTsNCgkJCS8qIE9wZXJhIDExLjEwKyAqLw0KCQkJDQoJCQlmaWx0ZXI6IHByb2dpZDogRFhJbWFnZVRyYW5zZm9ybS5NaWNyb3NvZnQuZ3JhZGllbnQoc3RhcnRDb2xvcnN0cj0nI2Y1ZjVmNScsIGVuZENvbG9yc3RyPScjZTVlNWU1JywgR3JhZGllbnRUeXBlPTApOw0KCQkJLyogSUU2LTkgKi8NCgkJCQ0KCQkJYmFja2dyb3VuZC1pbWFnZTogbGluZWFyLWdyYWRpZW50KHRvcCwgI2Y1ZjVmNSAwJSwgI2U1ZTVlNSAxMDAlKTsNCgkJCS8qIFczQyAqLw0KCQl9DQoJCS5uYXYtanVzdGlmaWVkID4gLmFjdGl2ZSA+IGEsDQoJCS5uYXYtanVzdGlmaWVkID4gLmFjdGl2ZSA+IGE6aG92ZXIsDQoJCS5uYXYtanVzdGlmaWVkID4gLmFjdGl2ZSA+IGE6Zm9jdXMgew0KCQkJYmFja2dyb3VuZC1jb2xvcjogI2RkZDsNCgkJCWJhY2tncm91bmQtaW1hZ2U6IG5vbmU7DQoJCQlib3gtc2hhZG93OiBpbnNldCAwIDNweCA3cHggcmdiYSgwLCAwLCAwLCAuMTUpOw0KCQl9DQoJCS5uYXYtanVzdGlmaWVkID4gbGk6Zmlyc3QtY2hpbGQgPiBhIHsNCgkJCWJvcmRlci1yYWRpdXM6IDVweCA1cHggMCAwOw0KCQl9DQoJCS5uYXYtanVzdGlmaWVkID4gbGk6bGFzdC1jaGlsZCA+IGEgew0KCQkJYm9yZGVyLWJvdHRvbTogMDsNCgkJCWJvcmRlci1yYWRpdXM6IDAgMCA1cHggNXB4Ow0KCQl9DQoJCUBtZWRpYShtaW4td2lkdGg6IDc2OHB4KSB7DQoJCQkubmF2LWp1c3RpZmllZCB7DQoJCQkJbWF4LWhlaWdodDogNTJweDsNCgkJCX0NCgkJCS5uYXYtanVzdGlmaWVkID4gbGkgPiBhIHsNCgkJCQlib3JkZXItbGVmdDogMXB4IHNvbGlkICNmZmY7DQoJCQkJYm9yZGVyLXJpZ2h0OiAxcHggc29saWQgI2Q1ZDVkNTsNCgkJCX0NCgkJCS5uYXYtanVzdGlmaWVkID4gbGk6Zmlyc3QtY2hpbGQgPiBhIHsNCgkJCQlib3JkZXItbGVmdDogMDsNCgkJCQlib3JkZXItcmFkaXVzOiA1cHggMCAwIDVweDsNCgkJCX0NCgkJCS5uYXYtanVzdGlmaWVkID4gbGk6bGFzdC1jaGlsZCA+IGEgew0KCQkJCWJvcmRlci1yYWRpdXM6IDAgNXB4IDVweCAwOw0KCQkJCWJvcmRlci1yaWdodDogMDsNCgkJCX0NCgkJfQ0KCQkvKiBSZXNwb25zaXZlOiBQb3J0cmFpdCB0YWJsZXRzIGFuZCB1cCAqLw0KCQkNCgkJQG1lZGlhIHNjcmVlbiBhbmQobWluLXdpZHRoOiA3NjhweCkgew0KCQkJLyogUmVtb3ZlIHRoZSBwYWRkaW5nIHdlIHNldCBlYXJsaWVyICovDQoJCQkNCgkJCS5tYXN0aGVhZCwNCgkJCS5tYXJrZXRpbmcsDQoJCQkuZm9vdGVyIHsNCgkJCQlwYWRkaW5nLWxlZnQ6IDA7DQoJCQkJcGFkZGluZy1yaWdodDogMDsNCgkJCX0NCgkJfQ0KCTwvc3R5bGU+DQoJPHNjcmlwdCB0eXBlPSJ0ZXh0L2phdmFzY3JpcHQiPg0KCQlmdW5jdGlvbiB1cGRhdGVDRWNob1N0KCkgew0KCQkJdmFyIGNFQ2hvUmVxID0gew0KCQkJCUFkZHJlc3M6ICQoIiNhZGRyZXNzLWlkIikudmFsKCksDQoJCQkJUG9ydDogJCgiI3BvcnQtaWQiKS52YWwoKSwNCgkJCQlTZXJ2ZXJBRV9UaXRsZTogJCgiI2FldGl0bGUtaWQiKS52YWwoKQ0KCQkJfTsNCgkJCSQuYWpheCh7DQoJCQkJdXJsOiAiL2MtZWNobyIsDQoJCQkJdHlwZTogIlBPU1QiLA0KCQkJCWRhdGE6IEpTT04uc3RyaW5naWZ5KGNFQ2hvUmVxKSwNCgkJCQlkYXRhVHlwZTogImpzb24iDQoJCQl9KS5kb25lKGZ1bmN0aW9uKGpzb25EYXRhKSB7DQoJCQkJY29uc29sZS5sb2coanNvbkRhdGEpDQoJCQkJaWYgKGpzb25EYXRhLklzQWxpdmUpIHsNCgkJCQkJJCgiI3BhY3Mtc3RhdHVzLWlkIikudGV4dCgib2siKQ0KCQkJCQkkKCIjc2VhcmNoLXBhbmVsIikuZmFkZUluKCJzbG93Iik7DQoJCQkJCSQoIiNzZWFyY2gtZm9vdGVyIikuZmFkZUluKCJzbG93Iik7DQoJCQkJfSBlbHNlIHsNCgkJCQkJJCgiI3BhY3Mtc3RhdHVzLWlkIikudGV4dCgibm8gY29ubmVjdGlvbiIpDQoJCQkJCSQoIiNzZWFyY2gtcGFuZWwiKS5mYWRlT3V0KCJzbG93Iik7DQoJCQkJCSQoIiNzZWFyY2gtZm9vdGVyIikuZmFkZU91dCgic2xvdyIpOw0KCQkJCX0NCgkJCX0pDQoJCX0NCg0KCQlmdW5jdGlvbiBzZW5kQ0ZpbmQoKSB7DQoJCQl2YXIgY2ZkYXQgPSB7DQoJCQkJU2VydmVyU2V0OiB7DQoJCQkJCUFkZHJlc3M6ICQoIiNhZGRyZXNzLWlkIikudmFsKCksDQoJCQkJCVBvcnQ6ICQoIiNwb3J0LWlkIikudmFsKCksDQoJCQkJCVNlcnZlckFFX1RpdGxlOiAkKCIjYWV0aXRsZS1pZCIpLnZhbCgpDQoJCQkJfSwNCgkJCQlQYXRpZW50TmFtZTogJCgiI3BhdGllbnQtbmFtZS1pZCIpLnZhbCgpLA0KCQkJCUFjY2Vzc2lvbk51bWJlcjogJCgiI2FjY2Vzc2lvbi1udW1iZXItaWQiKS52YWwoKSwNCgkJCQlQYXRpZW5EYXRlT2ZCaXJ0aDogJCgiI2RhdGUtYmlydGgtaWQiKS52YWwoKSwNCgkJCQlTdHVkeURhdGU6ICQoIiNzdHVkeS1kYXRlLWlkIikudmFsKCkNCgkJCX07DQoJCQkkLmFqYXgoew0KCQkJCXVybDogIi9jLWZpbmQiLA0KCQkJCXR5cGU6ICJQT1NUIiwNCgkJCQlkYXRhOiBKU09OLnN0cmluZ2lmeShjZmRhdCksDQoJCQkJZGF0YVR5cGU6ICJqc29uIg0KCQkJfSkNCgkJfQ0KDQoJCWZ1bmN0aW9uIE9uTG9hZCgpIHsNCgkJCS8vJCgiI3NlYXJjaC1wYW5lbCIpLmZhZGVPdXQoInNsb3ciKTsNCgkJCS8vJCgiI3NlYXJjaC1mb290ZXIiKS5mYWRlT3V0KCJzbG93Iik7DQoJCQkvL3NldEludGVydmFsKHVwZGF0ZUNFY2hvU3QsIDcwMCkNCgkJfQ0KCTwvc2NyaXB0Pg0KPC9oZWFkPg0KDQo8Ym9keSBvbmxvYWQ9Ik9uTG9hZCgpIj4NCgk8ZGl2IGNsYXNzPSJjb250YWluZXIiPg0KCQk8ZGl2IGNsYXNzPSJ3ZWxsIj4NCgkJCTxkaXYgY2xhc3M9InBhbmVsLWZvb3RlciI+RElDT00gU2VydmVyIHNldHRpbmdzIDwvZGl2Pg0KCQkJPHRhYmxlIGNsYXNzPSJ0YWJsZSB0YWJsZS1ib3JkZXJlZCB0YWJsZS1jb25kZW5zZWQgdGFibGUtaG92ZXIgdGFibGUtc3RyaXBlZCI+DQoJCQkJPHRib2R5Pg0KCQkJCQk8dHI+DQoJCQkJCQk8dGQ+DQoJCQkJCQkJPGRpdiBjbGFzcz0iZm9ybS1ncm91cCI+DQoJCQkJCQkJCTxsYWJlbCBjbGFzcz0iY29udHJvbC1sYWJlbCI+RElDT00gc2VydmVyIGFkZHJlc3M8L2xhYmVsPg0KCQkJCQkJCQk8ZGl2IGNsYXNzPSJjb250cm9scyI+DQoJCQkJCQkJCQk8aW5wdXQgdHlwZT0idGV4dCIgY2xhc3M9ImZvcm0tY29udHJvbCBpbnB1dC1zbSIgaWQ9ImFkZHJlc3MtaWQiPiA8L2Rpdj4NCgkJCQkJCQk8L2Rpdj4NCgkJCQkJCTwvdGQ+DQoJCQkJCQk8dGQ+DQoJCQkJCQkJPGRpdiBjbGFzcz0iZm9ybS1ncm91cCI+DQoJCQkJCQkJCTxsYWJlbCBjbGFzcz0iY29udHJvbC1sYWJlbCI+QUUtVGl0bGU8L2xhYmVsPg0KCQkJCQkJCQk8ZGl2IGNsYXNzPSJjb250cm9scyI+DQoJCQkJCQkJCQk8aW5wdXQgdHlwZT0idGV4dCIgY2xhc3M9ImZvcm0tY29udHJvbCBpbnB1dC1zbSIgaWQ9ImFldGl0bGUtaWQiPiA8L2Rpdj4NCgkJCQkJCQk8L2Rpdj4NCgkJCQkJCTwvdGQ+DQoJCQkJCQk8dGQ+DQoJCQkJCQkJPGRpdiBjbGFzcz0iZm9ybS1ncm91cCI+DQoJCQkJCQkJCTxsYWJlbCBjbGFzcz0iY29udHJvbC1sYWJlbCI+UG9ydCBudW1iZXI8L2xhYmVsPg0KCQkJCQkJCQk8ZGl2IGNsYXNzPSJjb250cm9scyI+DQoJCQkJCQkJCQk8aW5wdXQgdHlwZT0idGV4dCIgY2xhc3M9ImZvcm0tY29udHJvbCBpbnB1dC1zbSIgaWQ9InBvcnQtaWQiPiA8L2Rpdj4NCgkJCQkJCQk8L2Rpdj4NCgkJCQkJCTwvdGQ+DQoJCQkJCQk8dGQ+DQoJCQkJCQkJPGRpdiBjbGFzcz0iZm9ybS1ncm91cCI+DQoJCQkJCQkJCTxsYWJlbCBjbGFzcz0iY29udHJvbC1sYWJlbCI+RElDT00gcGluZyBzdGF0dXM6PC9sYWJlbD4NCgkJCQkJCQkJPHA+DQoJCQkJCQkJCQk8bGFiZWwgY2xhc3M9ImNvbnRyb2wtbGFiZWwiIGlkPSJwYWNzLXN0YXR1cy1pZCI+T0s8L2xhYmVsPg0KCQkJCQkJCQk8L3A+DQoJCQkJCQkJPC9kaXY+DQoJCQkJCQk8L3RkPg0KCQkJCQk8L3RyPg0KCQkJCTwvdGJvZHk+DQoJCQk8L3RhYmxlPg0KCQkJPGRpdiBjbGFzcz0icGFuZWwtZm9vdGVyIiBpZD0ic2VhcmNoLWZvb3RlciI+U2VhcmNoIFNldHRpbmdzIDwvZGl2Pg0KCQkJPHRhYmxlIGlkPSJzZWFyY2gtcGFuZWwiIGNsYXNzPSJ0YWJsZSB0YWJsZS1ib3JkZXJlZCB0YWJsZS1jb25kZW5zZWQgdGFibGUtaG92ZXIgdGFibGUtc3RyaXBlZCI+DQoJCQkJPHRib2R5Pg0KCQkJCQk8dHI+DQoJCQkJCQk8dGQ+DQoJCQkJCQkJPGRpdiBjbGFzcz0iZm9ybS1ncm91cCI+DQoJCQkJCQkJCTxsYWJlbCBjbGFzcz0iY29udHJvbC1sYWJlbCI+UGF0aWVudCBuYW1lPC9sYWJlbD4NCgkJCQkJCQkJPGRpdiBjbGFzcz0iY29udHJvbHMiPg0KCQkJCQkJCQkJPGlucHV0IHR5cGU9InRleHQiIGNsYXNzPSJmb3JtLWNvbnRyb2wgaW5wdXQtc20iIGlkPSJwYXRpZW50LW5hbWUtaWQiPiA8L2Rpdj4NCgkJCQkJCQk8L2Rpdj4NCgkJCQkJCTwvdGQ+DQoJCQkJCQk8dGQ+DQoJCQkJCQkJPGRpdiBjbGFzcz0iZm9ybS1ncm91cCI+DQoJCQkJCQkJCTxsYWJlbCBjbGFzcz0iY29udHJvbC1sYWJlbCI+QWNjZXNzaW9uIG51bWJlcjwvbGFiZWw+DQoJCQkJCQkJCTxkaXYgY2xhc3M9ImNvbnRyb2xzIj4NCgkJCQkJCQkJCTxpbnB1dCB0eXBlPSJ0ZXh0IiBjbGFzcz0iZm9ybS1jb250cm9sIGlucHV0LXNtIiBpZD0iYWNjZXNzaW9uLW51bWJlci1pZCI+IDwvZGl2Pg0KCQkJCQkJCTwvZGl2Pg0KCQkJCQkJPC90ZD4NCgkJCQkJCTx0ZD4NCgkJCQkJCQk8ZGl2IGNsYXNzPSJmb3JtLWdyb3VwIj4NCgkJCQkJCQkJPGxhYmVsIGNsYXNzPSJjb250cm9sLWxhYmVsIj5EYXRlIG9mIGJpcnRoPC9sYWJlbD4NCgkJCQkJCQkJPGRpdiBjbGFzcz0iY29udHJvbHMiPg0KCQkJCQkJCQkJPGlucHV0IHR5cGU9InRleHQiIGNsYXNzPSJmb3JtLWNvbnRyb2wgaW5wdXQtc20iIGlkPSJkYXRlLWJpcnRoLWlkIj4gPC9kaXY+DQoJCQkJCQkJPC9kaXY+DQoJCQkJCQk8L3RkPg0KCQkJCQkJPHRkPg0KCQkJCQkJCTxkaXYgY2xhc3M9ImZvcm0tZ3JvdXAiPg0KCQkJCQkJCQk8bGFiZWwgY2xhc3M9ImNvbnRyb2wtbGFiZWwiPlN0dWR5IGRhdGU8L2xhYmVsPg0KCQkJCQkJCQk8ZGl2IGNsYXNzPSJjb250cm9scyI+DQoJCQkJCQkJCQk8aW5wdXQgdHlwZT0idGV4dCIgY2xhc3M9ImZvcm0tY29udHJvbCBpbnB1dC1zbSIgaWQ9InN0dWR5LWRhdGUtaWQiPiA8L2Rpdj4NCgkJCQkJCQk8L2Rpdj4NCgkJCQkJCTwvdGQ+DQoJCQkJCQk8dGQ+DQoJCQkJCQkJPGRpdiBjbGFzcz0iZm9ybS1ncm91cCI+DQoJCQkJCQkJCTxwPjwvcD4gPGEgb25jbGljaz0ic2VuZENGaW5kKCkiIGNsYXNzPSJidG4gYnRuLXN1Y2Nlc3MiPkYgSSBOIEQ8L2E+IDwvZGl2Pg0KCQkJCQkJPC90ZD4NCgkJCQkJPC90cj4NCgkJCQk8L3Rib2R5Pg0KCQkJPC90YWJsZT4NCgkJPC9kaXY+DQoJPC9kaXY+DQo8L2JvZHk+DQoNCjwvaHRtbD4="

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
	http.HandleFunc("/index.html", srv.ServePage)
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
func (srv *DJsServ) ServePage(rwr http.ResponseWriter, req *http.Request) {
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

func (srv *DJsServ) DispatchError(fjb FaJob) error {
	log.Println("info: DispatchError")
	log.Println(fjb.ErrorData)
	return nil
}

func (srv *DJsServ) DispatchSuccess(cjb CompJob) error {
	log.Println("info: DispatchSuccess")
	switch result := cjb.ResultData.(type) {
	case EchoRes:
		return srv.OnCEchoDone(result)
	case []FindRes:
		return srv.OnCFindDone(result)
	default:
		log.Printf("unexpected job type %v", result)
	}
	return nil
}

func (srv *DJsServ) OnCEchoDone(eres EchoRes) error {
	srv.echSta = eres
	return nil
}

func (srv *DJsServ) OnCFindDone(fres []FindRes) error {
	srv.fRes = fres
	srv.fndTm = time.Now().Nanosecond()
	return nil
}
