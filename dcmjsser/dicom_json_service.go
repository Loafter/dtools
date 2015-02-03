package main

import "io/ioutil"
import "net/http"
import "strconv"
import "log"
import "encoding/json"
import "errors"
import "time"

import "encoding/base64"

const htmlData = "PCFkb2N0eXBlIGh0bWw+DQo8aHRtbD4NCg0KPGhlYWQ+DQoJPHRpdGxlPkp1c3RpZmllZCBOYXY8L3RpdGxlPg0KCTxtZXRhIG5hbWU9InZpZXdwb3J0IiBjb250ZW50PSJ3aWR0aD1kZXZpY2Utd2lkdGgiPg0KCTxsaW5rIHJlbD0ic3R5bGVzaGVldCIgaHJlZj0iaHR0cHM6Ly9uZXRkbmEuYm9vdHN0cmFwY2RuLmNvbS9ib290c3dhdGNoLzMuMC4wL3NsYXRlL2Jvb3RzdHJhcC5taW4uY3NzIj4NCgk8c2NyaXB0IHR5cGU9InRleHQvamF2YXNjcmlwdCIgc3JjPSJodHRwczovL2FqYXguZ29vZ2xlYXBpcy5jb20vYWpheC9saWJzL2pxdWVyeS8yLjAuMy9qcXVlcnkubWluLmpzIj48L3NjcmlwdD4NCgk8c2NyaXB0IHR5cGU9InRleHQvamF2YXNjcmlwdCIgc3JjPSJodHRwczovL25ldGRuYS5ib290c3RyYXBjZG4uY29tL2Jvb3RzdHJhcC8zLjEuMS9qcy9ib290c3RyYXAubWluLmpzIj48L3NjcmlwdD4NCgk8c3R5bGUgdHlwZT0idGV4dC9jc3MiPg0KCQlib2R5IHsNCgkJCXBhZGRpbmctdG9wOiAyMHB4Ow0KCQl9DQoJCS5mb290ZXIgew0KCQkJYm9yZGVyLXRvcDogMXB4IHNvbGlkICNlZWU7DQoJCQltYXJnaW4tdG9wOiA0MHB4Ow0KCQkJcGFkZGluZy10b3A6IDQwcHg7DQoJCQlwYWRkaW5nLWJvdHRvbTogNDBweDsNCgkJfQ0KCQkvKiBNYWluIG1hcmtldGluZyBtZXNzYWdlIGFuZCBzaWduIHVwIGJ1dHRvbiAqLw0KDQoJCS5qdW1ib3Ryb24gew0KCQkJdGV4dC1hbGlnbjogY2VudGVyOw0KCQkJYmFja2dyb3VuZC1jb2xvcjogdHJhbnNwYXJlbnQ7DQoJCX0NCgkJLmp1bWJvdHJvbiAuYnRuIHsNCgkJCWZvbnQtc2l6ZTogMjFweDsNCgkJCXBhZGRpbmc6IDE0cHggMjRweDsNCgkJfQ0KCQkvKiBDdXN0b21pemUgdGhlIG5hdi1qdXN0aWZpZWQgbGlua3MgdG8gYmUgZmlsbCB0aGUgZW50aXJlIHNwYWNlIG9mIHRoZSAubmF2YmFyICovDQoNCgkJLm5hdi1qdXN0aWZpZWQgew0KCQkJYmFja2dyb3VuZC1jb2xvcjogI2VlZTsNCgkJCWJvcmRlci1yYWRpdXM6IDVweDsNCgkJCWJvcmRlcjogMXB4IHNvbGlkICNjY2M7DQoJCX0NCgkJLm5hdi1qdXN0aWZpZWQgPiBsaSA+IGEgew0KCQkJcGFkZGluZy10b3A6IDE1cHg7DQoJCQlwYWRkaW5nLWJvdHRvbTogMTVweDsNCgkJCWNvbG9yOiAjNzc3Ow0KCQkJZm9udC13ZWlnaHQ6IGJvbGQ7DQoJCQl0ZXh0LWFsaWduOiBjZW50ZXI7DQoJCQlib3JkZXItYm90dG9tOiAxcHggc29saWQgI2Q1ZDVkNTsNCgkJCWJhY2tncm91bmQtY29sb3I6ICNlNWU1ZTU7DQoJCQkvKiBPbGQgYnJvd3NlcnMgKi8NCg0KCQkJYmFja2dyb3VuZC1yZXBlYXQ6IHJlcGVhdC14Ow0KCQkJLyogUmVwZWF0IHRoZSBncmFkaWVudCAqLw0KDQoJCQliYWNrZ3JvdW5kLWltYWdlOiAtbW96LWxpbmVhci1ncmFkaWVudCh0b3AsICNmNWY1ZjUgMCUsICNlNWU1ZTUgMTAwJSk7DQoJCQkvKiBGRjMuNisgKi8NCg0KCQkJYmFja2dyb3VuZC1pbWFnZTogLXdlYmtpdC1ncmFkaWVudChsaW5lYXIsIGxlZnQgdG9wLCBsZWZ0IGJvdHRvbSwgY29sb3Itc3RvcCgwJSwgI2Y1ZjVmNSksIGNvbG9yLXN0b3AoMTAwJSwgI2U1ZTVlNSkpOw0KCQkJLyogQ2hyb21lLFNhZmFyaTQrICovDQoNCgkJCWJhY2tncm91bmQtaW1hZ2U6IC13ZWJraXQtbGluZWFyLWdyYWRpZW50KHRvcCwgI2Y1ZjVmNSAwJSwgI2U1ZTVlNSAxMDAlKTsNCgkJCS8qIENocm9tZSAxMCssU2FmYXJpIDUuMSsgKi8NCg0KCQkJYmFja2dyb3VuZC1pbWFnZTogLW1zLWxpbmVhci1ncmFkaWVudCh0b3AsICNmNWY1ZjUgMCUsICNlNWU1ZTUgMTAwJSk7DQoJCQkvKiBJRTEwKyAqLw0KDQoJCQliYWNrZ3JvdW5kLWltYWdlOiAtby1saW5lYXItZ3JhZGllbnQodG9wLCAjZjVmNWY1IDAlLCAjZTVlNWU1IDEwMCUpOw0KCQkJLyogT3BlcmEgMTEuMTArICovDQoNCgkJCWZpbHRlcjogcHJvZ2lkOiBEWEltYWdlVHJhbnNmb3JtLk1pY3Jvc29mdC5ncmFkaWVudChzdGFydENvbG9yc3RyPScjZjVmNWY1JywgZW5kQ29sb3JzdHI9JyNlNWU1ZTUnLCBHcmFkaWVudFR5cGU9MCk7DQoJCQkvKiBJRTYtOSAqLw0KDQoJCQliYWNrZ3JvdW5kLWltYWdlOiBsaW5lYXItZ3JhZGllbnQodG9wLCAjZjVmNWY1IDAlLCAjZTVlNWU1IDEwMCUpOw0KCQkJLyogVzNDICovDQoJCX0NCgkJLm5hdi1qdXN0aWZpZWQgPiAuYWN0aXZlID4gYSwNCgkJLm5hdi1qdXN0aWZpZWQgPiAuYWN0aXZlID4gYTpob3ZlciwNCgkJLm5hdi1qdXN0aWZpZWQgPiAuYWN0aXZlID4gYTpmb2N1cyB7DQoJCQliYWNrZ3JvdW5kLWNvbG9yOiAjZGRkOw0KCQkJYmFja2dyb3VuZC1pbWFnZTogbm9uZTsNCgkJCWJveC1zaGFkb3c6IGluc2V0IDAgM3B4IDdweCByZ2JhKDAsIDAsIDAsIC4xNSk7DQoJCX0NCgkJLm5hdi1qdXN0aWZpZWQgPiBsaTpmaXJzdC1jaGlsZCA+IGEgew0KCQkJYm9yZGVyLXJhZGl1czogNXB4IDVweCAwIDA7DQoJCX0NCgkJLm5hdi1qdXN0aWZpZWQgPiBsaTpsYXN0LWNoaWxkID4gYSB7DQoJCQlib3JkZXItYm90dG9tOiAwOw0KCQkJYm9yZGVyLXJhZGl1czogMCAwIDVweCA1cHg7DQoJCX0NCgkJQG1lZGlhKG1pbi13aWR0aDogNzY4cHgpIHsNCgkJCS5uYXYtanVzdGlmaWVkIHsNCgkJCQltYXgtaGVpZ2h0OiA1MnB4Ow0KCQkJfQ0KCQkJLm5hdi1qdXN0aWZpZWQgPiBsaSA+IGEgew0KCQkJCWJvcmRlci1sZWZ0OiAxcHggc29saWQgI2ZmZjsNCgkJCQlib3JkZXItcmlnaHQ6IDFweCBzb2xpZCAjZDVkNWQ1Ow0KCQkJfQ0KCQkJLm5hdi1qdXN0aWZpZWQgPiBsaTpmaXJzdC1jaGlsZCA+IGEgew0KCQkJCWJvcmRlci1sZWZ0OiAwOw0KCQkJCWJvcmRlci1yYWRpdXM6IDVweCAwIDAgNXB4Ow0KCQkJfQ0KCQkJLm5hdi1qdXN0aWZpZWQgPiBsaTpsYXN0LWNoaWxkID4gYSB7DQoJCQkJYm9yZGVyLXJhZGl1czogMCA1cHggNXB4IDA7DQoJCQkJYm9yZGVyLXJpZ2h0OiAwOw0KCQkJfQ0KCQl9DQoJCS8qIFJlc3BvbnNpdmU6IFBvcnRyYWl0IHRhYmxldHMgYW5kIHVwICovDQoNCgkJQG1lZGlhIHNjcmVlbiBhbmQobWluLXdpZHRoOiA3NjhweCkgew0KCQkJLyogUmVtb3ZlIHRoZSBwYWRkaW5nIHdlIHNldCBlYXJsaWVyICovDQoNCgkJCS5tYXN0aGVhZCwNCgkJCS5tYXJrZXRpbmcsDQoJCQkuZm9vdGVyIHsNCgkJCQlwYWRkaW5nLWxlZnQ6IDA7DQoJCQkJcGFkZGluZy1yaWdodDogMDsNCgkJCX0NCgkJfQ0KCTwvc3R5bGU+DQoJPHNjcmlwdCB0eXBlPSJ0ZXh0L2phdmFzY3JpcHQiPg0KCQlmdW5jdGlvbiB1cGRhdGVDRWNob1N0KCkgew0KCQkJdmFyIGNFQ2hvUmVxID0gew0KCQkJCUFkZHJlc3M6ICQoIiNhZGRyZXNzLWlkIikudmFsKCksDQoJCQkJUG9ydDogJCgiI3BvcnQtaWQiKS52YWwoKSwNCgkJCQlTZXJ2ZXJBRV9UaXRsZTogJCgiI2FldGl0bGUtaWQiKS52YWwoKQ0KCQkJfTsNCgkJCSQuYWpheCh7DQoJCQkJdXJsOiAiL2MtZWNobyIsDQoJCQkJdHlwZTogIlBPU1QiLA0KCQkJCWRhdGE6IEpTT04uc3RyaW5naWZ5KGNFQ2hvUmVxKSwNCgkJCQlkYXRhVHlwZTogImpzb24iDQoJCQl9KS5kb25lKGZ1bmN0aW9uKGpzb25EYXRhKSB7DQoJCQkJY29uc29sZS5sb2coanNvbkRhdGEpDQoJCQkJaWYgKGpzb25EYXRhLklzQWxpdmUpIHsNCgkJCQkJJCgiI3BhY3Mtc3RhdHVzLWlkIikudGV4dCgib2siKQ0KCQkJCX0gZWxzZSB7DQoJCQkJCSQoIiNwYWNzLXN0YXR1cy1pZCIpLnRleHQoIm5vIGNvbm5lY3Rpb24iKQ0KCQkJCX0NCg0KCQkJfSkNCgkJfQ0KDQoJCWZ1bmN0aW9uIE9uTG9hZCgpIHsNCgkJCXNldEludGVydmFsKHVwZGF0ZUNFY2hvU3QsIDQwMCkNCgkJfQ0KCTwvc2NyaXB0Pg0KPC9oZWFkPg0KDQo8Ym9keSBvbmxvYWQ9Ik9uTG9hZCgpIj4NCgk8ZGl2IGNsYXNzPSJjb250YWluZXIiPg0KCQk8ZGl2IGNsYXNzPSJ3ZWxsIj4NCgkJCTxkaXYgY2xhc3M9InBhbmVsLWZvb3RlciI+RElDT00gU2VydmVyIHNldHRpbmdzIDwvZGl2Pg0KCQkJPHRhYmxlIGNsYXNzPSJ0YWJsZSB0YWJsZS1ib3JkZXJlZCB0YWJsZS1jb25kZW5zZWQgdGFibGUtaG92ZXIgdGFibGUtc3RyaXBlZCI+DQoJCQkJPHRib2R5Pg0KCQkJCQk8dHI+DQoJCQkJCQk8dGQ+DQoJCQkJCQkJPGRpdiBjbGFzcz0iZm9ybS1ncm91cCI+DQoJCQkJCQkJCTxsYWJlbCBjbGFzcz0iY29udHJvbC1sYWJlbCI+RElDT00gc2VydmVyIGFkZHJlc3M8L2xhYmVsPg0KCQkJCQkJCQk8ZGl2IGNsYXNzPSJjb250cm9scyI+DQoJCQkJCQkJCQk8aW5wdXQgdHlwZT0idGV4dCIgY2xhc3M9ImZvcm0tY29udHJvbCBpbnB1dC1zbSIgaWQ9ImFkZHJlc3MtaWQiPiA8L2Rpdj4NCgkJCQkJCQk8L2Rpdj4NCgkJCQkJCTwvdGQ+DQoJCQkJCQk8dGQ+DQoJCQkJCQkJPGRpdiBjbGFzcz0iZm9ybS1ncm91cCI+DQoJCQkJCQkJCTxsYWJlbCBjbGFzcz0iY29udHJvbC1sYWJlbCI+QUUtVGl0bGU8L2xhYmVsPg0KCQkJCQkJCQk8ZGl2IGNsYXNzPSJjb250cm9scyI+DQoJCQkJCQkJCQk8aW5wdXQgdHlwZT0idGV4dCIgY2xhc3M9ImZvcm0tY29udHJvbCBpbnB1dC1zbSIgaWQ9ImFldGl0bGUtaWQiPiA8L2Rpdj4NCgkJCQkJCQk8L2Rpdj4NCgkJCQkJCTwvdGQ+DQoJCQkJCQk8dGQ+DQoJCQkJCQkJPGRpdiBjbGFzcz0iZm9ybS1ncm91cCI+DQoJCQkJCQkJCTxsYWJlbCBjbGFzcz0iY29udHJvbC1sYWJlbCI+UG9ydCBudW1iZXI8L2xhYmVsPg0KCQkJCQkJCQk8ZGl2IGNsYXNzPSJjb250cm9scyI+DQoJCQkJCQkJCQk8aW5wdXQgdHlwZT0idGV4dCIgY2xhc3M9ImZvcm0tY29udHJvbCBpbnB1dC1zbSIgaWQ9InBvcnQtaWQiPiA8L2Rpdj4NCgkJCQkJCQk8L2Rpdj4NCgkJCQkJCTwvdGQ+DQoJCQkJCQk8dGQ+DQoJCQkJCQkJPGRpdiBjbGFzcz0iZm9ybS1ncm91cCI+DQoJCQkJCQkJCTxsYWJlbCBjbGFzcz0iY29udHJvbC1sYWJlbCI+RElDT00gcGluZyBzdGF0dXM6PC9sYWJlbD4NCgkJCQkJCQkJPHA+DQoJCQkJCQkJCQk8bGFiZWwgY2xhc3M9ImNvbnRyb2wtbGFiZWwiIGlkPSJwYWNzLXN0YXR1cy1pZCI+T0s8L2xhYmVsPg0KCQkJCQkJCQk8L3A+DQoJCQkJCQkJPC9kaXY+DQoJCQkJCQk8L3RkPg0KCQkJCQk8L3RyPg0KCQkJCTwvdGJvZHk+DQoJCQk8L3RhYmxlPg0KCQkJPGRpdiBjbGFzcz0icGFuZWwtZm9vdGVyIj5TZWFyY2ggU2V0dGluZ3MgPC9kaXY+DQoJCQk8dGFibGUgY2xhc3M9InRhYmxlIHRhYmxlLWJvcmRlcmVkIHRhYmxlLWNvbmRlbnNlZCB0YWJsZS1ob3ZlciB0YWJsZS1zdHJpcGVkIj4NCgkJCQk8dGJvZHk+DQoJCQkJCTx0cj4NCgkJCQkJCTx0ZD4NCgkJCQkJCQk8ZGl2IGNsYXNzPSJmb3JtLWdyb3VwIj4NCgkJCQkJCQkJPGxhYmVsIGNsYXNzPSJjb250cm9sLWxhYmVsIj5QYXRpZW50IG5hbWU8L2xhYmVsPg0KCQkJCQkJCQk8ZGl2IGNsYXNzPSJjb250cm9scyI+DQoJCQkJCQkJCQk8aW5wdXQgdHlwZT0idGV4dCIgY2xhc3M9ImZvcm0tY29udHJvbCBpbnB1dC1zbSI+IDwvZGl2Pg0KCQkJCQkJCTwvZGl2Pg0KCQkJCQkJPC90ZD4NCgkJCQkJCTx0ZD4NCgkJCQkJCQk8ZGl2IGNsYXNzPSJmb3JtLWdyb3VwIj4NCgkJCQkJCQkJPGxhYmVsIGNsYXNzPSJjb250cm9sLWxhYmVsIj5QYXRpZW50IE1STjwvbGFiZWw+DQoJCQkJCQkJCTxkaXYgY2xhc3M9ImNvbnRyb2xzIj4NCgkJCQkJCQkJCTxpbnB1dCB0eXBlPSJ0ZXh0IiBjbGFzcz0iZm9ybS1jb250cm9sIGlucHV0LXNtIj4gPC9kaXY+DQoJCQkJCQkJPC9kaXY+DQoJCQkJCQk8L3RkPg0KCQkJCQkJPHRkPg0KCQkJCQkJCTxkaXYgY2xhc3M9ImZvcm0tZ3JvdXAiPg0KCQkJCQkJCQk8bGFiZWwgY2xhc3M9ImNvbnRyb2wtbGFiZWwiPlN0dWR5IElEPC9sYWJlbD4NCgkJCQkJCQkJPGRpdiBjbGFzcz0iY29udHJvbHMiPg0KCQkJCQkJCQkJPGlucHV0IHR5cGU9InRleHQiIGNsYXNzPSJmb3JtLWNvbnRyb2wgaW5wdXQtc20iPiA8L2Rpdj4NCgkJCQkJCQk8L2Rpdj4NCgkJCQkJCTwvdGQ+DQoJCQkJCQk8dGQ+DQoJCQkJCQkJPGRpdiBjbGFzcz0iZm9ybS1ncm91cCI+DQoJCQkJCQkJCTxsYWJlbCBjbGFzcz0iY29udHJvbC1sYWJlbCI+RGF0ZSBvZiBiaXJ0aDwvbGFiZWw+DQoJCQkJCQkJCTxkaXYgY2xhc3M9ImNvbnRyb2xzIj4NCgkJCQkJCQkJCTxpbnB1dCB0eXBlPSJ0ZXh0IiBjbGFzcz0iZm9ybS1jb250cm9sIGlucHV0LXNtIj4gPC9kaXY+DQoJCQkJCQkJPC9kaXY+DQoJCQkJCQk8L3RkPg0KCQkJCQkJPHRkPg0KCQkJCQkJCTxkaXYgY2xhc3M9ImZvcm0tZ3JvdXAiPg0KCQkJCQkJCQk8bGFiZWwgY2xhc3M9ImNvbnRyb2wtbGFiZWwiPlN0dWR5IGRhdGU8L2xhYmVsPg0KCQkJCQkJCQk8ZGl2IGNsYXNzPSJjb250cm9scyI+DQoJCQkJCQkJCQk8aW5wdXQgdHlwZT0idGV4dCIgY2xhc3M9ImZvcm0tY29udHJvbCBpbnB1dC1zbSI+IDwvZGl2Pg0KCQkJCQkJCTwvZGl2Pg0KCQkJCQkJPC90ZD4NCgkJCQkJCTx0ZD4NCgkJCQkJCQk8ZGl2IGNsYXNzPSJmb3JtLWdyb3VwIj4NCgkJCQkJCQkJPHA+PC9wPiA8YSBvbmNsaWNrPSJ1cGRhdGVDRWNob1N0KCkiIGNsYXNzPSJidG4gYnRuLXN1Y2Nlc3MiPkYgSSBOIEQ8L2E+IDwvZGl2Pg0KCQkJCQkJPC90ZD4NCgkJCQkJPC90cj4NCgkJCQk8L3Rib2R5Pg0KCQkJPC90YWJsZT4NCgkJPC9kaXY+DQoJPC9kaXY+DQo8L2JvZHk+DQoNCjwvaHRtbD4NCg=="

type FindData struct {
	FTime    time.Time
	CfindRes []FindRes
	Refresh  bool
}

//main srv class
type DJsServ struct {
	jbBal  JobBallancer
	dDisp  DDisp
	echSta EchoRes
	fndTm  time.Time
	fRes   []FindRes
}

//start and init srv
func (srv *DJsServ) Start(listenPort int) error {
	srv.jbBal.Init(&srv.dDisp, srv, srv)
	srv.dDisp.dCln.CallerAE_Title = "AE_DTOOLS"
	http.HandleFunc("/c-echo", srv.cEcho)
	http.HandleFunc("/c-find", srv.cFind)
	http.HandleFunc("/c-findres", srv.cFindData)
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
	var lctim time.Time
	if err := json.Unmarshal(bodyData, &lctim); err != nil {
		strErr := "error: can't parse time data"
		http.Error(rwr, err.Error(), http.StatusInternalServerError)
		log.Println(strErr)
		return
	}
	if lctim.Before(srv.fndTm) {
		fdat := FindData{Refresh: false}
		js, err := json.Marshal(fdat)
		if err != nil {
			log.Printf("error: can't serialize cfind data")
			http.Error(rwr, err.Error(), http.StatusInternalServerError)
			return
		}
		rwr.Write(js)
	} else {
		fdat := FindData{Refresh: true, CfindRes: srv.fRes, FTime: srv.fndTm}
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
	srv.fndTm = time.Now()
	return nil
}
