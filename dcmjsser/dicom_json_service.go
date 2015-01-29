package main

import "io/ioutil"
import "net/http"
import "strconv"
import "log"
import "encoding/json"
import "errors"

import "encoding/base64"

const htmlData = "PCFkb2N0eXBlIGh0bWw+DQoNCjxodG1sPg0KICANCiAgPGhlYWQ+DQogICAgPHRpdGxlPkp1c3RpZmllZCBOYXY8L3RpdGxlPg0KICAgIDxtZXRhIG5hbWU9InZpZXdwb3J0IiBjb250ZW50PSJ3aWR0aD1kZXZpY2Utd2lkdGgiPg0KICAgIDxsaW5rIHJlbD0ic3R5bGVzaGVldCIgaHJlZj0iaHR0cHM6Ly9uZXRkbmEuYm9vdHN0cmFwY2RuLmNvbS9ib290c3dhdGNoLzMuMC4wL3NsYXRlL2Jvb3RzdHJhcC5taW4uY3NzIj4NCiAgICA8c2NyaXB0IHR5cGU9InRleHQvamF2YXNjcmlwdCIgc3JjPSJodHRwczovL2FqYXguZ29vZ2xlYXBpcy5jb20vYWpheC9saWJzL2pxdWVyeS8yLjAuMy9qcXVlcnkubWluLmpzIj48L3NjcmlwdD4NCiAgICA8c2NyaXB0IHR5cGU9InRleHQvamF2YXNjcmlwdCIgc3JjPSJodHRwczovL25ldGRuYS5ib290c3RyYXBjZG4uY29tL2Jvb3RzdHJhcC8zLjEuMS9qcy9ib290c3RyYXAubWluLmpzIj48L3NjcmlwdD4NCiAgICA8c3R5bGUgdHlwZT0idGV4dC9jc3MiPg0KICAgICAgYm9keSB7DQogICAgICAgIHBhZGRpbmctdG9wOiAyMHB4Ow0KICAgICAgfQ0KICAgICAgLmZvb3RlciB7DQogICAgICAgIGJvcmRlci10b3A6IDFweCBzb2xpZCAjZWVlOw0KICAgICAgICBtYXJnaW4tdG9wOiA0MHB4Ow0KICAgICAgICBwYWRkaW5nLXRvcDogNDBweDsNCiAgICAgICAgcGFkZGluZy1ib3R0b206IDQwcHg7DQogICAgICB9DQogICAgICAvKiBNYWluIG1hcmtldGluZyBtZXNzYWdlIGFuZCBzaWduIHVwIGJ1dHRvbiAqLw0KICAgICAgLmp1bWJvdHJvbiB7DQogICAgICAgIHRleHQtYWxpZ246IGNlbnRlcjsNCiAgICAgICAgYmFja2dyb3VuZC1jb2xvcjogdHJhbnNwYXJlbnQ7DQogICAgICB9DQogICAgICAuanVtYm90cm9uIC5idG4gew0KICAgICAgICBmb250LXNpemU6IDIxcHg7DQogICAgICAgIHBhZGRpbmc6IDE0cHggMjRweDsNCiAgICAgIH0NCiAgICAgIC8qIEN1c3RvbWl6ZSB0aGUgbmF2LWp1c3RpZmllZCBsaW5rcyB0byBiZSBmaWxsIHRoZSBlbnRpcmUgc3BhY2Ugb2YgdGhlIC5uYXZiYXIgKi8NCiAgICAgIC5uYXYtanVzdGlmaWVkIHsNCiAgICAgICAgYmFja2dyb3VuZC1jb2xvcjogI2VlZTsNCiAgICAgICAgYm9yZGVyLXJhZGl1czogNXB4Ow0KICAgICAgICBib3JkZXI6IDFweCBzb2xpZCAjY2NjOw0KICAgICAgfQ0KICAgICAgLm5hdi1qdXN0aWZpZWQgPiBsaSA+IGEgew0KICAgICAgICBwYWRkaW5nLXRvcDogMTVweDsNCiAgICAgICAgcGFkZGluZy1ib3R0b206IDE1cHg7DQogICAgICAgIGNvbG9yOiAjNzc3Ow0KICAgICAgICBmb250LXdlaWdodDogYm9sZDsNCiAgICAgICAgdGV4dC1hbGlnbjogY2VudGVyOw0KICAgICAgICBib3JkZXItYm90dG9tOiAxcHggc29saWQgI2Q1ZDVkNTsNCiAgICAgICAgYmFja2dyb3VuZC1jb2xvcjogI2U1ZTVlNTsNCiAgICAgICAgLyogT2xkIGJyb3dzZXJzICovDQogICAgICAgIGJhY2tncm91bmQtcmVwZWF0OiByZXBlYXQteDsNCiAgICAgICAgLyogUmVwZWF0IHRoZSBncmFkaWVudCAqLw0KICAgICAgICBiYWNrZ3JvdW5kLWltYWdlOiAtbW96LWxpbmVhci1ncmFkaWVudCh0b3AsICNmNWY1ZjUgMCUsICNlNWU1ZTUgMTAwJSk7DQogICAgICAgIC8qIEZGMy42KyAqLw0KICAgICAgICBiYWNrZ3JvdW5kLWltYWdlOiAtd2Via2l0LWdyYWRpZW50KGxpbmVhciwgbGVmdCB0b3AsIGxlZnQgYm90dG9tLCBjb2xvci1zdG9wKDAlLCAjZjVmNWY1KSwgY29sb3Itc3RvcCgxMDAlLCAjZTVlNWU1KSk7DQogICAgICAgIC8qIENocm9tZSxTYWZhcmk0KyAqLw0KICAgICAgICBiYWNrZ3JvdW5kLWltYWdlOiAtd2Via2l0LWxpbmVhci1ncmFkaWVudCh0b3AsICNmNWY1ZjUgMCUsICNlNWU1ZTUgMTAwJSk7DQogICAgICAgIC8qIENocm9tZSAxMCssU2FmYXJpIDUuMSsgKi8NCiAgICAgICAgYmFja2dyb3VuZC1pbWFnZTogLW1zLWxpbmVhci1ncmFkaWVudCh0b3AsICNmNWY1ZjUgMCUsICNlNWU1ZTUgMTAwJSk7DQogICAgICAgIC8qIElFMTArICovDQogICAgICAgIGJhY2tncm91bmQtaW1hZ2U6IC1vLWxpbmVhci1ncmFkaWVudCh0b3AsICNmNWY1ZjUgMCUsICNlNWU1ZTUgMTAwJSk7DQogICAgICAgIC8qIE9wZXJhIDExLjEwKyAqLw0KICAgICAgICBmaWx0ZXI6IHByb2dpZDpEWEltYWdlVHJhbnNmb3JtLk1pY3Jvc29mdC5ncmFkaWVudChzdGFydENvbG9yc3RyPScjZjVmNWY1JywgZW5kQ29sb3JzdHI9JyNlNWU1ZTUnLCBHcmFkaWVudFR5cGU9MCk7DQogICAgICAgIC8qIElFNi05ICovDQogICAgICAgIGJhY2tncm91bmQtaW1hZ2U6IGxpbmVhci1ncmFkaWVudCh0b3AsICNmNWY1ZjUgMCUsICNlNWU1ZTUgMTAwJSk7DQogICAgICAgIC8qIFczQyAqLw0KICAgICAgfQ0KICAgICAgLm5hdi1qdXN0aWZpZWQgPiAuYWN0aXZlID4gYSwgLm5hdi1qdXN0aWZpZWQgPiAuYWN0aXZlID4gYTpob3ZlciwgLm5hdi1qdXN0aWZpZWQgPiAuYWN0aXZlID4gYTpmb2N1cyB7DQogICAgICAgIGJhY2tncm91bmQtY29sb3I6ICNkZGQ7DQogICAgICAgIGJhY2tncm91bmQtaW1hZ2U6IG5vbmU7DQogICAgICAgIGJveC1zaGFkb3c6IGluc2V0IDAgM3B4IDdweCByZ2JhKDAsIDAsIDAsIC4xNSk7DQogICAgICB9DQogICAgICAubmF2LWp1c3RpZmllZCA+IGxpOmZpcnN0LWNoaWxkID4gYSB7DQogICAgICAgIGJvcmRlci1yYWRpdXM6IDVweCA1cHggMCAwOw0KICAgICAgfQ0KICAgICAgLm5hdi1qdXN0aWZpZWQgPiBsaTpsYXN0LWNoaWxkID4gYSB7DQogICAgICAgIGJvcmRlci1ib3R0b206IDA7DQogICAgICAgIGJvcmRlci1yYWRpdXM6IDAgMCA1cHggNXB4Ow0KICAgICAgfQ0KICAgICAgQG1lZGlhKG1pbi13aWR0aDogNzY4cHgpIHsNCiAgICAgICAgLm5hdi1qdXN0aWZpZWQgew0KICAgICAgICAgIG1heC1oZWlnaHQ6IDUycHg7DQogICAgICAgIH0NCiAgICAgICAgLm5hdi1qdXN0aWZpZWQgPiBsaSA+IGEgew0KICAgICAgICAgIGJvcmRlci1sZWZ0OiAxcHggc29saWQgI2ZmZjsNCiAgICAgICAgICBib3JkZXItcmlnaHQ6IDFweCBzb2xpZCAjZDVkNWQ1Ow0KICAgICAgICB9DQogICAgICAgIC5uYXYtanVzdGlmaWVkID4gbGk6Zmlyc3QtY2hpbGQgPiBhIHsNCiAgICAgICAgICBib3JkZXItbGVmdDogMDsNCiAgICAgICAgICBib3JkZXItcmFkaXVzOiA1cHggMCAwIDVweDsNCiAgICAgICAgfQ0KICAgICAgICAubmF2LWp1c3RpZmllZCA+IGxpOmxhc3QtY2hpbGQgPiBhIHsNCiAgICAgICAgICBib3JkZXItcmFkaXVzOiAwIDVweCA1cHggMDsNCiAgICAgICAgICBib3JkZXItcmlnaHQ6IDA7DQogICAgICAgIH0NCiAgICAgIH0NCiAgICAgIC8qIFJlc3BvbnNpdmU6IFBvcnRyYWl0IHRhYmxldHMgYW5kIHVwICovDQogICAgICBAbWVkaWEgc2NyZWVuIGFuZChtaW4td2lkdGg6IDc2OHB4KSB7DQogICAgICAgIC8qIFJlbW92ZSB0aGUgcGFkZGluZyB3ZSBzZXQgZWFybGllciAqLw0KICAgICAgICAubWFzdGhlYWQsIC5tYXJrZXRpbmcsIC5mb290ZXIgew0KICAgICAgICAgIHBhZGRpbmctbGVmdDogMDsNCiAgICAgICAgICBwYWRkaW5nLXJpZ2h0OiAwOw0KICAgICAgICB9DQogICAgICB9DQogICAgPC9zdHlsZT4NCgk8c2NyaXB0IHR5cGU9InRleHQvamF2YXNjcmlwdCI+DQoJDQoJDQoJZnVuY3Rpb24gdXBkYXRlQ0VjaG9TdCgpIHsNCgkJdmFyIGNFQ2hvUmVxID0gewkJCQkNCiAgICAgICAgCQkgICBBZGRyZXNzOiAkKCIjYWRkcmVzcy1pZCIpLnZhbCgpLA0KICAgICAgICAJCSAgIFBvcnQ6ICQoIiNwb3J0LWlkIikudmFsKCksDQoJCQkgICBTZXJ2ZXJBRV9UaXRsZTogJCgiI2FldGl0bGUtaWQiKS52YWwoKQ0KICAgIAkJCSAgIH07DQoJCSAgICAgICAkLmFqYXgoew0KICAgICAgICAgICAgICAgIHVybDogIi9jLWVjaG8iLA0KICAgICAgICAgICAgICAgIHR5cGU6ICJQT1NUIiwNCiAgICAgICAgICAgICAgICBkYXRhOkpTT04uc3RyaW5naWZ5KGNFQ2hvUmVxKSwNCiAgICAgICAgICAgICAgICBkYXRhVHlwZTogImpzb24iDQogICAgICAgICAgICB9KS5kb25lKGZ1bmN0aW9uKGpzb25EYXRhKSB7DQogDQogICAgICAgICAgICB9KQ0KCQl9DQoJPC9zY3JpcHQ+DQogIDwvaGVhZD4NCiAgDQogIDxib2R5Pg0KICAgIDxkaXYgY2xhc3M9ImNvbnRhaW5lciI+DQogICAgICA8ZGl2IGNsYXNzPSJ3ZWxsIj4NCiAgICAgICAgPGRpdiBjbGFzcz0icGFuZWwtZm9vdGVyIj5ESUNPTSBTZXJ2ZXIgc2V0dGluZ3MNCiAgICAgICAgPC9kaXY+DQogICAgICAgIDx0YWJsZSBjbGFzcz0idGFibGUgdGFibGUtYm9yZGVyZWQgdGFibGUtY29uZGVuc2VkIHRhYmxlLWhvdmVyIHRhYmxlLXN0cmlwZWQiPg0KICAgICAgICAgIDx0Ym9keT4NCiAgICAgICAgICAgIDx0cj4NCiAgICAgICAgICAgICAgPHRkPg0KICAgICAgICAgICAgICAgIDxkaXYgY2xhc3M9ImZvcm0tZ3JvdXAiPg0KICAgICAgICAgICAgICAgICAgPGxhYmVsIGNsYXNzPSJjb250cm9sLWxhYmVsIj5ESUNPTSBzZXJ2ZXIgYWRkcmVzczwvbGFiZWw+DQogICAgICAgICAgICAgICAgICA8ZGl2IGNsYXNzPSJjb250cm9scyI+DQogICAgICAgICAgICAgICAgICAgIDxpbnB1dCB0eXBlPSJ0ZXh0IiBjbGFzcz0iZm9ybS1jb250cm9sIGlucHV0LXNtIiBpZD0iYWRkcmVzcy1pZCI+DQogICAgICAgICAgICAgICAgICA8L2Rpdj4NCiAgICAgICAgICAgICAgICA8L2Rpdj4NCiAgICAgICAgICAgICAgPC90ZD4NCiAgICAgICAgICAgICAgPHRkPg0KICAgICAgICAgICAgICAgIDxkaXYgY2xhc3M9ImZvcm0tZ3JvdXAiPg0KICAgICAgICAgICAgICAgICAgPGxhYmVsIGNsYXNzPSJjb250cm9sLWxhYmVsIj5BRS1UaXRsZTwvbGFiZWw+DQogICAgICAgICAgICAgICAgICA8ZGl2IGNsYXNzPSJjb250cm9scyI+DQogICAgICAgICAgICAgICAgICAgIDxpbnB1dCB0eXBlPSJ0ZXh0IiBjbGFzcz0iZm9ybS1jb250cm9sIGlucHV0LXNtIiBpZD0iYWV0aXRsZS1pZCI+DQogICAgICAgICAgICAgICAgICA8L2Rpdj4NCiAgICAgICAgICAgICAgICA8L2Rpdj4NCiAgICAgICAgICAgICAgPC90ZD4NCiAgICAgICAgICAgICAgPHRkPg0KICAgICAgICAgICAgICAgIDxkaXYgY2xhc3M9ImZvcm0tZ3JvdXAiPg0KICAgICAgICAgICAgICAgICAgPGxhYmVsIGNsYXNzPSJjb250cm9sLWxhYmVsIj5Qb3J0IG51bWJlcjwvbGFiZWw+DQogICAgICAgICAgICAgICAgICA8ZGl2IGNsYXNzPSJjb250cm9scyI+DQogICAgICAgICAgICAgICAgICAgIDxpbnB1dCB0eXBlPSJ0ZXh0IiBjbGFzcz0iZm9ybS1jb250cm9sIGlucHV0LXNtIiBpZD0icG9ydC1pZCI+DQogICAgICAgICAgICAgICAgICA8L2Rpdj4NCiAgICAgICAgICAgICAgICA8L2Rpdj4NCiAgICAgICAgICAgICAgPC90ZD4NCiAgICAgICAgICAgICAgPHRkPg0KICAgICAgICAgICAgICAgIDxkaXYgY2xhc3M9ImZvcm0tZ3JvdXAiPg0KICAgICAgICAgICAgICAgICAgPGxhYmVsIGNsYXNzPSJjb250cm9sLWxhYmVsIj5ESUNPTSBwaW5nIHN0YXR1czo8L2xhYmVsPg0KICAgICAgICAgICAgICAgICAgPHA+DQogICAgICAgICAgICAgICAgICAgIDxsYWJlbCBjbGFzcz0iY29udHJvbC1sYWJlbCI+T0s8L2xhYmVsPjwvcD4NCiAgICAgICAgICAgICAgICA8L2Rpdj4NCiAgICAgICAgICAgICAgPC90ZD4NCiAgICAgICAgICAgIDwvdHI+DQogICAgICAgICAgPC90Ym9keT4NCiAgICAgICAgPC90YWJsZT4NCiAgICAgICAgPGRpdiBjbGFzcz0icGFuZWwtZm9vdGVyIj5TZWFyY2ggU2V0dGluZ3MNCiAgICAgICAgPC9kaXY+DQogICAgICAgIDx0YWJsZSBjbGFzcz0idGFibGUgdGFibGUtYm9yZGVyZWQgdGFibGUtY29uZGVuc2VkIHRhYmxlLWhvdmVyIHRhYmxlLXN0cmlwZWQiPg0KICAgICAgICAgIDx0Ym9keT4NCiAgICAgICAgICAgIDx0cj4NCiAgICAgICAgICAgICAgPHRkPg0KICAgICAgICAgICAgICAgIDxkaXYgY2xhc3M9ImZvcm0tZ3JvdXAiPg0KICAgICAgICAgICAgICAgICAgPGxhYmVsIGNsYXNzPSJjb250cm9sLWxhYmVsIj5QYXRpZW50IG5hbWU8L2xhYmVsPg0KICAgICAgICAgICAgICAgICAgPGRpdiBjbGFzcz0iY29udHJvbHMiPg0KICAgICAgICAgICAgICAgICAgICA8aW5wdXQgdHlwZT0idGV4dCIgY2xhc3M9ImZvcm0tY29udHJvbCBpbnB1dC1zbSI+DQogICAgICAgICAgICAgICAgICA8L2Rpdj4NCiAgICAgICAgICAgICAgICA8L2Rpdj4NCiAgICAgICAgICAgICAgPC90ZD4NCiAgICAgICAgICAgICAgPHRkPg0KICAgICAgICAgICAgICAgIDxkaXYgY2xhc3M9ImZvcm0tZ3JvdXAiPg0KICAgICAgICAgICAgICAgICAgPGxhYmVsIGNsYXNzPSJjb250cm9sLWxhYmVsIj5QYXRpZW50IE1STjwvbGFiZWw+DQogICAgICAgICAgICAgICAgICA8ZGl2IGNsYXNzPSJjb250cm9scyI+DQogICAgICAgICAgICAgICAgICAgIDxpbnB1dCB0eXBlPSJ0ZXh0IiBjbGFzcz0iZm9ybS1jb250cm9sIGlucHV0LXNtIj4NCiAgICAgICAgICAgICAgICAgIDwvZGl2Pg0KICAgICAgICAgICAgICAgIDwvZGl2Pg0KICAgICAgICAgICAgICA8L3RkPg0KICAgICAgICAgICAgICA8dGQ+DQogICAgICAgICAgICAgICAgPGRpdiBjbGFzcz0iZm9ybS1ncm91cCI+DQogICAgICAgICAgICAgICAgICA8bGFiZWwgY2xhc3M9ImNvbnRyb2wtbGFiZWwiPlN0dWR5IElEPC9sYWJlbD4NCiAgICAgICAgICAgICAgICAgIDxkaXYgY2xhc3M9ImNvbnRyb2xzIj4NCiAgICAgICAgICAgICAgICAgICAgPGlucHV0IHR5cGU9InRleHQiIGNsYXNzPSJmb3JtLWNvbnRyb2wgaW5wdXQtc20iPg0KICAgICAgICAgICAgICAgICAgPC9kaXY+DQogICAgICAgICAgICAgICAgPC9kaXY+DQogICAgICAgICAgICAgIDwvdGQ+DQogICAgICAgICAgICAgIDx0ZD4NCiAgICAgICAgICAgICAgICA8ZGl2IGNsYXNzPSJmb3JtLWdyb3VwIj4NCiAgICAgICAgICAgICAgICAgIDxsYWJlbCBjbGFzcz0iY29udHJvbC1sYWJlbCI+RGF0ZSBvZiBiaXJ0aDwvbGFiZWw+DQogICAgICAgICAgICAgICAgICA8ZGl2IGNsYXNzPSJjb250cm9scyI+DQogICAgICAgICAgICAgICAgICAgIDxpbnB1dCB0eXBlPSJ0ZXh0IiBjbGFzcz0iZm9ybS1jb250cm9sIGlucHV0LXNtIj4NCiAgICAgICAgICAgICAgICAgIDwvZGl2Pg0KICAgICAgICAgICAgICAgIDwvZGl2Pg0KICAgICAgICAgICAgICA8L3RkPg0KICAgICAgICAgICAgICA8dGQ+DQogICAgICAgICAgICAgICAgPGRpdiBjbGFzcz0iZm9ybS1ncm91cCI+DQogICAgICAgICAgICAgICAgICA8bGFiZWwgY2xhc3M9ImNvbnRyb2wtbGFiZWwiPlN0dWR5IGRhdGU8L2xhYmVsPg0KICAgICAgICAgICAgICAgICAgPGRpdiBjbGFzcz0iY29udHJvbHMiPg0KICAgICAgICAgICAgICAgICAgICA8aW5wdXQgdHlwZT0idGV4dCIgY2xhc3M9ImZvcm0tY29udHJvbCBpbnB1dC1zbSI+DQogICAgICAgICAgICAgICAgICA8L2Rpdj4NCiAgICAgICAgICAgICAgICA8L2Rpdj4NCiAgICAgICAgICAgICAgPC90ZD4NCiAgICAgICAgICAgICAgPHRkPg0KICAgICAgICAgICAgICAgIDxkaXYgY2xhc3M9ImZvcm0tZ3JvdXAiPg0KICAgICAgICAgICAgICAgICAgPHA+PC9wPg0KICAgICAgICAgICAgICAgICAgPGEgb25jbGljaz0idXBkYXRlQ0VjaG9TdCgpIiBjbGFzcz0iYnRuIGJ0bi1zdWNjZXNzIj5GIEkgTiBEPC9hPg0KICAgICAgICAgICAgICAgIDwvZGl2Pg0KICAgICAgICAgICAgICA8L3RkPg0KICAgICAgICAgICAgPC90cj4NCiAgICAgICAgICA8L3Rib2R5Pg0KICAgICAgICA8L3RhYmxlPg0KICAgICAgPC9kaXY+DQogICAgPC9kaXY+DQogIDwvYm9keT4NCg0KPC9odG1sPg=="

type HttpResReq struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
}

//main service class
type DicomJsonService struct {
	jobBallancer    JobBallancer
	dicomDispatcher DicomDispatcher
	responses       map[string]HttpResReq
}

//start and init service
func (service *DicomJsonService) Start(listenPort int) error {
	service.responses = make(map[string]HttpResReq)
	onCompletedResp := new(OnCompletedResp)
	onCompletedResp.Init(service.responses)
	onErrorResp := new(OnErrorResp)
	onErrorResp.Init(service.responses)
	service.jobBallancer.Init(&service.dicomDispatcher, onCompletedResp, onErrorResp)
	http.HandleFunc("/c-echo", service.cEcho)
	http.HandleFunc("/index.html", service.ServePage)
	if err := http.ListenAndServe(":"+strconv.Itoa(listenPort), nil); err != nil {
		return errors.New("error: can't start listen http server")
	}
	return nil
}

//serve cEcho responce
func (service *DicomJsonService) cEcho(responseWriter http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	bodyData, err := ioutil.ReadAll(request.Body)
	if err != nil {
		strErr := "error: Can't read http body data"
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		log.Println(strErr)
		return
	}
	var dicomCEchoRequest DicomCEchoRequest
	if err := json.Unmarshal(bodyData, &dicomCEchoRequest); err != nil {
		strErr := "error: can't parse DicomCEchoRequest data"
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		log.Println(strErr)
		return
	}

	if guid, err := service.jobBallancer.PushJob(dicomCEchoRequest); err == nil {
		service.responses[guid] = HttpResReq{Request: request, ResponseWriter: responseWriter}
	} else {
		log.Printf("error: can't push job")
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
	}

}

//serve main page request
func (service *DicomJsonService) ServePage(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type: text/html", "*")

	content, err := ioutil.ReadFile("index.html")
	if err != nil {
		val, _ := base64.StdEncoding.DecodeString(htmlData)
		responseWriter.Write(val)
		return
	}
	responseWriter.Write(content)
}
