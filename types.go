package main

import "errors"

// See: https://developers.google.com/analytics/devguides/collection/protocol/v1/parameters
type MeasurementProtocol struct {
	// General
	V   int    `form:"v" json:"v"`                         // *Protocol Version
	Tid string `form:"tid" json:"tid"`                     // *Tracking ID
	Aip bool   `form:"aip,omitempty" json:"aip,omitempty"` // Anonymize IP
	Ds  string `form:"ds,omitempty" json:"ds,omitempty"`   // Datasource
	Qt  int    `form:"qt,omitempty" json:"qt,omitempty"`   // Queue Time
	Z   string `form:"z,omitempty" json:"z,omitempty"`     // Cache Buster

	// User
	Cid string `form:"cid" json:"cid"`                     // *Client ID
	Uid string `form:"uid,omitempty" json:"uid,omitempty"` // User ID

	// Session
	Sc    string `form:"sc,omitempty" json:"sc,omitempty"`       // Session Control
	Uip   string `form:"uip,omitempty" json:"uip,omitempty"`     // IP Override
	Ua    string `form:"ua,omitempty" json:"ua,omitempty"`       // User Agent Override
	Geoid string `form:"geoid,omitempty" json:"geoid,omitempty"` // Geogrpahical Override

	// Traffic Sources
	Dr    string `form:"dr,omitempty" json:"dr,omitempty"`       // Document Referrer
	Cn    string `form:"cn,omitempty" json:"cn,omitempty"`       // Campaign Name
	Cs    string `form:"cs,omitempty" json:"cs,omitempty"`       // Campaign Source
	Cm    string `form:"cm,omitempty" json:"cm,omitempty"`       // Campaign Medium
	Ck    string `form:"ck,omitempty" json:"ck,omitempty"`       // Campaign Keyword
	Cc    string `form:"cc,omitempty" json:"cc,omitempty"`       // Campaign Content
	Ci    string `form:"ci,omitempty" json:"ci,omitempty"`       // Campaign ID
	Gclid string `form:"gclid,omitempty" json:"gclid,omitempty"` // Google AdWords ID
	Dclid string `form:"dclid,omitempty" json:"dclid,omitempty"` // Google Display Ads ID

	// System Info
	Sr string `form:"sr,omitempty" json:"sr,omitempty"` // Screen Resolution
	Vp string `form:"vp,omitempty" json:"vp,omitempty"` // Viewport Size
	De string `form:"de,omitempty" json:"de,omitempty"` // Document Encoding
	Sd string `form:"sd,omitempty" json:"sd,omitempty"` // String Color
	Ul string `form:"ul,omitempty" json:"ul,omitempty"` // User Language
	Je bool   `form:"je,omitempty" json:"je,omitempty"` // Java Enabled
	Fl string `form:"fl,omitempty" json:"fl,omitempty"` // Flash Version

	// Hit
	T  string `form:"t" json:"t"`                       // *Hit Type
	Ni bool   `form:"ni,omitempty" json:"ni,omitempty"` // Non-Interaction Hit

	// Content Information
	Dl     string `form:"dl,omitempty" json:"dl,omitempty"`         // Document Location URL
	Dh     string `form:"dh,omitempty" json:"dh,omitempty"`         // Document Hostname
	Dp     string `form:"dp,omitempty" json:"dp,omitempty"`         // Document Path
	Dt     string `form:"dt,omitempty" json:"dt,omitempty"`         // Document Title
	Dd     string `form:"dd,omitempty" json:"dd,omitempty"`         // Screen Name *Required for Mobile properties
	Linkid string `form:"linkid,omitempty" json:"linkid,omitempty"` // Link ID

	// App Tracking
	An   string `form:"an,omitempty" json:"an,omitempty"`     // Application Name *Not required for Web properties
	Aid  string `form:"aid,omitempty" json:"aid,omitempty"`   // Application ID
	Av   string `form:"av,omitempty" json:"av,omitempty"`     // Application Version
	Aiid string `form:"aiid,omitempty" json:"aiid,omitempty"` // Application Installer ID

	// Event Tracking
	Ec string `form:"ec,omitempty" json:"ec,omitempty"` // Event Category *Required for Event type only
	Ea string `form:"ea,omitempty" json:"ea,omitempty"` // Event Action
	El string `form:"el,omitempty" json:"el,omitempty"` // Event Label
	Ev string `form:"ev,omitempty" json:"ev,omitempty"` // Event Value

	// E-Commerce
	Ti string `form:"ti,omitempty" json:"ti,omitempty"` // Transaction ID *Required for Transacation and Item hit type
	Ta string `form:"ta,omitempty" json:"ta,omitempty"` // Transaction Affiliation
	Tr string `form:"tr,omitempty" json:"tr,omitempty"` // Transaction Revenue
	Ts string `form:"ts,omitempty" json:"ts,omitempty"` // Transaction Shipping
	Tt string `form:"tt,omitempty" json:"tt,omitempty"` // Transaction Tax
	In string `form:"in,omitempty" json:"in,omitempty"` // Item Name *Required for Item hit type
	Ip string `form:"ip,omitempty" json:"ip,omitempty"` // Item Price
	Iq string `form:"iq,omitempty" json:"iq,omitempty"` // Item Quantity
	Ic string `form:"ic,omitempty" json:"ic,omitempty"` // Item Code
	Iv string `form:"iv,omitempty" json:"iv,omitempty"` // Item Category
	Cu string `form:"cu,omitempty" json:"cu,omitempty"` // Currency Code

	// Enhanced E-Commerce
	// Not Implemented Yet

	// Social Interactions
	Sn string `form:"sn,omitempty" json:"sn,omitempty"` // Social Network *Required for Social hit type
	Sa string `form:"sa,omitempty" json:"sa,omitempty"` // Social Action *Required for Social hit type
	St string `form:"st,omitempty" json:"st,omitempty"` // Social Action Target *Required for Social hit type

	// Timing
	Utc string `form:"utc,omitempty" json:"utc,omitempty"` // User Timing Category *Required for Timing hit type
	Utv string `form:"utv,omitempty" json:"utv,omitempty"` // User Timing Variable *Required for Timing hit type
	Utt string `form:"utt,omitempty" json:"utt,omitempty"` // User Timing Time *Required for Timing hit type
	Utl string `form:"utl,omitempty" json:"utl,omitempty"` // User Timing Label
	Plt int    `form:"plt,omitempty" json:"plt,omitempty"` // Page Load Time
	Dns int    `form:"dns,omitempty" json:"dns,omitempty"` // DNS Time
	Pdt int    `form:"pdt,omitempty" json:"pdt,omitempty"` // Page Download Time
	Rrt int    `form:"rrt,omitempty" json:"rrt,omitempty"` // Redirect Reponse Time
	Tcp int    `form:"tcp,omitempty" json:"tcp,omitempty"` // TCP Connect Time
	Srt int    `form:"srt,omitempty" json:"srt,omitempty"` // Server Response Time
	Dit int    `form:"dit,omitempty" json:"dit,omitempty"` // DOM Interactive Time
	Clt int    `form:"clt,omitempty" json:"clt,omitempty"` // Content Load Time

	// Exceptions
	Exd string `form:"exd,omitempty" json:"exd,omitempty"` // Exception Description
	Exf bool   `form:"exf,omitempty" json:"exf,omitempty"` // Is Exception Fatal?

	// Custom Dimensions / Metrics
	// Not Implemented Yet

	// Content Experiments
	Xid  string `form:"xid,omitempty" json:"xid,omitempty"`   // Experiment ID
	Xvar string `form:"xvar,omitempty" json:"xvar,omitempty"` // Experiment Variant
}

func (mp *MeasurementProtocol) Validate() error {
	if mp.V == 0 {
		return errors.New("'v' (Protocol Version) is required for all hit types")
	}
	if mp.Tid == "" {
		return errors.New("'tid' (Tracking ID)  is required for all hit types")
	}
	if mp.Cid == "" {
		return errors.New("'cid' (Client ID)  is required for all hit types")
	}

	return nil
}
