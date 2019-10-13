// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/ad_group.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "google.golang.org/genproto/googleapis/ads/googleads/v2/common"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v2/enums"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// An ad group.
type AdGroup struct {
	// The resource name of the ad group.
	// Ad group resource names have the form:
	//
	// `customers/{customer_id}/adGroups/{ad_group_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ID of the ad group.
	Id *wrappers.Int64Value `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	// The name of the ad group.
	//
	// This field is required and should not be empty when creating new ad
	// groups.
	//
	// It must contain fewer than 255 UTF-8 full-width characters.
	//
	// It must not contain any null (code point 0x0), NL line feed
	// (code point 0xA) or carriage return (code point 0xD) characters.
	Name *wrappers.StringValue `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// The status of the ad group.
	Status enums.AdGroupStatusEnum_AdGroupStatus `protobuf:"varint,5,opt,name=status,proto3,enum=google.ads.googleads.v2.enums.AdGroupStatusEnum_AdGroupStatus" json:"status,omitempty"`
	// The type of the ad group.
	Type enums.AdGroupTypeEnum_AdGroupType `protobuf:"varint,12,opt,name=type,proto3,enum=google.ads.googleads.v2.enums.AdGroupTypeEnum_AdGroupType" json:"type,omitempty"`
	// The ad rotation mode of the ad group.
	AdRotationMode enums.AdGroupAdRotationModeEnum_AdGroupAdRotationMode `protobuf:"varint,22,opt,name=ad_rotation_mode,json=adRotationMode,proto3,enum=google.ads.googleads.v2.enums.AdGroupAdRotationModeEnum_AdGroupAdRotationMode" json:"ad_rotation_mode,omitempty"`
	// For draft or experiment ad groups, this field is the resource name of the
	// base ad group from which this ad group was created. If a draft or
	// experiment ad group does not have a base ad group, then this field is null.
	//
	// For base ad groups, this field equals the ad group resource name.
	//
	// This field is read-only.
	BaseAdGroup *wrappers.StringValue `protobuf:"bytes,18,opt,name=base_ad_group,json=baseAdGroup,proto3" json:"base_ad_group,omitempty"`
	// The URL template for constructing a tracking URL.
	TrackingUrlTemplate *wrappers.StringValue `protobuf:"bytes,13,opt,name=tracking_url_template,json=trackingUrlTemplate,proto3" json:"tracking_url_template,omitempty"`
	// The list of mappings used to substitute custom parameter tags in a
	// `tracking_url_template`, `final_urls`, or `mobile_final_urls`.
	UrlCustomParameters []*common.CustomParameter `protobuf:"bytes,6,rep,name=url_custom_parameters,json=urlCustomParameters,proto3" json:"url_custom_parameters,omitempty"`
	// The campaign to which the ad group belongs.
	Campaign *wrappers.StringValue `protobuf:"bytes,10,opt,name=campaign,proto3" json:"campaign,omitempty"`
	// The maximum CPC (cost-per-click) bid.
	CpcBidMicros *wrappers.Int64Value `protobuf:"bytes,14,opt,name=cpc_bid_micros,json=cpcBidMicros,proto3" json:"cpc_bid_micros,omitempty"`
	// The maximum CPM (cost-per-thousand viewable impressions) bid.
	CpmBidMicros *wrappers.Int64Value `protobuf:"bytes,15,opt,name=cpm_bid_micros,json=cpmBidMicros,proto3" json:"cpm_bid_micros,omitempty"`
	// The target CPA (cost-per-acquisition).
	TargetCpaMicros *wrappers.Int64Value `protobuf:"bytes,27,opt,name=target_cpa_micros,json=targetCpaMicros,proto3" json:"target_cpa_micros,omitempty"`
	// The CPV (cost-per-view) bid.
	CpvBidMicros *wrappers.Int64Value `protobuf:"bytes,17,opt,name=cpv_bid_micros,json=cpvBidMicros,proto3" json:"cpv_bid_micros,omitempty"`
	// Average amount in micros that the advertiser is willing to pay for every
	// thousand times the ad is shown.
	TargetCpmMicros *wrappers.Int64Value `protobuf:"bytes,26,opt,name=target_cpm_micros,json=targetCpmMicros,proto3" json:"target_cpm_micros,omitempty"`
	// The target ROAS (return-on-ad-spend) override. If the ad group's campaign
	// bidding strategy is a standard Target ROAS strategy, then this field
	// overrides the target ROAS specified in the campaign's bidding strategy.
	// Otherwise, this value is ignored.
	TargetRoas *wrappers.DoubleValue `protobuf:"bytes,30,opt,name=target_roas,json=targetRoas,proto3" json:"target_roas,omitempty"`
	// The percent cpc bid amount, expressed as a fraction of the advertised price
	// for some good or service. The valid range for the fraction is [0,1) and the
	// value stored here is 1,000,000 * [fraction].
	PercentCpcBidMicros *wrappers.Int64Value `protobuf:"bytes,20,opt,name=percent_cpc_bid_micros,json=percentCpcBidMicros,proto3" json:"percent_cpc_bid_micros,omitempty"`
	// Settings for the Display Campaign Optimizer, initially termed "Explorer".
	ExplorerAutoOptimizerSetting *common.ExplorerAutoOptimizerSetting `protobuf:"bytes,21,opt,name=explorer_auto_optimizer_setting,json=explorerAutoOptimizerSetting,proto3" json:"explorer_auto_optimizer_setting,omitempty"`
	// Allows advertisers to specify a targeting dimension on which to place
	// absolute bids. This is only applicable for campaigns that target only the
	// display network and not search.
	DisplayCustomBidDimension enums.TargetingDimensionEnum_TargetingDimension `protobuf:"varint,23,opt,name=display_custom_bid_dimension,json=displayCustomBidDimension,proto3,enum=google.ads.googleads.v2.enums.TargetingDimensionEnum_TargetingDimension" json:"display_custom_bid_dimension,omitempty"`
	// URL template for appending params to Final URL.
	FinalUrlSuffix *wrappers.StringValue `protobuf:"bytes,24,opt,name=final_url_suffix,json=finalUrlSuffix,proto3" json:"final_url_suffix,omitempty"`
	// Setting for targeting related features.
	TargetingSetting *common.TargetingSetting `protobuf:"bytes,25,opt,name=targeting_setting,json=targetingSetting,proto3" json:"targeting_setting,omitempty"`
	// The effective target CPA (cost-per-acquisition).
	// This field is read-only.
	EffectiveTargetCpaMicros *wrappers.Int64Value `protobuf:"bytes,28,opt,name=effective_target_cpa_micros,json=effectiveTargetCpaMicros,proto3" json:"effective_target_cpa_micros,omitempty"`
	// Source of the effective target CPA.
	// This field is read-only.
	EffectiveTargetCpaSource enums.BiddingSourceEnum_BiddingSource `protobuf:"varint,29,opt,name=effective_target_cpa_source,json=effectiveTargetCpaSource,proto3,enum=google.ads.googleads.v2.enums.BiddingSourceEnum_BiddingSource" json:"effective_target_cpa_source,omitempty"`
	// The effective target ROAS (return-on-ad-spend).
	// This field is read-only.
	EffectiveTargetRoas *wrappers.DoubleValue `protobuf:"bytes,31,opt,name=effective_target_roas,json=effectiveTargetRoas,proto3" json:"effective_target_roas,omitempty"`
	// Source of the effective target ROAS.
	// This field is read-only.
	EffectiveTargetRoasSource enums.BiddingSourceEnum_BiddingSource `protobuf:"varint,32,opt,name=effective_target_roas_source,json=effectiveTargetRoasSource,proto3,enum=google.ads.googleads.v2.enums.BiddingSourceEnum_BiddingSource" json:"effective_target_roas_source,omitempty"`
	// The resource names of labels attached to this ad group.
	Labels               []*wrappers.StringValue `protobuf:"bytes,33,rep,name=labels,proto3" json:"labels,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *AdGroup) Reset()         { *m = AdGroup{} }
func (m *AdGroup) String() string { return proto.CompactTextString(m) }
func (*AdGroup) ProtoMessage()    {}
func (*AdGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1544ca6322174f1, []int{0}
}

func (m *AdGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroup.Unmarshal(m, b)
}
func (m *AdGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroup.Marshal(b, m, deterministic)
}
func (m *AdGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroup.Merge(m, src)
}
func (m *AdGroup) XXX_Size() int {
	return xxx_messageInfo_AdGroup.Size(m)
}
func (m *AdGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroup.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroup proto.InternalMessageInfo

func (m *AdGroup) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *AdGroup) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *AdGroup) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *AdGroup) GetStatus() enums.AdGroupStatusEnum_AdGroupStatus {
	if m != nil {
		return m.Status
	}
	return enums.AdGroupStatusEnum_UNSPECIFIED
}

func (m *AdGroup) GetType() enums.AdGroupTypeEnum_AdGroupType {
	if m != nil {
		return m.Type
	}
	return enums.AdGroupTypeEnum_UNSPECIFIED
}

func (m *AdGroup) GetAdRotationMode() enums.AdGroupAdRotationModeEnum_AdGroupAdRotationMode {
	if m != nil {
		return m.AdRotationMode
	}
	return enums.AdGroupAdRotationModeEnum_UNSPECIFIED
}

func (m *AdGroup) GetBaseAdGroup() *wrappers.StringValue {
	if m != nil {
		return m.BaseAdGroup
	}
	return nil
}

func (m *AdGroup) GetTrackingUrlTemplate() *wrappers.StringValue {
	if m != nil {
		return m.TrackingUrlTemplate
	}
	return nil
}

func (m *AdGroup) GetUrlCustomParameters() []*common.CustomParameter {
	if m != nil {
		return m.UrlCustomParameters
	}
	return nil
}

func (m *AdGroup) GetCampaign() *wrappers.StringValue {
	if m != nil {
		return m.Campaign
	}
	return nil
}

func (m *AdGroup) GetCpcBidMicros() *wrappers.Int64Value {
	if m != nil {
		return m.CpcBidMicros
	}
	return nil
}

func (m *AdGroup) GetCpmBidMicros() *wrappers.Int64Value {
	if m != nil {
		return m.CpmBidMicros
	}
	return nil
}

func (m *AdGroup) GetTargetCpaMicros() *wrappers.Int64Value {
	if m != nil {
		return m.TargetCpaMicros
	}
	return nil
}

func (m *AdGroup) GetCpvBidMicros() *wrappers.Int64Value {
	if m != nil {
		return m.CpvBidMicros
	}
	return nil
}

func (m *AdGroup) GetTargetCpmMicros() *wrappers.Int64Value {
	if m != nil {
		return m.TargetCpmMicros
	}
	return nil
}

func (m *AdGroup) GetTargetRoas() *wrappers.DoubleValue {
	if m != nil {
		return m.TargetRoas
	}
	return nil
}

func (m *AdGroup) GetPercentCpcBidMicros() *wrappers.Int64Value {
	if m != nil {
		return m.PercentCpcBidMicros
	}
	return nil
}

func (m *AdGroup) GetExplorerAutoOptimizerSetting() *common.ExplorerAutoOptimizerSetting {
	if m != nil {
		return m.ExplorerAutoOptimizerSetting
	}
	return nil
}

func (m *AdGroup) GetDisplayCustomBidDimension() enums.TargetingDimensionEnum_TargetingDimension {
	if m != nil {
		return m.DisplayCustomBidDimension
	}
	return enums.TargetingDimensionEnum_UNSPECIFIED
}

func (m *AdGroup) GetFinalUrlSuffix() *wrappers.StringValue {
	if m != nil {
		return m.FinalUrlSuffix
	}
	return nil
}

func (m *AdGroup) GetTargetingSetting() *common.TargetingSetting {
	if m != nil {
		return m.TargetingSetting
	}
	return nil
}

func (m *AdGroup) GetEffectiveTargetCpaMicros() *wrappers.Int64Value {
	if m != nil {
		return m.EffectiveTargetCpaMicros
	}
	return nil
}

func (m *AdGroup) GetEffectiveTargetCpaSource() enums.BiddingSourceEnum_BiddingSource {
	if m != nil {
		return m.EffectiveTargetCpaSource
	}
	return enums.BiddingSourceEnum_UNSPECIFIED
}

func (m *AdGroup) GetEffectiveTargetRoas() *wrappers.DoubleValue {
	if m != nil {
		return m.EffectiveTargetRoas
	}
	return nil
}

func (m *AdGroup) GetEffectiveTargetRoasSource() enums.BiddingSourceEnum_BiddingSource {
	if m != nil {
		return m.EffectiveTargetRoasSource
	}
	return enums.BiddingSourceEnum_UNSPECIFIED
}

func (m *AdGroup) GetLabels() []*wrappers.StringValue {
	if m != nil {
		return m.Labels
	}
	return nil
}

func init() {
	proto.RegisterType((*AdGroup)(nil), "google.ads.googleads.v2.resources.AdGroup")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/ad_group.proto", fileDescriptor_a1544ca6322174f1)
}

var fileDescriptor_a1544ca6322174f1 = []byte{
	// 984 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x96, 0xdd, 0x6e, 0x23, 0x35,
	0x14, 0xc7, 0x95, 0xb6, 0x14, 0x70, 0xdb, 0x6c, 0x77, 0x4a, 0x97, 0x69, 0x1b, 0x76, 0x53, 0xd0,
	0x4a, 0x95, 0x90, 0x26, 0x25, 0xbb, 0x2c, 0x28, 0xb0, 0x88, 0xa4, 0x5d, 0x0a, 0x48, 0x5b, 0xa2,
	0x49, 0x37, 0x17, 0xab, 0xa2, 0x91, 0x33, 0x76, 0x46, 0x16, 0x33, 0x63, 0xcb, 0xf6, 0x84, 0x16,
	0x09, 0x71, 0xc1, 0x15, 0xaf, 0xc1, 0x25, 0x8f, 0xc2, 0xa3, 0xc0, 0x13, 0x70, 0x87, 0xc6, 0x1f,
	0xd3, 0x24, 0x6d, 0x32, 0x83, 0xc4, 0x9d, 0xc7, 0x3e, 0xff, 0x9f, 0x8f, 0x8f, 0xcf, 0x39, 0x1e,
	0x70, 0x1c, 0x51, 0x1a, 0xc5, 0xb8, 0x05, 0x91, 0x68, 0xe9, 0x61, 0x3e, 0x9a, 0xb4, 0x5b, 0x1c,
	0x0b, 0x9a, 0xf1, 0x10, 0x8b, 0x16, 0x44, 0x41, 0xc4, 0x69, 0xc6, 0x3c, 0xc6, 0xa9, 0xa4, 0xce,
	0xa1, 0x36, 0xf3, 0x20, 0x12, 0x5e, 0xa1, 0xf0, 0x26, 0x6d, 0xaf, 0x50, 0xec, 0x7f, 0xbc, 0x08,
	0x1a, 0xd2, 0x24, 0xa1, 0x69, 0x2b, 0xcc, 0x84, 0xa4, 0x49, 0xc0, 0x20, 0x87, 0x09, 0x96, 0x98,
	0x6b, 0xf2, 0xfe, 0x69, 0x89, 0x0c, 0x5f, 0xb1, 0x98, 0x72, 0xcc, 0x03, 0x98, 0x49, 0x1a, 0x50,
	0x26, 0x49, 0x42, 0x7e, 0xc2, 0x3c, 0x10, 0x58, 0x4a, 0x92, 0x46, 0x86, 0xf2, 0xac, 0x84, 0x22,
	0x21, 0x8f, 0x70, 0x6e, 0x3f, 0xa7, 0x7b, 0xbe, 0x48, 0x87, 0xd3, 0x2c, 0xb9, 0x89, 0x42, 0x00,
	0x51, 0xc0, 0xa9, 0x84, 0x92, 0xd0, 0x34, 0x48, 0x28, 0xc2, 0x46, 0xfe, 0xa4, 0xa2, 0x5c, 0x48,
	0x28, 0x33, 0x61, 0x44, 0x1f, 0x55, 0x14, 0xc9, 0x6b, 0x66, 0xf7, 0x69, 0x2f, 0x97, 0x8c, 0x08,
	0x42, 0xea, 0x6c, 0xea, 0x2a, 0x8c, 0xe6, 0x93, 0xe5, 0x9a, 0x9b, 0x88, 0x20, 0x92, 0xe0, 0x54,
	0x10, 0x9a, 0x1a, 0xe1, 0x43, 0x23, 0x54, 0x5f, 0xa3, 0x6c, 0xdc, 0xfa, 0x91, 0x43, 0xc6, 0x30,
	0xb7, 0xfe, 0x37, 0x2c, 0x98, 0x91, 0x16, 0x4c, 0x53, 0x13, 0x16, 0xb3, 0xfa, 0xfe, 0xdf, 0xdb,
	0xe0, 0xcd, 0x2e, 0x3a, 0xcb, 0x4f, 0xe0, 0x7c, 0x00, 0xb6, 0x6c, 0x7e, 0x04, 0x29, 0x4c, 0xb0,
	0x5b, 0x6b, 0xd6, 0x8e, 0xde, 0xf6, 0x37, 0xed, 0xe4, 0x39, 0x4c, 0xb0, 0xf3, 0x21, 0x58, 0x21,
	0xc8, 0x5d, 0x6d, 0xd6, 0x8e, 0x36, 0xda, 0x07, 0x26, 0xb9, 0x3c, 0xbb, 0xb7, 0xf7, 0x4d, 0x2a,
	0x9f, 0x3d, 0x1d, 0xc2, 0x38, 0xc3, 0xfe, 0x0a, 0x41, 0xce, 0x31, 0x58, 0x53, 0xa0, 0x35, 0x65,
	0xde, 0xb8, 0x65, 0x3e, 0x90, 0x9c, 0xa4, 0x91, 0xb6, 0x57, 0x96, 0xce, 0x10, 0xac, 0xeb, 0xe8,
	0xbb, 0x6f, 0x34, 0x6b, 0x47, 0xf5, 0xf6, 0x17, 0xde, 0xa2, 0x54, 0x56, 0x71, 0xf1, 0x8c, 0xef,
	0x03, 0xa5, 0x79, 0x91, 0x66, 0xc9, 0xec, 0x8c, 0x6f, 0x68, 0xce, 0x39, 0x58, 0xcb, 0x2f, 0xc8,
	0xdd, 0x54, 0xd4, 0x4e, 0x35, 0xea, 0xc5, 0x35, 0xc3, 0xd3, 0xcc, 0xfc, 0xdb, 0x57, 0x1c, 0xe7,
	0x0a, 0x6c, 0xcf, 0x27, 0x99, 0xfb, 0x40, 0xb1, 0xcf, 0xab, 0xb1, 0xbb, 0xc8, 0x37, 0xe2, 0x97,
	0x14, 0xcd, 0xec, 0x32, 0xbb, 0xe2, 0xd7, 0xe1, 0xcc, 0xb7, 0xf3, 0x25, 0xd8, 0x1a, 0x41, 0x81,
	0x03, 0x9b, 0x78, 0xae, 0x53, 0x21, 0xb8, 0x1b, 0xb9, 0xc4, 0xde, 0x73, 0x1f, 0xec, 0x4a, 0x0e,
	0xc3, 0x1f, 0xf2, 0x6c, 0xca, 0x78, 0x1c, 0x48, 0x9c, 0xb0, 0x18, 0x4a, 0xec, 0x6e, 0x55, 0x20,
	0xed, 0x58, 0xe9, 0x2b, 0x1e, 0x5f, 0x18, 0xa1, 0x13, 0x82, 0xdd, 0x1c, 0x34, 0xdf, 0x33, 0x84,
	0xbb, 0xde, 0x5c, 0x3d, 0xda, 0x68, 0xb7, 0x16, 0x86, 0x44, 0xd7, 0xbb, 0x77, 0xa2, 0x84, 0x7d,
	0xab, 0xf3, 0x77, 0x32, 0x1e, 0xcf, 0xcd, 0x09, 0xe7, 0x53, 0xf0, 0x56, 0x08, 0x13, 0x06, 0x49,
	0x94, 0xba, 0xa0, 0x82, 0xa7, 0x85, 0xb5, 0xd3, 0x05, 0xf5, 0x90, 0x85, 0xc1, 0x88, 0xa0, 0x20,
	0x21, 0x21, 0xa7, 0xc2, 0xad, 0x97, 0xe7, 0xef, 0x66, 0xc8, 0xc2, 0x1e, 0x41, 0x2f, 0x95, 0x40,
	0x23, 0x92, 0x69, 0xc4, 0xbd, 0x4a, 0x88, 0xe4, 0x06, 0x71, 0x06, 0xee, 0xeb, 0x2a, 0x0e, 0x42,
	0x06, 0x2d, 0xe5, 0xa0, 0x9c, 0x72, 0x4f, 0xab, 0x4e, 0x18, 0x9c, 0xf6, 0x65, 0x32, 0xed, 0xcb,
	0xfd, 0x4a, 0xbe, 0x4c, 0xee, 0xf6, 0x25, 0xb1, 0x94, 0xfd, 0xff, 0xe0, 0x4b, 0x62, 0x40, 0xcf,
	0xc1, 0x86, 0x01, 0x71, 0x0a, 0x85, 0xfb, 0x70, 0xc1, 0xbd, 0x9c, 0xd2, 0x6c, 0x14, 0x63, 0xcd,
	0x00, 0x5a, 0xe0, 0x53, 0x28, 0x9c, 0x3e, 0x78, 0xc0, 0x30, 0x0f, 0x71, 0x9a, 0x3b, 0x32, 0x73,
	0x43, 0xef, 0x94, 0x3b, 0xb3, 0x63, 0xa4, 0x27, 0xd3, 0x17, 0xf5, 0x6b, 0x0d, 0x3c, 0x2a, 0x79,
	0x84, 0xdc, 0x5d, 0xc5, 0xfe, 0xbc, 0x2c, 0x2b, 0x5f, 0x18, 0x4c, 0x37, 0x93, 0xf4, 0x3b, 0x0b,
	0x19, 0x68, 0x86, 0xdf, 0xc0, 0x4b, 0x56, 0x9d, 0xdf, 0x6a, 0xa0, 0x81, 0x88, 0x60, 0x31, 0xbc,
	0xb6, 0x55, 0x91, 0x9f, 0xad, 0xe8, 0xdd, 0xee, 0xbb, 0xaa, 0x57, 0x7c, 0x5d, 0xd2, 0x2b, 0x2e,
	0x6c, 0xd7, 0x3f, 0xb5, 0x42, 0xd5, 0x28, 0x6e, 0x4f, 0xfb, 0x7b, 0x66, 0x37, 0x5d, 0x35, 0x3d,
	0x82, 0x8a, 0x25, 0xe7, 0x2b, 0xb0, 0x3d, 0x26, 0x29, 0x8c, 0x55, 0xad, 0x8b, 0x6c, 0x3c, 0x26,
	0x57, 0xae, 0x5b, 0xa1, 0x7e, 0xea, 0x4a, 0xf5, 0x8a, 0xc7, 0x03, 0xa5, 0x71, 0xbe, 0xb7, 0x39,
	0x33, 0xf5, 0x2e, 0xbb, 0x7b, 0x0a, 0x74, 0x5c, 0x16, 0xca, 0xc2, 0x63, 0x1b, 0xbe, 0x6d, 0x39,
	0x37, 0xe3, 0xbc, 0x06, 0x07, 0x78, 0x3c, 0xc6, 0xa1, 0x24, 0x13, 0x1c, 0xdc, 0x2e, 0x94, 0x46,
	0x79, 0x3e, 0xb8, 0x85, 0xfe, 0x62, 0xae, 0x62, 0x7e, 0x5e, 0xc0, 0xd6, 0xef, 0x9a, 0xfb, 0x5e,
	0xa5, 0xa7, 0xa6, 0xa7, 0x9f, 0xed, 0x81, 0xd2, 0xa8, 0x7b, 0x98, 0x99, 0xb9, 0x6b, 0x7b, 0xbd,
	0x92, 0x37, 0xdc, 0x5b, 0xdb, 0xab, 0x72, 0x79, 0x54, 0xa1, 0x5c, 0x76, 0xe6, 0xb0, 0xaa, 0x6e,
	0x7e, 0x01, 0x8d, 0x3b, 0x89, 0xf6, 0x44, 0xcd, 0xff, 0xe5, 0x44, 0x7b, 0x77, 0x6c, 0x6d, 0x8e,
	0xf4, 0x14, 0xac, 0xc7, 0x70, 0x84, 0x63, 0xe1, 0x1e, 0xaa, 0x16, 0xbf, 0x3c, 0x95, 0x8c, 0x6d,
	0xef, 0x9f, 0x1a, 0x78, 0x1c, 0xd2, 0xc4, 0x2b, 0xfd, 0x3d, 0xed, 0x6d, 0x9a, 0xc7, 0xaa, 0x9f,
	0xe3, 0xfa, 0xb5, 0xd7, 0xdf, 0x1a, 0x49, 0x44, 0x63, 0x98, 0x46, 0x1e, 0xe5, 0x51, 0x2b, 0xc2,
	0xa9, 0xda, 0xcc, 0xfe, 0x2e, 0x31, 0x22, 0x96, 0xfc, 0x22, 0x7f, 0x56, 0x8c, 0x7e, 0x5f, 0x59,
	0x3d, 0xeb, 0x76, 0xff, 0x58, 0x39, 0x3c, 0xd3, 0xc8, 0x2e, 0x12, 0x9e, 0x1e, 0xe6, 0xa3, 0x61,
	0xdb, 0xf3, 0xad, 0xe5, 0x9f, 0xd6, 0xe6, 0xb2, 0x8b, 0xc4, 0x65, 0x61, 0x73, 0x39, 0x6c, 0x5f,
	0x16, 0x36, 0x7f, 0xad, 0x3c, 0xd6, 0x0b, 0x9d, 0x4e, 0x17, 0x89, 0x4e, 0xa7, 0xb0, 0xea, 0x74,
	0x86, 0xed, 0x4e, 0xa7, 0xb0, 0x1b, 0xad, 0x2b, 0x67, 0x9f, 0xfc, 0x1b, 0x00, 0x00, 0xff, 0xff,
	0x53, 0xd1, 0xe8, 0x1c, 0xce, 0x0b, 0x00, 0x00,
}
