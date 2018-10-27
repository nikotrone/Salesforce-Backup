package SalesforceWSDL

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

type SObject struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com sObject"`

	FieldsToNull []string `xml:"fieldsToNull,omitempty"`
	Id           *ID      `xml:"Id,omitempty"`
}

type AggregateResult struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com AggregateResult"`

	*SObject
}

type AcceptedEventRelation struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com AcceptedEventRelation"`

	*SObject

	CreatedBy        *User     `xml:"CreatedBy,omitempty"`
	CreatedById      *ID       `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time `xml:"CreatedDate,omitempty"`
	Event            *Event    `xml:"Event,omitempty"`
	EventId          *ID       `xml:"EventId,omitempty"`
	IsDeleted        bool      `xml:"IsDeleted,omitempty"`
	LastModifiedBy   *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time `xml:"LastModifiedDate,omitempty"`
	Relation         *SObject  `xml:"Relation,omitempty"`
	RelationId       *ID       `xml:"RelationId,omitempty"`
	RespondedDate    time.Time `xml:"RespondedDate,omitempty"`
	Response         string    `xml:"Response,omitempty"`
	SystemModstamp   time.Time `xml:"SystemModstamp,omitempty"`
	Type             string    `xml:"Type,omitempty"`
}

type Account struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com Account"`

	*SObject

	AccountCleanInfos          *QueryResult      `xml:"AccountCleanInfos,omitempty"`
	AccountContactRoles        *QueryResult      `xml:"AccountContactRoles,omitempty"`
	AccountNumber              string            `xml:"AccountNumber,omitempty"`
	AccountPartnersFrom        *QueryResult      `xml:"AccountPartnersFrom,omitempty"`
	AccountPartnersTo          *QueryResult      `xml:"AccountPartnersTo,omitempty"`
	AccountSource              string            `xml:"AccountSource,omitempty"`
	Activec                    string            `xml:"Active__c,omitempty"`
	ActivityHistories          *QueryResult      `xml:"ActivityHistories,omitempty"`
	AnnualRevenue              float64           `xml:"AnnualRevenue,omitempty"`
	Assets                     *QueryResult      `xml:"Assets,omitempty"`
	AttachedContentDocuments   *QueryResult      `xml:"AttachedContentDocuments,omitempty"`
	Attachments                *QueryResult      `xml:"Attachments,omitempty"`
	BillingAddress             *Address          `xml:"BillingAddress,omitempty"`
	BillingCity                string            `xml:"BillingCity,omitempty"`
	BillingCountry             string            `xml:"BillingCountry,omitempty"`
	BillingGeocodeAccuracy     string            `xml:"BillingGeocodeAccuracy,omitempty"`
	BillingLatitude            float64           `xml:"BillingLatitude,omitempty"`
	BillingLongitude           float64           `xml:"BillingLongitude,omitempty"`
	BillingPostalCode          string            `xml:"BillingPostalCode,omitempty"`
	BillingState               string            `xml:"BillingState,omitempty"`
	BillingStreet              string            `xml:"BillingStreet,omitempty"`
	Cases                      *QueryResult      `xml:"Cases,omitempty"`
	ChildAccounts              *QueryResult      `xml:"ChildAccounts,omitempty"`
	CleanStatus                string            `xml:"CleanStatus,omitempty"`
	CombinedAttachments        *QueryResult      `xml:"CombinedAttachments,omitempty"`
	Contacts                   *QueryResult      `xml:"Contacts,omitempty"`
	ContentDocumentLinks       *QueryResult      `xml:"ContentDocumentLinks,omitempty"`
	Contracts                  *QueryResult      `xml:"Contracts,omitempty"`
	CreatedBy                  *User             `xml:"CreatedBy,omitempty"`
	CreatedById                *ID               `xml:"CreatedById,omitempty"`
	CreatedDate                time.Time         `xml:"CreatedDate,omitempty"`
	CustomerPriorityc          string            `xml:"CustomerPriority__c,omitempty"`
	DandbCompany               *DandBCompany     `xml:"DandbCompany,omitempty"`
	DandbCompanyId             *ID               `xml:"DandbCompanyId,omitempty"`
	Description                string            `xml:"Description,omitempty"`
	DunsNumber                 string            `xml:"DunsNumber,omitempty"`
	DuplicateRecordItems       *QueryResult      `xml:"DuplicateRecordItems,omitempty"`
	Emails                     *QueryResult      `xml:"Emails,omitempty"`
	Events                     *QueryResult      `xml:"Events,omitempty"`
	Fax                        string            `xml:"Fax,omitempty"`
	FeedSubscriptionsForEntity *QueryResult      `xml:"FeedSubscriptionsForEntity,omitempty"`
	Feeds                      *QueryResult      `xml:"Feeds,omitempty"`
	Histories                  *QueryResult      `xml:"Histories,omitempty"`
	Industry                   string            `xml:"Industry,omitempty"`
	IsDeleted                  bool              `xml:"IsDeleted,omitempty"`
	Jigsaw                     string            `xml:"Jigsaw,omitempty"`
	JigsawCompanyId            string            `xml:"JigsawCompanyId,omitempty"`
	LastActivityDate           time.Time         `xml:"LastActivityDate,omitempty"`
	LastModifiedBy             *User             `xml:"LastModifiedBy,omitempty"`
	LastModifiedById           *ID               `xml:"LastModifiedById,omitempty"`
	LastModifiedDate           time.Time         `xml:"LastModifiedDate,omitempty"`
	LastReferencedDate         time.Time         `xml:"LastReferencedDate,omitempty"`
	LastViewedDate             time.Time         `xml:"LastViewedDate,omitempty"`
	LookedUpFromActivities     *QueryResult      `xml:"LookedUpFromActivities,omitempty"`
	MasterRecord               *Account          `xml:"MasterRecord,omitempty"`
	MasterRecordId             *ID               `xml:"MasterRecordId,omitempty"`
	NaicsCode                  string            `xml:"NaicsCode,omitempty"`
	NaicsDesc                  string            `xml:"NaicsDesc,omitempty"`
	Name                       string            `xml:"Name,omitempty"`
	Notes                      *QueryResult      `xml:"Notes,omitempty"`
	NotesAndAttachments        *QueryResult      `xml:"NotesAndAttachments,omitempty"`
	NumberOfEmployees          int32             `xml:"NumberOfEmployees,omitempty"`
	NumberofLocationsc         float64           `xml:"NumberofLocations__c,omitempty"`
	OpenActivities             *QueryResult      `xml:"OpenActivities,omitempty"`
	Opportunities              *QueryResult      `xml:"Opportunities,omitempty"`
	OpportunityPartnersTo      *QueryResult      `xml:"OpportunityPartnersTo,omitempty"`
	Orders                     *QueryResult      `xml:"Orders,omitempty"`
	Owner                      *User             `xml:"Owner,omitempty"`
	OwnerId                    *ID               `xml:"OwnerId,omitempty"`
	Ownership                  string            `xml:"Ownership,omitempty"`
	Parent                     *Account          `xml:"Parent,omitempty"`
	ParentId                   *ID               `xml:"ParentId,omitempty"`
	PartnersFrom               *QueryResult      `xml:"PartnersFrom,omitempty"`
	PartnersTo                 *QueryResult      `xml:"PartnersTo,omitempty"`
	Phone                      string            `xml:"Phone,omitempty"`
	PhotoUrl                   string            `xml:"PhotoUrl,omitempty"`
	ProcessInstances           *QueryResult      `xml:"ProcessInstances,omitempty"`
	ProcessSteps               *QueryResult      `xml:"ProcessSteps,omitempty"`
	ProvidedAssets             *QueryResult      `xml:"ProvidedAssets,omitempty"`
	Rating                     string            `xml:"Rating,omitempty"`
	RecordActionHistories      *QueryResult      `xml:"RecordActionHistories,omitempty"`
	RecordActions              *QueryResult      `xml:"RecordActions,omitempty"`
	RecordAssociatedGroups     *QueryResult      `xml:"RecordAssociatedGroups,omitempty"`
	SLAExpirationDatec         time.Time         `xml:"SLAExpirationDate__c,omitempty"`
	SLASerialNumberc           string            `xml:"SLASerialNumber__c,omitempty"`
	SLAc                       string            `xml:"SLA__c,omitempty"`
	ServicedAssets             *QueryResult      `xml:"ServicedAssets,omitempty"`
	Shares                     *QueryResult      `xml:"Shares,omitempty"`
	ShippingAddress            *Address          `xml:"ShippingAddress,omitempty"`
	ShippingCity               string            `xml:"ShippingCity,omitempty"`
	ShippingCountry            string            `xml:"ShippingCountry,omitempty"`
	ShippingGeocodeAccuracy    string            `xml:"ShippingGeocodeAccuracy,omitempty"`
	ShippingLatitude           float64           `xml:"ShippingLatitude,omitempty"`
	ShippingLongitude          float64           `xml:"ShippingLongitude,omitempty"`
	ShippingPostalCode         string            `xml:"ShippingPostalCode,omitempty"`
	ShippingState              string            `xml:"ShippingState,omitempty"`
	ShippingStreet             string            `xml:"ShippingStreet,omitempty"`
	Sic                        string            `xml:"Sic,omitempty"`
	SicDesc                    string            `xml:"SicDesc,omitempty"`
	Site                       string            `xml:"Site,omitempty"`
	SystemModstamp             time.Time         `xml:"SystemModstamp,omitempty"`
	Tasks                      *QueryResult      `xml:"Tasks,omitempty"`
	TickerSymbol               string            `xml:"TickerSymbol,omitempty"`
	TopicAssignments           *QueryResult      `xml:"TopicAssignments,omitempty"`
	Tradestyle                 string            `xml:"Tradestyle,omitempty"`
	Type                       string            `xml:"Type,omitempty"`
	UpsellOpportunityc         string            `xml:"UpsellOpportunity__c,omitempty"`
	UserRecordAccess           *UserRecordAccess `xml:"UserRecordAccess,omitempty"`
	Users                      *QueryResult      `xml:"Users,omitempty"`
	Website                    string            `xml:"Website,omitempty"`
	YearStarted                string            `xml:"YearStarted,omitempty"`
}

type AccountCleanInfo struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com AccountCleanInfo"`

	*SObject

	Account                           *Account          `xml:"Account,omitempty"`
	AccountId                         *ID               `xml:"AccountId,omitempty"`
	AccountSite                       string            `xml:"AccountSite,omitempty"`
	Address                           *Address          `xml:"Address,omitempty"`
	AnnualRevenue                     float64           `xml:"AnnualRevenue,omitempty"`
	City                              string            `xml:"City,omitempty"`
	CleanedByJob                      bool              `xml:"CleanedByJob,omitempty"`
	CleanedByUser                     bool              `xml:"CleanedByUser,omitempty"`
	CompanyName                       string            `xml:"CompanyName,omitempty"`
	CompanyStatusDataDotCom           string            `xml:"CompanyStatusDataDotCom,omitempty"`
	Country                           string            `xml:"Country,omitempty"`
	CreatedBy                         *User             `xml:"CreatedBy,omitempty"`
	CreatedById                       *ID               `xml:"CreatedById,omitempty"`
	CreatedDate                       time.Time         `xml:"CreatedDate,omitempty"`
	DandBCompanyDunsNumber            string            `xml:"DandBCompanyDunsNumber,omitempty"`
	DataDotComId                      string            `xml:"DataDotComId,omitempty"`
	Description                       string            `xml:"Description,omitempty"`
	DunsNumber                        string            `xml:"DunsNumber,omitempty"`
	DunsRightMatchConfidence          int32             `xml:"DunsRightMatchConfidence,omitempty"`
	DunsRightMatchGrade               string            `xml:"DunsRightMatchGrade,omitempty"`
	Fax                               string            `xml:"Fax,omitempty"`
	GeocodeAccuracy                   string            `xml:"GeocodeAccuracy,omitempty"`
	Industry                          string            `xml:"Industry,omitempty"`
	IsDeleted                         bool              `xml:"IsDeleted,omitempty"`
	IsDifferentAccountSite            bool              `xml:"IsDifferentAccountSite,omitempty"`
	IsDifferentAnnualRevenue          bool              `xml:"IsDifferentAnnualRevenue,omitempty"`
	IsDifferentCity                   bool              `xml:"IsDifferentCity,omitempty"`
	IsDifferentCompanyName            bool              `xml:"IsDifferentCompanyName,omitempty"`
	IsDifferentCountry                bool              `xml:"IsDifferentCountry,omitempty"`
	IsDifferentCountryCode            bool              `xml:"IsDifferentCountryCode,omitempty"`
	IsDifferentDandBCompanyDunsNumber bool              `xml:"IsDifferentDandBCompanyDunsNumber,omitempty"`
	IsDifferentDescription            bool              `xml:"IsDifferentDescription,omitempty"`
	IsDifferentDunsNumber             bool              `xml:"IsDifferentDunsNumber,omitempty"`
	IsDifferentFax                    bool              `xml:"IsDifferentFax,omitempty"`
	IsDifferentIndustry               bool              `xml:"IsDifferentIndustry,omitempty"`
	IsDifferentNaicsCode              bool              `xml:"IsDifferentNaicsCode,omitempty"`
	IsDifferentNaicsDescription       bool              `xml:"IsDifferentNaicsDescription,omitempty"`
	IsDifferentNumberOfEmployees      bool              `xml:"IsDifferentNumberOfEmployees,omitempty"`
	IsDifferentOwnership              bool              `xml:"IsDifferentOwnership,omitempty"`
	IsDifferentPhone                  bool              `xml:"IsDifferentPhone,omitempty"`
	IsDifferentPostalCode             bool              `xml:"IsDifferentPostalCode,omitempty"`
	IsDifferentSic                    bool              `xml:"IsDifferentSic,omitempty"`
	IsDifferentSicDescription         bool              `xml:"IsDifferentSicDescription,omitempty"`
	IsDifferentState                  bool              `xml:"IsDifferentState,omitempty"`
	IsDifferentStateCode              bool              `xml:"IsDifferentStateCode,omitempty"`
	IsDifferentStreet                 bool              `xml:"IsDifferentStreet,omitempty"`
	IsDifferentTickerSymbol           bool              `xml:"IsDifferentTickerSymbol,omitempty"`
	IsDifferentTradestyle             bool              `xml:"IsDifferentTradestyle,omitempty"`
	IsDifferentWebsite                bool              `xml:"IsDifferentWebsite,omitempty"`
	IsDifferentYearStarted            bool              `xml:"IsDifferentYearStarted,omitempty"`
	IsFlaggedWrongAccountSite         bool              `xml:"IsFlaggedWrongAccountSite,omitempty"`
	IsFlaggedWrongAddress             bool              `xml:"IsFlaggedWrongAddress,omitempty"`
	IsFlaggedWrongAnnualRevenue       bool              `xml:"IsFlaggedWrongAnnualRevenue,omitempty"`
	IsFlaggedWrongCompanyName         bool              `xml:"IsFlaggedWrongCompanyName,omitempty"`
	IsFlaggedWrongDescription         bool              `xml:"IsFlaggedWrongDescription,omitempty"`
	IsFlaggedWrongDunsNumber          bool              `xml:"IsFlaggedWrongDunsNumber,omitempty"`
	IsFlaggedWrongFax                 bool              `xml:"IsFlaggedWrongFax,omitempty"`
	IsFlaggedWrongIndustry            bool              `xml:"IsFlaggedWrongIndustry,omitempty"`
	IsFlaggedWrongNaicsCode           bool              `xml:"IsFlaggedWrongNaicsCode,omitempty"`
	IsFlaggedWrongNaicsDescription    bool              `xml:"IsFlaggedWrongNaicsDescription,omitempty"`
	IsFlaggedWrongNumberOfEmployees   bool              `xml:"IsFlaggedWrongNumberOfEmployees,omitempty"`
	IsFlaggedWrongOwnership           bool              `xml:"IsFlaggedWrongOwnership,omitempty"`
	IsFlaggedWrongPhone               bool              `xml:"IsFlaggedWrongPhone,omitempty"`
	IsFlaggedWrongSic                 bool              `xml:"IsFlaggedWrongSic,omitempty"`
	IsFlaggedWrongSicDescription      bool              `xml:"IsFlaggedWrongSicDescription,omitempty"`
	IsFlaggedWrongTickerSymbol        bool              `xml:"IsFlaggedWrongTickerSymbol,omitempty"`
	IsFlaggedWrongTradestyle          bool              `xml:"IsFlaggedWrongTradestyle,omitempty"`
	IsFlaggedWrongWebsite             bool              `xml:"IsFlaggedWrongWebsite,omitempty"`
	IsFlaggedWrongYearStarted         bool              `xml:"IsFlaggedWrongYearStarted,omitempty"`
	IsInactive                        bool              `xml:"IsInactive,omitempty"`
	IsReviewedAccountSite             bool              `xml:"IsReviewedAccountSite,omitempty"`
	IsReviewedAddress                 bool              `xml:"IsReviewedAddress,omitempty"`
	IsReviewedAnnualRevenue           bool              `xml:"IsReviewedAnnualRevenue,omitempty"`
	IsReviewedCompanyName             bool              `xml:"IsReviewedCompanyName,omitempty"`
	IsReviewedDandBCompanyDunsNumber  bool              `xml:"IsReviewedDandBCompanyDunsNumber,omitempty"`
	IsReviewedDescription             bool              `xml:"IsReviewedDescription,omitempty"`
	IsReviewedDunsNumber              bool              `xml:"IsReviewedDunsNumber,omitempty"`
	IsReviewedFax                     bool              `xml:"IsReviewedFax,omitempty"`
	IsReviewedIndustry                bool              `xml:"IsReviewedIndustry,omitempty"`
	IsReviewedNaicsCode               bool              `xml:"IsReviewedNaicsCode,omitempty"`
	IsReviewedNaicsDescription        bool              `xml:"IsReviewedNaicsDescription,omitempty"`
	IsReviewedNumberOfEmployees       bool              `xml:"IsReviewedNumberOfEmployees,omitempty"`
	IsReviewedOwnership               bool              `xml:"IsReviewedOwnership,omitempty"`
	IsReviewedPhone                   bool              `xml:"IsReviewedPhone,omitempty"`
	IsReviewedSic                     bool              `xml:"IsReviewedSic,omitempty"`
	IsReviewedSicDescription          bool              `xml:"IsReviewedSicDescription,omitempty"`
	IsReviewedTickerSymbol            bool              `xml:"IsReviewedTickerSymbol,omitempty"`
	IsReviewedTradestyle              bool              `xml:"IsReviewedTradestyle,omitempty"`
	IsReviewedWebsite                 bool              `xml:"IsReviewedWebsite,omitempty"`
	IsReviewedYearStarted             bool              `xml:"IsReviewedYearStarted,omitempty"`
	LastMatchedDate                   time.Time         `xml:"LastMatchedDate,omitempty"`
	LastModifiedBy                    *User             `xml:"LastModifiedBy,omitempty"`
	LastModifiedById                  *ID               `xml:"LastModifiedById,omitempty"`
	LastModifiedDate                  time.Time         `xml:"LastModifiedDate,omitempty"`
	LastStatusChangedBy               *User             `xml:"LastStatusChangedBy,omitempty"`
	LastStatusChangedById             *ID               `xml:"LastStatusChangedById,omitempty"`
	LastStatusChangedDate             time.Time         `xml:"LastStatusChangedDate,omitempty"`
	Latitude                          float64           `xml:"Latitude,omitempty"`
	Longitude                         float64           `xml:"Longitude,omitempty"`
	NaicsCode                         string            `xml:"NaicsCode,omitempty"`
	NaicsDescription                  string            `xml:"NaicsDescription,omitempty"`
	Name                              string            `xml:"Name,omitempty"`
	NumberOfEmployees                 int32             `xml:"NumberOfEmployees,omitempty"`
	Ownership                         string            `xml:"Ownership,omitempty"`
	Phone                             string            `xml:"Phone,omitempty"`
	PostalCode                        string            `xml:"PostalCode,omitempty"`
	Sic                               string            `xml:"Sic,omitempty"`
	SicDescription                    string            `xml:"SicDescription,omitempty"`
	State                             string            `xml:"State,omitempty"`
	Street                            string            `xml:"Street,omitempty"`
	SystemModstamp                    time.Time         `xml:"SystemModstamp,omitempty"`
	TickerSymbol                      string            `xml:"TickerSymbol,omitempty"`
	Tradestyle                        string            `xml:"Tradestyle,omitempty"`
	UserRecordAccess                  *UserRecordAccess `xml:"UserRecordAccess,omitempty"`
	Website                           string            `xml:"Website,omitempty"`
	YearStarted                       string            `xml:"YearStarted,omitempty"`
}

type AccountContactRole struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com AccountContactRole"`

	*SObject

	Account          *Account  `xml:"Account,omitempty"`
	AccountId        *ID       `xml:"AccountId,omitempty"`
	Contact          *Contact  `xml:"Contact,omitempty"`
	ContactId        *ID       `xml:"ContactId,omitempty"`
	CreatedBy        *User     `xml:"CreatedBy,omitempty"`
	CreatedById      *ID       `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time `xml:"CreatedDate,omitempty"`
	IsDeleted        bool      `xml:"IsDeleted,omitempty"`
	IsPrimary        bool      `xml:"IsPrimary,omitempty"`
	LastModifiedBy   *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time `xml:"LastModifiedDate,omitempty"`
	Role             string    `xml:"Role,omitempty"`
	SystemModstamp   time.Time `xml:"SystemModstamp,omitempty"`
}

type AccountFeed struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com AccountFeed"`

	*SObject

	BestComment        *FeedComment `xml:"BestComment,omitempty"`
	BestCommentId      *ID          `xml:"BestCommentId,omitempty"`
	Body               string       `xml:"Body,omitempty"`
	CommentCount       int32        `xml:"CommentCount,omitempty"`
	CreatedBy          *SObject     `xml:"CreatedBy,omitempty"`
	CreatedById        *ID          `xml:"CreatedById,omitempty"`
	CreatedDate        time.Time    `xml:"CreatedDate,omitempty"`
	FeedAttachments    *QueryResult `xml:"FeedAttachments,omitempty"`
	FeedComments       *QueryResult `xml:"FeedComments,omitempty"`
	FeedLikes          *QueryResult `xml:"FeedLikes,omitempty"`
	FeedSignals        *QueryResult `xml:"FeedSignals,omitempty"`
	FeedTrackedChanges *QueryResult `xml:"FeedTrackedChanges,omitempty"`
	InsertedBy         *SObject     `xml:"InsertedBy,omitempty"`
	InsertedById       *ID          `xml:"InsertedById,omitempty"`
	IsDeleted          bool         `xml:"IsDeleted,omitempty"`
	IsRichText         bool         `xml:"IsRichText,omitempty"`
	LastModifiedDate   time.Time    `xml:"LastModifiedDate,omitempty"`
	LikeCount          int32        `xml:"LikeCount,omitempty"`
	LinkUrl            string       `xml:"LinkUrl,omitempty"`
	Parent             *Account     `xml:"Parent,omitempty"`
	ParentId           *ID          `xml:"ParentId,omitempty"`
	RelatedRecordId    *ID          `xml:"RelatedRecordId,omitempty"`
	SystemModstamp     time.Time    `xml:"SystemModstamp,omitempty"`
	Title              string       `xml:"Title,omitempty"`
	Type               string       `xml:"Type,omitempty"`
}

type AccountHistory struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com AccountHistory"`

	*SObject

	Account     *Account    `xml:"Account,omitempty"`
	AccountId   *ID         `xml:"AccountId,omitempty"`
	CreatedBy   *SObject    `xml:"CreatedBy,omitempty"`
	CreatedById *ID         `xml:"CreatedById,omitempty"`
	CreatedDate time.Time   `xml:"CreatedDate,omitempty"`
	Field       string      `xml:"Field,omitempty"`
	IsDeleted   bool        `xml:"IsDeleted,omitempty"`
	NewValue    interface{} `xml:"NewValue,omitempty"`
	OldValue    interface{} `xml:"OldValue,omitempty"`
}

type AccountPartner struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com AccountPartner"`

	*SObject

	AccountFrom      *Account     `xml:"AccountFrom,omitempty"`
	AccountFromId    *ID          `xml:"AccountFromId,omitempty"`
	AccountTo        *Account     `xml:"AccountTo,omitempty"`
	AccountToId      *ID          `xml:"AccountToId,omitempty"`
	CreatedBy        *User        `xml:"CreatedBy,omitempty"`
	CreatedById      *ID          `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time    `xml:"CreatedDate,omitempty"`
	IsDeleted        bool         `xml:"IsDeleted,omitempty"`
	IsPrimary        bool         `xml:"IsPrimary,omitempty"`
	LastModifiedBy   *User        `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID          `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time    `xml:"LastModifiedDate,omitempty"`
	Opportunity      *Opportunity `xml:"Opportunity,omitempty"`
	OpportunityId    *ID          `xml:"OpportunityId,omitempty"`
	ReversePartnerId *ID          `xml:"ReversePartnerId,omitempty"`
	Role             string       `xml:"Role,omitempty"`
	SystemModstamp   time.Time    `xml:"SystemModstamp,omitempty"`
}

type AccountShare struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com AccountShare"`

	*SObject

	Account                *Account  `xml:"Account,omitempty"`
	AccountAccessLevel     string    `xml:"AccountAccessLevel,omitempty"`
	AccountId              *ID       `xml:"AccountId,omitempty"`
	CaseAccessLevel        string    `xml:"CaseAccessLevel,omitempty"`
	ContactAccessLevel     string    `xml:"ContactAccessLevel,omitempty"`
	IsDeleted              bool      `xml:"IsDeleted,omitempty"`
	LastModifiedBy         *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById       *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate       time.Time `xml:"LastModifiedDate,omitempty"`
	OpportunityAccessLevel string    `xml:"OpportunityAccessLevel,omitempty"`
	RowCause               string    `xml:"RowCause,omitempty"`
	UserOrGroup            *SObject  `xml:"UserOrGroup,omitempty"`
	UserOrGroupId          *ID       `xml:"UserOrGroupId,omitempty"`
}

type ActionLinkGroupTemplate struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ActionLinkGroupTemplate"`

	*SObject

	ActionLinkTemplates  *QueryResult `xml:"ActionLinkTemplates,omitempty"`
	Category             string       `xml:"Category,omitempty"`
	CreatedBy            *User        `xml:"CreatedBy,omitempty"`
	CreatedById          *ID          `xml:"CreatedById,omitempty"`
	CreatedDate          time.Time    `xml:"CreatedDate,omitempty"`
	DeveloperName        string       `xml:"DeveloperName,omitempty"`
	ExecutionsAllowed    string       `xml:"ExecutionsAllowed,omitempty"`
	HoursUntilExpiration int32        `xml:"HoursUntilExpiration,omitempty"`
	IsDeleted            bool         `xml:"IsDeleted,omitempty"`
	IsPublished          bool         `xml:"IsPublished,omitempty"`
	Language             string       `xml:"Language,omitempty"`
	LastModifiedBy       *User        `xml:"LastModifiedBy,omitempty"`
	LastModifiedById     *ID          `xml:"LastModifiedById,omitempty"`
	LastModifiedDate     time.Time    `xml:"LastModifiedDate,omitempty"`
	MasterLabel          string       `xml:"MasterLabel,omitempty"`
	NamespacePrefix      string       `xml:"NamespacePrefix,omitempty"`
	SystemModstamp       time.Time    `xml:"SystemModstamp,omitempty"`
}

type ActionLinkTemplate struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ActionLinkTemplate"`

	*SObject

	ActionLinkGroupTemplate   *ActionLinkGroupTemplate `xml:"ActionLinkGroupTemplate,omitempty"`
	ActionLinkGroupTemplateId *ID                      `xml:"ActionLinkGroupTemplateId,omitempty"`
	ActionUrl                 string                   `xml:"ActionUrl,omitempty"`
	CreatedBy                 *User                    `xml:"CreatedBy,omitempty"`
	CreatedById               *ID                      `xml:"CreatedById,omitempty"`
	CreatedDate               time.Time                `xml:"CreatedDate,omitempty"`
	Headers                   string                   `xml:"Headers,omitempty"`
	IsConfirmationRequired    bool                     `xml:"IsConfirmationRequired,omitempty"`
	IsDeleted                 bool                     `xml:"IsDeleted,omitempty"`
	IsGroupDefault            bool                     `xml:"IsGroupDefault,omitempty"`
	Label                     string                   `xml:"Label,omitempty"`
	LabelKey                  string                   `xml:"LabelKey,omitempty"`
	LastModifiedBy            *User                    `xml:"LastModifiedBy,omitempty"`
	LastModifiedById          *ID                      `xml:"LastModifiedById,omitempty"`
	LastModifiedDate          time.Time                `xml:"LastModifiedDate,omitempty"`
	LinkType                  string                   `xml:"LinkType,omitempty"`
	Method                    string                   `xml:"Method,omitempty"`
	Position                  int32                    `xml:"Position,omitempty"`
	RequestBody               string                   `xml:"RequestBody,omitempty"`
	SystemModstamp            time.Time                `xml:"SystemModstamp,omitempty"`
	UserAlias                 string                   `xml:"UserAlias,omitempty"`
	UserVisibility            string                   `xml:"UserVisibility,omitempty"`
}

type ActivityHistory struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ActivityHistory"`

	*SObject

	Account                *Account      `xml:"Account,omitempty"`
	AccountId              *ID           `xml:"AccountId,omitempty"`
	ActivityDate           time.Time     `xml:"ActivityDate,omitempty"`
	ActivitySubtype        string        `xml:"ActivitySubtype,omitempty"`
	ActivityType           string        `xml:"ActivityType,omitempty"`
	AlternateDetail        *EmailMessage `xml:"AlternateDetail,omitempty"`
	AlternateDetailId      *ID           `xml:"AlternateDetailId,omitempty"`
	CallDisposition        string        `xml:"CallDisposition,omitempty"`
	CallDurationInSeconds  int32         `xml:"CallDurationInSeconds,omitempty"`
	CallObject             string        `xml:"CallObject,omitempty"`
	CallType               string        `xml:"CallType,omitempty"`
	CreatedBy              *User         `xml:"CreatedBy,omitempty"`
	CreatedById            *ID           `xml:"CreatedById,omitempty"`
	CreatedDate            time.Time     `xml:"CreatedDate,omitempty"`
	Description            string        `xml:"Description,omitempty"`
	DurationInMinutes      int32         `xml:"DurationInMinutes,omitempty"`
	EndDateTime            time.Time     `xml:"EndDateTime,omitempty"`
	IsAllDayEvent          bool          `xml:"IsAllDayEvent,omitempty"`
	IsClosed               bool          `xml:"IsClosed,omitempty"`
	IsDeleted              bool          `xml:"IsDeleted,omitempty"`
	IsHighPriority         bool          `xml:"IsHighPriority,omitempty"`
	IsReminderSet          bool          `xml:"IsReminderSet,omitempty"`
	IsTask                 bool          `xml:"IsTask,omitempty"`
	IsVisibleInSelfService bool          `xml:"IsVisibleInSelfService,omitempty"`
	LastModifiedBy         *User         `xml:"LastModifiedBy,omitempty"`
	LastModifiedById       *ID           `xml:"LastModifiedById,omitempty"`
	LastModifiedDate       time.Time     `xml:"LastModifiedDate,omitempty"`
	Location               string        `xml:"Location,omitempty"`
	Owner                  *SObject      `xml:"Owner,omitempty"`
	OwnerId                *ID           `xml:"OwnerId,omitempty"`
	Priority               string        `xml:"Priority,omitempty"`
	ReminderDateTime       time.Time     `xml:"ReminderDateTime,omitempty"`
	StartDateTime          time.Time     `xml:"StartDateTime,omitempty"`
	Status                 string        `xml:"Status,omitempty"`
	Subject                string        `xml:"Subject,omitempty"`
	SystemModstamp         time.Time     `xml:"SystemModstamp,omitempty"`
	What                   *SObject      `xml:"What,omitempty"`
	WhatId                 *ID           `xml:"WhatId,omitempty"`
	Who                    *SObject      `xml:"Who,omitempty"`
	WhoId                  *ID           `xml:"WhoId,omitempty"`
}

type AdditionalNumber struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com AdditionalNumber"`

	*SObject

	CallCenterId     *ID       `xml:"CallCenterId,omitempty"`
	CreatedBy        *User     `xml:"CreatedBy,omitempty"`
	CreatedById      *ID       `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time `xml:"CreatedDate,omitempty"`
	Description      string    `xml:"Description,omitempty"`
	IsDeleted        bool      `xml:"IsDeleted,omitempty"`
	LastModifiedBy   *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time `xml:"LastModifiedDate,omitempty"`
	Name             string    `xml:"Name,omitempty"`
	Phone            string    `xml:"Phone,omitempty"`
	SystemModstamp   time.Time `xml:"SystemModstamp,omitempty"`
}

type Announcement struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com Announcement"`

	*SObject

	CreatedBy        *User               `xml:"CreatedBy,omitempty"`
	CreatedById      *ID                 `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time           `xml:"CreatedDate,omitempty"`
	ExpirationDate   time.Time           `xml:"ExpirationDate,omitempty"`
	FeedItem         *FeedItem           `xml:"FeedItem,omitempty"`
	FeedItemId       *ID                 `xml:"FeedItemId,omitempty"`
	IsArchived       bool                `xml:"IsArchived,omitempty"`
	IsDeleted        bool                `xml:"IsDeleted,omitempty"`
	LastModifiedBy   *User               `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID                 `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time           `xml:"LastModifiedDate,omitempty"`
	Parent           *CollaborationGroup `xml:"Parent,omitempty"`
	ParentId         *ID                 `xml:"ParentId,omitempty"`
	SendEmails       bool                `xml:"SendEmails,omitempty"`
	SystemModstamp   time.Time           `xml:"SystemModstamp,omitempty"`
	UserRecordAccess *UserRecordAccess   `xml:"UserRecordAccess,omitempty"`
}

type ApexClass struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ApexClass"`

	*SObject

	ApiVersion             float64      `xml:"ApiVersion,omitempty"`
	Body                   string       `xml:"Body,omitempty"`
	BodyCrc                float64      `xml:"BodyCrc,omitempty"`
	CreatedBy              *User        `xml:"CreatedBy,omitempty"`
	CreatedById            *ID          `xml:"CreatedById,omitempty"`
	CreatedDate            time.Time    `xml:"CreatedDate,omitempty"`
	IsValid                bool         `xml:"IsValid,omitempty"`
	LastModifiedBy         *User        `xml:"LastModifiedBy,omitempty"`
	LastModifiedById       *ID          `xml:"LastModifiedById,omitempty"`
	LastModifiedDate       time.Time    `xml:"LastModifiedDate,omitempty"`
	LengthWithoutComments  int32        `xml:"LengthWithoutComments,omitempty"`
	Name                   string       `xml:"Name,omitempty"`
	NamespacePrefix        string       `xml:"NamespacePrefix,omitempty"`
	SetupEntityAccessItems *QueryResult `xml:"SetupEntityAccessItems,omitempty"`
	Status                 string       `xml:"Status,omitempty"`
	SystemModstamp         time.Time    `xml:"SystemModstamp,omitempty"`
}

type ApexComponent struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ApexComponent"`

	*SObject

	ApiVersion       float64   `xml:"ApiVersion,omitempty"`
	ControllerKey    string    `xml:"ControllerKey,omitempty"`
	ControllerType   string    `xml:"ControllerType,omitempty"`
	CreatedBy        *User     `xml:"CreatedBy,omitempty"`
	CreatedById      *ID       `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time `xml:"CreatedDate,omitempty"`
	Description      string    `xml:"Description,omitempty"`
	LastModifiedBy   *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time `xml:"LastModifiedDate,omitempty"`
	Markup           string    `xml:"Markup,omitempty"`
	MasterLabel      string    `xml:"MasterLabel,omitempty"`
	Name             string    `xml:"Name,omitempty"`
	NamespacePrefix  string    `xml:"NamespacePrefix,omitempty"`
	SystemModstamp   time.Time `xml:"SystemModstamp,omitempty"`
}

type ApexEmailNotification struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ApexEmailNotification"`

	*SObject

	CreatedBy        *User     `xml:"CreatedBy,omitempty"`
	CreatedById      *ID       `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time `xml:"CreatedDate,omitempty"`
	Email            string    `xml:"Email,omitempty"`
	IsDeleted        bool      `xml:"IsDeleted,omitempty"`
	LastModifiedBy   *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time `xml:"LastModifiedDate,omitempty"`
	SystemModstamp   time.Time `xml:"SystemModstamp,omitempty"`
	User             *User     `xml:"User,omitempty"`
	UserId           *ID       `xml:"UserId,omitempty"`
}

type ApexLog struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ApexLog"`

	*SObject

	Application          string    `xml:"Application,omitempty"`
	DurationMilliseconds int32     `xml:"DurationMilliseconds,omitempty"`
	LastModifiedDate     time.Time `xml:"LastModifiedDate,omitempty"`
	Location             string    `xml:"Location,omitempty"`
	LogLength            int32     `xml:"LogLength,omitempty"`
	LogUser              *SObject  `xml:"LogUser,omitempty"`
	LogUserId            *ID       `xml:"LogUserId,omitempty"`
	Operation            string    `xml:"Operation,omitempty"`
	Request              string    `xml:"Request,omitempty"`
	StartTime            time.Time `xml:"StartTime,omitempty"`
	Status               string    `xml:"Status,omitempty"`
	SystemModstamp       time.Time `xml:"SystemModstamp,omitempty"`
}

type ApexPage struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ApexPage"`

	*SObject

	ApiVersion                  float64      `xml:"ApiVersion,omitempty"`
	ControllerKey               string       `xml:"ControllerKey,omitempty"`
	ControllerType              string       `xml:"ControllerType,omitempty"`
	CreatedBy                   *User        `xml:"CreatedBy,omitempty"`
	CreatedById                 *ID          `xml:"CreatedById,omitempty"`
	CreatedDate                 time.Time    `xml:"CreatedDate,omitempty"`
	Description                 string       `xml:"Description,omitempty"`
	IsAvailableInTouch          bool         `xml:"IsAvailableInTouch,omitempty"`
	IsConfirmationTokenRequired bool         `xml:"IsConfirmationTokenRequired,omitempty"`
	LastModifiedBy              *User        `xml:"LastModifiedBy,omitempty"`
	LastModifiedById            *ID          `xml:"LastModifiedById,omitempty"`
	LastModifiedDate            time.Time    `xml:"LastModifiedDate,omitempty"`
	Markup                      string       `xml:"Markup,omitempty"`
	MasterLabel                 string       `xml:"MasterLabel,omitempty"`
	Name                        string       `xml:"Name,omitempty"`
	NamespacePrefix             string       `xml:"NamespacePrefix,omitempty"`
	SetupEntityAccessItems      *QueryResult `xml:"SetupEntityAccessItems,omitempty"`
	SystemModstamp              time.Time    `xml:"SystemModstamp,omitempty"`
}

type ApexPageInfo struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ApexPageInfo"`

	*SObject

	ApexPageId         string  `xml:"ApexPageId,omitempty"`
	ApiVersion         float64 `xml:"ApiVersion,omitempty"`
	Description        string  `xml:"Description,omitempty"`
	DurableId          string  `xml:"DurableId,omitempty"`
	IsAvailableInTouch bool    `xml:"IsAvailableInTouch,omitempty"`
	IsShowHeader       string  `xml:"IsShowHeader,omitempty"`
	MasterLabel        string  `xml:"MasterLabel,omitempty"`
	Name               string  `xml:"Name,omitempty"`
	NameSpacePrefix    string  `xml:"NameSpacePrefix,omitempty"`
}

type ApexTestQueueItem struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ApexTestQueueItem"`

	*SObject

	ApexClass              *ApexClass `xml:"ApexClass,omitempty"`
	ApexClassId            *ID        `xml:"ApexClassId,omitempty"`
	CreatedBy              *User      `xml:"CreatedBy,omitempty"`
	CreatedById            *ID        `xml:"CreatedById,omitempty"`
	CreatedDate            time.Time  `xml:"CreatedDate,omitempty"`
	ExtendedStatus         string     `xml:"ExtendedStatus,omitempty"`
	ParentJobId            *ID        `xml:"ParentJobId,omitempty"`
	ShouldSkipCodeCoverage bool       `xml:"ShouldSkipCodeCoverage,omitempty"`
	Status                 string     `xml:"Status,omitempty"`
	SystemModstamp         time.Time  `xml:"SystemModstamp,omitempty"`
	TestRunResultId        *ID        `xml:"TestRunResultId,omitempty"`
}

type ApexTestResult struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ApexTestResult"`

	*SObject

	ApexClass           *ApexClass         `xml:"ApexClass,omitempty"`
	ApexClassId         *ID                `xml:"ApexClassId,omitempty"`
	ApexLog             *ApexLog           `xml:"ApexLog,omitempty"`
	ApexLogId           *ID                `xml:"ApexLogId,omitempty"`
	ApexTestResults     *QueryResult       `xml:"ApexTestResults,omitempty"`
	ApexTestRunResult   *ApexTestRunResult `xml:"ApexTestRunResult,omitempty"`
	ApexTestRunResultId *ID                `xml:"ApexTestRunResultId,omitempty"`
	AsyncApexJob        *AsyncApexJob      `xml:"AsyncApexJob,omitempty"`
	AsyncApexJobId      *ID                `xml:"AsyncApexJobId,omitempty"`
	Message             string             `xml:"Message,omitempty"`
	MethodName          string             `xml:"MethodName,omitempty"`
	Outcome             string             `xml:"Outcome,omitempty"`
	QueueItem           *ApexTestQueueItem `xml:"QueueItem,omitempty"`
	QueueItemId         *ID                `xml:"QueueItemId,omitempty"`
	RunTime             int32              `xml:"RunTime,omitempty"`
	StackTrace          string             `xml:"StackTrace,omitempty"`
	SystemModstamp      time.Time          `xml:"SystemModstamp,omitempty"`
	TestTimestamp       time.Time          `xml:"TestTimestamp,omitempty"`
}

type ApexTestResultLimits struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ApexTestResultLimits"`

	*SObject

	ApexTestResult   *ApexTestResult `xml:"ApexTestResult,omitempty"`
	ApexTestResultId *ID             `xml:"ApexTestResultId,omitempty"`
	AsyncCalls       int32           `xml:"AsyncCalls,omitempty"`
	Callouts         int32           `xml:"Callouts,omitempty"`
	Cpu              int32           `xml:"Cpu,omitempty"`
	CreatedBy        *User           `xml:"CreatedBy,omitempty"`
	CreatedById      *ID             `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time       `xml:"CreatedDate,omitempty"`
	Dml              int32           `xml:"Dml,omitempty"`
	DmlRows          int32           `xml:"DmlRows,omitempty"`
	Email            int32           `xml:"Email,omitempty"`
	IsDeleted        bool            `xml:"IsDeleted,omitempty"`
	LastModifiedBy   *User           `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID             `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time       `xml:"LastModifiedDate,omitempty"`
	LimitContext     string          `xml:"LimitContext,omitempty"`
	LimitExceptions  string          `xml:"LimitExceptions,omitempty"`
	MobilePush       int32           `xml:"MobilePush,omitempty"`
	QueryRows        int32           `xml:"QueryRows,omitempty"`
	Soql             int32           `xml:"Soql,omitempty"`
	Sosl             int32           `xml:"Sosl,omitempty"`
	SystemModstamp   time.Time       `xml:"SystemModstamp,omitempty"`
}

type ApexTestRunResult struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ApexTestRunResult"`

	*SObject

	AsyncApexJob     *AsyncApexJob `xml:"AsyncApexJob,omitempty"`
	AsyncApexJobId   *ID           `xml:"AsyncApexJobId,omitempty"`
	ClassesCompleted int32         `xml:"ClassesCompleted,omitempty"`
	ClassesEnqueued  int32         `xml:"ClassesEnqueued,omitempty"`
	CreatedBy        *User         `xml:"CreatedBy,omitempty"`
	CreatedById      *ID           `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time     `xml:"CreatedDate,omitempty"`
	EndTime          time.Time     `xml:"EndTime,omitempty"`
	IsAllTests       bool          `xml:"IsAllTests,omitempty"`
	IsDeleted        bool          `xml:"IsDeleted,omitempty"`
	JobName          string        `xml:"JobName,omitempty"`
	LastModifiedBy   *User         `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID           `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time     `xml:"LastModifiedDate,omitempty"`
	MethodsCompleted int32         `xml:"MethodsCompleted,omitempty"`
	MethodsEnqueued  int32         `xml:"MethodsEnqueued,omitempty"`
	MethodsFailed    int32         `xml:"MethodsFailed,omitempty"`
	Source           string        `xml:"Source,omitempty"`
	StartTime        time.Time     `xml:"StartTime,omitempty"`
	Status           string        `xml:"Status,omitempty"`
	SystemModstamp   time.Time     `xml:"SystemModstamp,omitempty"`
	TestTime         int32         `xml:"TestTime,omitempty"`
	User             *User         `xml:"User,omitempty"`
	UserId           *ID           `xml:"UserId,omitempty"`
}

type ApexTestSuite struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ApexTestSuite"`

	*SObject

	ApexClassIds       []*ID        `xml:"ApexClassIds,omitempty"`
	ApexClassJunctions *QueryResult `xml:"ApexClassJunctions,omitempty"`
	CreatedBy          *User        `xml:"CreatedBy,omitempty"`
	CreatedById        *ID          `xml:"CreatedById,omitempty"`
	CreatedDate        time.Time    `xml:"CreatedDate,omitempty"`
	IsDeleted          bool         `xml:"IsDeleted,omitempty"`
	LastModifiedBy     *User        `xml:"LastModifiedBy,omitempty"`
	LastModifiedById   *ID          `xml:"LastModifiedById,omitempty"`
	LastModifiedDate   time.Time    `xml:"LastModifiedDate,omitempty"`
	SystemModstamp     time.Time    `xml:"SystemModstamp,omitempty"`
	TestSuiteName      string       `xml:"TestSuiteName,omitempty"`
}

type ApexTrigger struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ApexTrigger"`

	*SObject

	ApiVersion            float64   `xml:"ApiVersion,omitempty"`
	Body                  string    `xml:"Body,omitempty"`
	BodyCrc               float64   `xml:"BodyCrc,omitempty"`
	CreatedBy             *User     `xml:"CreatedBy,omitempty"`
	CreatedById           *ID       `xml:"CreatedById,omitempty"`
	CreatedDate           time.Time `xml:"CreatedDate,omitempty"`
	IsValid               bool      `xml:"IsValid,omitempty"`
	LastModifiedBy        *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById      *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate      time.Time `xml:"LastModifiedDate,omitempty"`
	LengthWithoutComments int32     `xml:"LengthWithoutComments,omitempty"`
	Name                  string    `xml:"Name,omitempty"`
	NamespacePrefix       string    `xml:"NamespacePrefix,omitempty"`
	Status                string    `xml:"Status,omitempty"`
	SystemModstamp        time.Time `xml:"SystemModstamp,omitempty"`
	TableEnumOrId         string    `xml:"TableEnumOrId,omitempty"`
	UsageAfterDelete      bool      `xml:"UsageAfterDelete,omitempty"`
	UsageAfterInsert      bool      `xml:"UsageAfterInsert,omitempty"`
	UsageAfterUndelete    bool      `xml:"UsageAfterUndelete,omitempty"`
	UsageAfterUpdate      bool      `xml:"UsageAfterUpdate,omitempty"`
	UsageBeforeDelete     bool      `xml:"UsageBeforeDelete,omitempty"`
	UsageBeforeInsert     bool      `xml:"UsageBeforeInsert,omitempty"`
	UsageBeforeUpdate     bool      `xml:"UsageBeforeUpdate,omitempty"`
	UsageIsBulk           bool      `xml:"UsageIsBulk,omitempty"`
}

type AppDefinition struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com AppDefinition"`

	*SObject

	AppTabs                      *QueryResult `xml:"AppTabs,omitempty"`
	Description                  string       `xml:"Description,omitempty"`
	DeveloperName                string       `xml:"DeveloperName,omitempty"`
	DurableId                    string       `xml:"DurableId,omitempty"`
	HeaderColor                  string       `xml:"HeaderColor,omitempty"`
	IsLargeFormFactorSupported   bool         `xml:"IsLargeFormFactorSupported,omitempty"`
	IsMediumFormFactorSupported  bool         `xml:"IsMediumFormFactorSupported,omitempty"`
	IsNavAutoTempTabsDisabled    bool         `xml:"IsNavAutoTempTabsDisabled,omitempty"`
	IsNavPersonalizationDisabled bool         `xml:"IsNavPersonalizationDisabled,omitempty"`
	IsOverrideOrgTheme           bool         `xml:"IsOverrideOrgTheme,omitempty"`
	IsSmallFormFactorSupported   bool         `xml:"IsSmallFormFactorSupported,omitempty"`
	Label                        string       `xml:"Label,omitempty"`
	LogoUrl                      string       `xml:"LogoUrl,omitempty"`
	MasterLabel                  string       `xml:"MasterLabel,omitempty"`
	NamespacePrefix              string       `xml:"NamespacePrefix,omitempty"`
	NavType                      string       `xml:"NavType,omitempty"`
	UiType                       string       `xml:"UiType,omitempty"`
	UtilityBar                   string       `xml:"UtilityBar,omitempty"`
}

type AppMenuItem struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com AppMenuItem"`

	*SObject

	ApplicationId             *ID       `xml:"ApplicationId,omitempty"`
	CanvasAccessMethod        string    `xml:"CanvasAccessMethod,omitempty"`
	CanvasEnabled             bool      `xml:"CanvasEnabled,omitempty"`
	CanvasOptions             string    `xml:"CanvasOptions,omitempty"`
	CanvasReferenceId         string    `xml:"CanvasReferenceId,omitempty"`
	CanvasSelectedLocations   string    `xml:"CanvasSelectedLocations,omitempty"`
	CanvasUrl                 string    `xml:"CanvasUrl,omitempty"`
	CreatedBy                 *User     `xml:"CreatedBy,omitempty"`
	CreatedById               *ID       `xml:"CreatedById,omitempty"`
	CreatedDate               time.Time `xml:"CreatedDate,omitempty"`
	Description               string    `xml:"Description,omitempty"`
	IconUrl                   string    `xml:"IconUrl,omitempty"`
	InfoUrl                   string    `xml:"InfoUrl,omitempty"`
	IsAccessible              bool      `xml:"IsAccessible,omitempty"`
	IsDeleted                 bool      `xml:"IsDeleted,omitempty"`
	IsRegisteredDeviceOnly    bool      `xml:"IsRegisteredDeviceOnly,omitempty"`
	IsUsingAdminAuthorization bool      `xml:"IsUsingAdminAuthorization,omitempty"`
	IsVisible                 bool      `xml:"IsVisible,omitempty"`
	Label                     string    `xml:"Label,omitempty"`
	LastModifiedBy            *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById          *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate          time.Time `xml:"LastModifiedDate,omitempty"`
	LogoUrl                   string    `xml:"LogoUrl,omitempty"`
	MobileAppBinaryId         string    `xml:"MobileAppBinaryId,omitempty"`
	MobileAppInstallUrl       string    `xml:"MobileAppInstallUrl,omitempty"`
	MobileAppInstalledDate    time.Time `xml:"MobileAppInstalledDate,omitempty"`
	MobileAppInstalledVersion string    `xml:"MobileAppInstalledVersion,omitempty"`
	MobileAppVer              string    `xml:"MobileAppVer,omitempty"`
	MobileDeviceType          string    `xml:"MobileDeviceType,omitempty"`
	MobileMinOsVer            string    `xml:"MobileMinOsVer,omitempty"`
	MobilePlatform            string    `xml:"MobilePlatform,omitempty"`
	MobileStartUrl            string    `xml:"MobileStartUrl,omitempty"`
	Name                      string    `xml:"Name,omitempty"`
	NamespacePrefix           string    `xml:"NamespacePrefix,omitempty"`
	SortOrder                 int32     `xml:"SortOrder,omitempty"`
	StartUrl                  string    `xml:"StartUrl,omitempty"`
	SystemModstamp            time.Time `xml:"SystemModstamp,omitempty"`
	Type                      string    `xml:"Type,omitempty"`
	UserSortOrder             int32     `xml:"UserSortOrder,omitempty"`
}

type AppTabMember struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com AppTabMember"`

	*SObject

	AppDefinition        *AppDefinition `xml:"AppDefinition,omitempty"`
	AppDefinitionId      string         `xml:"AppDefinitionId,omitempty"`
	DurableId            string         `xml:"DurableId,omitempty"`
	SortOrder            int32          `xml:"SortOrder,omitempty"`
	TabDefinition        *TabDefinition `xml:"TabDefinition,omitempty"`
	TabDefinitionId      string         `xml:"TabDefinitionId,omitempty"`
	WorkspaceDriverField string         `xml:"WorkspaceDriverField,omitempty"`
}

type Asset struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com Asset"`

	*SObject

	Account                    *Account          `xml:"Account,omitempty"`
	AccountId                  *ID               `xml:"AccountId,omitempty"`
	ActivityHistories          *QueryResult      `xml:"ActivityHistories,omitempty"`
	AssetLevel                 int32             `xml:"AssetLevel,omitempty"`
	AssetProvidedBy            *Account          `xml:"AssetProvidedBy,omitempty"`
	AssetProvidedById          *ID               `xml:"AssetProvidedById,omitempty"`
	AssetServicedBy            *Account          `xml:"AssetServicedBy,omitempty"`
	AssetServicedById          *ID               `xml:"AssetServicedById,omitempty"`
	AttachedContentDocuments   *QueryResult      `xml:"AttachedContentDocuments,omitempty"`
	Attachments                *QueryResult      `xml:"Attachments,omitempty"`
	Cases                      *QueryResult      `xml:"Cases,omitempty"`
	ChildAssets                *QueryResult      `xml:"ChildAssets,omitempty"`
	CombinedAttachments        *QueryResult      `xml:"CombinedAttachments,omitempty"`
	Contact                    *Contact          `xml:"Contact,omitempty"`
	ContactId                  *ID               `xml:"ContactId,omitempty"`
	ContentDocumentLinks       *QueryResult      `xml:"ContentDocumentLinks,omitempty"`
	CreatedBy                  *User             `xml:"CreatedBy,omitempty"`
	CreatedById                *ID               `xml:"CreatedById,omitempty"`
	CreatedDate                time.Time         `xml:"CreatedDate,omitempty"`
	Description                string            `xml:"Description,omitempty"`
	Emails                     *QueryResult      `xml:"Emails,omitempty"`
	Events                     *QueryResult      `xml:"Events,omitempty"`
	FeedSubscriptionsForEntity *QueryResult      `xml:"FeedSubscriptionsForEntity,omitempty"`
	Feeds                      *QueryResult      `xml:"Feeds,omitempty"`
	Histories                  *QueryResult      `xml:"Histories,omitempty"`
	InstallDate                time.Time         `xml:"InstallDate,omitempty"`
	IsCompetitorProduct        bool              `xml:"IsCompetitorProduct,omitempty"`
	IsDeleted                  bool              `xml:"IsDeleted,omitempty"`
	IsInternal                 bool              `xml:"IsInternal,omitempty"`
	LastModifiedBy             *User             `xml:"LastModifiedBy,omitempty"`
	LastModifiedById           *ID               `xml:"LastModifiedById,omitempty"`
	LastModifiedDate           time.Time         `xml:"LastModifiedDate,omitempty"`
	LastReferencedDate         time.Time         `xml:"LastReferencedDate,omitempty"`
	LastViewedDate             time.Time         `xml:"LastViewedDate,omitempty"`
	LookedUpFromActivities     *QueryResult      `xml:"LookedUpFromActivities,omitempty"`
	Name                       string            `xml:"Name,omitempty"`
	Notes                      *QueryResult      `xml:"Notes,omitempty"`
	NotesAndAttachments        *QueryResult      `xml:"NotesAndAttachments,omitempty"`
	OpenActivities             *QueryResult      `xml:"OpenActivities,omitempty"`
	Owner                      *User             `xml:"Owner,omitempty"`
	OwnerId                    *ID               `xml:"OwnerId,omitempty"`
	Parent                     *Asset            `xml:"Parent,omitempty"`
	ParentId                   *ID               `xml:"ParentId,omitempty"`
	Price                      float64           `xml:"Price,omitempty"`
	PrimaryAssets              *QueryResult      `xml:"PrimaryAssets,omitempty"`
	ProcessInstances           *QueryResult      `xml:"ProcessInstances,omitempty"`
	ProcessSteps               *QueryResult      `xml:"ProcessSteps,omitempty"`
	Product2                   *Product2         `xml:"Product2,omitempty"`
	Product2Id                 *ID               `xml:"Product2Id,omitempty"`
	ProductCode                string            `xml:"ProductCode,omitempty"`
	PurchaseDate               time.Time         `xml:"PurchaseDate,omitempty"`
	Quantity                   float64           `xml:"Quantity,omitempty"`
	RecordActionHistories      *QueryResult      `xml:"RecordActionHistories,omitempty"`
	RecordActions              *QueryResult      `xml:"RecordActions,omitempty"`
	RelatedAssets              *QueryResult      `xml:"RelatedAssets,omitempty"`
	RootAsset                  *Asset            `xml:"RootAsset,omitempty"`
	RootAssetId                *ID               `xml:"RootAssetId,omitempty"`
	SerialNumber               string            `xml:"SerialNumber,omitempty"`
	Shares                     *QueryResult      `xml:"Shares,omitempty"`
	Status                     string            `xml:"Status,omitempty"`
	StockKeepingUnit           string            `xml:"StockKeepingUnit,omitempty"`
	SystemModstamp             time.Time         `xml:"SystemModstamp,omitempty"`
	Tasks                      *QueryResult      `xml:"Tasks,omitempty"`
	TopicAssignments           *QueryResult      `xml:"TopicAssignments,omitempty"`
	UsageEndDate               time.Time         `xml:"UsageEndDate,omitempty"`
	UserRecordAccess           *UserRecordAccess `xml:"UserRecordAccess,omitempty"`
}

type AssetFeed struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com AssetFeed"`

	*SObject

	BestComment        *FeedComment `xml:"BestComment,omitempty"`
	BestCommentId      *ID          `xml:"BestCommentId,omitempty"`
	Body               string       `xml:"Body,omitempty"`
	CommentCount       int32        `xml:"CommentCount,omitempty"`
	CreatedBy          *SObject     `xml:"CreatedBy,omitempty"`
	CreatedById        *ID          `xml:"CreatedById,omitempty"`
	CreatedDate        time.Time    `xml:"CreatedDate,omitempty"`
	FeedAttachments    *QueryResult `xml:"FeedAttachments,omitempty"`
	FeedComments       *QueryResult `xml:"FeedComments,omitempty"`
	FeedLikes          *QueryResult `xml:"FeedLikes,omitempty"`
	FeedSignals        *QueryResult `xml:"FeedSignals,omitempty"`
	FeedTrackedChanges *QueryResult `xml:"FeedTrackedChanges,omitempty"`
	InsertedBy         *SObject     `xml:"InsertedBy,omitempty"`
	InsertedById       *ID          `xml:"InsertedById,omitempty"`
	IsDeleted          bool         `xml:"IsDeleted,omitempty"`
	IsRichText         bool         `xml:"IsRichText,omitempty"`
	LastModifiedDate   time.Time    `xml:"LastModifiedDate,omitempty"`
	LikeCount          int32        `xml:"LikeCount,omitempty"`
	LinkUrl            string       `xml:"LinkUrl,omitempty"`
	Parent             *Asset       `xml:"Parent,omitempty"`
	ParentId           *ID          `xml:"ParentId,omitempty"`
	RelatedRecordId    *ID          `xml:"RelatedRecordId,omitempty"`
	SystemModstamp     time.Time    `xml:"SystemModstamp,omitempty"`
	Title              string       `xml:"Title,omitempty"`
	Type               string       `xml:"Type,omitempty"`
}

type AssetHistory struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com AssetHistory"`

	*SObject

	Asset       *Asset      `xml:"Asset,omitempty"`
	AssetId     *ID         `xml:"AssetId,omitempty"`
	CreatedBy   *SObject    `xml:"CreatedBy,omitempty"`
	CreatedById *ID         `xml:"CreatedById,omitempty"`
	CreatedDate time.Time   `xml:"CreatedDate,omitempty"`
	Field       string      `xml:"Field,omitempty"`
	IsDeleted   bool        `xml:"IsDeleted,omitempty"`
	NewValue    interface{} `xml:"NewValue,omitempty"`
	OldValue    interface{} `xml:"OldValue,omitempty"`
}

type AssetRelationship struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com AssetRelationship"`

	*SObject

	ActivityHistories          *QueryResult      `xml:"ActivityHistories,omitempty"`
	Asset                      *Asset            `xml:"Asset,omitempty"`
	AssetId                    *ID               `xml:"AssetId,omitempty"`
	AssetRelationshipNumber    string            `xml:"AssetRelationshipNumber,omitempty"`
	AttachedContentDocuments   *QueryResult      `xml:"AttachedContentDocuments,omitempty"`
	CombinedAttachments        *QueryResult      `xml:"CombinedAttachments,omitempty"`
	ContentDocumentLinks       *QueryResult      `xml:"ContentDocumentLinks,omitempty"`
	CreatedBy                  *User             `xml:"CreatedBy,omitempty"`
	CreatedById                *ID               `xml:"CreatedById,omitempty"`
	CreatedDate                time.Time         `xml:"CreatedDate,omitempty"`
	Emails                     *QueryResult      `xml:"Emails,omitempty"`
	Events                     *QueryResult      `xml:"Events,omitempty"`
	FeedSubscriptionsForEntity *QueryResult      `xml:"FeedSubscriptionsForEntity,omitempty"`
	Feeds                      *QueryResult      `xml:"Feeds,omitempty"`
	FromDate                   time.Time         `xml:"FromDate,omitempty"`
	Histories                  *QueryResult      `xml:"Histories,omitempty"`
	IsDeleted                  bool              `xml:"IsDeleted,omitempty"`
	LastModifiedBy             *User             `xml:"LastModifiedBy,omitempty"`
	LastModifiedById           *ID               `xml:"LastModifiedById,omitempty"`
	LastModifiedDate           time.Time         `xml:"LastModifiedDate,omitempty"`
	LastReferencedDate         time.Time         `xml:"LastReferencedDate,omitempty"`
	LastViewedDate             time.Time         `xml:"LastViewedDate,omitempty"`
	LookedUpFromActivities     *QueryResult      `xml:"LookedUpFromActivities,omitempty"`
	OpenActivities             *QueryResult      `xml:"OpenActivities,omitempty"`
	ProcessInstances           *QueryResult      `xml:"ProcessInstances,omitempty"`
	ProcessSteps               *QueryResult      `xml:"ProcessSteps,omitempty"`
	RelatedAsset               *Asset            `xml:"RelatedAsset,omitempty"`
	RelatedAssetId             *ID               `xml:"RelatedAssetId,omitempty"`
	RelationshipType           string            `xml:"RelationshipType,omitempty"`
	SystemModstamp             time.Time         `xml:"SystemModstamp,omitempty"`
	Tasks                      *QueryResult      `xml:"Tasks,omitempty"`
	ToDate                     time.Time         `xml:"ToDate,omitempty"`
	UserRecordAccess           *UserRecordAccess `xml:"UserRecordAccess,omitempty"`
}

type AssetRelationshipFeed struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com AssetRelationshipFeed"`

	*SObject

	BestComment        *FeedComment       `xml:"BestComment,omitempty"`
	BestCommentId      *ID                `xml:"BestCommentId,omitempty"`
	Body               string             `xml:"Body,omitempty"`
	CommentCount       int32              `xml:"CommentCount,omitempty"`
	CreatedBy          *SObject           `xml:"CreatedBy,omitempty"`
	CreatedById        *ID                `xml:"CreatedById,omitempty"`
	CreatedDate        time.Time          `xml:"CreatedDate,omitempty"`
	FeedAttachments    *QueryResult       `xml:"FeedAttachments,omitempty"`
	FeedComments       *QueryResult       `xml:"FeedComments,omitempty"`
	FeedLikes          *QueryResult       `xml:"FeedLikes,omitempty"`
	FeedSignals        *QueryResult       `xml:"FeedSignals,omitempty"`
	FeedTrackedChanges *QueryResult       `xml:"FeedTrackedChanges,omitempty"`
	InsertedBy         *SObject           `xml:"InsertedBy,omitempty"`
	InsertedById       *ID                `xml:"InsertedById,omitempty"`
	IsDeleted          bool               `xml:"IsDeleted,omitempty"`
	IsRichText         bool               `xml:"IsRichText,omitempty"`
	LastModifiedDate   time.Time          `xml:"LastModifiedDate,omitempty"`
	LikeCount          int32              `xml:"LikeCount,omitempty"`
	LinkUrl            string             `xml:"LinkUrl,omitempty"`
	Parent             *AssetRelationship `xml:"Parent,omitempty"`
	ParentId           *ID                `xml:"ParentId,omitempty"`
	RelatedRecordId    *ID                `xml:"RelatedRecordId,omitempty"`
	SystemModstamp     time.Time          `xml:"SystemModstamp,omitempty"`
	Title              string             `xml:"Title,omitempty"`
	Type               string             `xml:"Type,omitempty"`
}

type AssetRelationshipHistory struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com AssetRelationshipHistory"`

	*SObject

	AssetRelationship   *AssetRelationship `xml:"AssetRelationship,omitempty"`
	AssetRelationshipId *ID                `xml:"AssetRelationshipId,omitempty"`
	CreatedBy           *SObject           `xml:"CreatedBy,omitempty"`
	CreatedById         *ID                `xml:"CreatedById,omitempty"`
	CreatedDate         time.Time          `xml:"CreatedDate,omitempty"`
	Field               string             `xml:"Field,omitempty"`
	IsDeleted           bool               `xml:"IsDeleted,omitempty"`
	NewValue            interface{}        `xml:"NewValue,omitempty"`
	OldValue            interface{}        `xml:"OldValue,omitempty"`
}

type AssetShare struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com AssetShare"`

	*SObject

	Asset            *Asset    `xml:"Asset,omitempty"`
	AssetAccessLevel string    `xml:"AssetAccessLevel,omitempty"`
	AssetId          *ID       `xml:"AssetId,omitempty"`
	IsDeleted        bool      `xml:"IsDeleted,omitempty"`
	LastModifiedBy   *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time `xml:"LastModifiedDate,omitempty"`
	RowCause         string    `xml:"RowCause,omitempty"`
	UserOrGroup      *SObject  `xml:"UserOrGroup,omitempty"`
	UserOrGroupId    *ID       `xml:"UserOrGroupId,omitempty"`
}

type AssetTokenEvent struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com AssetTokenEvent"`

	*SObject

	ActorTokenPayload string                `xml:"ActorTokenPayload,omitempty"`
	Asset             *Asset                `xml:"Asset,omitempty"`
	AssetId           *ID                   `xml:"AssetId,omitempty"`
	AssetName         string                `xml:"AssetName,omitempty"`
	AssetSerialNumber string                `xml:"AssetSerialNumber,omitempty"`
	ConnectedApp      *ConnectedApplication `xml:"ConnectedApp,omitempty"`
	ConnectedAppId    *ID                   `xml:"ConnectedAppId,omitempty"`
	CreatedBy         *User                 `xml:"CreatedBy,omitempty"`
	CreatedById       *ID                   `xml:"CreatedById,omitempty"`
	CreatedDate       time.Time             `xml:"CreatedDate,omitempty"`
	DeviceId          string                `xml:"DeviceId,omitempty"`
	DeviceKey         string                `xml:"DeviceKey,omitempty"`
	Expiration        time.Time             `xml:"Expiration,omitempty"`
	Name              string                `xml:"Name,omitempty"`
	ReplayId          string                `xml:"ReplayId,omitempty"`
	User              *User                 `xml:"User,omitempty"`
	UserId            *ID                   `xml:"UserId,omitempty"`
}

type AssignmentRule struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com AssignmentRule"`

	*SObject

	Active           bool      `xml:"Active,omitempty"`
	CreatedBy        *User     `xml:"CreatedBy,omitempty"`
	CreatedById      *ID       `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time `xml:"CreatedDate,omitempty"`
	LastModifiedBy   *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time `xml:"LastModifiedDate,omitempty"`
	Name             string    `xml:"Name,omitempty"`
	SobjectType      string    `xml:"SobjectType,omitempty"`
	SystemModstamp   time.Time `xml:"SystemModstamp,omitempty"`
}

type AsyncApexJob struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com AsyncApexJob"`

	*SObject

	ApexClass           *ApexClass   `xml:"ApexClass,omitempty"`
	ApexClassId         *ID          `xml:"ApexClassId,omitempty"`
	AsyncApex           *QueryResult `xml:"AsyncApex,omitempty"`
	CompletedDate       time.Time    `xml:"CompletedDate,omitempty"`
	CreatedBy           *User        `xml:"CreatedBy,omitempty"`
	CreatedById         *ID          `xml:"CreatedById,omitempty"`
	CreatedDate         time.Time    `xml:"CreatedDate,omitempty"`
	ExtendedStatus      string       `xml:"ExtendedStatus,omitempty"`
	JobItemsProcessed   int32        `xml:"JobItemsProcessed,omitempty"`
	JobType             string       `xml:"JobType,omitempty"`
	LastProcessed       string       `xml:"LastProcessed,omitempty"`
	LastProcessedOffset int32        `xml:"LastProcessedOffset,omitempty"`
	MethodName          string       `xml:"MethodName,omitempty"`
	NumberOfErrors      int32        `xml:"NumberOfErrors,omitempty"`
	ParentJobId         *ID          `xml:"ParentJobId,omitempty"`
	Status              string       `xml:"Status,omitempty"`
	TotalJobItems       int32        `xml:"TotalJobItems,omitempty"`
}

type AttachedContentDocument struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com AttachedContentDocument"`

	*SObject

	ContentDocument        *ContentDocument `xml:"ContentDocument,omitempty"`
	ContentDocumentId      *ID              `xml:"ContentDocumentId,omitempty"`
	ContentSize            int32            `xml:"ContentSize,omitempty"`
	ContentUrl             string           `xml:"ContentUrl,omitempty"`
	CreatedBy              *User            `xml:"CreatedBy,omitempty"`
	CreatedById            *ID              `xml:"CreatedById,omitempty"`
	CreatedDate            time.Time        `xml:"CreatedDate,omitempty"`
	ExternalDataSourceName string           `xml:"ExternalDataSourceName,omitempty"`
	ExternalDataSourceType string           `xml:"ExternalDataSourceType,omitempty"`
	FileExtension          string           `xml:"FileExtension,omitempty"`
	FileType               string           `xml:"FileType,omitempty"`
	IsDeleted              bool             `xml:"IsDeleted,omitempty"`
	LastModifiedBy         *User            `xml:"LastModifiedBy,omitempty"`
	LastModifiedById       *ID              `xml:"LastModifiedById,omitempty"`
	LastModifiedDate       time.Time        `xml:"LastModifiedDate,omitempty"`
	LinkedEntity           *Contract        `xml:"LinkedEntity,omitempty"`
	LinkedEntityId         *ID              `xml:"LinkedEntityId,omitempty"`
	SharingOption          string           `xml:"SharingOption,omitempty"`
	Title                  string           `xml:"Title,omitempty"`
}

type Attachment struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com Attachment"`

	*SObject

	Body             []byte    `xml:"Body,omitempty"`
	BodyLength       int32     `xml:"BodyLength,omitempty"`
	ContentType      string    `xml:"ContentType,omitempty"`
	CreatedBy        *SObject  `xml:"CreatedBy,omitempty"`
	CreatedById      *ID       `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time `xml:"CreatedDate,omitempty"`
	Description      string    `xml:"Description,omitempty"`
	IsDeleted        bool      `xml:"IsDeleted,omitempty"`
	IsPrivate        bool      `xml:"IsPrivate,omitempty"`
	LastModifiedBy   *SObject  `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time `xml:"LastModifiedDate,omitempty"`
	Name             string    `xml:"Name,omitempty"`
	Owner            *SObject  `xml:"Owner,omitempty"`
	OwnerId          *ID       `xml:"OwnerId,omitempty"`
	Parent           *SObject  `xml:"Parent,omitempty"`
	ParentId         *ID       `xml:"ParentId,omitempty"`
	SystemModstamp   time.Time `xml:"SystemModstamp,omitempty"`
}

type AuraDefinition struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com AuraDefinition"`

	*SObject

	AuraDefinitionBundle   *AuraDefinitionBundle `xml:"AuraDefinitionBundle,omitempty"`
	AuraDefinitionBundleId *ID                   `xml:"AuraDefinitionBundleId,omitempty"`
	CreatedBy              *User                 `xml:"CreatedBy,omitempty"`
	CreatedById            *ID                   `xml:"CreatedById,omitempty"`
	CreatedDate            time.Time             `xml:"CreatedDate,omitempty"`
	DefType                string                `xml:"DefType,omitempty"`
	Format                 string                `xml:"Format,omitempty"`
	IsDeleted              bool                  `xml:"IsDeleted,omitempty"`
	LastModifiedBy         *User                 `xml:"LastModifiedBy,omitempty"`
	LastModifiedById       *ID                   `xml:"LastModifiedById,omitempty"`
	LastModifiedDate       time.Time             `xml:"LastModifiedDate,omitempty"`
	Source                 string                `xml:"Source,omitempty"`
	SystemModstamp         time.Time             `xml:"SystemModstamp,omitempty"`
}

type AuraDefinitionBundle struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com AuraDefinitionBundle"`

	*SObject

	ApiVersion       float64   `xml:"ApiVersion,omitempty"`
	CreatedBy        *User     `xml:"CreatedBy,omitempty"`
	CreatedById      *ID       `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time `xml:"CreatedDate,omitempty"`
	Description      string    `xml:"Description,omitempty"`
	DeveloperName    string    `xml:"DeveloperName,omitempty"`
	IsDeleted        bool      `xml:"IsDeleted,omitempty"`
	Language         string    `xml:"Language,omitempty"`
	LastModifiedBy   *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time `xml:"LastModifiedDate,omitempty"`
	MasterLabel      string    `xml:"MasterLabel,omitempty"`
	NamespacePrefix  string    `xml:"NamespacePrefix,omitempty"`
	SystemModstamp   time.Time `xml:"SystemModstamp,omitempty"`
}

type AuraDefinitionBundleInfo struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com AuraDefinitionBundleInfo"`

	*SObject

	ApiVersion             float64      `xml:"ApiVersion,omitempty"`
	AuraDefinitionBundleId string       `xml:"AuraDefinitionBundleId,omitempty"`
	Bundle                 *QueryResult `xml:"Bundle,omitempty"`
	DeveloperName          string       `xml:"DeveloperName,omitempty"`
	DurableId              string       `xml:"DurableId,omitempty"`
	NamespacePrefix        string       `xml:"NamespacePrefix,omitempty"`
}

type AuraDefinitionInfo struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com AuraDefinitionInfo"`

	*SObject

	AuraDefinitionBundleInfo   *AuraDefinitionBundleInfo `xml:"AuraDefinitionBundleInfo,omitempty"`
	AuraDefinitionBundleInfoId string                    `xml:"AuraDefinitionBundleInfoId,omitempty"`
	AuraDefinitionId           string                    `xml:"AuraDefinitionId,omitempty"`
	DefType                    string                    `xml:"DefType,omitempty"`
	DeveloperName              string                    `xml:"DeveloperName,omitempty"`
	DurableId                  string                    `xml:"DurableId,omitempty"`
	Format                     string                    `xml:"Format,omitempty"`
	LastModifiedDate           time.Time                 `xml:"LastModifiedDate,omitempty"`
	NamespacePrefix            string                    `xml:"NamespacePrefix,omitempty"`
	Source                     string                    `xml:"Source,omitempty"`
}

type AuthConfig struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com AuthConfig"`

	*SObject

	AuthOptionsAuthProvider     bool         `xml:"AuthOptionsAuthProvider,omitempty"`
	AuthOptionsSaml             bool         `xml:"AuthOptionsSaml,omitempty"`
	AuthOptionsUsernamePassword bool         `xml:"AuthOptionsUsernamePassword,omitempty"`
	AuthProvidersForConfig      *QueryResult `xml:"AuthProvidersForConfig,omitempty"`
	CreatedBy                   *User        `xml:"CreatedBy,omitempty"`
	CreatedById                 *ID          `xml:"CreatedById,omitempty"`
	CreatedDate                 time.Time    `xml:"CreatedDate,omitempty"`
	DeveloperName               string       `xml:"DeveloperName,omitempty"`
	IsActive                    bool         `xml:"IsActive,omitempty"`
	IsDeleted                   bool         `xml:"IsDeleted,omitempty"`
	Language                    string       `xml:"Language,omitempty"`
	LastModifiedBy              *User        `xml:"LastModifiedBy,omitempty"`
	LastModifiedById            *ID          `xml:"LastModifiedById,omitempty"`
	LastModifiedDate            time.Time    `xml:"LastModifiedDate,omitempty"`
	MasterLabel                 string       `xml:"MasterLabel,omitempty"`
	NamespacePrefix             string       `xml:"NamespacePrefix,omitempty"`
	SystemModstamp              time.Time    `xml:"SystemModstamp,omitempty"`
	Type                        string       `xml:"Type,omitempty"`
	Url                         string       `xml:"Url,omitempty"`
}

type AuthConfigProviders struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com AuthConfigProviders"`

	*SObject

	AuthConfig       *AuthConfig `xml:"AuthConfig,omitempty"`
	AuthConfigId     *ID         `xml:"AuthConfigId,omitempty"`
	AuthProvider     *SObject    `xml:"AuthProvider,omitempty"`
	AuthProviderId   *ID         `xml:"AuthProviderId,omitempty"`
	CreatedBy        *User       `xml:"CreatedBy,omitempty"`
	CreatedById      *ID         `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time   `xml:"CreatedDate,omitempty"`
	IsDeleted        bool        `xml:"IsDeleted,omitempty"`
	LastModifiedBy   *User       `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID         `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time   `xml:"LastModifiedDate,omitempty"`
	SystemModstamp   time.Time   `xml:"SystemModstamp,omitempty"`
}

type AuthProvider struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com AuthProvider"`

	*SObject

	AuthorizeUrl                         string    `xml:"AuthorizeUrl,omitempty"`
	ConsumerKey                          string    `xml:"ConsumerKey,omitempty"`
	ConsumerSecret                       string    `xml:"ConsumerSecret,omitempty"`
	CreatedDate                          time.Time `xml:"CreatedDate,omitempty"`
	CustomMetadataTypeRecord             string    `xml:"CustomMetadataTypeRecord,omitempty"`
	DefaultScopes                        string    `xml:"DefaultScopes,omitempty"`
	DeveloperName                        string    `xml:"DeveloperName,omitempty"`
	ErrorUrl                             string    `xml:"ErrorUrl,omitempty"`
	ExecutionUserId                      *ID       `xml:"ExecutionUserId,omitempty"`
	FriendlyName                         string    `xml:"FriendlyName,omitempty"`
	IconUrl                              string    `xml:"IconUrl,omitempty"`
	IdTokenIssuer                        string    `xml:"IdTokenIssuer,omitempty"`
	LinkKickoffUrl                       string    `xml:"LinkKickoffUrl,omitempty"`
	LogoutUrl                            string    `xml:"LogoutUrl,omitempty"`
	OauthKickoffUrl                      string    `xml:"OauthKickoffUrl,omitempty"`
	OptionsIncludeOrgIdInId              bool      `xml:"OptionsIncludeOrgIdInId,omitempty"`
	OptionsSendAccessTokenInHeader       bool      `xml:"OptionsSendAccessTokenInHeader,omitempty"`
	OptionsSendClientCredentialsInHeader bool      `xml:"OptionsSendClientCredentialsInHeader,omitempty"`
	PluginId                             *ID       `xml:"PluginId,omitempty"`
	ProviderType                         string    `xml:"ProviderType,omitempty"`
	RegistrationHandlerId                *ID       `xml:"RegistrationHandlerId,omitempty"`
	SsoKickoffUrl                        string    `xml:"SsoKickoffUrl,omitempty"`
	TokenUrl                             string    `xml:"TokenUrl,omitempty"`
	UserInfoUrl                          string    `xml:"UserInfoUrl,omitempty"`
}

type AuthSession struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com AuthSession"`

	*SObject

	CreatedDate               time.Time     `xml:"CreatedDate,omitempty"`
	IsCurrent                 bool          `xml:"IsCurrent,omitempty"`
	LastModifiedDate          time.Time     `xml:"LastModifiedDate,omitempty"`
	LoginGeo                  *LoginGeo     `xml:"LoginGeo,omitempty"`
	LoginGeoId                *ID           `xml:"LoginGeoId,omitempty"`
	LoginHistory              *LoginHistory `xml:"LoginHistory,omitempty"`
	LoginHistoryId            *ID           `xml:"LoginHistoryId,omitempty"`
	LoginType                 string        `xml:"LoginType,omitempty"`
	LogoutUrl                 string        `xml:"LogoutUrl,omitempty"`
	NumSecondsValid           int32         `xml:"NumSecondsValid,omitempty"`
	ParentId                  *ID           `xml:"ParentId,omitempty"`
	SessionPermSetActivations *QueryResult  `xml:"SessionPermSetActivations,omitempty"`
	SessionSecurityLevel      string        `xml:"SessionSecurityLevel,omitempty"`
	SessionType               string        `xml:"SessionType,omitempty"`
	SourceIp                  string        `xml:"SourceIp,omitempty"`
	UserType                  string        `xml:"UserType,omitempty"`
	Users                     *User         `xml:"Users,omitempty"`
	UsersId                   *ID           `xml:"UsersId,omitempty"`
}

type BackgroundOperation struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com BackgroundOperation"`

	*SObject

	CreatedBy        *User                `xml:"CreatedBy,omitempty"`
	CreatedById      *ID                  `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time            `xml:"CreatedDate,omitempty"`
	Error            string               `xml:"Error,omitempty"`
	ExecutionGroup   string               `xml:"ExecutionGroup,omitempty"`
	ExpiresAt        time.Time            `xml:"ExpiresAt,omitempty"`
	FinishedAt       time.Time            `xml:"FinishedAt,omitempty"`
	GroupLeader      *BackgroundOperation `xml:"GroupLeader,omitempty"`
	GroupLeaderId    *ID                  `xml:"GroupLeaderId,omitempty"`
	IsDeleted        bool                 `xml:"IsDeleted,omitempty"`
	LastModifiedBy   *User                `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID                  `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time            `xml:"LastModifiedDate,omitempty"`
	MergedOperations *QueryResult         `xml:"MergedOperations,omitempty"`
	Name             string               `xml:"Name,omitempty"`
	NumFollowers     int32                `xml:"NumFollowers,omitempty"`
	ParentKey        string               `xml:"ParentKey,omitempty"`
	ProcessAfter     time.Time            `xml:"ProcessAfter,omitempty"`
	RetryBackoff     int32                `xml:"RetryBackoff,omitempty"`
	RetryCount       int32                `xml:"RetryCount,omitempty"`
	RetryLimit       int32                `xml:"RetryLimit,omitempty"`
	SequenceGroup    string               `xml:"SequenceGroup,omitempty"`
	SequenceNumber   int32                `xml:"SequenceNumber,omitempty"`
	StartedAt        time.Time            `xml:"StartedAt,omitempty"`
	Status           string               `xml:"Status,omitempty"`
	SubmittedAt      time.Time            `xml:"SubmittedAt,omitempty"`
	SystemModstamp   time.Time            `xml:"SystemModstamp,omitempty"`
	Timeout          int32                `xml:"Timeout,omitempty"`
	UserRecordAccess *UserRecordAccess    `xml:"UserRecordAccess,omitempty"`
	WorkerUri        string               `xml:"WorkerUri,omitempty"`
}

type BatchApexErrorEvent struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com BatchApexErrorEvent"`

	*SObject

	AsyncApexJobId              string    `xml:"AsyncApexJobId,omitempty"`
	CreatedBy                   *User     `xml:"CreatedBy,omitempty"`
	CreatedById                 *ID       `xml:"CreatedById,omitempty"`
	CreatedDate                 time.Time `xml:"CreatedDate,omitempty"`
	DoesExceedJobScopeMaxLength bool      `xml:"DoesExceedJobScopeMaxLength,omitempty"`
	ExceptionType               string    `xml:"ExceptionType,omitempty"`
	JobScope                    string    `xml:"JobScope,omitempty"`
	Message                     string    `xml:"Message,omitempty"`
	ReplayId                    string    `xml:"ReplayId,omitempty"`
	RequestId                   string    `xml:"RequestId,omitempty"`
	StackTrace                  string    `xml:"StackTrace,omitempty"`
}

type BrandTemplate struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com BrandTemplate"`

	*SObject

	CreatedBy        *User     `xml:"CreatedBy,omitempty"`
	CreatedById      *ID       `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time `xml:"CreatedDate,omitempty"`
	Description      string    `xml:"Description,omitempty"`
	DeveloperName    string    `xml:"DeveloperName,omitempty"`
	IsActive         bool      `xml:"IsActive,omitempty"`
	LastModifiedBy   *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time `xml:"LastModifiedDate,omitempty"`
	Name             string    `xml:"Name,omitempty"`
	NamespacePrefix  string    `xml:"NamespacePrefix,omitempty"`
	SystemModstamp   time.Time `xml:"SystemModstamp,omitempty"`
	Value            string    `xml:"Value,omitempty"`
}

type BrandingSet struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com BrandingSet"`

	*SObject

	CreatedBy        *User     `xml:"CreatedBy,omitempty"`
	CreatedById      *ID       `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time `xml:"CreatedDate,omitempty"`
	Description      string    `xml:"Description,omitempty"`
	DeveloperName    string    `xml:"DeveloperName,omitempty"`
	IsDeleted        bool      `xml:"IsDeleted,omitempty"`
	Language         string    `xml:"Language,omitempty"`
	LastModifiedBy   *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time `xml:"LastModifiedDate,omitempty"`
	MasterLabel      string    `xml:"MasterLabel,omitempty"`
	NamespacePrefix  string    `xml:"NamespacePrefix,omitempty"`
	SystemModstamp   time.Time `xml:"SystemModstamp,omitempty"`
}

type BrandingSetProperty struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com BrandingSetProperty"`

	*SObject

	BrandingSet      *BrandingSet `xml:"BrandingSet,omitempty"`
	BrandingSetId    *ID          `xml:"BrandingSetId,omitempty"`
	CreatedBy        *User        `xml:"CreatedBy,omitempty"`
	CreatedById      *ID          `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time    `xml:"CreatedDate,omitempty"`
	IsDeleted        bool         `xml:"IsDeleted,omitempty"`
	LastModifiedBy   *User        `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID          `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time    `xml:"LastModifiedDate,omitempty"`
	PropertyName     string       `xml:"PropertyName,omitempty"`
	PropertyValue    string       `xml:"PropertyValue,omitempty"`
	SystemModstamp   time.Time    `xml:"SystemModstamp,omitempty"`
}

type BusinessHours struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com BusinessHours"`

	*SObject

	Cases              *QueryResult `xml:"Cases,omitempty"`
	CreatedBy          *User        `xml:"CreatedBy,omitempty"`
	CreatedById        *ID          `xml:"CreatedById,omitempty"`
	CreatedDate        time.Time    `xml:"CreatedDate,omitempty"`
	FridayEndTime      time.Time    `xml:"FridayEndTime,omitempty"`
	FridayStartTime    time.Time    `xml:"FridayStartTime,omitempty"`
	IsActive           bool         `xml:"IsActive,omitempty"`
	IsDefault          bool         `xml:"IsDefault,omitempty"`
	LastModifiedBy     *User        `xml:"LastModifiedBy,omitempty"`
	LastModifiedById   *ID          `xml:"LastModifiedById,omitempty"`
	LastModifiedDate   time.Time    `xml:"LastModifiedDate,omitempty"`
	LastViewedDate     time.Time    `xml:"LastViewedDate,omitempty"`
	MondayEndTime      time.Time    `xml:"MondayEndTime,omitempty"`
	MondayStartTime    time.Time    `xml:"MondayStartTime,omitempty"`
	Name               string       `xml:"Name,omitempty"`
	SaturdayEndTime    time.Time    `xml:"SaturdayEndTime,omitempty"`
	SaturdayStartTime  time.Time    `xml:"SaturdayStartTime,omitempty"`
	SundayEndTime      time.Time    `xml:"SundayEndTime,omitempty"`
	SundayStartTime    time.Time    `xml:"SundayStartTime,omitempty"`
	SystemModstamp     time.Time    `xml:"SystemModstamp,omitempty"`
	ThursdayEndTime    time.Time    `xml:"ThursdayEndTime,omitempty"`
	ThursdayStartTime  time.Time    `xml:"ThursdayStartTime,omitempty"`
	TimeZoneSidKey     string       `xml:"TimeZoneSidKey,omitempty"`
	TuesdayEndTime     time.Time    `xml:"TuesdayEndTime,omitempty"`
	TuesdayStartTime   time.Time    `xml:"TuesdayStartTime,omitempty"`
	WednesdayEndTime   time.Time    `xml:"WednesdayEndTime,omitempty"`
	WednesdayStartTime time.Time    `xml:"WednesdayStartTime,omitempty"`
}

type BusinessProcess struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com BusinessProcess"`

	*SObject

	CreatedBy        *User     `xml:"CreatedBy,omitempty"`
	CreatedById      *ID       `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time `xml:"CreatedDate,omitempty"`
	Description      string    `xml:"Description,omitempty"`
	IsActive         bool      `xml:"IsActive,omitempty"`
	LastModifiedBy   *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time `xml:"LastModifiedDate,omitempty"`
	Name             string    `xml:"Name,omitempty"`
	NamespacePrefix  string    `xml:"NamespacePrefix,omitempty"`
	SystemModstamp   time.Time `xml:"SystemModstamp,omitempty"`
	TableEnumOrId    string    `xml:"TableEnumOrId,omitempty"`
}

type CallCenter struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com CallCenter"`

	*SObject

	AdapterUrl       string    `xml:"AdapterUrl,omitempty"`
	CreatedBy        *User     `xml:"CreatedBy,omitempty"`
	CreatedById      *ID       `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time `xml:"CreatedDate,omitempty"`
	CustomSettings   string    `xml:"CustomSettings,omitempty"`
	InternalName     string    `xml:"InternalName,omitempty"`
	LastModifiedBy   *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time `xml:"LastModifiedDate,omitempty"`
	Name             string    `xml:"Name,omitempty"`
	SystemModstamp   time.Time `xml:"SystemModstamp,omitempty"`
	Version          float64   `xml:"Version,omitempty"`
}

type Campaign struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com Campaign"`

	*SObject

	ActivityHistories          *QueryResult      `xml:"ActivityHistories,omitempty"`
	ActualCost                 float64           `xml:"ActualCost,omitempty"`
	AmountAllOpportunities     float64           `xml:"AmountAllOpportunities,omitempty"`
	AmountWonOpportunities     float64           `xml:"AmountWonOpportunities,omitempty"`
	AttachedContentDocuments   *QueryResult      `xml:"AttachedContentDocuments,omitempty"`
	Attachments                *QueryResult      `xml:"Attachments,omitempty"`
	BudgetedCost               float64           `xml:"BudgetedCost,omitempty"`
	CampaignMemberRecordType   *RecordType       `xml:"CampaignMemberRecordType,omitempty"`
	CampaignMemberRecordTypeId *ID               `xml:"CampaignMemberRecordTypeId,omitempty"`
	CampaignMemberStatuses     *QueryResult      `xml:"CampaignMemberStatuses,omitempty"`
	CampaignMembers            *QueryResult      `xml:"CampaignMembers,omitempty"`
	ChildCampaigns             *QueryResult      `xml:"ChildCampaigns,omitempty"`
	CombinedAttachments        *QueryResult      `xml:"CombinedAttachments,omitempty"`
	ContentDocumentLinks       *QueryResult      `xml:"ContentDocumentLinks,omitempty"`
	CreatedBy                  *User             `xml:"CreatedBy,omitempty"`
	CreatedById                *ID               `xml:"CreatedById,omitempty"`
	CreatedDate                time.Time         `xml:"CreatedDate,omitempty"`
	Description                string            `xml:"Description,omitempty"`
	Emails                     *QueryResult      `xml:"Emails,omitempty"`
	EndDate                    time.Time         `xml:"EndDate,omitempty"`
	Events                     *QueryResult      `xml:"Events,omitempty"`
	ExpectedResponse           float64           `xml:"ExpectedResponse,omitempty"`
	ExpectedRevenue            float64           `xml:"ExpectedRevenue,omitempty"`
	FeedSubscriptionsForEntity *QueryResult      `xml:"FeedSubscriptionsForEntity,omitempty"`
	Feeds                      *QueryResult      `xml:"Feeds,omitempty"`
	Histories                  *QueryResult      `xml:"Histories,omitempty"`
	IsActive                   bool              `xml:"IsActive,omitempty"`
	IsDeleted                  bool              `xml:"IsDeleted,omitempty"`
	LastActivityDate           time.Time         `xml:"LastActivityDate,omitempty"`
	LastModifiedBy             *User             `xml:"LastModifiedBy,omitempty"`
	LastModifiedById           *ID               `xml:"LastModifiedById,omitempty"`
	LastModifiedDate           time.Time         `xml:"LastModifiedDate,omitempty"`
	LastReferencedDate         time.Time         `xml:"LastReferencedDate,omitempty"`
	LastViewedDate             time.Time         `xml:"LastViewedDate,omitempty"`
	Leads                      *QueryResult      `xml:"Leads,omitempty"`
	ListEmailRecipientSources  *QueryResult      `xml:"ListEmailRecipientSources,omitempty"`
	ListEmails                 *QueryResult      `xml:"ListEmails,omitempty"`
	LookedUpFromActivities     *QueryResult      `xml:"LookedUpFromActivities,omitempty"`
	Name                       string            `xml:"Name,omitempty"`
	NumberOfContacts           int32             `xml:"NumberOfContacts,omitempty"`
	NumberOfConvertedLeads     int32             `xml:"NumberOfConvertedLeads,omitempty"`
	NumberOfLeads              int32             `xml:"NumberOfLeads,omitempty"`
	NumberOfOpportunities      int32             `xml:"NumberOfOpportunities,omitempty"`
	NumberOfResponses          int32             `xml:"NumberOfResponses,omitempty"`
	NumberOfWonOpportunities   int32             `xml:"NumberOfWonOpportunities,omitempty"`
	NumberSent                 float64           `xml:"NumberSent,omitempty"`
	OpenActivities             *QueryResult      `xml:"OpenActivities,omitempty"`
	Opportunities              *QueryResult      `xml:"Opportunities,omitempty"`
	Owner                      *User             `xml:"Owner,omitempty"`
	OwnerId                    *ID               `xml:"OwnerId,omitempty"`
	Parent                     *Campaign         `xml:"Parent,omitempty"`
	ParentId                   *ID               `xml:"ParentId,omitempty"`
	ProcessInstances           *QueryResult      `xml:"ProcessInstances,omitempty"`
	ProcessSteps               *QueryResult      `xml:"ProcessSteps,omitempty"`
	RecordAssociatedGroups     *QueryResult      `xml:"RecordAssociatedGroups,omitempty"`
	Shares                     *QueryResult      `xml:"Shares,omitempty"`
	StartDate                  time.Time         `xml:"StartDate,omitempty"`
	Status                     string            `xml:"Status,omitempty"`
	SystemModstamp             time.Time         `xml:"SystemModstamp,omitempty"`
	Tasks                      *QueryResult      `xml:"Tasks,omitempty"`
	TopicAssignments           *QueryResult      `xml:"TopicAssignments,omitempty"`
	Type                       string            `xml:"Type,omitempty"`
	UserRecordAccess           *UserRecordAccess `xml:"UserRecordAccess,omitempty"`
}

type CampaignFeed struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com CampaignFeed"`

	*SObject

	BestComment        *FeedComment `xml:"BestComment,omitempty"`
	BestCommentId      *ID          `xml:"BestCommentId,omitempty"`
	Body               string       `xml:"Body,omitempty"`
	CommentCount       int32        `xml:"CommentCount,omitempty"`
	CreatedBy          *SObject     `xml:"CreatedBy,omitempty"`
	CreatedById        *ID          `xml:"CreatedById,omitempty"`
	CreatedDate        time.Time    `xml:"CreatedDate,omitempty"`
	FeedAttachments    *QueryResult `xml:"FeedAttachments,omitempty"`
	FeedComments       *QueryResult `xml:"FeedComments,omitempty"`
	FeedLikes          *QueryResult `xml:"FeedLikes,omitempty"`
	FeedSignals        *QueryResult `xml:"FeedSignals,omitempty"`
	FeedTrackedChanges *QueryResult `xml:"FeedTrackedChanges,omitempty"`
	InsertedBy         *SObject     `xml:"InsertedBy,omitempty"`
	InsertedById       *ID          `xml:"InsertedById,omitempty"`
	IsDeleted          bool         `xml:"IsDeleted,omitempty"`
	IsRichText         bool         `xml:"IsRichText,omitempty"`
	LastModifiedDate   time.Time    `xml:"LastModifiedDate,omitempty"`
	LikeCount          int32        `xml:"LikeCount,omitempty"`
	LinkUrl            string       `xml:"LinkUrl,omitempty"`
	Parent             *Campaign    `xml:"Parent,omitempty"`
	ParentId           *ID          `xml:"ParentId,omitempty"`
	RelatedRecordId    *ID          `xml:"RelatedRecordId,omitempty"`
	SystemModstamp     time.Time    `xml:"SystemModstamp,omitempty"`
	Title              string       `xml:"Title,omitempty"`
	Type               string       `xml:"Type,omitempty"`
}

type CampaignHistory struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com CampaignHistory"`

	*SObject

	Campaign    *Campaign   `xml:"Campaign,omitempty"`
	CampaignId  *ID         `xml:"CampaignId,omitempty"`
	CreatedBy   *SObject    `xml:"CreatedBy,omitempty"`
	CreatedById *ID         `xml:"CreatedById,omitempty"`
	CreatedDate time.Time   `xml:"CreatedDate,omitempty"`
	Field       string      `xml:"Field,omitempty"`
	IsDeleted   bool        `xml:"IsDeleted,omitempty"`
	NewValue    interface{} `xml:"NewValue,omitempty"`
	OldValue    interface{} `xml:"OldValue,omitempty"`
}

type CampaignMember struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com CampaignMember"`

	*SObject

	Campaign                      *Campaign    `xml:"Campaign,omitempty"`
	CampaignId                    *ID          `xml:"CampaignId,omitempty"`
	City                          string       `xml:"City,omitempty"`
	CompanyOrAccount              string       `xml:"CompanyOrAccount,omitempty"`
	Contact                       *Contact     `xml:"Contact,omitempty"`
	ContactId                     *ID          `xml:"ContactId,omitempty"`
	Country                       string       `xml:"Country,omitempty"`
	CreatedBy                     *User        `xml:"CreatedBy,omitempty"`
	CreatedById                   *ID          `xml:"CreatedById,omitempty"`
	CreatedDate                   time.Time    `xml:"CreatedDate,omitempty"`
	Description                   string       `xml:"Description,omitempty"`
	DoNotCall                     bool         `xml:"DoNotCall,omitempty"`
	Email                         string       `xml:"Email,omitempty"`
	Fax                           string       `xml:"Fax,omitempty"`
	FirstName                     string       `xml:"FirstName,omitempty"`
	FirstRespondedDate            time.Time    `xml:"FirstRespondedDate,omitempty"`
	HasOptedOutOfEmail            bool         `xml:"HasOptedOutOfEmail,omitempty"`
	HasOptedOutOfFax              bool         `xml:"HasOptedOutOfFax,omitempty"`
	HasResponded                  bool         `xml:"HasResponded,omitempty"`
	IsDeleted                     bool         `xml:"IsDeleted,omitempty"`
	LastModifiedBy                *User        `xml:"LastModifiedBy,omitempty"`
	LastModifiedById              *ID          `xml:"LastModifiedById,omitempty"`
	LastModifiedDate              time.Time    `xml:"LastModifiedDate,omitempty"`
	LastName                      string       `xml:"LastName,omitempty"`
	Lead                          *Lead        `xml:"Lead,omitempty"`
	LeadId                        *ID          `xml:"LeadId,omitempty"`
	LeadOrContactId               *ID          `xml:"LeadOrContactId,omitempty"`
	LeadOrContactOwner            *SObject     `xml:"LeadOrContactOwner,omitempty"`
	LeadOrContactOwnerId          *ID          `xml:"LeadOrContactOwnerId,omitempty"`
	LeadSource                    string       `xml:"LeadSource,omitempty"`
	ListEmailIndividualRecipients *QueryResult `xml:"ListEmailIndividualRecipients,omitempty"`
	MobilePhone                   string       `xml:"MobilePhone,omitempty"`
	Name                          string       `xml:"Name,omitempty"`
	Phone                         string       `xml:"Phone,omitempty"`
	PostalCode                    string       `xml:"PostalCode,omitempty"`
	Salutation                    string       `xml:"Salutation,omitempty"`
	State                         string       `xml:"State,omitempty"`
	Status                        string       `xml:"Status,omitempty"`
	Street                        string       `xml:"Street,omitempty"`
	SystemModstamp                time.Time    `xml:"SystemModstamp,omitempty"`
	Title                         string       `xml:"Title,omitempty"`
	Type                          string       `xml:"Type,omitempty"`
}

type CampaignMemberStatus struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com CampaignMemberStatus"`

	*SObject

	CampaignId       *ID       `xml:"CampaignId,omitempty"`
	CreatedBy        *User     `xml:"CreatedBy,omitempty"`
	CreatedById      *ID       `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time `xml:"CreatedDate,omitempty"`
	HasResponded     bool      `xml:"HasResponded,omitempty"`
	IsDefault        bool      `xml:"IsDefault,omitempty"`
	IsDeleted        bool      `xml:"IsDeleted,omitempty"`
	Label            string    `xml:"Label,omitempty"`
	LastModifiedBy   *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time `xml:"LastModifiedDate,omitempty"`
	SortOrder        int32     `xml:"SortOrder,omitempty"`
	SystemModstamp   time.Time `xml:"SystemModstamp,omitempty"`
}

type CampaignShare struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com CampaignShare"`

	*SObject

	Campaign            *Campaign `xml:"Campaign,omitempty"`
	CampaignAccessLevel string    `xml:"CampaignAccessLevel,omitempty"`
	CampaignId          *ID       `xml:"CampaignId,omitempty"`
	IsDeleted           bool      `xml:"IsDeleted,omitempty"`
	LastModifiedBy      *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById    *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate    time.Time `xml:"LastModifiedDate,omitempty"`
	RowCause            string    `xml:"RowCause,omitempty"`
	UserOrGroup         *SObject  `xml:"UserOrGroup,omitempty"`
	UserOrGroupId       *ID       `xml:"UserOrGroupId,omitempty"`
}

type Case struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com Case"`

	*SObject

	Account                    *Account          `xml:"Account,omitempty"`
	AccountId                  *ID               `xml:"AccountId,omitempty"`
	ActivityHistories          *QueryResult      `xml:"ActivityHistories,omitempty"`
	Asset                      *Asset            `xml:"Asset,omitempty"`
	AssetId                    *ID               `xml:"AssetId,omitempty"`
	AttachedContentDocuments   *QueryResult      `xml:"AttachedContentDocuments,omitempty"`
	Attachments                *QueryResult      `xml:"Attachments,omitempty"`
	CaseComments               *QueryResult      `xml:"CaseComments,omitempty"`
	CaseContactRoles           *QueryResult      `xml:"CaseContactRoles,omitempty"`
	CaseNumber                 string            `xml:"CaseNumber,omitempty"`
	CaseSolutions              *QueryResult      `xml:"CaseSolutions,omitempty"`
	Cases                      *QueryResult      `xml:"Cases,omitempty"`
	ClosedDate                 time.Time         `xml:"ClosedDate,omitempty"`
	CombinedAttachments        *QueryResult      `xml:"CombinedAttachments,omitempty"`
	Comments                   string            `xml:"Comments,omitempty"`
	Contact                    *Contact          `xml:"Contact,omitempty"`
	ContactEmail               string            `xml:"ContactEmail,omitempty"`
	ContactFax                 string            `xml:"ContactFax,omitempty"`
	ContactId                  *ID               `xml:"ContactId,omitempty"`
	ContactMobile              string            `xml:"ContactMobile,omitempty"`
	ContactPhone               string            `xml:"ContactPhone,omitempty"`
	ContentDocumentLinks       *QueryResult      `xml:"ContentDocumentLinks,omitempty"`
	CreatedBy                  *User             `xml:"CreatedBy,omitempty"`
	CreatedById                *ID               `xml:"CreatedById,omitempty"`
	CreatedDate                time.Time         `xml:"CreatedDate,omitempty"`
	Description                string            `xml:"Description,omitempty"`
	EmailMessages              *QueryResult      `xml:"EmailMessages,omitempty"`
	Emails                     *QueryResult      `xml:"Emails,omitempty"`
	EngineeringReqNumberc      string            `xml:"EngineeringReqNumber__c,omitempty"`
	Events                     *QueryResult      `xml:"Events,omitempty"`
	FeedSubscriptionsForEntity *QueryResult      `xml:"FeedSubscriptionsForEntity,omitempty"`
	Feeds                      *QueryResult      `xml:"Feeds,omitempty"`
	Histories                  *QueryResult      `xml:"Histories,omitempty"`
	IsClosed                   bool              `xml:"IsClosed,omitempty"`
	IsDeleted                  bool              `xml:"IsDeleted,omitempty"`
	IsEscalated                bool              `xml:"IsEscalated,omitempty"`
	LastModifiedBy             *User             `xml:"LastModifiedBy,omitempty"`
	LastModifiedById           *ID               `xml:"LastModifiedById,omitempty"`
	LastModifiedDate           time.Time         `xml:"LastModifiedDate,omitempty"`
	LastReferencedDate         time.Time         `xml:"LastReferencedDate,omitempty"`
	LastViewedDate             time.Time         `xml:"LastViewedDate,omitempty"`
	LookedUpFromActivities     *QueryResult      `xml:"LookedUpFromActivities,omitempty"`
	OpenActivities             *QueryResult      `xml:"OpenActivities,omitempty"`
	Origin                     string            `xml:"Origin,omitempty"`
	Owner                      *SObject          `xml:"Owner,omitempty"`
	OwnerId                    *ID               `xml:"OwnerId,omitempty"`
	Parent                     *Case             `xml:"Parent,omitempty"`
	ParentId                   *ID               `xml:"ParentId,omitempty"`
	PotentialLiabilityc        string            `xml:"PotentialLiability__c,omitempty"`
	Priority                   string            `xml:"Priority,omitempty"`
	ProcessInstances           *QueryResult      `xml:"ProcessInstances,omitempty"`
	ProcessSteps               *QueryResult      `xml:"ProcessSteps,omitempty"`
	Productc                   string            `xml:"Product__c,omitempty"`
	Reason                     string            `xml:"Reason,omitempty"`
	RecordActionHistories      *QueryResult      `xml:"RecordActionHistories,omitempty"`
	RecordActions              *QueryResult      `xml:"RecordActions,omitempty"`
	RecordAssociatedGroups     *QueryResult      `xml:"RecordAssociatedGroups,omitempty"`
	SLAViolationc              string            `xml:"SLAViolation__c,omitempty"`
	Shares                     *QueryResult      `xml:"Shares,omitempty"`
	Solutions                  *QueryResult      `xml:"Solutions,omitempty"`
	Status                     string            `xml:"Status,omitempty"`
	Subject                    string            `xml:"Subject,omitempty"`
	SuppliedCompany            string            `xml:"SuppliedCompany,omitempty"`
	SuppliedEmail              string            `xml:"SuppliedEmail,omitempty"`
	SuppliedName               string            `xml:"SuppliedName,omitempty"`
	SuppliedPhone              string            `xml:"SuppliedPhone,omitempty"`
	SystemModstamp             time.Time         `xml:"SystemModstamp,omitempty"`
	Tasks                      *QueryResult      `xml:"Tasks,omitempty"`
	TeamMembers                *QueryResult      `xml:"TeamMembers,omitempty"`
	TeamTemplateRecords        *QueryResult      `xml:"TeamTemplateRecords,omitempty"`
	TopicAssignments           *QueryResult      `xml:"TopicAssignments,omitempty"`
	Type                       string            `xml:"Type,omitempty"`
	UserRecordAccess           *UserRecordAccess `xml:"UserRecordAccess,omitempty"`
}

type CaseComment struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com CaseComment"`

	*SObject

	CommentBody      string    `xml:"CommentBody,omitempty"`
	CreatedBy        *SObject  `xml:"CreatedBy,omitempty"`
	CreatedById      *ID       `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time `xml:"CreatedDate,omitempty"`
	IsDeleted        bool      `xml:"IsDeleted,omitempty"`
	IsPublished      bool      `xml:"IsPublished,omitempty"`
	LastModifiedBy   *SObject  `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time `xml:"LastModifiedDate,omitempty"`
	Parent           *Case     `xml:"Parent,omitempty"`
	ParentId         *ID       `xml:"ParentId,omitempty"`
	SystemModstamp   time.Time `xml:"SystemModstamp,omitempty"`
}

type CaseContactRole struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com CaseContactRole"`

	*SObject

	Cases            *Case     `xml:"Cases,omitempty"`
	CasesId          *ID       `xml:"CasesId,omitempty"`
	Contact          *Contact  `xml:"Contact,omitempty"`
	ContactId        *ID       `xml:"ContactId,omitempty"`
	CreatedBy        *User     `xml:"CreatedBy,omitempty"`
	CreatedById      *ID       `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time `xml:"CreatedDate,omitempty"`
	IsDeleted        bool      `xml:"IsDeleted,omitempty"`
	LastModifiedBy   *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time `xml:"LastModifiedDate,omitempty"`
	Role             string    `xml:"Role,omitempty"`
	SystemModstamp   time.Time `xml:"SystemModstamp,omitempty"`
}

type CaseFeed struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com CaseFeed"`

	*SObject

	BestComment        *FeedComment `xml:"BestComment,omitempty"`
	BestCommentId      *ID          `xml:"BestCommentId,omitempty"`
	Body               string       `xml:"Body,omitempty"`
	CommentCount       int32        `xml:"CommentCount,omitempty"`
	CreatedBy          *SObject     `xml:"CreatedBy,omitempty"`
	CreatedById        *ID          `xml:"CreatedById,omitempty"`
	CreatedDate        time.Time    `xml:"CreatedDate,omitempty"`
	FeedAttachments    *QueryResult `xml:"FeedAttachments,omitempty"`
	FeedComments       *QueryResult `xml:"FeedComments,omitempty"`
	FeedLikes          *QueryResult `xml:"FeedLikes,omitempty"`
	FeedSignals        *QueryResult `xml:"FeedSignals,omitempty"`
	FeedTrackedChanges *QueryResult `xml:"FeedTrackedChanges,omitempty"`
	InsertedBy         *SObject     `xml:"InsertedBy,omitempty"`
	InsertedById       *ID          `xml:"InsertedById,omitempty"`
	IsDeleted          bool         `xml:"IsDeleted,omitempty"`
	IsRichText         bool         `xml:"IsRichText,omitempty"`
	LastModifiedDate   time.Time    `xml:"LastModifiedDate,omitempty"`
	LikeCount          int32        `xml:"LikeCount,omitempty"`
	LinkUrl            string       `xml:"LinkUrl,omitempty"`
	Parent             *Case        `xml:"Parent,omitempty"`
	ParentId           *ID          `xml:"ParentId,omitempty"`
	RelatedRecordId    *ID          `xml:"RelatedRecordId,omitempty"`
	SystemModstamp     time.Time    `xml:"SystemModstamp,omitempty"`
	Title              string       `xml:"Title,omitempty"`
	Type               string       `xml:"Type,omitempty"`
}

type CaseHistory struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com CaseHistory"`

	*SObject

	Case        *Case       `xml:"Case,omitempty"`
	CaseId      *ID         `xml:"CaseId,omitempty"`
	CreatedBy   *SObject    `xml:"CreatedBy,omitempty"`
	CreatedById *ID         `xml:"CreatedById,omitempty"`
	CreatedDate time.Time   `xml:"CreatedDate,omitempty"`
	Field       string      `xml:"Field,omitempty"`
	IsDeleted   bool        `xml:"IsDeleted,omitempty"`
	NewValue    interface{} `xml:"NewValue,omitempty"`
	OldValue    interface{} `xml:"OldValue,omitempty"`
}

type CaseShare struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com CaseShare"`

	*SObject

	Case             *Case     `xml:"Case,omitempty"`
	CaseAccessLevel  string    `xml:"CaseAccessLevel,omitempty"`
	CaseId           *ID       `xml:"CaseId,omitempty"`
	IsDeleted        bool      `xml:"IsDeleted,omitempty"`
	LastModifiedBy   *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time `xml:"LastModifiedDate,omitempty"`
	RowCause         string    `xml:"RowCause,omitempty"`
	UserOrGroup      *SObject  `xml:"UserOrGroup,omitempty"`
	UserOrGroupId    *ID       `xml:"UserOrGroupId,omitempty"`
}

type CaseSolution struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com CaseSolution"`

	*SObject

	Case           *Case     `xml:"Case,omitempty"`
	CaseId         *ID       `xml:"CaseId,omitempty"`
	CreatedBy      *User     `xml:"CreatedBy,omitempty"`
	CreatedById    *ID       `xml:"CreatedById,omitempty"`
	CreatedDate    time.Time `xml:"CreatedDate,omitempty"`
	IsDeleted      bool      `xml:"IsDeleted,omitempty"`
	Solution       *Solution `xml:"Solution,omitempty"`
	SolutionId     *ID       `xml:"SolutionId,omitempty"`
	SystemModstamp time.Time `xml:"SystemModstamp,omitempty"`
}

type CaseStatus struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com CaseStatus"`

	*SObject

	ApiName          string    `xml:"ApiName,omitempty"`
	CreatedBy        *User     `xml:"CreatedBy,omitempty"`
	CreatedById      *ID       `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time `xml:"CreatedDate,omitempty"`
	IsClosed         bool      `xml:"IsClosed,omitempty"`
	IsDefault        bool      `xml:"IsDefault,omitempty"`
	LastModifiedBy   *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time `xml:"LastModifiedDate,omitempty"`
	MasterLabel      string    `xml:"MasterLabel,omitempty"`
	SortOrder        int32     `xml:"SortOrder,omitempty"`
	SystemModstamp   time.Time `xml:"SystemModstamp,omitempty"`
}

type CaseTeamMember struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com CaseTeamMember"`

	*SObject

	CreatedBy            *User                   `xml:"CreatedBy,omitempty"`
	CreatedById          *ID                     `xml:"CreatedById,omitempty"`
	CreatedDate          time.Time               `xml:"CreatedDate,omitempty"`
	LastModifiedBy       *User                   `xml:"LastModifiedBy,omitempty"`
	LastModifiedById     *ID                     `xml:"LastModifiedById,omitempty"`
	LastModifiedDate     time.Time               `xml:"LastModifiedDate,omitempty"`
	Member               *SObject                `xml:"Member,omitempty"`
	MemberId             *ID                     `xml:"MemberId,omitempty"`
	Parent               *Case                   `xml:"Parent,omitempty"`
	ParentId             *ID                     `xml:"ParentId,omitempty"`
	SystemModstamp       time.Time               `xml:"SystemModstamp,omitempty"`
	TeamRole             *CaseTeamRole           `xml:"TeamRole,omitempty"`
	TeamRoleId           *ID                     `xml:"TeamRoleId,omitempty"`
	TeamTemplate         *CaseTeamTemplate       `xml:"TeamTemplate,omitempty"`
	TeamTemplateId       *ID                     `xml:"TeamTemplateId,omitempty"`
	TeamTemplateMember   *CaseTeamTemplateMember `xml:"TeamTemplateMember,omitempty"`
	TeamTemplateMemberId *ID                     `xml:"TeamTemplateMemberId,omitempty"`
}

type CaseTeamRole struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com CaseTeamRole"`

	*SObject

	AccessLevel             string    `xml:"AccessLevel,omitempty"`
	CreatedBy               *User     `xml:"CreatedBy,omitempty"`
	CreatedById             *ID       `xml:"CreatedById,omitempty"`
	CreatedDate             time.Time `xml:"CreatedDate,omitempty"`
	LastModifiedBy          *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById        *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate        time.Time `xml:"LastModifiedDate,omitempty"`
	Name                    string    `xml:"Name,omitempty"`
	PreferencesVisibleInCSP bool      `xml:"PreferencesVisibleInCSP,omitempty"`
	SystemModstamp          time.Time `xml:"SystemModstamp,omitempty"`
}

type CaseTeamTemplate struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com CaseTeamTemplate"`

	*SObject

	CreatedBy        *User     `xml:"CreatedBy,omitempty"`
	CreatedById      *ID       `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time `xml:"CreatedDate,omitempty"`
	Description      string    `xml:"Description,omitempty"`
	LastModifiedBy   *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time `xml:"LastModifiedDate,omitempty"`
	Name             string    `xml:"Name,omitempty"`
	SystemModstamp   time.Time `xml:"SystemModstamp,omitempty"`
}

type CaseTeamTemplateMember struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com CaseTeamTemplateMember"`

	*SObject

	CreatedBy        *User     `xml:"CreatedBy,omitempty"`
	CreatedById      *ID       `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time `xml:"CreatedDate,omitempty"`
	LastModifiedBy   *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time `xml:"LastModifiedDate,omitempty"`
	MemberId         *ID       `xml:"MemberId,omitempty"`
	SystemModstamp   time.Time `xml:"SystemModstamp,omitempty"`
	TeamRoleId       *ID       `xml:"TeamRoleId,omitempty"`
	TeamTemplateId   *ID       `xml:"TeamTemplateId,omitempty"`
}

type CaseTeamTemplateRecord struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com CaseTeamTemplateRecord"`

	*SObject

	CreatedBy      *User             `xml:"CreatedBy,omitempty"`
	CreatedById    *ID               `xml:"CreatedById,omitempty"`
	CreatedDate    time.Time         `xml:"CreatedDate,omitempty"`
	Parent         *Case             `xml:"Parent,omitempty"`
	ParentId       *ID               `xml:"ParentId,omitempty"`
	SystemModstamp time.Time         `xml:"SystemModstamp,omitempty"`
	TeamTemplate   *CaseTeamTemplate `xml:"TeamTemplate,omitempty"`
	TeamTemplateId *ID               `xml:"TeamTemplateId,omitempty"`
}

type CategoryData struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com CategoryData"`

	*SObject

	CategoryNodeId   *ID       `xml:"CategoryNodeId,omitempty"`
	CreatedBy        *User     `xml:"CreatedBy,omitempty"`
	CreatedById      *ID       `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time `xml:"CreatedDate,omitempty"`
	IsDeleted        bool      `xml:"IsDeleted,omitempty"`
	LastModifiedBy   *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time `xml:"LastModifiedDate,omitempty"`
	RelatedSobjectId *ID       `xml:"RelatedSobjectId,omitempty"`
	SystemModstamp   time.Time `xml:"SystemModstamp,omitempty"`
}

type CategoryNode struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com CategoryNode"`

	*SObject

	CreatedBy        *User     `xml:"CreatedBy,omitempty"`
	CreatedById      *ID       `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time `xml:"CreatedDate,omitempty"`
	LastModifiedBy   *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time `xml:"LastModifiedDate,omitempty"`
	MasterLabel      string    `xml:"MasterLabel,omitempty"`
	ParentId         *ID       `xml:"ParentId,omitempty"`
	SortOrder        int32     `xml:"SortOrder,omitempty"`
	SortStyle        string    `xml:"SortStyle,omitempty"`
	SystemModstamp   time.Time `xml:"SystemModstamp,omitempty"`
}

type ChatterActivity struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ChatterActivity"`

	*SObject

	CommentCount         int32     `xml:"CommentCount,omitempty"`
	CommentReceivedCount int32     `xml:"CommentReceivedCount,omitempty"`
	InfluenceRawRank     int32     `xml:"InfluenceRawRank,omitempty"`
	LikeReceivedCount    int32     `xml:"LikeReceivedCount,omitempty"`
	ParentId             *ID       `xml:"ParentId,omitempty"`
	PostCount            int32     `xml:"PostCount,omitempty"`
	SystemModstamp       time.Time `xml:"SystemModstamp,omitempty"`
}

type ChatterExtension struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ChatterExtension"`

	*SObject

	CompositionComponentEnumOrId string        `xml:"CompositionComponentEnumOrId,omitempty"`
	CreatedBy                    *User         `xml:"CreatedBy,omitempty"`
	CreatedById                  *ID           `xml:"CreatedById,omitempty"`
	CreatedDate                  time.Time     `xml:"CreatedDate,omitempty"`
	Description                  string        `xml:"Description,omitempty"`
	DeveloperName                string        `xml:"DeveloperName,omitempty"`
	ExtensionName                string        `xml:"ExtensionName,omitempty"`
	HeaderText                   string        `xml:"HeaderText,omitempty"`
	HoverText                    string        `xml:"HoverText,omitempty"`
	Icon                         *ContentAsset `xml:"Icon,omitempty"`
	IconId                       *ID           `xml:"IconId,omitempty"`
	IsDeleted                    bool          `xml:"IsDeleted,omitempty"`
	IsProtected                  bool          `xml:"IsProtected,omitempty"`
	Language                     string        `xml:"Language,omitempty"`
	LastModifiedBy               *User         `xml:"LastModifiedBy,omitempty"`
	LastModifiedById             *ID           `xml:"LastModifiedById,omitempty"`
	LastModifiedDate             time.Time     `xml:"LastModifiedDate,omitempty"`
	MasterLabel                  string        `xml:"MasterLabel,omitempty"`
	NamespacePrefix              string        `xml:"NamespacePrefix,omitempty"`
	RenderComponentEnumOrId      string        `xml:"RenderComponentEnumOrId,omitempty"`
	SystemModstamp               time.Time     `xml:"SystemModstamp,omitempty"`
	Type                         string        `xml:"Type,omitempty"`
}

type ChatterExtensionConfig struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ChatterExtensionConfig"`

	*SObject

	CanCreate          bool              `xml:"CanCreate,omitempty"`
	CanRead            bool              `xml:"CanRead,omitempty"`
	ChatterExtension   *ChatterExtension `xml:"ChatterExtension,omitempty"`
	ChatterExtensionId *ID               `xml:"ChatterExtensionId,omitempty"`
	CreatedBy          *User             `xml:"CreatedBy,omitempty"`
	CreatedById        *ID               `xml:"CreatedById,omitempty"`
	CreatedDate        time.Time         `xml:"CreatedDate,omitempty"`
	IsDeleted          bool              `xml:"IsDeleted,omitempty"`
	LastModifiedBy     *User             `xml:"LastModifiedBy,omitempty"`
	LastModifiedById   *ID               `xml:"LastModifiedById,omitempty"`
	LastModifiedDate   time.Time         `xml:"LastModifiedDate,omitempty"`
	Position           int32             `xml:"Position,omitempty"`
	SystemModstamp     time.Time         `xml:"SystemModstamp,omitempty"`
}

type ClientBrowser struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ClientBrowser"`

	*SObject

	CreatedDate   time.Time `xml:"CreatedDate,omitempty"`
	FullUserAgent string    `xml:"FullUserAgent,omitempty"`
	LastUpdate    time.Time `xml:"LastUpdate,omitempty"`
	ProxyInfo     string    `xml:"ProxyInfo,omitempty"`
	Users         *User     `xml:"Users,omitempty"`
	UsersId       *ID       `xml:"UsersId,omitempty"`
}

type CollaborationGroup struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com CollaborationGroup"`

	*SObject

	Announcement               *Announcement `xml:"Announcement,omitempty"`
	AnnouncementId             *ID           `xml:"AnnouncementId,omitempty"`
	AttachedContentDocuments   *QueryResult  `xml:"AttachedContentDocuments,omitempty"`
	BannerPhotoUrl             string        `xml:"BannerPhotoUrl,omitempty"`
	CanHaveGuests              bool          `xml:"CanHaveGuests,omitempty"`
	CollaborationGroupRecords  *QueryResult  `xml:"CollaborationGroupRecords,omitempty"`
	CollaborationType          string        `xml:"CollaborationType,omitempty"`
	CombinedAttachments        *QueryResult  `xml:"CombinedAttachments,omitempty"`
	ContentDocumentLinks       *QueryResult  `xml:"ContentDocumentLinks,omitempty"`
	CreatedBy                  *User         `xml:"CreatedBy,omitempty"`
	CreatedById                *ID           `xml:"CreatedById,omitempty"`
	CreatedDate                time.Time     `xml:"CreatedDate,omitempty"`
	Description                string        `xml:"Description,omitempty"`
	FeedSubscriptionsForEntity *QueryResult  `xml:"FeedSubscriptionsForEntity,omitempty"`
	Feeds                      *QueryResult  `xml:"Feeds,omitempty"`
	FullPhotoUrl               string        `xml:"FullPhotoUrl,omitempty"`
	GroupEmail                 string        `xml:"GroupEmail,omitempty"`
	GroupMemberRequests        *QueryResult  `xml:"GroupMemberRequests,omitempty"`
	GroupMembers               *QueryResult  `xml:"GroupMembers,omitempty"`
	HasPrivateFieldsAccess     bool          `xml:"HasPrivateFieldsAccess,omitempty"`
	InformationBody            string        `xml:"InformationBody,omitempty"`
	InformationTitle           string        `xml:"InformationTitle,omitempty"`
	IsArchived                 bool          `xml:"IsArchived,omitempty"`
	IsAutoArchiveDisabled      bool          `xml:"IsAutoArchiveDisabled,omitempty"`
	IsBroadcast                bool          `xml:"IsBroadcast,omitempty"`
	LastFeedModifiedDate       time.Time     `xml:"LastFeedModifiedDate,omitempty"`
	LastModifiedBy             *User         `xml:"LastModifiedBy,omitempty"`
	LastModifiedById           *ID           `xml:"LastModifiedById,omitempty"`
	LastModifiedDate           time.Time     `xml:"LastModifiedDate,omitempty"`
	LastReferencedDate         time.Time     `xml:"LastReferencedDate,omitempty"`
	LastViewedDate             time.Time     `xml:"LastViewedDate,omitempty"`
	MediumPhotoUrl             string        `xml:"MediumPhotoUrl,omitempty"`
	MemberCount                int32         `xml:"MemberCount,omitempty"`
	Name                       string        `xml:"Name,omitempty"`
	Owner                      *User         `xml:"Owner,omitempty"`
	OwnerId                    *ID           `xml:"OwnerId,omitempty"`
	SmallPhotoUrl              string        `xml:"SmallPhotoUrl,omitempty"`
	SystemModstamp             time.Time     `xml:"SystemModstamp,omitempty"`
}

type CollaborationGroupFeed struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com CollaborationGroupFeed"`

	*SObject

	BestComment        *FeedComment        `xml:"BestComment,omitempty"`
	BestCommentId      *ID                 `xml:"BestCommentId,omitempty"`
	Body               string              `xml:"Body,omitempty"`
	CommentCount       int32               `xml:"CommentCount,omitempty"`
	CreatedBy          *SObject            `xml:"CreatedBy,omitempty"`
	CreatedById        *ID                 `xml:"CreatedById,omitempty"`
	CreatedDate        time.Time           `xml:"CreatedDate,omitempty"`
	FeedAttachments    *QueryResult        `xml:"FeedAttachments,omitempty"`
	FeedComments       *QueryResult        `xml:"FeedComments,omitempty"`
	FeedLikes          *QueryResult        `xml:"FeedLikes,omitempty"`
	FeedSignals        *QueryResult        `xml:"FeedSignals,omitempty"`
	FeedTrackedChanges *QueryResult        `xml:"FeedTrackedChanges,omitempty"`
	InsertedBy         *SObject            `xml:"InsertedBy,omitempty"`
	InsertedById       *ID                 `xml:"InsertedById,omitempty"`
	IsDeleted          bool                `xml:"IsDeleted,omitempty"`
	IsRichText         bool                `xml:"IsRichText,omitempty"`
	LastModifiedDate   time.Time           `xml:"LastModifiedDate,omitempty"`
	LikeCount          int32               `xml:"LikeCount,omitempty"`
	LinkUrl            string              `xml:"LinkUrl,omitempty"`
	Parent             *CollaborationGroup `xml:"Parent,omitempty"`
	ParentId           *ID                 `xml:"ParentId,omitempty"`
	RelatedRecordId    *ID                 `xml:"RelatedRecordId,omitempty"`
	SystemModstamp     time.Time           `xml:"SystemModstamp,omitempty"`
	Title              string              `xml:"Title,omitempty"`
	Type               string              `xml:"Type,omitempty"`
}

type CollaborationGroupMember struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com CollaborationGroupMember"`

	*SObject

	CollaborationGroup    *CollaborationGroup `xml:"CollaborationGroup,omitempty"`
	CollaborationGroupId  *ID                 `xml:"CollaborationGroupId,omitempty"`
	CollaborationRole     string              `xml:"CollaborationRole,omitempty"`
	CreatedBy             *User               `xml:"CreatedBy,omitempty"`
	CreatedById           *ID                 `xml:"CreatedById,omitempty"`
	CreatedDate           time.Time           `xml:"CreatedDate,omitempty"`
	LastFeedAccessDate    time.Time           `xml:"LastFeedAccessDate,omitempty"`
	LastModifiedBy        *User               `xml:"LastModifiedBy,omitempty"`
	LastModifiedById      *ID                 `xml:"LastModifiedById,omitempty"`
	LastModifiedDate      time.Time           `xml:"LastModifiedDate,omitempty"`
	Member                *User               `xml:"Member,omitempty"`
	MemberId              *ID                 `xml:"MemberId,omitempty"`
	NotificationFrequency string              `xml:"NotificationFrequency,omitempty"`
	SystemModstamp        time.Time           `xml:"SystemModstamp,omitempty"`
}

type CollaborationGroupMemberRequest struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com CollaborationGroupMemberRequest"`

	*SObject

	CollaborationGroup   *CollaborationGroup `xml:"CollaborationGroup,omitempty"`
	CollaborationGroupId *ID                 `xml:"CollaborationGroupId,omitempty"`
	CreatedBy            *User               `xml:"CreatedBy,omitempty"`
	CreatedById          *ID                 `xml:"CreatedById,omitempty"`
	CreatedDate          time.Time           `xml:"CreatedDate,omitempty"`
	LastModifiedBy       *User               `xml:"LastModifiedBy,omitempty"`
	LastModifiedById     *ID                 `xml:"LastModifiedById,omitempty"`
	LastModifiedDate     time.Time           `xml:"LastModifiedDate,omitempty"`
	Requester            *User               `xml:"Requester,omitempty"`
	RequesterId          *ID                 `xml:"RequesterId,omitempty"`
	ResponseMessage      string              `xml:"ResponseMessage,omitempty"`
	Status               string              `xml:"Status,omitempty"`
	SystemModstamp       time.Time           `xml:"SystemModstamp,omitempty"`
}

type CollaborationGroupRecord struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com CollaborationGroupRecord"`

	*SObject

	CollaborationGroup   *CollaborationGroup `xml:"CollaborationGroup,omitempty"`
	CollaborationGroupId *ID                 `xml:"CollaborationGroupId,omitempty"`
	CreatedBy            *User               `xml:"CreatedBy,omitempty"`
	CreatedById          *ID                 `xml:"CreatedById,omitempty"`
	CreatedDate          time.Time           `xml:"CreatedDate,omitempty"`
	IsDeleted            bool                `xml:"IsDeleted,omitempty"`
	LastModifiedBy       *User               `xml:"LastModifiedBy,omitempty"`
	LastModifiedById     *ID                 `xml:"LastModifiedById,omitempty"`
	LastModifiedDate     time.Time           `xml:"LastModifiedDate,omitempty"`
	Record               *SObject            `xml:"Record,omitempty"`
	RecordId             *ID                 `xml:"RecordId,omitempty"`
	SystemModstamp       time.Time           `xml:"SystemModstamp,omitempty"`
	UserRecordAccess     *UserRecordAccess   `xml:"UserRecordAccess,omitempty"`
}

type CollaborationInvitation struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com CollaborationInvitation"`

	*SObject

	CreatedBy                  *User     `xml:"CreatedBy,omitempty"`
	CreatedById                *ID       `xml:"CreatedById,omitempty"`
	CreatedDate                time.Time `xml:"CreatedDate,omitempty"`
	InvitedUserEmail           string    `xml:"InvitedUserEmail,omitempty"`
	InvitedUserEmailNormalized string    `xml:"InvitedUserEmailNormalized,omitempty"`
	InviterId                  *ID       `xml:"InviterId,omitempty"`
	LastModifiedBy             *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById           *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate           time.Time `xml:"LastModifiedDate,omitempty"`
	OptionalMessage            string    `xml:"OptionalMessage,omitempty"`
	ParentId                   *ID       `xml:"ParentId,omitempty"`
	SharedEntityId             *ID       `xml:"SharedEntityId,omitempty"`
	Status                     string    `xml:"Status,omitempty"`
	SystemModstamp             time.Time `xml:"SystemModstamp,omitempty"`
}

type ColorDefinition struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ColorDefinition"`

	*SObject

	Color           string         `xml:"Color,omitempty"`
	Context         string         `xml:"Context,omitempty"`
	DurableId       string         `xml:"DurableId,omitempty"`
	TabDefinition   *TabDefinition `xml:"TabDefinition,omitempty"`
	TabDefinitionId string         `xml:"TabDefinitionId,omitempty"`
	Theme           string         `xml:"Theme,omitempty"`
}

type CombinedAttachment struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com CombinedAttachment"`

	*SObject

	ContentSize            int32     `xml:"ContentSize,omitempty"`
	ContentUrl             string    `xml:"ContentUrl,omitempty"`
	CreatedBy              *User     `xml:"CreatedBy,omitempty"`
	CreatedById            *ID       `xml:"CreatedById,omitempty"`
	CreatedDate            time.Time `xml:"CreatedDate,omitempty"`
	ExternalDataSourceName string    `xml:"ExternalDataSourceName,omitempty"`
	ExternalDataSourceType string    `xml:"ExternalDataSourceType,omitempty"`
	FileExtension          string    `xml:"FileExtension,omitempty"`
	FileType               string    `xml:"FileType,omitempty"`
	IsDeleted              bool      `xml:"IsDeleted,omitempty"`
	LastModifiedBy         *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById       *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate       time.Time `xml:"LastModifiedDate,omitempty"`
	Parent                 *Contract `xml:"Parent,omitempty"`
	ParentId               *ID       `xml:"ParentId,omitempty"`
	RecordType             string    `xml:"RecordType,omitempty"`
	SharingOption          string    `xml:"SharingOption,omitempty"`
	Title                  string    `xml:"Title,omitempty"`
}

type Community struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com Community"`

	*SObject

	CreatedBy        *User     `xml:"CreatedBy,omitempty"`
	CreatedById      *ID       `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time `xml:"CreatedDate,omitempty"`
	Description      string    `xml:"Description,omitempty"`
	IsActive         bool      `xml:"IsActive,omitempty"`
	IsPublished      bool      `xml:"IsPublished,omitempty"`
	LastModifiedBy   *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time `xml:"LastModifiedDate,omitempty"`
	Name             string    `xml:"Name,omitempty"`
	SystemModstamp   time.Time `xml:"SystemModstamp,omitempty"`
}

type ConnectedApplication struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ConnectedApplication"`

	*SObject

	CreatedBy                           *User        `xml:"CreatedBy,omitempty"`
	CreatedById                         *ID          `xml:"CreatedById,omitempty"`
	CreatedDate                         time.Time    `xml:"CreatedDate,omitempty"`
	InstalledMobileApps                 *QueryResult `xml:"InstalledMobileApps,omitempty"`
	LastModifiedBy                      *User        `xml:"LastModifiedBy,omitempty"`
	LastModifiedById                    *ID          `xml:"LastModifiedById,omitempty"`
	LastModifiedDate                    time.Time    `xml:"LastModifiedDate,omitempty"`
	MobileSessionTimeout                string       `xml:"MobileSessionTimeout,omitempty"`
	MobileStartUrl                      string       `xml:"MobileStartUrl,omitempty"`
	Name                                string       `xml:"Name,omitempty"`
	OptionsAllowAdminApprovedUsersOnly  bool         `xml:"OptionsAllowAdminApprovedUsersOnly,omitempty"`
	OptionsFullContentPushNotifications bool         `xml:"OptionsFullContentPushNotifications,omitempty"`
	OptionsHasSessionLevelPolicy        bool         `xml:"OptionsHasSessionLevelPolicy,omitempty"`
	OptionsIsInternal                   bool         `xml:"OptionsIsInternal,omitempty"`
	OptionsRefreshTokenValidityMetric   bool         `xml:"OptionsRefreshTokenValidityMetric,omitempty"`
	PinLength                           string       `xml:"PinLength,omitempty"`
	RefreshTokenValidityPeriod          int32        `xml:"RefreshTokenValidityPeriod,omitempty"`
	SetupEntityAccessItems              *QueryResult `xml:"SetupEntityAccessItems,omitempty"`
	StartUrl                            string       `xml:"StartUrl,omitempty"`
	SystemModstamp                      time.Time    `xml:"SystemModstamp,omitempty"`
}

type Contact struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com Contact"`

	*SObject

	AcceptedEventRelations        *QueryResult      `xml:"AcceptedEventRelations,omitempty"`
	Account                       *Account          `xml:"Account,omitempty"`
	AccountContactRoles           *QueryResult      `xml:"AccountContactRoles,omitempty"`
	AccountId                     *ID               `xml:"AccountId,omitempty"`
	ActivityHistories             *QueryResult      `xml:"ActivityHistories,omitempty"`
	Assets                        *QueryResult      `xml:"Assets,omitempty"`
	AssistantName                 string            `xml:"AssistantName,omitempty"`
	AssistantPhone                string            `xml:"AssistantPhone,omitempty"`
	AttachedContentDocuments      *QueryResult      `xml:"AttachedContentDocuments,omitempty"`
	Attachments                   *QueryResult      `xml:"Attachments,omitempty"`
	Birthdate                     time.Time         `xml:"Birthdate,omitempty"`
	CampaignMembers               *QueryResult      `xml:"CampaignMembers,omitempty"`
	CaseContactRoles              *QueryResult      `xml:"CaseContactRoles,omitempty"`
	Cases                         *QueryResult      `xml:"Cases,omitempty"`
	CleanStatus                   string            `xml:"CleanStatus,omitempty"`
	CombinedAttachments           *QueryResult      `xml:"CombinedAttachments,omitempty"`
	ContactCleanInfos             *QueryResult      `xml:"ContactCleanInfos,omitempty"`
	ContentDocumentLinks          *QueryResult      `xml:"ContentDocumentLinks,omitempty"`
	ContractContactRoles          *QueryResult      `xml:"ContractContactRoles,omitempty"`
	ContractsSigned               *QueryResult      `xml:"ContractsSigned,omitempty"`
	CreatedBy                     *User             `xml:"CreatedBy,omitempty"`
	CreatedById                   *ID               `xml:"CreatedById,omitempty"`
	CreatedDate                   time.Time         `xml:"CreatedDate,omitempty"`
	DeclinedEventRelations        *QueryResult      `xml:"DeclinedEventRelations,omitempty"`
	Department                    string            `xml:"Department,omitempty"`
	Description                   string            `xml:"Description,omitempty"`
	DuplicateRecordItems          *QueryResult      `xml:"DuplicateRecordItems,omitempty"`
	Email                         string            `xml:"Email,omitempty"`
	EmailBouncedDate              time.Time         `xml:"EmailBouncedDate,omitempty"`
	EmailBouncedReason            string            `xml:"EmailBouncedReason,omitempty"`
	EmailMessageRelations         *QueryResult      `xml:"EmailMessageRelations,omitempty"`
	EmailStatuses                 *QueryResult      `xml:"EmailStatuses,omitempty"`
	EventRelations                *QueryResult      `xml:"EventRelations,omitempty"`
	Events                        *QueryResult      `xml:"Events,omitempty"`
	Fax                           string            `xml:"Fax,omitempty"`
	FeedSubscriptionsForEntity    *QueryResult      `xml:"FeedSubscriptionsForEntity,omitempty"`
	Feeds                         *QueryResult      `xml:"Feeds,omitempty"`
	FirstName                     string            `xml:"FirstName,omitempty"`
	Histories                     *QueryResult      `xml:"Histories,omitempty"`
	HomePhone                     string            `xml:"HomePhone,omitempty"`
	IsDeleted                     bool              `xml:"IsDeleted,omitempty"`
	IsEmailBounced                bool              `xml:"IsEmailBounced,omitempty"`
	Jigsaw                        string            `xml:"Jigsaw,omitempty"`
	JigsawContactId               string            `xml:"JigsawContactId,omitempty"`
	Languagesc                    string            `xml:"Languages__c,omitempty"`
	LastActivityDate              time.Time         `xml:"LastActivityDate,omitempty"`
	LastCURequestDate             time.Time         `xml:"LastCURequestDate,omitempty"`
	LastCUUpdateDate              time.Time         `xml:"LastCUUpdateDate,omitempty"`
	LastModifiedBy                *User             `xml:"LastModifiedBy,omitempty"`
	LastModifiedById              *ID               `xml:"LastModifiedById,omitempty"`
	LastModifiedDate              time.Time         `xml:"LastModifiedDate,omitempty"`
	LastName                      string            `xml:"LastName,omitempty"`
	LastReferencedDate            time.Time         `xml:"LastReferencedDate,omitempty"`
	LastViewedDate                time.Time         `xml:"LastViewedDate,omitempty"`
	LeadSource                    string            `xml:"LeadSource,omitempty"`
	Levelc                        string            `xml:"Level__c,omitempty"`
	ListEmailIndividualRecipients *QueryResult      `xml:"ListEmailIndividualRecipients,omitempty"`
	LookedUpFromActivities        *QueryResult      `xml:"LookedUpFromActivities,omitempty"`
	MailingAddress                *Address          `xml:"MailingAddress,omitempty"`
	MailingCity                   string            `xml:"MailingCity,omitempty"`
	MailingCountry                string            `xml:"MailingCountry,omitempty"`
	MailingGeocodeAccuracy        string            `xml:"MailingGeocodeAccuracy,omitempty"`
	MailingLatitude               float64           `xml:"MailingLatitude,omitempty"`
	MailingLongitude              float64           `xml:"MailingLongitude,omitempty"`
	MailingPostalCode             string            `xml:"MailingPostalCode,omitempty"`
	MailingState                  string            `xml:"MailingState,omitempty"`
	MailingStreet                 string            `xml:"MailingStreet,omitempty"`
	MasterRecord                  *Contact          `xml:"MasterRecord,omitempty"`
	MasterRecordId                *ID               `xml:"MasterRecordId,omitempty"`
	MobilePhone                   string            `xml:"MobilePhone,omitempty"`
	Name                          string            `xml:"Name,omitempty"`
	Notes                         *QueryResult      `xml:"Notes,omitempty"`
	NotesAndAttachments           *QueryResult      `xml:"NotesAndAttachments,omitempty"`
	OpenActivities                *QueryResult      `xml:"OpenActivities,omitempty"`
	Opportunities                 *QueryResult      `xml:"Opportunities,omitempty"`
	OpportunityContactRoles       *QueryResult      `xml:"OpportunityContactRoles,omitempty"`
	OtherAddress                  *Address          `xml:"OtherAddress,omitempty"`
	OtherCity                     string            `xml:"OtherCity,omitempty"`
	OtherCountry                  string            `xml:"OtherCountry,omitempty"`
	OtherGeocodeAccuracy          string            `xml:"OtherGeocodeAccuracy,omitempty"`
	OtherLatitude                 float64           `xml:"OtherLatitude,omitempty"`
	OtherLongitude                float64           `xml:"OtherLongitude,omitempty"`
	OtherPhone                    string            `xml:"OtherPhone,omitempty"`
	OtherPostalCode               string            `xml:"OtherPostalCode,omitempty"`
	OtherState                    string            `xml:"OtherState,omitempty"`
	OtherStreet                   string            `xml:"OtherStreet,omitempty"`
	OutgoingEmailRelations        *QueryResult      `xml:"OutgoingEmailRelations,omitempty"`
	Owner                         *User             `xml:"Owner,omitempty"`
	OwnerId                       *ID               `xml:"OwnerId,omitempty"`
	PersonRecord                  *QueryResult      `xml:"PersonRecord,omitempty"`
	Phone                         string            `xml:"Phone,omitempty"`
	PhotoUrl                      string            `xml:"PhotoUrl,omitempty"`
	ProcessInstances              *QueryResult      `xml:"ProcessInstances,omitempty"`
	ProcessSteps                  *QueryResult      `xml:"ProcessSteps,omitempty"`
	RecordActionHistories         *QueryResult      `xml:"RecordActionHistories,omitempty"`
	RecordActions                 *QueryResult      `xml:"RecordActions,omitempty"`
	RecordAssociatedGroups        *QueryResult      `xml:"RecordAssociatedGroups,omitempty"`
	ReportsTo                     *Contact          `xml:"ReportsTo,omitempty"`
	ReportsToId                   *ID               `xml:"ReportsToId,omitempty"`
	Salutation                    string            `xml:"Salutation,omitempty"`
	Shares                        *QueryResult      `xml:"Shares,omitempty"`
	SystemModstamp                time.Time         `xml:"SystemModstamp,omitempty"`
	Tasks                         *QueryResult      `xml:"Tasks,omitempty"`
	Title                         string            `xml:"Title,omitempty"`
	TopicAssignments              *QueryResult      `xml:"TopicAssignments,omitempty"`
	UndecidedEventRelations       *QueryResult      `xml:"UndecidedEventRelations,omitempty"`
	UserRecordAccess              *UserRecordAccess `xml:"UserRecordAccess,omitempty"`
	Users                         *QueryResult      `xml:"Users,omitempty"`
}

type ContactCleanInfo struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ContactCleanInfo"`

	*SObject

	Address                 *Address          `xml:"Address,omitempty"`
	City                    string            `xml:"City,omitempty"`
	CleanedByJob            bool              `xml:"CleanedByJob,omitempty"`
	CleanedByUser           bool              `xml:"CleanedByUser,omitempty"`
	Contact                 *Contact          `xml:"Contact,omitempty"`
	ContactId               *ID               `xml:"ContactId,omitempty"`
	ContactStatusDataDotCom string            `xml:"ContactStatusDataDotCom,omitempty"`
	Country                 string            `xml:"Country,omitempty"`
	CreatedBy               *User             `xml:"CreatedBy,omitempty"`
	CreatedById             *ID               `xml:"CreatedById,omitempty"`
	CreatedDate             time.Time         `xml:"CreatedDate,omitempty"`
	DataDotComId            string            `xml:"DataDotComId,omitempty"`
	Email                   string            `xml:"Email,omitempty"`
	FirstName               string            `xml:"FirstName,omitempty"`
	GeocodeAccuracy         string            `xml:"GeocodeAccuracy,omitempty"`
	IsDeleted               bool              `xml:"IsDeleted,omitempty"`
	IsDifferentCity         bool              `xml:"IsDifferentCity,omitempty"`
	IsDifferentCountry      bool              `xml:"IsDifferentCountry,omitempty"`
	IsDifferentCountryCode  bool              `xml:"IsDifferentCountryCode,omitempty"`
	IsDifferentEmail        bool              `xml:"IsDifferentEmail,omitempty"`
	IsDifferentFirstName    bool              `xml:"IsDifferentFirstName,omitempty"`
	IsDifferentLastName     bool              `xml:"IsDifferentLastName,omitempty"`
	IsDifferentPhone        bool              `xml:"IsDifferentPhone,omitempty"`
	IsDifferentPostalCode   bool              `xml:"IsDifferentPostalCode,omitempty"`
	IsDifferentState        bool              `xml:"IsDifferentState,omitempty"`
	IsDifferentStateCode    bool              `xml:"IsDifferentStateCode,omitempty"`
	IsDifferentStreet       bool              `xml:"IsDifferentStreet,omitempty"`
	IsDifferentTitle        bool              `xml:"IsDifferentTitle,omitempty"`
	IsFlaggedWrongAddress   bool              `xml:"IsFlaggedWrongAddress,omitempty"`
	IsFlaggedWrongEmail     bool              `xml:"IsFlaggedWrongEmail,omitempty"`
	IsFlaggedWrongName      bool              `xml:"IsFlaggedWrongName,omitempty"`
	IsFlaggedWrongPhone     bool              `xml:"IsFlaggedWrongPhone,omitempty"`
	IsFlaggedWrongTitle     bool              `xml:"IsFlaggedWrongTitle,omitempty"`
	IsInactive              bool              `xml:"IsInactive,omitempty"`
	IsReviewedAddress       bool              `xml:"IsReviewedAddress,omitempty"`
	IsReviewedEmail         bool              `xml:"IsReviewedEmail,omitempty"`
	IsReviewedName          bool              `xml:"IsReviewedName,omitempty"`
	IsReviewedPhone         bool              `xml:"IsReviewedPhone,omitempty"`
	IsReviewedTitle         bool              `xml:"IsReviewedTitle,omitempty"`
	LastMatchedDate         time.Time         `xml:"LastMatchedDate,omitempty"`
	LastModifiedBy          *User             `xml:"LastModifiedBy,omitempty"`
	LastModifiedById        *ID               `xml:"LastModifiedById,omitempty"`
	LastModifiedDate        time.Time         `xml:"LastModifiedDate,omitempty"`
	LastName                string            `xml:"LastName,omitempty"`
	LastStatusChangedBy     *User             `xml:"LastStatusChangedBy,omitempty"`
	LastStatusChangedById   *ID               `xml:"LastStatusChangedById,omitempty"`
	LastStatusChangedDate   time.Time         `xml:"LastStatusChangedDate,omitempty"`
	Latitude                float64           `xml:"Latitude,omitempty"`
	Longitude               float64           `xml:"Longitude,omitempty"`
	Name                    string            `xml:"Name,omitempty"`
	Phone                   string            `xml:"Phone,omitempty"`
	PostalCode              string            `xml:"PostalCode,omitempty"`
	State                   string            `xml:"State,omitempty"`
	Street                  string            `xml:"Street,omitempty"`
	SystemModstamp          time.Time         `xml:"SystemModstamp,omitempty"`
	Title                   string            `xml:"Title,omitempty"`
	UserRecordAccess        *UserRecordAccess `xml:"UserRecordAccess,omitempty"`
}

type ContactFeed struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ContactFeed"`

	*SObject

	BestComment        *FeedComment `xml:"BestComment,omitempty"`
	BestCommentId      *ID          `xml:"BestCommentId,omitempty"`
	Body               string       `xml:"Body,omitempty"`
	CommentCount       int32        `xml:"CommentCount,omitempty"`
	CreatedBy          *SObject     `xml:"CreatedBy,omitempty"`
	CreatedById        *ID          `xml:"CreatedById,omitempty"`
	CreatedDate        time.Time    `xml:"CreatedDate,omitempty"`
	FeedAttachments    *QueryResult `xml:"FeedAttachments,omitempty"`
	FeedComments       *QueryResult `xml:"FeedComments,omitempty"`
	FeedLikes          *QueryResult `xml:"FeedLikes,omitempty"`
	FeedSignals        *QueryResult `xml:"FeedSignals,omitempty"`
	FeedTrackedChanges *QueryResult `xml:"FeedTrackedChanges,omitempty"`
	InsertedBy         *SObject     `xml:"InsertedBy,omitempty"`
	InsertedById       *ID          `xml:"InsertedById,omitempty"`
	IsDeleted          bool         `xml:"IsDeleted,omitempty"`
	IsRichText         bool         `xml:"IsRichText,omitempty"`
	LastModifiedDate   time.Time    `xml:"LastModifiedDate,omitempty"`
	LikeCount          int32        `xml:"LikeCount,omitempty"`
	LinkUrl            string       `xml:"LinkUrl,omitempty"`
	Parent             *Contact     `xml:"Parent,omitempty"`
	ParentId           *ID          `xml:"ParentId,omitempty"`
	RelatedRecordId    *ID          `xml:"RelatedRecordId,omitempty"`
	SystemModstamp     time.Time    `xml:"SystemModstamp,omitempty"`
	Title              string       `xml:"Title,omitempty"`
	Type               string       `xml:"Type,omitempty"`
}

type ContactHistory struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ContactHistory"`

	*SObject

	Contact     *Contact    `xml:"Contact,omitempty"`
	ContactId   *ID         `xml:"ContactId,omitempty"`
	CreatedBy   *SObject    `xml:"CreatedBy,omitempty"`
	CreatedById *ID         `xml:"CreatedById,omitempty"`
	CreatedDate time.Time   `xml:"CreatedDate,omitempty"`
	Field       string      `xml:"Field,omitempty"`
	IsDeleted   bool        `xml:"IsDeleted,omitempty"`
	NewValue    interface{} `xml:"NewValue,omitempty"`
	OldValue    interface{} `xml:"OldValue,omitempty"`
}

type ContactShare struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ContactShare"`

	*SObject

	Contact            *Contact  `xml:"Contact,omitempty"`
	ContactAccessLevel string    `xml:"ContactAccessLevel,omitempty"`
	ContactId          *ID       `xml:"ContactId,omitempty"`
	IsDeleted          bool      `xml:"IsDeleted,omitempty"`
	LastModifiedBy     *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById   *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate   time.Time `xml:"LastModifiedDate,omitempty"`
	RowCause           string    `xml:"RowCause,omitempty"`
	UserOrGroup        *SObject  `xml:"UserOrGroup,omitempty"`
	UserOrGroupId      *ID       `xml:"UserOrGroupId,omitempty"`
}

type ContentAsset struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ContentAsset"`

	*SObject

	ContentDocument          *ContentDocument `xml:"ContentDocument,omitempty"`
	ContentDocumentId        *ID              `xml:"ContentDocumentId,omitempty"`
	CreatedBy                *User            `xml:"CreatedBy,omitempty"`
	CreatedById              *ID              `xml:"CreatedById,omitempty"`
	CreatedDate              time.Time        `xml:"CreatedDate,omitempty"`
	DeveloperName            string           `xml:"DeveloperName,omitempty"`
	IsDeleted                bool             `xml:"IsDeleted,omitempty"`
	IsVisibleByExternalUsers bool             `xml:"IsVisibleByExternalUsers,omitempty"`
	Language                 string           `xml:"Language,omitempty"`
	LastModifiedBy           *User            `xml:"LastModifiedBy,omitempty"`
	LastModifiedById         *ID              `xml:"LastModifiedById,omitempty"`
	LastModifiedDate         time.Time        `xml:"LastModifiedDate,omitempty"`
	MasterLabel              string           `xml:"MasterLabel,omitempty"`
	NamespacePrefix          string           `xml:"NamespacePrefix,omitempty"`
	SystemModstamp           time.Time        `xml:"SystemModstamp,omitempty"`
}

type ContentBody struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ContentBody"`

	*SObject
}

type ContentDistribution struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ContentDistribution"`

	*SObject

	ContentDistributionViews         *QueryResult    `xml:"ContentDistributionViews,omitempty"`
	ContentDocumentId                *ID             `xml:"ContentDocumentId,omitempty"`
	ContentDownloadUrl               string          `xml:"ContentDownloadUrl,omitempty"`
	ContentVersion                   *ContentVersion `xml:"ContentVersion,omitempty"`
	ContentVersionId                 *ID             `xml:"ContentVersionId,omitempty"`
	CreatedBy                        *User           `xml:"CreatedBy,omitempty"`
	CreatedById                      *ID             `xml:"CreatedById,omitempty"`
	CreatedDate                      time.Time       `xml:"CreatedDate,omitempty"`
	DistributionPublicUrl            string          `xml:"DistributionPublicUrl,omitempty"`
	ExpiryDate                       time.Time       `xml:"ExpiryDate,omitempty"`
	FirstViewDate                    time.Time       `xml:"FirstViewDate,omitempty"`
	IsDeleted                        bool            `xml:"IsDeleted,omitempty"`
	LastModifiedBy                   *User           `xml:"LastModifiedBy,omitempty"`
	LastModifiedById                 *ID             `xml:"LastModifiedById,omitempty"`
	LastModifiedDate                 time.Time       `xml:"LastModifiedDate,omitempty"`
	LastViewDate                     time.Time       `xml:"LastViewDate,omitempty"`
	Name                             string          `xml:"Name,omitempty"`
	Owner                            *User           `xml:"Owner,omitempty"`
	OwnerId                          *ID             `xml:"OwnerId,omitempty"`
	Password                         string          `xml:"Password,omitempty"`
	PdfDownloadUrl                   string          `xml:"PdfDownloadUrl,omitempty"`
	PreferencesAllowOriginalDownload bool            `xml:"PreferencesAllowOriginalDownload,omitempty"`
	PreferencesAllowPDFDownload      bool            `xml:"PreferencesAllowPDFDownload,omitempty"`
	PreferencesAllowViewInBrowser    bool            `xml:"PreferencesAllowViewInBrowser,omitempty"`
	PreferencesExpires               bool            `xml:"PreferencesExpires,omitempty"`
	PreferencesLinkLatestVersion     bool            `xml:"PreferencesLinkLatestVersion,omitempty"`
	PreferencesNotifyOnVisit         bool            `xml:"PreferencesNotifyOnVisit,omitempty"`
	PreferencesNotifyRndtnComplete   bool            `xml:"PreferencesNotifyRndtnComplete,omitempty"`
	PreferencesPasswordRequired      bool            `xml:"PreferencesPasswordRequired,omitempty"`
	RelatedRecord                    *SObject        `xml:"RelatedRecord,omitempty"`
	RelatedRecordId                  *ID             `xml:"RelatedRecordId,omitempty"`
	SystemModstamp                   time.Time       `xml:"SystemModstamp,omitempty"`
	ViewCount                        int32           `xml:"ViewCount,omitempty"`
}

type ContentDistributionView struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ContentDistributionView"`

	*SObject

	CreatedBy      *User                `xml:"CreatedBy,omitempty"`
	CreatedById    *ID                  `xml:"CreatedById,omitempty"`
	CreatedDate    time.Time            `xml:"CreatedDate,omitempty"`
	Distribution   *ContentDistribution `xml:"Distribution,omitempty"`
	DistributionId *ID                  `xml:"DistributionId,omitempty"`
	IsDeleted      bool                 `xml:"IsDeleted,omitempty"`
	IsDownload     bool                 `xml:"IsDownload,omitempty"`
	IsInternal     bool                 `xml:"IsInternal,omitempty"`
	ParentViewId   *ID                  `xml:"ParentViewId,omitempty"`
	SystemModstamp time.Time            `xml:"SystemModstamp,omitempty"`
}

type ContentDocument struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ContentDocument"`

	*SObject

	ArchivedById               *ID             `xml:"ArchivedById,omitempty"`
	ArchivedDate               time.Time       `xml:"ArchivedDate,omitempty"`
	ContentAsset               *ContentAsset   `xml:"ContentAsset,omitempty"`
	ContentAssetId             *ID             `xml:"ContentAssetId,omitempty"`
	ContentDistributions       *QueryResult    `xml:"ContentDistributions,omitempty"`
	ContentDocumentLinks       *QueryResult    `xml:"ContentDocumentLinks,omitempty"`
	ContentModifiedDate        time.Time       `xml:"ContentModifiedDate,omitempty"`
	ContentSize                int32           `xml:"ContentSize,omitempty"`
	ContentVersions            *QueryResult    `xml:"ContentVersions,omitempty"`
	CreatedBy                  *User           `xml:"CreatedBy,omitempty"`
	CreatedById                *ID             `xml:"CreatedById,omitempty"`
	CreatedDate                time.Time       `xml:"CreatedDate,omitempty"`
	Description                string          `xml:"Description,omitempty"`
	FeedSubscriptionsForEntity *QueryResult    `xml:"FeedSubscriptionsForEntity,omitempty"`
	Feeds                      *QueryResult    `xml:"Feeds,omitempty"`
	FileExtension              string          `xml:"FileExtension,omitempty"`
	FileType                   string          `xml:"FileType,omitempty"`
	Histories                  *QueryResult    `xml:"Histories,omitempty"`
	IsArchived                 bool            `xml:"IsArchived,omitempty"`
	IsDeleted                  bool            `xml:"IsDeleted,omitempty"`
	LastModifiedBy             *User           `xml:"LastModifiedBy,omitempty"`
	LastModifiedById           *ID             `xml:"LastModifiedById,omitempty"`
	LastModifiedDate           time.Time       `xml:"LastModifiedDate,omitempty"`
	LastReferencedDate         time.Time       `xml:"LastReferencedDate,omitempty"`
	LastViewedDate             time.Time       `xml:"LastViewedDate,omitempty"`
	LatestPublishedVersion     *ContentVersion `xml:"LatestPublishedVersion,omitempty"`
	LatestPublishedVersionId   *ID             `xml:"LatestPublishedVersionId,omitempty"`
	Owner                      *User           `xml:"Owner,omitempty"`
	OwnerId                    *ID             `xml:"OwnerId,omitempty"`
	ParentId                   *ID             `xml:"ParentId,omitempty"`
	PublishStatus              string          `xml:"PublishStatus,omitempty"`
	SharingOption              string          `xml:"SharingOption,omitempty"`
	SharingPrivacy             string          `xml:"SharingPrivacy,omitempty"`
	SystemModstamp             time.Time       `xml:"SystemModstamp,omitempty"`
	Title                      string          `xml:"Title,omitempty"`
	TopicAssignments           *QueryResult    `xml:"TopicAssignments,omitempty"`
}

type ContentDocumentFeed struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ContentDocumentFeed"`

	*SObject

	BestComment        *FeedComment     `xml:"BestComment,omitempty"`
	BestCommentId      *ID              `xml:"BestCommentId,omitempty"`
	Body               string           `xml:"Body,omitempty"`
	CommentCount       int32            `xml:"CommentCount,omitempty"`
	CreatedBy          *SObject         `xml:"CreatedBy,omitempty"`
	CreatedById        *ID              `xml:"CreatedById,omitempty"`
	CreatedDate        time.Time        `xml:"CreatedDate,omitempty"`
	FeedAttachments    *QueryResult     `xml:"FeedAttachments,omitempty"`
	FeedComments       *QueryResult     `xml:"FeedComments,omitempty"`
	FeedLikes          *QueryResult     `xml:"FeedLikes,omitempty"`
	FeedSignals        *QueryResult     `xml:"FeedSignals,omitempty"`
	FeedTrackedChanges *QueryResult     `xml:"FeedTrackedChanges,omitempty"`
	InsertedBy         *SObject         `xml:"InsertedBy,omitempty"`
	InsertedById       *ID              `xml:"InsertedById,omitempty"`
	IsDeleted          bool             `xml:"IsDeleted,omitempty"`
	IsRichText         bool             `xml:"IsRichText,omitempty"`
	LastModifiedDate   time.Time        `xml:"LastModifiedDate,omitempty"`
	LikeCount          int32            `xml:"LikeCount,omitempty"`
	LinkUrl            string           `xml:"LinkUrl,omitempty"`
	Parent             *ContentDocument `xml:"Parent,omitempty"`
	ParentId           *ID              `xml:"ParentId,omitempty"`
	RelatedRecordId    *ID              `xml:"RelatedRecordId,omitempty"`
	SystemModstamp     time.Time        `xml:"SystemModstamp,omitempty"`
	Title              string           `xml:"Title,omitempty"`
	Type               string           `xml:"Type,omitempty"`
}

type ContentDocumentHistory struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ContentDocumentHistory"`

	*SObject

	ContentDocument   *ContentDocument `xml:"ContentDocument,omitempty"`
	ContentDocumentId *ID              `xml:"ContentDocumentId,omitempty"`
	CreatedBy         *SObject         `xml:"CreatedBy,omitempty"`
	CreatedById       *ID              `xml:"CreatedById,omitempty"`
	CreatedDate       time.Time        `xml:"CreatedDate,omitempty"`
	Field             string           `xml:"Field,omitempty"`
	IsDeleted         bool             `xml:"IsDeleted,omitempty"`
	NewValue          interface{}      `xml:"NewValue,omitempty"`
	OldValue          interface{}      `xml:"OldValue,omitempty"`
}

type ContentDocumentLink struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ContentDocumentLink"`

	*SObject

	ContentDocument   *ContentDocument `xml:"ContentDocument,omitempty"`
	ContentDocumentId *ID              `xml:"ContentDocumentId,omitempty"`
	IsDeleted         bool             `xml:"IsDeleted,omitempty"`
	LinkedEntity      *SObject         `xml:"LinkedEntity,omitempty"`
	LinkedEntityId    *ID              `xml:"LinkedEntityId,omitempty"`
	ShareType         string           `xml:"ShareType,omitempty"`
	SystemModstamp    time.Time        `xml:"SystemModstamp,omitempty"`
	Visibility        string           `xml:"Visibility,omitempty"`
}

type ContentDocumentSubscription struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ContentDocumentSubscription"`

	*SObject

	ContentDocument   *ContentDocument `xml:"ContentDocument,omitempty"`
	ContentDocumentId *ID              `xml:"ContentDocumentId,omitempty"`
	IsCommentSub      bool             `xml:"IsCommentSub,omitempty"`
	IsDocumentSub     bool             `xml:"IsDocumentSub,omitempty"`
	User              *User            `xml:"User,omitempty"`
	UserId            *ID              `xml:"UserId,omitempty"`
}

type ContentFolder struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ContentFolder"`

	*SObject

	ContentFolderLinks    *QueryResult      `xml:"ContentFolderLinks,omitempty"`
	CreatedBy             *User             `xml:"CreatedBy,omitempty"`
	CreatedById           *ID               `xml:"CreatedById,omitempty"`
	CreatedDate           time.Time         `xml:"CreatedDate,omitempty"`
	IsDeleted             bool              `xml:"IsDeleted,omitempty"`
	LastModifiedBy        *User             `xml:"LastModifiedBy,omitempty"`
	LastModifiedById      *ID               `xml:"LastModifiedById,omitempty"`
	LastModifiedDate      time.Time         `xml:"LastModifiedDate,omitempty"`
	Name                  string            `xml:"Name,omitempty"`
	ParentContentFolder   *ContentFolder    `xml:"ParentContentFolder,omitempty"`
	ParentContentFolderId *ID               `xml:"ParentContentFolderId,omitempty"`
	SystemModstamp        time.Time         `xml:"SystemModstamp,omitempty"`
	UserRecordAccess      *UserRecordAccess `xml:"UserRecordAccess,omitempty"`
}

type ContentFolderItem struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ContentFolderItem"`

	*SObject

	ContentSize           int32          `xml:"ContentSize,omitempty"`
	CreatedBy             *User          `xml:"CreatedBy,omitempty"`
	CreatedById           *ID            `xml:"CreatedById,omitempty"`
	CreatedDate           time.Time      `xml:"CreatedDate,omitempty"`
	FileExtension         string         `xml:"FileExtension,omitempty"`
	FileType              string         `xml:"FileType,omitempty"`
	IsDeleted             bool           `xml:"IsDeleted,omitempty"`
	IsFolder              bool           `xml:"IsFolder,omitempty"`
	LastModifiedBy        *User          `xml:"LastModifiedBy,omitempty"`
	LastModifiedById      *ID            `xml:"LastModifiedById,omitempty"`
	LastModifiedDate      time.Time      `xml:"LastModifiedDate,omitempty"`
	ParentContentFolder   *ContentFolder `xml:"ParentContentFolder,omitempty"`
	ParentContentFolderId *ID            `xml:"ParentContentFolderId,omitempty"`
	SystemModstamp        time.Time      `xml:"SystemModstamp,omitempty"`
	Title                 string         `xml:"Title,omitempty"`
}

type ContentFolderLink struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ContentFolderLink"`

	*SObject

	ContentFolder      *ContentFolder `xml:"ContentFolder,omitempty"`
	ContentFolderId    *ID            `xml:"ContentFolderId,omitempty"`
	EnableFolderStatus string         `xml:"EnableFolderStatus,omitempty"`
	IsDeleted          bool           `xml:"IsDeleted,omitempty"`
	ParentEntityId     *ID            `xml:"ParentEntityId,omitempty"`
}

type ContentFolderMember struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ContentFolderMember"`

	*SObject

	ChildRecord           *ContentDocument `xml:"ChildRecord,omitempty"`
	ChildRecordId         *ID              `xml:"ChildRecordId,omitempty"`
	CreatedBy             *User            `xml:"CreatedBy,omitempty"`
	CreatedById           *ID              `xml:"CreatedById,omitempty"`
	CreatedDate           time.Time        `xml:"CreatedDate,omitempty"`
	IsDeleted             bool             `xml:"IsDeleted,omitempty"`
	LastModifiedBy        *User            `xml:"LastModifiedBy,omitempty"`
	LastModifiedById      *ID              `xml:"LastModifiedById,omitempty"`
	LastModifiedDate      time.Time        `xml:"LastModifiedDate,omitempty"`
	ParentContentFolder   *ContentFolder   `xml:"ParentContentFolder,omitempty"`
	ParentContentFolderId *ID              `xml:"ParentContentFolderId,omitempty"`
	SystemModstamp        time.Time        `xml:"SystemModstamp,omitempty"`
}

type ContentNotification struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ContentNotification"`

	*SObject

	CreatedDate        time.Time `xml:"CreatedDate,omitempty"`
	EntityIdentifierId *ID       `xml:"EntityIdentifierId,omitempty"`
	EntityType         string    `xml:"EntityType,omitempty"`
	Nature             string    `xml:"Nature,omitempty"`
	Subject            string    `xml:"Subject,omitempty"`
	Text               string    `xml:"Text,omitempty"`
	Users              *User     `xml:"Users,omitempty"`
	UsersId            *ID       `xml:"UsersId,omitempty"`
}

type ContentTagSubscription struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ContentTagSubscription"`

	*SObject

	User   *User `xml:"User,omitempty"`
	UserId *ID   `xml:"UserId,omitempty"`
}

type ContentUserSubscription struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ContentUserSubscription"`

	*SObject

	SubscribedToUser   *User `xml:"SubscribedToUser,omitempty"`
	SubscribedToUserId *ID   `xml:"SubscribedToUserId,omitempty"`
	SubscriberUser     *User `xml:"SubscriberUser,omitempty"`
	SubscriberUserId   *ID   `xml:"SubscriberUserId,omitempty"`
}

type ContentVersion struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ContentVersion"`

	*SObject

	Checksum               string              `xml:"Checksum,omitempty"`
	ContentBody            *ContentBody        `xml:"ContentBody,omitempty"`
	ContentBodyId          *ID                 `xml:"ContentBodyId,omitempty"`
	ContentDocument        *ContentDocument    `xml:"ContentDocument,omitempty"`
	ContentDocumentId      *ID                 `xml:"ContentDocumentId,omitempty"`
	ContentLocation        string              `xml:"ContentLocation,omitempty"`
	ContentModifiedBy      *User               `xml:"ContentModifiedBy,omitempty"`
	ContentModifiedById    *ID                 `xml:"ContentModifiedById,omitempty"`
	ContentModifiedDate    time.Time           `xml:"ContentModifiedDate,omitempty"`
	ContentSize            int32               `xml:"ContentSize,omitempty"`
	ContentUrl             string              `xml:"ContentUrl,omitempty"`
	CreatedBy              *User               `xml:"CreatedBy,omitempty"`
	CreatedById            *ID                 `xml:"CreatedById,omitempty"`
	CreatedDate            time.Time           `xml:"CreatedDate,omitempty"`
	Description            string              `xml:"Description,omitempty"`
	ExternalDataSource     *ExternalDataSource `xml:"ExternalDataSource,omitempty"`
	ExternalDataSourceId   *ID                 `xml:"ExternalDataSourceId,omitempty"`
	ExternalDocumentInfo1  string              `xml:"ExternalDocumentInfo1,omitempty"`
	ExternalDocumentInfo2  string              `xml:"ExternalDocumentInfo2,omitempty"`
	FeaturedContentBoost   int32               `xml:"FeaturedContentBoost,omitempty"`
	FeaturedContentDate    time.Time           `xml:"FeaturedContentDate,omitempty"`
	FileExtension          string              `xml:"FileExtension,omitempty"`
	FileType               string              `xml:"FileType,omitempty"`
	FirstPublishLocation   *SObject            `xml:"FirstPublishLocation,omitempty"`
	FirstPublishLocationId *ID                 `xml:"FirstPublishLocationId,omitempty"`
	Histories              *QueryResult        `xml:"Histories,omitempty"`
	IsAssetEnabled         bool                `xml:"IsAssetEnabled,omitempty"`
	IsDeleted              bool                `xml:"IsDeleted,omitempty"`
	IsLatest               bool                `xml:"IsLatest,omitempty"`
	IsMajorVersion         bool                `xml:"IsMajorVersion,omitempty"`
	LastModifiedBy         *User               `xml:"LastModifiedBy,omitempty"`
	LastModifiedById       *ID                 `xml:"LastModifiedById,omitempty"`
	LastModifiedDate       time.Time           `xml:"LastModifiedDate,omitempty"`
	NegativeRatingCount    int32               `xml:"NegativeRatingCount,omitempty"`
	Origin                 string              `xml:"Origin,omitempty"`
	Owner                  *User               `xml:"Owner,omitempty"`
	OwnerId                *ID                 `xml:"OwnerId,omitempty"`
	PathOnClient           string              `xml:"PathOnClient,omitempty"`
	PositiveRatingCount    int32               `xml:"PositiveRatingCount,omitempty"`
	PublishStatus          string              `xml:"PublishStatus,omitempty"`
	RatingCount            int32               `xml:"RatingCount,omitempty"`
	ReasonForChange        string              `xml:"ReasonForChange,omitempty"`
	SharingOption          string              `xml:"SharingOption,omitempty"`
	SharingPrivacy         string              `xml:"SharingPrivacy,omitempty"`
	SystemModstamp         time.Time           `xml:"SystemModstamp,omitempty"`
	TagCsv                 string              `xml:"TagCsv,omitempty"`
	TextPreview            string              `xml:"TextPreview,omitempty"`
	Title                  string              `xml:"Title,omitempty"`
	VersionData            []byte              `xml:"VersionData,omitempty"`
	VersionNumber          string              `xml:"VersionNumber,omitempty"`
}

type ContentVersionComment struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ContentVersionComment"`

	*SObject

	ContentDocument   *ContentDocument `xml:"ContentDocument,omitempty"`
	ContentDocumentId *ID              `xml:"ContentDocumentId,omitempty"`
	ContentVersion    *ContentVersion  `xml:"ContentVersion,omitempty"`
	ContentVersionId  *ID              `xml:"ContentVersionId,omitempty"`
	CreatedDate       time.Time        `xml:"CreatedDate,omitempty"`
	UserComment       string           `xml:"UserComment,omitempty"`
}

type ContentVersionHistory struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ContentVersionHistory"`

	*SObject

	ContentVersion   *ContentVersion `xml:"ContentVersion,omitempty"`
	ContentVersionId *ID             `xml:"ContentVersionId,omitempty"`
	CreatedBy        *SObject        `xml:"CreatedBy,omitempty"`
	CreatedById      *ID             `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time       `xml:"CreatedDate,omitempty"`
	Field            string          `xml:"Field,omitempty"`
	IsDeleted        bool            `xml:"IsDeleted,omitempty"`
	NewValue         interface{}     `xml:"NewValue,omitempty"`
	OldValue         interface{}     `xml:"OldValue,omitempty"`
}

type ContentVersionRating struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ContentVersionRating"`

	*SObject

	ContentVersion   *ContentVersion `xml:"ContentVersion,omitempty"`
	ContentVersionId *ID             `xml:"ContentVersionId,omitempty"`
	LastModifiedDate time.Time       `xml:"LastModifiedDate,omitempty"`
	Rating           int32           `xml:"Rating,omitempty"`
	User             *User           `xml:"User,omitempty"`
	UserComment      string          `xml:"UserComment,omitempty"`
	UserId           *ID             `xml:"UserId,omitempty"`
}

type ContentWorkspace struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ContentWorkspace"`

	*SObject

	AttachedContentDocuments     *QueryResult  `xml:"AttachedContentDocuments,omitempty"`
	CombinedAttachments          *QueryResult  `xml:"CombinedAttachments,omitempty"`
	ContentDocumentLinks         *QueryResult  `xml:"ContentDocumentLinks,omitempty"`
	ContentFolderLinks           *QueryResult  `xml:"ContentFolderLinks,omitempty"`
	ContentWorkspaceMembers      *QueryResult  `xml:"ContentWorkspaceMembers,omitempty"`
	CreatedBy                    *User         `xml:"CreatedBy,omitempty"`
	CreatedById                  *ID           `xml:"CreatedById,omitempty"`
	CreatedDate                  time.Time     `xml:"CreatedDate,omitempty"`
	DefaultRecordTypeId          *ID           `xml:"DefaultRecordTypeId,omitempty"`
	Description                  string        `xml:"Description,omitempty"`
	DeveloperName                string        `xml:"DeveloperName,omitempty"`
	IsRestrictContentTypes       bool          `xml:"IsRestrictContentTypes,omitempty"`
	IsRestrictLinkedContentTypes bool          `xml:"IsRestrictLinkedContentTypes,omitempty"`
	LastModifiedBy               *User         `xml:"LastModifiedBy,omitempty"`
	LastModifiedById             *ID           `xml:"LastModifiedById,omitempty"`
	LastModifiedDate             time.Time     `xml:"LastModifiedDate,omitempty"`
	LastWorkspaceActivityDate    time.Time     `xml:"LastWorkspaceActivityDate,omitempty"`
	Name                         string        `xml:"Name,omitempty"`
	NamespacePrefix              string        `xml:"NamespacePrefix,omitempty"`
	RootContentFolderId          *ID           `xml:"RootContentFolderId,omitempty"`
	ShouldAddCreatorMembership   bool          `xml:"ShouldAddCreatorMembership,omitempty"`
	SystemModstamp               time.Time     `xml:"SystemModstamp,omitempty"`
	TagModel                     string        `xml:"TagModel,omitempty"`
	WorkspaceImage               *ContentAsset `xml:"WorkspaceImage,omitempty"`
	WorkspaceImageId             *ID           `xml:"WorkspaceImageId,omitempty"`
	WorkspaceType                string        `xml:"WorkspaceType,omitempty"`
}

type ContentWorkspaceDoc struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ContentWorkspaceDoc"`

	*SObject

	ContentDocument    *ContentDocument  `xml:"ContentDocument,omitempty"`
	ContentDocumentId  *ID               `xml:"ContentDocumentId,omitempty"`
	ContentWorkspace   *ContentWorkspace `xml:"ContentWorkspace,omitempty"`
	ContentWorkspaceId *ID               `xml:"ContentWorkspaceId,omitempty"`
	CreatedDate        time.Time         `xml:"CreatedDate,omitempty"`
	IsDeleted          bool              `xml:"IsDeleted,omitempty"`
	IsOwner            bool              `xml:"IsOwner,omitempty"`
	SystemModstamp     time.Time         `xml:"SystemModstamp,omitempty"`
}

type ContentWorkspaceMember struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ContentWorkspaceMember"`

	*SObject

	ContentWorkspace             *ContentWorkspace           `xml:"ContentWorkspace,omitempty"`
	ContentWorkspaceId           *ID                         `xml:"ContentWorkspaceId,omitempty"`
	ContentWorkspacePermission   *ContentWorkspacePermission `xml:"ContentWorkspacePermission,omitempty"`
	ContentWorkspacePermissionId *ID                         `xml:"ContentWorkspacePermissionId,omitempty"`
	CreatedBy                    *User                       `xml:"CreatedBy,omitempty"`
	CreatedById                  *ID                         `xml:"CreatedById,omitempty"`
	CreatedDate                  time.Time                   `xml:"CreatedDate,omitempty"`
	Member                       *SObject                    `xml:"Member,omitempty"`
	MemberId                     *ID                         `xml:"MemberId,omitempty"`
	MemberType                   string                      `xml:"MemberType,omitempty"`
}

type ContentWorkspacePermission struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ContentWorkspacePermission"`

	*SObject

	CreatedBy                        *User     `xml:"CreatedBy,omitempty"`
	CreatedById                      *ID       `xml:"CreatedById,omitempty"`
	CreatedDate                      time.Time `xml:"CreatedDate,omitempty"`
	Description                      string    `xml:"Description,omitempty"`
	LastModifiedBy                   *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById                 *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate                 time.Time `xml:"LastModifiedDate,omitempty"`
	Name                             string    `xml:"Name,omitempty"`
	PermissionsAddComment            bool      `xml:"PermissionsAddComment,omitempty"`
	PermissionsAddContent            bool      `xml:"PermissionsAddContent,omitempty"`
	PermissionsAddContentOBO         bool      `xml:"PermissionsAddContentOBO,omitempty"`
	PermissionsArchiveContent        bool      `xml:"PermissionsArchiveContent,omitempty"`
	PermissionsChatterSharing        bool      `xml:"PermissionsChatterSharing,omitempty"`
	PermissionsDeleteContent         bool      `xml:"PermissionsDeleteContent,omitempty"`
	PermissionsDeliverContent        bool      `xml:"PermissionsDeliverContent,omitempty"`
	PermissionsFeatureContent        bool      `xml:"PermissionsFeatureContent,omitempty"`
	PermissionsManageWorkspace       bool      `xml:"PermissionsManageWorkspace,omitempty"`
	PermissionsModifyComments        bool      `xml:"PermissionsModifyComments,omitempty"`
	PermissionsOrganizeFileAndFolder bool      `xml:"PermissionsOrganizeFileAndFolder,omitempty"`
	PermissionsTagContent            bool      `xml:"PermissionsTagContent,omitempty"`
	PermissionsViewComments          bool      `xml:"PermissionsViewComments,omitempty"`
	SystemModstamp                   time.Time `xml:"SystemModstamp,omitempty"`
	Type                             string    `xml:"Type,omitempty"`
}

type ContentWorkspaceSubscription struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ContentWorkspaceSubscription"`

	*SObject

	ContentWorkspace   *ContentWorkspace `xml:"ContentWorkspace,omitempty"`
	ContentWorkspaceId *ID               `xml:"ContentWorkspaceId,omitempty"`
	User               *User             `xml:"User,omitempty"`
	UserId             *ID               `xml:"UserId,omitempty"`
}

type Contract struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com Contract"`

	*SObject

	Account                    *Account          `xml:"Account,omitempty"`
	AccountId                  *ID               `xml:"AccountId,omitempty"`
	ActivatedBy                *User             `xml:"ActivatedBy,omitempty"`
	ActivatedById              *ID               `xml:"ActivatedById,omitempty"`
	ActivatedDate              time.Time         `xml:"ActivatedDate,omitempty"`
	ActivityHistories          *QueryResult      `xml:"ActivityHistories,omitempty"`
	AttachedContentDocuments   *QueryResult      `xml:"AttachedContentDocuments,omitempty"`
	Attachments                *QueryResult      `xml:"Attachments,omitempty"`
	BillingAddress             *Address          `xml:"BillingAddress,omitempty"`
	BillingCity                string            `xml:"BillingCity,omitempty"`
	BillingCountry             string            `xml:"BillingCountry,omitempty"`
	BillingGeocodeAccuracy     string            `xml:"BillingGeocodeAccuracy,omitempty"`
	BillingLatitude            float64           `xml:"BillingLatitude,omitempty"`
	BillingLongitude           float64           `xml:"BillingLongitude,omitempty"`
	BillingPostalCode          string            `xml:"BillingPostalCode,omitempty"`
	BillingState               string            `xml:"BillingState,omitempty"`
	BillingStreet              string            `xml:"BillingStreet,omitempty"`
	CombinedAttachments        *QueryResult      `xml:"CombinedAttachments,omitempty"`
	CompanySigned              *User             `xml:"CompanySigned,omitempty"`
	CompanySignedDate          time.Time         `xml:"CompanySignedDate,omitempty"`
	CompanySignedId            *ID               `xml:"CompanySignedId,omitempty"`
	ContentDocumentLinks       *QueryResult      `xml:"ContentDocumentLinks,omitempty"`
	ContractContactRoles       *QueryResult      `xml:"ContractContactRoles,omitempty"`
	ContractNumber             string            `xml:"ContractNumber,omitempty"`
	ContractTerm               int32             `xml:"ContractTerm,omitempty"`
	CreatedBy                  *User             `xml:"CreatedBy,omitempty"`
	CreatedById                *ID               `xml:"CreatedById,omitempty"`
	CreatedDate                time.Time         `xml:"CreatedDate,omitempty"`
	CustomerSigned             *Contact          `xml:"CustomerSigned,omitempty"`
	CustomerSignedDate         time.Time         `xml:"CustomerSignedDate,omitempty"`
	CustomerSignedId           *ID               `xml:"CustomerSignedId,omitempty"`
	CustomerSignedTitle        string            `xml:"CustomerSignedTitle,omitempty"`
	Description                string            `xml:"Description,omitempty"`
	Emails                     *QueryResult      `xml:"Emails,omitempty"`
	EndDate                    time.Time         `xml:"EndDate,omitempty"`
	Events                     *QueryResult      `xml:"Events,omitempty"`
	FeedSubscriptionsForEntity *QueryResult      `xml:"FeedSubscriptionsForEntity,omitempty"`
	Feeds                      *QueryResult      `xml:"Feeds,omitempty"`
	Histories                  *QueryResult      `xml:"Histories,omitempty"`
	IsDeleted                  bool              `xml:"IsDeleted,omitempty"`
	LastActivityDate           time.Time         `xml:"LastActivityDate,omitempty"`
	LastApprovedDate           time.Time         `xml:"LastApprovedDate,omitempty"`
	LastModifiedBy             *User             `xml:"LastModifiedBy,omitempty"`
	LastModifiedById           *ID               `xml:"LastModifiedById,omitempty"`
	LastModifiedDate           time.Time         `xml:"LastModifiedDate,omitempty"`
	LastReferencedDate         time.Time         `xml:"LastReferencedDate,omitempty"`
	LastViewedDate             time.Time         `xml:"LastViewedDate,omitempty"`
	LookedUpFromActivities     *QueryResult      `xml:"LookedUpFromActivities,omitempty"`
	Notes                      *QueryResult      `xml:"Notes,omitempty"`
	NotesAndAttachments        *QueryResult      `xml:"NotesAndAttachments,omitempty"`
	OpenActivities             *QueryResult      `xml:"OpenActivities,omitempty"`
	Opportunities              *QueryResult      `xml:"Opportunities,omitempty"`
	Orders                     *QueryResult      `xml:"Orders,omitempty"`
	Owner                      *User             `xml:"Owner,omitempty"`
	OwnerExpirationNotice      string            `xml:"OwnerExpirationNotice,omitempty"`
	OwnerId                    *ID               `xml:"OwnerId,omitempty"`
	Pricebook2                 *Pricebook2       `xml:"Pricebook2,omitempty"`
	Pricebook2Id               *ID               `xml:"Pricebook2Id,omitempty"`
	ProcessInstances           *QueryResult      `xml:"ProcessInstances,omitempty"`
	ProcessSteps               *QueryResult      `xml:"ProcessSteps,omitempty"`
	RecordActionHistories      *QueryResult      `xml:"RecordActionHistories,omitempty"`
	RecordActions              *QueryResult      `xml:"RecordActions,omitempty"`
	RecordAssociatedGroups     *QueryResult      `xml:"RecordAssociatedGroups,omitempty"`
	SpecialTerms               string            `xml:"SpecialTerms,omitempty"`
	StartDate                  time.Time         `xml:"StartDate,omitempty"`
	Status                     string            `xml:"Status,omitempty"`
	StatusCode                 string            `xml:"StatusCode,omitempty"`
	SystemModstamp             time.Time         `xml:"SystemModstamp,omitempty"`
	Tasks                      *QueryResult      `xml:"Tasks,omitempty"`
	TopicAssignments           *QueryResult      `xml:"TopicAssignments,omitempty"`
	UserRecordAccess           *UserRecordAccess `xml:"UserRecordAccess,omitempty"`
}

type ContractContactRole struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ContractContactRole"`

	*SObject

	Contact          *Contact  `xml:"Contact,omitempty"`
	ContactId        *ID       `xml:"ContactId,omitempty"`
	Contract         *Contract `xml:"Contract,omitempty"`
	ContractId       *ID       `xml:"ContractId,omitempty"`
	CreatedBy        *User     `xml:"CreatedBy,omitempty"`
	CreatedById      *ID       `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time `xml:"CreatedDate,omitempty"`
	IsDeleted        bool      `xml:"IsDeleted,omitempty"`
	IsPrimary        bool      `xml:"IsPrimary,omitempty"`
	LastModifiedBy   *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time `xml:"LastModifiedDate,omitempty"`
	Role             string    `xml:"Role,omitempty"`
	SystemModstamp   time.Time `xml:"SystemModstamp,omitempty"`
}

type ContractFeed struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ContractFeed"`

	*SObject

	BestComment        *FeedComment `xml:"BestComment,omitempty"`
	BestCommentId      *ID          `xml:"BestCommentId,omitempty"`
	Body               string       `xml:"Body,omitempty"`
	CommentCount       int32        `xml:"CommentCount,omitempty"`
	CreatedBy          *SObject     `xml:"CreatedBy,omitempty"`
	CreatedById        *ID          `xml:"CreatedById,omitempty"`
	CreatedDate        time.Time    `xml:"CreatedDate,omitempty"`
	FeedAttachments    *QueryResult `xml:"FeedAttachments,omitempty"`
	FeedComments       *QueryResult `xml:"FeedComments,omitempty"`
	FeedLikes          *QueryResult `xml:"FeedLikes,omitempty"`
	FeedSignals        *QueryResult `xml:"FeedSignals,omitempty"`
	FeedTrackedChanges *QueryResult `xml:"FeedTrackedChanges,omitempty"`
	InsertedBy         *SObject     `xml:"InsertedBy,omitempty"`
	InsertedById       *ID          `xml:"InsertedById,omitempty"`
	IsDeleted          bool         `xml:"IsDeleted,omitempty"`
	IsRichText         bool         `xml:"IsRichText,omitempty"`
	LastModifiedDate   time.Time    `xml:"LastModifiedDate,omitempty"`
	LikeCount          int32        `xml:"LikeCount,omitempty"`
	LinkUrl            string       `xml:"LinkUrl,omitempty"`
	Parent             *Contract    `xml:"Parent,omitempty"`
	ParentId           *ID          `xml:"ParentId,omitempty"`
	RelatedRecordId    *ID          `xml:"RelatedRecordId,omitempty"`
	SystemModstamp     time.Time    `xml:"SystemModstamp,omitempty"`
	Title              string       `xml:"Title,omitempty"`
	Type               string       `xml:"Type,omitempty"`
}

type ContractHistory struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ContractHistory"`

	*SObject

	Contract    *Contract   `xml:"Contract,omitempty"`
	ContractId  *ID         `xml:"ContractId,omitempty"`
	CreatedBy   *SObject    `xml:"CreatedBy,omitempty"`
	CreatedById *ID         `xml:"CreatedById,omitempty"`
	CreatedDate time.Time   `xml:"CreatedDate,omitempty"`
	Field       string      `xml:"Field,omitempty"`
	IsDeleted   bool        `xml:"IsDeleted,omitempty"`
	NewValue    interface{} `xml:"NewValue,omitempty"`
	OldValue    interface{} `xml:"OldValue,omitempty"`
}

type ContractStatus struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ContractStatus"`

	*SObject

	ApiName          string    `xml:"ApiName,omitempty"`
	CreatedBy        *User     `xml:"CreatedBy,omitempty"`
	CreatedById      *ID       `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time `xml:"CreatedDate,omitempty"`
	IsDefault        bool      `xml:"IsDefault,omitempty"`
	LastModifiedBy   *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time `xml:"LastModifiedDate,omitempty"`
	MasterLabel      string    `xml:"MasterLabel,omitempty"`
	SortOrder        int32     `xml:"SortOrder,omitempty"`
	StatusCode       string    `xml:"StatusCode,omitempty"`
	SystemModstamp   time.Time `xml:"SystemModstamp,omitempty"`
}

type CorsWhitelistEntry struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com CorsWhitelistEntry"`

	*SObject

	CreatedBy        *User     `xml:"CreatedBy,omitempty"`
	CreatedById      *ID       `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time `xml:"CreatedDate,omitempty"`
	DeveloperName    string    `xml:"DeveloperName,omitempty"`
	IsDeleted        bool      `xml:"IsDeleted,omitempty"`
	Language         string    `xml:"Language,omitempty"`
	LastModifiedBy   *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time `xml:"LastModifiedDate,omitempty"`
	MasterLabel      string    `xml:"MasterLabel,omitempty"`
	NamespacePrefix  string    `xml:"NamespacePrefix,omitempty"`
	SystemModstamp   time.Time `xml:"SystemModstamp,omitempty"`
	UrlPattern       string    `xml:"UrlPattern,omitempty"`
}

type CronJobDetail struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com CronJobDetail"`

	*SObject

	JobType string `xml:"JobType,omitempty"`
	Name    string `xml:"Name,omitempty"`
}

type CronTrigger struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com CronTrigger"`

	*SObject

	CreatedBy        *User          `xml:"CreatedBy,omitempty"`
	CreatedById      *ID            `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time      `xml:"CreatedDate,omitempty"`
	CronExpression   string         `xml:"CronExpression,omitempty"`
	CronJobDetail    *CronJobDetail `xml:"CronJobDetail,omitempty"`
	CronJobDetailId  *ID            `xml:"CronJobDetailId,omitempty"`
	EndTime          time.Time      `xml:"EndTime,omitempty"`
	LastModifiedBy   *User          `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID            `xml:"LastModifiedById,omitempty"`
	NextFireTime     time.Time      `xml:"NextFireTime,omitempty"`
	OwnerId          *ID            `xml:"OwnerId,omitempty"`
	PreviousFireTime time.Time      `xml:"PreviousFireTime,omitempty"`
	StartTime        time.Time      `xml:"StartTime,omitempty"`
	State            string         `xml:"State,omitempty"`
	TimeZoneSidKey   string         `xml:"TimeZoneSidKey,omitempty"`
	TimesTriggered   int32          `xml:"TimesTriggered,omitempty"`
}

type CspTrustedSite struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com CspTrustedSite"`

	*SObject

	Context          string    `xml:"Context,omitempty"`
	CreatedBy        *User     `xml:"CreatedBy,omitempty"`
	CreatedById      *ID       `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time `xml:"CreatedDate,omitempty"`
	Description      string    `xml:"Description,omitempty"`
	DeveloperName    string    `xml:"DeveloperName,omitempty"`
	EndpointUrl      string    `xml:"EndpointUrl,omitempty"`
	IsActive         bool      `xml:"IsActive,omitempty"`
	IsDeleted        bool      `xml:"IsDeleted,omitempty"`
	Language         string    `xml:"Language,omitempty"`
	LastModifiedBy   *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time `xml:"LastModifiedDate,omitempty"`
	MasterLabel      string    `xml:"MasterLabel,omitempty"`
	NamespacePrefix  string    `xml:"NamespacePrefix,omitempty"`
	SystemModstamp   time.Time `xml:"SystemModstamp,omitempty"`
}

type CustomBrand struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com CustomBrand"`

	*SObject

	CreatedBy         *User        `xml:"CreatedBy,omitempty"`
	CreatedById       *ID          `xml:"CreatedById,omitempty"`
	CreatedDate       time.Time    `xml:"CreatedDate,omitempty"`
	CustomBrandAssets *QueryResult `xml:"CustomBrandAssets,omitempty"`
	LastModifiedBy    *User        `xml:"LastModifiedBy,omitempty"`
	LastModifiedById  *ID          `xml:"LastModifiedById,omitempty"`
	LastModifiedDate  time.Time    `xml:"LastModifiedDate,omitempty"`
	ParentId          *ID          `xml:"ParentId,omitempty"`
}

type CustomBrandAsset struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com CustomBrandAsset"`

	*SObject

	AssetCategory    string       `xml:"AssetCategory,omitempty"`
	AssetSourceId    *ID          `xml:"AssetSourceId,omitempty"`
	CreatedBy        *User        `xml:"CreatedBy,omitempty"`
	CreatedById      *ID          `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time    `xml:"CreatedDate,omitempty"`
	CustomBrand      *CustomBrand `xml:"CustomBrand,omitempty"`
	CustomBrandId    *ID          `xml:"CustomBrandId,omitempty"`
	LastModifiedBy   *User        `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID          `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time    `xml:"LastModifiedDate,omitempty"`
	TextAsset        string       `xml:"TextAsset,omitempty"`
}

type CustomHttpHeader struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com CustomHttpHeader"`

	*SObject

	CreatedBy        *User     `xml:"CreatedBy,omitempty"`
	CreatedById      *ID       `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time `xml:"CreatedDate,omitempty"`
	Description      string    `xml:"Description,omitempty"`
	HeaderFieldName  string    `xml:"HeaderFieldName,omitempty"`
	HeaderFieldValue string    `xml:"HeaderFieldValue,omitempty"`
	IsActive         bool      `xml:"IsActive,omitempty"`
	LastModifiedBy   *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time `xml:"LastModifiedDate,omitempty"`
	Parent           *SObject  `xml:"Parent,omitempty"`
	ParentId         *ID       `xml:"ParentId,omitempty"`
	SystemModstamp   time.Time `xml:"SystemModstamp,omitempty"`
}

type CustomObjectUserLicenseMetrics struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com CustomObjectUserLicenseMetrics"`

	*SObject

	CustomObjectId   string       `xml:"CustomObjectId,omitempty"`
	CustomObjectName string       `xml:"CustomObjectName,omitempty"`
	CustomObjectType string       `xml:"CustomObjectType,omitempty"`
	MetricsDate      time.Time    `xml:"MetricsDate,omitempty"`
	ObjectCount      int32        `xml:"ObjectCount,omitempty"`
	SystemModstamp   time.Time    `xml:"SystemModstamp,omitempty"`
	UserLicense      *UserLicense `xml:"UserLicense,omitempty"`
	UserLicenseId    *ID          `xml:"UserLicenseId,omitempty"`
}

type CustomPermission struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com CustomPermission"`

	*SObject

	CreatedBy                      *User        `xml:"CreatedBy,omitempty"`
	CreatedById                    *ID          `xml:"CreatedById,omitempty"`
	CreatedDate                    time.Time    `xml:"CreatedDate,omitempty"`
	CustomPermissionDependencyItem *QueryResult `xml:"CustomPermissionDependencyItem,omitempty"`
	CustomPermissionItem           *QueryResult `xml:"CustomPermissionItem,omitempty"`
	Description                    string       `xml:"Description,omitempty"`
	DeveloperName                  string       `xml:"DeveloperName,omitempty"`
	GrantedByLicenses              *QueryResult `xml:"GrantedByLicenses,omitempty"`
	IsDeleted                      bool         `xml:"IsDeleted,omitempty"`
	IsProtected                    bool         `xml:"IsProtected,omitempty"`
	Language                       string       `xml:"Language,omitempty"`
	LastModifiedBy                 *User        `xml:"LastModifiedBy,omitempty"`
	LastModifiedById               *ID          `xml:"LastModifiedById,omitempty"`
	LastModifiedDate               time.Time    `xml:"LastModifiedDate,omitempty"`
	MasterLabel                    string       `xml:"MasterLabel,omitempty"`
	NamespacePrefix                string       `xml:"NamespacePrefix,omitempty"`
	SetupEntityAccessItems         *QueryResult `xml:"SetupEntityAccessItems,omitempty"`
	SystemModstamp                 time.Time    `xml:"SystemModstamp,omitempty"`
}

type CustomPermissionDependency struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com CustomPermissionDependency"`

	*SObject

	CreatedBy                  *User             `xml:"CreatedBy,omitempty"`
	CreatedById                *ID               `xml:"CreatedById,omitempty"`
	CreatedDate                time.Time         `xml:"CreatedDate,omitempty"`
	CustomPermission           *CustomPermission `xml:"CustomPermission,omitempty"`
	CustomPermissionId         *ID               `xml:"CustomPermissionId,omitempty"`
	IsDeleted                  bool              `xml:"IsDeleted,omitempty"`
	LastModifiedBy             *User             `xml:"LastModifiedBy,omitempty"`
	LastModifiedById           *ID               `xml:"LastModifiedById,omitempty"`
	LastModifiedDate           time.Time         `xml:"LastModifiedDate,omitempty"`
	RequiredCustomPermission   *CustomPermission `xml:"RequiredCustomPermission,omitempty"`
	RequiredCustomPermissionId *ID               `xml:"RequiredCustomPermissionId,omitempty"`
	SystemModstamp             time.Time         `xml:"SystemModstamp,omitempty"`
}

type DandBCompany struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com DandBCompany"`

	*SObject

	Accounts                     *QueryResult      `xml:"Accounts,omitempty"`
	Address                      *Address          `xml:"Address,omitempty"`
	City                         string            `xml:"City,omitempty"`
	CompanyCurrencyIsoCode       string            `xml:"CompanyCurrencyIsoCode,omitempty"`
	Country                      string            `xml:"Country,omitempty"`
	CountryAccessCode            string            `xml:"CountryAccessCode,omitempty"`
	CreatedBy                    *User             `xml:"CreatedBy,omitempty"`
	CreatedById                  *ID               `xml:"CreatedById,omitempty"`
	CreatedDate                  time.Time         `xml:"CreatedDate,omitempty"`
	CurrencyCode                 string            `xml:"CurrencyCode,omitempty"`
	Description                  string            `xml:"Description,omitempty"`
	DomesticUltimateBusinessName string            `xml:"DomesticUltimateBusinessName,omitempty"`
	DomesticUltimateDunsNumber   string            `xml:"DomesticUltimateDunsNumber,omitempty"`
	DunsNumber                   string            `xml:"DunsNumber,omitempty"`
	EmployeeQuantityGrowthRate   float64           `xml:"EmployeeQuantityGrowthRate,omitempty"`
	EmployeesHere                float64           `xml:"EmployeesHere,omitempty"`
	EmployeesHereReliability     string            `xml:"EmployeesHereReliability,omitempty"`
	EmployeesTotal               float64           `xml:"EmployeesTotal,omitempty"`
	EmployeesTotalReliability    string            `xml:"EmployeesTotalReliability,omitempty"`
	FamilyMembers                int32             `xml:"FamilyMembers,omitempty"`
	Fax                          string            `xml:"Fax,omitempty"`
	FifthNaics                   string            `xml:"FifthNaics,omitempty"`
	FifthNaicsDesc               string            `xml:"FifthNaicsDesc,omitempty"`
	FifthSic                     string            `xml:"FifthSic,omitempty"`
	FifthSic8                    string            `xml:"FifthSic8,omitempty"`
	FifthSic8Desc                string            `xml:"FifthSic8Desc,omitempty"`
	FifthSicDesc                 string            `xml:"FifthSicDesc,omitempty"`
	FipsMsaCode                  string            `xml:"FipsMsaCode,omitempty"`
	FipsMsaDesc                  string            `xml:"FipsMsaDesc,omitempty"`
	FortuneRank                  int32             `xml:"FortuneRank,omitempty"`
	FourthNaics                  string            `xml:"FourthNaics,omitempty"`
	FourthNaicsDesc              string            `xml:"FourthNaicsDesc,omitempty"`
	FourthSic                    string            `xml:"FourthSic,omitempty"`
	FourthSic8                   string            `xml:"FourthSic8,omitempty"`
	FourthSic8Desc               string            `xml:"FourthSic8Desc,omitempty"`
	FourthSicDesc                string            `xml:"FourthSicDesc,omitempty"`
	GeoCodeAccuracy              string            `xml:"GeoCodeAccuracy,omitempty"`
	GeocodeAccuracyStandard      string            `xml:"GeocodeAccuracyStandard,omitempty"`
	GlobalUltimateBusinessName   string            `xml:"GlobalUltimateBusinessName,omitempty"`
	GlobalUltimateDunsNumber     string            `xml:"GlobalUltimateDunsNumber,omitempty"`
	GlobalUltimateTotalEmployees float64           `xml:"GlobalUltimateTotalEmployees,omitempty"`
	ImportExportAgent            string            `xml:"ImportExportAgent,omitempty"`
	IncludedInSnP500             string            `xml:"IncludedInSnP500,omitempty"`
	IsDeleted                    bool              `xml:"IsDeleted,omitempty"`
	LastModifiedBy               *User             `xml:"LastModifiedBy,omitempty"`
	LastModifiedById             *ID               `xml:"LastModifiedById,omitempty"`
	LastModifiedDate             time.Time         `xml:"LastModifiedDate,omitempty"`
	LastReferencedDate           time.Time         `xml:"LastReferencedDate,omitempty"`
	LastViewedDate               time.Time         `xml:"LastViewedDate,omitempty"`
	Latitude                     string            `xml:"Latitude,omitempty"`
	Leads                        *QueryResult      `xml:"Leads,omitempty"`
	LegalStatus                  string            `xml:"LegalStatus,omitempty"`
	LocationStatus               string            `xml:"LocationStatus,omitempty"`
	Longitude                    string            `xml:"Longitude,omitempty"`
	MailingAddress               *Address          `xml:"MailingAddress,omitempty"`
	MailingCity                  string            `xml:"MailingCity,omitempty"`
	MailingCountry               string            `xml:"MailingCountry,omitempty"`
	MailingGeocodeAccuracy       string            `xml:"MailingGeocodeAccuracy,omitempty"`
	MailingPostalCode            string            `xml:"MailingPostalCode,omitempty"`
	MailingState                 string            `xml:"MailingState,omitempty"`
	MailingStreet                string            `xml:"MailingStreet,omitempty"`
	MarketingPreScreen           string            `xml:"MarketingPreScreen,omitempty"`
	MarketingSegmentationCluster string            `xml:"MarketingSegmentationCluster,omitempty"`
	MinorityOwned                string            `xml:"MinorityOwned,omitempty"`
	Name                         string            `xml:"Name,omitempty"`
	NationalId                   string            `xml:"NationalId,omitempty"`
	NationalIdType               string            `xml:"NationalIdType,omitempty"`
	OutOfBusiness                string            `xml:"OutOfBusiness,omitempty"`
	OwnOrRent                    string            `xml:"OwnOrRent,omitempty"`
	ParentOrHqBusinessName       string            `xml:"ParentOrHqBusinessName,omitempty"`
	ParentOrHqDunsNumber         string            `xml:"ParentOrHqDunsNumber,omitempty"`
	Phone                        string            `xml:"Phone,omitempty"`
	PostalCode                   string            `xml:"PostalCode,omitempty"`
	PremisesMeasure              int32             `xml:"PremisesMeasure,omitempty"`
	PremisesMeasureReliability   string            `xml:"PremisesMeasureReliability,omitempty"`
	PremisesMeasureUnit          string            `xml:"PremisesMeasureUnit,omitempty"`
	PrimaryNaics                 string            `xml:"PrimaryNaics,omitempty"`
	PrimaryNaicsDesc             string            `xml:"PrimaryNaicsDesc,omitempty"`
	PrimarySic                   string            `xml:"PrimarySic,omitempty"`
	PrimarySic8                  string            `xml:"PrimarySic8,omitempty"`
	PrimarySic8Desc              string            `xml:"PrimarySic8Desc,omitempty"`
	PrimarySicDesc               string            `xml:"PrimarySicDesc,omitempty"`
	PriorYearEmployees           int32             `xml:"PriorYearEmployees,omitempty"`
	PriorYearRevenue             float64           `xml:"PriorYearRevenue,omitempty"`
	PublicIndicator              string            `xml:"PublicIndicator,omitempty"`
	SalesTurnoverGrowthRate      float64           `xml:"SalesTurnoverGrowthRate,omitempty"`
	SalesVolume                  float64           `xml:"SalesVolume,omitempty"`
	SalesVolumeReliability       string            `xml:"SalesVolumeReliability,omitempty"`
	SecondNaics                  string            `xml:"SecondNaics,omitempty"`
	SecondNaicsDesc              string            `xml:"SecondNaicsDesc,omitempty"`
	SecondSic                    string            `xml:"SecondSic,omitempty"`
	SecondSic8                   string            `xml:"SecondSic8,omitempty"`
	SecondSic8Desc               string            `xml:"SecondSic8Desc,omitempty"`
	SecondSicDesc                string            `xml:"SecondSicDesc,omitempty"`
	SixthNaics                   string            `xml:"SixthNaics,omitempty"`
	SixthNaicsDesc               string            `xml:"SixthNaicsDesc,omitempty"`
	SixthSic                     string            `xml:"SixthSic,omitempty"`
	SixthSic8                    string            `xml:"SixthSic8,omitempty"`
	SixthSic8Desc                string            `xml:"SixthSic8Desc,omitempty"`
	SixthSicDesc                 string            `xml:"SixthSicDesc,omitempty"`
	SmallBusiness                string            `xml:"SmallBusiness,omitempty"`
	State                        string            `xml:"State,omitempty"`
	StockExchange                string            `xml:"StockExchange,omitempty"`
	StockSymbol                  string            `xml:"StockSymbol,omitempty"`
	Street                       string            `xml:"Street,omitempty"`
	Subsidiary                   string            `xml:"Subsidiary,omitempty"`
	SystemModstamp               time.Time         `xml:"SystemModstamp,omitempty"`
	ThirdNaics                   string            `xml:"ThirdNaics,omitempty"`
	ThirdNaicsDesc               string            `xml:"ThirdNaicsDesc,omitempty"`
	ThirdSic                     string            `xml:"ThirdSic,omitempty"`
	ThirdSic8                    string            `xml:"ThirdSic8,omitempty"`
	ThirdSic8Desc                string            `xml:"ThirdSic8Desc,omitempty"`
	ThirdSicDesc                 string            `xml:"ThirdSicDesc,omitempty"`
	TradeStyle1                  string            `xml:"TradeStyle1,omitempty"`
	TradeStyle2                  string            `xml:"TradeStyle2,omitempty"`
	TradeStyle3                  string            `xml:"TradeStyle3,omitempty"`
	TradeStyle4                  string            `xml:"TradeStyle4,omitempty"`
	TradeStyle5                  string            `xml:"TradeStyle5,omitempty"`
	URL                          string            `xml:"URL,omitempty"`
	UsTaxId                      string            `xml:"UsTaxId,omitempty"`
	UserRecordAccess             *UserRecordAccess `xml:"UserRecordAccess,omitempty"`
	WomenOwned                   string            `xml:"WomenOwned,omitempty"`
	YearStarted                  string            `xml:"YearStarted,omitempty"`
}

type Dashboard struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com Dashboard"`

	*SObject

	AttachedContentDocuments     *QueryResult `xml:"AttachedContentDocuments,omitempty"`
	BackgroundDirection          string       `xml:"BackgroundDirection,omitempty"`
	BackgroundEnd                int32        `xml:"BackgroundEnd,omitempty"`
	BackgroundStart              int32        `xml:"BackgroundStart,omitempty"`
	ChartTheme                   string       `xml:"ChartTheme,omitempty"`
	ColorPalette                 string       `xml:"ColorPalette,omitempty"`
	CombinedAttachments          *QueryResult `xml:"CombinedAttachments,omitempty"`
	ContentDocumentLinks         *QueryResult `xml:"ContentDocumentLinks,omitempty"`
	CreatedBy                    *User        `xml:"CreatedBy,omitempty"`
	CreatedById                  *ID          `xml:"CreatedById,omitempty"`
	CreatedDate                  time.Time    `xml:"CreatedDate,omitempty"`
	DashboardComponents          *QueryResult `xml:"DashboardComponents,omitempty"`
	DashboardResultRefreshedDate string       `xml:"DashboardResultRefreshedDate,omitempty"`
	DashboardResultRunningUser   string       `xml:"DashboardResultRunningUser,omitempty"`
	Description                  string       `xml:"Description,omitempty"`
	DeveloperName                string       `xml:"DeveloperName,omitempty"`
	FeedSubscriptionsForEntity   *QueryResult `xml:"FeedSubscriptionsForEntity,omitempty"`
	Feeds                        *QueryResult `xml:"Feeds,omitempty"`
	Folder                       *Folder      `xml:"Folder,omitempty"`
	FolderId                     *ID          `xml:"FolderId,omitempty"`
	FolderName                   string       `xml:"FolderName,omitempty"`
	IsDeleted                    bool         `xml:"IsDeleted,omitempty"`
	LastModifiedBy               *User        `xml:"LastModifiedBy,omitempty"`
	LastModifiedById             *ID          `xml:"LastModifiedById,omitempty"`
	LastModifiedDate             time.Time    `xml:"LastModifiedDate,omitempty"`
	LastReferencedDate           time.Time    `xml:"LastReferencedDate,omitempty"`
	LastViewedDate               time.Time    `xml:"LastViewedDate,omitempty"`
	LeftSize                     string       `xml:"LeftSize,omitempty"`
	MiddleSize                   string       `xml:"MiddleSize,omitempty"`
	NamespacePrefix              string       `xml:"NamespacePrefix,omitempty"`
	RightSize                    string       `xml:"RightSize,omitempty"`
	RunningUser                  *User        `xml:"RunningUser,omitempty"`
	RunningUserId                *ID          `xml:"RunningUserId,omitempty"`
	SystemModstamp               time.Time    `xml:"SystemModstamp,omitempty"`
	TextColor                    int32        `xml:"TextColor,omitempty"`
	Title                        string       `xml:"Title,omitempty"`
	TitleColor                   int32        `xml:"TitleColor,omitempty"`
	TitleSize                    int32        `xml:"TitleSize,omitempty"`
	Type                         string       `xml:"Type,omitempty"`
}

type DashboardComponent struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com DashboardComponent"`

	*SObject

	AttachedContentDocuments   *QueryResult `xml:"AttachedContentDocuments,omitempty"`
	CombinedAttachments        *QueryResult `xml:"CombinedAttachments,omitempty"`
	ContentDocumentLinks       *QueryResult `xml:"ContentDocumentLinks,omitempty"`
	CustomReportId             *ID          `xml:"CustomReportId,omitempty"`
	Dashboard                  *Dashboard   `xml:"Dashboard,omitempty"`
	DashboardId                *ID          `xml:"DashboardId,omitempty"`
	FeedSubscriptionsForEntity *QueryResult `xml:"FeedSubscriptionsForEntity,omitempty"`
	Feeds                      *QueryResult `xml:"Feeds,omitempty"`
	Name                       string       `xml:"Name,omitempty"`
}

type DashboardComponentFeed struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com DashboardComponentFeed"`

	*SObject

	BestComment        *FeedComment        `xml:"BestComment,omitempty"`
	BestCommentId      *ID                 `xml:"BestCommentId,omitempty"`
	Body               string              `xml:"Body,omitempty"`
	CommentCount       int32               `xml:"CommentCount,omitempty"`
	CreatedBy          *SObject            `xml:"CreatedBy,omitempty"`
	CreatedById        *ID                 `xml:"CreatedById,omitempty"`
	CreatedDate        time.Time           `xml:"CreatedDate,omitempty"`
	FeedAttachments    *QueryResult        `xml:"FeedAttachments,omitempty"`
	FeedComments       *QueryResult        `xml:"FeedComments,omitempty"`
	FeedLikes          *QueryResult        `xml:"FeedLikes,omitempty"`
	FeedSignals        *QueryResult        `xml:"FeedSignals,omitempty"`
	FeedTrackedChanges *QueryResult        `xml:"FeedTrackedChanges,omitempty"`
	InsertedBy         *SObject            `xml:"InsertedBy,omitempty"`
	InsertedById       *ID                 `xml:"InsertedById,omitempty"`
	IsDeleted          bool                `xml:"IsDeleted,omitempty"`
	IsRichText         bool                `xml:"IsRichText,omitempty"`
	LastModifiedDate   time.Time           `xml:"LastModifiedDate,omitempty"`
	LikeCount          int32               `xml:"LikeCount,omitempty"`
	LinkUrl            string              `xml:"LinkUrl,omitempty"`
	Parent             *DashboardComponent `xml:"Parent,omitempty"`
	ParentId           *ID                 `xml:"ParentId,omitempty"`
	RelatedRecordId    *ID                 `xml:"RelatedRecordId,omitempty"`
	SystemModstamp     time.Time           `xml:"SystemModstamp,omitempty"`
	Title              string              `xml:"Title,omitempty"`
	Type               string              `xml:"Type,omitempty"`
}

type DashboardFeed struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com DashboardFeed"`

	*SObject

	BestComment        *FeedComment `xml:"BestComment,omitempty"`
	BestCommentId      *ID          `xml:"BestCommentId,omitempty"`
	Body               string       `xml:"Body,omitempty"`
	CommentCount       int32        `xml:"CommentCount,omitempty"`
	CreatedBy          *SObject     `xml:"CreatedBy,omitempty"`
	CreatedById        *ID          `xml:"CreatedById,omitempty"`
	CreatedDate        time.Time    `xml:"CreatedDate,omitempty"`
	FeedAttachments    *QueryResult `xml:"FeedAttachments,omitempty"`
	FeedComments       *QueryResult `xml:"FeedComments,omitempty"`
	FeedLikes          *QueryResult `xml:"FeedLikes,omitempty"`
	FeedSignals        *QueryResult `xml:"FeedSignals,omitempty"`
	FeedTrackedChanges *QueryResult `xml:"FeedTrackedChanges,omitempty"`
	InsertedBy         *SObject     `xml:"InsertedBy,omitempty"`
	InsertedById       *ID          `xml:"InsertedById,omitempty"`
	IsDeleted          bool         `xml:"IsDeleted,omitempty"`
	IsRichText         bool         `xml:"IsRichText,omitempty"`
	LastModifiedDate   time.Time    `xml:"LastModifiedDate,omitempty"`
	LikeCount          int32        `xml:"LikeCount,omitempty"`
	LinkUrl            string       `xml:"LinkUrl,omitempty"`
	Parent             *Dashboard   `xml:"Parent,omitempty"`
	ParentId           *ID          `xml:"ParentId,omitempty"`
	RelatedRecordId    *ID          `xml:"RelatedRecordId,omitempty"`
	SystemModstamp     time.Time    `xml:"SystemModstamp,omitempty"`
	Title              string       `xml:"Title,omitempty"`
	Type               string       `xml:"Type,omitempty"`
}

type DataAssessmentFieldMetric struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com DataAssessmentFieldMetric"`

	*SObject

	CreatedBy                  *User                 `xml:"CreatedBy,omitempty"`
	CreatedById                *ID                   `xml:"CreatedById,omitempty"`
	CreatedDate                time.Time             `xml:"CreatedDate,omitempty"`
	DataAssessmentMetric       *DataAssessmentMetric `xml:"DataAssessmentMetric,omitempty"`
	DataAssessmentMetricId     *ID                   `xml:"DataAssessmentMetricId,omitempty"`
	DataAssessmentValueMetrics *QueryResult          `xml:"DataAssessmentValueMetrics,omitempty"`
	FieldName                  string                `xml:"FieldName,omitempty"`
	IsDeleted                  bool                  `xml:"IsDeleted,omitempty"`
	LastModifiedBy             *User                 `xml:"LastModifiedBy,omitempty"`
	LastModifiedById           *ID                   `xml:"LastModifiedById,omitempty"`
	LastModifiedDate           time.Time             `xml:"LastModifiedDate,omitempty"`
	Name                       string                `xml:"Name,omitempty"`
	NumMatchedBlanks           int32                 `xml:"NumMatchedBlanks,omitempty"`
	NumMatchedDifferent        int32                 `xml:"NumMatchedDifferent,omitempty"`
	NumMatchedInSync           int32                 `xml:"NumMatchedInSync,omitempty"`
	NumUnmatchedBlanks         int32                 `xml:"NumUnmatchedBlanks,omitempty"`
	SystemModstamp             time.Time             `xml:"SystemModstamp,omitempty"`
	UserRecordAccess           *UserRecordAccess     `xml:"UserRecordAccess,omitempty"`
}

type DataAssessmentMetric struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com DataAssessmentMetric"`

	*SObject

	CreatedBy             *User             `xml:"CreatedBy,omitempty"`
	CreatedById           *ID               `xml:"CreatedById,omitempty"`
	CreatedDate           time.Time         `xml:"CreatedDate,omitempty"`
	DataAssessmentMetrics *QueryResult      `xml:"DataAssessmentMetrics,omitempty"`
	IsDeleted             bool              `xml:"IsDeleted,omitempty"`
	LastModifiedBy        *User             `xml:"LastModifiedBy,omitempty"`
	LastModifiedById      *ID               `xml:"LastModifiedById,omitempty"`
	LastModifiedDate      time.Time         `xml:"LastModifiedDate,omitempty"`
	Name                  string            `xml:"Name,omitempty"`
	NumDuplicates         int32             `xml:"NumDuplicates,omitempty"`
	NumMatched            int32             `xml:"NumMatched,omitempty"`
	NumMatchedDifferent   int32             `xml:"NumMatchedDifferent,omitempty"`
	NumProcessed          int32             `xml:"NumProcessed,omitempty"`
	NumTotal              int32             `xml:"NumTotal,omitempty"`
	NumUnmatched          int32             `xml:"NumUnmatched,omitempty"`
	SystemModstamp        time.Time         `xml:"SystemModstamp,omitempty"`
	UserRecordAccess      *UserRecordAccess `xml:"UserRecordAccess,omitempty"`
}

type DataAssessmentValueMetric struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com DataAssessmentValueMetric"`

	*SObject

	CreatedBy                   *User                      `xml:"CreatedBy,omitempty"`
	CreatedById                 *ID                        `xml:"CreatedById,omitempty"`
	CreatedDate                 time.Time                  `xml:"CreatedDate,omitempty"`
	DataAssessmentFieldMetric   *DataAssessmentFieldMetric `xml:"DataAssessmentFieldMetric,omitempty"`
	DataAssessmentFieldMetricId *ID                        `xml:"DataAssessmentFieldMetricId,omitempty"`
	FieldValue                  string                     `xml:"FieldValue,omitempty"`
	IsDeleted                   bool                       `xml:"IsDeleted,omitempty"`
	LastModifiedBy              *User                      `xml:"LastModifiedBy,omitempty"`
	LastModifiedById            *ID                        `xml:"LastModifiedById,omitempty"`
	LastModifiedDate            time.Time                  `xml:"LastModifiedDate,omitempty"`
	Name                        string                     `xml:"Name,omitempty"`
	SystemModstamp              time.Time                  `xml:"SystemModstamp,omitempty"`
	UserRecordAccess            *UserRecordAccess          `xml:"UserRecordAccess,omitempty"`
	ValueCount                  int32                      `xml:"ValueCount,omitempty"`
}

type DataStatistics struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com DataStatistics"`

	*SObject

	ExternalId string `xml:"ExternalId,omitempty"`
	StatType   string `xml:"StatType,omitempty"`
	StatValue  int32  `xml:"StatValue,omitempty"`
	Type       string `xml:"Type,omitempty"`
	User       *User  `xml:"User,omitempty"`
	UserId     *ID    `xml:"UserId,omitempty"`
}

type DataType struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com DataType"`

	*SObject

	ContextServiceDataTypeId string `xml:"ContextServiceDataTypeId,omitempty"`
	ContextWsdlDataTypeId    string `xml:"ContextWsdlDataTypeId,omitempty"`
	DeveloperName            string `xml:"DeveloperName,omitempty"`
	DurableId                string `xml:"DurableId,omitempty"`
	IsComplex                bool   `xml:"IsComplex,omitempty"`
}

type DatacloudAddress struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com DatacloudAddress"`

	*SObject

	AddressLine1    string `xml:"AddressLine1,omitempty"`
	AddressLine2    string `xml:"AddressLine2,omitempty"`
	City            string `xml:"City,omitempty"`
	Country         string `xml:"Country,omitempty"`
	ExternalId      string `xml:"ExternalId,omitempty"`
	GeoAccuracyCode string `xml:"GeoAccuracyCode,omitempty"`
	GeoAccuracyNum  string `xml:"GeoAccuracyNum,omitempty"`
	Latitude        string `xml:"Latitude,omitempty"`
	Longitude       string `xml:"Longitude,omitempty"`
	PostalCode      string `xml:"PostalCode,omitempty"`
	State           string `xml:"State,omitempty"`
}

type DatacloudCompany struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com DatacloudCompany"`

	*SObject

	ActiveContacts             int32     `xml:"ActiveContacts,omitempty"`
	AnnualRevenue              float64   `xml:"AnnualRevenue,omitempty"`
	City                       string    `xml:"City,omitempty"`
	CompanyId                  string    `xml:"CompanyId,omitempty"`
	Country                    string    `xml:"Country,omitempty"`
	CountryCode                string    `xml:"CountryCode,omitempty"`
	Description                string    `xml:"Description,omitempty"`
	DunsNumber                 string    `xml:"DunsNumber,omitempty"`
	EmployeeQuantityGrowthRate float64   `xml:"EmployeeQuantityGrowthRate,omitempty"`
	ExternalId                 string    `xml:"ExternalId,omitempty"`
	Fax                        string    `xml:"Fax,omitempty"`
	FortuneRank                int32     `xml:"FortuneRank,omitempty"`
	FullAddress                string    `xml:"FullAddress,omitempty"`
	IncludedInSnP500           string    `xml:"IncludedInSnP500,omitempty"`
	Industry                   string    `xml:"Industry,omitempty"`
	IsInCrm                    bool      `xml:"IsInCrm,omitempty"`
	IsInactive                 bool      `xml:"IsInactive,omitempty"`
	IsOwned                    bool      `xml:"IsOwned,omitempty"`
	NaicsCode                  string    `xml:"NaicsCode,omitempty"`
	NaicsDesc                  string    `xml:"NaicsDesc,omitempty"`
	Name                       string    `xml:"Name,omitempty"`
	NumberOfEmployees          int32     `xml:"NumberOfEmployees,omitempty"`
	Ownership                  string    `xml:"Ownership,omitempty"`
	Phone                      string    `xml:"Phone,omitempty"`
	PremisesMeasure            int32     `xml:"PremisesMeasure,omitempty"`
	PremisesMeasureReliability string    `xml:"PremisesMeasureReliability,omitempty"`
	PremisesMeasureUnit        string    `xml:"PremisesMeasureUnit,omitempty"`
	PriorYearEmployees         int32     `xml:"PriorYearEmployees,omitempty"`
	PriorYearRevenue           float64   `xml:"PriorYearRevenue,omitempty"`
	SalesTurnoverGrowthRate    float64   `xml:"SalesTurnoverGrowthRate,omitempty"`
	Sic                        string    `xml:"Sic,omitempty"`
	SicCodeDesc                string    `xml:"SicCodeDesc,omitempty"`
	SicDesc                    string    `xml:"SicDesc,omitempty"`
	Site                       string    `xml:"Site,omitempty"`
	State                      string    `xml:"State,omitempty"`
	StateCode                  string    `xml:"StateCode,omitempty"`
	Street                     string    `xml:"Street,omitempty"`
	TickerSymbol               string    `xml:"TickerSymbol,omitempty"`
	TradeStyle                 string    `xml:"TradeStyle,omitempty"`
	UpdatedDate                time.Time `xml:"UpdatedDate,omitempty"`
	Website                    string    `xml:"Website,omitempty"`
	YearStarted                string    `xml:"YearStarted,omitempty"`
	Zip                        string    `xml:"Zip,omitempty"`
}

type DatacloudContact struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com DatacloudContact"`

	*SObject

	City        string    `xml:"City,omitempty"`
	CompanyId   string    `xml:"CompanyId,omitempty"`
	CompanyName string    `xml:"CompanyName,omitempty"`
	ContactId   string    `xml:"ContactId,omitempty"`
	Country     string    `xml:"Country,omitempty"`
	Department  string    `xml:"Department,omitempty"`
	Email       string    `xml:"Email,omitempty"`
	ExternalId  string    `xml:"ExternalId,omitempty"`
	FirstName   string    `xml:"FirstName,omitempty"`
	IsInCrm     bool      `xml:"IsInCrm,omitempty"`
	IsInactive  bool      `xml:"IsInactive,omitempty"`
	IsOwned     bool      `xml:"IsOwned,omitempty"`
	LastName    string    `xml:"LastName,omitempty"`
	Level       string    `xml:"Level,omitempty"`
	Phone       string    `xml:"Phone,omitempty"`
	State       string    `xml:"State,omitempty"`
	Street      string    `xml:"Street,omitempty"`
	Title       string    `xml:"Title,omitempty"`
	UpdatedDate time.Time `xml:"UpdatedDate,omitempty"`
	Zip         string    `xml:"Zip,omitempty"`
}

type DatacloudDandBCompany struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com DatacloudDandBCompany"`

	*SObject

	City                         string  `xml:"City,omitempty"`
	CompanyCurrencyIsoCode       string  `xml:"CompanyCurrencyIsoCode,omitempty"`
	CompanyId                    string  `xml:"CompanyId,omitempty"`
	Country                      string  `xml:"Country,omitempty"`
	CountryAccessCode            string  `xml:"CountryAccessCode,omitempty"`
	CountryCode                  string  `xml:"CountryCode,omitempty"`
	CurrencyCode                 string  `xml:"CurrencyCode,omitempty"`
	Description                  string  `xml:"Description,omitempty"`
	DomesticUltimateBusinessName string  `xml:"DomesticUltimateBusinessName,omitempty"`
	DomesticUltimateDunsNumber   string  `xml:"DomesticUltimateDunsNumber,omitempty"`
	DunsNumber                   string  `xml:"DunsNumber,omitempty"`
	EmployeeQuantityGrowthRate   float64 `xml:"EmployeeQuantityGrowthRate,omitempty"`
	EmployeesHere                float64 `xml:"EmployeesHere,omitempty"`
	EmployeesHereReliability     string  `xml:"EmployeesHereReliability,omitempty"`
	EmployeesTotal               float64 `xml:"EmployeesTotal,omitempty"`
	EmployeesTotalReliability    string  `xml:"EmployeesTotalReliability,omitempty"`
	ExternalId                   string  `xml:"ExternalId,omitempty"`
	FamilyMembers                int32   `xml:"FamilyMembers,omitempty"`
	Fax                          string  `xml:"Fax,omitempty"`
	FifthNaics                   string  `xml:"FifthNaics,omitempty"`
	FifthNaicsDesc               string  `xml:"FifthNaicsDesc,omitempty"`
	FifthSic                     string  `xml:"FifthSic,omitempty"`
	FifthSic8                    string  `xml:"FifthSic8,omitempty"`
	FifthSic8Desc                string  `xml:"FifthSic8Desc,omitempty"`
	FifthSicDesc                 string  `xml:"FifthSicDesc,omitempty"`
	FipsMsaCode                  string  `xml:"FipsMsaCode,omitempty"`
	FipsMsaDesc                  string  `xml:"FipsMsaDesc,omitempty"`
	FortuneRank                  int32   `xml:"FortuneRank,omitempty"`
	FourthNaics                  string  `xml:"FourthNaics,omitempty"`
	FourthNaicsDesc              string  `xml:"FourthNaicsDesc,omitempty"`
	FourthSic                    string  `xml:"FourthSic,omitempty"`
	FourthSic8                   string  `xml:"FourthSic8,omitempty"`
	FourthSic8Desc               string  `xml:"FourthSic8Desc,omitempty"`
	FourthSicDesc                string  `xml:"FourthSicDesc,omitempty"`
	FullAddress                  string  `xml:"FullAddress,omitempty"`
	GeoCodeAccuracy              string  `xml:"GeoCodeAccuracy,omitempty"`
	GlobalUltimateBusinessName   string  `xml:"GlobalUltimateBusinessName,omitempty"`
	GlobalUltimateDunsNumber     string  `xml:"GlobalUltimateDunsNumber,omitempty"`
	GlobalUltimateTotalEmployees float64 `xml:"GlobalUltimateTotalEmployees,omitempty"`
	ImportExportAgent            string  `xml:"ImportExportAgent,omitempty"`
	IncludedInSnP500             string  `xml:"IncludedInSnP500,omitempty"`
	Industry                     string  `xml:"Industry,omitempty"`
	IsInCrm                      bool    `xml:"IsInCrm,omitempty"`
	IsOwned                      bool    `xml:"IsOwned,omitempty"`
	IsParent                     bool    `xml:"IsParent,omitempty"`
	Latitude                     string  `xml:"Latitude,omitempty"`
	LegalStatus                  string  `xml:"LegalStatus,omitempty"`
	LocationStatus               string  `xml:"LocationStatus,omitempty"`
	Longitude                    string  `xml:"Longitude,omitempty"`
	MailingCity                  string  `xml:"MailingCity,omitempty"`
	MailingCountry               string  `xml:"MailingCountry,omitempty"`
	MailingCountryCode           string  `xml:"MailingCountryCode,omitempty"`
	MailingState                 string  `xml:"MailingState,omitempty"`
	MailingStateCode             string  `xml:"MailingStateCode,omitempty"`
	MailingStreet                string  `xml:"MailingStreet,omitempty"`
	MailingZip                   string  `xml:"MailingZip,omitempty"`
	MarketingPreScreen           string  `xml:"MarketingPreScreen,omitempty"`
	MarketingSegmentationCluster string  `xml:"MarketingSegmentationCluster,omitempty"`
	MinorityOwned                string  `xml:"MinorityOwned,omitempty"`
	Name                         string  `xml:"Name,omitempty"`
	NationalId                   string  `xml:"NationalId,omitempty"`
	NationalIdType               string  `xml:"NationalIdType,omitempty"`
	OutOfBusiness                string  `xml:"OutOfBusiness,omitempty"`
	OwnOrRent                    string  `xml:"OwnOrRent,omitempty"`
	ParentOrHqBusinessName       string  `xml:"ParentOrHqBusinessName,omitempty"`
	ParentOrHqDunsNumber         string  `xml:"ParentOrHqDunsNumber,omitempty"`
	Phone                        string  `xml:"Phone,omitempty"`
	PremisesMeasure              int32   `xml:"PremisesMeasure,omitempty"`
	PremisesMeasureReliability   string  `xml:"PremisesMeasureReliability,omitempty"`
	PremisesMeasureUnit          string  `xml:"PremisesMeasureUnit,omitempty"`
	PrimaryNaics                 string  `xml:"PrimaryNaics,omitempty"`
	PrimaryNaicsDesc             string  `xml:"PrimaryNaicsDesc,omitempty"`
	PrimarySic                   string  `xml:"PrimarySic,omitempty"`
	PrimarySic8                  string  `xml:"PrimarySic8,omitempty"`
	PrimarySic8Desc              string  `xml:"PrimarySic8Desc,omitempty"`
	PrimarySicDesc               string  `xml:"PrimarySicDesc,omitempty"`
	PriorYearEmployees           int32   `xml:"PriorYearEmployees,omitempty"`
	PriorYearRevenue             float64 `xml:"PriorYearRevenue,omitempty"`
	PublicIndicator              string  `xml:"PublicIndicator,omitempty"`
	Revenue                      float64 `xml:"Revenue,omitempty"`
	SalesTurnoverGrowthRate      float64 `xml:"SalesTurnoverGrowthRate,omitempty"`
	SalesVolume                  float64 `xml:"SalesVolume,omitempty"`
	SalesVolumeReliability       string  `xml:"SalesVolumeReliability,omitempty"`
	SecondNaics                  string  `xml:"SecondNaics,omitempty"`
	SecondNaicsDesc              string  `xml:"SecondNaicsDesc,omitempty"`
	SecondSic                    string  `xml:"SecondSic,omitempty"`
	SecondSic8                   string  `xml:"SecondSic8,omitempty"`
	SecondSic8Desc               string  `xml:"SecondSic8Desc,omitempty"`
	SecondSicDesc                string  `xml:"SecondSicDesc,omitempty"`
	SicCodeDesc                  string  `xml:"SicCodeDesc,omitempty"`
	SixthNaics                   string  `xml:"SixthNaics,omitempty"`
	SixthNaicsDesc               string  `xml:"SixthNaicsDesc,omitempty"`
	SixthSic                     string  `xml:"SixthSic,omitempty"`
	SixthSic8                    string  `xml:"SixthSic8,omitempty"`
	SixthSic8Desc                string  `xml:"SixthSic8Desc,omitempty"`
	SixthSicDesc                 string  `xml:"SixthSicDesc,omitempty"`
	SmallBusiness                string  `xml:"SmallBusiness,omitempty"`
	State                        string  `xml:"State,omitempty"`
	StateCode                    string  `xml:"StateCode,omitempty"`
	StockExchange                string  `xml:"StockExchange,omitempty"`
	StockSymbol                  string  `xml:"StockSymbol,omitempty"`
	Street                       string  `xml:"Street,omitempty"`
	Subsidiary                   string  `xml:"Subsidiary,omitempty"`
	ThirdNaics                   string  `xml:"ThirdNaics,omitempty"`
	ThirdNaicsDesc               string  `xml:"ThirdNaicsDesc,omitempty"`
	ThirdSic                     string  `xml:"ThirdSic,omitempty"`
	ThirdSic8                    string  `xml:"ThirdSic8,omitempty"`
	ThirdSic8Desc                string  `xml:"ThirdSic8Desc,omitempty"`
	ThirdSicDesc                 string  `xml:"ThirdSicDesc,omitempty"`
	TradeStyle1                  string  `xml:"TradeStyle1,omitempty"`
	TradeStyle2                  string  `xml:"TradeStyle2,omitempty"`
	TradeStyle3                  string  `xml:"TradeStyle3,omitempty"`
	TradeStyle4                  string  `xml:"TradeStyle4,omitempty"`
	TradeStyle5                  string  `xml:"TradeStyle5,omitempty"`
	URL                          string  `xml:"URL,omitempty"`
	UsTaxId                      string  `xml:"UsTaxId,omitempty"`
	WomenOwned                   string  `xml:"WomenOwned,omitempty"`
	YearStarted                  string  `xml:"YearStarted,omitempty"`
	Zip                          string  `xml:"Zip,omitempty"`
}

type DatacloudOwnedEntity struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com DatacloudOwnedEntity"`

	*SObject

	CreatedBy           *User                   `xml:"CreatedBy,omitempty"`
	CreatedById         *ID                     `xml:"CreatedById,omitempty"`
	CreatedDate         time.Time               `xml:"CreatedDate,omitempty"`
	DataDotComKey       string                  `xml:"DataDotComKey,omitempty"`
	DatacloudEntityType string                  `xml:"DatacloudEntityType,omitempty"`
	IsDeleted           bool                    `xml:"IsDeleted,omitempty"`
	LastModifiedBy      *User                   `xml:"LastModifiedBy,omitempty"`
	LastModifiedById    *ID                     `xml:"LastModifiedById,omitempty"`
	LastModifiedDate    time.Time               `xml:"LastModifiedDate,omitempty"`
	Name                string                  `xml:"Name,omitempty"`
	PurchaseType        string                  `xml:"PurchaseType,omitempty"`
	PurchaseUsage       *DatacloudPurchaseUsage `xml:"PurchaseUsage,omitempty"`
	PurchaseUsageId     *ID                     `xml:"PurchaseUsageId,omitempty"`
	SystemModstamp      time.Time               `xml:"SystemModstamp,omitempty"`
	User                *User                   `xml:"User,omitempty"`
	UserId              *ID                     `xml:"UserId,omitempty"`
	UserRecordAccess    *UserRecordAccess       `xml:"UserRecordAccess,omitempty"`
}

type DatacloudPurchaseUsage struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com DatacloudPurchaseUsage"`

	*SObject

	CreatedBy           *User             `xml:"CreatedBy,omitempty"`
	CreatedById         *ID               `xml:"CreatedById,omitempty"`
	CreatedDate         time.Time         `xml:"CreatedDate,omitempty"`
	DatacloudEntityType string            `xml:"DatacloudEntityType,omitempty"`
	Description         string            `xml:"Description,omitempty"`
	IsDeleted           bool              `xml:"IsDeleted,omitempty"`
	LastModifiedBy      *User             `xml:"LastModifiedBy,omitempty"`
	LastModifiedById    *ID               `xml:"LastModifiedById,omitempty"`
	LastModifiedDate    time.Time         `xml:"LastModifiedDate,omitempty"`
	Name                string            `xml:"Name,omitempty"`
	PurchaseType        string            `xml:"PurchaseType,omitempty"`
	SystemModstamp      time.Time         `xml:"SystemModstamp,omitempty"`
	Usage               float64           `xml:"Usage,omitempty"`
	User                *User             `xml:"User,omitempty"`
	UserId              *ID               `xml:"UserId,omitempty"`
	UserRecordAccess    *UserRecordAccess `xml:"UserRecordAccess,omitempty"`
	UserType            string            `xml:"UserType,omitempty"`
}

type DatasetExport struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com DatasetExport"`

	*SObject

	CompressedMetadataLength int32        `xml:"CompressedMetadataLength,omitempty"`
	CreatedBy                *User        `xml:"CreatedBy,omitempty"`
	CreatedById              *ID          `xml:"CreatedById,omitempty"`
	CreatedDate              time.Time    `xml:"CreatedDate,omitempty"`
	DatasetExportParts       *QueryResult `xml:"DatasetExportParts,omitempty"`
	IsDeleted                bool         `xml:"IsDeleted,omitempty"`
	LastModifiedBy           *User        `xml:"LastModifiedBy,omitempty"`
	LastModifiedById         *ID          `xml:"LastModifiedById,omitempty"`
	LastModifiedDate         time.Time    `xml:"LastModifiedDate,omitempty"`
	Metadata                 []byte       `xml:"Metadata,omitempty"`
	MetadataLength           int32        `xml:"MetadataLength,omitempty"`
	Owner                    string       `xml:"Owner,omitempty"`
	PublisherInfo            string       `xml:"PublisherInfo,omitempty"`
	PublisherType            string       `xml:"PublisherType,omitempty"`
	Status                   string       `xml:"Status,omitempty"`
	SystemModstamp           time.Time    `xml:"SystemModstamp,omitempty"`
}

type DatasetExportEvent struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com DatasetExportEvent"`

	*SObject

	CreatedBy          *User     `xml:"CreatedBy,omitempty"`
	CreatedById        *ID       `xml:"CreatedById,omitempty"`
	CreatedDate        time.Time `xml:"CreatedDate,omitempty"`
	DataflowInstanceId string    `xml:"DataflowInstanceId,omitempty"`
	DatasetExportId    string    `xml:"DatasetExportId,omitempty"`
	Message            string    `xml:"Message,omitempty"`
	Owner              string    `xml:"Owner,omitempty"`
	PublisherInfo      string    `xml:"PublisherInfo,omitempty"`
	PublisherType      string    `xml:"PublisherType,omitempty"`
	ReplayId           string    `xml:"ReplayId,omitempty"`
	Status             string    `xml:"Status,omitempty"`
}

type DatasetExportPart struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com DatasetExportPart"`

	*SObject

	CompressedDataFileLength int32          `xml:"CompressedDataFileLength,omitempty"`
	CreatedBy                *User          `xml:"CreatedBy,omitempty"`
	CreatedById              *ID            `xml:"CreatedById,omitempty"`
	CreatedDate              time.Time      `xml:"CreatedDate,omitempty"`
	DataFile                 []byte         `xml:"DataFile,omitempty"`
	DataFileLength           int32          `xml:"DataFileLength,omitempty"`
	DatasetExport            *DatasetExport `xml:"DatasetExport,omitempty"`
	DatasetExportId          *ID            `xml:"DatasetExportId,omitempty"`
	IsDeleted                bool           `xml:"IsDeleted,omitempty"`
	LastModifiedBy           *User          `xml:"LastModifiedBy,omitempty"`
	LastModifiedById         *ID            `xml:"LastModifiedById,omitempty"`
	LastModifiedDate         time.Time      `xml:"LastModifiedDate,omitempty"`
	Owner                    string         `xml:"Owner,omitempty"`
	PartNumber               int32          `xml:"PartNumber,omitempty"`
	SystemModstamp           time.Time      `xml:"SystemModstamp,omitempty"`
}

type DeclinedEventRelation struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com DeclinedEventRelation"`

	*SObject

	CreatedBy        *User     `xml:"CreatedBy,omitempty"`
	CreatedById      *ID       `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time `xml:"CreatedDate,omitempty"`
	Event            *Event    `xml:"Event,omitempty"`
	EventId          *ID       `xml:"EventId,omitempty"`
	IsDeleted        bool      `xml:"IsDeleted,omitempty"`
	LastModifiedBy   *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time `xml:"LastModifiedDate,omitempty"`
	Relation         *SObject  `xml:"Relation,omitempty"`
	RelationId       *ID       `xml:"RelationId,omitempty"`
	RespondedDate    time.Time `xml:"RespondedDate,omitempty"`
	Response         string    `xml:"Response,omitempty"`
	SystemModstamp   time.Time `xml:"SystemModstamp,omitempty"`
	Type             string    `xml:"Type,omitempty"`
}

type Document struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com Document"`

	*SObject

	Author             *User     `xml:"Author,omitempty"`
	AuthorId           *ID       `xml:"AuthorId,omitempty"`
	Body               []byte    `xml:"Body,omitempty"`
	BodyLength         int32     `xml:"BodyLength,omitempty"`
	ContentType        string    `xml:"ContentType,omitempty"`
	CreatedBy          *User     `xml:"CreatedBy,omitempty"`
	CreatedById        *ID       `xml:"CreatedById,omitempty"`
	CreatedDate        time.Time `xml:"CreatedDate,omitempty"`
	Description        string    `xml:"Description,omitempty"`
	DeveloperName      string    `xml:"DeveloperName,omitempty"`
	Folder             *Folder   `xml:"Folder,omitempty"`
	FolderId           *ID       `xml:"FolderId,omitempty"`
	IsBodySearchable   bool      `xml:"IsBodySearchable,omitempty"`
	IsDeleted          bool      `xml:"IsDeleted,omitempty"`
	IsInternalUseOnly  bool      `xml:"IsInternalUseOnly,omitempty"`
	IsPublic           bool      `xml:"IsPublic,omitempty"`
	Keywords           string    `xml:"Keywords,omitempty"`
	LastModifiedBy     *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById   *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate   time.Time `xml:"LastModifiedDate,omitempty"`
	LastReferencedDate time.Time `xml:"LastReferencedDate,omitempty"`
	LastViewedDate     time.Time `xml:"LastViewedDate,omitempty"`
	Name               string    `xml:"Name,omitempty"`
	NamespacePrefix    string    `xml:"NamespacePrefix,omitempty"`
	SystemModstamp     time.Time `xml:"SystemModstamp,omitempty"`
	Type               string    `xml:"Type,omitempty"`
	Url                string    `xml:"Url,omitempty"`
}

type DocumentAttachmentMap struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com DocumentAttachmentMap"`

	*SObject

	CreatedBy        *User     `xml:"CreatedBy,omitempty"`
	CreatedById      *ID       `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time `xml:"CreatedDate,omitempty"`
	DocumentId       *ID       `xml:"DocumentId,omitempty"`
	DocumentSequence int32     `xml:"DocumentSequence,omitempty"`
	ParentId         *ID       `xml:"ParentId,omitempty"`
}

type Domain struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com Domain"`

	*SObject

	CnameTarget          string       `xml:"CnameTarget,omitempty"`
	CreatedBy            *User        `xml:"CreatedBy,omitempty"`
	CreatedById          *ID          `xml:"CreatedById,omitempty"`
	CreatedDate          time.Time    `xml:"CreatedDate,omitempty"`
	Domain               string       `xml:"Domain,omitempty"`
	DomainSites          *QueryResult `xml:"DomainSites,omitempty"`
	DomainType           string       `xml:"DomainType,omitempty"`
	LastModifiedBy       *User        `xml:"LastModifiedBy,omitempty"`
	LastModifiedById     *ID          `xml:"LastModifiedById,omitempty"`
	LastModifiedDate     time.Time    `xml:"LastModifiedDate,omitempty"`
	OptionsExternalHttps bool         `xml:"OptionsExternalHttps,omitempty"`
	SystemModstamp       time.Time    `xml:"SystemModstamp,omitempty"`
}

type DomainSite struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com DomainSite"`

	*SObject

	CreatedBy        *User     `xml:"CreatedBy,omitempty"`
	CreatedById      *ID       `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time `xml:"CreatedDate,omitempty"`
	Domain           *Domain   `xml:"Domain,omitempty"`
	DomainId         *ID       `xml:"DomainId,omitempty"`
	LastModifiedBy   *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time `xml:"LastModifiedDate,omitempty"`
	PathPrefix       string    `xml:"PathPrefix,omitempty"`
	Site             *Site     `xml:"Site,omitempty"`
	SiteId           *ID       `xml:"SiteId,omitempty"`
	SystemModstamp   time.Time `xml:"SystemModstamp,omitempty"`
}

type DuplicateRecordItem struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com DuplicateRecordItem"`

	*SObject

	CreatedBy            *User               `xml:"CreatedBy,omitempty"`
	CreatedById          *ID                 `xml:"CreatedById,omitempty"`
	CreatedDate          time.Time           `xml:"CreatedDate,omitempty"`
	DuplicateRecordSet   *DuplicateRecordSet `xml:"DuplicateRecordSet,omitempty"`
	DuplicateRecordSetId *ID                 `xml:"DuplicateRecordSetId,omitempty"`
	IsDeleted            bool                `xml:"IsDeleted,omitempty"`
	LastModifiedBy       *User               `xml:"LastModifiedBy,omitempty"`
	LastModifiedById     *ID                 `xml:"LastModifiedById,omitempty"`
	LastModifiedDate     time.Time           `xml:"LastModifiedDate,omitempty"`
	Name                 string              `xml:"Name,omitempty"`
	ProcessInstances     *QueryResult        `xml:"ProcessInstances,omitempty"`
	ProcessSteps         *QueryResult        `xml:"ProcessSteps,omitempty"`
	Record               *SObject            `xml:"Record,omitempty"`
	RecordId             *ID                 `xml:"RecordId,omitempty"`
	SystemModstamp       time.Time           `xml:"SystemModstamp,omitempty"`
	UserRecordAccess     *UserRecordAccess   `xml:"UserRecordAccess,omitempty"`
}

type DuplicateRecordSet struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com DuplicateRecordSet"`

	*SObject

	CreatedBy            *User             `xml:"CreatedBy,omitempty"`
	CreatedById          *ID               `xml:"CreatedById,omitempty"`
	CreatedDate          time.Time         `xml:"CreatedDate,omitempty"`
	DuplicateRecordItems *QueryResult      `xml:"DuplicateRecordItems,omitempty"`
	DuplicateRule        *DuplicateRule    `xml:"DuplicateRule,omitempty"`
	DuplicateRuleId      *ID               `xml:"DuplicateRuleId,omitempty"`
	IsDeleted            bool              `xml:"IsDeleted,omitempty"`
	LastModifiedBy       *User             `xml:"LastModifiedBy,omitempty"`
	LastModifiedById     *ID               `xml:"LastModifiedById,omitempty"`
	LastModifiedDate     time.Time         `xml:"LastModifiedDate,omitempty"`
	LastReferencedDate   time.Time         `xml:"LastReferencedDate,omitempty"`
	LastViewedDate       time.Time         `xml:"LastViewedDate,omitempty"`
	Name                 string            `xml:"Name,omitempty"`
	ProcessInstances     *QueryResult      `xml:"ProcessInstances,omitempty"`
	ProcessSteps         *QueryResult      `xml:"ProcessSteps,omitempty"`
	RecordCount          int32             `xml:"RecordCount,omitempty"`
	SystemModstamp       time.Time         `xml:"SystemModstamp,omitempty"`
	UserRecordAccess     *UserRecordAccess `xml:"UserRecordAccess,omitempty"`
}

type DuplicateRule struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com DuplicateRule"`

	*SObject

	CreatedBy           *User        `xml:"CreatedBy,omitempty"`
	CreatedById         *ID          `xml:"CreatedById,omitempty"`
	CreatedDate         time.Time    `xml:"CreatedDate,omitempty"`
	DeveloperName       string       `xml:"DeveloperName,omitempty"`
	DuplicateRecordSets *QueryResult `xml:"DuplicateRecordSets,omitempty"`
	IsActive            bool         `xml:"IsActive,omitempty"`
	IsDeleted           bool         `xml:"IsDeleted,omitempty"`
	Language            string       `xml:"Language,omitempty"`
	LastModifiedBy      *User        `xml:"LastModifiedBy,omitempty"`
	LastModifiedById    *ID          `xml:"LastModifiedById,omitempty"`
	LastModifiedDate    time.Time    `xml:"LastModifiedDate,omitempty"`
	LastViewedDate      time.Time    `xml:"LastViewedDate,omitempty"`
	MasterLabel         string       `xml:"MasterLabel,omitempty"`
	NamespacePrefix     string       `xml:"NamespacePrefix,omitempty"`
	SobjectSubtype      string       `xml:"SobjectSubtype,omitempty"`
	SobjectType         string       `xml:"SobjectType,omitempty"`
	SystemModstamp      time.Time    `xml:"SystemModstamp,omitempty"`
}

type EmailCapture struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com EmailCapture"`

	*SObject

	CaptureDate      time.Time `xml:"CaptureDate,omitempty"`
	CreatedBy        *User     `xml:"CreatedBy,omitempty"`
	CreatedById      *ID       `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time `xml:"CreatedDate,omitempty"`
	FromPattern      string    `xml:"FromPattern,omitempty"`
	IsActive         bool      `xml:"IsActive,omitempty"`
	IsDeleted        bool      `xml:"IsDeleted,omitempty"`
	LastModifiedBy   *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time `xml:"LastModifiedDate,omitempty"`
	RawMessage       []byte    `xml:"RawMessage,omitempty"`
	RawMessageLength int32     `xml:"RawMessageLength,omitempty"`
	Recipient        string    `xml:"Recipient,omitempty"`
	Sender           string    `xml:"Sender,omitempty"`
	SystemModstamp   time.Time `xml:"SystemModstamp,omitempty"`
	ToPattern        string    `xml:"ToPattern,omitempty"`
}

type EmailDomainKey struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com EmailDomainKey"`

	*SObject

	CreatedBy              *User     `xml:"CreatedBy,omitempty"`
	CreatedById            *ID       `xml:"CreatedById,omitempty"`
	CreatedDate            time.Time `xml:"CreatedDate,omitempty"`
	Domain                 string    `xml:"Domain,omitempty"`
	DomainMatch            string    `xml:"DomainMatch,omitempty"`
	IsActive               bool      `xml:"IsActive,omitempty"`
	IsDeleted              bool      `xml:"IsDeleted,omitempty"`
	LastModifiedBy         *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById       *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate       time.Time `xml:"LastModifiedDate,omitempty"`
	PrivateKey             string    `xml:"PrivateKey,omitempty"`
	PublicKey              string    `xml:"PublicKey,omitempty"`
	Selector               string    `xml:"Selector,omitempty"`
	SystemModstamp         time.Time `xml:"SystemModstamp,omitempty"`
	TxtRecordsPublishState string    `xml:"TxtRecordsPublishState,omitempty"`
}

type EmailMessage struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com EmailMessage"`

	*SObject

	ActivityId               *ID           `xml:"ActivityId,omitempty"`
	AttachedContentDocuments *QueryResult  `xml:"AttachedContentDocuments,omitempty"`
	Attachments              *QueryResult  `xml:"Attachments,omitempty"`
	BccAddress               string        `xml:"BccAddress,omitempty"`
	BccIds                   []*ID         `xml:"BccIds,omitempty"`
	CcAddress                string        `xml:"CcAddress,omitempty"`
	CcIds                    []*ID         `xml:"CcIds,omitempty"`
	CombinedAttachments      *QueryResult  `xml:"CombinedAttachments,omitempty"`
	ContentDocumentIds       []*ID         `xml:"ContentDocumentIds,omitempty"`
	ContentDocumentLinks     *QueryResult  `xml:"ContentDocumentLinks,omitempty"`
	CreatedBy                *User         `xml:"CreatedBy,omitempty"`
	CreatedById              *ID           `xml:"CreatedById,omitempty"`
	CreatedDate              time.Time     `xml:"CreatedDate,omitempty"`
	EmailMessageRelations    *QueryResult  `xml:"EmailMessageRelations,omitempty"`
	FirstOpenedDate          time.Time     `xml:"FirstOpenedDate,omitempty"`
	FromAddress              string        `xml:"FromAddress,omitempty"`
	FromName                 string        `xml:"FromName,omitempty"`
	HasAttachment            bool          `xml:"HasAttachment,omitempty"`
	Headers                  string        `xml:"Headers,omitempty"`
	HtmlBody                 string        `xml:"HtmlBody,omitempty"`
	Incoming                 bool          `xml:"Incoming,omitempty"`
	IsBounced                bool          `xml:"IsBounced,omitempty"`
	IsClientManaged          bool          `xml:"IsClientManaged,omitempty"`
	IsDeleted                bool          `xml:"IsDeleted,omitempty"`
	IsExternallyVisible      bool          `xml:"IsExternallyVisible,omitempty"`
	IsOpened                 bool          `xml:"IsOpened,omitempty"`
	IsTracked                bool          `xml:"IsTracked,omitempty"`
	LastModifiedBy           *User         `xml:"LastModifiedBy,omitempty"`
	LastModifiedById         *ID           `xml:"LastModifiedById,omitempty"`
	LastModifiedDate         time.Time     `xml:"LastModifiedDate,omitempty"`
	LastOpenedDate           time.Time     `xml:"LastOpenedDate,omitempty"`
	MessageDate              time.Time     `xml:"MessageDate,omitempty"`
	MessageIdentifier        string        `xml:"MessageIdentifier,omitempty"`
	Parent                   *Case         `xml:"Parent,omitempty"`
	ParentId                 *ID           `xml:"ParentId,omitempty"`
	ProcessInstances         *QueryResult  `xml:"ProcessInstances,omitempty"`
	ProcessSteps             *QueryResult  `xml:"ProcessSteps,omitempty"`
	RelatedTo                *SObject      `xml:"RelatedTo,omitempty"`
	RelatedToId              *ID           `xml:"RelatedToId,omitempty"`
	ReplyToEmailMessage      *EmailMessage `xml:"ReplyToEmailMessage,omitempty"`
	ReplyToEmailMessageId    *ID           `xml:"ReplyToEmailMessageId,omitempty"`
	Status                   string        `xml:"Status,omitempty"`
	Subject                  string        `xml:"Subject,omitempty"`
	SystemModstamp           time.Time     `xml:"SystemModstamp,omitempty"`
	TextBody                 string        `xml:"TextBody,omitempty"`
	ThreadIdentifier         string        `xml:"ThreadIdentifier,omitempty"`
	ToAddress                string        `xml:"ToAddress,omitempty"`
	ToIds                    []*ID         `xml:"ToIds,omitempty"`
	ValidatedFromAddress     string        `xml:"ValidatedFromAddress,omitempty"`
}

type EmailMessageRelation struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com EmailMessageRelation"`

	*SObject

	CreatedBy          *User         `xml:"CreatedBy,omitempty"`
	CreatedById        *ID           `xml:"CreatedById,omitempty"`
	CreatedDate        time.Time     `xml:"CreatedDate,omitempty"`
	EmailMessage       *EmailMessage `xml:"EmailMessage,omitempty"`
	EmailMessageId     *ID           `xml:"EmailMessageId,omitempty"`
	IsDeleted          bool          `xml:"IsDeleted,omitempty"`
	Relation           *SObject      `xml:"Relation,omitempty"`
	RelationAddress    string        `xml:"RelationAddress,omitempty"`
	RelationId         *ID           `xml:"RelationId,omitempty"`
	RelationObjectType string        `xml:"RelationObjectType,omitempty"`
	RelationType       string        `xml:"RelationType,omitempty"`
	SystemModstamp     time.Time     `xml:"SystemModstamp,omitempty"`
}

type EmailServicesAddress struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com EmailServicesAddress"`

	*SObject

	AuthorizedSenders string                 `xml:"AuthorizedSenders,omitempty"`
	CreatedBy         *User                  `xml:"CreatedBy,omitempty"`
	CreatedById       *ID                    `xml:"CreatedById,omitempty"`
	CreatedDate       time.Time              `xml:"CreatedDate,omitempty"`
	DeveloperName     string                 `xml:"DeveloperName,omitempty"`
	EmailDomainName   string                 `xml:"EmailDomainName,omitempty"`
	Function          *EmailServicesFunction `xml:"Function,omitempty"`
	FunctionId        *ID                    `xml:"FunctionId,omitempty"`
	IsActive          bool                   `xml:"IsActive,omitempty"`
	LastModifiedBy    *User                  `xml:"LastModifiedBy,omitempty"`
	LastModifiedById  *ID                    `xml:"LastModifiedById,omitempty"`
	LastModifiedDate  time.Time              `xml:"LastModifiedDate,omitempty"`
	LocalPart         string                 `xml:"LocalPart,omitempty"`
	RunAsUserId       *ID                    `xml:"RunAsUserId,omitempty"`
	SystemModstamp    time.Time              `xml:"SystemModstamp,omitempty"`
}

type EmailServicesFunction struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com EmailServicesFunction"`

	*SObject

	AddressInactiveAction       string       `xml:"AddressInactiveAction,omitempty"`
	Addresses                   *QueryResult `xml:"Addresses,omitempty"`
	ApexClassId                 *ID          `xml:"ApexClassId,omitempty"`
	AttachmentOption            string       `xml:"AttachmentOption,omitempty"`
	AuthenticationFailureAction string       `xml:"AuthenticationFailureAction,omitempty"`
	AuthorizationFailureAction  string       `xml:"AuthorizationFailureAction,omitempty"`
	AuthorizedSenders           string       `xml:"AuthorizedSenders,omitempty"`
	CreatedBy                   *User        `xml:"CreatedBy,omitempty"`
	CreatedById                 *ID          `xml:"CreatedById,omitempty"`
	CreatedDate                 time.Time    `xml:"CreatedDate,omitempty"`
	ErrorRoutingAddress         string       `xml:"ErrorRoutingAddress,omitempty"`
	FunctionInactiveAction      string       `xml:"FunctionInactiveAction,omitempty"`
	FunctionName                string       `xml:"FunctionName,omitempty"`
	IsActive                    bool         `xml:"IsActive,omitempty"`
	IsAuthenticationRequired    bool         `xml:"IsAuthenticationRequired,omitempty"`
	IsErrorRoutingEnabled       bool         `xml:"IsErrorRoutingEnabled,omitempty"`
	IsTextAttachmentsAsBinary   bool         `xml:"IsTextAttachmentsAsBinary,omitempty"`
	IsTlsRequired               bool         `xml:"IsTlsRequired,omitempty"`
	LastModifiedBy              *User        `xml:"LastModifiedBy,omitempty"`
	LastModifiedById            *ID          `xml:"LastModifiedById,omitempty"`
	LastModifiedDate            time.Time    `xml:"LastModifiedDate,omitempty"`
	OverLimitAction             string       `xml:"OverLimitAction,omitempty"`
	SystemModstamp              time.Time    `xml:"SystemModstamp,omitempty"`
}

type EmailStatus struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com EmailStatus"`

	*SObject

	CreatedBy         *User     `xml:"CreatedBy,omitempty"`
	CreatedById       *ID       `xml:"CreatedById,omitempty"`
	CreatedDate       time.Time `xml:"CreatedDate,omitempty"`
	EmailTemplateName string    `xml:"EmailTemplateName,omitempty"`
	FirstOpenDate     time.Time `xml:"FirstOpenDate,omitempty"`
	LastModifiedBy    *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById  *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate  time.Time `xml:"LastModifiedDate,omitempty"`
	LastOpenDate      time.Time `xml:"LastOpenDate,omitempty"`
	Task              *Task     `xml:"Task,omitempty"`
	TaskId            *ID       `xml:"TaskId,omitempty"`
	TimesOpened       int32     `xml:"TimesOpened,omitempty"`
	Who               *SObject  `xml:"Who,omitempty"`
	WhoId             *ID       `xml:"WhoId,omitempty"`
}

type EmailTemplate struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com EmailTemplate"`

	*SObject

	ApiVersion               float64      `xml:"ApiVersion,omitempty"`
	AttachedContentDocuments *QueryResult `xml:"AttachedContentDocuments,omitempty"`
	Attachments              *QueryResult `xml:"Attachments,omitempty"`
	Body                     string       `xml:"Body,omitempty"`
	BrandTemplateId          *ID          `xml:"BrandTemplateId,omitempty"`
	CombinedAttachments      *QueryResult `xml:"CombinedAttachments,omitempty"`
	ContentDocumentLinks     *QueryResult `xml:"ContentDocumentLinks,omitempty"`
	CreatedBy                *User        `xml:"CreatedBy,omitempty"`
	CreatedById              *ID          `xml:"CreatedById,omitempty"`
	CreatedDate              time.Time    `xml:"CreatedDate,omitempty"`
	Description              string       `xml:"Description,omitempty"`
	DeveloperName            string       `xml:"DeveloperName,omitempty"`
	Encoding                 string       `xml:"Encoding,omitempty"`
	Folder                   *Folder      `xml:"Folder,omitempty"`
	FolderId                 *ID          `xml:"FolderId,omitempty"`
	HtmlValue                string       `xml:"HtmlValue,omitempty"`
	IsActive                 bool         `xml:"IsActive,omitempty"`
	LastModifiedBy           *User        `xml:"LastModifiedBy,omitempty"`
	LastModifiedById         *ID          `xml:"LastModifiedById,omitempty"`
	LastModifiedDate         time.Time    `xml:"LastModifiedDate,omitempty"`
	LastUsedDate             time.Time    `xml:"LastUsedDate,omitempty"`
	ListEmailTemplates       *QueryResult `xml:"ListEmailTemplates,omitempty"`
	Markup                   string       `xml:"Markup,omitempty"`
	Name                     string       `xml:"Name,omitempty"`
	NamespacePrefix          string       `xml:"NamespacePrefix,omitempty"`
	Owner                    *User        `xml:"Owner,omitempty"`
	OwnerId                  *ID          `xml:"OwnerId,omitempty"`
	RelatedEntityType        string       `xml:"RelatedEntityType,omitempty"`
	Subject                  string       `xml:"Subject,omitempty"`
	SystemModstamp           time.Time    `xml:"SystemModstamp,omitempty"`
	TemplateStyle            string       `xml:"TemplateStyle,omitempty"`
	TemplateType             string       `xml:"TemplateType,omitempty"`
	TimesUsed                int32        `xml:"TimesUsed,omitempty"`
	UiType                   string       `xml:"UiType,omitempty"`
}

type EmbeddedServiceDetail struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com EmbeddedServiceDetail"`

	*SObject

	AvatarImg                     string `xml:"AvatarImg,omitempty"`
	CancelApptBookingFlowName     string `xml:"CancelApptBookingFlowName,omitempty"`
	ContrastInvertedColor         string `xml:"ContrastInvertedColor,omitempty"`
	ContrastPrimaryColor          string `xml:"ContrastPrimaryColor,omitempty"`
	CustomMinimizedComponent      string `xml:"CustomMinimizedComponent,omitempty"`
	CustomPrechatComponent        string `xml:"CustomPrechatComponent,omitempty"`
	DurableId                     string `xml:"DurableId,omitempty"`
	FieldServiceConfirmCardImg    string `xml:"FieldServiceConfirmCardImg,omitempty"`
	FieldServiceHomeImg           string `xml:"FieldServiceHomeImg,omitempty"`
	FieldServiceLogoImg           string `xml:"FieldServiceLogoImg,omitempty"`
	FlowDeveloperName             string `xml:"FlowDeveloperName,omitempty"`
	Font                          string `xml:"Font,omitempty"`
	FontSize                      string `xml:"FontSize,omitempty"`
	HeaderBackgroundImg           string `xml:"HeaderBackgroundImg,omitempty"`
	Height                        int32  `xml:"Height,omitempty"`
	IsFieldServiceEnabled         bool   `xml:"IsFieldServiceEnabled,omitempty"`
	IsLiveAgentEnabled            bool   `xml:"IsLiveAgentEnabled,omitempty"`
	IsOfflineCaseEnabled          bool   `xml:"IsOfflineCaseEnabled,omitempty"`
	IsPrechatEnabled              bool   `xml:"IsPrechatEnabled,omitempty"`
	IsQueuePositionEnabled        bool   `xml:"IsQueuePositionEnabled,omitempty"`
	ModifyApptBookingFlowName     string `xml:"ModifyApptBookingFlowName,omitempty"`
	NavBarColor                   string `xml:"NavBarColor,omitempty"`
	OfflineCaseBackgroundImg      string `xml:"OfflineCaseBackgroundImg,omitempty"`
	PrechatBackgroundImg          string `xml:"PrechatBackgroundImg,omitempty"`
	PrimaryColor                  string `xml:"PrimaryColor,omitempty"`
	SecondaryColor                string `xml:"SecondaryColor,omitempty"`
	ShouldHideAuthDialog          bool   `xml:"ShouldHideAuthDialog,omitempty"`
	ShouldShowExistingAppointment bool   `xml:"ShouldShowExistingAppointment,omitempty"`
	ShouldShowNewAppointment      bool   `xml:"ShouldShowNewAppointment,omitempty"`
	Site                          string `xml:"Site,omitempty"`
	SmallCompanyLogoImg           string `xml:"SmallCompanyLogoImg,omitempty"`
	WaitingStateBackgroundImg     string `xml:"WaitingStateBackgroundImg,omitempty"`
	Width                         int32  `xml:"Width,omitempty"`
}

type EmbeddedServiceLabel struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com EmbeddedServiceLabel"`

	*SObject

	CustomLabelName                    string `xml:"CustomLabelName,omitempty"`
	DurableId                          string `xml:"DurableId,omitempty"`
	EmbeddedServiceConfigDeveloperName string `xml:"EmbeddedServiceConfigDeveloperName,omitempty"`
	LabelKey                           string `xml:"LabelKey,omitempty"`
}

type EntityDefinition struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com EntityDefinition"`

	*SObject

	ChildRelationships           *QueryResult          `xml:"ChildRelationships,omitempty"`
	DataSteward                  *SObject              `xml:"DataSteward,omitempty"`
	DataStewardId                *ID                   `xml:"DataStewardId,omitempty"`
	DefaultCompactLayoutId       string                `xml:"DefaultCompactLayoutId,omitempty"`
	DetailUrl                    string                `xml:"DetailUrl,omitempty"`
	DeveloperName                string                `xml:"DeveloperName,omitempty"`
	DurableId                    string                `xml:"DurableId,omitempty"`
	EditDefinitionUrl            string                `xml:"EditDefinitionUrl,omitempty"`
	EditUrl                      string                `xml:"EditUrl,omitempty"`
	ExternalSharingModel         string                `xml:"ExternalSharingModel,omitempty"`
	Fields                       *QueryResult          `xml:"Fields,omitempty"`
	HasSubtypes                  bool                  `xml:"HasSubtypes,omitempty"`
	HelpSettingPageName          string                `xml:"HelpSettingPageName,omitempty"`
	HelpSettingPageUrl           string                `xml:"HelpSettingPageUrl,omitempty"`
	InternalSharingModel         string                `xml:"InternalSharingModel,omitempty"`
	IsApexTriggerable            bool                  `xml:"IsApexTriggerable,omitempty"`
	IsAutoActivityCaptureEnabled bool                  `xml:"IsAutoActivityCaptureEnabled,omitempty"`
	IsCompactLayoutable          bool                  `xml:"IsCompactLayoutable,omitempty"`
	IsCustomSetting              bool                  `xml:"IsCustomSetting,omitempty"`
	IsCustomizable               bool                  `xml:"IsCustomizable,omitempty"`
	IsDeprecatedAndHidden        bool                  `xml:"IsDeprecatedAndHidden,omitempty"`
	IsEverCreatable              bool                  `xml:"IsEverCreatable,omitempty"`
	IsEverDeletable              bool                  `xml:"IsEverDeletable,omitempty"`
	IsEverUpdatable              bool                  `xml:"IsEverUpdatable,omitempty"`
	IsFeedEnabled                bool                  `xml:"IsFeedEnabled,omitempty"`
	IsIdEnabled                  bool                  `xml:"IsIdEnabled,omitempty"`
	IsLayoutable                 bool                  `xml:"IsLayoutable,omitempty"`
	IsMruEnabled                 bool                  `xml:"IsMruEnabled,omitempty"`
	IsProcessEnabled             bool                  `xml:"IsProcessEnabled,omitempty"`
	IsQueryable                  bool                  `xml:"IsQueryable,omitempty"`
	IsReplicateable              bool                  `xml:"IsReplicateable,omitempty"`
	IsRetrieveable               bool                  `xml:"IsRetrieveable,omitempty"`
	IsSearchLayoutable           bool                  `xml:"IsSearchLayoutable,omitempty"`
	IsSearchable                 bool                  `xml:"IsSearchable,omitempty"`
	IsSubtype                    bool                  `xml:"IsSubtype,omitempty"`
	IsTriggerable                bool                  `xml:"IsTriggerable,omitempty"`
	IsWorkflowEnabled            bool                  `xml:"IsWorkflowEnabled,omitempty"`
	KeyPrefix                    string                `xml:"KeyPrefix,omitempty"`
	Label                        string                `xml:"Label,omitempty"`
	LastModifiedBy               *User                 `xml:"LastModifiedBy,omitempty"`
	LastModifiedById             *ID                   `xml:"LastModifiedById,omitempty"`
	LastModifiedDate             time.Time             `xml:"LastModifiedDate,omitempty"`
	MasterLabel                  string                `xml:"MasterLabel,omitempty"`
	NamespacePrefix              string                `xml:"NamespacePrefix,omitempty"`
	NewUrl                       string                `xml:"NewUrl,omitempty"`
	OwnerChangeOptions           *QueryResult          `xml:"OwnerChangeOptions,omitempty"`
	Particles                    *QueryResult          `xml:"Particles,omitempty"`
	PluralLabel                  string                `xml:"PluralLabel,omitempty"`
	Publisher                    *Publisher            `xml:"Publisher,omitempty"`
	PublisherId                  string                `xml:"PublisherId,omitempty"`
	QualifiedApiName             string                `xml:"QualifiedApiName,omitempty"`
	RecordTypesSupported         *RecordTypesSupported `xml:"RecordTypesSupported,omitempty"`
	RelationshipDomains          *QueryResult          `xml:"RelationshipDomains,omitempty"`
	RunningUserEntityAccess      *UserEntityAccess     `xml:"RunningUserEntityAccess,omitempty"`
	RunningUserEntityAccessId    string                `xml:"RunningUserEntityAccessId,omitempty"`
	SearchLayouts                *QueryResult          `xml:"SearchLayouts,omitempty"`
}

type EntityParticle struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com EntityParticle"`

	*SObject

	ByteLength                 int32                    `xml:"ByteLength,omitempty"`
	DataType                   string                   `xml:"DataType,omitempty"`
	DefaultValueFormula        string                   `xml:"DefaultValueFormula,omitempty"`
	DeveloperName              string                   `xml:"DeveloperName,omitempty"`
	Digits                     int32                    `xml:"Digits,omitempty"`
	DurableId                  string                   `xml:"DurableId,omitempty"`
	EntityDefinition           *EntityDefinition        `xml:"EntityDefinition,omitempty"`
	EntityDefinitionId         string                   `xml:"EntityDefinitionId,omitempty"`
	ExtraTypeInfo              string                   `xml:"ExtraTypeInfo,omitempty"`
	FieldDefinition            *FieldDefinition         `xml:"FieldDefinition,omitempty"`
	FieldDefinitionId          string                   `xml:"FieldDefinitionId,omitempty"`
	InlineHelpText             string                   `xml:"InlineHelpText,omitempty"`
	IsApiFilterable            bool                     `xml:"IsApiFilterable,omitempty"`
	IsApiGroupable             bool                     `xml:"IsApiGroupable,omitempty"`
	IsApiSortable              bool                     `xml:"IsApiSortable,omitempty"`
	IsAutonumber               bool                     `xml:"IsAutonumber,omitempty"`
	IsCalculated               bool                     `xml:"IsCalculated,omitempty"`
	IsCaseSensitive            bool                     `xml:"IsCaseSensitive,omitempty"`
	IsCompactLayoutable        bool                     `xml:"IsCompactLayoutable,omitempty"`
	IsComponent                bool                     `xml:"IsComponent,omitempty"`
	IsCompound                 bool                     `xml:"IsCompound,omitempty"`
	IsCreatable                bool                     `xml:"IsCreatable,omitempty"`
	IsDefaultedOnCreate        bool                     `xml:"IsDefaultedOnCreate,omitempty"`
	IsDependentPicklist        bool                     `xml:"IsDependentPicklist,omitempty"`
	IsDeprecatedAndHidden      bool                     `xml:"IsDeprecatedAndHidden,omitempty"`
	IsDisplayLocationInDecimal bool                     `xml:"IsDisplayLocationInDecimal,omitempty"`
	IsEncrypted                bool                     `xml:"IsEncrypted,omitempty"`
	IsFieldHistoryTracked      bool                     `xml:"IsFieldHistoryTracked,omitempty"`
	IsHighScaleNumber          bool                     `xml:"IsHighScaleNumber,omitempty"`
	IsHtmlFormatted            bool                     `xml:"IsHtmlFormatted,omitempty"`
	IsIdLookup                 bool                     `xml:"IsIdLookup,omitempty"`
	IsLayoutable               bool                     `xml:"IsLayoutable,omitempty"`
	IsListVisible              bool                     `xml:"IsListVisible,omitempty"`
	IsNameField                bool                     `xml:"IsNameField,omitempty"`
	IsNamePointing             bool                     `xml:"IsNamePointing,omitempty"`
	IsNillable                 bool                     `xml:"IsNillable,omitempty"`
	IsPermissionable           bool                     `xml:"IsPermissionable,omitempty"`
	IsUnique                   bool                     `xml:"IsUnique,omitempty"`
	IsUpdatable                bool                     `xml:"IsUpdatable,omitempty"`
	IsWorkflowFilterable       bool                     `xml:"IsWorkflowFilterable,omitempty"`
	IsWriteRequiresMasterRead  bool                     `xml:"IsWriteRequiresMasterRead,omitempty"`
	Label                      string                   `xml:"Label,omitempty"`
	Length                     int32                    `xml:"Length,omitempty"`
	Mask                       string                   `xml:"Mask,omitempty"`
	MaskType                   string                   `xml:"MaskType,omitempty"`
	MasterLabel                string                   `xml:"MasterLabel,omitempty"`
	Name                       string                   `xml:"Name,omitempty"`
	NamespacePrefix            string                   `xml:"NamespacePrefix,omitempty"`
	PicklistValues             *QueryResult             `xml:"PicklistValues,omitempty"`
	Precision                  int32                    `xml:"Precision,omitempty"`
	QualifiedApiName           string                   `xml:"QualifiedApiName,omitempty"`
	ReferenceTargetField       string                   `xml:"ReferenceTargetField,omitempty"`
	ReferenceTo                *RelationshipReferenceTo `xml:"ReferenceTo,omitempty"`
	RelationshipName           string                   `xml:"RelationshipName,omitempty"`
	RelationshipOrder          int32                    `xml:"RelationshipOrder,omitempty"`
	Scale                      int32                    `xml:"Scale,omitempty"`
	ServiceDataType            *DataType                `xml:"ServiceDataType,omitempty"`
	ServiceDataTypeId          string                   `xml:"ServiceDataTypeId,omitempty"`
	ValueType                  *DataType                `xml:"ValueType,omitempty"`
	ValueTypeId                string                   `xml:"ValueTypeId,omitempty"`
}

type EntitySubscription struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com EntitySubscription"`

	*SObject

	CreatedBy    *SObject  `xml:"CreatedBy,omitempty"`
	CreatedById  *ID       `xml:"CreatedById,omitempty"`
	CreatedDate  time.Time `xml:"CreatedDate,omitempty"`
	IsDeleted    bool      `xml:"IsDeleted,omitempty"`
	Parent       *SObject  `xml:"Parent,omitempty"`
	ParentId     *ID       `xml:"ParentId,omitempty"`
	Subscriber   *User     `xml:"Subscriber,omitempty"`
	SubscriberId *ID       `xml:"SubscriberId,omitempty"`
}

type Event struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com Event"`

	*SObject

	AcceptedEventInviteeIds     []*ID        `xml:"AcceptedEventInviteeIds,omitempty"`
	AcceptedEventRelations      *QueryResult `xml:"AcceptedEventRelations,omitempty"`
	Account                     *Account     `xml:"Account,omitempty"`
	AccountId                   *ID          `xml:"AccountId,omitempty"`
	ActivityDate                time.Time    `xml:"ActivityDate,omitempty"`
	ActivityDateTime            time.Time    `xml:"ActivityDateTime,omitempty"`
	AttachedContentDocuments    *QueryResult `xml:"AttachedContentDocuments,omitempty"`
	Attachments                 *QueryResult `xml:"Attachments,omitempty"`
	CombinedAttachments         *QueryResult `xml:"CombinedAttachments,omitempty"`
	ContentDocumentLinks        *QueryResult `xml:"ContentDocumentLinks,omitempty"`
	CreatedBy                   *User        `xml:"CreatedBy,omitempty"`
	CreatedById                 *ID          `xml:"CreatedById,omitempty"`
	CreatedDate                 time.Time    `xml:"CreatedDate,omitempty"`
	DeclinedEventInviteeIds     []*ID        `xml:"DeclinedEventInviteeIds,omitempty"`
	DeclinedEventRelations      *QueryResult `xml:"DeclinedEventRelations,omitempty"`
	Description                 string       `xml:"Description,omitempty"`
	DurationInMinutes           int32        `xml:"DurationInMinutes,omitempty"`
	EndDateTime                 time.Time    `xml:"EndDateTime,omitempty"`
	EventRelations              *QueryResult `xml:"EventRelations,omitempty"`
	EventSubtype                string       `xml:"EventSubtype,omitempty"`
	FeedSubscriptionsForEntity  *QueryResult `xml:"FeedSubscriptionsForEntity,omitempty"`
	Feeds                       *QueryResult `xml:"Feeds,omitempty"`
	GroupEventType              string       `xml:"GroupEventType,omitempty"`
	IsAllDayEvent               bool         `xml:"IsAllDayEvent,omitempty"`
	IsArchived                  bool         `xml:"IsArchived,omitempty"`
	IsChild                     bool         `xml:"IsChild,omitempty"`
	IsDeleted                   bool         `xml:"IsDeleted,omitempty"`
	IsGroupEvent                bool         `xml:"IsGroupEvent,omitempty"`
	IsPrivate                   bool         `xml:"IsPrivate,omitempty"`
	IsRecurrence                bool         `xml:"IsRecurrence,omitempty"`
	IsRecurrence2               bool         `xml:"IsRecurrence2,omitempty"`
	IsRecurrence2Exception      bool         `xml:"IsRecurrence2Exception,omitempty"`
	IsRecurrence2Exclusion      bool         `xml:"IsRecurrence2Exclusion,omitempty"`
	IsReminderSet               bool         `xml:"IsReminderSet,omitempty"`
	LastModifiedBy              *User        `xml:"LastModifiedBy,omitempty"`
	LastModifiedById            *ID          `xml:"LastModifiedById,omitempty"`
	LastModifiedDate            time.Time    `xml:"LastModifiedDate,omitempty"`
	Location                    string       `xml:"Location,omitempty"`
	Owner                       *SObject     `xml:"Owner,omitempty"`
	OwnerId                     *ID          `xml:"OwnerId,omitempty"`
	Recurrence2PatternStartDate time.Time    `xml:"Recurrence2PatternStartDate,omitempty"`
	Recurrence2PatternText      string       `xml:"Recurrence2PatternText,omitempty"`
	Recurrence2PatternTimeZone  string       `xml:"Recurrence2PatternTimeZone,omitempty"`
	Recurrence2PatternVersion   string       `xml:"Recurrence2PatternVersion,omitempty"`
	RecurrenceActivityId        *ID          `xml:"RecurrenceActivityId,omitempty"`
	RecurrenceDayOfMonth        int32        `xml:"RecurrenceDayOfMonth,omitempty"`
	RecurrenceDayOfWeekMask     int32        `xml:"RecurrenceDayOfWeekMask,omitempty"`
	RecurrenceEndDateOnly       time.Time    `xml:"RecurrenceEndDateOnly,omitempty"`
	RecurrenceInstance          string       `xml:"RecurrenceInstance,omitempty"`
	RecurrenceInterval          int32        `xml:"RecurrenceInterval,omitempty"`
	RecurrenceMonthOfYear       string       `xml:"RecurrenceMonthOfYear,omitempty"`
	RecurrenceStartDateTime     time.Time    `xml:"RecurrenceStartDateTime,omitempty"`
	RecurrenceTimeZoneSidKey    string       `xml:"RecurrenceTimeZoneSidKey,omitempty"`
	RecurrenceType              string       `xml:"RecurrenceType,omitempty"`
	RecurringEvents             *QueryResult `xml:"RecurringEvents,omitempty"`
	ReminderDateTime            time.Time    `xml:"ReminderDateTime,omitempty"`
	ShowAs                      string       `xml:"ShowAs,omitempty"`
	StartDateTime               time.Time    `xml:"StartDateTime,omitempty"`
	Subject                     string       `xml:"Subject,omitempty"`
	SystemModstamp              time.Time    `xml:"SystemModstamp,omitempty"`
	TopicAssignments            *QueryResult `xml:"TopicAssignments,omitempty"`
	UndecidedEventInviteeIds    []*ID        `xml:"UndecidedEventInviteeIds,omitempty"`
	UndecidedEventRelations     *QueryResult `xml:"UndecidedEventRelations,omitempty"`
	What                        *SObject     `xml:"What,omitempty"`
	WhatId                      *ID          `xml:"WhatId,omitempty"`
	Who                         *SObject     `xml:"Who,omitempty"`
	WhoId                       *ID          `xml:"WhoId,omitempty"`
}

type EventBusSubscriber struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com EventBusSubscriber"`

	*SObject

	ExternalId string `xml:"ExternalId,omitempty"`
	LastError  string `xml:"LastError,omitempty"`
	Name       string `xml:"Name,omitempty"`
	Position   int32  `xml:"Position,omitempty"`
	Retries    int32  `xml:"Retries,omitempty"`
	Status     string `xml:"Status,omitempty"`
	Tip        int32  `xml:"Tip,omitempty"`
	Topic      string `xml:"Topic,omitempty"`
	Type       string `xml:"Type,omitempty"`
}

type EventFeed struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com EventFeed"`

	*SObject

	BestComment        *FeedComment `xml:"BestComment,omitempty"`
	BestCommentId      *ID          `xml:"BestCommentId,omitempty"`
	Body               string       `xml:"Body,omitempty"`
	CommentCount       int32        `xml:"CommentCount,omitempty"`
	CreatedBy          *SObject     `xml:"CreatedBy,omitempty"`
	CreatedById        *ID          `xml:"CreatedById,omitempty"`
	CreatedDate        time.Time    `xml:"CreatedDate,omitempty"`
	FeedAttachments    *QueryResult `xml:"FeedAttachments,omitempty"`
	FeedComments       *QueryResult `xml:"FeedComments,omitempty"`
	FeedLikes          *QueryResult `xml:"FeedLikes,omitempty"`
	FeedSignals        *QueryResult `xml:"FeedSignals,omitempty"`
	FeedTrackedChanges *QueryResult `xml:"FeedTrackedChanges,omitempty"`
	InsertedBy         *SObject     `xml:"InsertedBy,omitempty"`
	InsertedById       *ID          `xml:"InsertedById,omitempty"`
	IsDeleted          bool         `xml:"IsDeleted,omitempty"`
	IsRichText         bool         `xml:"IsRichText,omitempty"`
	LastModifiedDate   time.Time    `xml:"LastModifiedDate,omitempty"`
	LikeCount          int32        `xml:"LikeCount,omitempty"`
	LinkUrl            string       `xml:"LinkUrl,omitempty"`
	Parent             *Event       `xml:"Parent,omitempty"`
	ParentId           *ID          `xml:"ParentId,omitempty"`
	RelatedRecordId    *ID          `xml:"RelatedRecordId,omitempty"`
	SystemModstamp     time.Time    `xml:"SystemModstamp,omitempty"`
	Title              string       `xml:"Title,omitempty"`
	Type               string       `xml:"Type,omitempty"`
}

type EventLogFile struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com EventLogFile"`

	*SObject

	ApiVersion         float64   `xml:"ApiVersion,omitempty"`
	CreatedBy          *User     `xml:"CreatedBy,omitempty"`
	CreatedById        *ID       `xml:"CreatedById,omitempty"`
	CreatedDate        time.Time `xml:"CreatedDate,omitempty"`
	EventType          string    `xml:"EventType,omitempty"`
	Interval           string    `xml:"Interval,omitempty"`
	IsDeleted          bool      `xml:"IsDeleted,omitempty"`
	LastModifiedBy     *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById   *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate   time.Time `xml:"LastModifiedDate,omitempty"`
	LogDate            time.Time `xml:"LogDate,omitempty"`
	LogFile            []byte    `xml:"LogFile,omitempty"`
	LogFileContentType string    `xml:"LogFileContentType,omitempty"`
	LogFileFieldNames  string    `xml:"LogFileFieldNames,omitempty"`
	LogFileFieldTypes  string    `xml:"LogFileFieldTypes,omitempty"`
	LogFileLength      float64   `xml:"LogFileLength,omitempty"`
	Sequence           int32     `xml:"Sequence,omitempty"`
	SystemModstamp     time.Time `xml:"SystemModstamp,omitempty"`
}

type EventRelation struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com EventRelation"`

	*SObject

	CreatedBy        *User     `xml:"CreatedBy,omitempty"`
	CreatedById      *ID       `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time `xml:"CreatedDate,omitempty"`
	Event            *Event    `xml:"Event,omitempty"`
	EventId          *ID       `xml:"EventId,omitempty"`
	IsDeleted        bool      `xml:"IsDeleted,omitempty"`
	LastModifiedBy   *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time `xml:"LastModifiedDate,omitempty"`
	Relation         *SObject  `xml:"Relation,omitempty"`
	RelationId       *ID       `xml:"RelationId,omitempty"`
	RespondedDate    time.Time `xml:"RespondedDate,omitempty"`
	Response         string    `xml:"Response,omitempty"`
	Status           string    `xml:"Status,omitempty"`
	SystemModstamp   time.Time `xml:"SystemModstamp,omitempty"`
}

type ExternalDataSource struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ExternalDataSource"`

	*SObject

	AuthProvider           *AuthProvider   `xml:"AuthProvider,omitempty"`
	AuthProviderId         *ID             `xml:"AuthProviderId,omitempty"`
	CreatedBy              *User           `xml:"CreatedBy,omitempty"`
	CreatedById            *ID             `xml:"CreatedById,omitempty"`
	CreatedDate            time.Time       `xml:"CreatedDate,omitempty"`
	CustomConfiguration    string          `xml:"CustomConfiguration,omitempty"`
	CustomHttpHeaders      *QueryResult    `xml:"CustomHttpHeaders,omitempty"`
	DeveloperName          string          `xml:"DeveloperName,omitempty"`
	Endpoint               string          `xml:"Endpoint,omitempty"`
	IsDeleted              bool            `xml:"IsDeleted,omitempty"`
	IsWritable             bool            `xml:"IsWritable,omitempty"`
	Language               string          `xml:"Language,omitempty"`
	LargeIcon              *StaticResource `xml:"LargeIcon,omitempty"`
	LargeIconId            *ID             `xml:"LargeIconId,omitempty"`
	LastModifiedBy         *User           `xml:"LastModifiedBy,omitempty"`
	LastModifiedById       *ID             `xml:"LastModifiedById,omitempty"`
	LastModifiedDate       time.Time       `xml:"LastModifiedDate,omitempty"`
	MasterLabel            string          `xml:"MasterLabel,omitempty"`
	NamespacePrefix        string          `xml:"NamespacePrefix,omitempty"`
	PrincipalType          string          `xml:"PrincipalType,omitempty"`
	Protocol               string          `xml:"Protocol,omitempty"`
	Repository             string          `xml:"Repository,omitempty"`
	SetupEntityAccessItems *QueryResult    `xml:"SetupEntityAccessItems,omitempty"`
	SmallIcon              *StaticResource `xml:"SmallIcon,omitempty"`
	SmallIconId            *ID             `xml:"SmallIconId,omitempty"`
	SystemModstamp         time.Time       `xml:"SystemModstamp,omitempty"`
	Type                   string          `xml:"Type,omitempty"`
	UserAuths              *QueryResult    `xml:"UserAuths,omitempty"`
}

type ExternalDataUserAuth struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ExternalDataUserAuth"`

	*SObject

	AuthProvider         *AuthProvider `xml:"AuthProvider,omitempty"`
	AuthProviderId       *ID           `xml:"AuthProviderId,omitempty"`
	CreatedBy            *User         `xml:"CreatedBy,omitempty"`
	CreatedById          *ID           `xml:"CreatedById,omitempty"`
	CreatedDate          time.Time     `xml:"CreatedDate,omitempty"`
	ExternalDataSource   *SObject      `xml:"ExternalDataSource,omitempty"`
	ExternalDataSourceId *ID           `xml:"ExternalDataSourceId,omitempty"`
	IsDeleted            bool          `xml:"IsDeleted,omitempty"`
	LastModifiedBy       *User         `xml:"LastModifiedBy,omitempty"`
	LastModifiedById     *ID           `xml:"LastModifiedById,omitempty"`
	LastModifiedDate     time.Time     `xml:"LastModifiedDate,omitempty"`
	Password             string        `xml:"Password,omitempty"`
	Protocol             string        `xml:"Protocol,omitempty"`
	SystemModstamp       time.Time     `xml:"SystemModstamp,omitempty"`
	User                 *User         `xml:"User,omitempty"`
	UserId               *ID           `xml:"UserId,omitempty"`
	Username             string        `xml:"Username,omitempty"`
}

type FeedAttachment struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com FeedAttachment"`

	*SObject

	FeedEntityId *ID    `xml:"FeedEntityId,omitempty"`
	IsDeleted    bool   `xml:"IsDeleted,omitempty"`
	RecordId     *ID    `xml:"RecordId,omitempty"`
	Title        string `xml:"Title,omitempty"`
	Type         string `xml:"Type,omitempty"`
	Value        string `xml:"Value,omitempty"`
}

type FeedComment struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com FeedComment"`

	*SObject

	CommentBody           string       `xml:"CommentBody,omitempty"`
	CommentType           string       `xml:"CommentType,omitempty"`
	CreatedBy             *SObject     `xml:"CreatedBy,omitempty"`
	CreatedById           *ID          `xml:"CreatedById,omitempty"`
	CreatedDate           time.Time    `xml:"CreatedDate,omitempty"`
	FeedAttachments       *QueryResult `xml:"FeedAttachments,omitempty"`
	FeedItemId            *ID          `xml:"FeedItemId,omitempty"`
	FeedRevisions         *QueryResult `xml:"FeedRevisions,omitempty"`
	FeedThreadedComments  *QueryResult `xml:"FeedThreadedComments,omitempty"`
	HasEntityLinks        bool         `xml:"HasEntityLinks,omitempty"`
	InsertedBy            *User        `xml:"InsertedBy,omitempty"`
	InsertedById          *ID          `xml:"InsertedById,omitempty"`
	IsDeleted             bool         `xml:"IsDeleted,omitempty"`
	IsRichText            bool         `xml:"IsRichText,omitempty"`
	IsVerified            bool         `xml:"IsVerified,omitempty"`
	LastEditById          *ID          `xml:"LastEditById,omitempty"`
	LastEditDate          time.Time    `xml:"LastEditDate,omitempty"`
	ParentId              *ID          `xml:"ParentId,omitempty"`
	RelatedRecordId       *ID          `xml:"RelatedRecordId,omitempty"`
	Revision              int32        `xml:"Revision,omitempty"`
	Status                string       `xml:"Status,omitempty"`
	SystemModstamp        time.Time    `xml:"SystemModstamp,omitempty"`
	ThreadChildrenCount   int32        `xml:"ThreadChildrenCount,omitempty"`
	ThreadLastUpdatedDate time.Time    `xml:"ThreadLastUpdatedDate,omitempty"`
	ThreadLevel           int32        `xml:"ThreadLevel,omitempty"`
	ThreadParent          *FeedComment `xml:"ThreadParent,omitempty"`
	ThreadParentId        *ID          `xml:"ThreadParentId,omitempty"`
}

type FeedItem struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com FeedItem"`

	*SObject

	BestComment        *FeedComment `xml:"BestComment,omitempty"`
	BestCommentId      *ID          `xml:"BestCommentId,omitempty"`
	Body               string       `xml:"Body,omitempty"`
	CommentCount       int32        `xml:"CommentCount,omitempty"`
	CreatedBy          *SObject     `xml:"CreatedBy,omitempty"`
	CreatedById        *ID          `xml:"CreatedById,omitempty"`
	CreatedDate        time.Time    `xml:"CreatedDate,omitempty"`
	FeedAttachments    *QueryResult `xml:"FeedAttachments,omitempty"`
	FeedComments       *QueryResult `xml:"FeedComments,omitempty"`
	FeedLikes          *QueryResult `xml:"FeedLikes,omitempty"`
	FeedRevisions      *QueryResult `xml:"FeedRevisions,omitempty"`
	FeedSignals        *QueryResult `xml:"FeedSignals,omitempty"`
	FeedTrackedChanges *QueryResult `xml:"FeedTrackedChanges,omitempty"`
	HasContent         bool         `xml:"HasContent,omitempty"`
	HasFeedEntity      bool         `xml:"HasFeedEntity,omitempty"`
	HasLink            bool         `xml:"HasLink,omitempty"`
	HasVerifiedComment bool         `xml:"HasVerifiedComment,omitempty"`
	InsertedBy         *SObject     `xml:"InsertedBy,omitempty"`
	InsertedById       *ID          `xml:"InsertedById,omitempty"`
	IsClosed           bool         `xml:"IsClosed,omitempty"`
	IsDeleted          bool         `xml:"IsDeleted,omitempty"`
	IsRichText         bool         `xml:"IsRichText,omitempty"`
	LastEditById       *ID          `xml:"LastEditById,omitempty"`
	LastEditDate       time.Time    `xml:"LastEditDate,omitempty"`
	LastModifiedDate   time.Time    `xml:"LastModifiedDate,omitempty"`
	LikeCount          int32        `xml:"LikeCount,omitempty"`
	LinkUrl            string       `xml:"LinkUrl,omitempty"`
	Parent             *SObject     `xml:"Parent,omitempty"`
	ParentId           *ID          `xml:"ParentId,omitempty"`
	RelatedRecordId    *ID          `xml:"RelatedRecordId,omitempty"`
	Revision           int32        `xml:"Revision,omitempty"`
	Status             string       `xml:"Status,omitempty"`
	SystemModstamp     time.Time    `xml:"SystemModstamp,omitempty"`
	Title              string       `xml:"Title,omitempty"`
	TopicAssignments   *QueryResult `xml:"TopicAssignments,omitempty"`
	Type               string       `xml:"Type,omitempty"`
}

type FeedLike struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com FeedLike"`

	*SObject

	CreatedBy    *User     `xml:"CreatedBy,omitempty"`
	CreatedById  *ID       `xml:"CreatedById,omitempty"`
	CreatedDate  time.Time `xml:"CreatedDate,omitempty"`
	FeedEntityId *ID       `xml:"FeedEntityId,omitempty"`
	FeedItemId   *ID       `xml:"FeedItemId,omitempty"`
	InsertedBy   *User     `xml:"InsertedBy,omitempty"`
	InsertedById *ID       `xml:"InsertedById,omitempty"`
	IsDeleted    bool      `xml:"IsDeleted,omitempty"`
}

type FeedPollChoice struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com FeedPollChoice"`

	*SObject

	ChoiceBody    string       `xml:"ChoiceBody,omitempty"`
	CreatedBy     *User        `xml:"CreatedBy,omitempty"`
	CreatedById   *ID          `xml:"CreatedById,omitempty"`
	CreatedDate   time.Time    `xml:"CreatedDate,omitempty"`
	FeedItemId    *ID          `xml:"FeedItemId,omitempty"`
	FeedPollVotes *QueryResult `xml:"FeedPollVotes,omitempty"`
	IsDeleted     bool         `xml:"IsDeleted,omitempty"`
	Position      int32        `xml:"Position,omitempty"`
}

type FeedPollVote struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com FeedPollVote"`

	*SObject

	Choice           *FeedPollChoice `xml:"Choice,omitempty"`
	ChoiceId         *ID             `xml:"ChoiceId,omitempty"`
	CreatedBy        *User           `xml:"CreatedBy,omitempty"`
	CreatedById      *ID             `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time       `xml:"CreatedDate,omitempty"`
	FeedItemId       *ID             `xml:"FeedItemId,omitempty"`
	IsDeleted        bool            `xml:"IsDeleted,omitempty"`
	LastModifiedDate time.Time       `xml:"LastModifiedDate,omitempty"`
}

type FeedRevision struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com FeedRevision"`

	*SObject

	Action          string    `xml:"Action,omitempty"`
	CreatedBy       *SObject  `xml:"CreatedBy,omitempty"`
	CreatedById     *ID       `xml:"CreatedById,omitempty"`
	CreatedDate     time.Time `xml:"CreatedDate,omitempty"`
	EditedAttribute string    `xml:"EditedAttribute,omitempty"`
	FeedEntityId    *ID       `xml:"FeedEntityId,omitempty"`
	IsDeleted       bool      `xml:"IsDeleted,omitempty"`
	IsValueRichText bool      `xml:"IsValueRichText,omitempty"`
	Revision        int32     `xml:"Revision,omitempty"`
	SystemModstamp  time.Time `xml:"SystemModstamp,omitempty"`
	Value           string    `xml:"Value,omitempty"`
}

type FeedSignal struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com FeedSignal"`

	*SObject

	CreatedBy    *User     `xml:"CreatedBy,omitempty"`
	CreatedById  *ID       `xml:"CreatedById,omitempty"`
	CreatedDate  time.Time `xml:"CreatedDate,omitempty"`
	FeedEntityId *ID       `xml:"FeedEntityId,omitempty"`
	FeedItemId   *ID       `xml:"FeedItemId,omitempty"`
	InsertedBy   *User     `xml:"InsertedBy,omitempty"`
	InsertedById *ID       `xml:"InsertedById,omitempty"`
	IsDeleted    bool      `xml:"IsDeleted,omitempty"`
	SignalType   string    `xml:"SignalType,omitempty"`
	SignalValue  int32     `xml:"SignalValue,omitempty"`
}

type FeedTrackedChange struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com FeedTrackedChange"`

	*SObject

	FeedItemId *ID         `xml:"FeedItemId,omitempty"`
	FieldName  string      `xml:"FieldName,omitempty"`
	NewValue   interface{} `xml:"NewValue,omitempty"`
	OldValue   interface{} `xml:"OldValue,omitempty"`
}

type FieldDefinition struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com FieldDefinition"`

	*SObject

	BusinessOwner                *SObject                 `xml:"BusinessOwner,omitempty"`
	BusinessOwnerId              *ID                      `xml:"BusinessOwnerId,omitempty"`
	BusinessStatus               string                   `xml:"BusinessStatus,omitempty"`
	ControlledFields             *QueryResult             `xml:"ControlledFields,omitempty"`
	ControllingFieldDefinition   *FieldDefinition         `xml:"ControllingFieldDefinition,omitempty"`
	ControllingFieldDefinitionId string                   `xml:"ControllingFieldDefinitionId,omitempty"`
	DataType                     string                   `xml:"DataType,omitempty"`
	Description                  string                   `xml:"Description,omitempty"`
	DeveloperName                string                   `xml:"DeveloperName,omitempty"`
	DurableId                    string                   `xml:"DurableId,omitempty"`
	EntityDefinition             *EntityDefinition        `xml:"EntityDefinition,omitempty"`
	EntityDefinitionId           string                   `xml:"EntityDefinitionId,omitempty"`
	ExtraTypeInfo                string                   `xml:"ExtraTypeInfo,omitempty"`
	IsApiFilterable              bool                     `xml:"IsApiFilterable,omitempty"`
	IsApiGroupable               bool                     `xml:"IsApiGroupable,omitempty"`
	IsApiSortable                bool                     `xml:"IsApiSortable,omitempty"`
	IsCalculated                 bool                     `xml:"IsCalculated,omitempty"`
	IsCompactLayoutable          bool                     `xml:"IsCompactLayoutable,omitempty"`
	IsCompound                   bool                     `xml:"IsCompound,omitempty"`
	IsFieldHistoryTracked        bool                     `xml:"IsFieldHistoryTracked,omitempty"`
	IsHighScaleNumber            bool                     `xml:"IsHighScaleNumber,omitempty"`
	IsHtmlFormatted              bool                     `xml:"IsHtmlFormatted,omitempty"`
	IsIndexed                    bool                     `xml:"IsIndexed,omitempty"`
	IsListFilterable             bool                     `xml:"IsListFilterable,omitempty"`
	IsListSortable               bool                     `xml:"IsListSortable,omitempty"`
	IsListVisible                bool                     `xml:"IsListVisible,omitempty"`
	IsNameField                  bool                     `xml:"IsNameField,omitempty"`
	IsNillable                   bool                     `xml:"IsNillable,omitempty"`
	IsPolymorphicForeignKey      bool                     `xml:"IsPolymorphicForeignKey,omitempty"`
	IsSearchPrefilterable        bool                     `xml:"IsSearchPrefilterable,omitempty"`
	IsWorkflowFilterable         bool                     `xml:"IsWorkflowFilterable,omitempty"`
	Label                        string                   `xml:"Label,omitempty"`
	LastModifiedBy               *User                    `xml:"LastModifiedBy,omitempty"`
	LastModifiedById             *ID                      `xml:"LastModifiedById,omitempty"`
	LastModifiedDate             time.Time                `xml:"LastModifiedDate,omitempty"`
	Length                       int32                    `xml:"Length,omitempty"`
	MasterLabel                  string                   `xml:"MasterLabel,omitempty"`
	NamespacePrefix              string                   `xml:"NamespacePrefix,omitempty"`
	Particles                    *QueryResult             `xml:"Particles,omitempty"`
	Precision                    int32                    `xml:"Precision,omitempty"`
	Publisher                    *Publisher               `xml:"Publisher,omitempty"`
	PublisherId                  string                   `xml:"PublisherId,omitempty"`
	QualifiedApiName             string                   `xml:"QualifiedApiName,omitempty"`
	ReferenceTargetField         string                   `xml:"ReferenceTargetField,omitempty"`
	ReferenceTo                  *RelationshipReferenceTo `xml:"ReferenceTo,omitempty"`
	RelationshipDomains          *QueryResult             `xml:"RelationshipDomains,omitempty"`
	RelationshipName             string                   `xml:"RelationshipName,omitempty"`
	RunningUserFieldAccess       *UserFieldAccess         `xml:"RunningUserFieldAccess,omitempty"`
	RunningUserFieldAccessId     string                   `xml:"RunningUserFieldAccessId,omitempty"`
	Scale                        int32                    `xml:"Scale,omitempty"`
	SecurityClassification       string                   `xml:"SecurityClassification,omitempty"`
	ServiceDataType              *DataType                `xml:"ServiceDataType,omitempty"`
	ServiceDataTypeId            string                   `xml:"ServiceDataTypeId,omitempty"`
	ValueType                    *DataType                `xml:"ValueType,omitempty"`
	ValueTypeId                  string                   `xml:"ValueTypeId,omitempty"`
}

type FieldPermissions struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com FieldPermissions"`

	*SObject

	Field           string         `xml:"Field,omitempty"`
	Parent          *PermissionSet `xml:"Parent,omitempty"`
	ParentId        *ID            `xml:"ParentId,omitempty"`
	PermissionsEdit bool           `xml:"PermissionsEdit,omitempty"`
	PermissionsRead bool           `xml:"PermissionsRead,omitempty"`
	SobjectType     string         `xml:"SobjectType,omitempty"`
	SystemModstamp  time.Time      `xml:"SystemModstamp,omitempty"`
}

type FileSearchActivity struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com FileSearchActivity"`

	*SObject

	AvgNumResults    float64           `xml:"AvgNumResults,omitempty"`
	ClickRank        float64           `xml:"ClickRank,omitempty"`
	CountQueries     int32             `xml:"CountQueries,omitempty"`
	CountUsers       int32             `xml:"CountUsers,omitempty"`
	CreatedBy        *User             `xml:"CreatedBy,omitempty"`
	CreatedById      *ID               `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time         `xml:"CreatedDate,omitempty"`
	IsDeleted        bool              `xml:"IsDeleted,omitempty"`
	LastModifiedBy   *User             `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID               `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time         `xml:"LastModifiedDate,omitempty"`
	Name             string            `xml:"Name,omitempty"`
	Period           string            `xml:"Period,omitempty"`
	QueryDate        time.Time         `xml:"QueryDate,omitempty"`
	QueryLanguage    string            `xml:"QueryLanguage,omitempty"`
	SearchTerm       string            `xml:"SearchTerm,omitempty"`
	SystemModstamp   time.Time         `xml:"SystemModstamp,omitempty"`
	UserRecordAccess *UserRecordAccess `xml:"UserRecordAccess,omitempty"`
}

type FiscalYearSettings struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com FiscalYearSettings"`

	*SObject

	Description        string       `xml:"Description,omitempty"`
	EndDate            time.Time    `xml:"EndDate,omitempty"`
	IsStandardYear     bool         `xml:"IsStandardYear,omitempty"`
	Name               string       `xml:"Name,omitempty"`
	PeriodId           *ID          `xml:"PeriodId,omitempty"`
	PeriodLabelScheme  string       `xml:"PeriodLabelScheme,omitempty"`
	PeriodPrefix       string       `xml:"PeriodPrefix,omitempty"`
	Periods            *QueryResult `xml:"Periods,omitempty"`
	QuarterLabelScheme string       `xml:"QuarterLabelScheme,omitempty"`
	QuarterPrefix      string       `xml:"QuarterPrefix,omitempty"`
	StartDate          time.Time    `xml:"StartDate,omitempty"`
	SystemModstamp     time.Time    `xml:"SystemModstamp,omitempty"`
	WeekLabelScheme    string       `xml:"WeekLabelScheme,omitempty"`
	WeekStartDay       int32        `xml:"WeekStartDay,omitempty"`
	YearType           string       `xml:"YearType,omitempty"`
}

type FlexQueueItem struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com FlexQueueItem"`

	*SObject

	AsyncApexJob    *AsyncApexJob `xml:"AsyncApexJob,omitempty"`
	AsyncApexJobId  *ID           `xml:"AsyncApexJobId,omitempty"`
	FlexQueueItemId string        `xml:"FlexQueueItemId,omitempty"`
	JobPosition     int32         `xml:"JobPosition,omitempty"`
	JobType         string        `xml:"JobType,omitempty"`
}

type FlowInterview struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com FlowInterview"`

	*SObject

	CreatedBy        *User             `xml:"CreatedBy,omitempty"`
	CreatedById      *ID               `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time         `xml:"CreatedDate,omitempty"`
	CurrentElement   string            `xml:"CurrentElement,omitempty"`
	Guid             string            `xml:"Guid,omitempty"`
	InterviewLabel   string            `xml:"InterviewLabel,omitempty"`
	IsDeleted        bool              `xml:"IsDeleted,omitempty"`
	LastModifiedBy   *User             `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID               `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time         `xml:"LastModifiedDate,omitempty"`
	Name             string            `xml:"Name,omitempty"`
	Owner            *SObject          `xml:"Owner,omitempty"`
	OwnerId          *ID               `xml:"OwnerId,omitempty"`
	PauseLabel       string            `xml:"PauseLabel,omitempty"`
	RecordActions    *QueryResult      `xml:"RecordActions,omitempty"`
	RecordRelations  *QueryResult      `xml:"RecordRelations,omitempty"`
	StageRelations   *QueryResult      `xml:"StageRelations,omitempty"`
	SystemModstamp   time.Time         `xml:"SystemModstamp,omitempty"`
	UserRecordAccess *UserRecordAccess `xml:"UserRecordAccess,omitempty"`
}

type FlowInterviewShare struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com FlowInterviewShare"`

	*SObject

	AccessLevel      string         `xml:"AccessLevel,omitempty"`
	IsDeleted        bool           `xml:"IsDeleted,omitempty"`
	LastModifiedBy   *User          `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID            `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time      `xml:"LastModifiedDate,omitempty"`
	Parent           *FlowInterview `xml:"Parent,omitempty"`
	ParentId         *ID            `xml:"ParentId,omitempty"`
	RowCause         string         `xml:"RowCause,omitempty"`
	UserOrGroup      *SObject       `xml:"UserOrGroup,omitempty"`
	UserOrGroupId    *ID            `xml:"UserOrGroupId,omitempty"`
}

type FlowRecordRelation struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com FlowRecordRelation"`

	*SObject

	CreatedBy        *User             `xml:"CreatedBy,omitempty"`
	CreatedById      *ID               `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time         `xml:"CreatedDate,omitempty"`
	IsDeleted        bool              `xml:"IsDeleted,omitempty"`
	LastModifiedBy   *User             `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID               `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time         `xml:"LastModifiedDate,omitempty"`
	Name             string            `xml:"Name,omitempty"`
	Parent           *FlowInterview    `xml:"Parent,omitempty"`
	ParentId         *ID               `xml:"ParentId,omitempty"`
	RelatedRecord    *Contract         `xml:"RelatedRecord,omitempty"`
	RelatedRecordId  *ID               `xml:"RelatedRecordId,omitempty"`
	SystemModstamp   time.Time         `xml:"SystemModstamp,omitempty"`
	UserRecordAccess *UserRecordAccess `xml:"UserRecordAccess,omitempty"`
}

type FlowStageRelation struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com FlowStageRelation"`

	*SObject

	CreatedBy        *User             `xml:"CreatedBy,omitempty"`
	CreatedById      *ID               `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time         `xml:"CreatedDate,omitempty"`
	FlexIndex        string            `xml:"FlexIndex,omitempty"`
	IsDeleted        bool              `xml:"IsDeleted,omitempty"`
	LastModifiedBy   *User             `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID               `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time         `xml:"LastModifiedDate,omitempty"`
	Name             string            `xml:"Name,omitempty"`
	Parent           *FlowInterview    `xml:"Parent,omitempty"`
	ParentId         *ID               `xml:"ParentId,omitempty"`
	StageLabel       string            `xml:"StageLabel,omitempty"`
	StageOrder       int32             `xml:"StageOrder,omitempty"`
	StageType        string            `xml:"StageType,omitempty"`
	SystemModstamp   time.Time         `xml:"SystemModstamp,omitempty"`
	UserRecordAccess *UserRecordAccess `xml:"UserRecordAccess,omitempty"`
}

type Folder struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com Folder"`

	*SObject

	AccessType       string            `xml:"AccessType,omitempty"`
	CreatedBy        *User             `xml:"CreatedBy,omitempty"`
	CreatedById      *ID               `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time         `xml:"CreatedDate,omitempty"`
	DeveloperName    string            `xml:"DeveloperName,omitempty"`
	IsReadonly       bool              `xml:"IsReadonly,omitempty"`
	LastModifiedBy   *User             `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID               `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time         `xml:"LastModifiedDate,omitempty"`
	Name             string            `xml:"Name,omitempty"`
	NamespacePrefix  string            `xml:"NamespacePrefix,omitempty"`
	ParentId         *ID               `xml:"ParentId,omitempty"`
	SubFolders       *QueryResult      `xml:"SubFolders,omitempty"`
	SystemModstamp   time.Time         `xml:"SystemModstamp,omitempty"`
	Type             string            `xml:"Type,omitempty"`
	UserRecordAccess *UserRecordAccess `xml:"UserRecordAccess,omitempty"`
}

type FolderedContentDocument struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com FolderedContentDocument"`

	*SObject

	ContentDocument       *ContentDocument `xml:"ContentDocument,omitempty"`
	ContentDocumentId     *ID              `xml:"ContentDocumentId,omitempty"`
	ContentSize           int32            `xml:"ContentSize,omitempty"`
	CreatedBy             *User            `xml:"CreatedBy,omitempty"`
	CreatedById           *ID              `xml:"CreatedById,omitempty"`
	CreatedDate           time.Time        `xml:"CreatedDate,omitempty"`
	FileExtension         string           `xml:"FileExtension,omitempty"`
	FileType              string           `xml:"FileType,omitempty"`
	IsDeleted             bool             `xml:"IsDeleted,omitempty"`
	IsFolder              bool             `xml:"IsFolder,omitempty"`
	LastModifiedBy        *User            `xml:"LastModifiedBy,omitempty"`
	LastModifiedById      *ID              `xml:"LastModifiedById,omitempty"`
	LastModifiedDate      time.Time        `xml:"LastModifiedDate,omitempty"`
	ParentContentFolder   *ContentFolder   `xml:"ParentContentFolder,omitempty"`
	ParentContentFolderId *ID              `xml:"ParentContentFolderId,omitempty"`
	SystemModstamp        time.Time        `xml:"SystemModstamp,omitempty"`
	Title                 string           `xml:"Title,omitempty"`
}

type ForecastShare struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ForecastShare"`

	*SObject

	AccessLevel      string    `xml:"AccessLevel,omitempty"`
	CanSubmit        bool      `xml:"CanSubmit,omitempty"`
	LastModifiedBy   *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time `xml:"LastModifiedDate,omitempty"`
	RowCause         string    `xml:"RowCause,omitempty"`
	UserOrGroupId    *ID       `xml:"UserOrGroupId,omitempty"`
	UserRoleId       *ID       `xml:"UserRoleId,omitempty"`
}

type ForecastingShare struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ForecastingShare"`

	*SObject

	AccessLevel                 string    `xml:"AccessLevel,omitempty"`
	CreatedBy                   *User     `xml:"CreatedBy,omitempty"`
	CreatedById                 *ID       `xml:"CreatedById,omitempty"`
	CreatedDate                 time.Time `xml:"CreatedDate,omitempty"`
	IsDeleted                   bool      `xml:"IsDeleted,omitempty"`
	LastModifiedBy              *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById            *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate            time.Time `xml:"LastModifiedDate,omitempty"`
	SharedForecastManagerRole   *SObject  `xml:"SharedForecastManagerRole,omitempty"`
	SharedForecastManagerRoleId *ID       `xml:"SharedForecastManagerRoleId,omitempty"`
	SystemModstamp              time.Time `xml:"SystemModstamp,omitempty"`
	UserOrGroup                 *User     `xml:"UserOrGroup,omitempty"`
	UserOrGroupId               *ID       `xml:"UserOrGroupId,omitempty"`
}

type GrantedByLicense struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com GrantedByLicense"`

	*SObject

	CreatedBy              *User                 `xml:"CreatedBy,omitempty"`
	CreatedById            *ID                   `xml:"CreatedById,omitempty"`
	CreatedDate            time.Time             `xml:"CreatedDate,omitempty"`
	CustomPermission       *CustomPermission     `xml:"CustomPermission,omitempty"`
	CustomPermissionId     *ID                   `xml:"CustomPermissionId,omitempty"`
	IsDeleted              bool                  `xml:"IsDeleted,omitempty"`
	LastModifiedBy         *User                 `xml:"LastModifiedBy,omitempty"`
	LastModifiedById       *ID                   `xml:"LastModifiedById,omitempty"`
	LastModifiedDate       time.Time             `xml:"LastModifiedDate,omitempty"`
	PermissionSetLicense   *PermissionSetLicense `xml:"PermissionSetLicense,omitempty"`
	PermissionSetLicenseId *ID                   `xml:"PermissionSetLicenseId,omitempty"`
	SystemModstamp         time.Time             `xml:"SystemModstamp,omitempty"`
}

type Group struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com Group"`

	*SObject

	CreatedBy              *User        `xml:"CreatedBy,omitempty"`
	CreatedById            *ID          `xml:"CreatedById,omitempty"`
	CreatedDate            time.Time    `xml:"CreatedDate,omitempty"`
	DelegatedUsers         *QueryResult `xml:"DelegatedUsers,omitempty"`
	DeveloperName          string       `xml:"DeveloperName,omitempty"`
	DoesIncludeBosses      bool         `xml:"DoesIncludeBosses,omitempty"`
	DoesSendEmailToMembers bool         `xml:"DoesSendEmailToMembers,omitempty"`
	Email                  string       `xml:"Email,omitempty"`
	GroupMembers           *QueryResult `xml:"GroupMembers,omitempty"`
	LastModifiedBy         *User        `xml:"LastModifiedBy,omitempty"`
	LastModifiedById       *ID          `xml:"LastModifiedById,omitempty"`
	LastModifiedDate       time.Time    `xml:"LastModifiedDate,omitempty"`
	Name                   string       `xml:"Name,omitempty"`
	Owner                  *SObject     `xml:"Owner,omitempty"`
	OwnerId                *ID          `xml:"OwnerId,omitempty"`
	QueueSobjects          *QueryResult `xml:"QueueSobjects,omitempty"`
	Related                *SObject     `xml:"Related,omitempty"`
	RelatedId              *ID          `xml:"RelatedId,omitempty"`
	SystemModstamp         time.Time    `xml:"SystemModstamp,omitempty"`
	Type                   string       `xml:"Type,omitempty"`
}

type GroupMember struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com GroupMember"`

	*SObject

	Group          *Group    `xml:"Group,omitempty"`
	GroupId        *ID       `xml:"GroupId,omitempty"`
	SystemModstamp time.Time `xml:"SystemModstamp,omitempty"`
	UserOrGroupId  *ID       `xml:"UserOrGroupId,omitempty"`
}

type Holiday struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com Holiday"`

	*SObject

	ActivityDate            time.Time `xml:"ActivityDate,omitempty"`
	CreatedBy               *User     `xml:"CreatedBy,omitempty"`
	CreatedById             *ID       `xml:"CreatedById,omitempty"`
	CreatedDate             time.Time `xml:"CreatedDate,omitempty"`
	Description             string    `xml:"Description,omitempty"`
	EndTimeInMinutes        int32     `xml:"EndTimeInMinutes,omitempty"`
	IsAllDay                bool      `xml:"IsAllDay,omitempty"`
	IsRecurrence            bool      `xml:"IsRecurrence,omitempty"`
	LastModifiedBy          *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById        *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate        time.Time `xml:"LastModifiedDate,omitempty"`
	Name                    string    `xml:"Name,omitempty"`
	RecurrenceDayOfMonth    int32     `xml:"RecurrenceDayOfMonth,omitempty"`
	RecurrenceDayOfWeekMask int32     `xml:"RecurrenceDayOfWeekMask,omitempty"`
	RecurrenceEndDateOnly   time.Time `xml:"RecurrenceEndDateOnly,omitempty"`
	RecurrenceInstance      string    `xml:"RecurrenceInstance,omitempty"`
	RecurrenceInterval      int32     `xml:"RecurrenceInterval,omitempty"`
	RecurrenceMonthOfYear   string    `xml:"RecurrenceMonthOfYear,omitempty"`
	RecurrenceStartDate     time.Time `xml:"RecurrenceStartDate,omitempty"`
	RecurrenceType          string    `xml:"RecurrenceType,omitempty"`
	StartTimeInMinutes      int32     `xml:"StartTimeInMinutes,omitempty"`
	SystemModstamp          time.Time `xml:"SystemModstamp,omitempty"`
}

type IconDefinition struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com IconDefinition"`

	*SObject

	ContentType     string         `xml:"ContentType,omitempty"`
	DurableId       string         `xml:"DurableId,omitempty"`
	Height          int32          `xml:"Height,omitempty"`
	TabDefinition   *TabDefinition `xml:"TabDefinition,omitempty"`
	TabDefinitionId string         `xml:"TabDefinitionId,omitempty"`
	Theme           string         `xml:"Theme,omitempty"`
	Url             string         `xml:"Url,omitempty"`
	Width           int32          `xml:"Width,omitempty"`
}

type Idea struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com Idea"`

	*SObject

	Body                 string            `xml:"Body,omitempty"`
	Categories           string            `xml:"Categories,omitempty"`
	Comments             *QueryResult      `xml:"Comments,omitempty"`
	Community            *Community        `xml:"Community,omitempty"`
	CommunityId          *ID               `xml:"CommunityId,omitempty"`
	CreatedBy            *User             `xml:"CreatedBy,omitempty"`
	CreatedById          *ID               `xml:"CreatedById,omitempty"`
	CreatedDate          time.Time         `xml:"CreatedDate,omitempty"`
	CreatorFullPhotoUrl  string            `xml:"CreatorFullPhotoUrl,omitempty"`
	CreatorName          string            `xml:"CreatorName,omitempty"`
	CreatorSmallPhotoUrl string            `xml:"CreatorSmallPhotoUrl,omitempty"`
	IsDeleted            bool              `xml:"IsDeleted,omitempty"`
	IsHtml               bool              `xml:"IsHtml,omitempty"`
	IsMerged             bool              `xml:"IsMerged,omitempty"`
	LastComment          *IdeaComment      `xml:"LastComment,omitempty"`
	LastCommentDate      time.Time         `xml:"LastCommentDate,omitempty"`
	LastCommentId        *ID               `xml:"LastCommentId,omitempty"`
	LastModifiedBy       *User             `xml:"LastModifiedBy,omitempty"`
	LastModifiedById     *ID               `xml:"LastModifiedById,omitempty"`
	LastModifiedDate     time.Time         `xml:"LastModifiedDate,omitempty"`
	LastReferencedDate   time.Time         `xml:"LastReferencedDate,omitempty"`
	LastViewedDate       time.Time         `xml:"LastViewedDate,omitempty"`
	NumComments          int32             `xml:"NumComments,omitempty"`
	ParentIdea           *Idea             `xml:"ParentIdea,omitempty"`
	ParentIdeaId         *ID               `xml:"ParentIdeaId,omitempty"`
	RecordType           *RecordType       `xml:"RecordType,omitempty"`
	RecordTypeId         *ID               `xml:"RecordTypeId,omitempty"`
	Status               string            `xml:"Status,omitempty"`
	SystemModstamp       time.Time         `xml:"SystemModstamp,omitempty"`
	Title                string            `xml:"Title,omitempty"`
	UserRecordAccess     *UserRecordAccess `xml:"UserRecordAccess,omitempty"`
	VoteScore            float64           `xml:"VoteScore,omitempty"`
	VoteTotal            float64           `xml:"VoteTotal,omitempty"`
	Votes                *QueryResult      `xml:"Votes,omitempty"`
}

type IdeaComment struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com IdeaComment"`

	*SObject

	CommentBody          string       `xml:"CommentBody,omitempty"`
	CommunityId          *ID          `xml:"CommunityId,omitempty"`
	CreatedBy            *User        `xml:"CreatedBy,omitempty"`
	CreatedById          *ID          `xml:"CreatedById,omitempty"`
	CreatedDate          time.Time    `xml:"CreatedDate,omitempty"`
	CreatorFullPhotoUrl  string       `xml:"CreatorFullPhotoUrl,omitempty"`
	CreatorName          string       `xml:"CreatorName,omitempty"`
	CreatorSmallPhotoUrl string       `xml:"CreatorSmallPhotoUrl,omitempty"`
	Idea                 *Idea        `xml:"Idea,omitempty"`
	IdeaId               *ID          `xml:"IdeaId,omitempty"`
	IsDeleted            bool         `xml:"IsDeleted,omitempty"`
	IsHtml               bool         `xml:"IsHtml,omitempty"`
	SystemModstamp       time.Time    `xml:"SystemModstamp,omitempty"`
	UpVotes              int32        `xml:"UpVotes,omitempty"`
	Votes                *QueryResult `xml:"Votes,omitempty"`
}

type IdpEventLog struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com IdpEventLog"`

	*SObject

	AppId               *ID       `xml:"AppId,omitempty"`
	AuthSessionId       *ID       `xml:"AuthSessionId,omitempty"`
	ErrorCode           string    `xml:"ErrorCode,omitempty"`
	IdentityUsed        string    `xml:"IdentityUsed,omitempty"`
	InitiatedBy         string    `xml:"InitiatedBy,omitempty"`
	OptionsHasLogoutUrl bool      `xml:"OptionsHasLogoutUrl,omitempty"`
	SamlEntityUrl       string    `xml:"SamlEntityUrl,omitempty"`
	SsoType             string    `xml:"SsoType,omitempty"`
	Timestamp           time.Time `xml:"Timestamp,omitempty"`
	UserId              *ID       `xml:"UserId,omitempty"`
}

type InstalledMobileApp struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com InstalledMobileApp"`

	*SObject

	ConnectedApplication   *ConnectedApplication `xml:"ConnectedApplication,omitempty"`
	ConnectedApplicationId *ID                   `xml:"ConnectedApplicationId,omitempty"`
	CreatedBy              *User                 `xml:"CreatedBy,omitempty"`
	CreatedById            *ID                   `xml:"CreatedById,omitempty"`
	CreatedDate            time.Time             `xml:"CreatedDate,omitempty"`
	IsDeleted              bool                  `xml:"IsDeleted,omitempty"`
	LastModifiedBy         *User                 `xml:"LastModifiedBy,omitempty"`
	LastModifiedById       *ID                   `xml:"LastModifiedById,omitempty"`
	LastModifiedDate       time.Time             `xml:"LastModifiedDate,omitempty"`
	Name                   string                `xml:"Name,omitempty"`
	Status                 string                `xml:"Status,omitempty"`
	SystemModstamp         time.Time             `xml:"SystemModstamp,omitempty"`
	User                   *User                 `xml:"User,omitempty"`
	UserId                 *ID                   `xml:"UserId,omitempty"`
	UserRecordAccess       *UserRecordAccess     `xml:"UserRecordAccess,omitempty"`
	Version                string                `xml:"Version,omitempty"`
}

type KnowledgeableUser struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com KnowledgeableUser"`

	*SObject

	RawRank        int32     `xml:"RawRank,omitempty"`
	SystemModstamp time.Time `xml:"SystemModstamp,omitempty"`
	TopicId        *ID       `xml:"TopicId,omitempty"`
	UserId         *ID       `xml:"UserId,omitempty"`
}

type Lead struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com Lead"`

	*SObject

	AcceptedEventRelations        *QueryResult      `xml:"AcceptedEventRelations,omitempty"`
	ActivityHistories             *QueryResult      `xml:"ActivityHistories,omitempty"`
	Address                       *Address          `xml:"Address,omitempty"`
	AnnualRevenue                 float64           `xml:"AnnualRevenue,omitempty"`
	AttachedContentDocuments      *QueryResult      `xml:"AttachedContentDocuments,omitempty"`
	Attachments                   *QueryResult      `xml:"Attachments,omitempty"`
	Campaign                      *Campaign         `xml:"Campaign,omitempty"`
	CampaignMembers               *QueryResult      `xml:"CampaignMembers,omitempty"`
	City                          string            `xml:"City,omitempty"`
	CleanStatus                   string            `xml:"CleanStatus,omitempty"`
	CombinedAttachments           *QueryResult      `xml:"CombinedAttachments,omitempty"`
	Company                       string            `xml:"Company,omitempty"`
	CompanyDunsNumber             string            `xml:"CompanyDunsNumber,omitempty"`
	ContentDocumentLinks          *QueryResult      `xml:"ContentDocumentLinks,omitempty"`
	ConvertedAccount              *Account          `xml:"ConvertedAccount,omitempty"`
	ConvertedAccountId            *ID               `xml:"ConvertedAccountId,omitempty"`
	ConvertedContact              *Contact          `xml:"ConvertedContact,omitempty"`
	ConvertedContactId            *ID               `xml:"ConvertedContactId,omitempty"`
	ConvertedDate                 time.Time         `xml:"ConvertedDate,omitempty"`
	ConvertedOpportunity          *Opportunity      `xml:"ConvertedOpportunity,omitempty"`
	ConvertedOpportunityId        *ID               `xml:"ConvertedOpportunityId,omitempty"`
	Country                       string            `xml:"Country,omitempty"`
	CreatedBy                     *User             `xml:"CreatedBy,omitempty"`
	CreatedById                   *ID               `xml:"CreatedById,omitempty"`
	CreatedDate                   time.Time         `xml:"CreatedDate,omitempty"`
	CurrentGeneratorsc            string            `xml:"CurrentGenerators__c,omitempty"`
	DandbCompany                  *DandBCompany     `xml:"DandbCompany,omitempty"`
	DandbCompanyId                *ID               `xml:"DandbCompanyId,omitempty"`
	DeclinedEventRelations        *QueryResult      `xml:"DeclinedEventRelations,omitempty"`
	Description                   string            `xml:"Description,omitempty"`
	DuplicateRecordItems          *QueryResult      `xml:"DuplicateRecordItems,omitempty"`
	Email                         string            `xml:"Email,omitempty"`
	EmailBouncedDate              time.Time         `xml:"EmailBouncedDate,omitempty"`
	EmailBouncedReason            string            `xml:"EmailBouncedReason,omitempty"`
	EmailMessageRelations         *QueryResult      `xml:"EmailMessageRelations,omitempty"`
	EmailStatuses                 *QueryResult      `xml:"EmailStatuses,omitempty"`
	EventRelations                *QueryResult      `xml:"EventRelations,omitempty"`
	Events                        *QueryResult      `xml:"Events,omitempty"`
	Fax                           string            `xml:"Fax,omitempty"`
	FeedSubscriptionsForEntity    *QueryResult      `xml:"FeedSubscriptionsForEntity,omitempty"`
	Feeds                         *QueryResult      `xml:"Feeds,omitempty"`
	FirstName                     string            `xml:"FirstName,omitempty"`
	GeocodeAccuracy               string            `xml:"GeocodeAccuracy,omitempty"`
	Histories                     *QueryResult      `xml:"Histories,omitempty"`
	Industry                      string            `xml:"Industry,omitempty"`
	IsConverted                   bool              `xml:"IsConverted,omitempty"`
	IsDeleted                     bool              `xml:"IsDeleted,omitempty"`
	IsUnreadByOwner               bool              `xml:"IsUnreadByOwner,omitempty"`
	Jigsaw                        string            `xml:"Jigsaw,omitempty"`
	JigsawContactId               string            `xml:"JigsawContactId,omitempty"`
	LastActivityDate              time.Time         `xml:"LastActivityDate,omitempty"`
	LastModifiedBy                *User             `xml:"LastModifiedBy,omitempty"`
	LastModifiedById              *ID               `xml:"LastModifiedById,omitempty"`
	LastModifiedDate              time.Time         `xml:"LastModifiedDate,omitempty"`
	LastName                      string            `xml:"LastName,omitempty"`
	LastReferencedDate            time.Time         `xml:"LastReferencedDate,omitempty"`
	LastViewedDate                time.Time         `xml:"LastViewedDate,omitempty"`
	Latitude                      float64           `xml:"Latitude,omitempty"`
	LeadCleanInfos                *QueryResult      `xml:"LeadCleanInfos,omitempty"`
	LeadSource                    string            `xml:"LeadSource,omitempty"`
	ListEmailIndividualRecipients *QueryResult      `xml:"ListEmailIndividualRecipients,omitempty"`
	Longitude                     float64           `xml:"Longitude,omitempty"`
	LookedUpFromActivities        *QueryResult      `xml:"LookedUpFromActivities,omitempty"`
	MasterRecord                  *Lead             `xml:"MasterRecord,omitempty"`
	MasterRecordId                *ID               `xml:"MasterRecordId,omitempty"`
	MobilePhone                   string            `xml:"MobilePhone,omitempty"`
	Name                          string            `xml:"Name,omitempty"`
	Notes                         *QueryResult      `xml:"Notes,omitempty"`
	NotesAndAttachments           *QueryResult      `xml:"NotesAndAttachments,omitempty"`
	NumberOfEmployees             int32             `xml:"NumberOfEmployees,omitempty"`
	NumberofLocationsc            float64           `xml:"NumberofLocations__c,omitempty"`
	OpenActivities                *QueryResult      `xml:"OpenActivities,omitempty"`
	OutgoingEmailRelations        *QueryResult      `xml:"OutgoingEmailRelations,omitempty"`
	Owner                         *SObject          `xml:"Owner,omitempty"`
	OwnerId                       *ID               `xml:"OwnerId,omitempty"`
	PersonRecord                  *QueryResult      `xml:"PersonRecord,omitempty"`
	Phone                         string            `xml:"Phone,omitempty"`
	PhotoUrl                      string            `xml:"PhotoUrl,omitempty"`
	PostalCode                    string            `xml:"PostalCode,omitempty"`
	Primaryc                      string            `xml:"Primary__c,omitempty"`
	ProcessInstances              *QueryResult      `xml:"ProcessInstances,omitempty"`
	ProcessSteps                  *QueryResult      `xml:"ProcessSteps,omitempty"`
	ProductInterestc              string            `xml:"ProductInterest__c,omitempty"`
	Rating                        string            `xml:"Rating,omitempty"`
	RecordActionHistories         *QueryResult      `xml:"RecordActionHistories,omitempty"`
	RecordActions                 *QueryResult      `xml:"RecordActions,omitempty"`
	RecordAssociatedGroups        *QueryResult      `xml:"RecordAssociatedGroups,omitempty"`
	SICCodec                      string            `xml:"SICCode__c,omitempty"`
	Salutation                    string            `xml:"Salutation,omitempty"`
	Shares                        *QueryResult      `xml:"Shares,omitempty"`
	State                         string            `xml:"State,omitempty"`
	Status                        string            `xml:"Status,omitempty"`
	Street                        string            `xml:"Street,omitempty"`
	SystemModstamp                time.Time         `xml:"SystemModstamp,omitempty"`
	Tasks                         *QueryResult      `xml:"Tasks,omitempty"`
	Title                         string            `xml:"Title,omitempty"`
	TopicAssignments              *QueryResult      `xml:"TopicAssignments,omitempty"`
	UndecidedEventRelations       *QueryResult      `xml:"UndecidedEventRelations,omitempty"`
	UserRecordAccess              *UserRecordAccess `xml:"UserRecordAccess,omitempty"`
	Website                       string            `xml:"Website,omitempty"`
}

type LeadCleanInfo struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com LeadCleanInfo"`

	*SObject

	Address                           *Address          `xml:"Address,omitempty"`
	AnnualRevenue                     float64           `xml:"AnnualRevenue,omitempty"`
	City                              string            `xml:"City,omitempty"`
	CleanedByJob                      bool              `xml:"CleanedByJob,omitempty"`
	CleanedByUser                     bool              `xml:"CleanedByUser,omitempty"`
	CompanyDunsNumber                 string            `xml:"CompanyDunsNumber,omitempty"`
	CompanyName                       string            `xml:"CompanyName,omitempty"`
	ContactStatusDataDotCom           string            `xml:"ContactStatusDataDotCom,omitempty"`
	Country                           string            `xml:"Country,omitempty"`
	CreatedBy                         *User             `xml:"CreatedBy,omitempty"`
	CreatedById                       *ID               `xml:"CreatedById,omitempty"`
	CreatedDate                       time.Time         `xml:"CreatedDate,omitempty"`
	DandBCompanyDunsNumber            string            `xml:"DandBCompanyDunsNumber,omitempty"`
	DataDotComCompanyId               string            `xml:"DataDotComCompanyId,omitempty"`
	DataDotComId                      string            `xml:"DataDotComId,omitempty"`
	Email                             string            `xml:"Email,omitempty"`
	FirstName                         string            `xml:"FirstName,omitempty"`
	GeocodeAccuracy                   string            `xml:"GeocodeAccuracy,omitempty"`
	Industry                          string            `xml:"Industry,omitempty"`
	IsDeleted                         bool              `xml:"IsDeleted,omitempty"`
	IsDifferentAnnualRevenue          bool              `xml:"IsDifferentAnnualRevenue,omitempty"`
	IsDifferentCity                   bool              `xml:"IsDifferentCity,omitempty"`
	IsDifferentCompanyDunsNumber      bool              `xml:"IsDifferentCompanyDunsNumber,omitempty"`
	IsDifferentCompanyName            bool              `xml:"IsDifferentCompanyName,omitempty"`
	IsDifferentCountry                bool              `xml:"IsDifferentCountry,omitempty"`
	IsDifferentCountryCode            bool              `xml:"IsDifferentCountryCode,omitempty"`
	IsDifferentDandBCompanyDunsNumber bool              `xml:"IsDifferentDandBCompanyDunsNumber,omitempty"`
	IsDifferentEmail                  bool              `xml:"IsDifferentEmail,omitempty"`
	IsDifferentFirstName              bool              `xml:"IsDifferentFirstName,omitempty"`
	IsDifferentIndustry               bool              `xml:"IsDifferentIndustry,omitempty"`
	IsDifferentLastName               bool              `xml:"IsDifferentLastName,omitempty"`
	IsDifferentNumberOfEmployees      bool              `xml:"IsDifferentNumberOfEmployees,omitempty"`
	IsDifferentPhone                  bool              `xml:"IsDifferentPhone,omitempty"`
	IsDifferentPostalCode             bool              `xml:"IsDifferentPostalCode,omitempty"`
	IsDifferentState                  bool              `xml:"IsDifferentState,omitempty"`
	IsDifferentStateCode              bool              `xml:"IsDifferentStateCode,omitempty"`
	IsDifferentStreet                 bool              `xml:"IsDifferentStreet,omitempty"`
	IsDifferentTitle                  bool              `xml:"IsDifferentTitle,omitempty"`
	IsFlaggedWrongAddress             bool              `xml:"IsFlaggedWrongAddress,omitempty"`
	IsFlaggedWrongAnnualRevenue       bool              `xml:"IsFlaggedWrongAnnualRevenue,omitempty"`
	IsFlaggedWrongCompanyDunsNumber   bool              `xml:"IsFlaggedWrongCompanyDunsNumber,omitempty"`
	IsFlaggedWrongCompanyName         bool              `xml:"IsFlaggedWrongCompanyName,omitempty"`
	IsFlaggedWrongEmail               bool              `xml:"IsFlaggedWrongEmail,omitempty"`
	IsFlaggedWrongIndustry            bool              `xml:"IsFlaggedWrongIndustry,omitempty"`
	IsFlaggedWrongName                bool              `xml:"IsFlaggedWrongName,omitempty"`
	IsFlaggedWrongNumberOfEmployees   bool              `xml:"IsFlaggedWrongNumberOfEmployees,omitempty"`
	IsFlaggedWrongPhone               bool              `xml:"IsFlaggedWrongPhone,omitempty"`
	IsFlaggedWrongTitle               bool              `xml:"IsFlaggedWrongTitle,omitempty"`
	IsInactive                        bool              `xml:"IsInactive,omitempty"`
	IsReviewedAddress                 bool              `xml:"IsReviewedAddress,omitempty"`
	IsReviewedAnnualRevenue           bool              `xml:"IsReviewedAnnualRevenue,omitempty"`
	IsReviewedCompanyDunsNumber       bool              `xml:"IsReviewedCompanyDunsNumber,omitempty"`
	IsReviewedCompanyName             bool              `xml:"IsReviewedCompanyName,omitempty"`
	IsReviewedDandBCompanyDunsNumber  bool              `xml:"IsReviewedDandBCompanyDunsNumber,omitempty"`
	IsReviewedEmail                   bool              `xml:"IsReviewedEmail,omitempty"`
	IsReviewedIndustry                bool              `xml:"IsReviewedIndustry,omitempty"`
	IsReviewedName                    bool              `xml:"IsReviewedName,omitempty"`
	IsReviewedNumberOfEmployees       bool              `xml:"IsReviewedNumberOfEmployees,omitempty"`
	IsReviewedPhone                   bool              `xml:"IsReviewedPhone,omitempty"`
	IsReviewedTitle                   bool              `xml:"IsReviewedTitle,omitempty"`
	LastMatchedDate                   time.Time         `xml:"LastMatchedDate,omitempty"`
	LastModifiedBy                    *User             `xml:"LastModifiedBy,omitempty"`
	LastModifiedById                  *ID               `xml:"LastModifiedById,omitempty"`
	LastModifiedDate                  time.Time         `xml:"LastModifiedDate,omitempty"`
	LastName                          string            `xml:"LastName,omitempty"`
	LastStatusChangedBy               *User             `xml:"LastStatusChangedBy,omitempty"`
	LastStatusChangedById             *ID               `xml:"LastStatusChangedById,omitempty"`
	LastStatusChangedDate             time.Time         `xml:"LastStatusChangedDate,omitempty"`
	Latitude                          float64           `xml:"Latitude,omitempty"`
	Lead                              *Lead             `xml:"Lead,omitempty"`
	LeadId                            *ID               `xml:"LeadId,omitempty"`
	Longitude                         float64           `xml:"Longitude,omitempty"`
	Name                              string            `xml:"Name,omitempty"`
	NumberOfEmployees                 int32             `xml:"NumberOfEmployees,omitempty"`
	Phone                             string            `xml:"Phone,omitempty"`
	PostalCode                        string            `xml:"PostalCode,omitempty"`
	State                             string            `xml:"State,omitempty"`
	Street                            string            `xml:"Street,omitempty"`
	SystemModstamp                    time.Time         `xml:"SystemModstamp,omitempty"`
	Title                             string            `xml:"Title,omitempty"`
	UserRecordAccess                  *UserRecordAccess `xml:"UserRecordAccess,omitempty"`
}

type LeadFeed struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com LeadFeed"`

	*SObject

	BestComment        *FeedComment `xml:"BestComment,omitempty"`
	BestCommentId      *ID          `xml:"BestCommentId,omitempty"`
	Body               string       `xml:"Body,omitempty"`
	CommentCount       int32        `xml:"CommentCount,omitempty"`
	CreatedBy          *SObject     `xml:"CreatedBy,omitempty"`
	CreatedById        *ID          `xml:"CreatedById,omitempty"`
	CreatedDate        time.Time    `xml:"CreatedDate,omitempty"`
	FeedAttachments    *QueryResult `xml:"FeedAttachments,omitempty"`
	FeedComments       *QueryResult `xml:"FeedComments,omitempty"`
	FeedLikes          *QueryResult `xml:"FeedLikes,omitempty"`
	FeedSignals        *QueryResult `xml:"FeedSignals,omitempty"`
	FeedTrackedChanges *QueryResult `xml:"FeedTrackedChanges,omitempty"`
	InsertedBy         *SObject     `xml:"InsertedBy,omitempty"`
	InsertedById       *ID          `xml:"InsertedById,omitempty"`
	IsDeleted          bool         `xml:"IsDeleted,omitempty"`
	IsRichText         bool         `xml:"IsRichText,omitempty"`
	LastModifiedDate   time.Time    `xml:"LastModifiedDate,omitempty"`
	LikeCount          int32        `xml:"LikeCount,omitempty"`
	LinkUrl            string       `xml:"LinkUrl,omitempty"`
	Parent             *Lead        `xml:"Parent,omitempty"`
	ParentId           *ID          `xml:"ParentId,omitempty"`
	RelatedRecordId    *ID          `xml:"RelatedRecordId,omitempty"`
	SystemModstamp     time.Time    `xml:"SystemModstamp,omitempty"`
	Title              string       `xml:"Title,omitempty"`
	Type               string       `xml:"Type,omitempty"`
}

type LeadHistory struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com LeadHistory"`

	*SObject

	CreatedBy   *SObject    `xml:"CreatedBy,omitempty"`
	CreatedById *ID         `xml:"CreatedById,omitempty"`
	CreatedDate time.Time   `xml:"CreatedDate,omitempty"`
	Field       string      `xml:"Field,omitempty"`
	IsDeleted   bool        `xml:"IsDeleted,omitempty"`
	Lead        *Lead       `xml:"Lead,omitempty"`
	LeadId      *ID         `xml:"LeadId,omitempty"`
	NewValue    interface{} `xml:"NewValue,omitempty"`
	OldValue    interface{} `xml:"OldValue,omitempty"`
}

type LeadShare struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com LeadShare"`

	*SObject

	IsDeleted        bool      `xml:"IsDeleted,omitempty"`
	LastModifiedBy   *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time `xml:"LastModifiedDate,omitempty"`
	Lead             *Lead     `xml:"Lead,omitempty"`
	LeadAccessLevel  string    `xml:"LeadAccessLevel,omitempty"`
	LeadId           *ID       `xml:"LeadId,omitempty"`
	RowCause         string    `xml:"RowCause,omitempty"`
	UserOrGroup      *SObject  `xml:"UserOrGroup,omitempty"`
	UserOrGroupId    *ID       `xml:"UserOrGroupId,omitempty"`
}

type LeadStatus struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com LeadStatus"`

	*SObject

	ApiName          string    `xml:"ApiName,omitempty"`
	CreatedBy        *User     `xml:"CreatedBy,omitempty"`
	CreatedById      *ID       `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time `xml:"CreatedDate,omitempty"`
	IsConverted      bool      `xml:"IsConverted,omitempty"`
	IsDefault        bool      `xml:"IsDefault,omitempty"`
	LastModifiedBy   *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time `xml:"LastModifiedDate,omitempty"`
	MasterLabel      string    `xml:"MasterLabel,omitempty"`
	SortOrder        int32     `xml:"SortOrder,omitempty"`
	SystemModstamp   time.Time `xml:"SystemModstamp,omitempty"`
}

type LightningComponentBundle struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com LightningComponentBundle"`

	*SObject

	ApiVersion       float64   `xml:"ApiVersion,omitempty"`
	AvailableFor     string    `xml:"AvailableFor,omitempty"`
	CreatedBy        *User     `xml:"CreatedBy,omitempty"`
	CreatedById      *ID       `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time `xml:"CreatedDate,omitempty"`
	Description      string    `xml:"Description,omitempty"`
	DeveloperName    string    `xml:"DeveloperName,omitempty"`
	IsDeleted        bool      `xml:"IsDeleted,omitempty"`
	IsExposed        bool      `xml:"IsExposed,omitempty"`
	Language         string    `xml:"Language,omitempty"`
	LastModifiedBy   *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time `xml:"LastModifiedDate,omitempty"`
	MasterLabel      string    `xml:"MasterLabel,omitempty"`
	MinVersion       float64   `xml:"MinVersion,omitempty"`
	NamespacePrefix  string    `xml:"NamespacePrefix,omitempty"`
	SystemModstamp   time.Time `xml:"SystemModstamp,omitempty"`
	TagConfigs       string    `xml:"TagConfigs,omitempty"`
}

type LightningComponentResource struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com LightningComponentResource"`

	*SObject

	CreatedBy                  *User                     `xml:"CreatedBy,omitempty"`
	CreatedById                *ID                       `xml:"CreatedById,omitempty"`
	CreatedDate                time.Time                 `xml:"CreatedDate,omitempty"`
	FilePath                   string                    `xml:"FilePath,omitempty"`
	Format                     string                    `xml:"Format,omitempty"`
	IsDeleted                  bool                      `xml:"IsDeleted,omitempty"`
	LastModifiedBy             *User                     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById           *ID                       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate           time.Time                 `xml:"LastModifiedDate,omitempty"`
	LightningComponentBundle   *LightningComponentBundle `xml:"LightningComponentBundle,omitempty"`
	LightningComponentBundleId *ID                       `xml:"LightningComponentBundleId,omitempty"`
	Source                     string                    `xml:"Source,omitempty"`
	SystemModstamp             time.Time                 `xml:"SystemModstamp,omitempty"`
}

type LightningComponentTag struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com LightningComponentTag"`

	*SObject

	CreatedBy                  *User                     `xml:"CreatedBy,omitempty"`
	CreatedById                *ID                       `xml:"CreatedById,omitempty"`
	CreatedDate                time.Time                 `xml:"CreatedDate,omitempty"`
	IsDeleted                  bool                      `xml:"IsDeleted,omitempty"`
	LastModifiedBy             *User                     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById           *ID                       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate           time.Time                 `xml:"LastModifiedDate,omitempty"`
	LightningComponentBundle   *LightningComponentBundle `xml:"LightningComponentBundle,omitempty"`
	LightningComponentBundleId *ID                       `xml:"LightningComponentBundleId,omitempty"`
	SystemModstamp             time.Time                 `xml:"SystemModstamp,omitempty"`
	Tag                        string                    `xml:"Tag,omitempty"`
}

type LightningExitByPageMetrics struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com LightningExitByPageMetrics"`

	*SObject

	MetricsDate    time.Time `xml:"MetricsDate,omitempty"`
	PageName       string    `xml:"PageName,omitempty"`
	RecordCount    int32     `xml:"RecordCount,omitempty"`
	SystemModstamp time.Time `xml:"SystemModstamp,omitempty"`
	User           *User     `xml:"User,omitempty"`
	UserId         *ID       `xml:"UserId,omitempty"`
}

type LightningExperienceTheme struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com LightningExperienceTheme"`

	*SObject

	CreatedBy                  *User        `xml:"CreatedBy,omitempty"`
	CreatedById                *ID          `xml:"CreatedById,omitempty"`
	CreatedDate                time.Time    `xml:"CreatedDate,omitempty"`
	DefaultBrandingSet         *BrandingSet `xml:"DefaultBrandingSet,omitempty"`
	DefaultBrandingSetId       *ID          `xml:"DefaultBrandingSetId,omitempty"`
	Description                string       `xml:"Description,omitempty"`
	DeveloperName              string       `xml:"DeveloperName,omitempty"`
	IsDeleted                  bool         `xml:"IsDeleted,omitempty"`
	Language                   string       `xml:"Language,omitempty"`
	LastModifiedBy             *User        `xml:"LastModifiedBy,omitempty"`
	LastModifiedById           *ID          `xml:"LastModifiedById,omitempty"`
	LastModifiedDate           time.Time    `xml:"LastModifiedDate,omitempty"`
	MasterLabel                string       `xml:"MasterLabel,omitempty"`
	NamespacePrefix            string       `xml:"NamespacePrefix,omitempty"`
	ShouldOverrideLoadingImage bool         `xml:"ShouldOverrideLoadingImage,omitempty"`
	SystemModstamp             time.Time    `xml:"SystemModstamp,omitempty"`
}

type LightningToggleMetrics struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com LightningToggleMetrics"`

	*SObject

	Action         string    `xml:"Action,omitempty"`
	MetricsDate    time.Time `xml:"MetricsDate,omitempty"`
	RecordCount    int32     `xml:"RecordCount,omitempty"`
	SystemModstamp time.Time `xml:"SystemModstamp,omitempty"`
	User           *User     `xml:"User,omitempty"`
	UserId         *ID       `xml:"UserId,omitempty"`
}

type LightningUsageByAppTypeMetrics struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com LightningUsageByAppTypeMetrics"`

	*SObject

	AppExperience  string    `xml:"AppExperience,omitempty"`
	MetricsDate    time.Time `xml:"MetricsDate,omitempty"`
	SystemModstamp time.Time `xml:"SystemModstamp,omitempty"`
	User           *User     `xml:"User,omitempty"`
	UserId         *ID       `xml:"UserId,omitempty"`
}

type LightningUsageByBrowserMetrics struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com LightningUsageByBrowserMetrics"`

	*SObject

	Browser        string    `xml:"Browser,omitempty"`
	EptBin3To5     int32     `xml:"EptBin3To5,omitempty"`
	EptBin5To8     int32     `xml:"EptBin5To8,omitempty"`
	EptBin8To10    int32     `xml:"EptBin8To10,omitempty"`
	EptBinOver10   int32     `xml:"EptBinOver10,omitempty"`
	EptBinUnder3   int32     `xml:"EptBinUnder3,omitempty"`
	MetricsDate    time.Time `xml:"MetricsDate,omitempty"`
	PageName       string    `xml:"PageName,omitempty"`
	RecordCountEPT int32     `xml:"RecordCountEPT,omitempty"`
	SumEPT         int32     `xml:"SumEPT,omitempty"`
	SystemModstamp time.Time `xml:"SystemModstamp,omitempty"`
	TotalCount     int32     `xml:"TotalCount,omitempty"`
}

type LightningUsageByFlexiPageMetrics struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com LightningUsageByFlexiPageMetrics"`

	*SObject

	FlexiPageNameOrId string    `xml:"FlexiPageNameOrId,omitempty"`
	FlexiPageType     string    `xml:"FlexiPageType,omitempty"`
	MetricsDate       time.Time `xml:"MetricsDate,omitempty"`
	RecordCountEPT    int32     `xml:"RecordCountEPT,omitempty"`
	SumEPT            int32     `xml:"SumEPT,omitempty"`
	SystemModstamp    time.Time `xml:"SystemModstamp,omitempty"`
	TotalCount        int32     `xml:"TotalCount,omitempty"`
}

type LightningUsageByPageMetrics struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com LightningUsageByPageMetrics"`

	*SObject

	EptBin3To5     int32     `xml:"EptBin3To5,omitempty"`
	EptBin5To8     int32     `xml:"EptBin5To8,omitempty"`
	EptBin8To10    int32     `xml:"EptBin8To10,omitempty"`
	EptBinOver10   int32     `xml:"EptBinOver10,omitempty"`
	EptBinUnder3   int32     `xml:"EptBinUnder3,omitempty"`
	MetricsDate    time.Time `xml:"MetricsDate,omitempty"`
	PageName       string    `xml:"PageName,omitempty"`
	RecordCountEPT int32     `xml:"RecordCountEPT,omitempty"`
	SumEPT         int32     `xml:"SumEPT,omitempty"`
	SystemModstamp time.Time `xml:"SystemModstamp,omitempty"`
	TotalCount     int32     `xml:"TotalCount,omitempty"`
	User           *User     `xml:"User,omitempty"`
	UserId         *ID       `xml:"UserId,omitempty"`
}

type ListEmail struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ListEmail"`

	*SObject

	ActivityHistories        *QueryResult      `xml:"ActivityHistories,omitempty"`
	AttachedContentDocuments *QueryResult      `xml:"AttachedContentDocuments,omitempty"`
	Campaign                 *Campaign         `xml:"Campaign,omitempty"`
	CampaignId               *ID               `xml:"CampaignId,omitempty"`
	CombinedAttachments      *QueryResult      `xml:"CombinedAttachments,omitempty"`
	ContentDocumentIds       []*ID             `xml:"ContentDocumentIds,omitempty"`
	ContentDocumentLinks     *QueryResult      `xml:"ContentDocumentLinks,omitempty"`
	CreatedBy                *User             `xml:"CreatedBy,omitempty"`
	CreatedById              *ID               `xml:"CreatedById,omitempty"`
	CreatedDate              time.Time         `xml:"CreatedDate,omitempty"`
	Emails                   *QueryResult      `xml:"Emails,omitempty"`
	Events                   *QueryResult      `xml:"Events,omitempty"`
	FromAddress              string            `xml:"FromAddress,omitempty"`
	FromName                 string            `xml:"FromName,omitempty"`
	HasAttachment            bool              `xml:"HasAttachment,omitempty"`
	HtmlBody                 string            `xml:"HtmlBody,omitempty"`
	IsDeleted                bool              `xml:"IsDeleted,omitempty"`
	IsTracked                bool              `xml:"IsTracked,omitempty"`
	LastModifiedBy           *User             `xml:"LastModifiedBy,omitempty"`
	LastModifiedById         *ID               `xml:"LastModifiedById,omitempty"`
	LastModifiedDate         time.Time         `xml:"LastModifiedDate,omitempty"`
	LastReferencedDate       time.Time         `xml:"LastReferencedDate,omitempty"`
	LastViewedDate           time.Time         `xml:"LastViewedDate,omitempty"`
	LookedUpFromActivities   *QueryResult      `xml:"LookedUpFromActivities,omitempty"`
	Name                     string            `xml:"Name,omitempty"`
	OpenActivities           *QueryResult      `xml:"OpenActivities,omitempty"`
	Owner                    *SObject          `xml:"Owner,omitempty"`
	OwnerId                  *ID               `xml:"OwnerId,omitempty"`
	ScheduledDate            time.Time         `xml:"ScheduledDate,omitempty"`
	Status                   string            `xml:"Status,omitempty"`
	Subject                  string            `xml:"Subject,omitempty"`
	SystemModstamp           time.Time         `xml:"SystemModstamp,omitempty"`
	Tasks                    *QueryResult      `xml:"Tasks,omitempty"`
	TextBody                 string            `xml:"TextBody,omitempty"`
	TotalSent                int32             `xml:"TotalSent,omitempty"`
	UserRecordAccess         *UserRecordAccess `xml:"UserRecordAccess,omitempty"`
}

type ListEmailIndividualRecipient struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ListEmailIndividualRecipient"`

	*SObject

	CreatedBy        *User             `xml:"CreatedBy,omitempty"`
	CreatedById      *ID               `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time         `xml:"CreatedDate,omitempty"`
	IsDeleted        bool              `xml:"IsDeleted,omitempty"`
	LastModifiedBy   *User             `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID               `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time         `xml:"LastModifiedDate,omitempty"`
	ListEmail        *ListEmail        `xml:"ListEmail,omitempty"`
	ListEmailId      *ID               `xml:"ListEmailId,omitempty"`
	Name             string            `xml:"Name,omitempty"`
	Recipient        *Contact          `xml:"Recipient,omitempty"`
	RecipientId      *ID               `xml:"RecipientId,omitempty"`
	SystemModstamp   time.Time         `xml:"SystemModstamp,omitempty"`
	UserRecordAccess *UserRecordAccess `xml:"UserRecordAccess,omitempty"`
}

type ListEmailRecipientSource struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ListEmailRecipientSource"`

	*SObject

	CreatedBy        *User             `xml:"CreatedBy,omitempty"`
	CreatedById      *ID               `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time         `xml:"CreatedDate,omitempty"`
	IsDeleted        bool              `xml:"IsDeleted,omitempty"`
	LastModifiedBy   *User             `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID               `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time         `xml:"LastModifiedDate,omitempty"`
	ListEmail        *ListEmail        `xml:"ListEmail,omitempty"`
	ListEmailId      *ID               `xml:"ListEmailId,omitempty"`
	Name             string            `xml:"Name,omitempty"`
	SourceList       *SObject          `xml:"SourceList,omitempty"`
	SourceListId     *ID               `xml:"SourceListId,omitempty"`
	SourceType       string            `xml:"SourceType,omitempty"`
	SystemModstamp   time.Time         `xml:"SystemModstamp,omitempty"`
	UserRecordAccess *UserRecordAccess `xml:"UserRecordAccess,omitempty"`
}

type ListEmailShare struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ListEmailShare"`

	*SObject

	AccessLevel      string     `xml:"AccessLevel,omitempty"`
	IsDeleted        bool       `xml:"IsDeleted,omitempty"`
	LastModifiedBy   *User      `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID        `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time  `xml:"LastModifiedDate,omitempty"`
	Parent           *ListEmail `xml:"Parent,omitempty"`
	ParentId         *ID        `xml:"ParentId,omitempty"`
	RowCause         string     `xml:"RowCause,omitempty"`
	UserOrGroup      *SObject   `xml:"UserOrGroup,omitempty"`
	UserOrGroupId    *ID        `xml:"UserOrGroupId,omitempty"`
}

type ListView struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ListView"`

	*SObject

	CreatedBy                 *User        `xml:"CreatedBy,omitempty"`
	CreatedById               *ID          `xml:"CreatedById,omitempty"`
	CreatedDate               time.Time    `xml:"CreatedDate,omitempty"`
	DeveloperName             string       `xml:"DeveloperName,omitempty"`
	IsSoqlCompatible          bool         `xml:"IsSoqlCompatible,omitempty"`
	LastModifiedBy            *User        `xml:"LastModifiedBy,omitempty"`
	LastModifiedById          *ID          `xml:"LastModifiedById,omitempty"`
	LastModifiedDate          time.Time    `xml:"LastModifiedDate,omitempty"`
	LastReferencedDate        time.Time    `xml:"LastReferencedDate,omitempty"`
	LastViewedDate            time.Time    `xml:"LastViewedDate,omitempty"`
	ListEmailRecipientSources *QueryResult `xml:"ListEmailRecipientSources,omitempty"`
	Name                      string       `xml:"Name,omitempty"`
	NamespacePrefix           string       `xml:"NamespacePrefix,omitempty"`
	SobjectType               string       `xml:"SobjectType,omitempty"`
	SystemModstamp            time.Time    `xml:"SystemModstamp,omitempty"`
}

type ListViewChart struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ListViewChart"`

	*SObject

	AggregateField   string    `xml:"AggregateField,omitempty"`
	AggregateType    string    `xml:"AggregateType,omitempty"`
	ChartType        string    `xml:"ChartType,omitempty"`
	CreatedBy        *User     `xml:"CreatedBy,omitempty"`
	CreatedById      *ID       `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time `xml:"CreatedDate,omitempty"`
	DeveloperName    string    `xml:"DeveloperName,omitempty"`
	GroupingField    string    `xml:"GroupingField,omitempty"`
	IsDeleted        bool      `xml:"IsDeleted,omitempty"`
	Language         string    `xml:"Language,omitempty"`
	LastModifiedBy   *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time `xml:"LastModifiedDate,omitempty"`
	MasterLabel      string    `xml:"MasterLabel,omitempty"`
	Owner            *User     `xml:"Owner,omitempty"`
	OwnerId          *ID       `xml:"OwnerId,omitempty"`
	SobjectType      string    `xml:"SobjectType,omitempty"`
	SystemModstamp   time.Time `xml:"SystemModstamp,omitempty"`
}

type ListViewChartInstance struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ListViewChartInstance"`

	*SObject

	AggregateField              string         `xml:"AggregateField,omitempty"`
	AggregateType               string         `xml:"AggregateType,omitempty"`
	ChartType                   string         `xml:"ChartType,omitempty"`
	DataQuery                   string         `xml:"DataQuery,omitempty"`
	DataQueryWithoutUserFilters string         `xml:"DataQueryWithoutUserFilters,omitempty"`
	DeveloperName               string         `xml:"DeveloperName,omitempty"`
	ExternalId                  string         `xml:"ExternalId,omitempty"`
	GroupingField               string         `xml:"GroupingField,omitempty"`
	IsDeletable                 bool           `xml:"IsDeletable,omitempty"`
	IsEditable                  bool           `xml:"IsEditable,omitempty"`
	IsLastViewed                bool           `xml:"IsLastViewed,omitempty"`
	Label                       string         `xml:"Label,omitempty"`
	ListViewChart               *ListViewChart `xml:"ListViewChart,omitempty"`
	ListViewChartId             *ID            `xml:"ListViewChartId,omitempty"`
	ListViewContext             *ListView      `xml:"ListViewContext,omitempty"`
	ListViewContextId           *ID            `xml:"ListViewContextId,omitempty"`
	SourceEntity                string         `xml:"SourceEntity,omitempty"`
}

type LoginGeo struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com LoginGeo"`

	*SObject

	City             string    `xml:"City,omitempty"`
	Country          string    `xml:"Country,omitempty"`
	CountryIso       string    `xml:"CountryIso,omitempty"`
	CreatedBy        *User     `xml:"CreatedBy,omitempty"`
	CreatedById      *ID       `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time `xml:"CreatedDate,omitempty"`
	IsDeleted        bool      `xml:"IsDeleted,omitempty"`
	LastModifiedBy   *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time `xml:"LastModifiedDate,omitempty"`
	Latitude         float64   `xml:"Latitude,omitempty"`
	LoginTime        time.Time `xml:"LoginTime,omitempty"`
	Longitude        float64   `xml:"Longitude,omitempty"`
	PostalCode       string    `xml:"PostalCode,omitempty"`
	Subdivision      string    `xml:"Subdivision,omitempty"`
	SystemModstamp   time.Time `xml:"SystemModstamp,omitempty"`
}

type LoginHistory struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com LoginHistory"`

	*SObject

	ApiType                 string    `xml:"ApiType,omitempty"`
	ApiVersion              string    `xml:"ApiVersion,omitempty"`
	Application             string    `xml:"Application,omitempty"`
	AuthenticationServiceId *ID       `xml:"AuthenticationServiceId,omitempty"`
	Browser                 string    `xml:"Browser,omitempty"`
	CipherSuite             string    `xml:"CipherSuite,omitempty"`
	ClientVersion           string    `xml:"ClientVersion,omitempty"`
	CountryIso              string    `xml:"CountryIso,omitempty"`
	LoginGeo                *LoginGeo `xml:"LoginGeo,omitempty"`
	LoginGeoId              *ID       `xml:"LoginGeoId,omitempty"`
	LoginTime               time.Time `xml:"LoginTime,omitempty"`
	LoginType               string    `xml:"LoginType,omitempty"`
	LoginUrl                string    `xml:"LoginUrl,omitempty"`
	Platform                string    `xml:"Platform,omitempty"`
	SourceIp                string    `xml:"SourceIp,omitempty"`
	Status                  string    `xml:"Status,omitempty"`
	TlsProtocol             string    `xml:"TlsProtocol,omitempty"`
	UserId                  *ID       `xml:"UserId,omitempty"`
}

type LoginIp struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com LoginIp"`

	*SObject

	ChallengeMethod   string    `xml:"ChallengeMethod,omitempty"`
	ChallengeSentDate time.Time `xml:"ChallengeSentDate,omitempty"`
	CreatedDate       time.Time `xml:"CreatedDate,omitempty"`
	IsAuthenticated   bool      `xml:"IsAuthenticated,omitempty"`
	SourceIp          string    `xml:"SourceIp,omitempty"`
	Users             *User     `xml:"Users,omitempty"`
	UsersId           *ID       `xml:"UsersId,omitempty"`
}

type LookedUpFromActivity struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com LookedUpFromActivity"`

	*SObject

	Account                *Account  `xml:"Account,omitempty"`
	AccountId              *ID       `xml:"AccountId,omitempty"`
	ActivityDate           time.Time `xml:"ActivityDate,omitempty"`
	ActivitySubtype        string    `xml:"ActivitySubtype,omitempty"`
	ActivityType           string    `xml:"ActivityType,omitempty"`
	CallDisposition        string    `xml:"CallDisposition,omitempty"`
	CallDurationInSeconds  int32     `xml:"CallDurationInSeconds,omitempty"`
	CallObject             string    `xml:"CallObject,omitempty"`
	CallType               string    `xml:"CallType,omitempty"`
	CreatedBy              *User     `xml:"CreatedBy,omitempty"`
	CreatedById            *ID       `xml:"CreatedById,omitempty"`
	CreatedDate            time.Time `xml:"CreatedDate,omitempty"`
	Description            string    `xml:"Description,omitempty"`
	DurationInMinutes      int32     `xml:"DurationInMinutes,omitempty"`
	EndDateTime            time.Time `xml:"EndDateTime,omitempty"`
	IsAllDayEvent          bool      `xml:"IsAllDayEvent,omitempty"`
	IsClosed               bool      `xml:"IsClosed,omitempty"`
	IsDeleted              bool      `xml:"IsDeleted,omitempty"`
	IsHighPriority         bool      `xml:"IsHighPriority,omitempty"`
	IsReminderSet          bool      `xml:"IsReminderSet,omitempty"`
	IsTask                 bool      `xml:"IsTask,omitempty"`
	IsVisibleInSelfService bool      `xml:"IsVisibleInSelfService,omitempty"`
	LastModifiedBy         *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById       *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate       time.Time `xml:"LastModifiedDate,omitempty"`
	Location               string    `xml:"Location,omitempty"`
	Owner                  *SObject  `xml:"Owner,omitempty"`
	OwnerId                *ID       `xml:"OwnerId,omitempty"`
	Priority               string    `xml:"Priority,omitempty"`
	ReminderDateTime       time.Time `xml:"ReminderDateTime,omitempty"`
	StartDateTime          time.Time `xml:"StartDateTime,omitempty"`
	Status                 string    `xml:"Status,omitempty"`
	Subject                string    `xml:"Subject,omitempty"`
	SystemModstamp         time.Time `xml:"SystemModstamp,omitempty"`
	What                   *SObject  `xml:"What,omitempty"`
	WhatId                 *ID       `xml:"WhatId,omitempty"`
	Who                    *SObject  `xml:"Who,omitempty"`
	WhoId                  *ID       `xml:"WhoId,omitempty"`
}

type Macro struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com Macro"`

	*SObject

	CreatedBy            *User             `xml:"CreatedBy,omitempty"`
	CreatedById          *ID               `xml:"CreatedById,omitempty"`
	CreatedDate          time.Time         `xml:"CreatedDate,omitempty"`
	Description          string            `xml:"Description,omitempty"`
	Histories            *QueryResult      `xml:"Histories,omitempty"`
	IsAlohaSupported     bool              `xml:"IsAlohaSupported,omitempty"`
	IsDeleted            bool              `xml:"IsDeleted,omitempty"`
	IsLightningSupported bool              `xml:"IsLightningSupported,omitempty"`
	LastModifiedBy       *User             `xml:"LastModifiedBy,omitempty"`
	LastModifiedById     *ID               `xml:"LastModifiedById,omitempty"`
	LastModifiedDate     time.Time         `xml:"LastModifiedDate,omitempty"`
	LastReferencedDate   time.Time         `xml:"LastReferencedDate,omitempty"`
	LastViewedDate       time.Time         `xml:"LastViewedDate,omitempty"`
	Name                 string            `xml:"Name,omitempty"`
	Owner                *SObject          `xml:"Owner,omitempty"`
	OwnerId              *ID               `xml:"OwnerId,omitempty"`
	StartingContext      string            `xml:"StartingContext,omitempty"`
	SystemModstamp       time.Time         `xml:"SystemModstamp,omitempty"`
	UserRecordAccess     *UserRecordAccess `xml:"UserRecordAccess,omitempty"`
}

type MacroHistory struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com MacroHistory"`

	*SObject

	CreatedBy   *SObject    `xml:"CreatedBy,omitempty"`
	CreatedById *ID         `xml:"CreatedById,omitempty"`
	CreatedDate time.Time   `xml:"CreatedDate,omitempty"`
	Field       string      `xml:"Field,omitempty"`
	IsDeleted   bool        `xml:"IsDeleted,omitempty"`
	Macro       *Macro      `xml:"Macro,omitempty"`
	MacroId     *ID         `xml:"MacroId,omitempty"`
	NewValue    interface{} `xml:"NewValue,omitempty"`
	OldValue    interface{} `xml:"OldValue,omitempty"`
}

type MacroInstruction struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com MacroInstruction"`

	*SObject

	CreatedBy        *User             `xml:"CreatedBy,omitempty"`
	CreatedById      *ID               `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time         `xml:"CreatedDate,omitempty"`
	IsDeleted        bool              `xml:"IsDeleted,omitempty"`
	LastModifiedBy   *User             `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID               `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time         `xml:"LastModifiedDate,omitempty"`
	Macro            *Macro            `xml:"Macro,omitempty"`
	MacroId          *ID               `xml:"MacroId,omitempty"`
	Name             string            `xml:"Name,omitempty"`
	Operation        string            `xml:"Operation,omitempty"`
	SortOrder        int32             `xml:"SortOrder,omitempty"`
	SystemModstamp   time.Time         `xml:"SystemModstamp,omitempty"`
	Target           string            `xml:"Target,omitempty"`
	UserRecordAccess *UserRecordAccess `xml:"UserRecordAccess,omitempty"`
	Value            string            `xml:"Value,omitempty"`
	ValueRecord      string            `xml:"ValueRecord,omitempty"`
}

type MacroShare struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com MacroShare"`

	*SObject

	AccessLevel      string    `xml:"AccessLevel,omitempty"`
	IsDeleted        bool      `xml:"IsDeleted,omitempty"`
	LastModifiedBy   *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time `xml:"LastModifiedDate,omitempty"`
	Parent           *Macro    `xml:"Parent,omitempty"`
	ParentId         *ID       `xml:"ParentId,omitempty"`
	RowCause         string    `xml:"RowCause,omitempty"`
	UserOrGroup      *SObject  `xml:"UserOrGroup,omitempty"`
	UserOrGroupId    *ID       `xml:"UserOrGroupId,omitempty"`
}

type MailmergeTemplate struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com MailmergeTemplate"`

	*SObject

	Body                                     []byte    `xml:"Body,omitempty"`
	BodyLength                               int32     `xml:"BodyLength,omitempty"`
	CreatedBy                                *User     `xml:"CreatedBy,omitempty"`
	CreatedById                              *ID       `xml:"CreatedById,omitempty"`
	CreatedDate                              time.Time `xml:"CreatedDate,omitempty"`
	Description                              string    `xml:"Description,omitempty"`
	Filename                                 string    `xml:"Filename,omitempty"`
	IsDeleted                                bool      `xml:"IsDeleted,omitempty"`
	LastModifiedBy                           *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById                         *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate                         time.Time `xml:"LastModifiedDate,omitempty"`
	LastUsedDate                             time.Time `xml:"LastUsedDate,omitempty"`
	Name                                     string    `xml:"Name,omitempty"`
	SecurityOptionsAttachmentHasFlash        bool      `xml:"SecurityOptionsAttachmentHasFlash,omitempty"`
	SecurityOptionsAttachmentHasXSSThreat    bool      `xml:"SecurityOptionsAttachmentHasXSSThreat,omitempty"`
	SecurityOptionsAttachmentScannedForXSS   bool      `xml:"SecurityOptionsAttachmentScannedForXSS,omitempty"`
	SecurityOptionsAttachmentScannedforFlash bool      `xml:"SecurityOptionsAttachmentScannedforFlash,omitempty"`
	SystemModstamp                           time.Time `xml:"SystemModstamp,omitempty"`
}

type MatchingRule struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com MatchingRule"`

	*SObject

	BooleanFilter     string       `xml:"BooleanFilter,omitempty"`
	CreatedBy         *User        `xml:"CreatedBy,omitempty"`
	CreatedById       *ID          `xml:"CreatedById,omitempty"`
	CreatedDate       time.Time    `xml:"CreatedDate,omitempty"`
	Description       string       `xml:"Description,omitempty"`
	DeveloperName     string       `xml:"DeveloperName,omitempty"`
	IsDeleted         bool         `xml:"IsDeleted,omitempty"`
	Language          string       `xml:"Language,omitempty"`
	LastModifiedBy    *User        `xml:"LastModifiedBy,omitempty"`
	LastModifiedById  *ID          `xml:"LastModifiedById,omitempty"`
	LastModifiedDate  time.Time    `xml:"LastModifiedDate,omitempty"`
	MasterLabel       string       `xml:"MasterLabel,omitempty"`
	MatchEngine       string       `xml:"MatchEngine,omitempty"`
	MatchingRuleItems *QueryResult `xml:"MatchingRuleItems,omitempty"`
	NamespacePrefix   string       `xml:"NamespacePrefix,omitempty"`
	RuleStatus        string       `xml:"RuleStatus,omitempty"`
	SobjectSubtype    string       `xml:"SobjectSubtype,omitempty"`
	SobjectType       string       `xml:"SobjectType,omitempty"`
	SystemModstamp    time.Time    `xml:"SystemModstamp,omitempty"`
}

type MatchingRuleItem struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com MatchingRuleItem"`

	*SObject

	BlankValueBehavior string        `xml:"BlankValueBehavior,omitempty"`
	CreatedBy          *User         `xml:"CreatedBy,omitempty"`
	CreatedById        *ID           `xml:"CreatedById,omitempty"`
	CreatedDate        time.Time     `xml:"CreatedDate,omitempty"`
	Field              string        `xml:"Field,omitempty"`
	IsDeleted          bool          `xml:"IsDeleted,omitempty"`
	LastModifiedBy     *User         `xml:"LastModifiedBy,omitempty"`
	LastModifiedById   *ID           `xml:"LastModifiedById,omitempty"`
	LastModifiedDate   time.Time     `xml:"LastModifiedDate,omitempty"`
	MatchingMethod     string        `xml:"MatchingMethod,omitempty"`
	MatchingRule       *MatchingRule `xml:"MatchingRule,omitempty"`
	MatchingRuleId     *ID           `xml:"MatchingRuleId,omitempty"`
	SortOrder          int32         `xml:"SortOrder,omitempty"`
	SystemModstamp     time.Time     `xml:"SystemModstamp,omitempty"`
}

type Name struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com Name"`

	*SObject

	Alias              string      `xml:"Alias,omitempty"`
	Email              string      `xml:"Email,omitempty"`
	FirstName          string      `xml:"FirstName,omitempty"`
	IsActive           bool        `xml:"IsActive,omitempty"`
	LastName           string      `xml:"LastName,omitempty"`
	LastReferencedDate time.Time   `xml:"LastReferencedDate,omitempty"`
	LastViewedDate     time.Time   `xml:"LastViewedDate,omitempty"`
	Name               string      `xml:"Name,omitempty"`
	NameOrAlias        string      `xml:"NameOrAlias,omitempty"`
	Phone              string      `xml:"Phone,omitempty"`
	Profile            *Profile    `xml:"Profile,omitempty"`
	ProfileId          *ID         `xml:"ProfileId,omitempty"`
	RecordType         *RecordType `xml:"RecordType,omitempty"`
	RecordTypeId       *ID         `xml:"RecordTypeId,omitempty"`
	Title              string      `xml:"Title,omitempty"`
	Type               string      `xml:"Type,omitempty"`
	UserRole           *UserRole   `xml:"UserRole,omitempty"`
	UserRoleId         *ID         `xml:"UserRoleId,omitempty"`
	Username           string      `xml:"Username,omitempty"`
}

type NamedCredential struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com NamedCredential"`

	*SObject

	AuthProvider                              *AuthProvider `xml:"AuthProvider,omitempty"`
	AuthProviderId                            *ID           `xml:"AuthProviderId,omitempty"`
	CalloutOptionsAllowMergeFieldsInBody      bool          `xml:"CalloutOptionsAllowMergeFieldsInBody,omitempty"`
	CalloutOptionsAllowMergeFieldsInHeader    bool          `xml:"CalloutOptionsAllowMergeFieldsInHeader,omitempty"`
	CalloutOptionsGenerateAuthorizationHeader bool          `xml:"CalloutOptionsGenerateAuthorizationHeader,omitempty"`
	CreatedBy                                 *User         `xml:"CreatedBy,omitempty"`
	CreatedById                               *ID           `xml:"CreatedById,omitempty"`
	CreatedDate                               time.Time     `xml:"CreatedDate,omitempty"`
	CustomHttpHeaders                         *QueryResult  `xml:"CustomHttpHeaders,omitempty"`
	DeveloperName                             string        `xml:"DeveloperName,omitempty"`
	Endpoint                                  string        `xml:"Endpoint,omitempty"`
	IsDeleted                                 bool          `xml:"IsDeleted,omitempty"`
	Language                                  string        `xml:"Language,omitempty"`
	LastModifiedBy                            *User         `xml:"LastModifiedBy,omitempty"`
	LastModifiedById                          *ID           `xml:"LastModifiedById,omitempty"`
	LastModifiedDate                          time.Time     `xml:"LastModifiedDate,omitempty"`
	MasterLabel                               string        `xml:"MasterLabel,omitempty"`
	NamespacePrefix                           string        `xml:"NamespacePrefix,omitempty"`
	PrincipalType                             string        `xml:"PrincipalType,omitempty"`
	SetupEntityAccessItems                    *QueryResult  `xml:"SetupEntityAccessItems,omitempty"`
	SystemModstamp                            time.Time     `xml:"SystemModstamp,omitempty"`
	UserAuths                                 *QueryResult  `xml:"UserAuths,omitempty"`
}

type Note struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com Note"`

	*SObject

	Body             string    `xml:"Body,omitempty"`
	CreatedBy        *User     `xml:"CreatedBy,omitempty"`
	CreatedById      *ID       `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time `xml:"CreatedDate,omitempty"`
	IsDeleted        bool      `xml:"IsDeleted,omitempty"`
	IsPrivate        bool      `xml:"IsPrivate,omitempty"`
	LastModifiedBy   *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time `xml:"LastModifiedDate,omitempty"`
	Owner            *User     `xml:"Owner,omitempty"`
	OwnerId          *ID       `xml:"OwnerId,omitempty"`
	Parent           *SObject  `xml:"Parent,omitempty"`
	ParentId         *ID       `xml:"ParentId,omitempty"`
	SystemModstamp   time.Time `xml:"SystemModstamp,omitempty"`
	Title            string    `xml:"Title,omitempty"`
}

type NoteAndAttachment struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com NoteAndAttachment"`

	*SObject

	CreatedBy        *User     `xml:"CreatedBy,omitempty"`
	CreatedById      *ID       `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time `xml:"CreatedDate,omitempty"`
	IsDeleted        bool      `xml:"IsDeleted,omitempty"`
	IsNote           bool      `xml:"IsNote,omitempty"`
	IsPrivate        bool      `xml:"IsPrivate,omitempty"`
	LastModifiedBy   *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time `xml:"LastModifiedDate,omitempty"`
	Owner            *User     `xml:"Owner,omitempty"`
	OwnerId          *ID       `xml:"OwnerId,omitempty"`
	Parent           *SObject  `xml:"Parent,omitempty"`
	ParentId         *ID       `xml:"ParentId,omitempty"`
	SystemModstamp   time.Time `xml:"SystemModstamp,omitempty"`
	Title            string    `xml:"Title,omitempty"`
}

type OauthToken struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com OauthToken"`

	*SObject

	AccessToken   string       `xml:"AccessToken,omitempty"`
	AppMenuItem   *AppMenuItem `xml:"AppMenuItem,omitempty"`
	AppMenuItemId *ID          `xml:"AppMenuItemId,omitempty"`
	AppName       string       `xml:"AppName,omitempty"`
	CreatedDate   time.Time    `xml:"CreatedDate,omitempty"`
	DeleteToken   string       `xml:"DeleteToken,omitempty"`
	LastUsedDate  time.Time    `xml:"LastUsedDate,omitempty"`
	RequestToken  string       `xml:"RequestToken,omitempty"`
	UseCount      int32        `xml:"UseCount,omitempty"`
	User          *User        `xml:"User,omitempty"`
	UserId        *ID          `xml:"UserId,omitempty"`
}

type ObjectPermissions struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ObjectPermissions"`

	*SObject

	CreatedBy                   *User          `xml:"CreatedBy,omitempty"`
	CreatedById                 *ID            `xml:"CreatedById,omitempty"`
	CreatedDate                 time.Time      `xml:"CreatedDate,omitempty"`
	LastModifiedBy              *User          `xml:"LastModifiedBy,omitempty"`
	LastModifiedById            *ID            `xml:"LastModifiedById,omitempty"`
	LastModifiedDate            time.Time      `xml:"LastModifiedDate,omitempty"`
	Parent                      *PermissionSet `xml:"Parent,omitempty"`
	ParentId                    *ID            `xml:"ParentId,omitempty"`
	PermissionsCreate           bool           `xml:"PermissionsCreate,omitempty"`
	PermissionsDelete           bool           `xml:"PermissionsDelete,omitempty"`
	PermissionsEdit             bool           `xml:"PermissionsEdit,omitempty"`
	PermissionsModifyAllRecords bool           `xml:"PermissionsModifyAllRecords,omitempty"`
	PermissionsRead             bool           `xml:"PermissionsRead,omitempty"`
	PermissionsViewAllRecords   bool           `xml:"PermissionsViewAllRecords,omitempty"`
	SobjectType                 string         `xml:"SobjectType,omitempty"`
	SystemModstamp              time.Time      `xml:"SystemModstamp,omitempty"`
}

type OnboardingMetrics struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com OnboardingMetrics"`

	*SObject

	CreatedBy        *User     `xml:"CreatedBy,omitempty"`
	CreatedById      *ID       `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time `xml:"CreatedDate,omitempty"`
	ExperienceName   string    `xml:"ExperienceName,omitempty"`
	IsDeleted        bool      `xml:"IsDeleted,omitempty"`
	LastModifiedBy   *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time `xml:"LastModifiedDate,omitempty"`
	SeenCount        int32     `xml:"SeenCount,omitempty"`
	SystemModstamp   time.Time `xml:"SystemModstamp,omitempty"`
	User             *User     `xml:"User,omitempty"`
	UserId           *ID       `xml:"UserId,omitempty"`
}

type OpenActivity struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com OpenActivity"`

	*SObject

	Account                *Account      `xml:"Account,omitempty"`
	AccountId              *ID           `xml:"AccountId,omitempty"`
	ActivityDate           time.Time     `xml:"ActivityDate,omitempty"`
	ActivitySubtype        string        `xml:"ActivitySubtype,omitempty"`
	ActivityType           string        `xml:"ActivityType,omitempty"`
	AlternateDetail        *EmailMessage `xml:"AlternateDetail,omitempty"`
	AlternateDetailId      *ID           `xml:"AlternateDetailId,omitempty"`
	CallDisposition        string        `xml:"CallDisposition,omitempty"`
	CallDurationInSeconds  int32         `xml:"CallDurationInSeconds,omitempty"`
	CallObject             string        `xml:"CallObject,omitempty"`
	CallType               string        `xml:"CallType,omitempty"`
	CreatedBy              *User         `xml:"CreatedBy,omitempty"`
	CreatedById            *ID           `xml:"CreatedById,omitempty"`
	CreatedDate            time.Time     `xml:"CreatedDate,omitempty"`
	Description            string        `xml:"Description,omitempty"`
	DurationInMinutes      int32         `xml:"DurationInMinutes,omitempty"`
	EndDateTime            time.Time     `xml:"EndDateTime,omitempty"`
	IsAllDayEvent          bool          `xml:"IsAllDayEvent,omitempty"`
	IsClosed               bool          `xml:"IsClosed,omitempty"`
	IsDeleted              bool          `xml:"IsDeleted,omitempty"`
	IsHighPriority         bool          `xml:"IsHighPriority,omitempty"`
	IsReminderSet          bool          `xml:"IsReminderSet,omitempty"`
	IsTask                 bool          `xml:"IsTask,omitempty"`
	IsVisibleInSelfService bool          `xml:"IsVisibleInSelfService,omitempty"`
	LastModifiedBy         *User         `xml:"LastModifiedBy,omitempty"`
	LastModifiedById       *ID           `xml:"LastModifiedById,omitempty"`
	LastModifiedDate       time.Time     `xml:"LastModifiedDate,omitempty"`
	Location               string        `xml:"Location,omitempty"`
	Owner                  *SObject      `xml:"Owner,omitempty"`
	OwnerId                *ID           `xml:"OwnerId,omitempty"`
	Priority               string        `xml:"Priority,omitempty"`
	ReminderDateTime       time.Time     `xml:"ReminderDateTime,omitempty"`
	StartDateTime          time.Time     `xml:"StartDateTime,omitempty"`
	Status                 string        `xml:"Status,omitempty"`
	Subject                string        `xml:"Subject,omitempty"`
	SystemModstamp         time.Time     `xml:"SystemModstamp,omitempty"`
	What                   *SObject      `xml:"What,omitempty"`
	WhatId                 *ID           `xml:"WhatId,omitempty"`
	Who                    *SObject      `xml:"Who,omitempty"`
	WhoId                  *ID           `xml:"WhoId,omitempty"`
}

type Opportunity struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com Opportunity"`

	*SObject

	Account                     *Account          `xml:"Account,omitempty"`
	AccountId                   *ID               `xml:"AccountId,omitempty"`
	AccountPartners             *QueryResult      `xml:"AccountPartners,omitempty"`
	ActivityHistories           *QueryResult      `xml:"ActivityHistories,omitempty"`
	Amount                      float64           `xml:"Amount,omitempty"`
	AttachedContentDocuments    *QueryResult      `xml:"AttachedContentDocuments,omitempty"`
	Attachments                 *QueryResult      `xml:"Attachments,omitempty"`
	Campaign                    *Campaign         `xml:"Campaign,omitempty"`
	CampaignId                  *ID               `xml:"CampaignId,omitempty"`
	CloseDate                   time.Time         `xml:"CloseDate,omitempty"`
	CombinedAttachments         *QueryResult      `xml:"CombinedAttachments,omitempty"`
	ContentDocumentLinks        *QueryResult      `xml:"ContentDocumentLinks,omitempty"`
	CreatedBy                   *User             `xml:"CreatedBy,omitempty"`
	CreatedById                 *ID               `xml:"CreatedById,omitempty"`
	CreatedDate                 time.Time         `xml:"CreatedDate,omitempty"`
	CurrentGeneratorsc          string            `xml:"CurrentGenerators__c,omitempty"`
	DeliveryInstallationStatusc string            `xml:"DeliveryInstallationStatus__c,omitempty"`
	Description                 string            `xml:"Description,omitempty"`
	Emails                      *QueryResult      `xml:"Emails,omitempty"`
	Events                      *QueryResult      `xml:"Events,omitempty"`
	ExpectedRevenue             float64           `xml:"ExpectedRevenue,omitempty"`
	FeedSubscriptionsForEntity  *QueryResult      `xml:"FeedSubscriptionsForEntity,omitempty"`
	Feeds                       *QueryResult      `xml:"Feeds,omitempty"`
	Fiscal                      string            `xml:"Fiscal,omitempty"`
	FiscalQuarter               int32             `xml:"FiscalQuarter,omitempty"`
	FiscalYear                  int32             `xml:"FiscalYear,omitempty"`
	ForecastCategory            string            `xml:"ForecastCategory,omitempty"`
	ForecastCategoryName        string            `xml:"ForecastCategoryName,omitempty"`
	HasOpenActivity             bool              `xml:"HasOpenActivity,omitempty"`
	HasOpportunityLineItem      bool              `xml:"HasOpportunityLineItem,omitempty"`
	HasOverdueTask              bool              `xml:"HasOverdueTask,omitempty"`
	Histories                   *QueryResult      `xml:"Histories,omitempty"`
	IsClosed                    bool              `xml:"IsClosed,omitempty"`
	IsDeleted                   bool              `xml:"IsDeleted,omitempty"`
	IsPrivate                   bool              `xml:"IsPrivate,omitempty"`
	IsWon                       bool              `xml:"IsWon,omitempty"`
	LastActivityDate            time.Time         `xml:"LastActivityDate,omitempty"`
	LastModifiedBy              *User             `xml:"LastModifiedBy,omitempty"`
	LastModifiedById            *ID               `xml:"LastModifiedById,omitempty"`
	LastModifiedDate            time.Time         `xml:"LastModifiedDate,omitempty"`
	LastReferencedDate          time.Time         `xml:"LastReferencedDate,omitempty"`
	LastViewedDate              time.Time         `xml:"LastViewedDate,omitempty"`
	LeadSource                  string            `xml:"LeadSource,omitempty"`
	LookedUpFromActivities      *QueryResult      `xml:"LookedUpFromActivities,omitempty"`
	MainCompetitorsc            string            `xml:"MainCompetitors__c,omitempty"`
	Name                        string            `xml:"Name,omitempty"`
	NextStep                    string            `xml:"NextStep,omitempty"`
	Notes                       *QueryResult      `xml:"Notes,omitempty"`
	NotesAndAttachments         *QueryResult      `xml:"NotesAndAttachments,omitempty"`
	OpenActivities              *QueryResult      `xml:"OpenActivities,omitempty"`
	OpportunityCompetitors      *QueryResult      `xml:"OpportunityCompetitors,omitempty"`
	OpportunityContactRoles     *QueryResult      `xml:"OpportunityContactRoles,omitempty"`
	OpportunityHistories        *QueryResult      `xml:"OpportunityHistories,omitempty"`
	OpportunityLineItems        *QueryResult      `xml:"OpportunityLineItems,omitempty"`
	OpportunityPartnersFrom     *QueryResult      `xml:"OpportunityPartnersFrom,omitempty"`
	OrderNumberc                string            `xml:"OrderNumber__c,omitempty"`
	Orders                      *QueryResult      `xml:"Orders,omitempty"`
	Owner                       *User             `xml:"Owner,omitempty"`
	OwnerId                     *ID               `xml:"OwnerId,omitempty"`
	Partners                    *QueryResult      `xml:"Partners,omitempty"`
	Pricebook2                  *Pricebook2       `xml:"Pricebook2,omitempty"`
	Pricebook2Id                *ID               `xml:"Pricebook2Id,omitempty"`
	Probability                 float64           `xml:"Probability,omitempty"`
	ProcessInstances            *QueryResult      `xml:"ProcessInstances,omitempty"`
	ProcessSteps                *QueryResult      `xml:"ProcessSteps,omitempty"`
	RecordActionHistories       *QueryResult      `xml:"RecordActionHistories,omitempty"`
	RecordActions               *QueryResult      `xml:"RecordActions,omitempty"`
	RecordAssociatedGroups      *QueryResult      `xml:"RecordAssociatedGroups,omitempty"`
	Shares                      *QueryResult      `xml:"Shares,omitempty"`
	StageName                   string            `xml:"StageName,omitempty"`
	SystemModstamp              time.Time         `xml:"SystemModstamp,omitempty"`
	Tasks                       *QueryResult      `xml:"Tasks,omitempty"`
	TopicAssignments            *QueryResult      `xml:"TopicAssignments,omitempty"`
	TotalOpportunityQuantity    float64           `xml:"TotalOpportunityQuantity,omitempty"`
	TrackingNumberc             string            `xml:"TrackingNumber__c,omitempty"`
	Type                        string            `xml:"Type,omitempty"`
	UserRecordAccess            *UserRecordAccess `xml:"UserRecordAccess,omitempty"`
}

type OpportunityCompetitor struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com OpportunityCompetitor"`

	*SObject

	CompetitorName   string       `xml:"CompetitorName,omitempty"`
	CreatedBy        *User        `xml:"CreatedBy,omitempty"`
	CreatedById      *ID          `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time    `xml:"CreatedDate,omitempty"`
	IsDeleted        bool         `xml:"IsDeleted,omitempty"`
	LastModifiedBy   *User        `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID          `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time    `xml:"LastModifiedDate,omitempty"`
	Opportunity      *Opportunity `xml:"Opportunity,omitempty"`
	OpportunityId    *ID          `xml:"OpportunityId,omitempty"`
	Strengths        string       `xml:"Strengths,omitempty"`
	SystemModstamp   time.Time    `xml:"SystemModstamp,omitempty"`
	Weaknesses       string       `xml:"Weaknesses,omitempty"`
}

type OpportunityContactRole struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com OpportunityContactRole"`

	*SObject

	Contact          *Contact     `xml:"Contact,omitempty"`
	ContactId        *ID          `xml:"ContactId,omitempty"`
	CreatedBy        *User        `xml:"CreatedBy,omitempty"`
	CreatedById      *ID          `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time    `xml:"CreatedDate,omitempty"`
	IsDeleted        bool         `xml:"IsDeleted,omitempty"`
	IsPrimary        bool         `xml:"IsPrimary,omitempty"`
	LastModifiedBy   *User        `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID          `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time    `xml:"LastModifiedDate,omitempty"`
	Opportunity      *Opportunity `xml:"Opportunity,omitempty"`
	OpportunityId    *ID          `xml:"OpportunityId,omitempty"`
	Role             string       `xml:"Role,omitempty"`
	SystemModstamp   time.Time    `xml:"SystemModstamp,omitempty"`
}

type OpportunityFeed struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com OpportunityFeed"`

	*SObject

	BestComment        *FeedComment `xml:"BestComment,omitempty"`
	BestCommentId      *ID          `xml:"BestCommentId,omitempty"`
	Body               string       `xml:"Body,omitempty"`
	CommentCount       int32        `xml:"CommentCount,omitempty"`
	CreatedBy          *SObject     `xml:"CreatedBy,omitempty"`
	CreatedById        *ID          `xml:"CreatedById,omitempty"`
	CreatedDate        time.Time    `xml:"CreatedDate,omitempty"`
	FeedAttachments    *QueryResult `xml:"FeedAttachments,omitempty"`
	FeedComments       *QueryResult `xml:"FeedComments,omitempty"`
	FeedLikes          *QueryResult `xml:"FeedLikes,omitempty"`
	FeedSignals        *QueryResult `xml:"FeedSignals,omitempty"`
	FeedTrackedChanges *QueryResult `xml:"FeedTrackedChanges,omitempty"`
	InsertedBy         *SObject     `xml:"InsertedBy,omitempty"`
	InsertedById       *ID          `xml:"InsertedById,omitempty"`
	IsDeleted          bool         `xml:"IsDeleted,omitempty"`
	IsRichText         bool         `xml:"IsRichText,omitempty"`
	LastModifiedDate   time.Time    `xml:"LastModifiedDate,omitempty"`
	LikeCount          int32        `xml:"LikeCount,omitempty"`
	LinkUrl            string       `xml:"LinkUrl,omitempty"`
	Parent             *Opportunity `xml:"Parent,omitempty"`
	ParentId           *ID          `xml:"ParentId,omitempty"`
	RelatedRecordId    *ID          `xml:"RelatedRecordId,omitempty"`
	SystemModstamp     time.Time    `xml:"SystemModstamp,omitempty"`
	Title              string       `xml:"Title,omitempty"`
	Type               string       `xml:"Type,omitempty"`
}

type OpportunityFieldHistory struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com OpportunityFieldHistory"`

	*SObject

	CreatedBy     *SObject     `xml:"CreatedBy,omitempty"`
	CreatedById   *ID          `xml:"CreatedById,omitempty"`
	CreatedDate   time.Time    `xml:"CreatedDate,omitempty"`
	Field         string       `xml:"Field,omitempty"`
	IsDeleted     bool         `xml:"IsDeleted,omitempty"`
	NewValue      interface{}  `xml:"NewValue,omitempty"`
	OldValue      interface{}  `xml:"OldValue,omitempty"`
	Opportunity   *Opportunity `xml:"Opportunity,omitempty"`
	OpportunityId *ID          `xml:"OpportunityId,omitempty"`
}

type OpportunityHistory struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com OpportunityHistory"`

	*SObject

	Amount           float64      `xml:"Amount,omitempty"`
	CloseDate        time.Time    `xml:"CloseDate,omitempty"`
	CreatedBy        *User        `xml:"CreatedBy,omitempty"`
	CreatedById      *ID          `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time    `xml:"CreatedDate,omitempty"`
	ExpectedRevenue  float64      `xml:"ExpectedRevenue,omitempty"`
	ForecastCategory string       `xml:"ForecastCategory,omitempty"`
	IsDeleted        bool         `xml:"IsDeleted,omitempty"`
	Opportunity      *Opportunity `xml:"Opportunity,omitempty"`
	OpportunityId    *ID          `xml:"OpportunityId,omitempty"`
	Probability      float64      `xml:"Probability,omitempty"`
	StageName        string       `xml:"StageName,omitempty"`
	SystemModstamp   time.Time    `xml:"SystemModstamp,omitempty"`
}

type OpportunityLineItem struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com OpportunityLineItem"`

	*SObject

	CreatedBy        *User             `xml:"CreatedBy,omitempty"`
	CreatedById      *ID               `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time         `xml:"CreatedDate,omitempty"`
	Description      string            `xml:"Description,omitempty"`
	IsDeleted        bool              `xml:"IsDeleted,omitempty"`
	LastModifiedBy   *User             `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID               `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time         `xml:"LastModifiedDate,omitempty"`
	ListPrice        float64           `xml:"ListPrice,omitempty"`
	Name             string            `xml:"Name,omitempty"`
	Opportunity      *Opportunity      `xml:"Opportunity,omitempty"`
	OpportunityId    *ID               `xml:"OpportunityId,omitempty"`
	PricebookEntry   *PricebookEntry   `xml:"PricebookEntry,omitempty"`
	PricebookEntryId *ID               `xml:"PricebookEntryId,omitempty"`
	Product2         *Product2         `xml:"Product2,omitempty"`
	Product2Id       *ID               `xml:"Product2Id,omitempty"`
	ProductCode      string            `xml:"ProductCode,omitempty"`
	Quantity         float64           `xml:"Quantity,omitempty"`
	ServiceDate      time.Time         `xml:"ServiceDate,omitempty"`
	SortOrder        int32             `xml:"SortOrder,omitempty"`
	SystemModstamp   time.Time         `xml:"SystemModstamp,omitempty"`
	TotalPrice       float64           `xml:"TotalPrice,omitempty"`
	UnitPrice        float64           `xml:"UnitPrice,omitempty"`
	UserRecordAccess *UserRecordAccess `xml:"UserRecordAccess,omitempty"`
}

type OpportunityPartner struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com OpportunityPartner"`

	*SObject

	AccountTo        *Account     `xml:"AccountTo,omitempty"`
	AccountToId      *ID          `xml:"AccountToId,omitempty"`
	CreatedBy        *User        `xml:"CreatedBy,omitempty"`
	CreatedById      *ID          `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time    `xml:"CreatedDate,omitempty"`
	IsDeleted        bool         `xml:"IsDeleted,omitempty"`
	IsPrimary        bool         `xml:"IsPrimary,omitempty"`
	LastModifiedBy   *User        `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID          `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time    `xml:"LastModifiedDate,omitempty"`
	Opportunity      *Opportunity `xml:"Opportunity,omitempty"`
	OpportunityId    *ID          `xml:"OpportunityId,omitempty"`
	ReversePartnerId *ID          `xml:"ReversePartnerId,omitempty"`
	Role             string       `xml:"Role,omitempty"`
	SystemModstamp   time.Time    `xml:"SystemModstamp,omitempty"`
}

type OpportunityShare struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com OpportunityShare"`

	*SObject

	IsDeleted              bool         `xml:"IsDeleted,omitempty"`
	LastModifiedBy         *User        `xml:"LastModifiedBy,omitempty"`
	LastModifiedById       *ID          `xml:"LastModifiedById,omitempty"`
	LastModifiedDate       time.Time    `xml:"LastModifiedDate,omitempty"`
	Opportunity            *Opportunity `xml:"Opportunity,omitempty"`
	OpportunityAccessLevel string       `xml:"OpportunityAccessLevel,omitempty"`
	OpportunityId          *ID          `xml:"OpportunityId,omitempty"`
	RowCause               string       `xml:"RowCause,omitempty"`
	UserOrGroup            *SObject     `xml:"UserOrGroup,omitempty"`
	UserOrGroupId          *ID          `xml:"UserOrGroupId,omitempty"`
}

type OpportunityStage struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com OpportunityStage"`

	*SObject

	ApiName              string    `xml:"ApiName,omitempty"`
	CreatedBy            *User     `xml:"CreatedBy,omitempty"`
	CreatedById          *ID       `xml:"CreatedById,omitempty"`
	CreatedDate          time.Time `xml:"CreatedDate,omitempty"`
	DefaultProbability   float64   `xml:"DefaultProbability,omitempty"`
	Description          string    `xml:"Description,omitempty"`
	ForecastCategory     string    `xml:"ForecastCategory,omitempty"`
	ForecastCategoryName string    `xml:"ForecastCategoryName,omitempty"`
	IsActive             bool      `xml:"IsActive,omitempty"`
	IsClosed             bool      `xml:"IsClosed,omitempty"`
	IsWon                bool      `xml:"IsWon,omitempty"`
	LastModifiedBy       *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById     *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate     time.Time `xml:"LastModifiedDate,omitempty"`
	MasterLabel          string    `xml:"MasterLabel,omitempty"`
	SortOrder            int32     `xml:"SortOrder,omitempty"`
	SystemModstamp       time.Time `xml:"SystemModstamp,omitempty"`
}

type Order struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com Order"`

	*SObject

	Account                    *Account          `xml:"Account,omitempty"`
	AccountId                  *ID               `xml:"AccountId,omitempty"`
	ActivatedBy                *User             `xml:"ActivatedBy,omitempty"`
	ActivatedById              *ID               `xml:"ActivatedById,omitempty"`
	ActivatedDate              time.Time         `xml:"ActivatedDate,omitempty"`
	ActivityHistories          *QueryResult      `xml:"ActivityHistories,omitempty"`
	AttachedContentDocuments   *QueryResult      `xml:"AttachedContentDocuments,omitempty"`
	Attachments                *QueryResult      `xml:"Attachments,omitempty"`
	BillToContact              *Contact          `xml:"BillToContact,omitempty"`
	BillToContactId            *ID               `xml:"BillToContactId,omitempty"`
	BillingAddress             *Address          `xml:"BillingAddress,omitempty"`
	BillingCity                string            `xml:"BillingCity,omitempty"`
	BillingCountry             string            `xml:"BillingCountry,omitempty"`
	BillingGeocodeAccuracy     string            `xml:"BillingGeocodeAccuracy,omitempty"`
	BillingLatitude            float64           `xml:"BillingLatitude,omitempty"`
	BillingLongitude           float64           `xml:"BillingLongitude,omitempty"`
	BillingPostalCode          string            `xml:"BillingPostalCode,omitempty"`
	BillingState               string            `xml:"BillingState,omitempty"`
	BillingStreet              string            `xml:"BillingStreet,omitempty"`
	CombinedAttachments        *QueryResult      `xml:"CombinedAttachments,omitempty"`
	CompanyAuthorizedBy        *User             `xml:"CompanyAuthorizedBy,omitempty"`
	CompanyAuthorizedById      *ID               `xml:"CompanyAuthorizedById,omitempty"`
	CompanyAuthorizedDate      time.Time         `xml:"CompanyAuthorizedDate,omitempty"`
	ContentDocumentLinks       *QueryResult      `xml:"ContentDocumentLinks,omitempty"`
	Contract                   *Contract         `xml:"Contract,omitempty"`
	ContractId                 *ID               `xml:"ContractId,omitempty"`
	CreatedBy                  *User             `xml:"CreatedBy,omitempty"`
	CreatedById                *ID               `xml:"CreatedById,omitempty"`
	CreatedDate                time.Time         `xml:"CreatedDate,omitempty"`
	CustomerAuthorizedBy       *Contact          `xml:"CustomerAuthorizedBy,omitempty"`
	CustomerAuthorizedById     *ID               `xml:"CustomerAuthorizedById,omitempty"`
	CustomerAuthorizedDate     time.Time         `xml:"CustomerAuthorizedDate,omitempty"`
	Description                string            `xml:"Description,omitempty"`
	EffectiveDate              time.Time         `xml:"EffectiveDate,omitempty"`
	Emails                     *QueryResult      `xml:"Emails,omitempty"`
	EndDate                    time.Time         `xml:"EndDate,omitempty"`
	Events                     *QueryResult      `xml:"Events,omitempty"`
	FeedSubscriptionsForEntity *QueryResult      `xml:"FeedSubscriptionsForEntity,omitempty"`
	Feeds                      *QueryResult      `xml:"Feeds,omitempty"`
	Histories                  *QueryResult      `xml:"Histories,omitempty"`
	IsDeleted                  bool              `xml:"IsDeleted,omitempty"`
	IsReductionOrder           bool              `xml:"IsReductionOrder,omitempty"`
	LastModifiedBy             *User             `xml:"LastModifiedBy,omitempty"`
	LastModifiedById           *ID               `xml:"LastModifiedById,omitempty"`
	LastModifiedDate           time.Time         `xml:"LastModifiedDate,omitempty"`
	LastReferencedDate         time.Time         `xml:"LastReferencedDate,omitempty"`
	LastViewedDate             time.Time         `xml:"LastViewedDate,omitempty"`
	LookedUpFromActivities     *QueryResult      `xml:"LookedUpFromActivities,omitempty"`
	Name                       string            `xml:"Name,omitempty"`
	Notes                      *QueryResult      `xml:"Notes,omitempty"`
	NotesAndAttachments        *QueryResult      `xml:"NotesAndAttachments,omitempty"`
	OpenActivities             *QueryResult      `xml:"OpenActivities,omitempty"`
	OrderItems                 *QueryResult      `xml:"OrderItems,omitempty"`
	OrderNumber                string            `xml:"OrderNumber,omitempty"`
	OrderReferenceNumber       string            `xml:"OrderReferenceNumber,omitempty"`
	Orders                     *QueryResult      `xml:"Orders,omitempty"`
	OriginalOrder              *Order            `xml:"OriginalOrder,omitempty"`
	OriginalOrderId            *ID               `xml:"OriginalOrderId,omitempty"`
	Owner                      *SObject          `xml:"Owner,omitempty"`
	OwnerId                    *ID               `xml:"OwnerId,omitempty"`
	PoDate                     time.Time         `xml:"PoDate,omitempty"`
	PoNumber                   string            `xml:"PoNumber,omitempty"`
	Pricebook2                 *Pricebook2       `xml:"Pricebook2,omitempty"`
	Pricebook2Id               *ID               `xml:"Pricebook2Id,omitempty"`
	ProcessInstances           *QueryResult      `xml:"ProcessInstances,omitempty"`
	ProcessSteps               *QueryResult      `xml:"ProcessSteps,omitempty"`
	RecordActionHistories      *QueryResult      `xml:"RecordActionHistories,omitempty"`
	RecordActions              *QueryResult      `xml:"RecordActions,omitempty"`
	Shares                     *QueryResult      `xml:"Shares,omitempty"`
	ShipToContact              *Contact          `xml:"ShipToContact,omitempty"`
	ShipToContactId            *ID               `xml:"ShipToContactId,omitempty"`
	ShippingAddress            *Address          `xml:"ShippingAddress,omitempty"`
	ShippingCity               string            `xml:"ShippingCity,omitempty"`
	ShippingCountry            string            `xml:"ShippingCountry,omitempty"`
	ShippingGeocodeAccuracy    string            `xml:"ShippingGeocodeAccuracy,omitempty"`
	ShippingLatitude           float64           `xml:"ShippingLatitude,omitempty"`
	ShippingLongitude          float64           `xml:"ShippingLongitude,omitempty"`
	ShippingPostalCode         string            `xml:"ShippingPostalCode,omitempty"`
	ShippingState              string            `xml:"ShippingState,omitempty"`
	ShippingStreet             string            `xml:"ShippingStreet,omitempty"`
	Status                     string            `xml:"Status,omitempty"`
	StatusCode                 string            `xml:"StatusCode,omitempty"`
	SystemModstamp             time.Time         `xml:"SystemModstamp,omitempty"`
	Tasks                      *QueryResult      `xml:"Tasks,omitempty"`
	TopicAssignments           *QueryResult      `xml:"TopicAssignments,omitempty"`
	TotalAmount                float64           `xml:"TotalAmount,omitempty"`
	Type                       string            `xml:"Type,omitempty"`
	UserRecordAccess           *UserRecordAccess `xml:"UserRecordAccess,omitempty"`
}

type OrderFeed struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com OrderFeed"`

	*SObject

	BestComment        *FeedComment `xml:"BestComment,omitempty"`
	BestCommentId      *ID          `xml:"BestCommentId,omitempty"`
	Body               string       `xml:"Body,omitempty"`
	CommentCount       int32        `xml:"CommentCount,omitempty"`
	CreatedBy          *SObject     `xml:"CreatedBy,omitempty"`
	CreatedById        *ID          `xml:"CreatedById,omitempty"`
	CreatedDate        time.Time    `xml:"CreatedDate,omitempty"`
	FeedAttachments    *QueryResult `xml:"FeedAttachments,omitempty"`
	FeedComments       *QueryResult `xml:"FeedComments,omitempty"`
	FeedLikes          *QueryResult `xml:"FeedLikes,omitempty"`
	FeedSignals        *QueryResult `xml:"FeedSignals,omitempty"`
	FeedTrackedChanges *QueryResult `xml:"FeedTrackedChanges,omitempty"`
	InsertedBy         *SObject     `xml:"InsertedBy,omitempty"`
	InsertedById       *ID          `xml:"InsertedById,omitempty"`
	IsDeleted          bool         `xml:"IsDeleted,omitempty"`
	IsRichText         bool         `xml:"IsRichText,omitempty"`
	LastModifiedDate   time.Time    `xml:"LastModifiedDate,omitempty"`
	LikeCount          int32        `xml:"LikeCount,omitempty"`
	LinkUrl            string       `xml:"LinkUrl,omitempty"`
	Parent             *Order       `xml:"Parent,omitempty"`
	ParentId           *ID          `xml:"ParentId,omitempty"`
	RelatedRecordId    *ID          `xml:"RelatedRecordId,omitempty"`
	SystemModstamp     time.Time    `xml:"SystemModstamp,omitempty"`
	Title              string       `xml:"Title,omitempty"`
	Type               string       `xml:"Type,omitempty"`
}

type OrderHistory struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com OrderHistory"`

	*SObject

	CreatedBy   *SObject    `xml:"CreatedBy,omitempty"`
	CreatedById *ID         `xml:"CreatedById,omitempty"`
	CreatedDate time.Time   `xml:"CreatedDate,omitempty"`
	Field       string      `xml:"Field,omitempty"`
	IsDeleted   bool        `xml:"IsDeleted,omitempty"`
	NewValue    interface{} `xml:"NewValue,omitempty"`
	OldValue    interface{} `xml:"OldValue,omitempty"`
	Order       *Order      `xml:"Order,omitempty"`
	OrderId     *ID         `xml:"OrderId,omitempty"`
}

type OrderItem struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com OrderItem"`

	*SObject

	AttachedContentDocuments   *QueryResult      `xml:"AttachedContentDocuments,omitempty"`
	AvailableQuantity          float64           `xml:"AvailableQuantity,omitempty"`
	ChildOrderItems            *QueryResult      `xml:"ChildOrderItems,omitempty"`
	CombinedAttachments        *QueryResult      `xml:"CombinedAttachments,omitempty"`
	ContentDocumentLinks       *QueryResult      `xml:"ContentDocumentLinks,omitempty"`
	CreatedBy                  *User             `xml:"CreatedBy,omitempty"`
	CreatedById                *ID               `xml:"CreatedById,omitempty"`
	CreatedDate                time.Time         `xml:"CreatedDate,omitempty"`
	Description                string            `xml:"Description,omitempty"`
	EndDate                    time.Time         `xml:"EndDate,omitempty"`
	FeedSubscriptionsForEntity *QueryResult      `xml:"FeedSubscriptionsForEntity,omitempty"`
	Feeds                      *QueryResult      `xml:"Feeds,omitempty"`
	Histories                  *QueryResult      `xml:"Histories,omitempty"`
	IsDeleted                  bool              `xml:"IsDeleted,omitempty"`
	LastModifiedBy             *User             `xml:"LastModifiedBy,omitempty"`
	LastModifiedById           *ID               `xml:"LastModifiedById,omitempty"`
	LastModifiedDate           time.Time         `xml:"LastModifiedDate,omitempty"`
	ListPrice                  float64           `xml:"ListPrice,omitempty"`
	Order                      *Order            `xml:"Order,omitempty"`
	OrderId                    *ID               `xml:"OrderId,omitempty"`
	OrderItemNumber            string            `xml:"OrderItemNumber,omitempty"`
	OriginalOrderItem          *OrderItem        `xml:"OriginalOrderItem,omitempty"`
	OriginalOrderItemId        *ID               `xml:"OriginalOrderItemId,omitempty"`
	PricebookEntry             *PricebookEntry   `xml:"PricebookEntry,omitempty"`
	PricebookEntryId           *ID               `xml:"PricebookEntryId,omitempty"`
	Product2                   *Product2         `xml:"Product2,omitempty"`
	Product2Id                 *ID               `xml:"Product2Id,omitempty"`
	Quantity                   float64           `xml:"Quantity,omitempty"`
	ServiceDate                time.Time         `xml:"ServiceDate,omitempty"`
	SystemModstamp             time.Time         `xml:"SystemModstamp,omitempty"`
	TotalPrice                 float64           `xml:"TotalPrice,omitempty"`
	UnitPrice                  float64           `xml:"UnitPrice,omitempty"`
	UserRecordAccess           *UserRecordAccess `xml:"UserRecordAccess,omitempty"`
}

type OrderItemFeed struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com OrderItemFeed"`

	*SObject

	BestComment        *FeedComment `xml:"BestComment,omitempty"`
	BestCommentId      *ID          `xml:"BestCommentId,omitempty"`
	Body               string       `xml:"Body,omitempty"`
	CommentCount       int32        `xml:"CommentCount,omitempty"`
	CreatedBy          *SObject     `xml:"CreatedBy,omitempty"`
	CreatedById        *ID          `xml:"CreatedById,omitempty"`
	CreatedDate        time.Time    `xml:"CreatedDate,omitempty"`
	FeedAttachments    *QueryResult `xml:"FeedAttachments,omitempty"`
	FeedComments       *QueryResult `xml:"FeedComments,omitempty"`
	FeedLikes          *QueryResult `xml:"FeedLikes,omitempty"`
	FeedSignals        *QueryResult `xml:"FeedSignals,omitempty"`
	FeedTrackedChanges *QueryResult `xml:"FeedTrackedChanges,omitempty"`
	InsertedBy         *SObject     `xml:"InsertedBy,omitempty"`
	InsertedById       *ID          `xml:"InsertedById,omitempty"`
	IsDeleted          bool         `xml:"IsDeleted,omitempty"`
	IsRichText         bool         `xml:"IsRichText,omitempty"`
	LastModifiedDate   time.Time    `xml:"LastModifiedDate,omitempty"`
	LikeCount          int32        `xml:"LikeCount,omitempty"`
	LinkUrl            string       `xml:"LinkUrl,omitempty"`
	Parent             *OrderItem   `xml:"Parent,omitempty"`
	ParentId           *ID          `xml:"ParentId,omitempty"`
	RelatedRecordId    *ID          `xml:"RelatedRecordId,omitempty"`
	SystemModstamp     time.Time    `xml:"SystemModstamp,omitempty"`
	Title              string       `xml:"Title,omitempty"`
	Type               string       `xml:"Type,omitempty"`
}

type OrderItemHistory struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com OrderItemHistory"`

	*SObject

	CreatedBy   *SObject    `xml:"CreatedBy,omitempty"`
	CreatedById *ID         `xml:"CreatedById,omitempty"`
	CreatedDate time.Time   `xml:"CreatedDate,omitempty"`
	Field       string      `xml:"Field,omitempty"`
	IsDeleted   bool        `xml:"IsDeleted,omitempty"`
	NewValue    interface{} `xml:"NewValue,omitempty"`
	OldValue    interface{} `xml:"OldValue,omitempty"`
	OrderItem   *OrderItem  `xml:"OrderItem,omitempty"`
	OrderItemId *ID         `xml:"OrderItemId,omitempty"`
}

type OrderShare struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com OrderShare"`

	*SObject

	IsDeleted        bool      `xml:"IsDeleted,omitempty"`
	LastModifiedBy   *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time `xml:"LastModifiedDate,omitempty"`
	Order            *Order    `xml:"Order,omitempty"`
	OrderAccessLevel string    `xml:"OrderAccessLevel,omitempty"`
	OrderId          *ID       `xml:"OrderId,omitempty"`
	RowCause         string    `xml:"RowCause,omitempty"`
	UserOrGroup      *SObject  `xml:"UserOrGroup,omitempty"`
	UserOrGroupId    *ID       `xml:"UserOrGroupId,omitempty"`
}

type OrgDeleteRequest struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com OrgDeleteRequest"`

	*SObject

	CreatedBy        *User             `xml:"CreatedBy,omitempty"`
	CreatedById      *ID               `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time         `xml:"CreatedDate,omitempty"`
	IsDeleted        bool              `xml:"IsDeleted,omitempty"`
	LastModifiedBy   *User             `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID               `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time         `xml:"LastModifiedDate,omitempty"`
	Name             string            `xml:"Name,omitempty"`
	Owner            *SObject          `xml:"Owner,omitempty"`
	OwnerId          *ID               `xml:"OwnerId,omitempty"`
	ProcessInstances *QueryResult      `xml:"ProcessInstances,omitempty"`
	ProcessSteps     *QueryResult      `xml:"ProcessSteps,omitempty"`
	RequestType      string            `xml:"RequestType,omitempty"`
	SystemModstamp   time.Time         `xml:"SystemModstamp,omitempty"`
	UserRecordAccess *UserRecordAccess `xml:"UserRecordAccess,omitempty"`
}

type OrgDeleteRequestShare struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com OrgDeleteRequestShare"`

	*SObject

	AccessLevel      string            `xml:"AccessLevel,omitempty"`
	IsDeleted        bool              `xml:"IsDeleted,omitempty"`
	LastModifiedBy   *User             `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID               `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time         `xml:"LastModifiedDate,omitempty"`
	Parent           *OrgDeleteRequest `xml:"Parent,omitempty"`
	ParentId         *ID               `xml:"ParentId,omitempty"`
	RowCause         string            `xml:"RowCause,omitempty"`
	UserOrGroup      *SObject          `xml:"UserOrGroup,omitempty"`
	UserOrGroupId    *ID               `xml:"UserOrGroupId,omitempty"`
}

type OrgLifecycleNotification struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com OrgLifecycleNotification"`

	*SObject

	CreatedBy            *User     `xml:"CreatedBy,omitempty"`
	CreatedById          *ID       `xml:"CreatedById,omitempty"`
	CreatedDate          time.Time `xml:"CreatedDate,omitempty"`
	LifecycleRequestId   string    `xml:"LifecycleRequestId,omitempty"`
	LifecycleRequestType string    `xml:"LifecycleRequestType,omitempty"`
	OrgId                string    `xml:"OrgId,omitempty"`
	ReplayId             string    `xml:"ReplayId,omitempty"`
	Status               string    `xml:"Status,omitempty"`
	StatusCode           string    `xml:"StatusCode,omitempty"`
}

type OrgWideEmailAddress struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com OrgWideEmailAddress"`

	*SObject

	Address            string    `xml:"Address,omitempty"`
	CreatedBy          *User     `xml:"CreatedBy,omitempty"`
	CreatedById        *ID       `xml:"CreatedById,omitempty"`
	CreatedDate        time.Time `xml:"CreatedDate,omitempty"`
	DisplayName        string    `xml:"DisplayName,omitempty"`
	IsAllowAllProfiles bool      `xml:"IsAllowAllProfiles,omitempty"`
	LastModifiedBy     *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById   *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate   time.Time `xml:"LastModifiedDate,omitempty"`
	SystemModstamp     time.Time `xml:"SystemModstamp,omitempty"`
}

type Organization struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com Organization"`

	*SObject

	Address                                *Address     `xml:"Address,omitempty"`
	AttachedContentDocuments               *QueryResult `xml:"AttachedContentDocuments,omitempty"`
	City                                   string       `xml:"City,omitempty"`
	CombinedAttachments                    *QueryResult `xml:"CombinedAttachments,omitempty"`
	ComplianceBccEmail                     string       `xml:"ComplianceBccEmail,omitempty"`
	ContentDocumentLinks                   *QueryResult `xml:"ContentDocumentLinks,omitempty"`
	Country                                string       `xml:"Country,omitempty"`
	CreatedBy                              *User        `xml:"CreatedBy,omitempty"`
	CreatedById                            *ID          `xml:"CreatedById,omitempty"`
	CreatedDate                            time.Time    `xml:"CreatedDate,omitempty"`
	CustomBrands                           *QueryResult `xml:"CustomBrands,omitempty"`
	DefaultAccountAccess                   string       `xml:"DefaultAccountAccess,omitempty"`
	DefaultCalendarAccess                  string       `xml:"DefaultCalendarAccess,omitempty"`
	DefaultCampaignAccess                  string       `xml:"DefaultCampaignAccess,omitempty"`
	DefaultCaseAccess                      string       `xml:"DefaultCaseAccess,omitempty"`
	DefaultContactAccess                   string       `xml:"DefaultContactAccess,omitempty"`
	DefaultLeadAccess                      string       `xml:"DefaultLeadAccess,omitempty"`
	DefaultLocaleSidKey                    string       `xml:"DefaultLocaleSidKey,omitempty"`
	DefaultOpportunityAccess               string       `xml:"DefaultOpportunityAccess,omitempty"`
	DefaultPricebookAccess                 string       `xml:"DefaultPricebookAccess,omitempty"`
	Division                               string       `xml:"Division,omitempty"`
	Fax                                    string       `xml:"Fax,omitempty"`
	FiscalYearStartMonth                   int32        `xml:"FiscalYearStartMonth,omitempty"`
	GeocodeAccuracy                        string       `xml:"GeocodeAccuracy,omitempty"`
	InstanceName                           string       `xml:"InstanceName,omitempty"`
	IsReadOnly                             bool         `xml:"IsReadOnly,omitempty"`
	IsSandbox                              bool         `xml:"IsSandbox,omitempty"`
	LanguageLocaleKey                      string       `xml:"LanguageLocaleKey,omitempty"`
	LastModifiedBy                         *User        `xml:"LastModifiedBy,omitempty"`
	LastModifiedById                       *ID          `xml:"LastModifiedById,omitempty"`
	LastModifiedDate                       time.Time    `xml:"LastModifiedDate,omitempty"`
	Latitude                               float64      `xml:"Latitude,omitempty"`
	Longitude                              float64      `xml:"Longitude,omitempty"`
	MonthlyPageViewsEntitlement            int32        `xml:"MonthlyPageViewsEntitlement,omitempty"`
	MonthlyPageViewsUsed                   int32        `xml:"MonthlyPageViewsUsed,omitempty"`
	Name                                   string       `xml:"Name,omitempty"`
	NamespacePrefix                        string       `xml:"NamespacePrefix,omitempty"`
	NumKnowledgeService                    int32        `xml:"NumKnowledgeService,omitempty"`
	OrganizationType                       string       `xml:"OrganizationType,omitempty"`
	Phone                                  string       `xml:"Phone,omitempty"`
	PostalCode                             string       `xml:"PostalCode,omitempty"`
	PreferencesActivityAnalyticsEnabled    bool         `xml:"PreferencesActivityAnalyticsEnabled,omitempty"`
	PreferencesAutoSelectIndividualOnMerge bool         `xml:"PreferencesAutoSelectIndividualOnMerge,omitempty"`
	PreferencesConsentManagementEnabled    bool         `xml:"PreferencesConsentManagementEnabled,omitempty"`
	PreferencesIndividualAutoCreateEnabled bool         `xml:"PreferencesIndividualAutoCreateEnabled,omitempty"`
	PreferencesLightningLoginEnabled       bool         `xml:"PreferencesLightningLoginEnabled,omitempty"`
	PreferencesOnlyLLPermUserAllowed       bool         `xml:"PreferencesOnlyLLPermUserAllowed,omitempty"`
	PreferencesRequireOpportunityProducts  bool         `xml:"PreferencesRequireOpportunityProducts,omitempty"`
	PreferencesTerminateOldestSession      bool         `xml:"PreferencesTerminateOldestSession,omitempty"`
	PreferencesTransactionSecurityPolicy   bool         `xml:"PreferencesTransactionSecurityPolicy,omitempty"`
	PrimaryContact                         string       `xml:"PrimaryContact,omitempty"`
	ReceivesAdminInfoEmails                bool         `xml:"ReceivesAdminInfoEmails,omitempty"`
	ReceivesInfoEmails                     bool         `xml:"ReceivesInfoEmails,omitempty"`
	SignupCountryIsoCode                   string       `xml:"SignupCountryIsoCode,omitempty"`
	State                                  string       `xml:"State,omitempty"`
	Street                                 string       `xml:"Street,omitempty"`
	SystemModstamp                         time.Time    `xml:"SystemModstamp,omitempty"`
	TimeZoneSidKey                         string       `xml:"TimeZoneSidKey,omitempty"`
	TrialExpirationDate                    time.Time    `xml:"TrialExpirationDate,omitempty"`
	UiSkin                                 string       `xml:"UiSkin,omitempty"`
	UsesStartDateAsFiscalYearName          bool         `xml:"UsesStartDateAsFiscalYearName,omitempty"`
	WebToCaseDefaultOrigin                 string       `xml:"WebToCaseDefaultOrigin,omitempty"`
}

type OutgoingEmail struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com OutgoingEmail"`

	*SObject

	BccAddress             string       `xml:"BccAddress,omitempty"`
	BccIds                 []*ID        `xml:"BccIds,omitempty"`
	CcAddress              string       `xml:"CcAddress,omitempty"`
	CcIds                  []*ID        `xml:"CcIds,omitempty"`
	ContentDocumentIds     []*ID        `xml:"ContentDocumentIds,omitempty"`
	ContentDocumentLinks   *QueryResult `xml:"ContentDocumentLinks,omitempty"`
	EmailTemplate          *SObject     `xml:"EmailTemplate,omitempty"`
	EmailTemplateId        *ID          `xml:"EmailTemplateId,omitempty"`
	ExternalId             string       `xml:"ExternalId,omitempty"`
	HtmlBody               string       `xml:"HtmlBody,omitempty"`
	OutgoingEmailRelations *QueryResult `xml:"OutgoingEmailRelations,omitempty"`
	RelatedTo              *SObject     `xml:"RelatedTo,omitempty"`
	RelatedToId            *ID          `xml:"RelatedToId,omitempty"`
	Subject                string       `xml:"Subject,omitempty"`
	TextBody               string       `xml:"TextBody,omitempty"`
	ToAddress              string       `xml:"ToAddress,omitempty"`
	ToIds                  []*ID        `xml:"ToIds,omitempty"`
	ValidatedFromAddress   string       `xml:"ValidatedFromAddress,omitempty"`
	Who                    *SObject     `xml:"Who,omitempty"`
	WhoId                  *ID          `xml:"WhoId,omitempty"`
}

type OutgoingEmailRelation struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com OutgoingEmailRelation"`

	*SObject

	ExternalId      string         `xml:"ExternalId,omitempty"`
	OutgoingEmail   *OutgoingEmail `xml:"OutgoingEmail,omitempty"`
	OutgoingEmailId string         `xml:"OutgoingEmailId,omitempty"`
	Relation        *SObject       `xml:"Relation,omitempty"`
	RelationAddress string         `xml:"RelationAddress,omitempty"`
	RelationId      *ID            `xml:"RelationId,omitempty"`
}

type OwnedContentDocument struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com OwnedContentDocument"`

	*SObject

	ContentDocument        *ContentDocument `xml:"ContentDocument,omitempty"`
	ContentDocumentId      *ID              `xml:"ContentDocumentId,omitempty"`
	ContentSize            int32            `xml:"ContentSize,omitempty"`
	ContentUrl             string           `xml:"ContentUrl,omitempty"`
	CreatedBy              *User            `xml:"CreatedBy,omitempty"`
	CreatedById            *ID              `xml:"CreatedById,omitempty"`
	CreatedDate            time.Time        `xml:"CreatedDate,omitempty"`
	ExternalDataSourceName string           `xml:"ExternalDataSourceName,omitempty"`
	ExternalDataSourceType string           `xml:"ExternalDataSourceType,omitempty"`
	FileExtension          string           `xml:"FileExtension,omitempty"`
	FileType               string           `xml:"FileType,omitempty"`
	IsDeleted              bool             `xml:"IsDeleted,omitempty"`
	LastModifiedBy         *User            `xml:"LastModifiedBy,omitempty"`
	LastModifiedById       *ID              `xml:"LastModifiedById,omitempty"`
	LastModifiedDate       time.Time        `xml:"LastModifiedDate,omitempty"`
	Owner                  *User            `xml:"Owner,omitempty"`
	OwnerId                *ID              `xml:"OwnerId,omitempty"`
	Title                  string           `xml:"Title,omitempty"`
}

type OwnerChangeOptionInfo struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com OwnerChangeOptionInfo"`

	*SObject

	DefaultValue       bool                   `xml:"DefaultValue,omitempty"`
	DurableId          string                 `xml:"DurableId,omitempty"`
	EntityDefinition   *EntityDefinition      `xml:"EntityDefinition,omitempty"`
	EntityDefinitionId string                 `xml:"EntityDefinitionId,omitempty"`
	IsEditable         bool                   `xml:"IsEditable,omitempty"`
	Label              string                 `xml:"Label,omitempty"`
	Name               string                 `xml:"Name,omitempty"`
	Parent             *OwnerChangeOptionInfo `xml:"Parent,omitempty"`
	ParentId           string                 `xml:"ParentId,omitempty"`
}

type PackageLicense struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com PackageLicense"`

	*SObject

	AllowedLicenses  int32     `xml:"AllowedLicenses,omitempty"`
	CreatedDate      time.Time `xml:"CreatedDate,omitempty"`
	ExpirationDate   time.Time `xml:"ExpirationDate,omitempty"`
	IsProvisioned    bool      `xml:"IsProvisioned,omitempty"`
	LastModifiedDate time.Time `xml:"LastModifiedDate,omitempty"`
	NamespacePrefix  string    `xml:"NamespacePrefix,omitempty"`
	Status           string    `xml:"Status,omitempty"`
	SystemModstamp   time.Time `xml:"SystemModstamp,omitempty"`
	UsedLicenses     int32     `xml:"UsedLicenses,omitempty"`
}

type Partner struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com Partner"`

	*SObject

	AccountFrom      *Account     `xml:"AccountFrom,omitempty"`
	AccountFromId    *ID          `xml:"AccountFromId,omitempty"`
	AccountTo        *Account     `xml:"AccountTo,omitempty"`
	AccountToId      *ID          `xml:"AccountToId,omitempty"`
	CreatedBy        *User        `xml:"CreatedBy,omitempty"`
	CreatedById      *ID          `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time    `xml:"CreatedDate,omitempty"`
	IsDeleted        bool         `xml:"IsDeleted,omitempty"`
	IsPrimary        bool         `xml:"IsPrimary,omitempty"`
	LastModifiedBy   *User        `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID          `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time    `xml:"LastModifiedDate,omitempty"`
	Opportunity      *Opportunity `xml:"Opportunity,omitempty"`
	OpportunityId    *ID          `xml:"OpportunityId,omitempty"`
	ReversePartnerId *ID          `xml:"ReversePartnerId,omitempty"`
	Role             string       `xml:"Role,omitempty"`
	SystemModstamp   time.Time    `xml:"SystemModstamp,omitempty"`
}

type PartnerRole struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com PartnerRole"`

	*SObject

	ApiName          string    `xml:"ApiName,omitempty"`
	CreatedBy        *User     `xml:"CreatedBy,omitempty"`
	CreatedById      *ID       `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time `xml:"CreatedDate,omitempty"`
	LastModifiedBy   *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time `xml:"LastModifiedDate,omitempty"`
	MasterLabel      string    `xml:"MasterLabel,omitempty"`
	ReverseRole      string    `xml:"ReverseRole,omitempty"`
	SortOrder        int32     `xml:"SortOrder,omitempty"`
	SystemModstamp   time.Time `xml:"SystemModstamp,omitempty"`
}

type Period struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com Period"`

	*SObject

	EndDate              time.Time           `xml:"EndDate,omitempty"`
	FiscalYearSettings   *FiscalYearSettings `xml:"FiscalYearSettings,omitempty"`
	FiscalYearSettingsId *ID                 `xml:"FiscalYearSettingsId,omitempty"`
	FullyQualifiedLabel  string              `xml:"FullyQualifiedLabel,omitempty"`
	IsForecastPeriod     bool                `xml:"IsForecastPeriod,omitempty"`
	Number               int32               `xml:"Number,omitempty"`
	PeriodLabel          string              `xml:"PeriodLabel,omitempty"`
	QuarterLabel         string              `xml:"QuarterLabel,omitempty"`
	StartDate            time.Time           `xml:"StartDate,omitempty"`
	SystemModstamp       time.Time           `xml:"SystemModstamp,omitempty"`
	Type                 string              `xml:"Type,omitempty"`
}

type PermissionSet struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com PermissionSet"`

	*SObject

	Assignments                               *QueryResult `xml:"Assignments,omitempty"`
	CreatedBy                                 *User        `xml:"CreatedBy,omitempty"`
	CreatedById                               *ID          `xml:"CreatedById,omitempty"`
	CreatedDate                               time.Time    `xml:"CreatedDate,omitempty"`
	Description                               string       `xml:"Description,omitempty"`
	FieldPerms                                *QueryResult `xml:"FieldPerms,omitempty"`
	HasActivationRequired                     bool         `xml:"HasActivationRequired,omitempty"`
	IsCustom                                  bool         `xml:"IsCustom,omitempty"`
	IsOwnedByProfile                          bool         `xml:"IsOwnedByProfile,omitempty"`
	Label                                     string       `xml:"Label,omitempty"`
	LastModifiedBy                            *User        `xml:"LastModifiedBy,omitempty"`
	LastModifiedById                          *ID          `xml:"LastModifiedById,omitempty"`
	LastModifiedDate                          time.Time    `xml:"LastModifiedDate,omitempty"`
	License                                   *SObject     `xml:"License,omitempty"`
	LicenseId                                 *ID          `xml:"LicenseId,omitempty"`
	Name                                      string       `xml:"Name,omitempty"`
	NamespacePrefix                           string       `xml:"NamespacePrefix,omitempty"`
	ObjectPerms                               *QueryResult `xml:"ObjectPerms,omitempty"`
	PermissionsAccessCMC                      bool         `xml:"PermissionsAccessCMC,omitempty"`
	PermissionsActivateContract               bool         `xml:"PermissionsActivateContract,omitempty"`
	PermissionsActivateOrder                  bool         `xml:"PermissionsActivateOrder,omitempty"`
	PermissionsAddAnalyticsRemoteConnections  bool         `xml:"PermissionsAddAnalyticsRemoteConnections,omitempty"`
	PermissionsAddDirectMessageMembers        bool         `xml:"PermissionsAddDirectMessageMembers,omitempty"`
	PermissionsAllowEmailIC                   bool         `xml:"PermissionsAllowEmailIC,omitempty"`
	PermissionsAllowLightningLogin            bool         `xml:"PermissionsAllowLightningLogin,omitempty"`
	PermissionsAllowUniversalSearch           bool         `xml:"PermissionsAllowUniversalSearch,omitempty"`
	PermissionsAllowViewEditConvertedLeads    bool         `xml:"PermissionsAllowViewEditConvertedLeads,omitempty"`
	PermissionsAllowViewKnowledge             bool         `xml:"PermissionsAllowViewKnowledge,omitempty"`
	PermissionsApexRestServices               bool         `xml:"PermissionsApexRestServices,omitempty"`
	PermissionsApiEnabled                     bool         `xml:"PermissionsApiEnabled,omitempty"`
	PermissionsAssignPermissionSets           bool         `xml:"PermissionsAssignPermissionSets,omitempty"`
	PermissionsAssignTopics                   bool         `xml:"PermissionsAssignTopics,omitempty"`
	PermissionsAuthorApex                     bool         `xml:"PermissionsAuthorApex,omitempty"`
	PermissionsAutomaticActivityCapture       bool         `xml:"PermissionsAutomaticActivityCapture,omitempty"`
	PermissionsBulkApiHardDelete              bool         `xml:"PermissionsBulkApiHardDelete,omitempty"`
	PermissionsBulkMacrosAllowed              bool         `xml:"PermissionsBulkMacrosAllowed,omitempty"`
	PermissionsCampaignInfluence2             bool         `xml:"PermissionsCampaignInfluence2,omitempty"`
	PermissionsCanApproveFeedPost             bool         `xml:"PermissionsCanApproveFeedPost,omitempty"`
	PermissionsCanEditDataPrepRecipe          bool         `xml:"PermissionsCanEditDataPrepRecipe,omitempty"`
	PermissionsCanInsertFeedSystemFields      bool         `xml:"PermissionsCanInsertFeedSystemFields,omitempty"`
	PermissionsCanManageMaps                  bool         `xml:"PermissionsCanManageMaps,omitempty"`
	PermissionsCanUseNewDashboardBuilder      bool         `xml:"PermissionsCanUseNewDashboardBuilder,omitempty"`
	PermissionsCanVerifyComment               bool         `xml:"PermissionsCanVerifyComment,omitempty"`
	PermissionsChangeDashboardColors          bool         `xml:"PermissionsChangeDashboardColors,omitempty"`
	PermissionsChatterComposeUiCodesnippet    bool         `xml:"PermissionsChatterComposeUiCodesnippet,omitempty"`
	PermissionsChatterEditOwnPost             bool         `xml:"PermissionsChatterEditOwnPost,omitempty"`
	PermissionsChatterEditOwnRecordPost       bool         `xml:"PermissionsChatterEditOwnRecordPost,omitempty"`
	PermissionsChatterFileLink                bool         `xml:"PermissionsChatterFileLink,omitempty"`
	PermissionsChatterInternalUser            bool         `xml:"PermissionsChatterInternalUser,omitempty"`
	PermissionsChatterInviteExternalUsers     bool         `xml:"PermissionsChatterInviteExternalUsers,omitempty"`
	PermissionsChatterOwnGroups               bool         `xml:"PermissionsChatterOwnGroups,omitempty"`
	PermissionsCloseConversations             bool         `xml:"PermissionsCloseConversations,omitempty"`
	PermissionsConfigCustomRecs               bool         `xml:"PermissionsConfigCustomRecs,omitempty"`
	PermissionsConnectOrgToEnvironmentHub     bool         `xml:"PermissionsConnectOrgToEnvironmentHub,omitempty"`
	PermissionsContentAdministrator           bool         `xml:"PermissionsContentAdministrator,omitempty"`
	PermissionsContentWorkspaces              bool         `xml:"PermissionsContentWorkspaces,omitempty"`
	PermissionsConvertLeads                   bool         `xml:"PermissionsConvertLeads,omitempty"`
	PermissionsCreateCustomizeDashboards      bool         `xml:"PermissionsCreateCustomizeDashboards,omitempty"`
	PermissionsCreateCustomizeFilters         bool         `xml:"PermissionsCreateCustomizeFilters,omitempty"`
	PermissionsCreateCustomizeReports         bool         `xml:"PermissionsCreateCustomizeReports,omitempty"`
	PermissionsCreateDashboardFolders         bool         `xml:"PermissionsCreateDashboardFolders,omitempty"`
	PermissionsCreateLtngTempInPub            bool         `xml:"PermissionsCreateLtngTempInPub,omitempty"`
	PermissionsCreatePackaging                bool         `xml:"PermissionsCreatePackaging,omitempty"`
	PermissionsCreateReportFolders            bool         `xml:"PermissionsCreateReportFolders,omitempty"`
	PermissionsCreateReportInLightning        bool         `xml:"PermissionsCreateReportInLightning,omitempty"`
	PermissionsCreateTopics                   bool         `xml:"PermissionsCreateTopics,omitempty"`
	PermissionsCreateWorkspaces               bool         `xml:"PermissionsCreateWorkspaces,omitempty"`
	PermissionsCustomMobileAppsAccess         bool         `xml:"PermissionsCustomMobileAppsAccess,omitempty"`
	PermissionsCustomSidebarOnAllPages        bool         `xml:"PermissionsCustomSidebarOnAllPages,omitempty"`
	PermissionsCustomizeApplication           bool         `xml:"PermissionsCustomizeApplication,omitempty"`
	PermissionsDataExport                     bool         `xml:"PermissionsDataExport,omitempty"`
	PermissionsDelegatedTwoFactor             bool         `xml:"PermissionsDelegatedTwoFactor,omitempty"`
	PermissionsDeleteActivatedContract        bool         `xml:"PermissionsDeleteActivatedContract,omitempty"`
	PermissionsDeleteTopics                   bool         `xml:"PermissionsDeleteTopics,omitempty"`
	PermissionsDistributeFromPersWksp         bool         `xml:"PermissionsDistributeFromPersWksp,omitempty"`
	PermissionsEditActivatedOrders            bool         `xml:"PermissionsEditActivatedOrders,omitempty"`
	PermissionsEditBrandTemplates             bool         `xml:"PermissionsEditBrandTemplates,omitempty"`
	PermissionsEditCaseComments               bool         `xml:"PermissionsEditCaseComments,omitempty"`
	PermissionsEditEvent                      bool         `xml:"PermissionsEditEvent,omitempty"`
	PermissionsEditHtmlTemplates              bool         `xml:"PermissionsEditHtmlTemplates,omitempty"`
	PermissionsEditKnowledge                  bool         `xml:"PermissionsEditKnowledge,omitempty"`
	PermissionsEditMyDashboards               bool         `xml:"PermissionsEditMyDashboards,omitempty"`
	PermissionsEditMyReports                  bool         `xml:"PermissionsEditMyReports,omitempty"`
	PermissionsEditOppLineItemUnitPrice       bool         `xml:"PermissionsEditOppLineItemUnitPrice,omitempty"`
	PermissionsEditPublicDocuments            bool         `xml:"PermissionsEditPublicDocuments,omitempty"`
	PermissionsEditPublicFilters              bool         `xml:"PermissionsEditPublicFilters,omitempty"`
	PermissionsEditPublicTemplates            bool         `xml:"PermissionsEditPublicTemplates,omitempty"`
	PermissionsEditReadonlyFields             bool         `xml:"PermissionsEditReadonlyFields,omitempty"`
	PermissionsEditTask                       bool         `xml:"PermissionsEditTask,omitempty"`
	PermissionsEditTopics                     bool         `xml:"PermissionsEditTopics,omitempty"`
	PermissionsEmailAdministration            bool         `xml:"PermissionsEmailAdministration,omitempty"`
	PermissionsEmailMass                      bool         `xml:"PermissionsEmailMass,omitempty"`
	PermissionsEmailSingle                    bool         `xml:"PermissionsEmailSingle,omitempty"`
	PermissionsEmailTemplateManagement        bool         `xml:"PermissionsEmailTemplateManagement,omitempty"`
	PermissionsEnableCommunityAppLauncher     bool         `xml:"PermissionsEnableCommunityAppLauncher,omitempty"`
	PermissionsEnableNotifications            bool         `xml:"PermissionsEnableNotifications,omitempty"`
	PermissionsExportReport                   bool         `xml:"PermissionsExportReport,omitempty"`
	PermissionsFeedPinning                    bool         `xml:"PermissionsFeedPinning,omitempty"`
	PermissionsFlowUFLRequired                bool         `xml:"PermissionsFlowUFLRequired,omitempty"`
	PermissionsForceTwoFactor                 bool         `xml:"PermissionsForceTwoFactor,omitempty"`
	PermissionsGiveRecognitionBadge           bool         `xml:"PermissionsGiveRecognitionBadge,omitempty"`
	PermissionsGovernNetworks                 bool         `xml:"PermissionsGovernNetworks,omitempty"`
	PermissionsHideReadByList                 bool         `xml:"PermissionsHideReadByList,omitempty"`
	PermissionsIdentityConnect                bool         `xml:"PermissionsIdentityConnect,omitempty"`
	PermissionsIdentityEnabled                bool         `xml:"PermissionsIdentityEnabled,omitempty"`
	PermissionsImportCustomObjects            bool         `xml:"PermissionsImportCustomObjects,omitempty"`
	PermissionsImportLeads                    bool         `xml:"PermissionsImportLeads,omitempty"`
	PermissionsImportPersonal                 bool         `xml:"PermissionsImportPersonal,omitempty"`
	PermissionsInsightsAppAdmin               bool         `xml:"PermissionsInsightsAppAdmin,omitempty"`
	PermissionsInsightsAppDashboardEditor     bool         `xml:"PermissionsInsightsAppDashboardEditor,omitempty"`
	PermissionsInsightsAppEltEditor           bool         `xml:"PermissionsInsightsAppEltEditor,omitempty"`
	PermissionsInsightsAppUploadUser          bool         `xml:"PermissionsInsightsAppUploadUser,omitempty"`
	PermissionsInsightsAppUser                bool         `xml:"PermissionsInsightsAppUser,omitempty"`
	PermissionsInsightsCreateApplication      bool         `xml:"PermissionsInsightsCreateApplication,omitempty"`
	PermissionsInstallPackaging               bool         `xml:"PermissionsInstallPackaging,omitempty"`
	PermissionsIotUser                        bool         `xml:"PermissionsIotUser,omitempty"`
	PermissionsLightningConsoleAllowedForUser bool         `xml:"PermissionsLightningConsoleAllowedForUser,omitempty"`
	PermissionsLightningExperienceUser        bool         `xml:"PermissionsLightningExperienceUser,omitempty"`
	PermissionsListEmailSend                  bool         `xml:"PermissionsListEmailSend,omitempty"`
	PermissionsLtngPromoReserved01UserPerm    bool         `xml:"PermissionsLtngPromoReserved01UserPerm,omitempty"`
	PermissionsManageAnalyticSnapshots        bool         `xml:"PermissionsManageAnalyticSnapshots,omitempty"`
	PermissionsManageAuthProviders            bool         `xml:"PermissionsManageAuthProviders,omitempty"`
	PermissionsManageBusinessHourHolidays     bool         `xml:"PermissionsManageBusinessHourHolidays,omitempty"`
	PermissionsManageCallCenters              bool         `xml:"PermissionsManageCallCenters,omitempty"`
	PermissionsManageCases                    bool         `xml:"PermissionsManageCases,omitempty"`
	PermissionsManageCategories               bool         `xml:"PermissionsManageCategories,omitempty"`
	PermissionsManageCertificates             bool         `xml:"PermissionsManageCertificates,omitempty"`
	PermissionsManageChatterMessages          bool         `xml:"PermissionsManageChatterMessages,omitempty"`
	PermissionsManageContentPermissions       bool         `xml:"PermissionsManageContentPermissions,omitempty"`
	PermissionsManageContentProperties        bool         `xml:"PermissionsManageContentProperties,omitempty"`
	PermissionsManageContentTypes             bool         `xml:"PermissionsManageContentTypes,omitempty"`
	PermissionsManageCustomPermissions        bool         `xml:"PermissionsManageCustomPermissions,omitempty"`
	PermissionsManageCustomReportTypes        bool         `xml:"PermissionsManageCustomReportTypes,omitempty"`
	PermissionsManageDashbdsInPubFolders      bool         `xml:"PermissionsManageDashbdsInPubFolders,omitempty"`
	PermissionsManageDataCategories           bool         `xml:"PermissionsManageDataCategories,omitempty"`
	PermissionsManageDataIntegrations         bool         `xml:"PermissionsManageDataIntegrations,omitempty"`
	PermissionsManageDynamicDashboards        bool         `xml:"PermissionsManageDynamicDashboards,omitempty"`
	PermissionsManageEmailClientConfig        bool         `xml:"PermissionsManageEmailClientConfig,omitempty"`
	PermissionsManageEncryptionKeys           bool         `xml:"PermissionsManageEncryptionKeys,omitempty"`
	PermissionsManageExchangeConfig           bool         `xml:"PermissionsManageExchangeConfig,omitempty"`
	PermissionsManageHealthCheck              bool         `xml:"PermissionsManageHealthCheck,omitempty"`
	PermissionsManageInteraction              bool         `xml:"PermissionsManageInteraction,omitempty"`
	PermissionsManageInternalUsers            bool         `xml:"PermissionsManageInternalUsers,omitempty"`
	PermissionsManageIpAddresses              bool         `xml:"PermissionsManageIpAddresses,omitempty"`
	PermissionsManageKnowledge                bool         `xml:"PermissionsManageKnowledge,omitempty"`
	PermissionsManageKnowledgeImportExport    bool         `xml:"PermissionsManageKnowledgeImportExport,omitempty"`
	PermissionsManageLeads                    bool         `xml:"PermissionsManageLeads,omitempty"`
	PermissionsManageLoginAccessPolicies      bool         `xml:"PermissionsManageLoginAccessPolicies,omitempty"`
	PermissionsManageMobile                   bool         `xml:"PermissionsManageMobile,omitempty"`
	PermissionsManageNetworks                 bool         `xml:"PermissionsManageNetworks,omitempty"`
	PermissionsManagePasswordPolicies         bool         `xml:"PermissionsManagePasswordPolicies,omitempty"`
	PermissionsManageProfilesPermissionsets   bool         `xml:"PermissionsManageProfilesPermissionsets,omitempty"`
	PermissionsManagePvtRptsAndDashbds        bool         `xml:"PermissionsManagePvtRptsAndDashbds,omitempty"`
	PermissionsManageRemoteAccess             bool         `xml:"PermissionsManageRemoteAccess,omitempty"`
	PermissionsManageReportsInPubFolders      bool         `xml:"PermissionsManageReportsInPubFolders,omitempty"`
	PermissionsManageRoles                    bool         `xml:"PermissionsManageRoles,omitempty"`
	PermissionsManageSearchPromotionRules     bool         `xml:"PermissionsManageSearchPromotionRules,omitempty"`
	PermissionsManageSessionPermissionSets    bool         `xml:"PermissionsManageSessionPermissionSets,omitempty"`
	PermissionsManageSharing                  bool         `xml:"PermissionsManageSharing,omitempty"`
	PermissionsManageSolutions                bool         `xml:"PermissionsManageSolutions,omitempty"`
	PermissionsManageSurveys                  bool         `xml:"PermissionsManageSurveys,omitempty"`
	PermissionsManageSynonyms                 bool         `xml:"PermissionsManageSynonyms,omitempty"`
	PermissionsManageTemplatedApp             bool         `xml:"PermissionsManageTemplatedApp,omitempty"`
	PermissionsManageTwoFactor                bool         `xml:"PermissionsManageTwoFactor,omitempty"`
	PermissionsManageUnlistedGroups           bool         `xml:"PermissionsManageUnlistedGroups,omitempty"`
	PermissionsManageUsers                    bool         `xml:"PermissionsManageUsers,omitempty"`
	PermissionsMassInlineEdit                 bool         `xml:"PermissionsMassInlineEdit,omitempty"`
	PermissionsMergeTopics                    bool         `xml:"PermissionsMergeTopics,omitempty"`
	PermissionsModerateChatter                bool         `xml:"PermissionsModerateChatter,omitempty"`
	PermissionsModerateNetworkUsers           bool         `xml:"PermissionsModerateNetworkUsers,omitempty"`
	PermissionsModifyAllData                  bool         `xml:"PermissionsModifyAllData,omitempty"`
	PermissionsModifyMetadata                 bool         `xml:"PermissionsModifyMetadata,omitempty"`
	PermissionsModifySecureAgents             bool         `xml:"PermissionsModifySecureAgents,omitempty"`
	PermissionsNewReportBuilder               bool         `xml:"PermissionsNewReportBuilder,omitempty"`
	PermissionsPackaging2                     bool         `xml:"PermissionsPackaging2,omitempty"`
	PermissionsPasswordNeverExpires           bool         `xml:"PermissionsPasswordNeverExpires,omitempty"`
	PermissionsPreventClassicExperience       bool         `xml:"PermissionsPreventClassicExperience,omitempty"`
	PermissionsPrivacyDataAccess              bool         `xml:"PermissionsPrivacyDataAccess,omitempty"`
	PermissionsPublishPackaging               bool         `xml:"PermissionsPublishPackaging,omitempty"`
	PermissionsRemoveDirectMessageMembers     bool         `xml:"PermissionsRemoveDirectMessageMembers,omitempty"`
	PermissionsResetPasswords                 bool         `xml:"PermissionsResetPasswords,omitempty"`
	PermissionsRunFlow                        bool         `xml:"PermissionsRunFlow,omitempty"`
	PermissionsRunReports                     bool         `xml:"PermissionsRunReports,omitempty"`
	PermissionsSalesConsole                   bool         `xml:"PermissionsSalesConsole,omitempty"`
	PermissionsScheduleReports                bool         `xml:"PermissionsScheduleReports,omitempty"`
	PermissionsSelectFilesFromSalesforce      bool         `xml:"PermissionsSelectFilesFromSalesforce,omitempty"`
	PermissionsSendAnnouncementEmails         bool         `xml:"PermissionsSendAnnouncementEmails,omitempty"`
	PermissionsSendSitRequests                bool         `xml:"PermissionsSendSitRequests,omitempty"`
	PermissionsShareInternalArticles          bool         `xml:"PermissionsShareInternalArticles,omitempty"`
	PermissionsShowCompanyNameAsUserBadge     bool         `xml:"PermissionsShowCompanyNameAsUserBadge,omitempty"`
	PermissionsSolutionImport                 bool         `xml:"PermissionsSolutionImport,omitempty"`
	PermissionsStdAutomaticActivityCapture    bool         `xml:"PermissionsStdAutomaticActivityCapture,omitempty"`
	PermissionsSubmitMacrosAllowed            bool         `xml:"PermissionsSubmitMacrosAllowed,omitempty"`
	PermissionsSubscribeDashboardToOtherUsers bool         `xml:"PermissionsSubscribeDashboardToOtherUsers,omitempty"`
	PermissionsSubscribeReportToOtherUsers    bool         `xml:"PermissionsSubscribeReportToOtherUsers,omitempty"`
	PermissionsSubscribeReportsRunAsUser      bool         `xml:"PermissionsSubscribeReportsRunAsUser,omitempty"`
	PermissionsSubscribeToLightningDashboards bool         `xml:"PermissionsSubscribeToLightningDashboards,omitempty"`
	PermissionsSubscribeToLightningReports    bool         `xml:"PermissionsSubscribeToLightningReports,omitempty"`
	PermissionsTransferAnyCase                bool         `xml:"PermissionsTransferAnyCase,omitempty"`
	PermissionsTransferAnyEntity              bool         `xml:"PermissionsTransferAnyEntity,omitempty"`
	PermissionsTransferAnyLead                bool         `xml:"PermissionsTransferAnyLead,omitempty"`
	PermissionsTwoFactorApi                   bool         `xml:"PermissionsTwoFactorApi,omitempty"`
	PermissionsUseTeamReassignWizards         bool         `xml:"PermissionsUseTeamReassignWizards,omitempty"`
	PermissionsUseTemplatedApp                bool         `xml:"PermissionsUseTemplatedApp,omitempty"`
	PermissionsUseWebLink                     bool         `xml:"PermissionsUseWebLink,omitempty"`
	PermissionsViewAllActivities              bool         `xml:"PermissionsViewAllActivities,omitempty"`
	PermissionsViewAllData                    bool         `xml:"PermissionsViewAllData,omitempty"`
	PermissionsViewAllUsers                   bool         `xml:"PermissionsViewAllUsers,omitempty"`
	PermissionsViewContent                    bool         `xml:"PermissionsViewContent,omitempty"`
	PermissionsViewDataAssessment             bool         `xml:"PermissionsViewDataAssessment,omitempty"`
	PermissionsViewDataCategories             bool         `xml:"PermissionsViewDataCategories,omitempty"`
	PermissionsViewEncryptedData              bool         `xml:"PermissionsViewEncryptedData,omitempty"`
	PermissionsViewEventLogFiles              bool         `xml:"PermissionsViewEventLogFiles,omitempty"`
	PermissionsViewHealthCheck                bool         `xml:"PermissionsViewHealthCheck,omitempty"`
	PermissionsViewHelpLink                   bool         `xml:"PermissionsViewHelpLink,omitempty"`
	PermissionsViewMyTeamsDashboards          bool         `xml:"PermissionsViewMyTeamsDashboards,omitempty"`
	PermissionsViewOnlyEmbeddedAppUser        bool         `xml:"PermissionsViewOnlyEmbeddedAppUser,omitempty"`
	PermissionsViewPublicDashboards           bool         `xml:"PermissionsViewPublicDashboards,omitempty"`
	PermissionsViewPublicReports              bool         `xml:"PermissionsViewPublicReports,omitempty"`
	PermissionsViewRoles                      bool         `xml:"PermissionsViewRoles,omitempty"`
	PermissionsViewSetup                      bool         `xml:"PermissionsViewSetup,omitempty"`
	PermissionsWaveTabularDownload            bool         `xml:"PermissionsWaveTabularDownload,omitempty"`
	PermissionsWorkCalibrationUser            bool         `xml:"PermissionsWorkCalibrationUser,omitempty"`
	PermissionsWorkDotComUserPerm             bool         `xml:"PermissionsWorkDotComUserPerm,omitempty"`
	Profile                                   *Profile     `xml:"Profile,omitempty"`
	ProfileId                                 *ID          `xml:"ProfileId,omitempty"`
	SessionActivations                        *QueryResult `xml:"SessionActivations,omitempty"`
	SetupEntityAccessItems                    *QueryResult `xml:"SetupEntityAccessItems,omitempty"`
	SystemModstamp                            time.Time    `xml:"SystemModstamp,omitempty"`
}

type PermissionSetAssignment struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com PermissionSetAssignment"`

	*SObject

	Assignee        *User          `xml:"Assignee,omitempty"`
	AssigneeId      *ID            `xml:"AssigneeId,omitempty"`
	PermissionSet   *PermissionSet `xml:"PermissionSet,omitempty"`
	PermissionSetId *ID            `xml:"PermissionSetId,omitempty"`
	SystemModstamp  time.Time      `xml:"SystemModstamp,omitempty"`
}

type PermissionSetLicense struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com PermissionSetLicense"`

	*SObject

	CreatedBy                                        *User        `xml:"CreatedBy,omitempty"`
	CreatedById                                      *ID          `xml:"CreatedById,omitempty"`
	CreatedDate                                      time.Time    `xml:"CreatedDate,omitempty"`
	DeveloperName                                    string       `xml:"DeveloperName,omitempty"`
	ExpirationDate                                   time.Time    `xml:"ExpirationDate,omitempty"`
	GrantedByLicenses                                *QueryResult `xml:"GrantedByLicenses,omitempty"`
	IsDeleted                                        bool         `xml:"IsDeleted,omitempty"`
	Language                                         string       `xml:"Language,omitempty"`
	LastModifiedBy                                   *User        `xml:"LastModifiedBy,omitempty"`
	LastModifiedById                                 *ID          `xml:"LastModifiedById,omitempty"`
	LastModifiedDate                                 time.Time    `xml:"LastModifiedDate,omitempty"`
	MasterLabel                                      string       `xml:"MasterLabel,omitempty"`
	MaximumPermissionsAccessCMC                      bool         `xml:"MaximumPermissionsAccessCMC,omitempty"`
	MaximumPermissionsActivateContract               bool         `xml:"MaximumPermissionsActivateContract,omitempty"`
	MaximumPermissionsActivateOrder                  bool         `xml:"MaximumPermissionsActivateOrder,omitempty"`
	MaximumPermissionsAddAnalyticsRemoteConnections  bool         `xml:"MaximumPermissionsAddAnalyticsRemoteConnections,omitempty"`
	MaximumPermissionsAddDirectMessageMembers        bool         `xml:"MaximumPermissionsAddDirectMessageMembers,omitempty"`
	MaximumPermissionsAllowEmailIC                   bool         `xml:"MaximumPermissionsAllowEmailIC,omitempty"`
	MaximumPermissionsAllowLightningLogin            bool         `xml:"MaximumPermissionsAllowLightningLogin,omitempty"`
	MaximumPermissionsAllowUniversalSearch           bool         `xml:"MaximumPermissionsAllowUniversalSearch,omitempty"`
	MaximumPermissionsAllowViewEditConvertedLeads    bool         `xml:"MaximumPermissionsAllowViewEditConvertedLeads,omitempty"`
	MaximumPermissionsAllowViewKnowledge             bool         `xml:"MaximumPermissionsAllowViewKnowledge,omitempty"`
	MaximumPermissionsApexRestServices               bool         `xml:"MaximumPermissionsApexRestServices,omitempty"`
	MaximumPermissionsApiEnabled                     bool         `xml:"MaximumPermissionsApiEnabled,omitempty"`
	MaximumPermissionsAssignPermissionSets           bool         `xml:"MaximumPermissionsAssignPermissionSets,omitempty"`
	MaximumPermissionsAssignTopics                   bool         `xml:"MaximumPermissionsAssignTopics,omitempty"`
	MaximumPermissionsAuthorApex                     bool         `xml:"MaximumPermissionsAuthorApex,omitempty"`
	MaximumPermissionsAutomaticActivityCapture       bool         `xml:"MaximumPermissionsAutomaticActivityCapture,omitempty"`
	MaximumPermissionsBulkApiHardDelete              bool         `xml:"MaximumPermissionsBulkApiHardDelete,omitempty"`
	MaximumPermissionsBulkMacrosAllowed              bool         `xml:"MaximumPermissionsBulkMacrosAllowed,omitempty"`
	MaximumPermissionsCampaignInfluence2             bool         `xml:"MaximumPermissionsCampaignInfluence2,omitempty"`
	MaximumPermissionsCanApproveFeedPost             bool         `xml:"MaximumPermissionsCanApproveFeedPost,omitempty"`
	MaximumPermissionsCanEditDataPrepRecipe          bool         `xml:"MaximumPermissionsCanEditDataPrepRecipe,omitempty"`
	MaximumPermissionsCanInsertFeedSystemFields      bool         `xml:"MaximumPermissionsCanInsertFeedSystemFields,omitempty"`
	MaximumPermissionsCanManageMaps                  bool         `xml:"MaximumPermissionsCanManageMaps,omitempty"`
	MaximumPermissionsCanUseNewDashboardBuilder      bool         `xml:"MaximumPermissionsCanUseNewDashboardBuilder,omitempty"`
	MaximumPermissionsCanVerifyComment               bool         `xml:"MaximumPermissionsCanVerifyComment,omitempty"`
	MaximumPermissionsChangeDashboardColors          bool         `xml:"MaximumPermissionsChangeDashboardColors,omitempty"`
	MaximumPermissionsChatterComposeUiCodesnippet    bool         `xml:"MaximumPermissionsChatterComposeUiCodesnippet,omitempty"`
	MaximumPermissionsChatterEditOwnPost             bool         `xml:"MaximumPermissionsChatterEditOwnPost,omitempty"`
	MaximumPermissionsChatterEditOwnRecordPost       bool         `xml:"MaximumPermissionsChatterEditOwnRecordPost,omitempty"`
	MaximumPermissionsChatterFileLink                bool         `xml:"MaximumPermissionsChatterFileLink,omitempty"`
	MaximumPermissionsChatterInternalUser            bool         `xml:"MaximumPermissionsChatterInternalUser,omitempty"`
	MaximumPermissionsChatterInviteExternalUsers     bool         `xml:"MaximumPermissionsChatterInviteExternalUsers,omitempty"`
	MaximumPermissionsChatterOwnGroups               bool         `xml:"MaximumPermissionsChatterOwnGroups,omitempty"`
	MaximumPermissionsCloseConversations             bool         `xml:"MaximumPermissionsCloseConversations,omitempty"`
	MaximumPermissionsConfigCustomRecs               bool         `xml:"MaximumPermissionsConfigCustomRecs,omitempty"`
	MaximumPermissionsConnectOrgToEnvironmentHub     bool         `xml:"MaximumPermissionsConnectOrgToEnvironmentHub,omitempty"`
	MaximumPermissionsContentAdministrator           bool         `xml:"MaximumPermissionsContentAdministrator,omitempty"`
	MaximumPermissionsContentWorkspaces              bool         `xml:"MaximumPermissionsContentWorkspaces,omitempty"`
	MaximumPermissionsConvertLeads                   bool         `xml:"MaximumPermissionsConvertLeads,omitempty"`
	MaximumPermissionsCreateCustomizeDashboards      bool         `xml:"MaximumPermissionsCreateCustomizeDashboards,omitempty"`
	MaximumPermissionsCreateCustomizeFilters         bool         `xml:"MaximumPermissionsCreateCustomizeFilters,omitempty"`
	MaximumPermissionsCreateCustomizeReports         bool         `xml:"MaximumPermissionsCreateCustomizeReports,omitempty"`
	MaximumPermissionsCreateDashboardFolders         bool         `xml:"MaximumPermissionsCreateDashboardFolders,omitempty"`
	MaximumPermissionsCreateLtngTempInPub            bool         `xml:"MaximumPermissionsCreateLtngTempInPub,omitempty"`
	MaximumPermissionsCreatePackaging                bool         `xml:"MaximumPermissionsCreatePackaging,omitempty"`
	MaximumPermissionsCreateReportFolders            bool         `xml:"MaximumPermissionsCreateReportFolders,omitempty"`
	MaximumPermissionsCreateReportInLightning        bool         `xml:"MaximumPermissionsCreateReportInLightning,omitempty"`
	MaximumPermissionsCreateTopics                   bool         `xml:"MaximumPermissionsCreateTopics,omitempty"`
	MaximumPermissionsCreateWorkspaces               bool         `xml:"MaximumPermissionsCreateWorkspaces,omitempty"`
	MaximumPermissionsCustomMobileAppsAccess         bool         `xml:"MaximumPermissionsCustomMobileAppsAccess,omitempty"`
	MaximumPermissionsCustomSidebarOnAllPages        bool         `xml:"MaximumPermissionsCustomSidebarOnAllPages,omitempty"`
	MaximumPermissionsCustomizeApplication           bool         `xml:"MaximumPermissionsCustomizeApplication,omitempty"`
	MaximumPermissionsDataExport                     bool         `xml:"MaximumPermissionsDataExport,omitempty"`
	MaximumPermissionsDelegatedTwoFactor             bool         `xml:"MaximumPermissionsDelegatedTwoFactor,omitempty"`
	MaximumPermissionsDeleteActivatedContract        bool         `xml:"MaximumPermissionsDeleteActivatedContract,omitempty"`
	MaximumPermissionsDeleteTopics                   bool         `xml:"MaximumPermissionsDeleteTopics,omitempty"`
	MaximumPermissionsDistributeFromPersWksp         bool         `xml:"MaximumPermissionsDistributeFromPersWksp,omitempty"`
	MaximumPermissionsEditActivatedOrders            bool         `xml:"MaximumPermissionsEditActivatedOrders,omitempty"`
	MaximumPermissionsEditBrandTemplates             bool         `xml:"MaximumPermissionsEditBrandTemplates,omitempty"`
	MaximumPermissionsEditCaseComments               bool         `xml:"MaximumPermissionsEditCaseComments,omitempty"`
	MaximumPermissionsEditEvent                      bool         `xml:"MaximumPermissionsEditEvent,omitempty"`
	MaximumPermissionsEditHtmlTemplates              bool         `xml:"MaximumPermissionsEditHtmlTemplates,omitempty"`
	MaximumPermissionsEditKnowledge                  bool         `xml:"MaximumPermissionsEditKnowledge,omitempty"`
	MaximumPermissionsEditMyDashboards               bool         `xml:"MaximumPermissionsEditMyDashboards,omitempty"`
	MaximumPermissionsEditMyReports                  bool         `xml:"MaximumPermissionsEditMyReports,omitempty"`
	MaximumPermissionsEditOppLineItemUnitPrice       bool         `xml:"MaximumPermissionsEditOppLineItemUnitPrice,omitempty"`
	MaximumPermissionsEditPublicDocuments            bool         `xml:"MaximumPermissionsEditPublicDocuments,omitempty"`
	MaximumPermissionsEditPublicFilters              bool         `xml:"MaximumPermissionsEditPublicFilters,omitempty"`
	MaximumPermissionsEditPublicTemplates            bool         `xml:"MaximumPermissionsEditPublicTemplates,omitempty"`
	MaximumPermissionsEditReadonlyFields             bool         `xml:"MaximumPermissionsEditReadonlyFields,omitempty"`
	MaximumPermissionsEditTask                       bool         `xml:"MaximumPermissionsEditTask,omitempty"`
	MaximumPermissionsEditTopics                     bool         `xml:"MaximumPermissionsEditTopics,omitempty"`
	MaximumPermissionsEmailAdministration            bool         `xml:"MaximumPermissionsEmailAdministration,omitempty"`
	MaximumPermissionsEmailMass                      bool         `xml:"MaximumPermissionsEmailMass,omitempty"`
	MaximumPermissionsEmailSingle                    bool         `xml:"MaximumPermissionsEmailSingle,omitempty"`
	MaximumPermissionsEmailTemplateManagement        bool         `xml:"MaximumPermissionsEmailTemplateManagement,omitempty"`
	MaximumPermissionsEnableCommunityAppLauncher     bool         `xml:"MaximumPermissionsEnableCommunityAppLauncher,omitempty"`
	MaximumPermissionsEnableNotifications            bool         `xml:"MaximumPermissionsEnableNotifications,omitempty"`
	MaximumPermissionsExportReport                   bool         `xml:"MaximumPermissionsExportReport,omitempty"`
	MaximumPermissionsFeedPinning                    bool         `xml:"MaximumPermissionsFeedPinning,omitempty"`
	MaximumPermissionsFlowUFLRequired                bool         `xml:"MaximumPermissionsFlowUFLRequired,omitempty"`
	MaximumPermissionsForceTwoFactor                 bool         `xml:"MaximumPermissionsForceTwoFactor,omitempty"`
	MaximumPermissionsGiveRecognitionBadge           bool         `xml:"MaximumPermissionsGiveRecognitionBadge,omitempty"`
	MaximumPermissionsGovernNetworks                 bool         `xml:"MaximumPermissionsGovernNetworks,omitempty"`
	MaximumPermissionsHideReadByList                 bool         `xml:"MaximumPermissionsHideReadByList,omitempty"`
	MaximumPermissionsIdentityConnect                bool         `xml:"MaximumPermissionsIdentityConnect,omitempty"`
	MaximumPermissionsIdentityEnabled                bool         `xml:"MaximumPermissionsIdentityEnabled,omitempty"`
	MaximumPermissionsImportCustomObjects            bool         `xml:"MaximumPermissionsImportCustomObjects,omitempty"`
	MaximumPermissionsImportLeads                    bool         `xml:"MaximumPermissionsImportLeads,omitempty"`
	MaximumPermissionsImportPersonal                 bool         `xml:"MaximumPermissionsImportPersonal,omitempty"`
	MaximumPermissionsInsightsAppAdmin               bool         `xml:"MaximumPermissionsInsightsAppAdmin,omitempty"`
	MaximumPermissionsInsightsAppDashboardEditor     bool         `xml:"MaximumPermissionsInsightsAppDashboardEditor,omitempty"`
	MaximumPermissionsInsightsAppEltEditor           bool         `xml:"MaximumPermissionsInsightsAppEltEditor,omitempty"`
	MaximumPermissionsInsightsAppUploadUser          bool         `xml:"MaximumPermissionsInsightsAppUploadUser,omitempty"`
	MaximumPermissionsInsightsAppUser                bool         `xml:"MaximumPermissionsInsightsAppUser,omitempty"`
	MaximumPermissionsInsightsCreateApplication      bool         `xml:"MaximumPermissionsInsightsCreateApplication,omitempty"`
	MaximumPermissionsInstallPackaging               bool         `xml:"MaximumPermissionsInstallPackaging,omitempty"`
	MaximumPermissionsIotUser                        bool         `xml:"MaximumPermissionsIotUser,omitempty"`
	MaximumPermissionsLightningConsoleAllowedForUser bool         `xml:"MaximumPermissionsLightningConsoleAllowedForUser,omitempty"`
	MaximumPermissionsLightningExperienceUser        bool         `xml:"MaximumPermissionsLightningExperienceUser,omitempty"`
	MaximumPermissionsListEmailSend                  bool         `xml:"MaximumPermissionsListEmailSend,omitempty"`
	MaximumPermissionsLtngPromoReserved01UserPerm    bool         `xml:"MaximumPermissionsLtngPromoReserved01UserPerm,omitempty"`
	MaximumPermissionsManageAnalyticSnapshots        bool         `xml:"MaximumPermissionsManageAnalyticSnapshots,omitempty"`
	MaximumPermissionsManageAuthProviders            bool         `xml:"MaximumPermissionsManageAuthProviders,omitempty"`
	MaximumPermissionsManageBusinessHourHolidays     bool         `xml:"MaximumPermissionsManageBusinessHourHolidays,omitempty"`
	MaximumPermissionsManageCallCenters              bool         `xml:"MaximumPermissionsManageCallCenters,omitempty"`
	MaximumPermissionsManageCases                    bool         `xml:"MaximumPermissionsManageCases,omitempty"`
	MaximumPermissionsManageCategories               bool         `xml:"MaximumPermissionsManageCategories,omitempty"`
	MaximumPermissionsManageCertificates             bool         `xml:"MaximumPermissionsManageCertificates,omitempty"`
	MaximumPermissionsManageChatterMessages          bool         `xml:"MaximumPermissionsManageChatterMessages,omitempty"`
	MaximumPermissionsManageContentPermissions       bool         `xml:"MaximumPermissionsManageContentPermissions,omitempty"`
	MaximumPermissionsManageContentProperties        bool         `xml:"MaximumPermissionsManageContentProperties,omitempty"`
	MaximumPermissionsManageContentTypes             bool         `xml:"MaximumPermissionsManageContentTypes,omitempty"`
	MaximumPermissionsManageCustomPermissions        bool         `xml:"MaximumPermissionsManageCustomPermissions,omitempty"`
	MaximumPermissionsManageCustomReportTypes        bool         `xml:"MaximumPermissionsManageCustomReportTypes,omitempty"`
	MaximumPermissionsManageDashbdsInPubFolders      bool         `xml:"MaximumPermissionsManageDashbdsInPubFolders,omitempty"`
	MaximumPermissionsManageDataCategories           bool         `xml:"MaximumPermissionsManageDataCategories,omitempty"`
	MaximumPermissionsManageDataIntegrations         bool         `xml:"MaximumPermissionsManageDataIntegrations,omitempty"`
	MaximumPermissionsManageDynamicDashboards        bool         `xml:"MaximumPermissionsManageDynamicDashboards,omitempty"`
	MaximumPermissionsManageEmailClientConfig        bool         `xml:"MaximumPermissionsManageEmailClientConfig,omitempty"`
	MaximumPermissionsManageEncryptionKeys           bool         `xml:"MaximumPermissionsManageEncryptionKeys,omitempty"`
	MaximumPermissionsManageExchangeConfig           bool         `xml:"MaximumPermissionsManageExchangeConfig,omitempty"`
	MaximumPermissionsManageHealthCheck              bool         `xml:"MaximumPermissionsManageHealthCheck,omitempty"`
	MaximumPermissionsManageInteraction              bool         `xml:"MaximumPermissionsManageInteraction,omitempty"`
	MaximumPermissionsManageInternalUsers            bool         `xml:"MaximumPermissionsManageInternalUsers,omitempty"`
	MaximumPermissionsManageIpAddresses              bool         `xml:"MaximumPermissionsManageIpAddresses,omitempty"`
	MaximumPermissionsManageKnowledge                bool         `xml:"MaximumPermissionsManageKnowledge,omitempty"`
	MaximumPermissionsManageKnowledgeImportExport    bool         `xml:"MaximumPermissionsManageKnowledgeImportExport,omitempty"`
	MaximumPermissionsManageLeads                    bool         `xml:"MaximumPermissionsManageLeads,omitempty"`
	MaximumPermissionsManageLoginAccessPolicies      bool         `xml:"MaximumPermissionsManageLoginAccessPolicies,omitempty"`
	MaximumPermissionsManageMobile                   bool         `xml:"MaximumPermissionsManageMobile,omitempty"`
	MaximumPermissionsManageNetworks                 bool         `xml:"MaximumPermissionsManageNetworks,omitempty"`
	MaximumPermissionsManagePasswordPolicies         bool         `xml:"MaximumPermissionsManagePasswordPolicies,omitempty"`
	MaximumPermissionsManageProfilesPermissionsets   bool         `xml:"MaximumPermissionsManageProfilesPermissionsets,omitempty"`
	MaximumPermissionsManagePvtRptsAndDashbds        bool         `xml:"MaximumPermissionsManagePvtRptsAndDashbds,omitempty"`
	MaximumPermissionsManageRemoteAccess             bool         `xml:"MaximumPermissionsManageRemoteAccess,omitempty"`
	MaximumPermissionsManageReportsInPubFolders      bool         `xml:"MaximumPermissionsManageReportsInPubFolders,omitempty"`
	MaximumPermissionsManageRoles                    bool         `xml:"MaximumPermissionsManageRoles,omitempty"`
	MaximumPermissionsManageSearchPromotionRules     bool         `xml:"MaximumPermissionsManageSearchPromotionRules,omitempty"`
	MaximumPermissionsManageSessionPermissionSets    bool         `xml:"MaximumPermissionsManageSessionPermissionSets,omitempty"`
	MaximumPermissionsManageSharing                  bool         `xml:"MaximumPermissionsManageSharing,omitempty"`
	MaximumPermissionsManageSolutions                bool         `xml:"MaximumPermissionsManageSolutions,omitempty"`
	MaximumPermissionsManageSurveys                  bool         `xml:"MaximumPermissionsManageSurveys,omitempty"`
	MaximumPermissionsManageSynonyms                 bool         `xml:"MaximumPermissionsManageSynonyms,omitempty"`
	MaximumPermissionsManageTemplatedApp             bool         `xml:"MaximumPermissionsManageTemplatedApp,omitempty"`
	MaximumPermissionsManageTwoFactor                bool         `xml:"MaximumPermissionsManageTwoFactor,omitempty"`
	MaximumPermissionsManageUnlistedGroups           bool         `xml:"MaximumPermissionsManageUnlistedGroups,omitempty"`
	MaximumPermissionsManageUsers                    bool         `xml:"MaximumPermissionsManageUsers,omitempty"`
	MaximumPermissionsMassInlineEdit                 bool         `xml:"MaximumPermissionsMassInlineEdit,omitempty"`
	MaximumPermissionsMergeTopics                    bool         `xml:"MaximumPermissionsMergeTopics,omitempty"`
	MaximumPermissionsModerateChatter                bool         `xml:"MaximumPermissionsModerateChatter,omitempty"`
	MaximumPermissionsModerateNetworkUsers           bool         `xml:"MaximumPermissionsModerateNetworkUsers,omitempty"`
	MaximumPermissionsModifyAllData                  bool         `xml:"MaximumPermissionsModifyAllData,omitempty"`
	MaximumPermissionsModifyMetadata                 bool         `xml:"MaximumPermissionsModifyMetadata,omitempty"`
	MaximumPermissionsModifySecureAgents             bool         `xml:"MaximumPermissionsModifySecureAgents,omitempty"`
	MaximumPermissionsNewReportBuilder               bool         `xml:"MaximumPermissionsNewReportBuilder,omitempty"`
	MaximumPermissionsPackaging2                     bool         `xml:"MaximumPermissionsPackaging2,omitempty"`
	MaximumPermissionsPasswordNeverExpires           bool         `xml:"MaximumPermissionsPasswordNeverExpires,omitempty"`
	MaximumPermissionsPreventClassicExperience       bool         `xml:"MaximumPermissionsPreventClassicExperience,omitempty"`
	MaximumPermissionsPrivacyDataAccess              bool         `xml:"MaximumPermissionsPrivacyDataAccess,omitempty"`
	MaximumPermissionsPublishPackaging               bool         `xml:"MaximumPermissionsPublishPackaging,omitempty"`
	MaximumPermissionsRemoveDirectMessageMembers     bool         `xml:"MaximumPermissionsRemoveDirectMessageMembers,omitempty"`
	MaximumPermissionsResetPasswords                 bool         `xml:"MaximumPermissionsResetPasswords,omitempty"`
	MaximumPermissionsRunFlow                        bool         `xml:"MaximumPermissionsRunFlow,omitempty"`
	MaximumPermissionsRunReports                     bool         `xml:"MaximumPermissionsRunReports,omitempty"`
	MaximumPermissionsSalesConsole                   bool         `xml:"MaximumPermissionsSalesConsole,omitempty"`
	MaximumPermissionsScheduleReports                bool         `xml:"MaximumPermissionsScheduleReports,omitempty"`
	MaximumPermissionsSelectFilesFromSalesforce      bool         `xml:"MaximumPermissionsSelectFilesFromSalesforce,omitempty"`
	MaximumPermissionsSendAnnouncementEmails         bool         `xml:"MaximumPermissionsSendAnnouncementEmails,omitempty"`
	MaximumPermissionsSendSitRequests                bool         `xml:"MaximumPermissionsSendSitRequests,omitempty"`
	MaximumPermissionsShareInternalArticles          bool         `xml:"MaximumPermissionsShareInternalArticles,omitempty"`
	MaximumPermissionsShowCompanyNameAsUserBadge     bool         `xml:"MaximumPermissionsShowCompanyNameAsUserBadge,omitempty"`
	MaximumPermissionsSolutionImport                 bool         `xml:"MaximumPermissionsSolutionImport,omitempty"`
	MaximumPermissionsStdAutomaticActivityCapture    bool         `xml:"MaximumPermissionsStdAutomaticActivityCapture,omitempty"`
	MaximumPermissionsSubmitMacrosAllowed            bool         `xml:"MaximumPermissionsSubmitMacrosAllowed,omitempty"`
	MaximumPermissionsSubscribeDashboardToOtherUsers bool         `xml:"MaximumPermissionsSubscribeDashboardToOtherUsers,omitempty"`
	MaximumPermissionsSubscribeReportToOtherUsers    bool         `xml:"MaximumPermissionsSubscribeReportToOtherUsers,omitempty"`
	MaximumPermissionsSubscribeReportsRunAsUser      bool         `xml:"MaximumPermissionsSubscribeReportsRunAsUser,omitempty"`
	MaximumPermissionsSubscribeToLightningDashboards bool         `xml:"MaximumPermissionsSubscribeToLightningDashboards,omitempty"`
	MaximumPermissionsSubscribeToLightningReports    bool         `xml:"MaximumPermissionsSubscribeToLightningReports,omitempty"`
	MaximumPermissionsTransferAnyCase                bool         `xml:"MaximumPermissionsTransferAnyCase,omitempty"`
	MaximumPermissionsTransferAnyEntity              bool         `xml:"MaximumPermissionsTransferAnyEntity,omitempty"`
	MaximumPermissionsTransferAnyLead                bool         `xml:"MaximumPermissionsTransferAnyLead,omitempty"`
	MaximumPermissionsTwoFactorApi                   bool         `xml:"MaximumPermissionsTwoFactorApi,omitempty"`
	MaximumPermissionsUseTeamReassignWizards         bool         `xml:"MaximumPermissionsUseTeamReassignWizards,omitempty"`
	MaximumPermissionsUseTemplatedApp                bool         `xml:"MaximumPermissionsUseTemplatedApp,omitempty"`
	MaximumPermissionsUseWebLink                     bool         `xml:"MaximumPermissionsUseWebLink,omitempty"`
	MaximumPermissionsViewAllActivities              bool         `xml:"MaximumPermissionsViewAllActivities,omitempty"`
	MaximumPermissionsViewAllData                    bool         `xml:"MaximumPermissionsViewAllData,omitempty"`
	MaximumPermissionsViewAllUsers                   bool         `xml:"MaximumPermissionsViewAllUsers,omitempty"`
	MaximumPermissionsViewContent                    bool         `xml:"MaximumPermissionsViewContent,omitempty"`
	MaximumPermissionsViewDataAssessment             bool         `xml:"MaximumPermissionsViewDataAssessment,omitempty"`
	MaximumPermissionsViewDataCategories             bool         `xml:"MaximumPermissionsViewDataCategories,omitempty"`
	MaximumPermissionsViewEncryptedData              bool         `xml:"MaximumPermissionsViewEncryptedData,omitempty"`
	MaximumPermissionsViewEventLogFiles              bool         `xml:"MaximumPermissionsViewEventLogFiles,omitempty"`
	MaximumPermissionsViewHealthCheck                bool         `xml:"MaximumPermissionsViewHealthCheck,omitempty"`
	MaximumPermissionsViewHelpLink                   bool         `xml:"MaximumPermissionsViewHelpLink,omitempty"`
	MaximumPermissionsViewMyTeamsDashboards          bool         `xml:"MaximumPermissionsViewMyTeamsDashboards,omitempty"`
	MaximumPermissionsViewOnlyEmbeddedAppUser        bool         `xml:"MaximumPermissionsViewOnlyEmbeddedAppUser,omitempty"`
	MaximumPermissionsViewPublicDashboards           bool         `xml:"MaximumPermissionsViewPublicDashboards,omitempty"`
	MaximumPermissionsViewPublicReports              bool         `xml:"MaximumPermissionsViewPublicReports,omitempty"`
	MaximumPermissionsViewRoles                      bool         `xml:"MaximumPermissionsViewRoles,omitempty"`
	MaximumPermissionsViewSetup                      bool         `xml:"MaximumPermissionsViewSetup,omitempty"`
	MaximumPermissionsWaveTabularDownload            bool         `xml:"MaximumPermissionsWaveTabularDownload,omitempty"`
	MaximumPermissionsWorkCalibrationUser            bool         `xml:"MaximumPermissionsWorkCalibrationUser,omitempty"`
	MaximumPermissionsWorkDotComUserPerm             bool         `xml:"MaximumPermissionsWorkDotComUserPerm,omitempty"`
	PermissionSetLicenseAssignments                  *QueryResult `xml:"PermissionSetLicenseAssignments,omitempty"`
	PermissionSetLicenseKey                          string       `xml:"PermissionSetLicenseKey,omitempty"`
	Status                                           string       `xml:"Status,omitempty"`
	SystemModstamp                                   time.Time    `xml:"SystemModstamp,omitempty"`
	TotalLicenses                                    int32        `xml:"TotalLicenses,omitempty"`
	UsedLicenses                                     int32        `xml:"UsedLicenses,omitempty"`
}

type PermissionSetLicenseAssign struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com PermissionSetLicenseAssign"`

	*SObject

	Assignee               *User                 `xml:"Assignee,omitempty"`
	AssigneeId             *ID                   `xml:"AssigneeId,omitempty"`
	CreatedBy              *User                 `xml:"CreatedBy,omitempty"`
	CreatedById            *ID                   `xml:"CreatedById,omitempty"`
	CreatedDate            time.Time             `xml:"CreatedDate,omitempty"`
	IsDeleted              bool                  `xml:"IsDeleted,omitempty"`
	LastModifiedBy         *User                 `xml:"LastModifiedBy,omitempty"`
	LastModifiedById       *ID                   `xml:"LastModifiedById,omitempty"`
	LastModifiedDate       time.Time             `xml:"LastModifiedDate,omitempty"`
	PermissionSetLicense   *PermissionSetLicense `xml:"PermissionSetLicense,omitempty"`
	PermissionSetLicenseId *ID                   `xml:"PermissionSetLicenseId,omitempty"`
	SystemModstamp         time.Time             `xml:"SystemModstamp,omitempty"`
}

type PicklistValueInfo struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com PicklistValueInfo"`

	*SObject

	DurableId        string          `xml:"DurableId,omitempty"`
	EntityParticle   *EntityParticle `xml:"EntityParticle,omitempty"`
	EntityParticleId string          `xml:"EntityParticleId,omitempty"`
	IsActive         bool            `xml:"IsActive,omitempty"`
	IsDefaultValue   bool            `xml:"IsDefaultValue,omitempty"`
	Label            string          `xml:"Label,omitempty"`
	ValidFor         string          `xml:"ValidFor,omitempty"`
	Value            string          `xml:"Value,omitempty"`
}

type PlatformAction struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com PlatformAction"`

	*SObject

	ActionListContext   string    `xml:"ActionListContext,omitempty"`
	ActionTarget        string    `xml:"ActionTarget,omitempty"`
	ActionTargetType    string    `xml:"ActionTargetType,omitempty"`
	ApiName             string    `xml:"ApiName,omitempty"`
	Category            string    `xml:"Category,omitempty"`
	ConfirmationMessage string    `xml:"ConfirmationMessage,omitempty"`
	DeviceFormat        string    `xml:"DeviceFormat,omitempty"`
	ExternalId          string    `xml:"ExternalId,omitempty"`
	GroupId             *ID       `xml:"GroupId,omitempty"`
	IconContentType     string    `xml:"IconContentType,omitempty"`
	IconHeight          int32     `xml:"IconHeight,omitempty"`
	IconUrl             string    `xml:"IconUrl,omitempty"`
	IconWidth           int32     `xml:"IconWidth,omitempty"`
	InvocationStatus    string    `xml:"InvocationStatus,omitempty"`
	InvokedByUser       *User     `xml:"InvokedByUser,omitempty"`
	InvokedByUserId     *ID       `xml:"InvokedByUserId,omitempty"`
	IsGroupDefault      bool      `xml:"IsGroupDefault,omitempty"`
	IsMassAction        bool      `xml:"IsMassAction,omitempty"`
	Label               string    `xml:"Label,omitempty"`
	LastModifiedDate    time.Time `xml:"LastModifiedDate,omitempty"`
	PrimaryColor        string    `xml:"PrimaryColor,omitempty"`
	RelatedListRecordId string    `xml:"RelatedListRecordId,omitempty"`
	RelatedSourceEntity string    `xml:"RelatedSourceEntity,omitempty"`
	Section             string    `xml:"Section,omitempty"`
	SourceEntity        string    `xml:"SourceEntity,omitempty"`
	Subtype             string    `xml:"Subtype,omitempty"`
	TargetObject        string    `xml:"TargetObject,omitempty"`
	TargetUrl           string    `xml:"TargetUrl,omitempty"`
	Type                string    `xml:"Type,omitempty"`
}

type PlatformCachePartition struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com PlatformCachePartition"`

	*SObject

	CreatedBy                  *User        `xml:"CreatedBy,omitempty"`
	CreatedById                *ID          `xml:"CreatedById,omitempty"`
	CreatedDate                time.Time    `xml:"CreatedDate,omitempty"`
	Description                string       `xml:"Description,omitempty"`
	DeveloperName              string       `xml:"DeveloperName,omitempty"`
	IsDefaultPartition         bool         `xml:"IsDefaultPartition,omitempty"`
	IsDeleted                  bool         `xml:"IsDeleted,omitempty"`
	Language                   string       `xml:"Language,omitempty"`
	LastModifiedBy             *User        `xml:"LastModifiedBy,omitempty"`
	LastModifiedById           *ID          `xml:"LastModifiedById,omitempty"`
	LastModifiedDate           time.Time    `xml:"LastModifiedDate,omitempty"`
	MasterLabel                string       `xml:"MasterLabel,omitempty"`
	NamespacePrefix            string       `xml:"NamespacePrefix,omitempty"`
	PlatforCachePartitionTypes *QueryResult `xml:"PlatforCachePartitionTypes,omitempty"`
	SystemModstamp             time.Time    `xml:"SystemModstamp,omitempty"`
}

type PlatformCachePartitionType struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com PlatformCachePartitionType"`

	*SObject

	AllocatedCapacity          int32                   `xml:"AllocatedCapacity,omitempty"`
	AllocatedPurchasedCapacity int32                   `xml:"AllocatedPurchasedCapacity,omitempty"`
	AllocatedTrialCapacity     int32                   `xml:"AllocatedTrialCapacity,omitempty"`
	CacheType                  string                  `xml:"CacheType,omitempty"`
	CreatedBy                  *User                   `xml:"CreatedBy,omitempty"`
	CreatedById                *ID                     `xml:"CreatedById,omitempty"`
	CreatedDate                time.Time               `xml:"CreatedDate,omitempty"`
	IsDeleted                  bool                    `xml:"IsDeleted,omitempty"`
	LastModifiedBy             *User                   `xml:"LastModifiedBy,omitempty"`
	LastModifiedById           *ID                     `xml:"LastModifiedById,omitempty"`
	LastModifiedDate           time.Time               `xml:"LastModifiedDate,omitempty"`
	PlatformCachePartition     *PlatformCachePartition `xml:"PlatformCachePartition,omitempty"`
	PlatformCachePartitionId   *ID                     `xml:"PlatformCachePartitionId,omitempty"`
	SystemModstamp             time.Time               `xml:"SystemModstamp,omitempty"`
}

type Pricebook2 struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com Pricebook2"`

	*SObject

	Contracts          *QueryResult      `xml:"Contracts,omitempty"`
	CreatedBy          *User             `xml:"CreatedBy,omitempty"`
	CreatedById        *ID               `xml:"CreatedById,omitempty"`
	CreatedDate        time.Time         `xml:"CreatedDate,omitempty"`
	Description        string            `xml:"Description,omitempty"`
	Histories          *QueryResult      `xml:"Histories,omitempty"`
	IsActive           bool              `xml:"IsActive,omitempty"`
	IsArchived         bool              `xml:"IsArchived,omitempty"`
	IsDeleted          bool              `xml:"IsDeleted,omitempty"`
	IsStandard         bool              `xml:"IsStandard,omitempty"`
	LastModifiedBy     *User             `xml:"LastModifiedBy,omitempty"`
	LastModifiedById   *ID               `xml:"LastModifiedById,omitempty"`
	LastModifiedDate   time.Time         `xml:"LastModifiedDate,omitempty"`
	LastReferencedDate time.Time         `xml:"LastReferencedDate,omitempty"`
	LastViewedDate     time.Time         `xml:"LastViewedDate,omitempty"`
	Name               string            `xml:"Name,omitempty"`
	Opportunities      *QueryResult      `xml:"Opportunities,omitempty"`
	Orders             *QueryResult      `xml:"Orders,omitempty"`
	PricebookEntries   *QueryResult      `xml:"PricebookEntries,omitempty"`
	SystemModstamp     time.Time         `xml:"SystemModstamp,omitempty"`
	UserRecordAccess   *UserRecordAccess `xml:"UserRecordAccess,omitempty"`
}

type Pricebook2History struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com Pricebook2History"`

	*SObject

	CreatedBy    *SObject    `xml:"CreatedBy,omitempty"`
	CreatedById  *ID         `xml:"CreatedById,omitempty"`
	CreatedDate  time.Time   `xml:"CreatedDate,omitempty"`
	Field        string      `xml:"Field,omitempty"`
	IsDeleted    bool        `xml:"IsDeleted,omitempty"`
	NewValue     interface{} `xml:"NewValue,omitempty"`
	OldValue     interface{} `xml:"OldValue,omitempty"`
	Pricebook2   *Pricebook2 `xml:"Pricebook2,omitempty"`
	Pricebook2Id *ID         `xml:"Pricebook2Id,omitempty"`
}

type PricebookEntry struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com PricebookEntry"`

	*SObject

	CreatedBy            *User        `xml:"CreatedBy,omitempty"`
	CreatedById          *ID          `xml:"CreatedById,omitempty"`
	CreatedDate          time.Time    `xml:"CreatedDate,omitempty"`
	IsActive             bool         `xml:"IsActive,omitempty"`
	IsDeleted            bool         `xml:"IsDeleted,omitempty"`
	LastModifiedBy       *User        `xml:"LastModifiedBy,omitempty"`
	LastModifiedById     *ID          `xml:"LastModifiedById,omitempty"`
	LastModifiedDate     time.Time    `xml:"LastModifiedDate,omitempty"`
	Name                 string       `xml:"Name,omitempty"`
	OpportunityLineItems *QueryResult `xml:"OpportunityLineItems,omitempty"`
	OrderItems           *QueryResult `xml:"OrderItems,omitempty"`
	Pricebook2           *Pricebook2  `xml:"Pricebook2,omitempty"`
	Pricebook2Id         *ID          `xml:"Pricebook2Id,omitempty"`
	Product2             *Product2    `xml:"Product2,omitempty"`
	Product2Id           *ID          `xml:"Product2Id,omitempty"`
	ProductCode          string       `xml:"ProductCode,omitempty"`
	SystemModstamp       time.Time    `xml:"SystemModstamp,omitempty"`
	UnitPrice            float64      `xml:"UnitPrice,omitempty"`
	UseStandardPrice     bool         `xml:"UseStandardPrice,omitempty"`
}

type ProcessDefinition struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ProcessDefinition"`

	*SObject

	CreatedBy        *User     `xml:"CreatedBy,omitempty"`
	CreatedById      *ID       `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time `xml:"CreatedDate,omitempty"`
	Description      string    `xml:"Description,omitempty"`
	DeveloperName    string    `xml:"DeveloperName,omitempty"`
	LastModifiedBy   *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time `xml:"LastModifiedDate,omitempty"`
	LockType         string    `xml:"LockType,omitempty"`
	Name             string    `xml:"Name,omitempty"`
	State            string    `xml:"State,omitempty"`
	SystemModstamp   time.Time `xml:"SystemModstamp,omitempty"`
	TableEnumOrId    string    `xml:"TableEnumOrId,omitempty"`
	Type             string    `xml:"Type,omitempty"`
}

type ProcessInstance struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ProcessInstance"`

	*SObject

	CompletedDate        time.Time          `xml:"CompletedDate,omitempty"`
	CreatedBy            *User              `xml:"CreatedBy,omitempty"`
	CreatedById          *ID                `xml:"CreatedById,omitempty"`
	CreatedDate          time.Time          `xml:"CreatedDate,omitempty"`
	ElapsedTimeInDays    float64            `xml:"ElapsedTimeInDays,omitempty"`
	ElapsedTimeInHours   float64            `xml:"ElapsedTimeInHours,omitempty"`
	ElapsedTimeInMinutes float64            `xml:"ElapsedTimeInMinutes,omitempty"`
	IsDeleted            bool               `xml:"IsDeleted,omitempty"`
	LastActor            *User              `xml:"LastActor,omitempty"`
	LastActorId          *ID                `xml:"LastActorId,omitempty"`
	LastModifiedBy       *User              `xml:"LastModifiedBy,omitempty"`
	LastModifiedById     *ID                `xml:"LastModifiedById,omitempty"`
	LastModifiedDate     time.Time          `xml:"LastModifiedDate,omitempty"`
	Nodes                *QueryResult       `xml:"Nodes,omitempty"`
	ProcessDefinition    *ProcessDefinition `xml:"ProcessDefinition,omitempty"`
	ProcessDefinitionId  *ID                `xml:"ProcessDefinitionId,omitempty"`
	Status               string             `xml:"Status,omitempty"`
	Steps                *QueryResult       `xml:"Steps,omitempty"`
	StepsAndWorkitems    *QueryResult       `xml:"StepsAndWorkitems,omitempty"`
	SubmittedBy          *User              `xml:"SubmittedBy,omitempty"`
	SubmittedById        *ID                `xml:"SubmittedById,omitempty"`
	SystemModstamp       time.Time          `xml:"SystemModstamp,omitempty"`
	TargetObject         *SObject           `xml:"TargetObject,omitempty"`
	TargetObjectId       *ID                `xml:"TargetObjectId,omitempty"`
	Workitems            *QueryResult       `xml:"Workitems,omitempty"`
}

type ProcessInstanceHistory struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ProcessInstanceHistory"`

	*SObject

	Actor                *SObject         `xml:"Actor,omitempty"`
	ActorId              *ID              `xml:"ActorId,omitempty"`
	Comments             string           `xml:"Comments,omitempty"`
	CreatedBy            *User            `xml:"CreatedBy,omitempty"`
	CreatedById          *ID              `xml:"CreatedById,omitempty"`
	CreatedDate          time.Time        `xml:"CreatedDate,omitempty"`
	ElapsedTimeInDays    float64          `xml:"ElapsedTimeInDays,omitempty"`
	ElapsedTimeInHours   float64          `xml:"ElapsedTimeInHours,omitempty"`
	ElapsedTimeInMinutes float64          `xml:"ElapsedTimeInMinutes,omitempty"`
	IsDeleted            bool             `xml:"IsDeleted,omitempty"`
	IsPending            bool             `xml:"IsPending,omitempty"`
	OriginalActor        *SObject         `xml:"OriginalActor,omitempty"`
	OriginalActorId      *ID              `xml:"OriginalActorId,omitempty"`
	ProcessInstance      *ProcessInstance `xml:"ProcessInstance,omitempty"`
	ProcessInstanceId    *ID              `xml:"ProcessInstanceId,omitempty"`
	ProcessNode          *ProcessNode     `xml:"ProcessNode,omitempty"`
	ProcessNodeId        *ID              `xml:"ProcessNodeId,omitempty"`
	RemindersSent        int32            `xml:"RemindersSent,omitempty"`
	StepStatus           string           `xml:"StepStatus,omitempty"`
	SystemModstamp       time.Time        `xml:"SystemModstamp,omitempty"`
	TargetObject         *SObject         `xml:"TargetObject,omitempty"`
	TargetObjectId       *ID              `xml:"TargetObjectId,omitempty"`
}

type ProcessInstanceNode struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ProcessInstanceNode"`

	*SObject

	CompletedDate        time.Time         `xml:"CompletedDate,omitempty"`
	CreatedBy            *User             `xml:"CreatedBy,omitempty"`
	CreatedById          *ID               `xml:"CreatedById,omitempty"`
	CreatedDate          time.Time         `xml:"CreatedDate,omitempty"`
	ElapsedTimeInDays    float64           `xml:"ElapsedTimeInDays,omitempty"`
	ElapsedTimeInHours   float64           `xml:"ElapsedTimeInHours,omitempty"`
	ElapsedTimeInMinutes float64           `xml:"ElapsedTimeInMinutes,omitempty"`
	IsDeleted            bool              `xml:"IsDeleted,omitempty"`
	LastActor            *User             `xml:"LastActor,omitempty"`
	LastActorId          *ID               `xml:"LastActorId,omitempty"`
	LastModifiedBy       *User             `xml:"LastModifiedBy,omitempty"`
	LastModifiedById     *ID               `xml:"LastModifiedById,omitempty"`
	LastModifiedDate     time.Time         `xml:"LastModifiedDate,omitempty"`
	NodeStatus           string            `xml:"NodeStatus,omitempty"`
	ProcessInstance      *ProcessInstance  `xml:"ProcessInstance,omitempty"`
	ProcessInstanceId    *ID               `xml:"ProcessInstanceId,omitempty"`
	ProcessNode          *ProcessNode      `xml:"ProcessNode,omitempty"`
	ProcessNodeId        *ID               `xml:"ProcessNodeId,omitempty"`
	ProcessNodeName      string            `xml:"ProcessNodeName,omitempty"`
	SystemModstamp       time.Time         `xml:"SystemModstamp,omitempty"`
	UserRecordAccess     *UserRecordAccess `xml:"UserRecordAccess,omitempty"`
}

type ProcessInstanceStep struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ProcessInstanceStep"`

	*SObject

	Actor                *SObject         `xml:"Actor,omitempty"`
	ActorId              *ID              `xml:"ActorId,omitempty"`
	Comments             string           `xml:"Comments,omitempty"`
	CreatedBy            *User            `xml:"CreatedBy,omitempty"`
	CreatedById          *ID              `xml:"CreatedById,omitempty"`
	CreatedDate          time.Time        `xml:"CreatedDate,omitempty"`
	ElapsedTimeInDays    float64          `xml:"ElapsedTimeInDays,omitempty"`
	ElapsedTimeInHours   float64          `xml:"ElapsedTimeInHours,omitempty"`
	ElapsedTimeInMinutes float64          `xml:"ElapsedTimeInMinutes,omitempty"`
	OriginalActor        *SObject         `xml:"OriginalActor,omitempty"`
	OriginalActorId      *ID              `xml:"OriginalActorId,omitempty"`
	ProcessInstance      *ProcessInstance `xml:"ProcessInstance,omitempty"`
	ProcessInstanceId    *ID              `xml:"ProcessInstanceId,omitempty"`
	StepNodeId           *ID              `xml:"StepNodeId,omitempty"`
	StepStatus           string           `xml:"StepStatus,omitempty"`
	SystemModstamp       time.Time        `xml:"SystemModstamp,omitempty"`
}

type ProcessInstanceWorkitem struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ProcessInstanceWorkitem"`

	*SObject

	Actor                *SObject         `xml:"Actor,omitempty"`
	ActorId              *ID              `xml:"ActorId,omitempty"`
	CreatedBy            *User            `xml:"CreatedBy,omitempty"`
	CreatedById          *ID              `xml:"CreatedById,omitempty"`
	CreatedDate          time.Time        `xml:"CreatedDate,omitempty"`
	ElapsedTimeInDays    float64          `xml:"ElapsedTimeInDays,omitempty"`
	ElapsedTimeInHours   float64          `xml:"ElapsedTimeInHours,omitempty"`
	ElapsedTimeInMinutes float64          `xml:"ElapsedTimeInMinutes,omitempty"`
	IsDeleted            bool             `xml:"IsDeleted,omitempty"`
	OriginalActor        *SObject         `xml:"OriginalActor,omitempty"`
	OriginalActorId      *ID              `xml:"OriginalActorId,omitempty"`
	ProcessInstance      *ProcessInstance `xml:"ProcessInstance,omitempty"`
	ProcessInstanceId    *ID              `xml:"ProcessInstanceId,omitempty"`
	SystemModstamp       time.Time        `xml:"SystemModstamp,omitempty"`
}

type ProcessNode struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ProcessNode"`

	*SObject

	Description         string             `xml:"Description,omitempty"`
	DeveloperName       string             `xml:"DeveloperName,omitempty"`
	Name                string             `xml:"Name,omitempty"`
	ProcessDefinition   *ProcessDefinition `xml:"ProcessDefinition,omitempty"`
	ProcessDefinitionId *ID                `xml:"ProcessDefinitionId,omitempty"`
	SystemModstamp      time.Time          `xml:"SystemModstamp,omitempty"`
}

type Product2 struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com Product2"`

	*SObject

	ActivityHistories          *QueryResult      `xml:"ActivityHistories,omitempty"`
	Assets                     *QueryResult      `xml:"Assets,omitempty"`
	AttachedContentDocuments   *QueryResult      `xml:"AttachedContentDocuments,omitempty"`
	Attachments                *QueryResult      `xml:"Attachments,omitempty"`
	CombinedAttachments        *QueryResult      `xml:"CombinedAttachments,omitempty"`
	ContentDocumentLinks       *QueryResult      `xml:"ContentDocumentLinks,omitempty"`
	CreatedBy                  *User             `xml:"CreatedBy,omitempty"`
	CreatedById                *ID               `xml:"CreatedById,omitempty"`
	CreatedDate                time.Time         `xml:"CreatedDate,omitempty"`
	Description                string            `xml:"Description,omitempty"`
	DisplayUrl                 string            `xml:"DisplayUrl,omitempty"`
	Emails                     *QueryResult      `xml:"Emails,omitempty"`
	Events                     *QueryResult      `xml:"Events,omitempty"`
	ExternalDataSourceId       *ID               `xml:"ExternalDataSourceId,omitempty"`
	ExternalId                 string            `xml:"ExternalId,omitempty"`
	Family                     string            `xml:"Family,omitempty"`
	FeedSubscriptionsForEntity *QueryResult      `xml:"FeedSubscriptionsForEntity,omitempty"`
	Feeds                      *QueryResult      `xml:"Feeds,omitempty"`
	Histories                  *QueryResult      `xml:"Histories,omitempty"`
	IsActive                   bool              `xml:"IsActive,omitempty"`
	IsArchived                 bool              `xml:"IsArchived,omitempty"`
	IsDeleted                  bool              `xml:"IsDeleted,omitempty"`
	LastModifiedBy             *User             `xml:"LastModifiedBy,omitempty"`
	LastModifiedById           *ID               `xml:"LastModifiedById,omitempty"`
	LastModifiedDate           time.Time         `xml:"LastModifiedDate,omitempty"`
	LastReferencedDate         time.Time         `xml:"LastReferencedDate,omitempty"`
	LastViewedDate             time.Time         `xml:"LastViewedDate,omitempty"`
	LookedUpFromActivities     *QueryResult      `xml:"LookedUpFromActivities,omitempty"`
	Name                       string            `xml:"Name,omitempty"`
	Notes                      *QueryResult      `xml:"Notes,omitempty"`
	NotesAndAttachments        *QueryResult      `xml:"NotesAndAttachments,omitempty"`
	OpenActivities             *QueryResult      `xml:"OpenActivities,omitempty"`
	PricebookEntries           *QueryResult      `xml:"PricebookEntries,omitempty"`
	ProcessInstances           *QueryResult      `xml:"ProcessInstances,omitempty"`
	ProcessSteps               *QueryResult      `xml:"ProcessSteps,omitempty"`
	ProductCode                string            `xml:"ProductCode,omitempty"`
	QuantityUnitOfMeasure      string            `xml:"QuantityUnitOfMeasure,omitempty"`
	RecordActionHistories      *QueryResult      `xml:"RecordActionHistories,omitempty"`
	RecordActions              *QueryResult      `xml:"RecordActions,omitempty"`
	StockKeepingUnit           string            `xml:"StockKeepingUnit,omitempty"`
	SystemModstamp             time.Time         `xml:"SystemModstamp,omitempty"`
	Tasks                      *QueryResult      `xml:"Tasks,omitempty"`
	UserRecordAccess           *UserRecordAccess `xml:"UserRecordAccess,omitempty"`
}

type Product2Feed struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com Product2Feed"`

	*SObject

	BestComment        *FeedComment `xml:"BestComment,omitempty"`
	BestCommentId      *ID          `xml:"BestCommentId,omitempty"`
	Body               string       `xml:"Body,omitempty"`
	CommentCount       int32        `xml:"CommentCount,omitempty"`
	CreatedBy          *SObject     `xml:"CreatedBy,omitempty"`
	CreatedById        *ID          `xml:"CreatedById,omitempty"`
	CreatedDate        time.Time    `xml:"CreatedDate,omitempty"`
	FeedAttachments    *QueryResult `xml:"FeedAttachments,omitempty"`
	FeedComments       *QueryResult `xml:"FeedComments,omitempty"`
	FeedLikes          *QueryResult `xml:"FeedLikes,omitempty"`
	FeedSignals        *QueryResult `xml:"FeedSignals,omitempty"`
	FeedTrackedChanges *QueryResult `xml:"FeedTrackedChanges,omitempty"`
	InsertedBy         *SObject     `xml:"InsertedBy,omitempty"`
	InsertedById       *ID          `xml:"InsertedById,omitempty"`
	IsDeleted          bool         `xml:"IsDeleted,omitempty"`
	IsRichText         bool         `xml:"IsRichText,omitempty"`
	LastModifiedDate   time.Time    `xml:"LastModifiedDate,omitempty"`
	LikeCount          int32        `xml:"LikeCount,omitempty"`
	LinkUrl            string       `xml:"LinkUrl,omitempty"`
	Parent             *Product2    `xml:"Parent,omitempty"`
	ParentId           *ID          `xml:"ParentId,omitempty"`
	RelatedRecordId    *ID          `xml:"RelatedRecordId,omitempty"`
	SystemModstamp     time.Time    `xml:"SystemModstamp,omitempty"`
	Title              string       `xml:"Title,omitempty"`
	Type               string       `xml:"Type,omitempty"`
}

type Product2History struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com Product2History"`

	*SObject

	CreatedBy   *SObject    `xml:"CreatedBy,omitempty"`
	CreatedById *ID         `xml:"CreatedById,omitempty"`
	CreatedDate time.Time   `xml:"CreatedDate,omitempty"`
	Field       string      `xml:"Field,omitempty"`
	IsDeleted   bool        `xml:"IsDeleted,omitempty"`
	NewValue    interface{} `xml:"NewValue,omitempty"`
	OldValue    interface{} `xml:"OldValue,omitempty"`
	Product2    *Product2   `xml:"Product2,omitempty"`
	Product2Id  *ID         `xml:"Product2Id,omitempty"`
}

type Profile struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com Profile"`

	*SObject

	CreatedBy                                 *User        `xml:"CreatedBy,omitempty"`
	CreatedById                               *ID          `xml:"CreatedById,omitempty"`
	CreatedDate                               time.Time    `xml:"CreatedDate,omitempty"`
	Description                               string       `xml:"Description,omitempty"`
	LastModifiedBy                            *User        `xml:"LastModifiedBy,omitempty"`
	LastModifiedById                          *ID          `xml:"LastModifiedById,omitempty"`
	LastModifiedDate                          time.Time    `xml:"LastModifiedDate,omitempty"`
	LastReferencedDate                        time.Time    `xml:"LastReferencedDate,omitempty"`
	LastViewedDate                            time.Time    `xml:"LastViewedDate,omitempty"`
	Name                                      string       `xml:"Name,omitempty"`
	PermissionsAccessCMC                      bool         `xml:"PermissionsAccessCMC,omitempty"`
	PermissionsActivateContract               bool         `xml:"PermissionsActivateContract,omitempty"`
	PermissionsActivateOrder                  bool         `xml:"PermissionsActivateOrder,omitempty"`
	PermissionsAddAnalyticsRemoteConnections  bool         `xml:"PermissionsAddAnalyticsRemoteConnections,omitempty"`
	PermissionsAddDirectMessageMembers        bool         `xml:"PermissionsAddDirectMessageMembers,omitempty"`
	PermissionsAllowEmailIC                   bool         `xml:"PermissionsAllowEmailIC,omitempty"`
	PermissionsAllowLightningLogin            bool         `xml:"PermissionsAllowLightningLogin,omitempty"`
	PermissionsAllowUniversalSearch           bool         `xml:"PermissionsAllowUniversalSearch,omitempty"`
	PermissionsAllowViewEditConvertedLeads    bool         `xml:"PermissionsAllowViewEditConvertedLeads,omitempty"`
	PermissionsAllowViewKnowledge             bool         `xml:"PermissionsAllowViewKnowledge,omitempty"`
	PermissionsApexRestServices               bool         `xml:"PermissionsApexRestServices,omitempty"`
	PermissionsApiEnabled                     bool         `xml:"PermissionsApiEnabled,omitempty"`
	PermissionsAssignPermissionSets           bool         `xml:"PermissionsAssignPermissionSets,omitempty"`
	PermissionsAssignTopics                   bool         `xml:"PermissionsAssignTopics,omitempty"`
	PermissionsAuthorApex                     bool         `xml:"PermissionsAuthorApex,omitempty"`
	PermissionsAutomaticActivityCapture       bool         `xml:"PermissionsAutomaticActivityCapture,omitempty"`
	PermissionsBulkApiHardDelete              bool         `xml:"PermissionsBulkApiHardDelete,omitempty"`
	PermissionsBulkMacrosAllowed              bool         `xml:"PermissionsBulkMacrosAllowed,omitempty"`
	PermissionsCampaignInfluence2             bool         `xml:"PermissionsCampaignInfluence2,omitempty"`
	PermissionsCanApproveFeedPost             bool         `xml:"PermissionsCanApproveFeedPost,omitempty"`
	PermissionsCanEditDataPrepRecipe          bool         `xml:"PermissionsCanEditDataPrepRecipe,omitempty"`
	PermissionsCanInsertFeedSystemFields      bool         `xml:"PermissionsCanInsertFeedSystemFields,omitempty"`
	PermissionsCanManageMaps                  bool         `xml:"PermissionsCanManageMaps,omitempty"`
	PermissionsCanUseNewDashboardBuilder      bool         `xml:"PermissionsCanUseNewDashboardBuilder,omitempty"`
	PermissionsCanVerifyComment               bool         `xml:"PermissionsCanVerifyComment,omitempty"`
	PermissionsChangeDashboardColors          bool         `xml:"PermissionsChangeDashboardColors,omitempty"`
	PermissionsChatterComposeUiCodesnippet    bool         `xml:"PermissionsChatterComposeUiCodesnippet,omitempty"`
	PermissionsChatterEditOwnPost             bool         `xml:"PermissionsChatterEditOwnPost,omitempty"`
	PermissionsChatterEditOwnRecordPost       bool         `xml:"PermissionsChatterEditOwnRecordPost,omitempty"`
	PermissionsChatterFileLink                bool         `xml:"PermissionsChatterFileLink,omitempty"`
	PermissionsChatterInternalUser            bool         `xml:"PermissionsChatterInternalUser,omitempty"`
	PermissionsChatterInviteExternalUsers     bool         `xml:"PermissionsChatterInviteExternalUsers,omitempty"`
	PermissionsChatterOwnGroups               bool         `xml:"PermissionsChatterOwnGroups,omitempty"`
	PermissionsCloseConversations             bool         `xml:"PermissionsCloseConversations,omitempty"`
	PermissionsConfigCustomRecs               bool         `xml:"PermissionsConfigCustomRecs,omitempty"`
	PermissionsConnectOrgToEnvironmentHub     bool         `xml:"PermissionsConnectOrgToEnvironmentHub,omitempty"`
	PermissionsContentAdministrator           bool         `xml:"PermissionsContentAdministrator,omitempty"`
	PermissionsContentWorkspaces              bool         `xml:"PermissionsContentWorkspaces,omitempty"`
	PermissionsConvertLeads                   bool         `xml:"PermissionsConvertLeads,omitempty"`
	PermissionsCreateCustomizeDashboards      bool         `xml:"PermissionsCreateCustomizeDashboards,omitempty"`
	PermissionsCreateCustomizeFilters         bool         `xml:"PermissionsCreateCustomizeFilters,omitempty"`
	PermissionsCreateCustomizeReports         bool         `xml:"PermissionsCreateCustomizeReports,omitempty"`
	PermissionsCreateDashboardFolders         bool         `xml:"PermissionsCreateDashboardFolders,omitempty"`
	PermissionsCreateLtngTempInPub            bool         `xml:"PermissionsCreateLtngTempInPub,omitempty"`
	PermissionsCreateMultiforce               bool         `xml:"PermissionsCreateMultiforce,omitempty"`
	PermissionsCreateReportFolders            bool         `xml:"PermissionsCreateReportFolders,omitempty"`
	PermissionsCreateReportInLightning        bool         `xml:"PermissionsCreateReportInLightning,omitempty"`
	PermissionsCreateTopics                   bool         `xml:"PermissionsCreateTopics,omitempty"`
	PermissionsCreateWorkspaces               bool         `xml:"PermissionsCreateWorkspaces,omitempty"`
	PermissionsCustomMobileAppsAccess         bool         `xml:"PermissionsCustomMobileAppsAccess,omitempty"`
	PermissionsCustomSidebarOnAllPages        bool         `xml:"PermissionsCustomSidebarOnAllPages,omitempty"`
	PermissionsCustomizeApplication           bool         `xml:"PermissionsCustomizeApplication,omitempty"`
	PermissionsDataExport                     bool         `xml:"PermissionsDataExport,omitempty"`
	PermissionsDelegatedTwoFactor             bool         `xml:"PermissionsDelegatedTwoFactor,omitempty"`
	PermissionsDeleteActivatedContract        bool         `xml:"PermissionsDeleteActivatedContract,omitempty"`
	PermissionsDeleteTopics                   bool         `xml:"PermissionsDeleteTopics,omitempty"`
	PermissionsDistributeFromPersWksp         bool         `xml:"PermissionsDistributeFromPersWksp,omitempty"`
	PermissionsEditActivatedOrders            bool         `xml:"PermissionsEditActivatedOrders,omitempty"`
	PermissionsEditBrandTemplates             bool         `xml:"PermissionsEditBrandTemplates,omitempty"`
	PermissionsEditCaseComments               bool         `xml:"PermissionsEditCaseComments,omitempty"`
	PermissionsEditEvent                      bool         `xml:"PermissionsEditEvent,omitempty"`
	PermissionsEditHtmlTemplates              bool         `xml:"PermissionsEditHtmlTemplates,omitempty"`
	PermissionsEditKnowledge                  bool         `xml:"PermissionsEditKnowledge,omitempty"`
	PermissionsEditMyDashboards               bool         `xml:"PermissionsEditMyDashboards,omitempty"`
	PermissionsEditMyReports                  bool         `xml:"PermissionsEditMyReports,omitempty"`
	PermissionsEditOppLineItemUnitPrice       bool         `xml:"PermissionsEditOppLineItemUnitPrice,omitempty"`
	PermissionsEditPublicDocuments            bool         `xml:"PermissionsEditPublicDocuments,omitempty"`
	PermissionsEditPublicFilters              bool         `xml:"PermissionsEditPublicFilters,omitempty"`
	PermissionsEditPublicTemplates            bool         `xml:"PermissionsEditPublicTemplates,omitempty"`
	PermissionsEditReadonlyFields             bool         `xml:"PermissionsEditReadonlyFields,omitempty"`
	PermissionsEditTask                       bool         `xml:"PermissionsEditTask,omitempty"`
	PermissionsEditTopics                     bool         `xml:"PermissionsEditTopics,omitempty"`
	PermissionsEmailAdministration            bool         `xml:"PermissionsEmailAdministration,omitempty"`
	PermissionsEmailMass                      bool         `xml:"PermissionsEmailMass,omitempty"`
	PermissionsEmailSingle                    bool         `xml:"PermissionsEmailSingle,omitempty"`
	PermissionsEmailTemplateManagement        bool         `xml:"PermissionsEmailTemplateManagement,omitempty"`
	PermissionsEnableCommunityAppLauncher     bool         `xml:"PermissionsEnableCommunityAppLauncher,omitempty"`
	PermissionsEnableNotifications            bool         `xml:"PermissionsEnableNotifications,omitempty"`
	PermissionsExportReport                   bool         `xml:"PermissionsExportReport,omitempty"`
	PermissionsFeedPinning                    bool         `xml:"PermissionsFeedPinning,omitempty"`
	PermissionsFlowUFLRequired                bool         `xml:"PermissionsFlowUFLRequired,omitempty"`
	PermissionsForceTwoFactor                 bool         `xml:"PermissionsForceTwoFactor,omitempty"`
	PermissionsGiveRecognitionBadge           bool         `xml:"PermissionsGiveRecognitionBadge,omitempty"`
	PermissionsGovernNetworks                 bool         `xml:"PermissionsGovernNetworks,omitempty"`
	PermissionsHideReadByList                 bool         `xml:"PermissionsHideReadByList,omitempty"`
	PermissionsIdentityConnect                bool         `xml:"PermissionsIdentityConnect,omitempty"`
	PermissionsIdentityEnabled                bool         `xml:"PermissionsIdentityEnabled,omitempty"`
	PermissionsImportCustomObjects            bool         `xml:"PermissionsImportCustomObjects,omitempty"`
	PermissionsImportLeads                    bool         `xml:"PermissionsImportLeads,omitempty"`
	PermissionsImportPersonal                 bool         `xml:"PermissionsImportPersonal,omitempty"`
	PermissionsInsightsAppAdmin               bool         `xml:"PermissionsInsightsAppAdmin,omitempty"`
	PermissionsInsightsAppDashboardEditor     bool         `xml:"PermissionsInsightsAppDashboardEditor,omitempty"`
	PermissionsInsightsAppEltEditor           bool         `xml:"PermissionsInsightsAppEltEditor,omitempty"`
	PermissionsInsightsAppUploadUser          bool         `xml:"PermissionsInsightsAppUploadUser,omitempty"`
	PermissionsInsightsAppUser                bool         `xml:"PermissionsInsightsAppUser,omitempty"`
	PermissionsInsightsCreateApplication      bool         `xml:"PermissionsInsightsCreateApplication,omitempty"`
	PermissionsInstallMultiforce              bool         `xml:"PermissionsInstallMultiforce,omitempty"`
	PermissionsIotUser                        bool         `xml:"PermissionsIotUser,omitempty"`
	PermissionsLightningConsoleAllowedForUser bool         `xml:"PermissionsLightningConsoleAllowedForUser,omitempty"`
	PermissionsLightningExperienceUser        bool         `xml:"PermissionsLightningExperienceUser,omitempty"`
	PermissionsListEmailSend                  bool         `xml:"PermissionsListEmailSend,omitempty"`
	PermissionsLtngPromoReserved01UserPerm    bool         `xml:"PermissionsLtngPromoReserved01UserPerm,omitempty"`
	PermissionsManageAnalyticSnapshots        bool         `xml:"PermissionsManageAnalyticSnapshots,omitempty"`
	PermissionsManageAuthProviders            bool         `xml:"PermissionsManageAuthProviders,omitempty"`
	PermissionsManageBusinessHourHolidays     bool         `xml:"PermissionsManageBusinessHourHolidays,omitempty"`
	PermissionsManageCallCenters              bool         `xml:"PermissionsManageCallCenters,omitempty"`
	PermissionsManageCases                    bool         `xml:"PermissionsManageCases,omitempty"`
	PermissionsManageCategories               bool         `xml:"PermissionsManageCategories,omitempty"`
	PermissionsManageCertificates             bool         `xml:"PermissionsManageCertificates,omitempty"`
	PermissionsManageChatterMessages          bool         `xml:"PermissionsManageChatterMessages,omitempty"`
	PermissionsManageContentPermissions       bool         `xml:"PermissionsManageContentPermissions,omitempty"`
	PermissionsManageContentProperties        bool         `xml:"PermissionsManageContentProperties,omitempty"`
	PermissionsManageContentTypes             bool         `xml:"PermissionsManageContentTypes,omitempty"`
	PermissionsManageCustomPermissions        bool         `xml:"PermissionsManageCustomPermissions,omitempty"`
	PermissionsManageCustomReportTypes        bool         `xml:"PermissionsManageCustomReportTypes,omitempty"`
	PermissionsManageDashbdsInPubFolders      bool         `xml:"PermissionsManageDashbdsInPubFolders,omitempty"`
	PermissionsManageDataCategories           bool         `xml:"PermissionsManageDataCategories,omitempty"`
	PermissionsManageDataIntegrations         bool         `xml:"PermissionsManageDataIntegrations,omitempty"`
	PermissionsManageDynamicDashboards        bool         `xml:"PermissionsManageDynamicDashboards,omitempty"`
	PermissionsManageEmailClientConfig        bool         `xml:"PermissionsManageEmailClientConfig,omitempty"`
	PermissionsManageEncryptionKeys           bool         `xml:"PermissionsManageEncryptionKeys,omitempty"`
	PermissionsManageExchangeConfig           bool         `xml:"PermissionsManageExchangeConfig,omitempty"`
	PermissionsManageHealthCheck              bool         `xml:"PermissionsManageHealthCheck,omitempty"`
	PermissionsManageInteraction              bool         `xml:"PermissionsManageInteraction,omitempty"`
	PermissionsManageInternalUsers            bool         `xml:"PermissionsManageInternalUsers,omitempty"`
	PermissionsManageIpAddresses              bool         `xml:"PermissionsManageIpAddresses,omitempty"`
	PermissionsManageKnowledge                bool         `xml:"PermissionsManageKnowledge,omitempty"`
	PermissionsManageKnowledgeImportExport    bool         `xml:"PermissionsManageKnowledgeImportExport,omitempty"`
	PermissionsManageLeads                    bool         `xml:"PermissionsManageLeads,omitempty"`
	PermissionsManageLoginAccessPolicies      bool         `xml:"PermissionsManageLoginAccessPolicies,omitempty"`
	PermissionsManageMobile                   bool         `xml:"PermissionsManageMobile,omitempty"`
	PermissionsManageNetworks                 bool         `xml:"PermissionsManageNetworks,omitempty"`
	PermissionsManagePasswordPolicies         bool         `xml:"PermissionsManagePasswordPolicies,omitempty"`
	PermissionsManageProfilesPermissionsets   bool         `xml:"PermissionsManageProfilesPermissionsets,omitempty"`
	PermissionsManagePvtRptsAndDashbds        bool         `xml:"PermissionsManagePvtRptsAndDashbds,omitempty"`
	PermissionsManageRemoteAccess             bool         `xml:"PermissionsManageRemoteAccess,omitempty"`
	PermissionsManageReportsInPubFolders      bool         `xml:"PermissionsManageReportsInPubFolders,omitempty"`
	PermissionsManageRoles                    bool         `xml:"PermissionsManageRoles,omitempty"`
	PermissionsManageSearchPromotionRules     bool         `xml:"PermissionsManageSearchPromotionRules,omitempty"`
	PermissionsManageSessionPermissionSets    bool         `xml:"PermissionsManageSessionPermissionSets,omitempty"`
	PermissionsManageSharing                  bool         `xml:"PermissionsManageSharing,omitempty"`
	PermissionsManageSolutions                bool         `xml:"PermissionsManageSolutions,omitempty"`
	PermissionsManageSurveys                  bool         `xml:"PermissionsManageSurveys,omitempty"`
	PermissionsManageSynonyms                 bool         `xml:"PermissionsManageSynonyms,omitempty"`
	PermissionsManageTemplatedApp             bool         `xml:"PermissionsManageTemplatedApp,omitempty"`
	PermissionsManageTwoFactor                bool         `xml:"PermissionsManageTwoFactor,omitempty"`
	PermissionsManageUnlistedGroups           bool         `xml:"PermissionsManageUnlistedGroups,omitempty"`
	PermissionsManageUsers                    bool         `xml:"PermissionsManageUsers,omitempty"`
	PermissionsMassInlineEdit                 bool         `xml:"PermissionsMassInlineEdit,omitempty"`
	PermissionsMergeTopics                    bool         `xml:"PermissionsMergeTopics,omitempty"`
	PermissionsModerateChatter                bool         `xml:"PermissionsModerateChatter,omitempty"`
	PermissionsModerateNetworkUsers           bool         `xml:"PermissionsModerateNetworkUsers,omitempty"`
	PermissionsModifyAllData                  bool         `xml:"PermissionsModifyAllData,omitempty"`
	PermissionsModifyMetadata                 bool         `xml:"PermissionsModifyMetadata,omitempty"`
	PermissionsModifySecureAgents             bool         `xml:"PermissionsModifySecureAgents,omitempty"`
	PermissionsNewReportBuilder               bool         `xml:"PermissionsNewReportBuilder,omitempty"`
	PermissionsPackaging2                     bool         `xml:"PermissionsPackaging2,omitempty"`
	PermissionsPasswordNeverExpires           bool         `xml:"PermissionsPasswordNeverExpires,omitempty"`
	PermissionsPreventClassicExperience       bool         `xml:"PermissionsPreventClassicExperience,omitempty"`
	PermissionsPrivacyDataAccess              bool         `xml:"PermissionsPrivacyDataAccess,omitempty"`
	PermissionsPublishMultiforce              bool         `xml:"PermissionsPublishMultiforce,omitempty"`
	PermissionsRemoveDirectMessageMembers     bool         `xml:"PermissionsRemoveDirectMessageMembers,omitempty"`
	PermissionsResetPasswords                 bool         `xml:"PermissionsResetPasswords,omitempty"`
	PermissionsRunFlow                        bool         `xml:"PermissionsRunFlow,omitempty"`
	PermissionsRunReports                     bool         `xml:"PermissionsRunReports,omitempty"`
	PermissionsSalesConsole                   bool         `xml:"PermissionsSalesConsole,omitempty"`
	PermissionsScheduleReports                bool         `xml:"PermissionsScheduleReports,omitempty"`
	PermissionsSelectFilesFromSalesforce      bool         `xml:"PermissionsSelectFilesFromSalesforce,omitempty"`
	PermissionsSendAnnouncementEmails         bool         `xml:"PermissionsSendAnnouncementEmails,omitempty"`
	PermissionsSendSitRequests                bool         `xml:"PermissionsSendSitRequests,omitempty"`
	PermissionsShareInternalArticles          bool         `xml:"PermissionsShareInternalArticles,omitempty"`
	PermissionsShowCompanyNameAsUserBadge     bool         `xml:"PermissionsShowCompanyNameAsUserBadge,omitempty"`
	PermissionsSolutionImport                 bool         `xml:"PermissionsSolutionImport,omitempty"`
	PermissionsStdAutomaticActivityCapture    bool         `xml:"PermissionsStdAutomaticActivityCapture,omitempty"`
	PermissionsSubmitMacrosAllowed            bool         `xml:"PermissionsSubmitMacrosAllowed,omitempty"`
	PermissionsSubscribeDashboardToOtherUsers bool         `xml:"PermissionsSubscribeDashboardToOtherUsers,omitempty"`
	PermissionsSubscribeReportToOtherUsers    bool         `xml:"PermissionsSubscribeReportToOtherUsers,omitempty"`
	PermissionsSubscribeReportsRunAsUser      bool         `xml:"PermissionsSubscribeReportsRunAsUser,omitempty"`
	PermissionsSubscribeToLightningDashboards bool         `xml:"PermissionsSubscribeToLightningDashboards,omitempty"`
	PermissionsSubscribeToLightningReports    bool         `xml:"PermissionsSubscribeToLightningReports,omitempty"`
	PermissionsTransferAnyCase                bool         `xml:"PermissionsTransferAnyCase,omitempty"`
	PermissionsTransferAnyEntity              bool         `xml:"PermissionsTransferAnyEntity,omitempty"`
	PermissionsTransferAnyLead                bool         `xml:"PermissionsTransferAnyLead,omitempty"`
	PermissionsTwoFactorApi                   bool         `xml:"PermissionsTwoFactorApi,omitempty"`
	PermissionsUseTeamReassignWizards         bool         `xml:"PermissionsUseTeamReassignWizards,omitempty"`
	PermissionsUseTemplatedApp                bool         `xml:"PermissionsUseTemplatedApp,omitempty"`
	PermissionsUseWebLink                     bool         `xml:"PermissionsUseWebLink,omitempty"`
	PermissionsViewAllActivities              bool         `xml:"PermissionsViewAllActivities,omitempty"`
	PermissionsViewAllData                    bool         `xml:"PermissionsViewAllData,omitempty"`
	PermissionsViewAllUsers                   bool         `xml:"PermissionsViewAllUsers,omitempty"`
	PermissionsViewContent                    bool         `xml:"PermissionsViewContent,omitempty"`
	PermissionsViewDataAssessment             bool         `xml:"PermissionsViewDataAssessment,omitempty"`
	PermissionsViewDataCategories             bool         `xml:"PermissionsViewDataCategories,omitempty"`
	PermissionsViewEncryptedData              bool         `xml:"PermissionsViewEncryptedData,omitempty"`
	PermissionsViewEventLogFiles              bool         `xml:"PermissionsViewEventLogFiles,omitempty"`
	PermissionsViewHealthCheck                bool         `xml:"PermissionsViewHealthCheck,omitempty"`
	PermissionsViewHelpLink                   bool         `xml:"PermissionsViewHelpLink,omitempty"`
	PermissionsViewMyTeamsDashboards          bool         `xml:"PermissionsViewMyTeamsDashboards,omitempty"`
	PermissionsViewOnlyEmbeddedAppUser        bool         `xml:"PermissionsViewOnlyEmbeddedAppUser,omitempty"`
	PermissionsViewPublicDashboards           bool         `xml:"PermissionsViewPublicDashboards,omitempty"`
	PermissionsViewPublicReports              bool         `xml:"PermissionsViewPublicReports,omitempty"`
	PermissionsViewRoles                      bool         `xml:"PermissionsViewRoles,omitempty"`
	PermissionsViewSetup                      bool         `xml:"PermissionsViewSetup,omitempty"`
	PermissionsWaveTabularDownload            bool         `xml:"PermissionsWaveTabularDownload,omitempty"`
	PermissionsWorkCalibrationUser            bool         `xml:"PermissionsWorkCalibrationUser,omitempty"`
	PermissionsWorkDotComUserPerm             bool         `xml:"PermissionsWorkDotComUserPerm,omitempty"`
	SystemModstamp                            time.Time    `xml:"SystemModstamp,omitempty"`
	UserLicense                               *UserLicense `xml:"UserLicense,omitempty"`
	UserLicenseId                             *ID          `xml:"UserLicenseId,omitempty"`
	UserType                                  string       `xml:"UserType,omitempty"`
	Users                                     *QueryResult `xml:"Users,omitempty"`
}

type Publisher struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com Publisher"`

	*SObject

	DurableId                  string       `xml:"DurableId,omitempty"`
	InstalledEntityDefinitions *QueryResult `xml:"InstalledEntityDefinitions,omitempty"`
	InstalledFieldDefinitions  *QueryResult `xml:"InstalledFieldDefinitions,omitempty"`
	IsSalesforce               bool         `xml:"IsSalesforce,omitempty"`
	MajorVersion               int32        `xml:"MajorVersion,omitempty"`
	MinorVersion               int32        `xml:"MinorVersion,omitempty"`
	Name                       string       `xml:"Name,omitempty"`
	NamespacePrefix            string       `xml:"NamespacePrefix,omitempty"`
}

type PushTopic struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com PushTopic"`

	*SObject

	ApiVersion                 float64   `xml:"ApiVersion,omitempty"`
	CreatedBy                  *User     `xml:"CreatedBy,omitempty"`
	CreatedById                *ID       `xml:"CreatedById,omitempty"`
	CreatedDate                time.Time `xml:"CreatedDate,omitempty"`
	Description                string    `xml:"Description,omitempty"`
	IsActive                   bool      `xml:"IsActive,omitempty"`
	IsDeleted                  bool      `xml:"IsDeleted,omitempty"`
	LastModifiedBy             *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById           *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate           time.Time `xml:"LastModifiedDate,omitempty"`
	Name                       string    `xml:"Name,omitempty"`
	NotifyForFields            string    `xml:"NotifyForFields,omitempty"`
	NotifyForOperationCreate   bool      `xml:"NotifyForOperationCreate,omitempty"`
	NotifyForOperationDelete   bool      `xml:"NotifyForOperationDelete,omitempty"`
	NotifyForOperationUndelete bool      `xml:"NotifyForOperationUndelete,omitempty"`
	NotifyForOperationUpdate   bool      `xml:"NotifyForOperationUpdate,omitempty"`
	NotifyForOperations        string    `xml:"NotifyForOperations,omitempty"`
	Query                      string    `xml:"Query,omitempty"`
	SystemModstamp             time.Time `xml:"SystemModstamp,omitempty"`
}

type QueueSobject struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com QueueSobject"`

	*SObject

	CreatedBy      *User     `xml:"CreatedBy,omitempty"`
	CreatedById    *ID       `xml:"CreatedById,omitempty"`
	Queue          *Group    `xml:"Queue,omitempty"`
	QueueId        *ID       `xml:"QueueId,omitempty"`
	SobjectType    string    `xml:"SobjectType,omitempty"`
	SystemModstamp time.Time `xml:"SystemModstamp,omitempty"`
}

type QuickText struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com QuickText"`

	*SObject

	Category           string            `xml:"Category,omitempty"`
	Channel            string            `xml:"Channel,omitempty"`
	CreatedBy          *User             `xml:"CreatedBy,omitempty"`
	CreatedById        *ID               `xml:"CreatedById,omitempty"`
	CreatedDate        time.Time         `xml:"CreatedDate,omitempty"`
	Histories          *QueryResult      `xml:"Histories,omitempty"`
	IsDeleted          bool              `xml:"IsDeleted,omitempty"`
	LastModifiedBy     *User             `xml:"LastModifiedBy,omitempty"`
	LastModifiedById   *ID               `xml:"LastModifiedById,omitempty"`
	LastModifiedDate   time.Time         `xml:"LastModifiedDate,omitempty"`
	LastReferencedDate time.Time         `xml:"LastReferencedDate,omitempty"`
	LastViewedDate     time.Time         `xml:"LastViewedDate,omitempty"`
	Message            string            `xml:"Message,omitempty"`
	Name               string            `xml:"Name,omitempty"`
	Owner              *SObject          `xml:"Owner,omitempty"`
	OwnerId            *ID               `xml:"OwnerId,omitempty"`
	SystemModstamp     time.Time         `xml:"SystemModstamp,omitempty"`
	UserRecordAccess   *UserRecordAccess `xml:"UserRecordAccess,omitempty"`
}

type QuickTextHistory struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com QuickTextHistory"`

	*SObject

	CreatedBy   *SObject    `xml:"CreatedBy,omitempty"`
	CreatedById *ID         `xml:"CreatedById,omitempty"`
	CreatedDate time.Time   `xml:"CreatedDate,omitempty"`
	Field       string      `xml:"Field,omitempty"`
	IsDeleted   bool        `xml:"IsDeleted,omitempty"`
	NewValue    interface{} `xml:"NewValue,omitempty"`
	OldValue    interface{} `xml:"OldValue,omitempty"`
	QuickText   *QuickText  `xml:"QuickText,omitempty"`
	QuickTextId *ID         `xml:"QuickTextId,omitempty"`
}

type QuickTextShare struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com QuickTextShare"`

	*SObject

	AccessLevel      string     `xml:"AccessLevel,omitempty"`
	IsDeleted        bool       `xml:"IsDeleted,omitempty"`
	LastModifiedBy   *User      `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID        `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time  `xml:"LastModifiedDate,omitempty"`
	Parent           *QuickText `xml:"Parent,omitempty"`
	ParentId         *ID        `xml:"ParentId,omitempty"`
	RowCause         string     `xml:"RowCause,omitempty"`
	UserOrGroup      *SObject   `xml:"UserOrGroup,omitempty"`
	UserOrGroupId    *ID        `xml:"UserOrGroupId,omitempty"`
}

type QuoteTemplateRichTextData struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com QuoteTemplateRichTextData"`

	*SObject

	CreatedBy        *User             `xml:"CreatedBy,omitempty"`
	CreatedById      *ID               `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time         `xml:"CreatedDate,omitempty"`
	Data             string            `xml:"Data,omitempty"`
	IsDeleted        bool              `xml:"IsDeleted,omitempty"`
	LastModifiedBy   *User             `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID               `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time         `xml:"LastModifiedDate,omitempty"`
	Name             string            `xml:"Name,omitempty"`
	SystemModstamp   time.Time         `xml:"SystemModstamp,omitempty"`
	UserRecordAccess *UserRecordAccess `xml:"UserRecordAccess,omitempty"`
}

type RecentlyViewed struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com RecentlyViewed"`

	*SObject

	Alias              string      `xml:"Alias,omitempty"`
	Email              string      `xml:"Email,omitempty"`
	FirstName          string      `xml:"FirstName,omitempty"`
	IsActive           bool        `xml:"IsActive,omitempty"`
	Language           string      `xml:"Language,omitempty"`
	LastName           string      `xml:"LastName,omitempty"`
	LastReferencedDate time.Time   `xml:"LastReferencedDate,omitempty"`
	LastViewedDate     time.Time   `xml:"LastViewedDate,omitempty"`
	Name               string      `xml:"Name,omitempty"`
	NameOrAlias        string      `xml:"NameOrAlias,omitempty"`
	Phone              string      `xml:"Phone,omitempty"`
	Profile            *Profile    `xml:"Profile,omitempty"`
	ProfileId          *ID         `xml:"ProfileId,omitempty"`
	RecordType         *RecordType `xml:"RecordType,omitempty"`
	RecordTypeId       *ID         `xml:"RecordTypeId,omitempty"`
	Title              string      `xml:"Title,omitempty"`
	Type               string      `xml:"Type,omitempty"`
	UserRole           *UserRole   `xml:"UserRole,omitempty"`
	UserRoleId         *ID         `xml:"UserRoleId,omitempty"`
}

type RecordAction struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com RecordAction"`

	*SObject

	ActionDefinition string            `xml:"ActionDefinition,omitempty"`
	CreatedBy        *User             `xml:"CreatedBy,omitempty"`
	CreatedById      *ID               `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time         `xml:"CreatedDate,omitempty"`
	FlowDefinition   string            `xml:"FlowDefinition,omitempty"`
	FlowInterview    *FlowInterview    `xml:"FlowInterview,omitempty"`
	FlowInterviewId  *ID               `xml:"FlowInterviewId,omitempty"`
	IsDeleted        bool              `xml:"IsDeleted,omitempty"`
	IsMandatory      bool              `xml:"IsMandatory,omitempty"`
	LastModifiedBy   *User             `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID               `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time         `xml:"LastModifiedDate,omitempty"`
	Order            int32             `xml:"Order,omitempty"`
	Pinned           string            `xml:"Pinned,omitempty"`
	Record           *Account          `xml:"Record,omitempty"`
	RecordId         *ID               `xml:"RecordId,omitempty"`
	Status           string            `xml:"Status,omitempty"`
	SystemModstamp   time.Time         `xml:"SystemModstamp,omitempty"`
	UserRecordAccess *UserRecordAccess `xml:"UserRecordAccess,omitempty"`
}

type RecordActionHistory struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com RecordActionHistory"`

	*SObject

	ActionDefinitionApiName string    `xml:"ActionDefinitionApiName,omitempty"`
	ActionDefinitionLabel   string    `xml:"ActionDefinitionLabel,omitempty"`
	ActionType              string    `xml:"ActionType,omitempty"`
	CreatedBy               *User     `xml:"CreatedBy,omitempty"`
	CreatedById             *ID       `xml:"CreatedById,omitempty"`
	CreatedDate             time.Time `xml:"CreatedDate,omitempty"`
	IsMandatory             bool      `xml:"IsMandatory,omitempty"`
	LastModifiedBy          *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById        *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate        time.Time `xml:"LastModifiedDate,omitempty"`
	LoggedTime              time.Time `xml:"LoggedTime,omitempty"`
	ParentRecord            *Account  `xml:"ParentRecord,omitempty"`
	ParentRecordId          *ID       `xml:"ParentRecordId,omitempty"`
	Pinned                  string    `xml:"Pinned,omitempty"`
	RecordActionId          string    `xml:"RecordActionId,omitempty"`
	State                   string    `xml:"State,omitempty"`
	SystemModstamp          time.Time `xml:"SystemModstamp,omitempty"`
	User                    *SObject  `xml:"User,omitempty"`
	UserId                  *ID       `xml:"UserId,omitempty"`
}

type RecordType struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com RecordType"`

	*SObject

	BusinessProcessId *ID       `xml:"BusinessProcessId,omitempty"`
	CreatedBy         *User     `xml:"CreatedBy,omitempty"`
	CreatedById       *ID       `xml:"CreatedById,omitempty"`
	CreatedDate       time.Time `xml:"CreatedDate,omitempty"`
	Description       string    `xml:"Description,omitempty"`
	DeveloperName     string    `xml:"DeveloperName,omitempty"`
	IsActive          bool      `xml:"IsActive,omitempty"`
	LastModifiedBy    *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById  *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate  time.Time `xml:"LastModifiedDate,omitempty"`
	Name              string    `xml:"Name,omitempty"`
	NamespacePrefix   string    `xml:"NamespacePrefix,omitempty"`
	SobjectType       string    `xml:"SobjectType,omitempty"`
	SystemModstamp    time.Time `xml:"SystemModstamp,omitempty"`
}

type RelationshipDomain struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com RelationshipDomain"`

	*SObject

	ChildSobject          *EntityDefinition    `xml:"ChildSobject,omitempty"`
	ChildSobjectId        string               `xml:"ChildSobjectId,omitempty"`
	DurableId             string               `xml:"DurableId,omitempty"`
	Field                 *FieldDefinition     `xml:"Field,omitempty"`
	FieldId               string               `xml:"FieldId,omitempty"`
	IsCascadeDelete       bool                 `xml:"IsCascadeDelete,omitempty"`
	IsDeprecatedAndHidden bool                 `xml:"IsDeprecatedAndHidden,omitempty"`
	IsRestrictedDelete    bool                 `xml:"IsRestrictedDelete,omitempty"`
	JunctionIdListNames   *JunctionIdListNames `xml:"JunctionIdListNames,omitempty"`
	ParentSobject         *EntityDefinition    `xml:"ParentSobject,omitempty"`
	ParentSobjectId       string               `xml:"ParentSobjectId,omitempty"`
	RelationshipInfo      *RelationshipInfo    `xml:"RelationshipInfo,omitempty"`
	RelationshipInfoId    string               `xml:"RelationshipInfoId,omitempty"`
	RelationshipName      string               `xml:"RelationshipName,omitempty"`
}

type RelationshipInfo struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com RelationshipInfo"`

	*SObject

	ChildSobject          *EntityDefinition    `xml:"ChildSobject,omitempty"`
	ChildSobjectId        string               `xml:"ChildSobjectId,omitempty"`
	DurableId             string               `xml:"DurableId,omitempty"`
	Field                 *FieldDefinition     `xml:"Field,omitempty"`
	FieldId               string               `xml:"FieldId,omitempty"`
	IsCascadeDelete       bool                 `xml:"IsCascadeDelete,omitempty"`
	IsDeprecatedAndHidden bool                 `xml:"IsDeprecatedAndHidden,omitempty"`
	IsRestrictedDelete    bool                 `xml:"IsRestrictedDelete,omitempty"`
	JunctionIdListNames   *JunctionIdListNames `xml:"JunctionIdListNames,omitempty"`
	RelationshipDomains   *QueryResult         `xml:"RelationshipDomains,omitempty"`
}

type Report struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com Report"`

	*SObject

	AttachedContentDocuments   *QueryResult `xml:"AttachedContentDocuments,omitempty"`
	CombinedAttachments        *QueryResult `xml:"CombinedAttachments,omitempty"`
	ContentDocumentLinks       *QueryResult `xml:"ContentDocumentLinks,omitempty"`
	CreatedBy                  *User        `xml:"CreatedBy,omitempty"`
	CreatedById                *ID          `xml:"CreatedById,omitempty"`
	CreatedDate                time.Time    `xml:"CreatedDate,omitempty"`
	Description                string       `xml:"Description,omitempty"`
	DeveloperName              string       `xml:"DeveloperName,omitempty"`
	FeedSubscriptionsForEntity *QueryResult `xml:"FeedSubscriptionsForEntity,omitempty"`
	Feeds                      *QueryResult `xml:"Feeds,omitempty"`
	FolderName                 string       `xml:"FolderName,omitempty"`
	Format                     string       `xml:"Format,omitempty"`
	IsDeleted                  bool         `xml:"IsDeleted,omitempty"`
	LastModifiedBy             *User        `xml:"LastModifiedBy,omitempty"`
	LastModifiedById           *ID          `xml:"LastModifiedById,omitempty"`
	LastModifiedDate           time.Time    `xml:"LastModifiedDate,omitempty"`
	LastReferencedDate         time.Time    `xml:"LastReferencedDate,omitempty"`
	LastRunDate                time.Time    `xml:"LastRunDate,omitempty"`
	LastViewedDate             time.Time    `xml:"LastViewedDate,omitempty"`
	Name                       string       `xml:"Name,omitempty"`
	NamespacePrefix            string       `xml:"NamespacePrefix,omitempty"`
	Owner                      *SObject     `xml:"Owner,omitempty"`
	OwnerId                    *ID          `xml:"OwnerId,omitempty"`
	SystemModstamp             time.Time    `xml:"SystemModstamp,omitempty"`
}

type ReportFeed struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ReportFeed"`

	*SObject

	BestComment        *FeedComment `xml:"BestComment,omitempty"`
	BestCommentId      *ID          `xml:"BestCommentId,omitempty"`
	Body               string       `xml:"Body,omitempty"`
	CommentCount       int32        `xml:"CommentCount,omitempty"`
	CreatedBy          *SObject     `xml:"CreatedBy,omitempty"`
	CreatedById        *ID          `xml:"CreatedById,omitempty"`
	CreatedDate        time.Time    `xml:"CreatedDate,omitempty"`
	FeedAttachments    *QueryResult `xml:"FeedAttachments,omitempty"`
	FeedComments       *QueryResult `xml:"FeedComments,omitempty"`
	FeedLikes          *QueryResult `xml:"FeedLikes,omitempty"`
	FeedSignals        *QueryResult `xml:"FeedSignals,omitempty"`
	FeedTrackedChanges *QueryResult `xml:"FeedTrackedChanges,omitempty"`
	InsertedBy         *SObject     `xml:"InsertedBy,omitempty"`
	InsertedById       *ID          `xml:"InsertedById,omitempty"`
	IsDeleted          bool         `xml:"IsDeleted,omitempty"`
	IsRichText         bool         `xml:"IsRichText,omitempty"`
	LastModifiedDate   time.Time    `xml:"LastModifiedDate,omitempty"`
	LikeCount          int32        `xml:"LikeCount,omitempty"`
	LinkUrl            string       `xml:"LinkUrl,omitempty"`
	Parent             *Report      `xml:"Parent,omitempty"`
	ParentId           *ID          `xml:"ParentId,omitempty"`
	RelatedRecordId    *ID          `xml:"RelatedRecordId,omitempty"`
	SystemModstamp     time.Time    `xml:"SystemModstamp,omitempty"`
	Title              string       `xml:"Title,omitempty"`
	Type               string       `xml:"Type,omitempty"`
}

type SamlSsoConfig struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com SamlSsoConfig"`

	*SObject

	AttributeFormat         string     `xml:"AttributeFormat,omitempty"`
	AttributeName           string     `xml:"AttributeName,omitempty"`
	Audience                string     `xml:"Audience,omitempty"`
	CreatedBy               *User      `xml:"CreatedBy,omitempty"`
	CreatedById             *ID        `xml:"CreatedById,omitempty"`
	CreatedDate             time.Time  `xml:"CreatedDate,omitempty"`
	DeveloperName           string     `xml:"DeveloperName,omitempty"`
	ErrorUrl                string     `xml:"ErrorUrl,omitempty"`
	ExecutionUser           *User      `xml:"ExecutionUser,omitempty"`
	ExecutionUserId         *ID        `xml:"ExecutionUserId,omitempty"`
	IdentityLocation        string     `xml:"IdentityLocation,omitempty"`
	IdentityMapping         string     `xml:"IdentityMapping,omitempty"`
	IsDeleted               bool       `xml:"IsDeleted,omitempty"`
	Issuer                  string     `xml:"Issuer,omitempty"`
	Language                string     `xml:"Language,omitempty"`
	LastModifiedBy          *User      `xml:"LastModifiedBy,omitempty"`
	LastModifiedById        *ID        `xml:"LastModifiedById,omitempty"`
	LastModifiedDate        time.Time  `xml:"LastModifiedDate,omitempty"`
	LoginUrl                string     `xml:"LoginUrl,omitempty"`
	LogoutUrl               string     `xml:"LogoutUrl,omitempty"`
	MasterLabel             string     `xml:"MasterLabel,omitempty"`
	NamespacePrefix         string     `xml:"NamespacePrefix,omitempty"`
	OptionsSpInitBinding    bool       `xml:"OptionsSpInitBinding,omitempty"`
	OptionsUserProvisioning bool       `xml:"OptionsUserProvisioning,omitempty"`
	RequestSignatureMethod  string     `xml:"RequestSignatureMethod,omitempty"`
	SamlJitHandler          *ApexClass `xml:"SamlJitHandler,omitempty"`
	SamlJitHandlerId        *ID        `xml:"SamlJitHandlerId,omitempty"`
	SingleLogoutBinding     string     `xml:"SingleLogoutBinding,omitempty"`
	SingleLogoutUrl         string     `xml:"SingleLogoutUrl,omitempty"`
	SystemModstamp          time.Time  `xml:"SystemModstamp,omitempty"`
	ValidationCert          string     `xml:"ValidationCert,omitempty"`
	Version                 string     `xml:"Version,omitempty"`
}

type Scontrol struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com Scontrol"`

	*SObject

	Binary           []byte    `xml:"Binary,omitempty"`
	BodyLength       int32     `xml:"BodyLength,omitempty"`
	ContentSource    string    `xml:"ContentSource,omitempty"`
	CreatedBy        *User     `xml:"CreatedBy,omitempty"`
	CreatedById      *ID       `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time `xml:"CreatedDate,omitempty"`
	Description      string    `xml:"Description,omitempty"`
	DeveloperName    string    `xml:"DeveloperName,omitempty"`
	EncodingKey      string    `xml:"EncodingKey,omitempty"`
	Filename         string    `xml:"Filename,omitempty"`
	HtmlWrapper      string    `xml:"HtmlWrapper,omitempty"`
	IsDeleted        bool      `xml:"IsDeleted,omitempty"`
	LastModifiedBy   *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time `xml:"LastModifiedDate,omitempty"`
	Name             string    `xml:"Name,omitempty"`
	NamespacePrefix  string    `xml:"NamespacePrefix,omitempty"`
	SupportsCaching  bool      `xml:"SupportsCaching,omitempty"`
	SystemModstamp   time.Time `xml:"SystemModstamp,omitempty"`
}

type SearchActivity struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com SearchActivity"`

	*SObject

	AvgNumResults     float64           `xml:"AvgNumResults,omitempty"`
	ClickRank         float64           `xml:"ClickRank,omitempty"`
	ClickedRecordName string            `xml:"ClickedRecordName,omitempty"`
	CountQueries      int32             `xml:"CountQueries,omitempty"`
	CountUsers        int32             `xml:"CountUsers,omitempty"`
	CreatedBy         *User             `xml:"CreatedBy,omitempty"`
	CreatedById       *ID               `xml:"CreatedById,omitempty"`
	CreatedDate       time.Time         `xml:"CreatedDate,omitempty"`
	IsDeleted         bool              `xml:"IsDeleted,omitempty"`
	KbChannel         string            `xml:"KbChannel,omitempty"`
	LastModifiedBy    *User             `xml:"LastModifiedBy,omitempty"`
	LastModifiedById  *ID               `xml:"LastModifiedById,omitempty"`
	LastModifiedDate  time.Time         `xml:"LastModifiedDate,omitempty"`
	Name              string            `xml:"Name,omitempty"`
	Period            string            `xml:"Period,omitempty"`
	QueryDate         time.Time         `xml:"QueryDate,omitempty"`
	QueryLanguage     string            `xml:"QueryLanguage,omitempty"`
	SearchTerm        string            `xml:"SearchTerm,omitempty"`
	SystemModstamp    time.Time         `xml:"SystemModstamp,omitempty"`
	UserRecordAccess  *UserRecordAccess `xml:"UserRecordAccess,omitempty"`
}

type SearchLayout struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com SearchLayout"`

	*SObject

	ButtonsDisplayed   *SearchLayoutButtonsDisplayed `xml:"ButtonsDisplayed,omitempty"`
	DurableId          string                        `xml:"DurableId,omitempty"`
	EntityDefinition   *EntityDefinition             `xml:"EntityDefinition,omitempty"`
	EntityDefinitionId string                        `xml:"EntityDefinitionId,omitempty"`
	FieldsDisplayed    *SearchLayoutFieldsDisplayed  `xml:"FieldsDisplayed,omitempty"`
	Label              string                        `xml:"Label,omitempty"`
	LastModifiedBy     *User                         `xml:"LastModifiedBy,omitempty"`
	LastModifiedById   *ID                           `xml:"LastModifiedById,omitempty"`
	LastModifiedDate   time.Time                     `xml:"LastModifiedDate,omitempty"`
	LayoutType         string                        `xml:"LayoutType,omitempty"`
}

type SearchPromotionRule struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com SearchPromotionRule"`

	*SObject

	CreatedBy        *User             `xml:"CreatedBy,omitempty"`
	CreatedById      *ID               `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time         `xml:"CreatedDate,omitempty"`
	IsDeleted        bool              `xml:"IsDeleted,omitempty"`
	LastModifiedBy   *User             `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID               `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time         `xml:"LastModifiedDate,omitempty"`
	Query            string            `xml:"Query,omitempty"`
	SystemModstamp   time.Time         `xml:"SystemModstamp,omitempty"`
	UserRecordAccess *UserRecordAccess `xml:"UserRecordAccess,omitempty"`
}

type SecureAgent struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com SecureAgent"`

	*SObject

	AgentKey              string               `xml:"AgentKey,omitempty"`
	CreatedBy             *User                `xml:"CreatedBy,omitempty"`
	CreatedById           *ID                  `xml:"CreatedById,omitempty"`
	CreatedDate           time.Time            `xml:"CreatedDate,omitempty"`
	DeveloperName         string               `xml:"DeveloperName,omitempty"`
	IsDeleted             bool                 `xml:"IsDeleted,omitempty"`
	Language              string               `xml:"Language,omitempty"`
	LastModifiedBy        *User                `xml:"LastModifiedBy,omitempty"`
	LastModifiedById      *ID                  `xml:"LastModifiedById,omitempty"`
	LastModifiedDate      time.Time            `xml:"LastModifiedDate,omitempty"`
	MasterLabel           string               `xml:"MasterLabel,omitempty"`
	Priority              int32                `xml:"Priority,omitempty"`
	ProxyUser             *User                `xml:"ProxyUser,omitempty"`
	ProxyUserId           *ID                  `xml:"ProxyUserId,omitempty"`
	SecureAgentPlugins    *QueryResult         `xml:"SecureAgentPlugins,omitempty"`
	SecureAgentsCluster   *SecureAgentsCluster `xml:"SecureAgentsCluster,omitempty"`
	SecureAgentsClusterId *ID                  `xml:"SecureAgentsClusterId,omitempty"`
	SystemModstamp        time.Time            `xml:"SystemModstamp,omitempty"`
}

type SecureAgentPlugin struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com SecureAgentPlugin"`

	*SObject

	CreatedBy                   *User        `xml:"CreatedBy,omitempty"`
	CreatedById                 *ID          `xml:"CreatedById,omitempty"`
	CreatedDate                 time.Time    `xml:"CreatedDate,omitempty"`
	IsDeleted                   bool         `xml:"IsDeleted,omitempty"`
	LastModifiedBy              *User        `xml:"LastModifiedBy,omitempty"`
	LastModifiedById            *ID          `xml:"LastModifiedById,omitempty"`
	LastModifiedDate            time.Time    `xml:"LastModifiedDate,omitempty"`
	PluginName                  string       `xml:"PluginName,omitempty"`
	PluginType                  string       `xml:"PluginType,omitempty"`
	RequestedVersion            string       `xml:"RequestedVersion,omitempty"`
	SecureAgent                 *SecureAgent `xml:"SecureAgent,omitempty"`
	SecureAgentId               *ID          `xml:"SecureAgentId,omitempty"`
	SecureAgentPluginProperties *QueryResult `xml:"SecureAgentPluginProperties,omitempty"`
	SystemModstamp              time.Time    `xml:"SystemModstamp,omitempty"`
	UpdateWindowEnd             time.Time    `xml:"UpdateWindowEnd,omitempty"`
	UpdateWindowStart           time.Time    `xml:"UpdateWindowStart,omitempty"`
}

type SecureAgentPluginProperty struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com SecureAgentPluginProperty"`

	*SObject

	CreatedBy           *User              `xml:"CreatedBy,omitempty"`
	CreatedById         *ID                `xml:"CreatedById,omitempty"`
	CreatedDate         time.Time          `xml:"CreatedDate,omitempty"`
	IsDeleted           bool               `xml:"IsDeleted,omitempty"`
	LastModifiedBy      *User              `xml:"LastModifiedBy,omitempty"`
	LastModifiedById    *ID                `xml:"LastModifiedById,omitempty"`
	LastModifiedDate    time.Time          `xml:"LastModifiedDate,omitempty"`
	PropertyName        string             `xml:"PropertyName,omitempty"`
	PropertyValue       string             `xml:"PropertyValue,omitempty"`
	SecureAgentPlugin   *SecureAgentPlugin `xml:"SecureAgentPlugin,omitempty"`
	SecureAgentPluginId *ID                `xml:"SecureAgentPluginId,omitempty"`
	SystemModstamp      time.Time          `xml:"SystemModstamp,omitempty"`
}

type SecureAgentsCluster struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com SecureAgentsCluster"`

	*SObject

	CreatedBy        *User     `xml:"CreatedBy,omitempty"`
	CreatedById      *ID       `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time `xml:"CreatedDate,omitempty"`
	Description      string    `xml:"Description,omitempty"`
	DeveloperName    string    `xml:"DeveloperName,omitempty"`
	IsDeleted        bool      `xml:"IsDeleted,omitempty"`
	Language         string    `xml:"Language,omitempty"`
	LastModifiedBy   *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time `xml:"LastModifiedDate,omitempty"`
	MasterLabel      string    `xml:"MasterLabel,omitempty"`
	SystemModstamp   time.Time `xml:"SystemModstamp,omitempty"`
}

type SecurityCustomBaseline struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com SecurityCustomBaseline"`

	*SObject

	Baseline         string    `xml:"Baseline,omitempty"`
	CreatedBy        *User     `xml:"CreatedBy,omitempty"`
	CreatedById      *ID       `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time `xml:"CreatedDate,omitempty"`
	DeveloperName    string    `xml:"DeveloperName,omitempty"`
	IsDefault        bool      `xml:"IsDefault,omitempty"`
	IsDeleted        bool      `xml:"IsDeleted,omitempty"`
	Language         string    `xml:"Language,omitempty"`
	LastModifiedBy   *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time `xml:"LastModifiedDate,omitempty"`
	MasterLabel      string    `xml:"MasterLabel,omitempty"`
	NamespacePrefix  string    `xml:"NamespacePrefix,omitempty"`
	SystemModstamp   time.Time `xml:"SystemModstamp,omitempty"`
}

type SessionPermSetActivation struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com SessionPermSetActivation"`

	*SObject

	AuthSession      *AuthSession   `xml:"AuthSession,omitempty"`
	AuthSessionId    *ID            `xml:"AuthSessionId,omitempty"`
	CreatedBy        *User          `xml:"CreatedBy,omitempty"`
	CreatedById      *ID            `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time      `xml:"CreatedDate,omitempty"`
	Description      string         `xml:"Description,omitempty"`
	IsDeleted        bool           `xml:"IsDeleted,omitempty"`
	LastModifiedBy   *User          `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID            `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time      `xml:"LastModifiedDate,omitempty"`
	PermissionSet    *PermissionSet `xml:"PermissionSet,omitempty"`
	PermissionSetId  *ID            `xml:"PermissionSetId,omitempty"`
	SystemModstamp   time.Time      `xml:"SystemModstamp,omitempty"`
	User             *User          `xml:"User,omitempty"`
	UserId           *ID            `xml:"UserId,omitempty"`
}

type SetupAuditTrail struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com SetupAuditTrail"`

	*SObject

	Action                     string    `xml:"Action,omitempty"`
	CreatedBy                  *User     `xml:"CreatedBy,omitempty"`
	CreatedById                *ID       `xml:"CreatedById,omitempty"`
	CreatedDate                time.Time `xml:"CreatedDate,omitempty"`
	DelegateUser               string    `xml:"DelegateUser,omitempty"`
	Display                    string    `xml:"Display,omitempty"`
	ResponsibleNamespacePrefix string    `xml:"ResponsibleNamespacePrefix,omitempty"`
	Section                    string    `xml:"Section,omitempty"`
}

type SetupEntityAccess struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com SetupEntityAccess"`

	*SObject

	Parent          *PermissionSet `xml:"Parent,omitempty"`
	ParentId        *ID            `xml:"ParentId,omitempty"`
	SetupEntityId   *ID            `xml:"SetupEntityId,omitempty"`
	SetupEntityType string         `xml:"SetupEntityType,omitempty"`
	SystemModstamp  time.Time      `xml:"SystemModstamp,omitempty"`
}

type Site struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com Site"`

	*SObject

	Admin                                      *User        `xml:"Admin,omitempty"`
	AdminId                                    *ID          `xml:"AdminId,omitempty"`
	AnalyticsTrackingCode                      string       `xml:"AnalyticsTrackingCode,omitempty"`
	AttachedContentDocuments                   *QueryResult `xml:"AttachedContentDocuments,omitempty"`
	ClickjackProtectionLevel                   string       `xml:"ClickjackProtectionLevel,omitempty"`
	CombinedAttachments                        *QueryResult `xml:"CombinedAttachments,omitempty"`
	ContentDocumentLinks                       *QueryResult `xml:"ContentDocumentLinks,omitempty"`
	CreatedBy                                  *User        `xml:"CreatedBy,omitempty"`
	CreatedById                                *ID          `xml:"CreatedById,omitempty"`
	CreatedDate                                time.Time    `xml:"CreatedDate,omitempty"`
	DailyBandwidthLimit                        int32        `xml:"DailyBandwidthLimit,omitempty"`
	DailyBandwidthUsed                         int32        `xml:"DailyBandwidthUsed,omitempty"`
	DailyRequestTimeLimit                      int32        `xml:"DailyRequestTimeLimit,omitempty"`
	DailyRequestTimeUsed                       int32        `xml:"DailyRequestTimeUsed,omitempty"`
	Description                                string       `xml:"Description,omitempty"`
	FeedSubscriptionsForEntity                 *QueryResult `xml:"FeedSubscriptionsForEntity,omitempty"`
	Feeds                                      *QueryResult `xml:"Feeds,omitempty"`
	GuestUser                                  *User        `xml:"GuestUser,omitempty"`
	GuestUserId                                *ID          `xml:"GuestUserId,omitempty"`
	Histories                                  *QueryResult `xml:"Histories,omitempty"`
	LastModifiedBy                             *User        `xml:"LastModifiedBy,omitempty"`
	LastModifiedById                           *ID          `xml:"LastModifiedById,omitempty"`
	LastModifiedDate                           time.Time    `xml:"LastModifiedDate,omitempty"`
	MasterLabel                                string       `xml:"MasterLabel,omitempty"`
	MonthlyPageViewsEntitlement                int32        `xml:"MonthlyPageViewsEntitlement,omitempty"`
	Name                                       string       `xml:"Name,omitempty"`
	OptionsAllowGuestSupportApi                bool         `xml:"OptionsAllowGuestSupportApi,omitempty"`
	OptionsAllowHomePage                       bool         `xml:"OptionsAllowHomePage,omitempty"`
	OptionsAllowStandardAnswersPages           bool         `xml:"OptionsAllowStandardAnswersPages,omitempty"`
	OptionsAllowStandardIdeasPages             bool         `xml:"OptionsAllowStandardIdeasPages,omitempty"`
	OptionsAllowStandardLookups                bool         `xml:"OptionsAllowStandardLookups,omitempty"`
	OptionsAllowStandardPortalPages            bool         `xml:"OptionsAllowStandardPortalPages,omitempty"`
	OptionsAllowStandardSearch                 bool         `xml:"OptionsAllowStandardSearch,omitempty"`
	OptionsBrowserXssProtection                bool         `xml:"OptionsBrowserXssProtection,omitempty"`
	OptionsContentSniffingProtection           bool         `xml:"OptionsContentSniffingProtection,omitempty"`
	OptionsCspUpgradeInsecureRequests          bool         `xml:"OptionsCspUpgradeInsecureRequests,omitempty"`
	OptionsEnableFeeds                         bool         `xml:"OptionsEnableFeeds,omitempty"`
	OptionsReferrerPolicyOriginWhenCrossOrigin bool         `xml:"OptionsReferrerPolicyOriginWhenCrossOrigin,omitempty"`
	OptionsRequireHttps                        bool         `xml:"OptionsRequireHttps,omitempty"`
	SiteDomainPaths                            *QueryResult `xml:"SiteDomainPaths,omitempty"`
	SiteIframeWhiteListUrls                    *QueryResult `xml:"SiteIframeWhiteListUrls,omitempty"`
	SiteType                                   string       `xml:"SiteType,omitempty"`
	Status                                     string       `xml:"Status,omitempty"`
	Subdomain                                  string       `xml:"Subdomain,omitempty"`
	SystemModstamp                             time.Time    `xml:"SystemModstamp,omitempty"`
	UrlPathPrefix                              string       `xml:"UrlPathPrefix,omitempty"`
}

type SiteFeed struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com SiteFeed"`

	*SObject

	BestComment        *FeedComment `xml:"BestComment,omitempty"`
	BestCommentId      *ID          `xml:"BestCommentId,omitempty"`
	Body               string       `xml:"Body,omitempty"`
	CommentCount       int32        `xml:"CommentCount,omitempty"`
	CreatedBy          *SObject     `xml:"CreatedBy,omitempty"`
	CreatedById        *ID          `xml:"CreatedById,omitempty"`
	CreatedDate        time.Time    `xml:"CreatedDate,omitempty"`
	FeedAttachments    *QueryResult `xml:"FeedAttachments,omitempty"`
	FeedComments       *QueryResult `xml:"FeedComments,omitempty"`
	FeedLikes          *QueryResult `xml:"FeedLikes,omitempty"`
	FeedSignals        *QueryResult `xml:"FeedSignals,omitempty"`
	FeedTrackedChanges *QueryResult `xml:"FeedTrackedChanges,omitempty"`
	InsertedBy         *SObject     `xml:"InsertedBy,omitempty"`
	InsertedById       *ID          `xml:"InsertedById,omitempty"`
	IsDeleted          bool         `xml:"IsDeleted,omitempty"`
	IsRichText         bool         `xml:"IsRichText,omitempty"`
	LastModifiedDate   time.Time    `xml:"LastModifiedDate,omitempty"`
	LikeCount          int32        `xml:"LikeCount,omitempty"`
	LinkUrl            string       `xml:"LinkUrl,omitempty"`
	Parent             *Site        `xml:"Parent,omitempty"`
	ParentId           *ID          `xml:"ParentId,omitempty"`
	RelatedRecordId    *ID          `xml:"RelatedRecordId,omitempty"`
	SystemModstamp     time.Time    `xml:"SystemModstamp,omitempty"`
	Title              string       `xml:"Title,omitempty"`
	Type               string       `xml:"Type,omitempty"`
}

type SiteHistory struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com SiteHistory"`

	*SObject

	CreatedBy   *SObject    `xml:"CreatedBy,omitempty"`
	CreatedById *ID         `xml:"CreatedById,omitempty"`
	CreatedDate time.Time   `xml:"CreatedDate,omitempty"`
	Field       string      `xml:"Field,omitempty"`
	IsDeleted   bool        `xml:"IsDeleted,omitempty"`
	NewValue    interface{} `xml:"NewValue,omitempty"`
	OldValue    interface{} `xml:"OldValue,omitempty"`
	Site        *Site       `xml:"Site,omitempty"`
	SiteId      *ID         `xml:"SiteId,omitempty"`
}

type SiteIframeWhiteListUrl struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com SiteIframeWhiteListUrl"`

	*SObject

	CreatedBy        *User     `xml:"CreatedBy,omitempty"`
	CreatedById      *ID       `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time `xml:"CreatedDate,omitempty"`
	IsDeleted        bool      `xml:"IsDeleted,omitempty"`
	LastModifiedBy   *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time `xml:"LastModifiedDate,omitempty"`
	Site             *Site     `xml:"Site,omitempty"`
	SiteId           *ID       `xml:"SiteId,omitempty"`
	SystemModstamp   time.Time `xml:"SystemModstamp,omitempty"`
	Url              string    `xml:"Url,omitempty"`
}

type Solution struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com Solution"`

	*SObject

	ActivityHistories          *QueryResult      `xml:"ActivityHistories,omitempty"`
	AttachedContentDocuments   *QueryResult      `xml:"AttachedContentDocuments,omitempty"`
	Attachments                *QueryResult      `xml:"Attachments,omitempty"`
	CaseSolutions              *QueryResult      `xml:"CaseSolutions,omitempty"`
	CombinedAttachments        *QueryResult      `xml:"CombinedAttachments,omitempty"`
	ContentDocumentLinks       *QueryResult      `xml:"ContentDocumentLinks,omitempty"`
	CreatedBy                  *User             `xml:"CreatedBy,omitempty"`
	CreatedById                *ID               `xml:"CreatedById,omitempty"`
	CreatedDate                time.Time         `xml:"CreatedDate,omitempty"`
	Emails                     *QueryResult      `xml:"Emails,omitempty"`
	Events                     *QueryResult      `xml:"Events,omitempty"`
	FeedSubscriptionsForEntity *QueryResult      `xml:"FeedSubscriptionsForEntity,omitempty"`
	Feeds                      *QueryResult      `xml:"Feeds,omitempty"`
	Histories                  *QueryResult      `xml:"Histories,omitempty"`
	IsDeleted                  bool              `xml:"IsDeleted,omitempty"`
	IsHtml                     bool              `xml:"IsHtml,omitempty"`
	IsPublished                bool              `xml:"IsPublished,omitempty"`
	IsPublishedInPublicKb      bool              `xml:"IsPublishedInPublicKb,omitempty"`
	IsReviewed                 bool              `xml:"IsReviewed,omitempty"`
	LastModifiedBy             *User             `xml:"LastModifiedBy,omitempty"`
	LastModifiedById           *ID               `xml:"LastModifiedById,omitempty"`
	LastModifiedDate           time.Time         `xml:"LastModifiedDate,omitempty"`
	LastReferencedDate         time.Time         `xml:"LastReferencedDate,omitempty"`
	LastViewedDate             time.Time         `xml:"LastViewedDate,omitempty"`
	LookedUpFromActivities     *QueryResult      `xml:"LookedUpFromActivities,omitempty"`
	OpenActivities             *QueryResult      `xml:"OpenActivities,omitempty"`
	Owner                      *User             `xml:"Owner,omitempty"`
	OwnerId                    *ID               `xml:"OwnerId,omitempty"`
	ProcessInstances           *QueryResult      `xml:"ProcessInstances,omitempty"`
	ProcessSteps               *QueryResult      `xml:"ProcessSteps,omitempty"`
	SolutionName               string            `xml:"SolutionName,omitempty"`
	SolutionNote               string            `xml:"SolutionNote,omitempty"`
	SolutionNumber             string            `xml:"SolutionNumber,omitempty"`
	Status                     string            `xml:"Status,omitempty"`
	SystemModstamp             time.Time         `xml:"SystemModstamp,omitempty"`
	Tasks                      *QueryResult      `xml:"Tasks,omitempty"`
	TimesUsed                  int32             `xml:"TimesUsed,omitempty"`
	TopicAssignments           *QueryResult      `xml:"TopicAssignments,omitempty"`
	UserRecordAccess           *UserRecordAccess `xml:"UserRecordAccess,omitempty"`
	Votes                      *QueryResult      `xml:"Votes,omitempty"`
}

type SolutionFeed struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com SolutionFeed"`

	*SObject

	BestComment        *FeedComment `xml:"BestComment,omitempty"`
	BestCommentId      *ID          `xml:"BestCommentId,omitempty"`
	Body               string       `xml:"Body,omitempty"`
	CommentCount       int32        `xml:"CommentCount,omitempty"`
	CreatedBy          *SObject     `xml:"CreatedBy,omitempty"`
	CreatedById        *ID          `xml:"CreatedById,omitempty"`
	CreatedDate        time.Time    `xml:"CreatedDate,omitempty"`
	FeedAttachments    *QueryResult `xml:"FeedAttachments,omitempty"`
	FeedComments       *QueryResult `xml:"FeedComments,omitempty"`
	FeedLikes          *QueryResult `xml:"FeedLikes,omitempty"`
	FeedSignals        *QueryResult `xml:"FeedSignals,omitempty"`
	FeedTrackedChanges *QueryResult `xml:"FeedTrackedChanges,omitempty"`
	InsertedBy         *SObject     `xml:"InsertedBy,omitempty"`
	InsertedById       *ID          `xml:"InsertedById,omitempty"`
	IsDeleted          bool         `xml:"IsDeleted,omitempty"`
	IsRichText         bool         `xml:"IsRichText,omitempty"`
	LastModifiedDate   time.Time    `xml:"LastModifiedDate,omitempty"`
	LikeCount          int32        `xml:"LikeCount,omitempty"`
	LinkUrl            string       `xml:"LinkUrl,omitempty"`
	Parent             *Solution    `xml:"Parent,omitempty"`
	ParentId           *ID          `xml:"ParentId,omitempty"`
	RelatedRecordId    *ID          `xml:"RelatedRecordId,omitempty"`
	SystemModstamp     time.Time    `xml:"SystemModstamp,omitempty"`
	Title              string       `xml:"Title,omitempty"`
	Type               string       `xml:"Type,omitempty"`
}

type SolutionHistory struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com SolutionHistory"`

	*SObject

	CreatedBy   *SObject    `xml:"CreatedBy,omitempty"`
	CreatedById *ID         `xml:"CreatedById,omitempty"`
	CreatedDate time.Time   `xml:"CreatedDate,omitempty"`
	Field       string      `xml:"Field,omitempty"`
	IsDeleted   bool        `xml:"IsDeleted,omitempty"`
	NewValue    interface{} `xml:"NewValue,omitempty"`
	OldValue    interface{} `xml:"OldValue,omitempty"`
	Solution    *Solution   `xml:"Solution,omitempty"`
	SolutionId  *ID         `xml:"SolutionId,omitempty"`
}

type SolutionStatus struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com SolutionStatus"`

	*SObject

	ApiName          string    `xml:"ApiName,omitempty"`
	CreatedBy        *User     `xml:"CreatedBy,omitempty"`
	CreatedById      *ID       `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time `xml:"CreatedDate,omitempty"`
	IsDefault        bool      `xml:"IsDefault,omitempty"`
	IsReviewed       bool      `xml:"IsReviewed,omitempty"`
	LastModifiedBy   *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time `xml:"LastModifiedDate,omitempty"`
	MasterLabel      string    `xml:"MasterLabel,omitempty"`
	SortOrder        int32     `xml:"SortOrder,omitempty"`
	SystemModstamp   time.Time `xml:"SystemModstamp,omitempty"`
}

type Stamp struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com Stamp"`

	*SObject

	CreatedBy        *User        `xml:"CreatedBy,omitempty"`
	CreatedById      *ID          `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time    `xml:"CreatedDate,omitempty"`
	CustomBrands     *QueryResult `xml:"CustomBrands,omitempty"`
	Description      string       `xml:"Description,omitempty"`
	IsDeleted        bool         `xml:"IsDeleted,omitempty"`
	LastModifiedBy   *User        `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID          `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time    `xml:"LastModifiedDate,omitempty"`
	MasterLabel      string       `xml:"MasterLabel,omitempty"`
	Parent           *SObject     `xml:"Parent,omitempty"`
	ParentId         *ID          `xml:"ParentId,omitempty"`
	SystemModstamp   time.Time    `xml:"SystemModstamp,omitempty"`
}

type StampAssignment struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com StampAssignment"`

	*SObject

	CreatedBy        *User     `xml:"CreatedBy,omitempty"`
	CreatedById      *ID       `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time `xml:"CreatedDate,omitempty"`
	IsDeleted        bool      `xml:"IsDeleted,omitempty"`
	LastModifiedBy   *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time `xml:"LastModifiedDate,omitempty"`
	Stamp            *Stamp    `xml:"Stamp,omitempty"`
	StampId          *ID       `xml:"StampId,omitempty"`
	Subject          *User     `xml:"Subject,omitempty"`
	SubjectId        *ID       `xml:"SubjectId,omitempty"`
	SystemModstamp   time.Time `xml:"SystemModstamp,omitempty"`
}

type StaticResource struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com StaticResource"`

	*SObject

	Body             []byte    `xml:"Body,omitempty"`
	BodyLength       int32     `xml:"BodyLength,omitempty"`
	CacheControl     string    `xml:"CacheControl,omitempty"`
	ContentType      string    `xml:"ContentType,omitempty"`
	CreatedBy        *User     `xml:"CreatedBy,omitempty"`
	CreatedById      *ID       `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time `xml:"CreatedDate,omitempty"`
	Description      string    `xml:"Description,omitempty"`
	LastModifiedBy   *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time `xml:"LastModifiedDate,omitempty"`
	Name             string    `xml:"Name,omitempty"`
	NamespacePrefix  string    `xml:"NamespacePrefix,omitempty"`
	SystemModstamp   time.Time `xml:"SystemModstamp,omitempty"`
}

type StreamingChannel struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com StreamingChannel"`

	*SObject

	CreatedBy          *User             `xml:"CreatedBy,omitempty"`
	CreatedById        *ID               `xml:"CreatedById,omitempty"`
	CreatedDate        time.Time         `xml:"CreatedDate,omitempty"`
	Description        string            `xml:"Description,omitempty"`
	IsDeleted          bool              `xml:"IsDeleted,omitempty"`
	IsDynamic          bool              `xml:"IsDynamic,omitempty"`
	LastModifiedBy     *User             `xml:"LastModifiedBy,omitempty"`
	LastModifiedById   *ID               `xml:"LastModifiedById,omitempty"`
	LastModifiedDate   time.Time         `xml:"LastModifiedDate,omitempty"`
	LastReferencedDate time.Time         `xml:"LastReferencedDate,omitempty"`
	LastViewedDate     time.Time         `xml:"LastViewedDate,omitempty"`
	Name               string            `xml:"Name,omitempty"`
	Owner              *SObject          `xml:"Owner,omitempty"`
	OwnerId            *ID               `xml:"OwnerId,omitempty"`
	ProcessInstances   *QueryResult      `xml:"ProcessInstances,omitempty"`
	ProcessSteps       *QueryResult      `xml:"ProcessSteps,omitempty"`
	SystemModstamp     time.Time         `xml:"SystemModstamp,omitempty"`
	UserRecordAccess   *UserRecordAccess `xml:"UserRecordAccess,omitempty"`
}

type StreamingChannelShare struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com StreamingChannelShare"`

	*SObject

	AccessLevel      string            `xml:"AccessLevel,omitempty"`
	IsDeleted        bool              `xml:"IsDeleted,omitempty"`
	LastModifiedBy   *User             `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID               `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time         `xml:"LastModifiedDate,omitempty"`
	Parent           *StreamingChannel `xml:"Parent,omitempty"`
	ParentId         *ID               `xml:"ParentId,omitempty"`
	RowCause         string            `xml:"RowCause,omitempty"`
	UserOrGroup      *SObject          `xml:"UserOrGroup,omitempty"`
	UserOrGroupId    *ID               `xml:"UserOrGroupId,omitempty"`
}

type TabDefinition struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com TabDefinition"`

	*SObject

	AppTabs                *QueryResult `xml:"AppTabs,omitempty"`
	Colors                 *QueryResult `xml:"Colors,omitempty"`
	DurableId              string       `xml:"DurableId,omitempty"`
	Icons                  *QueryResult `xml:"Icons,omitempty"`
	IsAvailableInAloha     bool         `xml:"IsAvailableInAloha,omitempty"`
	IsAvailableInLightning bool         `xml:"IsAvailableInLightning,omitempty"`
	IsCustom               bool         `xml:"IsCustom,omitempty"`
	Label                  string       `xml:"Label,omitempty"`
	Name                   string       `xml:"Name,omitempty"`
	SobjectName            string       `xml:"SobjectName,omitempty"`
	Url                    string       `xml:"Url,omitempty"`
}

type Task struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com Task"`

	*SObject

	Account                    *Account     `xml:"Account,omitempty"`
	AccountId                  *ID          `xml:"AccountId,omitempty"`
	ActivityDate               time.Time    `xml:"ActivityDate,omitempty"`
	AttachedContentDocuments   *QueryResult `xml:"AttachedContentDocuments,omitempty"`
	Attachments                *QueryResult `xml:"Attachments,omitempty"`
	CallDisposition            string       `xml:"CallDisposition,omitempty"`
	CallDurationInSeconds      int32        `xml:"CallDurationInSeconds,omitempty"`
	CallObject                 string       `xml:"CallObject,omitempty"`
	CallType                   string       `xml:"CallType,omitempty"`
	CombinedAttachments        *QueryResult `xml:"CombinedAttachments,omitempty"`
	ContentDocumentLinks       *QueryResult `xml:"ContentDocumentLinks,omitempty"`
	CreatedBy                  *User        `xml:"CreatedBy,omitempty"`
	CreatedById                *ID          `xml:"CreatedById,omitempty"`
	CreatedDate                time.Time    `xml:"CreatedDate,omitempty"`
	Description                string       `xml:"Description,omitempty"`
	FeedSubscriptionsForEntity *QueryResult `xml:"FeedSubscriptionsForEntity,omitempty"`
	Feeds                      *QueryResult `xml:"Feeds,omitempty"`
	IsArchived                 bool         `xml:"IsArchived,omitempty"`
	IsClosed                   bool         `xml:"IsClosed,omitempty"`
	IsDeleted                  bool         `xml:"IsDeleted,omitempty"`
	IsHighPriority             bool         `xml:"IsHighPriority,omitempty"`
	IsRecurrence               bool         `xml:"IsRecurrence,omitempty"`
	IsReminderSet              bool         `xml:"IsReminderSet,omitempty"`
	LastModifiedBy             *User        `xml:"LastModifiedBy,omitempty"`
	LastModifiedById           *ID          `xml:"LastModifiedById,omitempty"`
	LastModifiedDate           time.Time    `xml:"LastModifiedDate,omitempty"`
	ListEmailActivities        *QueryResult `xml:"ListEmailActivities,omitempty"`
	Owner                      *SObject     `xml:"Owner,omitempty"`
	OwnerId                    *ID          `xml:"OwnerId,omitempty"`
	Priority                   string       `xml:"Priority,omitempty"`
	RecordActionHistories      *QueryResult `xml:"RecordActionHistories,omitempty"`
	RecordActions              *QueryResult `xml:"RecordActions,omitempty"`
	RecurrenceActivityId       *ID          `xml:"RecurrenceActivityId,omitempty"`
	RecurrenceDayOfMonth       int32        `xml:"RecurrenceDayOfMonth,omitempty"`
	RecurrenceDayOfWeekMask    int32        `xml:"RecurrenceDayOfWeekMask,omitempty"`
	RecurrenceEndDateOnly      time.Time    `xml:"RecurrenceEndDateOnly,omitempty"`
	RecurrenceInstance         string       `xml:"RecurrenceInstance,omitempty"`
	RecurrenceInterval         int32        `xml:"RecurrenceInterval,omitempty"`
	RecurrenceMonthOfYear      string       `xml:"RecurrenceMonthOfYear,omitempty"`
	RecurrenceRegeneratedType  string       `xml:"RecurrenceRegeneratedType,omitempty"`
	RecurrenceStartDateOnly    time.Time    `xml:"RecurrenceStartDateOnly,omitempty"`
	RecurrenceTimeZoneSidKey   string       `xml:"RecurrenceTimeZoneSidKey,omitempty"`
	RecurrenceType             string       `xml:"RecurrenceType,omitempty"`
	RecurringTasks             *QueryResult `xml:"RecurringTasks,omitempty"`
	ReminderDateTime           time.Time    `xml:"ReminderDateTime,omitempty"`
	Status                     string       `xml:"Status,omitempty"`
	Subject                    string       `xml:"Subject,omitempty"`
	SystemModstamp             time.Time    `xml:"SystemModstamp,omitempty"`
	TaskSubtype                string       `xml:"TaskSubtype,omitempty"`
	TopicAssignments           *QueryResult `xml:"TopicAssignments,omitempty"`
	What                       *SObject     `xml:"What,omitempty"`
	WhatId                     *ID          `xml:"WhatId,omitempty"`
	Who                        *SObject     `xml:"Who,omitempty"`
	WhoId                      *ID          `xml:"WhoId,omitempty"`
}

type TaskFeed struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com TaskFeed"`

	*SObject

	BestComment        *FeedComment `xml:"BestComment,omitempty"`
	BestCommentId      *ID          `xml:"BestCommentId,omitempty"`
	Body               string       `xml:"Body,omitempty"`
	CommentCount       int32        `xml:"CommentCount,omitempty"`
	CreatedBy          *SObject     `xml:"CreatedBy,omitempty"`
	CreatedById        *ID          `xml:"CreatedById,omitempty"`
	CreatedDate        time.Time    `xml:"CreatedDate,omitempty"`
	FeedAttachments    *QueryResult `xml:"FeedAttachments,omitempty"`
	FeedComments       *QueryResult `xml:"FeedComments,omitempty"`
	FeedLikes          *QueryResult `xml:"FeedLikes,omitempty"`
	FeedSignals        *QueryResult `xml:"FeedSignals,omitempty"`
	FeedTrackedChanges *QueryResult `xml:"FeedTrackedChanges,omitempty"`
	InsertedBy         *SObject     `xml:"InsertedBy,omitempty"`
	InsertedById       *ID          `xml:"InsertedById,omitempty"`
	IsDeleted          bool         `xml:"IsDeleted,omitempty"`
	IsRichText         bool         `xml:"IsRichText,omitempty"`
	LastModifiedDate   time.Time    `xml:"LastModifiedDate,omitempty"`
	LikeCount          int32        `xml:"LikeCount,omitempty"`
	LinkUrl            string       `xml:"LinkUrl,omitempty"`
	Parent             *Task        `xml:"Parent,omitempty"`
	ParentId           *ID          `xml:"ParentId,omitempty"`
	RelatedRecordId    *ID          `xml:"RelatedRecordId,omitempty"`
	SystemModstamp     time.Time    `xml:"SystemModstamp,omitempty"`
	Title              string       `xml:"Title,omitempty"`
	Type               string       `xml:"Type,omitempty"`
}

type TaskPriority struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com TaskPriority"`

	*SObject

	ApiName          string    `xml:"ApiName,omitempty"`
	CreatedBy        *User     `xml:"CreatedBy,omitempty"`
	CreatedById      *ID       `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time `xml:"CreatedDate,omitempty"`
	IsDefault        bool      `xml:"IsDefault,omitempty"`
	IsHighPriority   bool      `xml:"IsHighPriority,omitempty"`
	LastModifiedBy   *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time `xml:"LastModifiedDate,omitempty"`
	MasterLabel      string    `xml:"MasterLabel,omitempty"`
	SortOrder        int32     `xml:"SortOrder,omitempty"`
	SystemModstamp   time.Time `xml:"SystemModstamp,omitempty"`
}

type TaskStatus struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com TaskStatus"`

	*SObject

	ApiName          string    `xml:"ApiName,omitempty"`
	CreatedBy        *User     `xml:"CreatedBy,omitempty"`
	CreatedById      *ID       `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time `xml:"CreatedDate,omitempty"`
	IsClosed         bool      `xml:"IsClosed,omitempty"`
	IsDefault        bool      `xml:"IsDefault,omitempty"`
	LastModifiedBy   *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time `xml:"LastModifiedDate,omitempty"`
	MasterLabel      string    `xml:"MasterLabel,omitempty"`
	SortOrder        int32     `xml:"SortOrder,omitempty"`
	SystemModstamp   time.Time `xml:"SystemModstamp,omitempty"`
}

type TenantUsageEntitlement struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com TenantUsageEntitlement"`

	*SObject

	AmountUsed           float64   `xml:"AmountUsed,omitempty"`
	CreatedBy            *User     `xml:"CreatedBy,omitempty"`
	CreatedById          *ID       `xml:"CreatedById,omitempty"`
	CreatedDate          time.Time `xml:"CreatedDate,omitempty"`
	CurrentAmountAllowed float64   `xml:"CurrentAmountAllowed,omitempty"`
	EndDate              time.Time `xml:"EndDate,omitempty"`
	Frequency            string    `xml:"Frequency,omitempty"`
	HasRollover          bool      `xml:"HasRollover,omitempty"`
	IsDeleted            bool      `xml:"IsDeleted,omitempty"`
	IsPersistentResource bool      `xml:"IsPersistentResource,omitempty"`
	LastModifiedBy       *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById     *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate     time.Time `xml:"LastModifiedDate,omitempty"`
	MasterLabel          string    `xml:"MasterLabel,omitempty"`
	OverageGrace         float64   `xml:"OverageGrace,omitempty"`
	ResourceGroupKey     string    `xml:"ResourceGroupKey,omitempty"`
	Setting              string    `xml:"Setting,omitempty"`
	StartDate            time.Time `xml:"StartDate,omitempty"`
	SystemModstamp       time.Time `xml:"SystemModstamp,omitempty"`
	UsageDate            time.Time `xml:"UsageDate,omitempty"`
}

type TestSuiteMembership struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com TestSuiteMembership"`

	*SObject

	ApexClass        *ApexClass     `xml:"ApexClass,omitempty"`
	ApexClassId      *ID            `xml:"ApexClassId,omitempty"`
	ApexTestSuite    *ApexTestSuite `xml:"ApexTestSuite,omitempty"`
	ApexTestSuiteId  *ID            `xml:"ApexTestSuiteId,omitempty"`
	CreatedBy        *User          `xml:"CreatedBy,omitempty"`
	CreatedById      *ID            `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time      `xml:"CreatedDate,omitempty"`
	IsDeleted        bool           `xml:"IsDeleted,omitempty"`
	LastModifiedBy   *User          `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID            `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time      `xml:"LastModifiedDate,omitempty"`
	SystemModstamp   time.Time      `xml:"SystemModstamp,omitempty"`
}

type ThirdPartyAccountLink struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com ThirdPartyAccountLink"`

	*SObject

	Handle                   string        `xml:"Handle,omitempty"`
	IsNotSsoUsable           bool          `xml:"IsNotSsoUsable,omitempty"`
	Provider                 string        `xml:"Provider,omitempty"`
	RemoteIdentifier         string        `xml:"RemoteIdentifier,omitempty"`
	SsoProvider              *AuthProvider `xml:"SsoProvider,omitempty"`
	SsoProviderId            *ID           `xml:"SsoProviderId,omitempty"`
	SsoProviderName          string        `xml:"SsoProviderName,omitempty"`
	ThirdPartyAccountLinkKey string        `xml:"ThirdPartyAccountLinkKey,omitempty"`
	User                     *User         `xml:"User,omitempty"`
	UserId                   *ID           `xml:"UserId,omitempty"`
}

type TodayGoal struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com TodayGoal"`

	*SObject

	CreatedBy        *User             `xml:"CreatedBy,omitempty"`
	CreatedById      *ID               `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time         `xml:"CreatedDate,omitempty"`
	IsDeleted        bool              `xml:"IsDeleted,omitempty"`
	LastModifiedBy   *User             `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID               `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time         `xml:"LastModifiedDate,omitempty"`
	Name             string            `xml:"Name,omitempty"`
	Owner            *SObject          `xml:"Owner,omitempty"`
	OwnerId          *ID               `xml:"OwnerId,omitempty"`
	SystemModstamp   time.Time         `xml:"SystemModstamp,omitempty"`
	User             *User             `xml:"User,omitempty"`
	UserId           *ID               `xml:"UserId,omitempty"`
	UserRecordAccess *UserRecordAccess `xml:"UserRecordAccess,omitempty"`
	Value            float64           `xml:"Value,omitempty"`
}

type TodayGoalShare struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com TodayGoalShare"`

	*SObject

	AccessLevel      string     `xml:"AccessLevel,omitempty"`
	IsDeleted        bool       `xml:"IsDeleted,omitempty"`
	LastModifiedBy   *User      `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID        `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time  `xml:"LastModifiedDate,omitempty"`
	Parent           *TodayGoal `xml:"Parent,omitempty"`
	ParentId         *ID        `xml:"ParentId,omitempty"`
	RowCause         string     `xml:"RowCause,omitempty"`
	UserOrGroup      *SObject   `xml:"UserOrGroup,omitempty"`
	UserOrGroupId    *ID        `xml:"UserOrGroupId,omitempty"`
}

type Topic struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com Topic"`

	*SObject

	ContentDocumentLinks       *QueryResult `xml:"ContentDocumentLinks,omitempty"`
	CreatedBy                  *User        `xml:"CreatedBy,omitempty"`
	CreatedById                *ID          `xml:"CreatedById,omitempty"`
	CreatedDate                time.Time    `xml:"CreatedDate,omitempty"`
	CustomBrands               *QueryResult `xml:"CustomBrands,omitempty"`
	Description                string       `xml:"Description,omitempty"`
	FeedSubscriptionsForEntity *QueryResult `xml:"FeedSubscriptionsForEntity,omitempty"`
	Feeds                      *QueryResult `xml:"Feeds,omitempty"`
	Name                       string       `xml:"Name,omitempty"`
	SystemModstamp             time.Time    `xml:"SystemModstamp,omitempty"`
	TalkingAbout               int32        `xml:"TalkingAbout,omitempty"`
	TopicAssignments           *QueryResult `xml:"TopicAssignments,omitempty"`
}

type TopicAssignment struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com TopicAssignment"`

	*SObject

	CreatedBy       *User     `xml:"CreatedBy,omitempty"`
	CreatedById     *ID       `xml:"CreatedById,omitempty"`
	CreatedDate     time.Time `xml:"CreatedDate,omitempty"`
	Entity          *Contract `xml:"Entity,omitempty"`
	EntityId        *ID       `xml:"EntityId,omitempty"`
	EntityKeyPrefix string    `xml:"EntityKeyPrefix,omitempty"`
	EntityType      string    `xml:"EntityType,omitempty"`
	IsDeleted       bool      `xml:"IsDeleted,omitempty"`
	SystemModstamp  time.Time `xml:"SystemModstamp,omitempty"`
	Topic           *Topic    `xml:"Topic,omitempty"`
	TopicId         *ID       `xml:"TopicId,omitempty"`
}

type TopicFeed struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com TopicFeed"`

	*SObject

	BestComment        *FeedComment `xml:"BestComment,omitempty"`
	BestCommentId      *ID          `xml:"BestCommentId,omitempty"`
	Body               string       `xml:"Body,omitempty"`
	CommentCount       int32        `xml:"CommentCount,omitempty"`
	CreatedBy          *SObject     `xml:"CreatedBy,omitempty"`
	CreatedById        *ID          `xml:"CreatedById,omitempty"`
	CreatedDate        time.Time    `xml:"CreatedDate,omitempty"`
	FeedAttachments    *QueryResult `xml:"FeedAttachments,omitempty"`
	FeedComments       *QueryResult `xml:"FeedComments,omitempty"`
	FeedLikes          *QueryResult `xml:"FeedLikes,omitempty"`
	FeedSignals        *QueryResult `xml:"FeedSignals,omitempty"`
	FeedTrackedChanges *QueryResult `xml:"FeedTrackedChanges,omitempty"`
	InsertedBy         *SObject     `xml:"InsertedBy,omitempty"`
	InsertedById       *ID          `xml:"InsertedById,omitempty"`
	IsDeleted          bool         `xml:"IsDeleted,omitempty"`
	IsRichText         bool         `xml:"IsRichText,omitempty"`
	LastModifiedDate   time.Time    `xml:"LastModifiedDate,omitempty"`
	LikeCount          int32        `xml:"LikeCount,omitempty"`
	LinkUrl            string       `xml:"LinkUrl,omitempty"`
	Parent             *Topic       `xml:"Parent,omitempty"`
	ParentId           *ID          `xml:"ParentId,omitempty"`
	RelatedRecordId    *ID          `xml:"RelatedRecordId,omitempty"`
	SystemModstamp     time.Time    `xml:"SystemModstamp,omitempty"`
	Title              string       `xml:"Title,omitempty"`
	Type               string       `xml:"Type,omitempty"`
}

type TopicUserEvent struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com TopicUserEvent"`

	*SObject

	ActionEnum  string    `xml:"ActionEnum,omitempty"`
	CreatedDate time.Time `xml:"CreatedDate,omitempty"`
	TopicId     *ID       `xml:"TopicId,omitempty"`
	UserId      *ID       `xml:"UserId,omitempty"`
}

type TransactionSecurityPolicy struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com TransactionSecurityPolicy"`

	*SObject

	ActionConfig     string     `xml:"ActionConfig,omitempty"`
	ApexPolicy       *ApexClass `xml:"ApexPolicy,omitempty"`
	ApexPolicyId     *ID        `xml:"ApexPolicyId,omitempty"`
	CreatedBy        *User      `xml:"CreatedBy,omitempty"`
	CreatedById      *ID        `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time  `xml:"CreatedDate,omitempty"`
	Description      string     `xml:"Description,omitempty"`
	DeveloperName    string     `xml:"DeveloperName,omitempty"`
	EventType        string     `xml:"EventType,omitempty"`
	ExecutionUser    *User      `xml:"ExecutionUser,omitempty"`
	ExecutionUserId  *ID        `xml:"ExecutionUserId,omitempty"`
	IsDeleted        bool       `xml:"IsDeleted,omitempty"`
	Language         string     `xml:"Language,omitempty"`
	LastModifiedBy   *User      `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID        `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time  `xml:"LastModifiedDate,omitempty"`
	MasterLabel      string     `xml:"MasterLabel,omitempty"`
	NamespacePrefix  string     `xml:"NamespacePrefix,omitempty"`
	ResourceName     string     `xml:"ResourceName,omitempty"`
	State            string     `xml:"State,omitempty"`
	SystemModstamp   time.Time  `xml:"SystemModstamp,omitempty"`
	Type             string     `xml:"Type,omitempty"`
}

type UndecidedEventRelation struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com UndecidedEventRelation"`

	*SObject

	CreatedBy        *User     `xml:"CreatedBy,omitempty"`
	CreatedById      *ID       `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time `xml:"CreatedDate,omitempty"`
	Event            *Event    `xml:"Event,omitempty"`
	EventId          *ID       `xml:"EventId,omitempty"`
	IsDeleted        bool      `xml:"IsDeleted,omitempty"`
	LastModifiedBy   *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time `xml:"LastModifiedDate,omitempty"`
	Relation         *SObject  `xml:"Relation,omitempty"`
	RelationId       *ID       `xml:"RelationId,omitempty"`
	RespondedDate    time.Time `xml:"RespondedDate,omitempty"`
	Response         string    `xml:"Response,omitempty"`
	SystemModstamp   time.Time `xml:"SystemModstamp,omitempty"`
	Type             string    `xml:"Type,omitempty"`
}

type User struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com User"`

	*SObject

	AboutMe                                            string       `xml:"AboutMe,omitempty"`
	AcceptedEventRelations                             *QueryResult `xml:"AcceptedEventRelations,omitempty"`
	Account                                            *Account     `xml:"Account,omitempty"`
	AccountCleanInfoReviewers                          *QueryResult `xml:"AccountCleanInfoReviewers,omitempty"`
	AccountId                                          *ID          `xml:"AccountId,omitempty"`
	Address                                            *Address     `xml:"Address,omitempty"`
	Alias                                              string       `xml:"Alias,omitempty"`
	AttachedContentDocuments                           *QueryResult `xml:"AttachedContentDocuments,omitempty"`
	BadgeText                                          string       `xml:"BadgeText,omitempty"`
	BannerPhotoUrl                                     string       `xml:"BannerPhotoUrl,omitempty"`
	CallCenterId                                       *ID          `xml:"CallCenterId,omitempty"`
	City                                               string       `xml:"City,omitempty"`
	CombinedAttachments                                *QueryResult `xml:"CombinedAttachments,omitempty"`
	CommunityNickname                                  string       `xml:"CommunityNickname,omitempty"`
	CompanyName                                        string       `xml:"CompanyName,omitempty"`
	Contact                                            *Contact     `xml:"Contact,omitempty"`
	ContactCleanInfoReviewers                          *QueryResult `xml:"ContactCleanInfoReviewers,omitempty"`
	ContactId                                          *ID          `xml:"ContactId,omitempty"`
	ContentDocumentLinks                               *QueryResult `xml:"ContentDocumentLinks,omitempty"`
	ContractsSigned                                    *QueryResult `xml:"ContractsSigned,omitempty"`
	Country                                            string       `xml:"Country,omitempty"`
	CreatedBy                                          *User        `xml:"CreatedBy,omitempty"`
	CreatedById                                        *ID          `xml:"CreatedById,omitempty"`
	CreatedDate                                        time.Time    `xml:"CreatedDate,omitempty"`
	DeclinedEventRelations                             *QueryResult `xml:"DeclinedEventRelations,omitempty"`
	DefaultGroupNotificationFrequency                  string       `xml:"DefaultGroupNotificationFrequency,omitempty"`
	DelegatedApproverId                                *ID          `xml:"DelegatedApproverId,omitempty"`
	DelegatedUsers                                     *QueryResult `xml:"DelegatedUsers,omitempty"`
	Department                                         string       `xml:"Department,omitempty"`
	DigestFrequency                                    string       `xml:"DigestFrequency,omitempty"`
	Division                                           string       `xml:"Division,omitempty"`
	Email                                              string       `xml:"Email,omitempty"`
	EmailEncodingKey                                   string       `xml:"EmailEncodingKey,omitempty"`
	EmailMessageRelations                              *QueryResult `xml:"EmailMessageRelations,omitempty"`
	EmailPreferencesAutoBcc                            bool         `xml:"EmailPreferencesAutoBcc,omitempty"`
	EmailPreferencesAutoBccStayInTouch                 bool         `xml:"EmailPreferencesAutoBccStayInTouch,omitempty"`
	EmailPreferencesStayInTouchReminder                bool         `xml:"EmailPreferencesStayInTouchReminder,omitempty"`
	EmployeeNumber                                     string       `xml:"EmployeeNumber,omitempty"`
	EventRelations                                     *QueryResult `xml:"EventRelations,omitempty"`
	Extension                                          string       `xml:"Extension,omitempty"`
	ExternalDataUserAuths                              *QueryResult `xml:"ExternalDataUserAuths,omitempty"`
	Fax                                                string       `xml:"Fax,omitempty"`
	FederationIdentifier                               string       `xml:"FederationIdentifier,omitempty"`
	FeedSubscriptions                                  *QueryResult `xml:"FeedSubscriptions,omitempty"`
	FeedSubscriptionsForEntity                         *QueryResult `xml:"FeedSubscriptionsForEntity,omitempty"`
	Feeds                                              *QueryResult `xml:"Feeds,omitempty"`
	FirstName                                          string       `xml:"FirstName,omitempty"`
	ForecastEnabled                                    bool         `xml:"ForecastEnabled,omitempty"`
	FullPhotoUrl                                       string       `xml:"FullPhotoUrl,omitempty"`
	GeocodeAccuracy                                    string       `xml:"GeocodeAccuracy,omitempty"`
	GroupMembershipRequests                            *QueryResult `xml:"GroupMembershipRequests,omitempty"`
	GroupMemberships                                   *QueryResult `xml:"GroupMemberships,omitempty"`
	InstalledMobileApps                                *QueryResult `xml:"InstalledMobileApps,omitempty"`
	IsActive                                           bool         `xml:"IsActive,omitempty"`
	IsExtIndicatorVisible                              bool         `xml:"IsExtIndicatorVisible,omitempty"`
	IsProfilePhotoActive                               bool         `xml:"IsProfilePhotoActive,omitempty"`
	JigsawImportLimitOverride                          int32        `xml:"JigsawImportLimitOverride,omitempty"`
	LanguageLocaleKey                                  string       `xml:"LanguageLocaleKey,omitempty"`
	LastLoginDate                                      time.Time    `xml:"LastLoginDate,omitempty"`
	LastModifiedBy                                     *User        `xml:"LastModifiedBy,omitempty"`
	LastModifiedById                                   *ID          `xml:"LastModifiedById,omitempty"`
	LastModifiedDate                                   time.Time    `xml:"LastModifiedDate,omitempty"`
	LastName                                           string       `xml:"LastName,omitempty"`
	LastPasswordChangeDate                             time.Time    `xml:"LastPasswordChangeDate,omitempty"`
	LastReferencedDate                                 time.Time    `xml:"LastReferencedDate,omitempty"`
	LastViewedDate                                     time.Time    `xml:"LastViewedDate,omitempty"`
	Latitude                                           float64      `xml:"Latitude,omitempty"`
	LeadCleanInfoReviewers                             *QueryResult `xml:"LeadCleanInfoReviewers,omitempty"`
	LocaleSidKey                                       string       `xml:"LocaleSidKey,omitempty"`
	Longitude                                          float64      `xml:"Longitude,omitempty"`
	ManagedUsers                                       *QueryResult `xml:"ManagedUsers,omitempty"`
	Manager                                            *User        `xml:"Manager,omitempty"`
	ManagerId                                          *ID          `xml:"ManagerId,omitempty"`
	MediumBannerPhotoUrl                               string       `xml:"MediumBannerPhotoUrl,omitempty"`
	MediumPhotoUrl                                     string       `xml:"MediumPhotoUrl,omitempty"`
	MobilePhone                                        string       `xml:"MobilePhone,omitempty"`
	Name                                               string       `xml:"Name,omitempty"`
	OfflinePdaTrialExpirationDate                      time.Time    `xml:"OfflinePdaTrialExpirationDate,omitempty"`
	OfflineTrialExpirationDate                         time.Time    `xml:"OfflineTrialExpirationDate,omitempty"`
	OutOfOfficeMessage                                 string       `xml:"OutOfOfficeMessage,omitempty"`
	OutgoingEmailRelations                             *QueryResult `xml:"OutgoingEmailRelations,omitempty"`
	OwnedContentDocuments                              *QueryResult `xml:"OwnedContentDocuments,omitempty"`
	PermissionSetAssignments                           *QueryResult `xml:"PermissionSetAssignments,omitempty"`
	PermissionSetLicenseAssignments                    *QueryResult `xml:"PermissionSetLicenseAssignments,omitempty"`
	PersonRecord                                       *QueryResult `xml:"PersonRecord,omitempty"`
	Phone                                              string       `xml:"Phone,omitempty"`
	PostalCode                                         string       `xml:"PostalCode,omitempty"`
	Profile                                            *Profile     `xml:"Profile,omitempty"`
	ProfileId                                          *ID          `xml:"ProfileId,omitempty"`
	ReceivesAdminInfoEmails                            bool         `xml:"ReceivesAdminInfoEmails,omitempty"`
	ReceivesInfoEmails                                 bool         `xml:"ReceivesInfoEmails,omitempty"`
	SenderEmail                                        string       `xml:"SenderEmail,omitempty"`
	SenderName                                         string       `xml:"SenderName,omitempty"`
	SessionPermSetActivations                          *QueryResult `xml:"SessionPermSetActivations,omitempty"`
	Shares                                             *QueryResult `xml:"Shares,omitempty"`
	Signature                                          string       `xml:"Signature,omitempty"`
	SmallBannerPhotoUrl                                string       `xml:"SmallBannerPhotoUrl,omitempty"`
	SmallPhotoUrl                                      string       `xml:"SmallPhotoUrl,omitempty"`
	State                                              string       `xml:"State,omitempty"`
	StayInTouchNote                                    string       `xml:"StayInTouchNote,omitempty"`
	StayInTouchSignature                               string       `xml:"StayInTouchSignature,omitempty"`
	StayInTouchSubject                                 string       `xml:"StayInTouchSubject,omitempty"`
	Street                                             string       `xml:"Street,omitempty"`
	SystemModstamp                                     time.Time    `xml:"SystemModstamp,omitempty"`
	TimeZoneSidKey                                     string       `xml:"TimeZoneSidKey,omitempty"`
	Title                                              string       `xml:"Title,omitempty"`
	UndecidedEventRelations                            *QueryResult `xml:"UndecidedEventRelations,omitempty"`
	UserEntityAccessRights                             *QueryResult `xml:"UserEntityAccessRights,omitempty"`
	UserFieldAccessRights                              *QueryResult `xml:"UserFieldAccessRights,omitempty"`
	UserPermissionsCallCenterAutoLogin                 bool         `xml:"UserPermissionsCallCenterAutoLogin,omitempty"`
	UserPermissionsInteractionUser                     bool         `xml:"UserPermissionsInteractionUser,omitempty"`
	UserPermissionsJigsawProspectingUser               bool         `xml:"UserPermissionsJigsawProspectingUser,omitempty"`
	UserPermissionsKnowledgeUser                       bool         `xml:"UserPermissionsKnowledgeUser,omitempty"`
	UserPermissionsMarketingUser                       bool         `xml:"UserPermissionsMarketingUser,omitempty"`
	UserPermissionsMobileUser                          bool         `xml:"UserPermissionsMobileUser,omitempty"`
	UserPermissionsOfflineUser                         bool         `xml:"UserPermissionsOfflineUser,omitempty"`
	UserPermissionsSFContentUser                       bool         `xml:"UserPermissionsSFContentUser,omitempty"`
	UserPermissionsSiteforceContributorUser            bool         `xml:"UserPermissionsSiteforceContributorUser,omitempty"`
	UserPermissionsSiteforcePublisherUser              bool         `xml:"UserPermissionsSiteforcePublisherUser,omitempty"`
	UserPermissionsSupportUser                         bool         `xml:"UserPermissionsSupportUser,omitempty"`
	UserPermissionsWorkDotComUserFeature               bool         `xml:"UserPermissionsWorkDotComUserFeature,omitempty"`
	UserPreferences                                    *QueryResult `xml:"UserPreferences,omitempty"`
	UserPreferencesActivityRemindersPopup              bool         `xml:"UserPreferencesActivityRemindersPopup,omitempty"`
	UserPreferencesApexPagesDeveloperMode              bool         `xml:"UserPreferencesApexPagesDeveloperMode,omitempty"`
	UserPreferencesCacheDiagnostics                    bool         `xml:"UserPreferencesCacheDiagnostics,omitempty"`
	UserPreferencesContentEmailAsAndWhen               bool         `xml:"UserPreferencesContentEmailAsAndWhen,omitempty"`
	UserPreferencesContentNoEmail                      bool         `xml:"UserPreferencesContentNoEmail,omitempty"`
	UserPreferencesCreateLEXAppsWTShown                bool         `xml:"UserPreferencesCreateLEXAppsWTShown,omitempty"`
	UserPreferencesDisCommentAfterLikeEmail            bool         `xml:"UserPreferencesDisCommentAfterLikeEmail,omitempty"`
	UserPreferencesDisMentionsCommentEmail             bool         `xml:"UserPreferencesDisMentionsCommentEmail,omitempty"`
	UserPreferencesDisProfPostCommentEmail             bool         `xml:"UserPreferencesDisProfPostCommentEmail,omitempty"`
	UserPreferencesDisableAllFeedsEmail                bool         `xml:"UserPreferencesDisableAllFeedsEmail,omitempty"`
	UserPreferencesDisableBookmarkEmail                bool         `xml:"UserPreferencesDisableBookmarkEmail,omitempty"`
	UserPreferencesDisableChangeCommentEmail           bool         `xml:"UserPreferencesDisableChangeCommentEmail,omitempty"`
	UserPreferencesDisableEndorsementEmail             bool         `xml:"UserPreferencesDisableEndorsementEmail,omitempty"`
	UserPreferencesDisableFeedbackEmail                bool         `xml:"UserPreferencesDisableFeedbackEmail,omitempty"`
	UserPreferencesDisableFileShareNotificationsForApi bool         `xml:"UserPreferencesDisableFileShareNotificationsForApi,omitempty"`
	UserPreferencesDisableFollowersEmail               bool         `xml:"UserPreferencesDisableFollowersEmail,omitempty"`
	UserPreferencesDisableLaterCommentEmail            bool         `xml:"UserPreferencesDisableLaterCommentEmail,omitempty"`
	UserPreferencesDisableLikeEmail                    bool         `xml:"UserPreferencesDisableLikeEmail,omitempty"`
	UserPreferencesDisableMentionsPostEmail            bool         `xml:"UserPreferencesDisableMentionsPostEmail,omitempty"`
	UserPreferencesDisableMessageEmail                 bool         `xml:"UserPreferencesDisableMessageEmail,omitempty"`
	UserPreferencesDisableProfilePostEmail             bool         `xml:"UserPreferencesDisableProfilePostEmail,omitempty"`
	UserPreferencesDisableSharePostEmail               bool         `xml:"UserPreferencesDisableSharePostEmail,omitempty"`
	UserPreferencesDisableWorkEmail                    bool         `xml:"UserPreferencesDisableWorkEmail,omitempty"`
	UserPreferencesEnableAutoSubForFeeds               bool         `xml:"UserPreferencesEnableAutoSubForFeeds,omitempty"`
	UserPreferencesEventRemindersCheckboxDefault       bool         `xml:"UserPreferencesEventRemindersCheckboxDefault,omitempty"`
	UserPreferencesExcludeMailAppAttachments           bool         `xml:"UserPreferencesExcludeMailAppAttachments,omitempty"`
	UserPreferencesFavoritesShowTopFavorites           bool         `xml:"UserPreferencesFavoritesShowTopFavorites,omitempty"`
	UserPreferencesFavoritesWTShown                    bool         `xml:"UserPreferencesFavoritesWTShown,omitempty"`
	UserPreferencesGlobalNavBarWTShown                 bool         `xml:"UserPreferencesGlobalNavBarWTShown,omitempty"`
	UserPreferencesGlobalNavGridMenuWTShown            bool         `xml:"UserPreferencesGlobalNavGridMenuWTShown,omitempty"`
	UserPreferencesHasCelebrationBadge                 bool         `xml:"UserPreferencesHasCelebrationBadge,omitempty"`
	UserPreferencesHideBiggerPhotoCallout              bool         `xml:"UserPreferencesHideBiggerPhotoCallout,omitempty"`
	UserPreferencesHideCSNDesktopTask                  bool         `xml:"UserPreferencesHideCSNDesktopTask,omitempty"`
	UserPreferencesHideCSNGetChatterMobileTask         bool         `xml:"UserPreferencesHideCSNGetChatterMobileTask,omitempty"`
	UserPreferencesHideChatterOnboardingSplash         bool         `xml:"UserPreferencesHideChatterOnboardingSplash,omitempty"`
	UserPreferencesHideEndUserOnboardingAssistantModal bool         `xml:"UserPreferencesHideEndUserOnboardingAssistantModal,omitempty"`
	UserPreferencesHideLightningMigrationModal         bool         `xml:"UserPreferencesHideLightningMigrationModal,omitempty"`
	UserPreferencesHideS1BrowserUI                     bool         `xml:"UserPreferencesHideS1BrowserUI,omitempty"`
	UserPreferencesHideSecondChatterOnboardingSplash   bool         `xml:"UserPreferencesHideSecondChatterOnboardingSplash,omitempty"`
	UserPreferencesHideSfxWelcomeMat                   bool         `xml:"UserPreferencesHideSfxWelcomeMat,omitempty"`
	UserPreferencesJigsawListUser                      bool         `xml:"UserPreferencesJigsawListUser,omitempty"`
	UserPreferencesLightningExperiencePreferred        bool         `xml:"UserPreferencesLightningExperiencePreferred,omitempty"`
	UserPreferencesNewLightningReportRunPageEnabled    bool         `xml:"UserPreferencesNewLightningReportRunPageEnabled,omitempty"`
	UserPreferencesPathAssistantCollapsed              bool         `xml:"UserPreferencesPathAssistantCollapsed,omitempty"`
	UserPreferencesPipelineViewHideHelpPopover         bool         `xml:"UserPreferencesPipelineViewHideHelpPopover,omitempty"`
	UserPreferencesPreviewCustomTheme                  bool         `xml:"UserPreferencesPreviewCustomTheme,omitempty"`
	UserPreferencesPreviewLightning                    bool         `xml:"UserPreferencesPreviewLightning,omitempty"`
	UserPreferencesRecordHomeReservedWTShown           bool         `xml:"UserPreferencesRecordHomeReservedWTShown,omitempty"`
	UserPreferencesRecordHomeSectionCollapseWTShown    bool         `xml:"UserPreferencesRecordHomeSectionCollapseWTShown,omitempty"`
	UserPreferencesReminderSoundOff                    bool         `xml:"UserPreferencesReminderSoundOff,omitempty"`
	UserPreferencesShowCityToExternalUsers             bool         `xml:"UserPreferencesShowCityToExternalUsers,omitempty"`
	UserPreferencesShowCityToGuestUsers                bool         `xml:"UserPreferencesShowCityToGuestUsers,omitempty"`
	UserPreferencesShowCountryToExternalUsers          bool         `xml:"UserPreferencesShowCountryToExternalUsers,omitempty"`
	UserPreferencesShowCountryToGuestUsers             bool         `xml:"UserPreferencesShowCountryToGuestUsers,omitempty"`
	UserPreferencesShowEmailToExternalUsers            bool         `xml:"UserPreferencesShowEmailToExternalUsers,omitempty"`
	UserPreferencesShowEmailToGuestUsers               bool         `xml:"UserPreferencesShowEmailToGuestUsers,omitempty"`
	UserPreferencesShowFaxToExternalUsers              bool         `xml:"UserPreferencesShowFaxToExternalUsers,omitempty"`
	UserPreferencesShowFaxToGuestUsers                 bool         `xml:"UserPreferencesShowFaxToGuestUsers,omitempty"`
	UserPreferencesShowManagerToExternalUsers          bool         `xml:"UserPreferencesShowManagerToExternalUsers,omitempty"`
	UserPreferencesShowManagerToGuestUsers             bool         `xml:"UserPreferencesShowManagerToGuestUsers,omitempty"`
	UserPreferencesShowMobilePhoneToExternalUsers      bool         `xml:"UserPreferencesShowMobilePhoneToExternalUsers,omitempty"`
	UserPreferencesShowMobilePhoneToGuestUsers         bool         `xml:"UserPreferencesShowMobilePhoneToGuestUsers,omitempty"`
	UserPreferencesShowPostalCodeToExternalUsers       bool         `xml:"UserPreferencesShowPostalCodeToExternalUsers,omitempty"`
	UserPreferencesShowPostalCodeToGuestUsers          bool         `xml:"UserPreferencesShowPostalCodeToGuestUsers,omitempty"`
	UserPreferencesShowProfilePicToGuestUsers          bool         `xml:"UserPreferencesShowProfilePicToGuestUsers,omitempty"`
	UserPreferencesShowStateToExternalUsers            bool         `xml:"UserPreferencesShowStateToExternalUsers,omitempty"`
	UserPreferencesShowStateToGuestUsers               bool         `xml:"UserPreferencesShowStateToGuestUsers,omitempty"`
	UserPreferencesShowStreetAddressToExternalUsers    bool         `xml:"UserPreferencesShowStreetAddressToExternalUsers,omitempty"`
	UserPreferencesShowStreetAddressToGuestUsers       bool         `xml:"UserPreferencesShowStreetAddressToGuestUsers,omitempty"`
	UserPreferencesShowTitleToExternalUsers            bool         `xml:"UserPreferencesShowTitleToExternalUsers,omitempty"`
	UserPreferencesShowTitleToGuestUsers               bool         `xml:"UserPreferencesShowTitleToGuestUsers,omitempty"`
	UserPreferencesShowWorkPhoneToExternalUsers        bool         `xml:"UserPreferencesShowWorkPhoneToExternalUsers,omitempty"`
	UserPreferencesShowWorkPhoneToGuestUsers           bool         `xml:"UserPreferencesShowWorkPhoneToGuestUsers,omitempty"`
	UserPreferencesSortFeedByComment                   bool         `xml:"UserPreferencesSortFeedByComment,omitempty"`
	UserPreferencesSuppressEventSFXReminders           bool         `xml:"UserPreferencesSuppressEventSFXReminders,omitempty"`
	UserPreferencesSuppressTaskSFXReminders            bool         `xml:"UserPreferencesSuppressTaskSFXReminders,omitempty"`
	UserPreferencesTaskRemindersCheckboxDefault        bool         `xml:"UserPreferencesTaskRemindersCheckboxDefault,omitempty"`
	UserPreferencesUserDebugModePref                   bool         `xml:"UserPreferencesUserDebugModePref,omitempty"`
	UserRole                                           *UserRole    `xml:"UserRole,omitempty"`
	UserRoleId                                         *ID          `xml:"UserRoleId,omitempty"`
	UserSites                                          *QueryResult `xml:"UserSites,omitempty"`
	UserType                                           string       `xml:"UserType,omitempty"`
	Username                                           string       `xml:"Username,omitempty"`
}

type UserAppInfo struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com UserAppInfo"`

	*SObject

	AppDefinition    *AppDefinition    `xml:"AppDefinition,omitempty"`
	AppDefinitionId  *ID               `xml:"AppDefinitionId,omitempty"`
	CreatedBy        *User             `xml:"CreatedBy,omitempty"`
	CreatedById      *ID               `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time         `xml:"CreatedDate,omitempty"`
	FormFactor       string            `xml:"FormFactor,omitempty"`
	IsDeleted        bool              `xml:"IsDeleted,omitempty"`
	LastModifiedBy   *User             `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID               `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time         `xml:"LastModifiedDate,omitempty"`
	SystemModstamp   time.Time         `xml:"SystemModstamp,omitempty"`
	User             *User             `xml:"User,omitempty"`
	UserId           *ID               `xml:"UserId,omitempty"`
	UserRecordAccess *UserRecordAccess `xml:"UserRecordAccess,omitempty"`
}

type UserAppMenuCustomization struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com UserAppMenuCustomization"`

	*SObject

	Application      *SObject          `xml:"Application,omitempty"`
	ApplicationId    *ID               `xml:"ApplicationId,omitempty"`
	CreatedBy        *User             `xml:"CreatedBy,omitempty"`
	CreatedById      *ID               `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time         `xml:"CreatedDate,omitempty"`
	IsDeleted        bool              `xml:"IsDeleted,omitempty"`
	LastModifiedBy   *User             `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID               `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time         `xml:"LastModifiedDate,omitempty"`
	Owner            *SObject          `xml:"Owner,omitempty"`
	OwnerId          *ID               `xml:"OwnerId,omitempty"`
	SortOrder        int32             `xml:"SortOrder,omitempty"`
	SystemModstamp   time.Time         `xml:"SystemModstamp,omitempty"`
	UserRecordAccess *UserRecordAccess `xml:"UserRecordAccess,omitempty"`
}

type UserAppMenuCustomizationShare struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com UserAppMenuCustomizationShare"`

	*SObject

	AccessLevel      string                    `xml:"AccessLevel,omitempty"`
	IsDeleted        bool                      `xml:"IsDeleted,omitempty"`
	LastModifiedBy   *User                     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID                       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time                 `xml:"LastModifiedDate,omitempty"`
	Parent           *UserAppMenuCustomization `xml:"Parent,omitempty"`
	ParentId         *ID                       `xml:"ParentId,omitempty"`
	RowCause         string                    `xml:"RowCause,omitempty"`
	UserOrGroup      *SObject                  `xml:"UserOrGroup,omitempty"`
	UserOrGroupId    *ID                       `xml:"UserOrGroupId,omitempty"`
}

type UserAppMenuItem struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com UserAppMenuItem"`

	*SObject

	AppMenuItemId             string `xml:"AppMenuItemId,omitempty"`
	ApplicationId             *ID    `xml:"ApplicationId,omitempty"`
	Description               string `xml:"Description,omitempty"`
	IconUrl                   string `xml:"IconUrl,omitempty"`
	InfoUrl                   string `xml:"InfoUrl,omitempty"`
	IsUsingAdminAuthorization bool   `xml:"IsUsingAdminAuthorization,omitempty"`
	IsVisible                 bool   `xml:"IsVisible,omitempty"`
	Label                     string `xml:"Label,omitempty"`
	LogoUrl                   string `xml:"LogoUrl,omitempty"`
	MobileStartUrl            string `xml:"MobileStartUrl,omitempty"`
	Name                      string `xml:"Name,omitempty"`
	SortOrder                 int32  `xml:"SortOrder,omitempty"`
	StartUrl                  string `xml:"StartUrl,omitempty"`
	Type                      string `xml:"Type,omitempty"`
	UserSortOrder             int32  `xml:"UserSortOrder,omitempty"`
}

type UserEmailPreferredPerson struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com UserEmailPreferredPerson"`

	*SObject

	CreatedBy        *User             `xml:"CreatedBy,omitempty"`
	CreatedById      *ID               `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time         `xml:"CreatedDate,omitempty"`
	Email            string            `xml:"Email,omitempty"`
	IsDeleted        bool              `xml:"IsDeleted,omitempty"`
	LastModifiedBy   *User             `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID               `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time         `xml:"LastModifiedDate,omitempty"`
	Name             string            `xml:"Name,omitempty"`
	Owner            *SObject          `xml:"Owner,omitempty"`
	OwnerId          *ID               `xml:"OwnerId,omitempty"`
	PersonRecord     *SObject          `xml:"PersonRecord,omitempty"`
	PersonRecordId   *ID               `xml:"PersonRecordId,omitempty"`
	SystemModstamp   time.Time         `xml:"SystemModstamp,omitempty"`
	UserRecordAccess *UserRecordAccess `xml:"UserRecordAccess,omitempty"`
}

type UserEmailPreferredPersonShare struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com UserEmailPreferredPersonShare"`

	*SObject

	AccessLevel      string                    `xml:"AccessLevel,omitempty"`
	IsDeleted        bool                      `xml:"IsDeleted,omitempty"`
	LastModifiedBy   *User                     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID                       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time                 `xml:"LastModifiedDate,omitempty"`
	Parent           *UserEmailPreferredPerson `xml:"Parent,omitempty"`
	ParentId         *ID                       `xml:"ParentId,omitempty"`
	RowCause         string                    `xml:"RowCause,omitempty"`
	UserOrGroup      *SObject                  `xml:"UserOrGroup,omitempty"`
	UserOrGroupId    *ID                       `xml:"UserOrGroupId,omitempty"`
}

type UserEntityAccess struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com UserEntityAccess"`

	*SObject

	DurableId          string            `xml:"DurableId,omitempty"`
	EntityDefinition   *EntityDefinition `xml:"EntityDefinition,omitempty"`
	EntityDefinitionId string            `xml:"EntityDefinitionId,omitempty"`
	IsActivateable     bool              `xml:"IsActivateable,omitempty"`
	IsCreatable        bool              `xml:"IsCreatable,omitempty"`
	IsDeletable        bool              `xml:"IsDeletable,omitempty"`
	IsEditable         bool              `xml:"IsEditable,omitempty"`
	IsFlsUpdatable     bool              `xml:"IsFlsUpdatable,omitempty"`
	IsMergeable        bool              `xml:"IsMergeable,omitempty"`
	IsReadable         bool              `xml:"IsReadable,omitempty"`
	IsUndeletable      bool              `xml:"IsUndeletable,omitempty"`
	IsUpdatable        bool              `xml:"IsUpdatable,omitempty"`
	User               *User             `xml:"User,omitempty"`
	UserId             *ID               `xml:"UserId,omitempty"`
}

type UserFeed struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com UserFeed"`

	*SObject

	BestComment        *FeedComment `xml:"BestComment,omitempty"`
	BestCommentId      *ID          `xml:"BestCommentId,omitempty"`
	Body               string       `xml:"Body,omitempty"`
	CommentCount       int32        `xml:"CommentCount,omitempty"`
	CreatedBy          *SObject     `xml:"CreatedBy,omitempty"`
	CreatedById        *ID          `xml:"CreatedById,omitempty"`
	CreatedDate        time.Time    `xml:"CreatedDate,omitempty"`
	FeedAttachments    *QueryResult `xml:"FeedAttachments,omitempty"`
	FeedComments       *QueryResult `xml:"FeedComments,omitempty"`
	FeedLikes          *QueryResult `xml:"FeedLikes,omitempty"`
	FeedSignals        *QueryResult `xml:"FeedSignals,omitempty"`
	FeedTrackedChanges *QueryResult `xml:"FeedTrackedChanges,omitempty"`
	InsertedBy         *SObject     `xml:"InsertedBy,omitempty"`
	InsertedById       *ID          `xml:"InsertedById,omitempty"`
	IsDeleted          bool         `xml:"IsDeleted,omitempty"`
	IsRichText         bool         `xml:"IsRichText,omitempty"`
	LastModifiedDate   time.Time    `xml:"LastModifiedDate,omitempty"`
	LikeCount          int32        `xml:"LikeCount,omitempty"`
	LinkUrl            string       `xml:"LinkUrl,omitempty"`
	Parent             *User        `xml:"Parent,omitempty"`
	ParentId           *ID          `xml:"ParentId,omitempty"`
	RelatedRecordId    *ID          `xml:"RelatedRecordId,omitempty"`
	SystemModstamp     time.Time    `xml:"SystemModstamp,omitempty"`
	Title              string       `xml:"Title,omitempty"`
	Type               string       `xml:"Type,omitempty"`
}

type UserFieldAccess struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com UserFieldAccess"`

	*SObject

	DurableId          string            `xml:"DurableId,omitempty"`
	EntityDefinition   *EntityDefinition `xml:"EntityDefinition,omitempty"`
	EntityDefinitionId string            `xml:"EntityDefinitionId,omitempty"`
	FieldDefinition    *FieldDefinition  `xml:"FieldDefinition,omitempty"`
	FieldDefinitionId  string            `xml:"FieldDefinitionId,omitempty"`
	IsAccessible       bool              `xml:"IsAccessible,omitempty"`
	IsCreatable        bool              `xml:"IsCreatable,omitempty"`
	IsUpdatable        bool              `xml:"IsUpdatable,omitempty"`
	User               *User             `xml:"User,omitempty"`
	UserId             *ID               `xml:"UserId,omitempty"`
}

type UserLicense struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com UserLicense"`

	*SObject

	CreatedDate             time.Time    `xml:"CreatedDate,omitempty"`
	LastModifiedDate        time.Time    `xml:"LastModifiedDate,omitempty"`
	LicenseDefinitionKey    string       `xml:"LicenseDefinitionKey,omitempty"`
	MasterLabel             string       `xml:"MasterLabel,omitempty"`
	Name                    string       `xml:"Name,omitempty"`
	Status                  string       `xml:"Status,omitempty"`
	SystemModstamp          time.Time    `xml:"SystemModstamp,omitempty"`
	TotalLicenses           int32        `xml:"TotalLicenses,omitempty"`
	UsedLicenses            int32        `xml:"UsedLicenses,omitempty"`
	UsedLicensesLastUpdated time.Time    `xml:"UsedLicensesLastUpdated,omitempty"`
	Userlicenses            *QueryResult `xml:"userlicenses,omitempty"`
}

type UserListView struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com UserListView"`

	*SObject

	CreatedBy        *User     `xml:"CreatedBy,omitempty"`
	CreatedById      *ID       `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time `xml:"CreatedDate,omitempty"`
	IsDeleted        bool      `xml:"IsDeleted,omitempty"`
	LastModifiedBy   *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time `xml:"LastModifiedDate,omitempty"`
	LastViewedChart  string    `xml:"LastViewedChart,omitempty"`
	ListView         *ListView `xml:"ListView,omitempty"`
	ListViewId       *ID       `xml:"ListViewId,omitempty"`
	SobjectType      string    `xml:"SobjectType,omitempty"`
	SystemModstamp   time.Time `xml:"SystemModstamp,omitempty"`
	User             *User     `xml:"User,omitempty"`
	UserId           *ID       `xml:"UserId,omitempty"`
}

type UserListViewCriterion struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com UserListViewCriterion"`

	*SObject

	ColumnName       string        `xml:"ColumnName,omitempty"`
	CreatedBy        *User         `xml:"CreatedBy,omitempty"`
	CreatedById      *ID           `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time     `xml:"CreatedDate,omitempty"`
	IsDeleted        bool          `xml:"IsDeleted,omitempty"`
	LastModifiedBy   *User         `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID           `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time     `xml:"LastModifiedDate,omitempty"`
	Operation        string        `xml:"Operation,omitempty"`
	SortOrder        int32         `xml:"SortOrder,omitempty"`
	SystemModstamp   time.Time     `xml:"SystemModstamp,omitempty"`
	UserListView     *UserListView `xml:"UserListView,omitempty"`
	UserListViewId   *ID           `xml:"UserListViewId,omitempty"`
	Value            string        `xml:"Value,omitempty"`
}

type UserLogin struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com UserLogin"`

	*SObject

	IsFrozen         bool      `xml:"IsFrozen,omitempty"`
	IsPasswordLocked bool      `xml:"IsPasswordLocked,omitempty"`
	LastModifiedBy   *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time `xml:"LastModifiedDate,omitempty"`
	UserId           *ID       `xml:"UserId,omitempty"`
}

type UserPackageLicense struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com UserPackageLicense"`

	*SObject

	CreatedBy        *User           `xml:"CreatedBy,omitempty"`
	CreatedById      *ID             `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time       `xml:"CreatedDate,omitempty"`
	LastModifiedBy   *User           `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID             `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time       `xml:"LastModifiedDate,omitempty"`
	PackageLicense   *PackageLicense `xml:"PackageLicense,omitempty"`
	PackageLicenseId *ID             `xml:"PackageLicenseId,omitempty"`
	SystemModstamp   time.Time       `xml:"SystemModstamp,omitempty"`
	UserId           *ID             `xml:"UserId,omitempty"`
}

type UserPermissionAccess struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com UserPermissionAccess"`

	*SObject

	LastCacheUpdate                           time.Time `xml:"LastCacheUpdate,omitempty"`
	PermissionsAccessCMC                      bool      `xml:"PermissionsAccessCMC,omitempty"`
	PermissionsActivateContract               bool      `xml:"PermissionsActivateContract,omitempty"`
	PermissionsActivateOrder                  bool      `xml:"PermissionsActivateOrder,omitempty"`
	PermissionsAddAnalyticsRemoteConnections  bool      `xml:"PermissionsAddAnalyticsRemoteConnections,omitempty"`
	PermissionsAddDirectMessageMembers        bool      `xml:"PermissionsAddDirectMessageMembers,omitempty"`
	PermissionsAllowEmailIC                   bool      `xml:"PermissionsAllowEmailIC,omitempty"`
	PermissionsAllowLightningLogin            bool      `xml:"PermissionsAllowLightningLogin,omitempty"`
	PermissionsAllowUniversalSearch           bool      `xml:"PermissionsAllowUniversalSearch,omitempty"`
	PermissionsAllowViewEditConvertedLeads    bool      `xml:"PermissionsAllowViewEditConvertedLeads,omitempty"`
	PermissionsAllowViewKnowledge             bool      `xml:"PermissionsAllowViewKnowledge,omitempty"`
	PermissionsApexRestServices               bool      `xml:"PermissionsApexRestServices,omitempty"`
	PermissionsApiEnabled                     bool      `xml:"PermissionsApiEnabled,omitempty"`
	PermissionsAssignPermissionSets           bool      `xml:"PermissionsAssignPermissionSets,omitempty"`
	PermissionsAssignTopics                   bool      `xml:"PermissionsAssignTopics,omitempty"`
	PermissionsAuthorApex                     bool      `xml:"PermissionsAuthorApex,omitempty"`
	PermissionsAutomaticActivityCapture       bool      `xml:"PermissionsAutomaticActivityCapture,omitempty"`
	PermissionsBulkApiHardDelete              bool      `xml:"PermissionsBulkApiHardDelete,omitempty"`
	PermissionsBulkMacrosAllowed              bool      `xml:"PermissionsBulkMacrosAllowed,omitempty"`
	PermissionsCampaignInfluence2             bool      `xml:"PermissionsCampaignInfluence2,omitempty"`
	PermissionsCanApproveFeedPost             bool      `xml:"PermissionsCanApproveFeedPost,omitempty"`
	PermissionsCanEditDataPrepRecipe          bool      `xml:"PermissionsCanEditDataPrepRecipe,omitempty"`
	PermissionsCanInsertFeedSystemFields      bool      `xml:"PermissionsCanInsertFeedSystemFields,omitempty"`
	PermissionsCanManageMaps                  bool      `xml:"PermissionsCanManageMaps,omitempty"`
	PermissionsCanUseNewDashboardBuilder      bool      `xml:"PermissionsCanUseNewDashboardBuilder,omitempty"`
	PermissionsCanVerifyComment               bool      `xml:"PermissionsCanVerifyComment,omitempty"`
	PermissionsChangeDashboardColors          bool      `xml:"PermissionsChangeDashboardColors,omitempty"`
	PermissionsChatterComposeUiCodesnippet    bool      `xml:"PermissionsChatterComposeUiCodesnippet,omitempty"`
	PermissionsChatterEditOwnPost             bool      `xml:"PermissionsChatterEditOwnPost,omitempty"`
	PermissionsChatterEditOwnRecordPost       bool      `xml:"PermissionsChatterEditOwnRecordPost,omitempty"`
	PermissionsChatterFileLink                bool      `xml:"PermissionsChatterFileLink,omitempty"`
	PermissionsChatterInternalUser            bool      `xml:"PermissionsChatterInternalUser,omitempty"`
	PermissionsChatterInviteExternalUsers     bool      `xml:"PermissionsChatterInviteExternalUsers,omitempty"`
	PermissionsChatterOwnGroups               bool      `xml:"PermissionsChatterOwnGroups,omitempty"`
	PermissionsCloseConversations             bool      `xml:"PermissionsCloseConversations,omitempty"`
	PermissionsConfigCustomRecs               bool      `xml:"PermissionsConfigCustomRecs,omitempty"`
	PermissionsConnectOrgToEnvironmentHub     bool      `xml:"PermissionsConnectOrgToEnvironmentHub,omitempty"`
	PermissionsContentAdministrator           bool      `xml:"PermissionsContentAdministrator,omitempty"`
	PermissionsContentWorkspaces              bool      `xml:"PermissionsContentWorkspaces,omitempty"`
	PermissionsConvertLeads                   bool      `xml:"PermissionsConvertLeads,omitempty"`
	PermissionsCreateCustomizeDashboards      bool      `xml:"PermissionsCreateCustomizeDashboards,omitempty"`
	PermissionsCreateCustomizeFilters         bool      `xml:"PermissionsCreateCustomizeFilters,omitempty"`
	PermissionsCreateCustomizeReports         bool      `xml:"PermissionsCreateCustomizeReports,omitempty"`
	PermissionsCreateDashboardFolders         bool      `xml:"PermissionsCreateDashboardFolders,omitempty"`
	PermissionsCreateLtngTempInPub            bool      `xml:"PermissionsCreateLtngTempInPub,omitempty"`
	PermissionsCreatePackaging                bool      `xml:"PermissionsCreatePackaging,omitempty"`
	PermissionsCreateReportFolders            bool      `xml:"PermissionsCreateReportFolders,omitempty"`
	PermissionsCreateReportInLightning        bool      `xml:"PermissionsCreateReportInLightning,omitempty"`
	PermissionsCreateTopics                   bool      `xml:"PermissionsCreateTopics,omitempty"`
	PermissionsCreateWorkspaces               bool      `xml:"PermissionsCreateWorkspaces,omitempty"`
	PermissionsCustomMobileAppsAccess         bool      `xml:"PermissionsCustomMobileAppsAccess,omitempty"`
	PermissionsCustomSidebarOnAllPages        bool      `xml:"PermissionsCustomSidebarOnAllPages,omitempty"`
	PermissionsCustomizeApplication           bool      `xml:"PermissionsCustomizeApplication,omitempty"`
	PermissionsDataExport                     bool      `xml:"PermissionsDataExport,omitempty"`
	PermissionsDelegatedTwoFactor             bool      `xml:"PermissionsDelegatedTwoFactor,omitempty"`
	PermissionsDeleteActivatedContract        bool      `xml:"PermissionsDeleteActivatedContract,omitempty"`
	PermissionsDeleteTopics                   bool      `xml:"PermissionsDeleteTopics,omitempty"`
	PermissionsDistributeFromPersWksp         bool      `xml:"PermissionsDistributeFromPersWksp,omitempty"`
	PermissionsEditActivatedOrders            bool      `xml:"PermissionsEditActivatedOrders,omitempty"`
	PermissionsEditBrandTemplates             bool      `xml:"PermissionsEditBrandTemplates,omitempty"`
	PermissionsEditCaseComments               bool      `xml:"PermissionsEditCaseComments,omitempty"`
	PermissionsEditEvent                      bool      `xml:"PermissionsEditEvent,omitempty"`
	PermissionsEditHtmlTemplates              bool      `xml:"PermissionsEditHtmlTemplates,omitempty"`
	PermissionsEditKnowledge                  bool      `xml:"PermissionsEditKnowledge,omitempty"`
	PermissionsEditMyDashboards               bool      `xml:"PermissionsEditMyDashboards,omitempty"`
	PermissionsEditMyReports                  bool      `xml:"PermissionsEditMyReports,omitempty"`
	PermissionsEditOppLineItemUnitPrice       bool      `xml:"PermissionsEditOppLineItemUnitPrice,omitempty"`
	PermissionsEditPublicDocuments            bool      `xml:"PermissionsEditPublicDocuments,omitempty"`
	PermissionsEditPublicFilters              bool      `xml:"PermissionsEditPublicFilters,omitempty"`
	PermissionsEditPublicTemplates            bool      `xml:"PermissionsEditPublicTemplates,omitempty"`
	PermissionsEditReadonlyFields             bool      `xml:"PermissionsEditReadonlyFields,omitempty"`
	PermissionsEditTask                       bool      `xml:"PermissionsEditTask,omitempty"`
	PermissionsEditTopics                     bool      `xml:"PermissionsEditTopics,omitempty"`
	PermissionsEmailAdministration            bool      `xml:"PermissionsEmailAdministration,omitempty"`
	PermissionsEmailMass                      bool      `xml:"PermissionsEmailMass,omitempty"`
	PermissionsEmailSingle                    bool      `xml:"PermissionsEmailSingle,omitempty"`
	PermissionsEmailTemplateManagement        bool      `xml:"PermissionsEmailTemplateManagement,omitempty"`
	PermissionsEnableCommunityAppLauncher     bool      `xml:"PermissionsEnableCommunityAppLauncher,omitempty"`
	PermissionsEnableNotifications            bool      `xml:"PermissionsEnableNotifications,omitempty"`
	PermissionsExportReport                   bool      `xml:"PermissionsExportReport,omitempty"`
	PermissionsFeedPinning                    bool      `xml:"PermissionsFeedPinning,omitempty"`
	PermissionsFlowUFLRequired                bool      `xml:"PermissionsFlowUFLRequired,omitempty"`
	PermissionsForceTwoFactor                 bool      `xml:"PermissionsForceTwoFactor,omitempty"`
	PermissionsGiveRecognitionBadge           bool      `xml:"PermissionsGiveRecognitionBadge,omitempty"`
	PermissionsGovernNetworks                 bool      `xml:"PermissionsGovernNetworks,omitempty"`
	PermissionsHideReadByList                 bool      `xml:"PermissionsHideReadByList,omitempty"`
	PermissionsIdentityConnect                bool      `xml:"PermissionsIdentityConnect,omitempty"`
	PermissionsIdentityEnabled                bool      `xml:"PermissionsIdentityEnabled,omitempty"`
	PermissionsImportCustomObjects            bool      `xml:"PermissionsImportCustomObjects,omitempty"`
	PermissionsImportLeads                    bool      `xml:"PermissionsImportLeads,omitempty"`
	PermissionsImportPersonal                 bool      `xml:"PermissionsImportPersonal,omitempty"`
	PermissionsInsightsAppAdmin               bool      `xml:"PermissionsInsightsAppAdmin,omitempty"`
	PermissionsInsightsAppDashboardEditor     bool      `xml:"PermissionsInsightsAppDashboardEditor,omitempty"`
	PermissionsInsightsAppEltEditor           bool      `xml:"PermissionsInsightsAppEltEditor,omitempty"`
	PermissionsInsightsAppUploadUser          bool      `xml:"PermissionsInsightsAppUploadUser,omitempty"`
	PermissionsInsightsAppUser                bool      `xml:"PermissionsInsightsAppUser,omitempty"`
	PermissionsInsightsCreateApplication      bool      `xml:"PermissionsInsightsCreateApplication,omitempty"`
	PermissionsInstallPackaging               bool      `xml:"PermissionsInstallPackaging,omitempty"`
	PermissionsIotUser                        bool      `xml:"PermissionsIotUser,omitempty"`
	PermissionsLightningConsoleAllowedForUser bool      `xml:"PermissionsLightningConsoleAllowedForUser,omitempty"`
	PermissionsLightningExperienceUser        bool      `xml:"PermissionsLightningExperienceUser,omitempty"`
	PermissionsListEmailSend                  bool      `xml:"PermissionsListEmailSend,omitempty"`
	PermissionsLtngPromoReserved01UserPerm    bool      `xml:"PermissionsLtngPromoReserved01UserPerm,omitempty"`
	PermissionsManageAnalyticSnapshots        bool      `xml:"PermissionsManageAnalyticSnapshots,omitempty"`
	PermissionsManageAuthProviders            bool      `xml:"PermissionsManageAuthProviders,omitempty"`
	PermissionsManageBusinessHourHolidays     bool      `xml:"PermissionsManageBusinessHourHolidays,omitempty"`
	PermissionsManageCallCenters              bool      `xml:"PermissionsManageCallCenters,omitempty"`
	PermissionsManageCases                    bool      `xml:"PermissionsManageCases,omitempty"`
	PermissionsManageCategories               bool      `xml:"PermissionsManageCategories,omitempty"`
	PermissionsManageCertificates             bool      `xml:"PermissionsManageCertificates,omitempty"`
	PermissionsManageChatterMessages          bool      `xml:"PermissionsManageChatterMessages,omitempty"`
	PermissionsManageContentPermissions       bool      `xml:"PermissionsManageContentPermissions,omitempty"`
	PermissionsManageContentProperties        bool      `xml:"PermissionsManageContentProperties,omitempty"`
	PermissionsManageContentTypes             bool      `xml:"PermissionsManageContentTypes,omitempty"`
	PermissionsManageCustomPermissions        bool      `xml:"PermissionsManageCustomPermissions,omitempty"`
	PermissionsManageCustomReportTypes        bool      `xml:"PermissionsManageCustomReportTypes,omitempty"`
	PermissionsManageDashbdsInPubFolders      bool      `xml:"PermissionsManageDashbdsInPubFolders,omitempty"`
	PermissionsManageDataCategories           bool      `xml:"PermissionsManageDataCategories,omitempty"`
	PermissionsManageDataIntegrations         bool      `xml:"PermissionsManageDataIntegrations,omitempty"`
	PermissionsManageDynamicDashboards        bool      `xml:"PermissionsManageDynamicDashboards,omitempty"`
	PermissionsManageEmailClientConfig        bool      `xml:"PermissionsManageEmailClientConfig,omitempty"`
	PermissionsManageEncryptionKeys           bool      `xml:"PermissionsManageEncryptionKeys,omitempty"`
	PermissionsManageExchangeConfig           bool      `xml:"PermissionsManageExchangeConfig,omitempty"`
	PermissionsManageHealthCheck              bool      `xml:"PermissionsManageHealthCheck,omitempty"`
	PermissionsManageInteraction              bool      `xml:"PermissionsManageInteraction,omitempty"`
	PermissionsManageInternalUsers            bool      `xml:"PermissionsManageInternalUsers,omitempty"`
	PermissionsManageIpAddresses              bool      `xml:"PermissionsManageIpAddresses,omitempty"`
	PermissionsManageKnowledge                bool      `xml:"PermissionsManageKnowledge,omitempty"`
	PermissionsManageKnowledgeImportExport    bool      `xml:"PermissionsManageKnowledgeImportExport,omitempty"`
	PermissionsManageLeads                    bool      `xml:"PermissionsManageLeads,omitempty"`
	PermissionsManageLoginAccessPolicies      bool      `xml:"PermissionsManageLoginAccessPolicies,omitempty"`
	PermissionsManageMobile                   bool      `xml:"PermissionsManageMobile,omitempty"`
	PermissionsManageNetworks                 bool      `xml:"PermissionsManageNetworks,omitempty"`
	PermissionsManagePasswordPolicies         bool      `xml:"PermissionsManagePasswordPolicies,omitempty"`
	PermissionsManageProfilesPermissionsets   bool      `xml:"PermissionsManageProfilesPermissionsets,omitempty"`
	PermissionsManagePvtRptsAndDashbds        bool      `xml:"PermissionsManagePvtRptsAndDashbds,omitempty"`
	PermissionsManageRemoteAccess             bool      `xml:"PermissionsManageRemoteAccess,omitempty"`
	PermissionsManageReportsInPubFolders      bool      `xml:"PermissionsManageReportsInPubFolders,omitempty"`
	PermissionsManageRoles                    bool      `xml:"PermissionsManageRoles,omitempty"`
	PermissionsManageSearchPromotionRules     bool      `xml:"PermissionsManageSearchPromotionRules,omitempty"`
	PermissionsManageSessionPermissionSets    bool      `xml:"PermissionsManageSessionPermissionSets,omitempty"`
	PermissionsManageSharing                  bool      `xml:"PermissionsManageSharing,omitempty"`
	PermissionsManageSolutions                bool      `xml:"PermissionsManageSolutions,omitempty"`
	PermissionsManageSurveys                  bool      `xml:"PermissionsManageSurveys,omitempty"`
	PermissionsManageSynonyms                 bool      `xml:"PermissionsManageSynonyms,omitempty"`
	PermissionsManageTemplatedApp             bool      `xml:"PermissionsManageTemplatedApp,omitempty"`
	PermissionsManageTwoFactor                bool      `xml:"PermissionsManageTwoFactor,omitempty"`
	PermissionsManageUnlistedGroups           bool      `xml:"PermissionsManageUnlistedGroups,omitempty"`
	PermissionsManageUsers                    bool      `xml:"PermissionsManageUsers,omitempty"`
	PermissionsMassInlineEdit                 bool      `xml:"PermissionsMassInlineEdit,omitempty"`
	PermissionsMergeTopics                    bool      `xml:"PermissionsMergeTopics,omitempty"`
	PermissionsModerateChatter                bool      `xml:"PermissionsModerateChatter,omitempty"`
	PermissionsModerateNetworkUsers           bool      `xml:"PermissionsModerateNetworkUsers,omitempty"`
	PermissionsModifyAllData                  bool      `xml:"PermissionsModifyAllData,omitempty"`
	PermissionsModifyMetadata                 bool      `xml:"PermissionsModifyMetadata,omitempty"`
	PermissionsModifySecureAgents             bool      `xml:"PermissionsModifySecureAgents,omitempty"`
	PermissionsNewReportBuilder               bool      `xml:"PermissionsNewReportBuilder,omitempty"`
	PermissionsPackaging2                     bool      `xml:"PermissionsPackaging2,omitempty"`
	PermissionsPasswordNeverExpires           bool      `xml:"PermissionsPasswordNeverExpires,omitempty"`
	PermissionsPreventClassicExperience       bool      `xml:"PermissionsPreventClassicExperience,omitempty"`
	PermissionsPrivacyDataAccess              bool      `xml:"PermissionsPrivacyDataAccess,omitempty"`
	PermissionsPublishPackaging               bool      `xml:"PermissionsPublishPackaging,omitempty"`
	PermissionsRemoveDirectMessageMembers     bool      `xml:"PermissionsRemoveDirectMessageMembers,omitempty"`
	PermissionsResetPasswords                 bool      `xml:"PermissionsResetPasswords,omitempty"`
	PermissionsRunFlow                        bool      `xml:"PermissionsRunFlow,omitempty"`
	PermissionsRunReports                     bool      `xml:"PermissionsRunReports,omitempty"`
	PermissionsSalesConsole                   bool      `xml:"PermissionsSalesConsole,omitempty"`
	PermissionsScheduleReports                bool      `xml:"PermissionsScheduleReports,omitempty"`
	PermissionsSelectFilesFromSalesforce      bool      `xml:"PermissionsSelectFilesFromSalesforce,omitempty"`
	PermissionsSendAnnouncementEmails         bool      `xml:"PermissionsSendAnnouncementEmails,omitempty"`
	PermissionsSendSitRequests                bool      `xml:"PermissionsSendSitRequests,omitempty"`
	PermissionsShareInternalArticles          bool      `xml:"PermissionsShareInternalArticles,omitempty"`
	PermissionsShowCompanyNameAsUserBadge     bool      `xml:"PermissionsShowCompanyNameAsUserBadge,omitempty"`
	PermissionsSolutionImport                 bool      `xml:"PermissionsSolutionImport,omitempty"`
	PermissionsStdAutomaticActivityCapture    bool      `xml:"PermissionsStdAutomaticActivityCapture,omitempty"`
	PermissionsSubmitMacrosAllowed            bool      `xml:"PermissionsSubmitMacrosAllowed,omitempty"`
	PermissionsSubscribeDashboardToOtherUsers bool      `xml:"PermissionsSubscribeDashboardToOtherUsers,omitempty"`
	PermissionsSubscribeReportToOtherUsers    bool      `xml:"PermissionsSubscribeReportToOtherUsers,omitempty"`
	PermissionsSubscribeReportsRunAsUser      bool      `xml:"PermissionsSubscribeReportsRunAsUser,omitempty"`
	PermissionsSubscribeToLightningDashboards bool      `xml:"PermissionsSubscribeToLightningDashboards,omitempty"`
	PermissionsSubscribeToLightningReports    bool      `xml:"PermissionsSubscribeToLightningReports,omitempty"`
	PermissionsTransferAnyCase                bool      `xml:"PermissionsTransferAnyCase,omitempty"`
	PermissionsTransferAnyEntity              bool      `xml:"PermissionsTransferAnyEntity,omitempty"`
	PermissionsTransferAnyLead                bool      `xml:"PermissionsTransferAnyLead,omitempty"`
	PermissionsTwoFactorApi                   bool      `xml:"PermissionsTwoFactorApi,omitempty"`
	PermissionsUseTeamReassignWizards         bool      `xml:"PermissionsUseTeamReassignWizards,omitempty"`
	PermissionsUseTemplatedApp                bool      `xml:"PermissionsUseTemplatedApp,omitempty"`
	PermissionsUseWebLink                     bool      `xml:"PermissionsUseWebLink,omitempty"`
	PermissionsViewAllActivities              bool      `xml:"PermissionsViewAllActivities,omitempty"`
	PermissionsViewAllData                    bool      `xml:"PermissionsViewAllData,omitempty"`
	PermissionsViewAllUsers                   bool      `xml:"PermissionsViewAllUsers,omitempty"`
	PermissionsViewContent                    bool      `xml:"PermissionsViewContent,omitempty"`
	PermissionsViewDataAssessment             bool      `xml:"PermissionsViewDataAssessment,omitempty"`
	PermissionsViewDataCategories             bool      `xml:"PermissionsViewDataCategories,omitempty"`
	PermissionsViewEncryptedData              bool      `xml:"PermissionsViewEncryptedData,omitempty"`
	PermissionsViewEventLogFiles              bool      `xml:"PermissionsViewEventLogFiles,omitempty"`
	PermissionsViewHealthCheck                bool      `xml:"PermissionsViewHealthCheck,omitempty"`
	PermissionsViewHelpLink                   bool      `xml:"PermissionsViewHelpLink,omitempty"`
	PermissionsViewMyTeamsDashboards          bool      `xml:"PermissionsViewMyTeamsDashboards,omitempty"`
	PermissionsViewOnlyEmbeddedAppUser        bool      `xml:"PermissionsViewOnlyEmbeddedAppUser,omitempty"`
	PermissionsViewPublicDashboards           bool      `xml:"PermissionsViewPublicDashboards,omitempty"`
	PermissionsViewPublicReports              bool      `xml:"PermissionsViewPublicReports,omitempty"`
	PermissionsViewRoles                      bool      `xml:"PermissionsViewRoles,omitempty"`
	PermissionsViewSetup                      bool      `xml:"PermissionsViewSetup,omitempty"`
	PermissionsWaveTabularDownload            bool      `xml:"PermissionsWaveTabularDownload,omitempty"`
	PermissionsWorkCalibrationUser            bool      `xml:"PermissionsWorkCalibrationUser,omitempty"`
	PermissionsWorkDotComUserPerm             bool      `xml:"PermissionsWorkDotComUserPerm,omitempty"`
}

type UserPreference struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com UserPreference"`

	*SObject

	Preference     string    `xml:"Preference,omitempty"`
	SystemModstamp time.Time `xml:"SystemModstamp,omitempty"`
	UserId         *ID       `xml:"UserId,omitempty"`
	Value          string    `xml:"Value,omitempty"`
}

type UserProvAccount struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com UserProvAccount"`

	*SObject

	ConnectedApp      *ConnectedApplication `xml:"ConnectedApp,omitempty"`
	ConnectedAppId    *ID                   `xml:"ConnectedAppId,omitempty"`
	CreatedBy         *User                 `xml:"CreatedBy,omitempty"`
	CreatedById       *ID                   `xml:"CreatedById,omitempty"`
	CreatedDate       time.Time             `xml:"CreatedDate,omitempty"`
	DeletedDate       time.Time             `xml:"DeletedDate,omitempty"`
	ExternalEmail     string                `xml:"ExternalEmail,omitempty"`
	ExternalFirstName string                `xml:"ExternalFirstName,omitempty"`
	ExternalLastName  string                `xml:"ExternalLastName,omitempty"`
	ExternalUserId    string                `xml:"ExternalUserId,omitempty"`
	ExternalUsername  string                `xml:"ExternalUsername,omitempty"`
	IsDeleted         bool                  `xml:"IsDeleted,omitempty"`
	IsKnownLink       bool                  `xml:"IsKnownLink,omitempty"`
	LastModifiedBy    *User                 `xml:"LastModifiedBy,omitempty"`
	LastModifiedById  *ID                   `xml:"LastModifiedById,omitempty"`
	LastModifiedDate  time.Time             `xml:"LastModifiedDate,omitempty"`
	LinkState         string                `xml:"LinkState,omitempty"`
	Name              string                `xml:"Name,omitempty"`
	SalesforceUser    *User                 `xml:"SalesforceUser,omitempty"`
	SalesforceUserId  *ID                   `xml:"SalesforceUserId,omitempty"`
	Status            string                `xml:"Status,omitempty"`
	SystemModstamp    time.Time             `xml:"SystemModstamp,omitempty"`
	UserRecordAccess  *UserRecordAccess     `xml:"UserRecordAccess,omitempty"`
}

type UserProvAccountStaging struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com UserProvAccountStaging"`

	*SObject

	ConnectedApp      *ConnectedApplication `xml:"ConnectedApp,omitempty"`
	ConnectedAppId    *ID                   `xml:"ConnectedAppId,omitempty"`
	CreatedBy         *User                 `xml:"CreatedBy,omitempty"`
	CreatedById       *ID                   `xml:"CreatedById,omitempty"`
	CreatedDate       time.Time             `xml:"CreatedDate,omitempty"`
	ExternalEmail     string                `xml:"ExternalEmail,omitempty"`
	ExternalFirstName string                `xml:"ExternalFirstName,omitempty"`
	ExternalLastName  string                `xml:"ExternalLastName,omitempty"`
	ExternalUserId    string                `xml:"ExternalUserId,omitempty"`
	ExternalUsername  string                `xml:"ExternalUsername,omitempty"`
	IsDeleted         bool                  `xml:"IsDeleted,omitempty"`
	LastModifiedBy    *User                 `xml:"LastModifiedBy,omitempty"`
	LastModifiedById  *ID                   `xml:"LastModifiedById,omitempty"`
	LastModifiedDate  time.Time             `xml:"LastModifiedDate,omitempty"`
	LinkState         string                `xml:"LinkState,omitempty"`
	Name              string                `xml:"Name,omitempty"`
	SalesforceUser    *User                 `xml:"SalesforceUser,omitempty"`
	SalesforceUserId  *ID                   `xml:"SalesforceUserId,omitempty"`
	Status            string                `xml:"Status,omitempty"`
	SystemModstamp    time.Time             `xml:"SystemModstamp,omitempty"`
	UserRecordAccess  *UserRecordAccess     `xml:"UserRecordAccess,omitempty"`
}

type UserProvMockTarget struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com UserProvMockTarget"`

	*SObject

	CreatedBy         *User             `xml:"CreatedBy,omitempty"`
	CreatedById       *ID               `xml:"CreatedById,omitempty"`
	CreatedDate       time.Time         `xml:"CreatedDate,omitempty"`
	ExternalEmail     string            `xml:"ExternalEmail,omitempty"`
	ExternalFirstName string            `xml:"ExternalFirstName,omitempty"`
	ExternalLastName  string            `xml:"ExternalLastName,omitempty"`
	ExternalUserId    string            `xml:"ExternalUserId,omitempty"`
	ExternalUsername  string            `xml:"ExternalUsername,omitempty"`
	IsDeleted         bool              `xml:"IsDeleted,omitempty"`
	LastModifiedBy    *User             `xml:"LastModifiedBy,omitempty"`
	LastModifiedById  *ID               `xml:"LastModifiedById,omitempty"`
	LastModifiedDate  time.Time         `xml:"LastModifiedDate,omitempty"`
	Name              string            `xml:"Name,omitempty"`
	SystemModstamp    time.Time         `xml:"SystemModstamp,omitempty"`
	UserRecordAccess  *UserRecordAccess `xml:"UserRecordAccess,omitempty"`
}

type UserProvisioningConfig struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com UserProvisioningConfig"`

	*SObject

	ApprovalRequired   string                `xml:"ApprovalRequired,omitempty"`
	ConnectedApp       *ConnectedApplication `xml:"ConnectedApp,omitempty"`
	ConnectedAppId     *ID                   `xml:"ConnectedAppId,omitempty"`
	CreatedBy          *User                 `xml:"CreatedBy,omitempty"`
	CreatedById        *ID                   `xml:"CreatedById,omitempty"`
	CreatedDate        time.Time             `xml:"CreatedDate,omitempty"`
	DeveloperName      string                `xml:"DeveloperName,omitempty"`
	Enabled            bool                  `xml:"Enabled,omitempty"`
	EnabledOperations  string                `xml:"EnabledOperations,omitempty"`
	IsDeleted          bool                  `xml:"IsDeleted,omitempty"`
	Language           string                `xml:"Language,omitempty"`
	LastModifiedBy     *User                 `xml:"LastModifiedBy,omitempty"`
	LastModifiedById   *ID                   `xml:"LastModifiedById,omitempty"`
	LastModifiedDate   time.Time             `xml:"LastModifiedDate,omitempty"`
	LastReconDateTime  time.Time             `xml:"LastReconDateTime,omitempty"`
	MasterLabel        string                `xml:"MasterLabel,omitempty"`
	NamedCredential    *NamedCredential      `xml:"NamedCredential,omitempty"`
	NamedCredentialId  *ID                   `xml:"NamedCredentialId,omitempty"`
	NamespacePrefix    string                `xml:"NamespacePrefix,omitempty"`
	Notes              string                `xml:"Notes,omitempty"`
	OnUpdateAttributes string                `xml:"OnUpdateAttributes,omitempty"`
	ReconFilter        string                `xml:"ReconFilter,omitempty"`
	SystemModstamp     time.Time             `xml:"SystemModstamp,omitempty"`
	UserAccountMapping string                `xml:"UserAccountMapping,omitempty"`
}

type UserProvisioningLog struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com UserProvisioningLog"`

	*SObject

	CreatedBy                 *User                    `xml:"CreatedBy,omitempty"`
	CreatedById               *ID                      `xml:"CreatedById,omitempty"`
	CreatedDate               time.Time                `xml:"CreatedDate,omitempty"`
	Details                   string                   `xml:"Details,omitempty"`
	ExternalUserId            string                   `xml:"ExternalUserId,omitempty"`
	ExternalUsername          string                   `xml:"ExternalUsername,omitempty"`
	IsDeleted                 bool                     `xml:"IsDeleted,omitempty"`
	LastModifiedBy            *User                    `xml:"LastModifiedBy,omitempty"`
	LastModifiedById          *ID                      `xml:"LastModifiedById,omitempty"`
	LastModifiedDate          time.Time                `xml:"LastModifiedDate,omitempty"`
	Name                      string                   `xml:"Name,omitempty"`
	Status                    string                   `xml:"Status,omitempty"`
	SystemModstamp            time.Time                `xml:"SystemModstamp,omitempty"`
	User                      *User                    `xml:"User,omitempty"`
	UserId                    *ID                      `xml:"UserId,omitempty"`
	UserProvisioningRequest   *UserProvisioningRequest `xml:"UserProvisioningRequest,omitempty"`
	UserProvisioningRequestId *ID                      `xml:"UserProvisioningRequestId,omitempty"`
	UserRecordAccess          *UserRecordAccess        `xml:"UserRecordAccess,omitempty"`
}

type UserProvisioningRequest struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com UserProvisioningRequest"`

	*SObject

	AppName           string                   `xml:"AppName,omitempty"`
	ApprovalStatus    string                   `xml:"ApprovalStatus,omitempty"`
	ConnectedApp      *ConnectedApplication    `xml:"ConnectedApp,omitempty"`
	ConnectedAppId    *ID                      `xml:"ConnectedAppId,omitempty"`
	CreatedBy         *User                    `xml:"CreatedBy,omitempty"`
	CreatedById       *ID                      `xml:"CreatedById,omitempty"`
	CreatedDate       time.Time                `xml:"CreatedDate,omitempty"`
	ExternalUserId    string                   `xml:"ExternalUserId,omitempty"`
	IsDeleted         bool                     `xml:"IsDeleted,omitempty"`
	LastModifiedBy    *User                    `xml:"LastModifiedBy,omitempty"`
	LastModifiedById  *ID                      `xml:"LastModifiedById,omitempty"`
	LastModifiedDate  time.Time                `xml:"LastModifiedDate,omitempty"`
	Manager           *User                    `xml:"Manager,omitempty"`
	ManagerId         *ID                      `xml:"ManagerId,omitempty"`
	Name              string                   `xml:"Name,omitempty"`
	Operation         string                   `xml:"Operation,omitempty"`
	Owner             *SObject                 `xml:"Owner,omitempty"`
	OwnerId           *ID                      `xml:"OwnerId,omitempty"`
	Parent            *UserProvisioningRequest `xml:"Parent,omitempty"`
	ParentId          *ID                      `xml:"ParentId,omitempty"`
	ProcessInstances  *QueryResult             `xml:"ProcessInstances,omitempty"`
	ProcessSteps      *QueryResult             `xml:"ProcessSteps,omitempty"`
	RetryCount        int32                    `xml:"RetryCount,omitempty"`
	SalesforceUser    *User                    `xml:"SalesforceUser,omitempty"`
	SalesforceUserId  *ID                      `xml:"SalesforceUserId,omitempty"`
	ScheduleDate      time.Time                `xml:"ScheduleDate,omitempty"`
	State             string                   `xml:"State,omitempty"`
	SystemModstamp    time.Time                `xml:"SystemModstamp,omitempty"`
	UserProvAccount   *UserProvAccount         `xml:"UserProvAccount,omitempty"`
	UserProvAccountId *ID                      `xml:"UserProvAccountId,omitempty"`
	UserProvConfig    *UserProvisioningConfig  `xml:"UserProvConfig,omitempty"`
	UserProvConfigId  *ID                      `xml:"UserProvConfigId,omitempty"`
	UserRecordAccess  *UserRecordAccess        `xml:"UserRecordAccess,omitempty"`
}

type UserProvisioningRequestShare struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com UserProvisioningRequestShare"`

	*SObject

	AccessLevel      string                   `xml:"AccessLevel,omitempty"`
	IsDeleted        bool                     `xml:"IsDeleted,omitempty"`
	LastModifiedBy   *User                    `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID                      `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time                `xml:"LastModifiedDate,omitempty"`
	Parent           *UserProvisioningRequest `xml:"Parent,omitempty"`
	ParentId         *ID                      `xml:"ParentId,omitempty"`
	RowCause         string                   `xml:"RowCause,omitempty"`
	UserOrGroup      *SObject                 `xml:"UserOrGroup,omitempty"`
	UserOrGroupId    *ID                      `xml:"UserOrGroupId,omitempty"`
}

type UserRecordAccess struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com UserRecordAccess"`

	*SObject

	HasAllAccess      bool   `xml:"HasAllAccess,omitempty"`
	HasDeleteAccess   bool   `xml:"HasDeleteAccess,omitempty"`
	HasEditAccess     bool   `xml:"HasEditAccess,omitempty"`
	HasReadAccess     bool   `xml:"HasReadAccess,omitempty"`
	HasTransferAccess bool   `xml:"HasTransferAccess,omitempty"`
	MaxAccessLevel    string `xml:"MaxAccessLevel,omitempty"`
	RecordId          string `xml:"RecordId,omitempty"`
	UserId            *ID    `xml:"UserId,omitempty"`
}

type UserRole struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com UserRole"`

	*SObject

	CaseAccessForAccountOwner        string       `xml:"CaseAccessForAccountOwner,omitempty"`
	ContactAccessForAccountOwner     string       `xml:"ContactAccessForAccountOwner,omitempty"`
	DeveloperName                    string       `xml:"DeveloperName,omitempty"`
	ForecastUserId                   *ID          `xml:"ForecastUserId,omitempty"`
	LastModifiedBy                   *User        `xml:"LastModifiedBy,omitempty"`
	LastModifiedById                 *ID          `xml:"LastModifiedById,omitempty"`
	LastModifiedDate                 time.Time    `xml:"LastModifiedDate,omitempty"`
	MayForecastManagerShare          bool         `xml:"MayForecastManagerShare,omitempty"`
	Name                             string       `xml:"Name,omitempty"`
	OpportunityAccessForAccountOwner string       `xml:"OpportunityAccessForAccountOwner,omitempty"`
	ParentRoleId                     *ID          `xml:"ParentRoleId,omitempty"`
	PortalAccountId                  *ID          `xml:"PortalAccountId,omitempty"`
	PortalAccountOwnerId             *ID          `xml:"PortalAccountOwnerId,omitempty"`
	PortalType                       string       `xml:"PortalType,omitempty"`
	RollupDescription                string       `xml:"RollupDescription,omitempty"`
	SystemModstamp                   time.Time    `xml:"SystemModstamp,omitempty"`
	Users                            *QueryResult `xml:"Users,omitempty"`
}

type UserShare struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com UserShare"`

	*SObject

	IsActive         bool      `xml:"IsActive,omitempty"`
	LastModifiedBy   *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time `xml:"LastModifiedDate,omitempty"`
	RowCause         string    `xml:"RowCause,omitempty"`
	User             *User     `xml:"User,omitempty"`
	UserAccessLevel  string    `xml:"UserAccessLevel,omitempty"`
	UserId           *ID       `xml:"UserId,omitempty"`
	UserOrGroup      *SObject  `xml:"UserOrGroup,omitempty"`
	UserOrGroupId    *ID       `xml:"UserOrGroupId,omitempty"`
}

type VerificationHistory struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com VerificationHistory"`

	*SObject

	Activity           string                `xml:"Activity,omitempty"`
	CreatedBy          *User                 `xml:"CreatedBy,omitempty"`
	CreatedById        *ID                   `xml:"CreatedById,omitempty"`
	CreatedDate        time.Time             `xml:"CreatedDate,omitempty"`
	EventGroup         int32                 `xml:"EventGroup,omitempty"`
	IsDeleted          bool                  `xml:"IsDeleted,omitempty"`
	LastModifiedBy     *User                 `xml:"LastModifiedBy,omitempty"`
	LastModifiedById   *ID                   `xml:"LastModifiedById,omitempty"`
	LastModifiedDate   time.Time             `xml:"LastModifiedDate,omitempty"`
	LoginGeo           *LoginGeo             `xml:"LoginGeo,omitempty"`
	LoginGeoId         *ID                   `xml:"LoginGeoId,omitempty"`
	LoginHistory       *LoginHistory         `xml:"LoginHistory,omitempty"`
	LoginHistoryId     *ID                   `xml:"LoginHistoryId,omitempty"`
	Policy             string                `xml:"Policy,omitempty"`
	Remarks            string                `xml:"Remarks,omitempty"`
	Resource           *ConnectedApplication `xml:"Resource,omitempty"`
	ResourceId         *ID                   `xml:"ResourceId,omitempty"`
	SourceIp           string                `xml:"SourceIp,omitempty"`
	Status             string                `xml:"Status,omitempty"`
	SystemModstamp     time.Time             `xml:"SystemModstamp,omitempty"`
	User               *User                 `xml:"User,omitempty"`
	UserId             *ID                   `xml:"UserId,omitempty"`
	VerificationMethod string                `xml:"VerificationMethod,omitempty"`
	VerificationTime   time.Time             `xml:"VerificationTime,omitempty"`
}

type VisualforceAccessMetrics struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com VisualforceAccessMetrics"`

	*SObject

	ApexPage           *ApexPage `xml:"ApexPage,omitempty"`
	ApexPageId         *ID       `xml:"ApexPageId,omitempty"`
	DailyPageViewCount int32     `xml:"DailyPageViewCount,omitempty"`
	LogDate            time.Time `xml:"LogDate,omitempty"`
	MetricsDate        time.Time `xml:"MetricsDate,omitempty"`
	Profile            *Profile  `xml:"Profile,omitempty"`
	ProfileId          *ID       `xml:"ProfileId,omitempty"`
	SystemModstamp     time.Time `xml:"SystemModstamp,omitempty"`
}

type Vote struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com Vote"`

	*SObject

	CreatedBy        *User     `xml:"CreatedBy,omitempty"`
	CreatedById      *ID       `xml:"CreatedById,omitempty"`
	CreatedDate      time.Time `xml:"CreatedDate,omitempty"`
	IsDeleted        bool      `xml:"IsDeleted,omitempty"`
	LastModifiedBy   *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate time.Time `xml:"LastModifiedDate,omitempty"`
	Parent           *SObject  `xml:"Parent,omitempty"`
	ParentId         *ID       `xml:"ParentId,omitempty"`
	SystemModstamp   time.Time `xml:"SystemModstamp,omitempty"`
	Type             string    `xml:"Type,omitempty"`
}

type WaveAutoInstallRequest struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com WaveAutoInstallRequest"`

	*SObject

	Configuration               string            `xml:"Configuration,omitempty"`
	CreatedBy                   *User             `xml:"CreatedBy,omitempty"`
	CreatedById                 *ID               `xml:"CreatedById,omitempty"`
	CreatedDate                 time.Time         `xml:"CreatedDate,omitempty"`
	FailedReason                string            `xml:"FailedReason,omitempty"`
	Folder                      *Folder           `xml:"Folder,omitempty"`
	FolderId                    *ID               `xml:"FolderId,omitempty"`
	IsDeleted                   bool              `xml:"IsDeleted,omitempty"`
	LastModifiedBy              *User             `xml:"LastModifiedBy,omitempty"`
	LastModifiedById            *ID               `xml:"LastModifiedById,omitempty"`
	LastModifiedDate            time.Time         `xml:"LastModifiedDate,omitempty"`
	Name                        string            `xml:"Name,omitempty"`
	RequestLog                  string            `xml:"RequestLog,omitempty"`
	RequestStatus               string            `xml:"RequestStatus,omitempty"`
	RequestType                 string            `xml:"RequestType,omitempty"`
	SystemModstamp              time.Time         `xml:"SystemModstamp,omitempty"`
	TemplateApiName             string            `xml:"TemplateApiName,omitempty"`
	TemplateVersion             string            `xml:"TemplateVersion,omitempty"`
	UserRecordAccess            *UserRecordAccess `xml:"UserRecordAccess,omitempty"`
	WaveCompatibilityCheckItems *QueryResult      `xml:"WaveCompatibilityCheckItems,omitempty"`
}

type WaveCompatibilityCheckItem struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com WaveCompatibilityCheckItem"`

	*SObject

	CreatedBy                *User                   `xml:"CreatedBy,omitempty"`
	CreatedById              *ID                     `xml:"CreatedById,omitempty"`
	CreatedDate              time.Time               `xml:"CreatedDate,omitempty"`
	IsDeleted                bool                    `xml:"IsDeleted,omitempty"`
	LastModifiedBy           *User                   `xml:"LastModifiedBy,omitempty"`
	LastModifiedById         *ID                     `xml:"LastModifiedById,omitempty"`
	LastModifiedDate         time.Time               `xml:"LastModifiedDate,omitempty"`
	Name                     string                  `xml:"Name,omitempty"`
	Payload                  string                  `xml:"Payload,omitempty"`
	SystemModstamp           time.Time               `xml:"SystemModstamp,omitempty"`
	TaskName                 string                  `xml:"TaskName,omitempty"`
	TaskResult               string                  `xml:"TaskResult,omitempty"`
	TemplateApiName          string                  `xml:"TemplateApiName,omitempty"`
	TemplateVersion          string                  `xml:"TemplateVersion,omitempty"`
	UserRecordAccess         *UserRecordAccess       `xml:"UserRecordAccess,omitempty"`
	WaveAutoInstallRequest   *WaveAutoInstallRequest `xml:"WaveAutoInstallRequest,omitempty"`
	WaveAutoInstallRequestId *ID                     `xml:"WaveAutoInstallRequestId,omitempty"`
}

type WebLink struct {
	XMLName xml.Name `xml:"urn:sobject.enterprise.soap.sforce.com WebLink"`

	*SObject

	CreatedBy           *User     `xml:"CreatedBy,omitempty"`
	CreatedById         *ID       `xml:"CreatedById,omitempty"`
	CreatedDate         time.Time `xml:"CreatedDate,omitempty"`
	Description         string    `xml:"Description,omitempty"`
	DisplayType         string    `xml:"DisplayType,omitempty"`
	EncodingKey         string    `xml:"EncodingKey,omitempty"`
	HasMenubar          bool      `xml:"HasMenubar,omitempty"`
	HasScrollbars       bool      `xml:"HasScrollbars,omitempty"`
	HasToolbar          bool      `xml:"HasToolbar,omitempty"`
	Height              int32     `xml:"Height,omitempty"`
	IsProtected         bool      `xml:"IsProtected,omitempty"`
	IsResizable         bool      `xml:"IsResizable,omitempty"`
	LastModifiedBy      *User     `xml:"LastModifiedBy,omitempty"`
	LastModifiedById    *ID       `xml:"LastModifiedById,omitempty"`
	LastModifiedDate    time.Time `xml:"LastModifiedDate,omitempty"`
	LinkType            string    `xml:"LinkType,omitempty"`
	MasterLabel         string    `xml:"MasterLabel,omitempty"`
	Name                string    `xml:"Name,omitempty"`
	NamespacePrefix     string    `xml:"NamespacePrefix,omitempty"`
	OpenType            string    `xml:"OpenType,omitempty"`
	PageOrSobjectType   string    `xml:"PageOrSobjectType,omitempty"`
	Position            string    `xml:"Position,omitempty"`
	RequireRowSelection bool      `xml:"RequireRowSelection,omitempty"`
	ScontrolId          *ID       `xml:"ScontrolId,omitempty"`
	ShowsLocation       bool      `xml:"ShowsLocation,omitempty"`
	ShowsStatus         bool      `xml:"ShowsStatus,omitempty"`
	SystemModstamp      time.Time `xml:"SystemModstamp,omitempty"`
	Url                 string    `xml:"Url,omitempty"`
	Width               int32     `xml:"Width,omitempty"`
}

type ID string

const ()

type QueryLocator string

const ()

type StatusCode string

const (
	StatusCodeALLORNONEOPERATIONROLLEDBACK StatusCode = "ALLORNONEOPERATIONROLLEDBACK"

	StatusCodeALREADYINPROCESS StatusCode = "ALREADYINPROCESS"

	StatusCodeAPEXDATAACCESSRESTRICTION StatusCode = "APEXDATAACCESSRESTRICTION"

	StatusCodeASSIGNEETYPEREQUIRED StatusCode = "ASSIGNEETYPEREQUIRED"

	StatusCodeAURACOMPILEERROR StatusCode = "AURACOMPILEERROR"

	StatusCodeAUTHPROVIDERNEEDSAUTH StatusCode = "AUTHPROVIDERNEEDSAUTH"

	StatusCodeAUTHPROVIDERNOTFOUND StatusCode = "AUTHPROVIDERNOTFOUND"

	StatusCodeBADCUSTOMENTITYPARENTDOMAIN StatusCode = "BADCUSTOMENTITYPARENTDOMAIN"

	StatusCodeBCCNOTALLOWEDIFBCCCOMPLIANCEENABLED StatusCode = "BCCNOTALLOWEDIFBCCCOMPLIANCEENABLED"

	StatusCodeCANNOTCASCADEPRODUCTACTIVE StatusCode = "CANNOTCASCADEPRODUCTACTIVE"

	StatusCodeCANNOTCHANGEFIELDTYPEOFAPEXREFERENCEDFIELD StatusCode = "CANNOTCHANGEFIELDTYPEOFAPEXREFERENCEDFIELD"

	StatusCodeCANNOTCHANGEFIELDTYPEOFREFERENCEDFIELD StatusCode = "CANNOTCHANGEFIELDTYPEOFREFERENCEDFIELD"

	StatusCodeCANNOTCREATEANOTHERMANAGEDPACKAGE StatusCode = "CANNOTCREATEANOTHERMANAGEDPACKAGE"

	StatusCodeCANNOTDEACTIVATEDIVISION StatusCode = "CANNOTDEACTIVATEDIVISION"

	StatusCodeCANNOTDELETEGLOBALACTIONLIST StatusCode = "CANNOTDELETEGLOBALACTIONLIST"

	StatusCodeCANNOTDELETELASTDATEDCONVERSIONRATE StatusCode = "CANNOTDELETELASTDATEDCONVERSIONRATE"

	StatusCodeCANNOTDELETEMANAGEDOBJECT StatusCode = "CANNOTDELETEMANAGEDOBJECT"

	StatusCodeCANNOTDISABLELASTADMIN StatusCode = "CANNOTDISABLELASTADMIN"

	StatusCodeCANNOTENABLEIPRESTRICTREQUESTS StatusCode = "CANNOTENABLEIPRESTRICTREQUESTS"

	StatusCodeCANNOTEXECUTEFLOWTRIGGER StatusCode = "CANNOTEXECUTEFLOWTRIGGER"

	StatusCodeCANNOTFREEZESELF StatusCode = "CANNOTFREEZESELF"

	StatusCodeCANNOTINSERTUPDATEACTIVATEENTITY StatusCode = "CANNOTINSERTUPDATEACTIVATEENTITY"

	StatusCodeCANNOTMODIFYMANAGEDOBJECT StatusCode = "CANNOTMODIFYMANAGEDOBJECT"

	StatusCodeCANNOTPASSWORDLOCKOUT StatusCode = "CANNOTPASSWORDLOCKOUT"

	StatusCodeCANNOTPOSTTOARCHIVEDGROUP StatusCode = "CANNOTPOSTTOARCHIVEDGROUP"

	StatusCodeCANNOTRENAMEAPEXREFERENCEDFIELD StatusCode = "CANNOTRENAMEAPEXREFERENCEDFIELD"

	StatusCodeCANNOTRENAMEAPEXREFERENCEDOBJECT StatusCode = "CANNOTRENAMEAPEXREFERENCEDOBJECT"

	StatusCodeCANNOTRENAMEREFERENCEDFIELD StatusCode = "CANNOTRENAMEREFERENCEDFIELD"

	StatusCodeCANNOTRENAMEREFERENCEDOBJECT StatusCode = "CANNOTRENAMEREFERENCEDOBJECT"

	StatusCodeCANNOTREPARENTRECORD StatusCode = "CANNOTREPARENTRECORD"

	StatusCodeCANNOTUPDATECONVERTEDLEAD StatusCode = "CANNOTUPDATECONVERTEDLEAD"

	StatusCodeCANTDISABLECORPCURRENCY StatusCode = "CANTDISABLECORPCURRENCY"

	StatusCodeCANTUNSETCORPCURRENCY StatusCode = "CANTUNSETCORPCURRENCY"

	StatusCodeCHILDSHAREFAILSPARENT StatusCode = "CHILDSHAREFAILSPARENT"

	StatusCodeCIRCULARDEPENDENCY StatusCode = "CIRCULARDEPENDENCY"

	StatusCodeCLEANSERVICEERROR StatusCode = "CLEANSERVICEERROR"

	StatusCodeCOLLISIONDETECTED StatusCode = "COLLISIONDETECTED"

	StatusCodeCOMMERCIALCONTROLERROR StatusCode = "COMMERCIALCONTROLERROR"

	StatusCodeCOMMUNITYNOTACCESSIBLE StatusCode = "COMMUNITYNOTACCESSIBLE"

	StatusCodeCONFLICTINGENVIRONMENTHUBMEMBER StatusCode = "CONFLICTINGENVIRONMENTHUBMEMBER"

	StatusCodeCONFLICTINGSSOUSERMAPPING StatusCode = "CONFLICTINGSSOUSERMAPPING"

	StatusCodeCUSTOMAPEXERROR StatusCode = "CUSTOMAPEXERROR"

	StatusCodeCUSTOMCLOBFIELDLIMITEXCEEDED StatusCode = "CUSTOMCLOBFIELDLIMITEXCEEDED"

	StatusCodeCUSTOMENTITYORFIELDLIMIT StatusCode = "CUSTOMENTITYORFIELDLIMIT"

	StatusCodeCUSTOMFIELDINDEXLIMITEXCEEDED StatusCode = "CUSTOMFIELDINDEXLIMITEXCEEDED"

	StatusCodeCUSTOMINDEXEXISTS StatusCode = "CUSTOMINDEXEXISTS"

	StatusCodeCUSTOMLINKLIMITEXCEEDED StatusCode = "CUSTOMLINKLIMITEXCEEDED"

	StatusCodeCUSTOMMETADATALIMITEXCEEDED StatusCode = "CUSTOMMETADATALIMITEXCEEDED"

	StatusCodeCUSTOMMETADATARELFIELDMANAGEABILITY StatusCode = "CUSTOMMETADATARELFIELDMANAGEABILITY"

	StatusCodeCUSTOMSETTINGSLIMITEXCEEDED StatusCode = "CUSTOMSETTINGSLIMITEXCEEDED"

	StatusCodeCUSTOMTABLIMITEXCEEDED StatusCode = "CUSTOMTABLIMITEXCEEDED"

	StatusCodeDATAASSESSMENTCONFIGASSESSMENTINPROGRESSERROR StatusCode = "DATAASSESSMENTCONFIGASSESSMENTINPROGRESSERROR"

	StatusCodeDATAASSESSMENTCONFIGSERVICEERROR StatusCode = "DATAASSESSMENTCONFIGSERVICEERROR"

	StatusCodeDATACLOUDADDRESSNORECORDSFOUND StatusCode = "DATACLOUDADDRESSNORECORDSFOUND"

	StatusCodeDATACLOUDADDRESSPROCESSINGERROR StatusCode = "DATACLOUDADDRESSPROCESSINGERROR"

	StatusCodeDATACLOUDADDRESSSERVERERROR StatusCode = "DATACLOUDADDRESSSERVERERROR"

	StatusCodeDELETEFAILED StatusCode = "DELETEFAILED"

	StatusCodeDELETENOTALLOWED StatusCode = "DELETENOTALLOWED"

	StatusCodeDELETEOPERATIONTOOLARGE StatusCode = "DELETEOPERATIONTOOLARGE"

	StatusCodeDELETEREQUIREDONCASCADE StatusCode = "DELETEREQUIREDONCASCADE"

	StatusCodeDEPENDENCYEXISTS StatusCode = "DEPENDENCYEXISTS"

	StatusCodeDUPLICATESDETECTED StatusCode = "DUPLICATESDETECTED"

	StatusCodeDUPLICATECASESOLUTION StatusCode = "DUPLICATECASESOLUTION"

	StatusCodeDUPLICATECOMMNICKNAME StatusCode = "DUPLICATECOMMNICKNAME"

	StatusCodeDUPLICATECUSTOMENTITYDEFINITION StatusCode = "DUPLICATECUSTOMENTITYDEFINITION"

	StatusCodeDUPLICATECUSTOMTABMOTIF StatusCode = "DUPLICATECUSTOMTABMOTIF"

	StatusCodeDUPLICATEDEVELOPERNAME StatusCode = "DUPLICATEDEVELOPERNAME"

	StatusCodeDUPLICATEEXTERNALID StatusCode = "DUPLICATEEXTERNALID"

	StatusCodeDUPLICATEMASTERLABEL StatusCode = "DUPLICATEMASTERLABEL"

	StatusCodeDUPLICATESENDERDISPLAYNAME StatusCode = "DUPLICATESENDERDISPLAYNAME"

	StatusCodeDUPLICATEUSERNAME StatusCode = "DUPLICATEUSERNAME"

	StatusCodeDUPLICATEVALUE StatusCode = "DUPLICATEVALUE"

	StatusCodeEMAILADDRESSBOUNCED StatusCode = "EMAILADDRESSBOUNCED"

	StatusCodeEMAILEXTERNALTRANSPORTCONNECTIONERROR StatusCode = "EMAILEXTERNALTRANSPORTCONNECTIONERROR"

	StatusCodeEMAILEXTERNALTRANSPORTPERMISSIONERROR StatusCode = "EMAILEXTERNALTRANSPORTPERMISSIONERROR"

	StatusCodeEMAILEXTERNALTRANSPORTTOKENERROR StatusCode = "EMAILEXTERNALTRANSPORTTOKENERROR"

	StatusCodeEMAILEXTERNALTRANSPORTTOOMANYREQUESTSERROR StatusCode = "EMAILEXTERNALTRANSPORTTOOMANYREQUESTSERROR"

	StatusCodeEMAILEXTERNALTRANSPORTUNKNOWNERROR StatusCode = "EMAILEXTERNALTRANSPORTUNKNOWNERROR"

	StatusCodeEMAILNOTPROCESSEDDUETOPRIORERROR StatusCode = "EMAILNOTPROCESSEDDUETOPRIORERROR"

	StatusCodeEMAILOPTEDOUT StatusCode = "EMAILOPTEDOUT"

	StatusCodeEMAILTEMPLATEFORMULAERROR StatusCode = "EMAILTEMPLATEFORMULAERROR"

	StatusCodeEMAILTEMPLATEMERGEFIELDACCESSERROR StatusCode = "EMAILTEMPLATEMERGEFIELDACCESSERROR"

	StatusCodeEMAILTEMPLATEMERGEFIELDERROR StatusCode = "EMAILTEMPLATEMERGEFIELDERROR"

	StatusCodeEMAILTEMPLATEMERGEFIELDVALUEERROR StatusCode = "EMAILTEMPLATEMERGEFIELDVALUEERROR"

	StatusCodeEMAILTEMPLATEPROCESSINGERROR StatusCode = "EMAILTEMPLATEPROCESSINGERROR"

	StatusCodeEMPTYSCONTROLFILENAME StatusCode = "EMPTYSCONTROLFILENAME"

	StatusCodeENHANCEDEMAILTEMPLATECOMPILATIONERROR StatusCode = "ENHANCEDEMAILTEMPLATECOMPILATIONERROR"

	StatusCodeENTITYFAILEDIFLASTMODIFIEDONUPDATE StatusCode = "ENTITYFAILEDIFLASTMODIFIEDONUPDATE"

	StatusCodeENTITYISARCHIVED StatusCode = "ENTITYISARCHIVED"

	StatusCodeENTITYISDELETED StatusCode = "ENTITYISDELETED"

	StatusCodeENTITYISLOCKED StatusCode = "ENTITYISLOCKED"

	StatusCodeENTITYSAVEERROR StatusCode = "ENTITYSAVEERROR"

	StatusCodeENTITYSAVEVALIDATIONERROR StatusCode = "ENTITYSAVEVALIDATIONERROR"

	StatusCodeENVIRONMENTHUBMEMBERSHIPCONFLICT StatusCode = "ENVIRONMENTHUBMEMBERSHIPCONFLICT"

	StatusCodeENVIRONMENTHUBMEMBERSHIPERRORJOININGHUB StatusCode = "ENVIRONMENTHUBMEMBERSHIPERRORJOININGHUB"

	StatusCodeENVIRONMENTHUBMEMBERSHIPUSERALREADYINHUB StatusCode = "ENVIRONMENTHUBMEMBERSHIPUSERALREADYINHUB"

	StatusCodeENVIRONMENTHUBMEMBERSHIPUSERNOTORGADMIN StatusCode = "ENVIRONMENTHUBMEMBERSHIPUSERNOTORGADMIN"

	StatusCodeERRORINMAILER StatusCode = "ERRORINMAILER"

	StatusCodeEXCHANGEWEBSERVICESURLINVALID StatusCode = "EXCHANGEWEBSERVICESURLINVALID"

	StatusCodeEXTERNALRESOURCEFORBIDDEN StatusCode = "EXTERNALRESOURCEFORBIDDEN"

	StatusCodeFAILEDACTIVATION StatusCode = "FAILEDACTIVATION"

	StatusCodeFIELDCUSTOMVALIDATIONEXCEPTION StatusCode = "FIELDCUSTOMVALIDATIONEXCEPTION"

	StatusCodeFIELDFILTERVALIDATIONEXCEPTION StatusCode = "FIELDFILTERVALIDATIONEXCEPTION"

	StatusCodeFIELDINTEGRITYEXCEPTION StatusCode = "FIELDINTEGRITYEXCEPTION"

	StatusCodeFIELDKEYWORDLISTMATCHLIMIT StatusCode = "FIELDKEYWORDLISTMATCHLIMIT"

	StatusCodeFIELDMAPPINGERROR StatusCode = "FIELDMAPPINGERROR"

	StatusCodeFIELDMODERATIONRULEBLOCK StatusCode = "FIELDMODERATIONRULEBLOCK"

	StatusCodeFIELDNOTUPDATABLE StatusCode = "FIELDNOTUPDATABLE"

	StatusCodeFILEEXTENSIONNOTALLOWED StatusCode = "FILEEXTENSIONNOTALLOWED"

	StatusCodeFILESIZELIMITEXCEEDED StatusCode = "FILESIZELIMITEXCEEDED"

	StatusCodeFILTEREDLOOKUPLIMITEXCEEDED StatusCode = "FILTEREDLOOKUPLIMITEXCEEDED"

	StatusCodeFINDDUPLICATESERROR StatusCode = "FINDDUPLICATESERROR"

	StatusCodeFLOWEXCEPTION StatusCode = "FLOWEXCEPTION"

	StatusCodeFUNCTIONALITYNOTENABLED StatusCode = "FUNCTIONALITYNOTENABLED"

	StatusCodeHASPUBLICREFERENCES StatusCode = "HASPUBLICREFERENCES"

	StatusCodeHTMLFILEUPLOADNOTALLOWED StatusCode = "HTMLFILEUPLOADNOTALLOWED"

	StatusCodeIMAGETOOLARGE StatusCode = "IMAGETOOLARGE"

	StatusCodeINACTIVEOWNERORUSER StatusCode = "INACTIVEOWNERORUSER"

	StatusCodeINACTIVERULEERROR StatusCode = "INACTIVERULEERROR"

	StatusCodeINSERTUPDATEDELETENOTALLOWEDDURINGMAINTENANCE StatusCode = "INSERTUPDATEDELETENOTALLOWEDDURINGMAINTENANCE"

	StatusCodeINSUFFICIENTACCESSONCROSSREFERENCEENTITY StatusCode = "INSUFFICIENTACCESSONCROSSREFERENCEENTITY"

	StatusCodeINSUFFICIENTACCESSORREADONLY StatusCode = "INSUFFICIENTACCESSORREADONLY"

	StatusCodeINSUFFICIENTACCESSTOINSIGHTSEXTERNALDATA StatusCode = "INSUFFICIENTACCESSTOINSIGHTSEXTERNALDATA"

	StatusCodeINSUFFICIENTCREDITS StatusCode = "INSUFFICIENTCREDITS"

	StatusCodeINTERNALERROR StatusCode = "INTERNALERROR"

	StatusCodeINVALIDACCESSLEVEL StatusCode = "INVALIDACCESSLEVEL"

	StatusCodeINVALIDACCESSTOKEN StatusCode = "INVALIDACCESSTOKEN"

	StatusCodeINVALIDAPIINPUT StatusCode = "INVALIDAPIINPUT"

	StatusCodeINVALIDARGUMENTTYPE StatusCode = "INVALIDARGUMENTTYPE"

	StatusCodeINVALIDASSIGNEETYPE StatusCode = "INVALIDASSIGNEETYPE"

	StatusCodeINVALIDASSIGNMENTRULE StatusCode = "INVALIDASSIGNMENTRULE"

	StatusCodeINVALIDAUTHHEADER StatusCode = "INVALIDAUTHHEADER"

	StatusCodeINVALIDBATCHOPERATION StatusCode = "INVALIDBATCHOPERATION"

	StatusCodeINVALIDCONTENTTYPE StatusCode = "INVALIDCONTENTTYPE"

	StatusCodeINVALIDCREDITCARDINFO StatusCode = "INVALIDCREDITCARDINFO"

	StatusCodeINVALIDCROSSREFERENCEKEY StatusCode = "INVALIDCROSSREFERENCEKEY"

	StatusCodeINVALIDCROSSREFERENCETYPEFORFIELD StatusCode = "INVALIDCROSSREFERENCETYPEFORFIELD"

	StatusCodeINVALIDCURRENCYCONVRATE StatusCode = "INVALIDCURRENCYCONVRATE"

	StatusCodeINVALIDCURRENCYCORPRATE StatusCode = "INVALIDCURRENCYCORPRATE"

	StatusCodeINVALIDCURRENCYISO StatusCode = "INVALIDCURRENCYISO"

	StatusCodeINVALIDDATASETREFERENCEINPUT StatusCode = "INVALIDDATASETREFERENCEINPUT"

	StatusCodeINVALIDDATACATEGORYGROUPREFERENCE StatusCode = "INVALIDDATACATEGORYGROUPREFERENCE"

	StatusCodeINVALIDDATAURI StatusCode = "INVALIDDATAURI"

	StatusCodeINVALIDEMAILADDRESS StatusCode = "INVALIDEMAILADDRESS"

	StatusCodeINVALIDEMPTYKEYOWNER StatusCode = "INVALIDEMPTYKEYOWNER"

	StatusCodeINVALIDENTITYFORMATCHENGINEERROR StatusCode = "INVALIDENTITYFORMATCHENGINEERROR"

	StatusCodeINVALIDENTITYFORMATCHOPERATIONERROR StatusCode = "INVALIDENTITYFORMATCHOPERATIONERROR"

	StatusCodeINVALIDENTITYFORUPSERT StatusCode = "INVALIDENTITYFORUPSERT"

	StatusCodeINVALIDENVIRONMENTHUBMEMBER StatusCode = "INVALIDENVIRONMENTHUBMEMBER"

	StatusCodeINVALIDEVENTDELIVERY StatusCode = "INVALIDEVENTDELIVERY"

	StatusCodeINVALIDEVENTINPUT StatusCode = "INVALIDEVENTINPUT"

	StatusCodeINVALIDEVENTSUBSCRIPTION StatusCode = "INVALIDEVENTSUBSCRIPTION"

	StatusCodeINVALIDEXTENSIONID StatusCode = "INVALIDEXTENSIONID"

	StatusCodeINVALIDFIELD StatusCode = "INVALIDFIELD"

	StatusCodeINVALIDFIELDFORINSERTUPDATE StatusCode = "INVALIDFIELDFORINSERTUPDATE"

	StatusCodeINVALIDFIELDWHENUSINGTEMPLATE StatusCode = "INVALIDFIELDWHENUSINGTEMPLATE"

	StatusCodeINVALIDFILTERACTION StatusCode = "INVALIDFILTERACTION"

	StatusCodeINVALIDGOOGLEDOCSURL StatusCode = "INVALIDGOOGLEDOCSURL"

	StatusCodeINVALIDIDFIELD StatusCode = "INVALIDIDFIELD"

	StatusCodeINVALIDINETADDRESS StatusCode = "INVALIDINETADDRESS"

	StatusCodeINVALIDINPUT StatusCode = "INVALIDINPUT"

	StatusCodeINVALIDKEYFIELDINPUT StatusCode = "INVALIDKEYFIELDINPUT"

	StatusCodeINVALIDLINEITEMCLONESTATE StatusCode = "INVALIDLINEITEMCLONESTATE"

	StatusCodeINVALIDMARKUP StatusCode = "INVALIDMARKUP"

	StatusCodeINVALIDMASTERORTRANSLATEDSOLUTION StatusCode = "INVALIDMASTERORTRANSLATEDSOLUTION"

	StatusCodeINVALIDMESSAGEIDREFERENCE StatusCode = "INVALIDMESSAGEIDREFERENCE"

	StatusCodeINVALIDNAMESPACEPREFIX StatusCode = "INVALIDNAMESPACEPREFIX"

	StatusCodeINVALIDOAUTHURL StatusCode = "INVALIDOAUTHURL"

	StatusCodeINVALIDOPERATION StatusCode = "INVALIDOPERATION"

	StatusCodeINVALIDOPERATOR StatusCode = "INVALIDOPERATOR"

	StatusCodeINVALIDORNULLFORRESTRICTEDPICKLIST StatusCode = "INVALIDORNULLFORRESTRICTEDPICKLIST"

	StatusCodeINVALIDOWNER StatusCode = "INVALIDOWNER"

	StatusCodeINVALIDPACKAGELICENSE StatusCode = "INVALIDPACKAGELICENSE"

	StatusCodeINVALIDPACKAGEVERSION StatusCode = "INVALIDPACKAGEVERSION"

	StatusCodeINVALIDPARTNERNETWORKSTATUS StatusCode = "INVALIDPARTNERNETWORKSTATUS"

	StatusCodeINVALIDPAYLOADVERSION StatusCode = "INVALIDPAYLOADVERSION"

	StatusCodeINVALIDPERSONACCOUNTOPERATION StatusCode = "INVALIDPERSONACCOUNTOPERATION"

	StatusCodeINVALIDPROVIDERTYPE StatusCode = "INVALIDPROVIDERTYPE"

	StatusCodeINVALIDQUERYLOCATOR StatusCode = "INVALIDQUERYLOCATOR"

	StatusCodeINVALIDREADONLYUSERDML StatusCode = "INVALIDREADONLYUSERDML"

	StatusCodeINVALIDREFRESHTOKEN StatusCode = "INVALIDREFRESHTOKEN"

	StatusCodeINVALIDRUNTIMEVALUE StatusCode = "INVALIDRUNTIMEVALUE"

	StatusCodeINVALIDSAVEASACTIVITYFLAG StatusCode = "INVALIDSAVEASACTIVITYFLAG"

	StatusCodeINVALIDSCSINBOUNDUSER StatusCode = "INVALIDSCSINBOUNDUSER"

	StatusCodeINVALIDSESSIONID StatusCode = "INVALIDSESSIONID"

	StatusCodeINVALIDSETUPOWNER StatusCode = "INVALIDSETUPOWNER"

	StatusCodeINVALIDSIGNUPCOUNTRY StatusCode = "INVALIDSIGNUPCOUNTRY"

	StatusCodeINVALIDSIGNUPOPTION StatusCode = "INVALIDSIGNUPOPTION"

	StatusCodeINVALIDSITEDELETEEXCEPTION StatusCode = "INVALIDSITEDELETEEXCEPTION"

	StatusCodeINVALIDSITEFILEIMPORTEDEXCEPTION StatusCode = "INVALIDSITEFILEIMPORTEDEXCEPTION"

	StatusCodeINVALIDSITEFILETYPEEXCEPTION StatusCode = "INVALIDSITEFILETYPEEXCEPTION"

	StatusCodeINVALIDSTATUS StatusCode = "INVALIDSTATUS"

	StatusCodeINVALIDSUBDOMAIN StatusCode = "INVALIDSUBDOMAIN"

	StatusCodeINVALIDTEXTREPRESENTATION StatusCode = "INVALIDTEXTREPRESENTATION"

	StatusCodeINVALIDTYPE StatusCode = "INVALIDTYPE"

	StatusCodeINVALIDTYPEFOROPERATION StatusCode = "INVALIDTYPEFOROPERATION"

	StatusCodeINVALIDTYPEONFIELDINRECORD StatusCode = "INVALIDTYPEONFIELDINRECORD"

	StatusCodeINVALIDUSERID StatusCode = "INVALIDUSERID"

	StatusCodeIPRANGELIMITEXCEEDED StatusCode = "IPRANGELIMITEXCEEDED"

	StatusCodeITEMNOTFOUND StatusCode = "ITEMNOTFOUND"

	StatusCodeJIGSAWIMPORTLIMITEXCEEDED StatusCode = "JIGSAWIMPORTLIMITEXCEEDED"

	StatusCodeLICENSELIMITEXCEEDED StatusCode = "LICENSELIMITEXCEEDED"

	StatusCodeLIGHTPORTALUSEREXCEPTION StatusCode = "LIGHTPORTALUSEREXCEPTION"

	StatusCodeLIMITEXCEEDED StatusCode = "LIMITEXCEEDED"

	StatusCodeMALFORMEDID StatusCode = "MALFORMEDID"

	StatusCodeMANAGERNOTDEFINED StatusCode = "MANAGERNOTDEFINED"

	StatusCodeMASSMAILRETRYLIMITEXCEEDED StatusCode = "MASSMAILRETRYLIMITEXCEEDED"

	StatusCodeMASSMAILLIMITEXCEEDED StatusCode = "MASSMAILLIMITEXCEEDED"

	StatusCodeMATCHDEFINITIONERROR StatusCode = "MATCHDEFINITIONERROR"

	StatusCodeMATCHOPERATIONERROR StatusCode = "MATCHOPERATIONERROR"

	StatusCodeMATCHOPERATIONINVALIDENGINEERROR StatusCode = "MATCHOPERATIONINVALIDENGINEERROR"

	StatusCodeMATCHOPERATIONINVALIDRULEERROR StatusCode = "MATCHOPERATIONINVALIDRULEERROR"

	StatusCodeMATCHOPERATIONMISSINGENGINEERROR StatusCode = "MATCHOPERATIONMISSINGENGINEERROR"

	StatusCodeMATCHOPERATIONMISSINGOBJECTTYPEERROR StatusCode = "MATCHOPERATIONMISSINGOBJECTTYPEERROR"

	StatusCodeMATCHOPERATIONMISSINGOPTIONSERROR StatusCode = "MATCHOPERATIONMISSINGOPTIONSERROR"

	StatusCodeMATCHOPERATIONMISSINGRULEERROR StatusCode = "MATCHOPERATIONMISSINGRULEERROR"

	StatusCodeMATCHOPERATIONUNKNOWNRULEERROR StatusCode = "MATCHOPERATIONUNKNOWNRULEERROR"

	StatusCodeMATCHOPERATIONUNSUPPORTEDVERSIONERROR StatusCode = "MATCHOPERATIONUNSUPPORTEDVERSIONERROR"

	StatusCodeMATCHPRECONDITIONFAILED StatusCode = "MATCHPRECONDITIONFAILED"

	StatusCodeMATCHRUNTIMEERROR StatusCode = "MATCHRUNTIMEERROR"

	StatusCodeMATCHSERVICEERROR StatusCode = "MATCHSERVICEERROR"

	StatusCodeMATCHSERVICETIMEDOUT StatusCode = "MATCHSERVICETIMEDOUT"

	StatusCodeMATCHSERVICEUNAVAILABLEERROR StatusCode = "MATCHSERVICEUNAVAILABLEERROR"

	StatusCodeMAXIMUMCCEMAILSEXCEEDED StatusCode = "MAXIMUMCCEMAILSEXCEEDED"

	StatusCodeMAXIMUMDASHBOARDCOMPONENTSEXCEEDED StatusCode = "MAXIMUMDASHBOARDCOMPONENTSEXCEEDED"

	StatusCodeMAXIMUMHIERARCHYCHILDRENREACHED StatusCode = "MAXIMUMHIERARCHYCHILDRENREACHED"

	StatusCodeMAXIMUMHIERARCHYLEVELSREACHED StatusCode = "MAXIMUMHIERARCHYLEVELSREACHED"

	StatusCodeMAXIMUMHIERARCHYTREESIZEREACHED StatusCode = "MAXIMUMHIERARCHYTREESIZEREACHED"

	StatusCodeMAXIMUMSIZEOFATTACHMENT StatusCode = "MAXIMUMSIZEOFATTACHMENT"

	StatusCodeMAXIMUMSIZEOFDOCUMENT StatusCode = "MAXIMUMSIZEOFDOCUMENT"

	StatusCodeMAXACTIONSPERRULEEXCEEDED StatusCode = "MAXACTIONSPERRULEEXCEEDED"

	StatusCodeMAXACTIVERULESEXCEEDED StatusCode = "MAXACTIVERULESEXCEEDED"

	StatusCodeMAXAPPROVALSTEPSEXCEEDED StatusCode = "MAXAPPROVALSTEPSEXCEEDED"

	StatusCodeMAXDEPTHINFLOWEXECUTION StatusCode = "MAXDEPTHINFLOWEXECUTION"

	StatusCodeMAXFORMULASPERRULEEXCEEDED StatusCode = "MAXFORMULASPERRULEEXCEEDED"

	StatusCodeMAXLIMITEXCEEDED StatusCode = "MAXLIMITEXCEEDED"

	StatusCodeMAXRULESEXCEEDED StatusCode = "MAXRULESEXCEEDED"

	StatusCodeMAXRULEENTRIESEXCEEDED StatusCode = "MAXRULEENTRIESEXCEEDED"

	StatusCodeMAXTASKDESCRIPTIONEXCEEEDED StatusCode = "MAXTASKDESCRIPTIONEXCEEEDED"

	StatusCodeMAXTMRULESEXCEEDED StatusCode = "MAXTMRULESEXCEEDED"

	StatusCodeMAXTMRULEITEMSEXCEEDED StatusCode = "MAXTMRULEITEMSEXCEEDED"

	StatusCodeMAXTRIGGERSEXCEEDED StatusCode = "MAXTRIGGERSEXCEEDED"

	StatusCodeMERGEFAILED StatusCode = "MERGEFAILED"

	StatusCodeMETADATAFIELDUPDATEERROR StatusCode = "METADATAFIELDUPDATEERROR"

	StatusCodeMISSINGARGUMENT StatusCode = "MISSINGARGUMENT"

	StatusCodeMISSINGRECORD StatusCode = "MISSINGRECORD"

	StatusCodeMIXEDDMLOPERATION StatusCode = "MIXEDDMLOPERATION"

	StatusCodeNONUNIQUESHIPPINGADDRESS StatusCode = "NONUNIQUESHIPPINGADDRESS"

	StatusCodeNOACCESSTOKEN StatusCode = "NOACCESSTOKEN"

	StatusCodeNOACCESSTOKENFROMREFRESH StatusCode = "NOACCESSTOKENFROMREFRESH"

	StatusCodeNOAPPLICABLEPROCESS StatusCode = "NOAPPLICABLEPROCESS"

	StatusCodeNOATTACHMENTPERMISSION StatusCode = "NOATTACHMENTPERMISSION"

	StatusCodeNOAUTHPROVIDER StatusCode = "NOAUTHPROVIDER"

	StatusCodeNOINACTIVEDIVISIONMEMBERS StatusCode = "NOINACTIVEDIVISIONMEMBERS"

	StatusCodeNOMASSMAILPERMISSION StatusCode = "NOMASSMAILPERMISSION"

	StatusCodeNOPARTNERPERMISSION StatusCode = "NOPARTNERPERMISSION"

	StatusCodeNOREFRESHTOKEN StatusCode = "NOREFRESHTOKEN"

	StatusCodeNOSUCHUSEREXISTS StatusCode = "NOSUCHUSEREXISTS"

	StatusCodeNOTOKENENDPOINT StatusCode = "NOTOKENENDPOINT"

	StatusCodeNUMBEROUTSIDEVALIDRANGE StatusCode = "NUMBEROUTSIDEVALIDRANGE"

	StatusCodeNUMHISTORYFIELDSBYSOBJECTEXCEEDED StatusCode = "NUMHISTORYFIELDSBYSOBJECTEXCEEDED"

	StatusCodeOPTEDOUTOFMASSMAIL StatusCode = "OPTEDOUTOFMASSMAIL"

	StatusCodeOPWITHINVALIDUSERTYPEEXCEPTION StatusCode = "OPWITHINVALIDUSERTYPEEXCEPTION"

	StatusCodeORCHESTRATIONINVALID StatusCode = "ORCHESTRATIONINVALID"

	StatusCodePACKAGELICENSEREQUIRED StatusCode = "PACKAGELICENSEREQUIRED"

	StatusCodePACKAGINGAPIINSTALLFAILED StatusCode = "PACKAGINGAPIINSTALLFAILED"

	StatusCodePACKAGINGAPIUNINSTALLFAILED StatusCode = "PACKAGINGAPIUNINSTALLFAILED"

	StatusCodePALIINVALIDACTIONID StatusCode = "PALIINVALIDACTIONID"

	StatusCodePALIINVALIDACTIONNAME StatusCode = "PALIINVALIDACTIONNAME"

	StatusCodePALIINVALIDACTIONTYPE StatusCode = "PALIINVALIDACTIONTYPE"

	StatusCodePALINVALIDASSISTANTRECOMMENDATIONTYPEID StatusCode = "PALINVALIDASSISTANTRECOMMENDATIONTYPEID"

	StatusCodePALINVALIDENTITYID StatusCode = "PALINVALIDENTITYID"

	StatusCodePALINVALIDFLEXIPAGEID StatusCode = "PALINVALIDFLEXIPAGEID"

	StatusCodePALINVALIDLAYOUTID StatusCode = "PALINVALIDLAYOUTID"

	StatusCodePALINVALIDPARAMETERS StatusCode = "PALINVALIDPARAMETERS"

	StatusCodePAYLOADSIZEEXCEEDED StatusCode = "PAYLOADSIZEEXCEEDED"

	StatusCodePAAPIEXCEPTION StatusCode = "PAAPIEXCEPTION"

	StatusCodePAAXISFAULT StatusCode = "PAAXISFAULT"

	StatusCodePAINVALIDIDEXCEPTION StatusCode = "PAINVALIDIDEXCEPTION"

	StatusCodePANOACCESSEXCEPTION StatusCode = "PANOACCESSEXCEPTION"

	StatusCodePANODATAFOUNDEXCEPTION StatusCode = "PANODATAFOUNDEXCEPTION"

	StatusCodePAURISYNTAXEXCEPTION StatusCode = "PAURISYNTAXEXCEPTION"

	StatusCodePAVISIBLEACTIONSFILTERORDERINGEXCEPTION StatusCode = "PAVISIBLEACTIONSFILTERORDERINGEXCEPTION"

	StatusCodePORTALNOACCESS StatusCode = "PORTALNOACCESS"

	StatusCodePORTALUSERALREADYEXISTSFORCONTACT StatusCode = "PORTALUSERALREADYEXISTSFORCONTACT"

	StatusCodePORTALUSERCREATIONRESTRICTEDWITHENCRYPTION StatusCode = "PORTALUSERCREATIONRESTRICTEDWITHENCRYPTION"

	StatusCodePRIVATECONTACTONASSET StatusCode = "PRIVATECONTACTONASSET"

	StatusCodePROCESSINGHALTED StatusCode = "PROCESSINGHALTED"

	StatusCodeQAINVALIDCREATEFEEDITEM StatusCode = "QAINVALIDCREATEFEEDITEM"

	StatusCodeQAINVALIDSUCCESSMESSAGE StatusCode = "QAINVALIDSUCCESSMESSAGE"

	StatusCodeQUERYTIMEOUT StatusCode = "QUERYTIMEOUT"

	StatusCodeQUICKACTIONLISTITEMNOTALLOWED StatusCode = "QUICKACTIONLISTITEMNOTALLOWED"

	StatusCodeQUICKACTIONLISTNOTALLOWED StatusCode = "QUICKACTIONLISTNOTALLOWED"

	StatusCodeRECORDINUSEBYWORKFLOW StatusCode = "RECORDINUSEBYWORKFLOW"

	StatusCodeRELATEDENTITYFILTERVALIDATIONEXCEPTION StatusCode = "RELATEDENTITYFILTERVALIDATIONEXCEPTION"

	StatusCodeRELFIELDBADACCESSIBILITY StatusCode = "RELFIELDBADACCESSIBILITY"

	StatusCodeREPUTATIONMINIMUMNUMBERNOTREACHED StatusCode = "REPUTATIONMINIMUMNUMBERNOTREACHED"

	StatusCodeREQUESTRUNNINGTOOLONG StatusCode = "REQUESTRUNNINGTOOLONG"

	StatusCodeREQUIREDFEATUREMISSING StatusCode = "REQUIREDFEATUREMISSING"

	StatusCodeREQUIREDFIELDMISSING StatusCode = "REQUIREDFIELDMISSING"

	StatusCodeREQUIRECONNECTEDAPPSCS StatusCode = "REQUIRECONNECTEDAPPSCS"

	StatusCodeREQUIRECONNECTEDAPPSESSIONSCS StatusCode = "REQUIRECONNECTEDAPPSESSIONSCS"

	StatusCodeREQUIRERUNASUSER StatusCode = "REQUIRERUNASUSER"

	StatusCodeRETRIEVEEXCHANGEATTACHMENTFAILED StatusCode = "RETRIEVEEXCHANGEATTACHMENTFAILED"

	StatusCodeRETRIEVEEXCHANGEEMAILFAILED StatusCode = "RETRIEVEEXCHANGEEMAILFAILED"

	StatusCodeRETRIEVEEXCHANGEEVENTFAILED StatusCode = "RETRIEVEEXCHANGEEVENTFAILED"

	StatusCodeRETRIEVEGOOGLEEMAILFAILED StatusCode = "RETRIEVEGOOGLEEMAILFAILED"

	StatusCodeRETRIEVEGOOGLEEVENTFAILED StatusCode = "RETRIEVEGOOGLEEVENTFAILED"

	StatusCodeRETRIEVEUSERCONFIGERROR StatusCode = "RETRIEVEUSERCONFIGERROR"

	StatusCodeSALESFORCEINBOXTRANSPORTCONNECTIONERROR StatusCode = "SALESFORCEINBOXTRANSPORTCONNECTIONERROR"

	StatusCodeSALESFORCEINBOXTRANSPORTTOKENERROR StatusCode = "SALESFORCEINBOXTRANSPORTTOKENERROR"

	StatusCodeSALESFORCEINBOXTRANSPORTUNKNOWNERROR StatusCode = "SALESFORCEINBOXTRANSPORTUNKNOWNERROR"

	StatusCodeSELFREFERENCEFROMFLOW StatusCode = "SELFREFERENCEFROMFLOW"

	StatusCodeSELFREFERENCEFROMTRIGGER StatusCode = "SELFREFERENCEFROMTRIGGER"

	StatusCodeSHARENEEDEDFORCHILDOWNER StatusCode = "SHARENEEDEDFORCHILDOWNER"

	StatusCodeSINGLEEMAILLIMITEXCEEDED StatusCode = "SINGLEEMAILLIMITEXCEEDED"

	StatusCodeSOCIALACCOUNTNOTFOUND StatusCode = "SOCIALACCOUNTNOTFOUND"

	StatusCodeSOCIALACTIONINVALID StatusCode = "SOCIALACTIONINVALID"

	StatusCodeSOCIALPERSONANOTFOUND StatusCode = "SOCIALPERSONANOTFOUND"

	StatusCodeSOCIALPOSTINVALID StatusCode = "SOCIALPOSTINVALID"

	StatusCodeSOCIALPOSTNOTFOUND StatusCode = "SOCIALPOSTNOTFOUND"

	StatusCodeSTANDARDPRICENOTDEFINED StatusCode = "STANDARDPRICENOTDEFINED"

	StatusCodeSTORAGELIMITEXCEEDED StatusCode = "STORAGELIMITEXCEEDED"

	StatusCodeSTRINGTOOLONG StatusCode = "STRINGTOOLONG"

	StatusCodeSUBDOMAININUSE StatusCode = "SUBDOMAININUSE"

	StatusCodeTABSETLIMITEXCEEDED StatusCode = "TABSETLIMITEXCEEDED"

	StatusCodeTEMPLATENOTACTIVE StatusCode = "TEMPLATENOTACTIVE"

	StatusCodeTEMPLATENOTFOUND StatusCode = "TEMPLATENOTFOUND"

	StatusCodeTERMSOFSERVICEUNREAD StatusCode = "TERMSOFSERVICEUNREAD"

	StatusCodeTERRITORYREALIGNINPROGRESS StatusCode = "TERRITORYREALIGNINPROGRESS"

	StatusCodeTEXTDATAOUTSIDESUPPORTEDCHARSET StatusCode = "TEXTDATAOUTSIDESUPPORTEDCHARSET"

	StatusCodeTOOMANYAPEXREQUESTS StatusCode = "TOOMANYAPEXREQUESTS"

	StatusCodeTOOMANYENUMVALUE StatusCode = "TOOMANYENUMVALUE"

	StatusCodeTOOMANYPOSSIBLEUSERSEXIST StatusCode = "TOOMANYPOSSIBLEUSERSEXIST"

	StatusCodeTRANSFERREQUIRESREAD StatusCode = "TRANSFERREQUIRESREAD"

	StatusCodeUNABLETOLOCKROW StatusCode = "UNABLETOLOCKROW"

	StatusCodeUNAVAILABLERECORDTYPEEXCEPTION StatusCode = "UNAVAILABLERECORDTYPEEXCEPTION"

	StatusCodeUNAVAILABLEREF StatusCode = "UNAVAILABLEREF"

	StatusCodeUNDELETEFAILED StatusCode = "UNDELETEFAILED"

	StatusCodeUNKNOWNEXCEPTION StatusCode = "UNKNOWNEXCEPTION"

	StatusCodeUNKNOWNTOKENERROR StatusCode = "UNKNOWNTOKENERROR"

	StatusCodeUNSAFEHTMLCONTENT StatusCode = "UNSAFEHTMLCONTENT"

	StatusCodeUNSPECIFIEDEMAILADDRESS StatusCode = "UNSPECIFIEDEMAILADDRESS"

	StatusCodeUNSUPPORTEDAPEXTRIGGEROPERATON StatusCode = "UNSUPPORTEDAPEXTRIGGEROPERATON"

	StatusCodeUNSUPPORTEDSOCIALPROVIDER StatusCode = "UNSUPPORTEDSOCIALPROVIDER"

	StatusCodeUNVERIFIEDSENDERADDRESS StatusCode = "UNVERIFIEDSENDERADDRESS"

	StatusCodeUPDATEGOOGLEEMAILLABELFAILED StatusCode = "UPDATEGOOGLEEMAILLABELFAILED"

	StatusCodeUSEROWNSPORTALACCOUNTEXCEPTION StatusCode = "USEROWNSPORTALACCOUNTEXCEPTION"

	StatusCodeUSERWITHAPEXSHARESEXCEPTION StatusCode = "USERWITHAPEXSHARESEXCEPTION"

	StatusCodeVFCOMPILEERROR StatusCode = "VFCOMPILEERROR"

	StatusCodeWEBLINKSIZELIMITEXCEEDED StatusCode = "WEBLINKSIZELIMITEXCEEDED"

	StatusCodeWEBLINKURLINVALID StatusCode = "WEBLINKURLINVALID"

	StatusCodeWRONGCONTROLLERTYPE StatusCode = "WRONGCONTROLLERTYPE"

	StatusCodeXCLEANDJMATCHIGNORABLEERROR StatusCode = "XCLEANDJMATCHIGNORABLEERROR"

	StatusCodeXCLEANDJMATCHINTERNALDJERROR StatusCode = "XCLEANDJMATCHINTERNALDJERROR"

	StatusCodeXCLEANDJMATCHNONRETRIABLEERROR StatusCode = "XCLEANDJMATCHNONRETRIABLEERROR"

	StatusCodeXCLEANDJMATCHRETRIABLEERROR StatusCode = "XCLEANDJMATCHRETRIABLEERROR"

	StatusCodeXCLEANDJMATCHUNKNOWNERROR StatusCode = "XCLEANDJMATCHUNKNOWNERROR"

	StatusCodeXCLEANUNEXPECTEDERROR StatusCode = "XCLEANUNEXPECTEDERROR"
)

type ExtendedErrorCode string

const (

	// Errors with this extended error code have the following properties: severity, actionCallName, parameterName
	ExtendedErrorCodeACTIONCALLDUPLICATEINPUTPARAM ExtendedErrorCode = "ACTIONCALLDUPLICATEINPUTPARAM"

	// Errors with this extended error code have the following properties: severity, actionCallName, parameterName
	ExtendedErrorCodeACTIONCALLDUPLICATEOUTPUTPARAM ExtendedErrorCode = "ACTIONCALLDUPLICATEOUTPUTPARAM"

	// Errors with this extended error code have the following properties: severity
	ExtendedErrorCodeACTIONCALLMISSINGNAME ExtendedErrorCode = "ACTIONCALLMISSINGNAME"

	// Errors with this extended error code have the following properties: severity, actionCallName, parameterName
	ExtendedErrorCodeACTIONCALLMISSINGREQUIREDPARAM ExtendedErrorCode = "ACTIONCALLMISSINGREQUIREDPARAM"

	// Errors with this extended error code have the following properties: severity, actionCallName
	ExtendedErrorCodeACTIONCALLMISSINGREQUIREDTYPE ExtendedErrorCode = "ACTIONCALLMISSINGREQUIREDTYPE"

	// Errors with this extended error code have the following properties: severity
	ExtendedErrorCodeACTIONCALLNOTFOUNDWITHNAMEANDTYPE ExtendedErrorCode = "ACTIONCALLNOTFOUNDWITHNAMEANDTYPE"

	// Errors with this extended error code have the following properties: severity, processType
	ExtendedErrorCodeACTIONCALLNOTSUPPORTEDFORPROCESSTYPE ExtendedErrorCode = "ACTIONCALLNOTSUPPORTEDFORPROCESSTYPE"

	// Errors with this extended error code have the following properties: severity, apexClassName, parameterName
	ExtendedErrorCodeAPEXCALLOUTINPUTDUPLICATE ExtendedErrorCode = "APEXCALLOUTINPUTDUPLICATE"

	// Errors with this extended error code have the following properties: severity, apexClassName, parameterName
	ExtendedErrorCodeAPEXCALLOUTINPUTINCOMPATIBLEDATATYPE ExtendedErrorCode = "APEXCALLOUTINPUTINCOMPATIBLEDATATYPE"

	// Errors with this extended error code have the following properties: severity, apexClassName
	ExtendedErrorCodeAPEXCALLOUTINVALID ExtendedErrorCode = "APEXCALLOUTINVALID"

	// Errors with this extended error code have the following properties: severity, apexClassName
	ExtendedErrorCodeAPEXCALLOUTMISSINGCLASSNAME ExtendedErrorCode = "APEXCALLOUTMISSINGCLASSNAME"

	// Errors with this extended error code have the following properties: severity, apexClassName
	ExtendedErrorCodeAPEXCALLOUTNOTFOUND ExtendedErrorCode = "APEXCALLOUTNOTFOUND"

	// Errors with this extended error code have the following properties: severity, apexClassName, parameterName
	ExtendedErrorCodeAPEXCALLOUTOUTPUTINCOMPATIBLEDATATYPE ExtendedErrorCode = "APEXCALLOUTOUTPUTINCOMPATIBLEDATATYPE"

	// Errors with this extended error code have the following properties: severity, apexClassName, parameterName
	ExtendedErrorCodeAPEXCALLOUTOUTPUTNOTFOUND ExtendedErrorCode = "APEXCALLOUTOUTPUTNOTFOUND"

	// Errors with this extended error code have the following properties: severity, apexClassName, parameterName
	ExtendedErrorCodeAPEXCALLOUTREQUIREDINPUTMISSING ExtendedErrorCode = "APEXCALLOUTREQUIREDINPUTMISSING"

	// Errors with this extended error code have the following properties: severity, apexClassName, parentScreenFieldName
	ExtendedErrorCodeAPEXCLASSMISSINGINTERFACE ExtendedErrorCode = "APEXCLASSMISSINGINTERFACE"

	// Errors with this extended error code have the following properties: severity, assignmentName, operatorName, elementName
	ExtendedErrorCodeASSIGNMENTITEMELEMENTMISSINGDATATYPE ExtendedErrorCode = "ASSIGNMENTITEMELEMENTMISSINGDATATYPE"

	// Errors with this extended error code have the following properties: severity, elementName, assignmentName, elementType
	ExtendedErrorCodeASSIGNMENTITEMELEMENTNOTSUPPORTED ExtendedErrorCode = "ASSIGNMENTITEMELEMENTNOTSUPPORTED"

	// Errors with this extended error code have the following properties: severity, fieldValue, dataType, incompatibleDataType
	ExtendedErrorCodeASSIGNMENTITEMFIELDINVALIDDATATYPE ExtendedErrorCode = "ASSIGNMENTITEMFIELDINVALIDDATATYPE"

	// Errors with this extended error code have the following properties: severity, elementName, acceptedDataType, dataType, fieldValue
	ExtendedErrorCodeASSIGNMENTITEMFIELDINVALIDDATATYPEWITHELEMENT ExtendedErrorCode = "ASSIGNMENTITEMFIELDINVALIDDATATYPEWITHELEMENT"

	// Errors with this extended error code have the following properties: severity, assignmentName, operatorName, leftElementName, leftElementType, rightElementName, rightElementType
	ExtendedErrorCodeASSIGNMENTITEMINCOMPATIBLEDATATYPES ExtendedErrorCode = "ASSIGNMENTITEMINCOMPATIBLEDATATYPES"

	// Errors with this extended error code have the following properties: severity, assignmentName, operatorName, leftElementName, rightElementName
	ExtendedErrorCodeASSIGNMENTITEMINVALIDCOLLECTION ExtendedErrorCode = "ASSIGNMENTITEMINVALIDCOLLECTION"

	// Errors with this extended error code have the following properties: severity, elementName, dataType, incompatibleDataType
	ExtendedErrorCodeASSIGNMENTITEMINVALIDDATATYPEINELEMENT ExtendedErrorCode = "ASSIGNMENTITEMINVALIDDATATYPEINELEMENT"

	// Errors with this extended error code have the following properties: severity, parameterName, operatorName
	ExtendedErrorCodeASSIGNMENTITEMINVALIDREFERENCE ExtendedErrorCode = "ASSIGNMENTITEMINVALIDREFERENCE"

	// Errors with this extended error code have the following properties: severity, assignmentName, operatorName, dataType, elementName
	ExtendedErrorCodeASSIGNMENTITEMLEFTDATATYPEINVALIDFOROPERATOR ExtendedErrorCode = "ASSIGNMENTITEMLEFTDATATYPEINVALIDFOROPERATOR"

	// Errors with this extended error code have the following properties: severity, assignmentName
	ExtendedErrorCodeASSIGNMENTITEMMODIFIESNONVARIABLE ExtendedErrorCode = "ASSIGNMENTITEMMODIFIESNONVARIABLE"

	// Errors with this extended error code have the following properties: severity, parameterName, operatorName
	ExtendedErrorCodeASSIGNMENTITEMNONEXISTENTREFERENCE ExtendedErrorCode = "ASSIGNMENTITEMNONEXISTENTREFERENCE"

	// Errors with this extended error code have the following properties: severity, assignmentName
	ExtendedErrorCodeASSIGNMENTITEMREQUIRED ExtendedErrorCode = "ASSIGNMENTITEMREQUIRED"

	// Errors with this extended error code have the following properties: severity, elementName
	ExtendedErrorCodeASSIGNMENTITEMRIGHTDATATYPEINVALIDFOROPERATOR ExtendedErrorCode = "ASSIGNMENTITEMRIGHTDATATYPEINVALIDFOROPERATOR"

	// Errors with this extended error code have the following properties: severity, choiceLookupName
	ExtendedErrorCodeAUTOLAUNCHEDCHOICELOOKUPNOTSUPPORTED ExtendedErrorCode = "AUTOLAUNCHEDCHOICELOOKUPNOTSUPPORTED"

	// Errors with this extended error code have the following properties: severity, choiceName
	ExtendedErrorCodeAUTOLAUNCHEDCHOICENOTSUPPORTED ExtendedErrorCode = "AUTOLAUNCHEDCHOICENOTSUPPORTED"

	// Errors with this extended error code have the following properties: severity
	ExtendedErrorCodeAUTOLAUNCHEDSCREENNOTSUPPORTED ExtendedErrorCode = "AUTOLAUNCHEDSCREENNOTSUPPORTED"

	// Errors with this extended error code have the following properties: severity
	ExtendedErrorCodeAUTOLAUNCHEDSTEPNOTSUPPORTED ExtendedErrorCode = "AUTOLAUNCHEDSTEPNOTSUPPORTED"

	// Errors with this extended error code have the following properties: severity, subflowType
	ExtendedErrorCodeAUTOLAUNCHEDSUBFLOWINCOMPATIBLEFLOWTYPE ExtendedErrorCode = "AUTOLAUNCHEDSUBFLOWINCOMPATIBLEFLOWTYPE"

	// Errors with this extended error code have the following properties: severity
	ExtendedErrorCodeAUTOLAUNCHEDWAITNOTSUPPORTED ExtendedErrorCode = "AUTOLAUNCHEDWAITNOTSUPPORTED"

	// Errors with this extended error code have the following properties: severity, screenFieldName
	ExtendedErrorCodeCHOICEFIELDDEFAULTCHOICENOTFOUND ExtendedErrorCode = "CHOICEFIELDDEFAULTCHOICENOTFOUND"

	// Errors with this extended error code have the following properties: severity, questionName
	ExtendedErrorCodeCHOICEFIELDMISSINGCHOICE ExtendedErrorCode = "CHOICEFIELDMISSINGCHOICE"

	// Errors with this extended error code have the following properties: severity, choiceName, parentScreenFieldName
	ExtendedErrorCodeCHOICELOOKUPDATATYPEINCOMPATIBLEWITHCHOICEFIELD ExtendedErrorCode = "CHOICELOOKUPDATATYPEINCOMPATIBLEWITHCHOICEFIELD"

	// Errors with this extended error code have the following properties: severity, choiceName, parentScreenFieldName
	ExtendedErrorCodeCHOICEDATATYPEINCOMPATIBLEWITHCHOICEFIELD ExtendedErrorCode = "CHOICEDATATYPEINCOMPATIBLEWITHCHOICEFIELD"

	// Errors with this extended error code have the following properties: severity, elementName, screenFieldName
	ExtendedErrorCodeCHOICENOTSUPPORTEDFORSCREENFIELDTYPE ExtendedErrorCode = "CHOICENOTSUPPORTEDFORSCREENFIELDTYPE"

	// Errors with this extended error code have the following properties: severity, choiceName
	ExtendedErrorCodeCHOICEUSEDMULTIPLETIMESINSAMEFIELD ExtendedErrorCode = "CHOICEUSEDMULTIPLETIMESINSAMEFIELD"

	// Errors with this extended error code have the following properties: severity
	ExtendedErrorCodeCONDITIONBUILDERMISSINGFLOWVARIABLE ExtendedErrorCode = "CONDITIONBUILDERMISSINGFLOWVARIABLE"

	// Errors with this extended error code have the following properties: severity
	ExtendedErrorCodeCONDITIONBUILDERMISSINGREQUIREDPERMISSIONS ExtendedErrorCode = "CONDITIONBUILDERMISSINGREQUIREDPERMISSIONS"

	// Errors with this extended error code have the following properties: severity
	ExtendedErrorCodeCONDITIONBUILDERUNSUPPORTEDFLOWVARIABLE ExtendedErrorCode = "CONDITIONBUILDERUNSUPPORTEDFLOWVARIABLE"

	// Errors with this extended error code have the following properties: severity, leftElementName, leftElementType, operatorName, rightElementName, rightElementType, ruleName
	ExtendedErrorCodeCONDITIONDATATYPEINCOMPATIBLE ExtendedErrorCode = "CONDITIONDATATYPEINCOMPATIBLE"

	// Errors with this extended error code have the following properties: severity, elementName, dataType, operatorName, parameterName, ruleName
	ExtendedErrorCodeCONDITIONDATATYPEINCOMPATIBLEWITHELEMENT ExtendedErrorCode = "CONDITIONDATATYPEINCOMPATIBLEWITHELEMENT"

	// Errors with this extended error code have the following properties: severity, elementName, leftElementType, operatorName, rightElementType, ruleName
	ExtendedErrorCodeCONDITIONELEMENTDATATYPESINCOMPATIBLE ExtendedErrorCode = "CONDITIONELEMENTDATATYPESINCOMPATIBLE"

	// Errors with this extended error code have the following properties: severity, ruleName
	ExtendedErrorCodeCONDITIONINVALIDLEFTOPERAND ExtendedErrorCode = "CONDITIONINVALIDLEFTOPERAND"

	// Errors with this extended error code have the following properties: severity, elementName, dataType, operatorName, parameterName, ruleName
	ExtendedErrorCodeCONDITIONINVALIDLEFTELEMENT ExtendedErrorCode = "CONDITIONINVALIDLEFTELEMENT"

	// Errors with this extended error code have the following properties: severity, elementName, characterLimit
	ExtendedErrorCodeCONDITIONLOGICEXCEEDSLIMIT ExtendedErrorCode = "CONDITIONLOGICEXCEEDSLIMIT"

	// Errors with this extended error code have the following properties: severity, elementName
	ExtendedErrorCodeCONDITIONLOGICINVALID ExtendedErrorCode = "CONDITIONLOGICINVALID"

	// Errors with this extended error code have the following properties: severity, elementName
	ExtendedErrorCodeCONDITIONLOGICMISSING ExtendedErrorCode = "CONDITIONLOGICMISSING"

	// Errors with this extended error code have the following properties: severity, elementName, dataType, operatorName, parameterName, ruleName
	ExtendedErrorCodeCONDITIONMISSINGDATATYPE ExtendedErrorCode = "CONDITIONMISSINGDATATYPE"

	// Errors with this extended error code have the following properties: severity, ruleName
	ExtendedErrorCodeCONDITIONMISSINGOPERATOR ExtendedErrorCode = "CONDITIONMISSINGOPERATOR"

	// Errors with this extended error code have the following properties: severity, ruleName
	ExtendedErrorCodeCONDITIONREFERENCEDELEMENTNOTFOUND ExtendedErrorCode = "CONDITIONREFERENCEDELEMENTNOTFOUND"

	// Errors with this extended error code have the following properties: severity, ruleName
	ExtendedErrorCodeCONDITIONRIGHTOPERANDNULL ExtendedErrorCode = "CONDITIONRIGHTOPERANDNULL"

	// Errors with this extended error code have the following properties: severity, elementName
	ExtendedErrorCodeCONNECTORMISSINGTARGET ExtendedErrorCode = "CONNECTORMISSINGTARGET"

	// Errors with this extended error code have the following properties: severity, constantName
	ExtendedErrorCodeCONSTANTINCLUDESREFERENCES ExtendedErrorCode = "CONSTANTINCLUDESREFERENCES"

	// Errors with this extended error code have the following properties: severity
	ExtendedErrorCodeCUSTOMEVENTSNOTENABLED ExtendedErrorCode = "CUSTOMEVENTSNOTENABLED"

	// Errors with this extended error code have the following properties: severity
	ExtendedErrorCodeCUSTOMEVENTMISSINGPROCESSMETADATAVALUES ExtendedErrorCode = "CUSTOMEVENTMISSINGPROCESSMETADATAVALUES"

	// Errors with this extended error code have the following properties: severity, objectType
	ExtendedErrorCodeCUSTOMEVENTOBJECTTYPENOTFOUND ExtendedErrorCode = "CUSTOMEVENTOBJECTTYPENOTFOUND"

	// Errors with this extended error code have the following properties: severity, objectType
	ExtendedErrorCodeCUSTOMEVENTOBJECTTYPENOTSUPPORTED ExtendedErrorCode = "CUSTOMEVENTOBJECTTYPENOTSUPPORTED"

	// Errors with this extended error code have the following properties: severity, metadataValue
	ExtendedErrorCodeCUSTOMEVENTPROCESSMETADATAVALUESMISSINGNAME ExtendedErrorCode = "CUSTOMEVENTPROCESSMETADATAVALUESMISSINGNAME"

	// Errors with this extended error code have the following properties: severity, metadataValue
	ExtendedErrorCodeCUSTOMEVENTPROCESSMETADATAVALUESMORETHANONENAME ExtendedErrorCode = "CUSTOMEVENTPROCESSMETADATAVALUESMORETHANONENAME"

	// Errors with this extended error code have the following properties: severity, elementName, dataType
	ExtendedErrorCodeDATATYPEINVALID ExtendedErrorCode = "DATATYPEINVALID"

	// Errors with this extended error code have the following properties: severity, elementName
	ExtendedErrorCodeDATATYPEMISSING ExtendedErrorCode = "DATATYPEMISSING"

	// Errors with this extended error code have the following properties: severity, flowDecision
	ExtendedErrorCodeDECISIONDEFAULTCONNECTORMISSINGLABEL ExtendedErrorCode = "DECISIONDEFAULTCONNECTORMISSINGLABEL"

	// Errors with this extended error code have the following properties: severity, flowDecision
	ExtendedErrorCodeDECISIONMISSINGOUTCOME ExtendedErrorCode = "DECISIONMISSINGOUTCOME"

	// Errors with this extended error code have the following properties: severity, elementName
	ExtendedErrorCodeELEMENTCONNECTSTOSELF ExtendedErrorCode = "ELEMENTCONNECTSTOSELF"

	// Errors with this extended error code have the following properties: severity, coordinateLimit, coordinateName
	ExtendedErrorCodeELEMENTCOORDINATESINVALID ExtendedErrorCode = "ELEMENTCOORDINATESINVALID"

	// Errors with this extended error code have the following properties: severity, elementName
	ExtendedErrorCodeELEMENTINVALIDCONNECTOR ExtendedErrorCode = "ELEMENTINVALIDCONNECTOR"

	// Errors with this extended error code have the following properties: severity, elementName
	ExtendedErrorCodeELEMENTINVALIDREFERENCE ExtendedErrorCode = "ELEMENTINVALIDREFERENCE"

	// Errors with this extended error code have the following properties: severity, elementName
	ExtendedErrorCodeELEMENTMISSINGCONNECTOR ExtendedErrorCode = "ELEMENTMISSINGCONNECTOR"

	// Errors with this extended error code have the following properties: severity, characterLimit, elementName
	ExtendedErrorCodeELEMENTMISSINGLABEL ExtendedErrorCode = "ELEMENTMISSINGLABEL"

	// Errors with this extended error code have the following properties: severity, characterLimit
	ExtendedErrorCodeELEMENTMISSINGNAME ExtendedErrorCode = "ELEMENTMISSINGNAME"

	// Errors with this extended error code have the following properties: severity, elementName
	ExtendedErrorCodeELEMENTMISSINGREFERENCE ExtendedErrorCode = "ELEMENTMISSINGREFERENCE"

	// Errors with this extended error code have the following properties: severity, elementName
	ExtendedErrorCodeELEMENTMORETHANONEFIELD ExtendedErrorCode = "ELEMENTMORETHANONEFIELD"

	// Errors with this extended error code have the following properties: severity
	ExtendedErrorCodeELEMENTNAMEINVALID ExtendedErrorCode = "ELEMENTNAMEINVALID"

	// Errors with this extended error code have the following properties: severity, elementName
	ExtendedErrorCodeELEMENTNEVERUSED ExtendedErrorCode = "ELEMENTNEVERUSED"

	// Errors with this extended error code have the following properties: severity, elementName
	ExtendedErrorCodeELEMENTSCALESMALLERTHANDEFAULTVALUE ExtendedErrorCode = "ELEMENTSCALESMALLERTHANDEFAULTVALUE"

	// Errors with this extended error code have the following properties: severity, objectName
	ExtendedErrorCodeEXTERNALOBJECTSNOTSUPPORTED ExtendedErrorCode = "EXTERNALOBJECTSNOTSUPPORTED"

	// Errors with this extended error code have the following properties: severity, fieldReference
	ExtendedErrorCodeEXTERNALOBJECTFIELDSNOTSUPPORTED ExtendedErrorCode = "EXTERNALOBJECTFIELDSNOTSUPPORTED"

	// Errors with this extended error code have the following properties: severity, fieldName, elementName
	ExtendedErrorCodeFIELDASSIGNMENTFIELDINCOMPATIBLEDATATYPE ExtendedErrorCode = "FIELDASSIGNMENTFIELDINCOMPATIBLEDATATYPE"

	// Errors with this extended error code have the following properties: severity, fieldName, elementName, assignmentName
	ExtendedErrorCodeFIELDASSIGNMENTINVALIDDATATYPE ExtendedErrorCode = "FIELDASSIGNMENTINVALIDDATATYPE"

	// Errors with this extended error code have the following properties: severity, fieldName, elementName, elementType
	ExtendedErrorCodeFIELDASSIGNMENTINVALIDELEMENT ExtendedErrorCode = "FIELDASSIGNMENTINVALIDELEMENT"

	// Errors with this extended error code have the following properties: severity, fieldName, parameterName
	ExtendedErrorCodeFIELDASSIGNMENTINVALIDREFERENCE ExtendedErrorCode = "FIELDASSIGNMENTINVALIDREFERENCE"

	// Errors with this extended error code have the following properties: severity, fieldName
	ExtendedErrorCodeFIELDASSIGNMENTMULTIPLEREFERENCESSAMEFIELD ExtendedErrorCode = "FIELDASSIGNMENTMULTIPLEREFERENCESSAMEFIELD"

	// Errors with this extended error code have the following properties: severity, fieldName, dataType
	ExtendedErrorCodeFIELDASSIGNMENTPICKLISTFIELDINCOMPATIBLEDATATYPE ExtendedErrorCode = "FIELDASSIGNMENTPICKLISTFIELDINCOMPATIBLEDATATYPE"

	// Errors with this extended error code have the following properties: severity, fieldName, elementName, elementType
	ExtendedErrorCodeFIELDASSIGNMENTREFERENCEDELEMENTMISSINGDATATYPE ExtendedErrorCode = "FIELDASSIGNMENTREFERENCEDELEMENTMISSINGDATATYPE"

	// Errors with this extended error code have the following properties: severity, elementName
	ExtendedErrorCodeFIELDSERVICEUNSUPPORTEDFIELDTYPE ExtendedErrorCode = "FIELDSERVICEUNSUPPORTEDFIELDTYPE"

	// Errors with this extended error code have the following properties: severity, fieldName, parameterName
	ExtendedErrorCodeFIELDINVALIDVALUE ExtendedErrorCode = "FIELDINVALIDVALUE"

	// Errors with this extended error code have the following properties: severity, objectName, fieldName
	ExtendedErrorCodeFIELDNOTFOUND ExtendedErrorCode = "FIELDNOTFOUND"

	// Errors with this extended error code have the following properties: severity, fieldRelationshipName
	ExtendedErrorCodeFIELDRELATIONSHIPNOTSUPPORTED ExtendedErrorCode = "FIELDRELATIONSHIPNOTSUPPORTED"

	// Errors with this extended error code have the following properties: componentName, propertyName, propertyType, errorCode, invalidTokens
	ExtendedErrorCodeFLEXIPAGECOMPONENTATTRIBUTEEXPRESSIONEXCEPTION ExtendedErrorCode = "FLEXIPAGECOMPONENTATTRIBUTEEXPRESSIONEXCEPTION"

	// Errors with this extended error code have the following properties: componentName, propertyName, propertyType, errorIdentifier, errorParams
	ExtendedErrorCodeFLEXIPAGECOMPONENTATTRIBUTEGENERICEXCEPTION ExtendedErrorCode = "FLEXIPAGECOMPONENTATTRIBUTEGENERICEXCEPTION"

	// Errors with this extended error code have the following properties: componentName, propertyName, propertyType
	ExtendedErrorCodeFLEXIPAGECOMPONENTATTRIBUTEMISSINGREQUIRED ExtendedErrorCode = "FLEXIPAGECOMPONENTATTRIBUTEMISSINGREQUIRED"

	// Errors with this extended error code have the following properties: componentName, propertyName, propertyType, maxLength
	ExtendedErrorCodeFLEXIPAGECOMPONENTATTRIBUTETOOLONG ExtendedErrorCode = "FLEXIPAGECOMPONENTATTRIBUTETOOLONG"

	// Errors with this extended error code have the following properties: componentName
	ExtendedErrorCodeFLEXIPAGECOMPONENTCUSTOMVALIDATIONEXCEPTION ExtendedErrorCode = "FLEXIPAGECOMPONENTCUSTOMVALIDATIONEXCEPTION"

	// Errors with this extended error code have the following properties:
	ExtendedErrorCodeFLEXIPAGECOMPONENTMAXLIMITEXCEPTION ExtendedErrorCode = "FLEXIPAGECOMPONENTMAXLIMITEXCEPTION"

	// Errors with this extended error code have the following properties: componentName, criterionIndex
	ExtendedErrorCodeFLEXIPAGECOMPONENTRULEVALIDATIONEXCEPTION ExtendedErrorCode = "FLEXIPAGECOMPONENTRULEVALIDATIONEXCEPTION"

	// Errors with this extended error code have the following properties: componentName, propertyName, propertyType, invalidValue
	ExtendedErrorCodeFLEXIPAGEPICKLISTINVALIDVALUEEXCEPTION ExtendedErrorCode = "FLEXIPAGEPICKLISTINVALIDVALUEEXCEPTION"

	// Errors with this extended error code have the following properties: severity, elementName
	ExtendedErrorCodeFLOWELEMENTSCALELESSTHANZERO ExtendedErrorCode = "FLOWELEMENTSCALELESSTHANZERO"

	// Errors with this extended error code have the following properties: severity, elementName
	ExtendedErrorCodeFLOWINCLUDESSTEP ExtendedErrorCode = "FLOWINCLUDESSTEP"

	// Errors with this extended error code have the following properties: severity
	ExtendedErrorCodeFLOWINTERVIEWBULKEXECUTION ExtendedErrorCode = "FLOWINTERVIEWBULKEXECUTION"

	// Errors with this extended error code have the following properties: severity
	ExtendedErrorCodeFLOWINTERVIEWHANDLEDERROR ExtendedErrorCode = "FLOWINTERVIEWHANDLEDERROR"

	// Errors with this extended error code have the following properties: severity
	ExtendedErrorCodeFLOWINTERVIEWINPUTVALIDATION ExtendedErrorCode = "FLOWINTERVIEWINPUTVALIDATION"

	// Errors with this extended error code have the following properties: severity
	ExtendedErrorCodeFLOWINTERVIEWINTERACTIONNOTFOUND ExtendedErrorCode = "FLOWINTERVIEWINTERACTIONNOTFOUND"

	// Errors with this extended error code have the following properties: severity
	ExtendedErrorCodeFLOWINTERVIEWINVALIDCHOICEUSERINPUT ExtendedErrorCode = "FLOWINTERVIEWINVALIDCHOICEUSERINPUT"

	// Errors with this extended error code have the following properties: severity
	ExtendedErrorCodeFLOWINTERVIEWINVALIDFIELDVALUE ExtendedErrorCode = "FLOWINTERVIEWINVALIDFIELDVALUE"

	// Errors with this extended error code have the following properties: severity
	ExtendedErrorCodeFLOWINTERVIEWINVALIDSTARTREQUEST ExtendedErrorCode = "FLOWINTERVIEWINVALIDSTARTREQUEST"

	// Errors with this extended error code have the following properties: severity
	ExtendedErrorCodeFLOWINTERVIEWLIMITEXCEEDED ExtendedErrorCode = "FLOWINTERVIEWLIMITEXCEEDED"

	// Errors with this extended error code have the following properties: severity
	ExtendedErrorCodeFLOWINTERVIEWMISSINGCHOICEFORREQUIREDCHOICEFIELD ExtendedErrorCode = "FLOWINTERVIEWMISSINGCHOICEFORREQUIREDCHOICEFIELD"

	// Errors with this extended error code have the following properties: severity
	ExtendedErrorCodeFLOWINTERVIEWMISSINGVALUEFORREQUIREDINPUTFIELD ExtendedErrorCode = "FLOWINTERVIEWMISSINGVALUEFORREQUIREDINPUTFIELD"

	// Errors with this extended error code have the following properties: severity
	ExtendedErrorCodeFLOWINTERVIEWNAVIGATE ExtendedErrorCode = "FLOWINTERVIEWNAVIGATE"

	// Errors with this extended error code have the following properties: severity
	ExtendedErrorCodeFLOWINTERVIEWRANGEVALIDATION ExtendedErrorCode = "FLOWINTERVIEWRANGEVALIDATION"

	// Errors with this extended error code have the following properties: severity
	ExtendedErrorCodeFLOWINTERVIEWREGEXVALIDATION ExtendedErrorCode = "FLOWINTERVIEWREGEXVALIDATION"

	// Errors with this extended error code have the following properties: severity
	ExtendedErrorCodeFLOWINTERVIEWRESUMEINTERVIEW ExtendedErrorCode = "FLOWINTERVIEWRESUMEINTERVIEW"

	// Errors with this extended error code have the following properties: severity
	ExtendedErrorCodeFLOWINTERVIEWSAVERESULT ExtendedErrorCode = "FLOWINTERVIEWSAVERESULT"

	// Errors with this extended error code have the following properties: severity
	ExtendedErrorCodeFLOWINTERVIEWSETCHOICESELECTED ExtendedErrorCode = "FLOWINTERVIEWSETCHOICESELECTED"

	// Errors with this extended error code have the following properties: severity
	ExtendedErrorCodeFLOWINTERVIEWSTARTINTERVIEW ExtendedErrorCode = "FLOWINTERVIEWSTARTINTERVIEW"

	// Errors with this extended error code have the following properties: severity
	ExtendedErrorCodeFLOWINTERVIEWTYPECONVERSION ExtendedErrorCode = "FLOWINTERVIEWTYPECONVERSION"

	// Errors with this extended error code have the following properties: severity, maxDevNameLength
	ExtendedErrorCodeFLOWINVALIDNAME ExtendedErrorCode = "FLOWINVALIDNAME"

	// Errors with this extended error code have the following properties: severity, flowName
	ExtendedErrorCodeFLOWNAMEUSEDINOTHERCLIENT ExtendedErrorCode = "FLOWNAMEUSEDINOTHERCLIENT"

	// Errors with this extended error code have the following properties: severity, stageName
	ExtendedErrorCodeFLOWSTAGEINCLUDESREFERENCES ExtendedErrorCode = "FLOWSTAGEINCLUDESREFERENCES"

	// Errors with this extended error code have the following properties: severity, stageName, stageOrder, stageWithSameOrder
	ExtendedErrorCodeFLOWSTAGEORDERDUPLICATE ExtendedErrorCode = "FLOWSTAGEORDERDUPLICATE"

	// Errors with this extended error code have the following properties: severity, stageName, invalidStageOrder, maxOrder, minOrder
	ExtendedErrorCodeFLOWSTAGEORDEROUTOFRANGE ExtendedErrorCode = "FLOWSTAGEORDEROUTOFRANGE"

	// Errors with this extended error code have the following properties: severity, formulaExpression
	ExtendedErrorCodeFORMULAEXPRESSIONINVALID ExtendedErrorCode = "FORMULAEXPRESSIONINVALID"

	// Errors with this extended error code have the following properties: severity, parameterName
	ExtendedErrorCodeINPUTPARAMINCOMPATIBLEDATATYPE ExtendedErrorCode = "INPUTPARAMINCOMPATIBLEDATATYPE"

	// Errors with this extended error code have the following properties: severity, parameterName
	ExtendedErrorCodeINPUTPARAMINCOMPATIBLEWITHCOLLECTIONVARIABLE ExtendedErrorCode = "INPUTPARAMINCOMPATIBLEWITHCOLLECTIONVARIABLE"

	// Errors with this extended error code have the following properties: severity, parameterName
	ExtendedErrorCodeINPUTPARAMINCOMPATIBLEWITHNONCOLLECTIONVARIABLE ExtendedErrorCode = "INPUTPARAMINCOMPATIBLEWITHNONCOLLECTIONVARIABLE"

	// Errors with this extended error code have the following properties: severity, parameterName
	ExtendedErrorCodeINPUTPARAMMISMATCHEDOBJECTTYPE ExtendedErrorCode = "INPUTPARAMMISMATCHEDOBJECTTYPE"

	// Errors with this extended error code have the following properties: severity
	ExtendedErrorCodeINVALIDFLOW ExtendedErrorCode = "INVALIDFLOW"

	// Errors with this extended error code have the following properties: severity
	ExtendedErrorCodeINVALIDFLOWINTERVIEW ExtendedErrorCode = "INVALIDFLOWINTERVIEW"

	// Errors with this extended error code have the following properties: severity, surveyName
	ExtendedErrorCodeINVALIDSURVEYVARIABLENAMEORTYPE ExtendedErrorCode = "INVALIDSURVEYVARIABLENAMEORTYPE"

	// Errors with this extended error code have the following properties: severity, elementName
	ExtendedErrorCodeLOOPASSIGNNEXTVALUETOMISMATCHEDDATATYPE ExtendedErrorCode = "LOOPASSIGNNEXTVALUETOMISMATCHEDDATATYPE"

	// Errors with this extended error code have the following properties: severity, elementName
	ExtendedErrorCodeLOOPASSIGNNEXTVALUETOMISMATCHEDOBJECTTYPE ExtendedErrorCode = "LOOPASSIGNNEXTVALUETOMISMATCHEDOBJECTTYPE"

	// Errors with this extended error code have the following properties: severity, elementName
	ExtendedErrorCodeLOOPASSIGNNEXTVALUETOMISSING ExtendedErrorCode = "LOOPASSIGNNEXTVALUETOMISSING"

	// Errors with this extended error code have the following properties: severity, elementName
	ExtendedErrorCodeLOOPASSIGNNEXTVALUETOMISSINGVARIABLE ExtendedErrorCode = "LOOPASSIGNNEXTVALUETOMISSINGVARIABLE"

	// Errors with this extended error code have the following properties: severity, fieldRelationshipName
	ExtendedErrorCodeLOOPASSIGNNEXTVALUETOREFERENCENOTFOUND ExtendedErrorCode = "LOOPASSIGNNEXTVALUETOREFERENCENOTFOUND"

	// Errors with this extended error code have the following properties: severity, elementName
	ExtendedErrorCodeLOOPCOLLECTIONELEMENTNOTFOUND ExtendedErrorCode = "LOOPCOLLECTIONELEMENTNOTFOUND"

	// Errors with this extended error code have the following properties: severity, elementName
	ExtendedErrorCodeLOOPCOLLECTIONNOTFOUND ExtendedErrorCode = "LOOPCOLLECTIONNOTFOUND"

	// Errors with this extended error code have the following properties: severity, fieldName
	ExtendedErrorCodeLOOPCOLLECTIONNOTSUPPORTEDFORFIELD ExtendedErrorCode = "LOOPCOLLECTIONNOTSUPPORTEDFORFIELD"

	// Errors with this extended error code have the following properties: severity
	ExtendedErrorCodeLOOPMISSINGCOLLECTION ExtendedErrorCode = "LOOPMISSINGCOLLECTION"

	// Errors with this extended error code have the following properties: severity, objectType
	ExtendedErrorCodeOBJECTTYPEINVALID ExtendedErrorCode = "OBJECTTYPEINVALID"

	// Errors with this extended error code have the following properties: severity, objectName
	ExtendedErrorCodeOBJECTCANNOTBECREATED ExtendedErrorCode = "OBJECTCANNOTBECREATED"

	// Errors with this extended error code have the following properties: severity, objectName
	ExtendedErrorCodeOBJECTCANNOTBEDELETED ExtendedErrorCode = "OBJECTCANNOTBEDELETED"

	// Errors with this extended error code have the following properties: severity, objectName
	ExtendedErrorCodeOBJECTCANNOTBEQUERIED ExtendedErrorCode = "OBJECTCANNOTBEQUERIED"

	// Errors with this extended error code have the following properties: severity, objectName
	ExtendedErrorCodeOBJECTCANNOTBEUPDATED ExtendedErrorCode = "OBJECTCANNOTBEUPDATED"

	// Errors with this extended error code have the following properties: severity, fieldName
	ExtendedErrorCodeOBJECTENCRYPTEDFIELDSNOTSUPPORTED ExtendedErrorCode = "OBJECTENCRYPTEDFIELDSNOTSUPPORTED"

	// Errors with this extended error code have the following properties: severity, objectName
	ExtendedErrorCodeOBJECTNOTFOUND ExtendedErrorCode = "OBJECTNOTFOUND"

	// Errors with this extended error code have the following properties: severity, parameterName
	ExtendedErrorCodeOUTPUTPARAMASSIGNTOREFERENCENOTFOUND ExtendedErrorCode = "OUTPUTPARAMASSIGNTOREFERENCENOTFOUND"

	// Errors with this extended error code have the following properties: severity, parameterName
	ExtendedErrorCodeOUTPUTPARAMINCOMPATIBLEDATATYPE ExtendedErrorCode = "OUTPUTPARAMINCOMPATIBLEDATATYPE"

	// Errors with this extended error code have the following properties: severity, parameterName
	ExtendedErrorCodeOUTPUTPARAMMISMATCHEDOBJECTTYPE ExtendedErrorCode = "OUTPUTPARAMMISMATCHEDOBJECTTYPE"

	// Errors with this extended error code have the following properties: severity, parameterName
	ExtendedErrorCodeOUTPUTPARAMMISMATCHEDWITHCOLLECTIONVARIABLE ExtendedErrorCode = "OUTPUTPARAMMISMATCHEDWITHCOLLECTIONVARIABLE"

	// Errors with this extended error code have the following properties: severity, parameterName
	ExtendedErrorCodeOUTPUTPARAMMISSINGASSIGNTOREFERENCE ExtendedErrorCode = "OUTPUTPARAMMISSINGASSIGNTOREFERENCE"

	// Errors with this extended error code have the following properties: severity, parameterName
	ExtendedErrorCodeOUTPUTPARAMMISTMATCHEDWITHNONCOLLECTIONVARIABLE ExtendedErrorCode = "OUTPUTPARAMMISTMATCHEDWITHNONCOLLECTIONVARIABLE"

	// Errors with this extended error code have the following properties: severity, parameterName
	ExtendedErrorCodePARAMDATATYPENOTSUPPORTED ExtendedErrorCode = "PARAMDATATYPENOTSUPPORTED"

	// Errors with this extended error code have the following properties: severity, processType, metadataValue
	ExtendedErrorCodePROCESSMETADATAVALUESNOTSUPPORTEDFORPROCESSTYPE ExtendedErrorCode = "PROCESSMETADATAVALUESNOTSUPPORTEDFORPROCESSTYPE"

	// Errors with this extended error code have the following properties: severity, metadataValue
	ExtendedErrorCodePROCESSMETADATAVALUENONEXISTENTELEMENT ExtendedErrorCode = "PROCESSMETADATAVALUENONEXISTENTELEMENT"

	// Errors with this extended error code have the following properties: severity, processType, elementType
	ExtendedErrorCodePROCESSTYPEELEMENTNOTSUPPORTED ExtendedErrorCode = "PROCESSTYPEELEMENTNOTSUPPORTED"

	// Errors with this extended error code have the following properties: severity, processType
	ExtendedErrorCodePROCESSTYPENOTSUPPORTED ExtendedErrorCode = "PROCESSTYPENOTSUPPORTED"

	// Errors with this extended error code have the following properties: severity, currentProcessType, flowName, incompatibleProcessType
	ExtendedErrorCodePROCESSTYPEINCOMPATIBLE ExtendedErrorCode = "PROCESSTYPEINCOMPATIBLE"

	// Errors with this extended error code have the following properties: severity, fieldName
	ExtendedErrorCodeRECORDFILTERENCRYPTEDFIELDSNOTSUPPORTED ExtendedErrorCode = "RECORDFILTERENCRYPTEDFIELDSNOTSUPPORTED"

	// Errors with this extended error code have the following properties: severity, fieldName, objectName
	ExtendedErrorCodeRECORDFILTERGEOLOCATIONFIELDSNOTSUPPORTED ExtendedErrorCode = "RECORDFILTERGEOLOCATIONFIELDSNOTSUPPORTED"

	// Errors with this extended error code have the following properties: severity, fieldName, elementName, elementType, operatorName
	ExtendedErrorCodeRECORDFILTERINVALIDDATATYPE ExtendedErrorCode = "RECORDFILTERINVALIDDATATYPE"

	// Errors with this extended error code have the following properties: severity, fieldName, assignmentName, elementName, elementType
	ExtendedErrorCodeRECORDFILTERINVALIDELEMENT ExtendedErrorCode = "RECORDFILTERINVALIDELEMENT"

	// Errors with this extended error code have the following properties: severity, fieldName, operatorName
	ExtendedErrorCodeRECORDFILTERINVALIDOPERATOR ExtendedErrorCode = "RECORDFILTERINVALIDOPERATOR"

	// Errors with this extended error code have the following properties: severity, fieldName, operatorName
	ExtendedErrorCodeRECORDFILTERINVALIDREFERENCE ExtendedErrorCode = "RECORDFILTERINVALIDREFERENCE"

	// Errors with this extended error code have the following properties: severity, fieldName, elementName, elementType, operatorName
	ExtendedErrorCodeRECORDFILTERMISSINGDATATYPE ExtendedErrorCode = "RECORDFILTERMISSINGDATATYPE"

	// Errors with this extended error code have the following properties: severity, fieldName
	ExtendedErrorCodeRECORDFILTERMULTIPLEQUERIESSAMEFIELD ExtendedErrorCode = "RECORDFILTERMULTIPLEQUERIESSAMEFIELD"

	// Errors with this extended error code have the following properties: severity, elementName
	ExtendedErrorCodeRECORDLOOKUPIDASSIGNMENTVARIABLEINCOMPATIBLEDATATYPE ExtendedErrorCode = "RECORDLOOKUPIDASSIGNMENTVARIABLEINCOMPATIBLEDATATYPE"

	// Errors with this extended error code have the following properties: severity, elementName
	ExtendedErrorCodeRECORDLOOKUPIDASSIGNMENTVARIABLENOTFOUND ExtendedErrorCode = "RECORDLOOKUPIDASSIGNMENTVARIABLENOTFOUND"

	// Errors with this extended error code have the following properties: severity, objectName
	ExtendedErrorCodeRECORDUPDATEMISSINGFILTERS ExtendedErrorCode = "RECORDUPDATEMISSINGFILTERS"

	// Errors with this extended error code have the following properties: severity, elementName, mergeFieldReference
	ExtendedErrorCodeREFERENCEDELEMENTNOTFOUND ExtendedErrorCode = "REFERENCEDELEMENTNOTFOUND"

	// Errors with this extended error code have the following properties: severity, elementName, ruleName
	ExtendedErrorCodeRULEMISSINGCONDITION ExtendedErrorCode = "RULEMISSINGCONDITION"

	// Errors with this extended error code have the following properties: severity, fieldName
	ExtendedErrorCodeSCREENFIELDAPIVERSIONNOTSUPPORTED ExtendedErrorCode = "SCREENFIELDAPIVERSIONNOTSUPPORTED"

	// Errors with this extended error code have the following properties: severity, fieldName
	ExtendedErrorCodeSCREENFIELDBOOLEANISREQUIREDISFALSE ExtendedErrorCode = "SCREENFIELDBOOLEANISREQUIREDISFALSE"

	// Errors with this extended error code have the following properties: severity, elementName
	ExtendedErrorCodeSCREENFIELDDEFAULTVALUENOTSUPPORTED ExtendedErrorCode = "SCREENFIELDDEFAULTVALUENOTSUPPORTED"

	// Errors with this extended error code have the following properties: severity, elementName
	ExtendedErrorCodeSCREENFIELDEXTENSIONCOMPONENTNOTGLOBAL ExtendedErrorCode = "SCREENFIELDEXTENSIONCOMPONENTNOTGLOBAL"

	// Errors with this extended error code have the following properties: severity, elementName, extensionName, parameterName
	ExtendedErrorCodeSCREENFIELDEXTENSIONDUPLICATEINPUTPARAM ExtendedErrorCode = "SCREENFIELDEXTENSIONDUPLICATEINPUTPARAM"

	// Errors with this extended error code have the following properties: severity, elementName, extensionName, parameterName
	ExtendedErrorCodeSCREENFIELDEXTENSIONDUPLICATEOUTPUTPARAM ExtendedErrorCode = "SCREENFIELDEXTENSIONDUPLICATEOUTPUTPARAM"

	// Errors with this extended error code have the following properties: severity, elementName, extensionName
	ExtendedErrorCodeSCREENFIELDEXTENSIONIMPLEMENTATIONINVALID ExtendedErrorCode = "SCREENFIELDEXTENSIONIMPLEMENTATIONINVALID"

	// Errors with this extended error code have the following properties: severity, elementName, extensionName, parameterName
	ExtendedErrorCodeSCREENFIELDEXTENSIONINPUTATTRIBUTEINVALID ExtendedErrorCode = "SCREENFIELDEXTENSIONINPUTATTRIBUTEINVALID"

	// Errors with this extended error code have the following properties: severity, elementName, extensionName
	ExtendedErrorCodeSCREENFIELDEXTENSIONNAMEINVALID ExtendedErrorCode = "SCREENFIELDEXTENSIONNAMEINVALID"

	// Errors with this extended error code have the following properties: severity, elementName, fieldType
	ExtendedErrorCodeSCREENFIELDEXTENSIONNAMEMISSING ExtendedErrorCode = "SCREENFIELDEXTENSIONNAMEMISSING"

	// Errors with this extended error code have the following properties: severity, elementName, fieldType
	ExtendedErrorCodeSCREENFIELDEXTENSIONNAMENOTSUPPORTED ExtendedErrorCode = "SCREENFIELDEXTENSIONNAMENOTSUPPORTED"

	// Errors with this extended error code have the following properties: severity, elementName, extensionName, parameterName
	ExtendedErrorCodeSCREENFIELDEXTENSIONOUTPUTATTRIBUTEINVALID ExtendedErrorCode = "SCREENFIELDEXTENSIONOUTPUTATTRIBUTEINVALID"

	// Errors with this extended error code have the following properties: severity, elementName, extensionName, parameterName
	ExtendedErrorCodeSCREENFIELDEXTENSIONREQUIREDINPUTMISSING ExtendedErrorCode = "SCREENFIELDEXTENSIONREQUIREDINPUTMISSING"

	// Errors with this extended error code have the following properties: severity, elementName, fieldType
	ExtendedErrorCodeSCREENFIELDINPUTSNOTSUPPORTED ExtendedErrorCode = "SCREENFIELDINPUTSNOTSUPPORTED"

	// Errors with this extended error code have the following properties: severity, dataType, fieldType
	ExtendedErrorCodeSCREENFIELDINVALIDDATATYPE ExtendedErrorCode = "SCREENFIELDINVALIDDATATYPE"

	// Errors with this extended error code have the following properties: severity, choiceName
	ExtendedErrorCodeSCREENFIELDMULTISELECTCHOICESEMICOLONNOTSUPPORTED ExtendedErrorCode = "SCREENFIELDMULTISELECTCHOICESEMICOLONNOTSUPPORTED"

	// Errors with this extended error code have the following properties: severity, elementName, fieldType
	ExtendedErrorCodeSCREENFIELDOUTPUTSNOTSUPPORTED ExtendedErrorCode = "SCREENFIELDOUTPUTSNOTSUPPORTED"

	// Errors with this extended error code have the following properties: severity, elementName, fieldType
	ExtendedErrorCodeSCREENFIELDTYPENOTSUPPORTED ExtendedErrorCode = "SCREENFIELDTYPENOTSUPPORTED"

	// Errors with this extended error code have the following properties: severity, choiceName
	ExtendedErrorCodeSCREENFIELDUSERINPUTNOTSUPPORTEDFORCHOICETYPE ExtendedErrorCode = "SCREENFIELDUSERINPUTNOTSUPPORTEDFORCHOICETYPE"

	// Errors with this extended error code have the following properties: severity, elementName
	ExtendedErrorCodeSCREENFIELDVALIDATIONRULENOTSUPPORTED ExtendedErrorCode = "SCREENFIELDVALIDATIONRULENOTSUPPORTED"

	// Errors with this extended error code have the following properties: severity, screenRuleName, attributeName
	ExtendedErrorCodeSCREENRULEACTIONINVALIDATTRIBUTE ExtendedErrorCode = "SCREENRULEACTIONINVALIDATTRIBUTE"

	// Errors with this extended error code have the following properties: severity, screenRuleName, attributeName
	ExtendedErrorCodeSCREENRULEACTIONINVALIDATTRIBUTEFORAPIVERSION ExtendedErrorCode = "SCREENRULEACTIONINVALIDATTRIBUTEFORAPIVERSION"

	// Errors with this extended error code have the following properties: severity, screenRuleName, acceptedValues, actionValue
	ExtendedErrorCodeSCREENRULEACTIONINVALIDVALUE ExtendedErrorCode = "SCREENRULEACTIONINVALIDVALUE"

	// Errors with this extended error code have the following properties: severity, screenRuleName
	ExtendedErrorCodeSCREENRULEACTIONMISSINGATTRIBUTE ExtendedErrorCode = "SCREENRULEACTIONMISSINGATTRIBUTE"

	// Errors with this extended error code have the following properties: severity, screenRuleName
	ExtendedErrorCodeSCREENRULEACTIONMISSINGFIELDREFERENCE ExtendedErrorCode = "SCREENRULEACTIONMISSINGFIELDREFERENCE"

	// Errors with this extended error code have the following properties: severity, screenRuleName
	ExtendedErrorCodeSCREENRULEACTIONMISSINGVALUE ExtendedErrorCode = "SCREENRULEACTIONMISSINGVALUE"

	// Errors with this extended error code have the following properties: severity, screenRuleName, attributeName, fieldName
	ExtendedErrorCodeSCREENRULEATTRIBUTENOTSUPPORTEDFORSCREENFIELD ExtendedErrorCode = "SCREENRULEATTRIBUTENOTSUPPORTEDFORSCREENFIELD"

	// Errors with this extended error code have the following properties: severity, screenRuleName, fieldValue
	ExtendedErrorCodeSCREENRULEFIELDNOTFOUNDONSCREEN ExtendedErrorCode = "SCREENRULEFIELDNOTFOUNDONSCREEN"

	// Errors with this extended error code have the following properties: severity, screenRuleName
	ExtendedErrorCodeSCREENRULEMISSINGACTION ExtendedErrorCode = "SCREENRULEMISSINGACTION"

	// Errors with this extended error code have the following properties: severity
	ExtendedErrorCodeSCREENRULENOTSUPPORTEDINORG ExtendedErrorCode = "SCREENRULENOTSUPPORTEDINORG"

	// Errors with this extended error code have the following properties: severity, fieldName
	ExtendedErrorCodeSCREENRULESCREENFIELDNOTVISIBLE ExtendedErrorCode = "SCREENRULESCREENFIELDNOTVISIBLE"

	// Errors with this extended error code have the following properties: severity
	ExtendedErrorCodeSCREENRULEVISIBILITYNOTSUPPORTEDINORG ExtendedErrorCode = "SCREENRULEVISIBILITYNOTSUPPORTEDINORG"

	// Errors with this extended error code have the following properties: severity
	ExtendedErrorCodeSCREENALLOWBACKALLOWFINISHBOTHFALSE ExtendedErrorCode = "SCREENALLOWBACKALLOWFINISHBOTHFALSE"

	// Errors with this extended error code have the following properties: severity, elementName
	ExtendedErrorCodeSCREENCONTAINSLIGHTNINGCOMPONENT ExtendedErrorCode = "SCREENCONTAINSLIGHTNINGCOMPONENT"

	// Errors with this extended error code have the following properties: severity
	ExtendedErrorCodeSCREENMISSINGFOOTERANDLIGHTNINGCOMPONENT ExtendedErrorCode = "SCREENMISSINGFOOTERANDLIGHTNINGCOMPONENT"

	// Errors with this extended error code have the following properties: severity, characterLimit
	ExtendedErrorCodeSCREENMISSINGLABEL ExtendedErrorCode = "SCREENMISSINGLABEL"

	// Errors with this extended error code have the following properties: severity, choiceName
	ExtendedErrorCodeSCREENMULTISELECTFIELDDOESNTSUPPORTCHOICEWITHUSERINPUT ExtendedErrorCode = "SCREENMULTISELECTFIELDDOESNTSUPPORTCHOICEWITHUSERINPUT"

	// Errors with this extended error code have the following properties: severity, fieldName
	ExtendedErrorCodeSCREENPAUSEDTEXTNOTSHOWNWHENALLOWPAUSEISFALSE ExtendedErrorCode = "SCREENPAUSEDTEXTNOTSHOWNWHENALLOWPAUSEISFALSE"

	// Errors with this extended error code have the following properties: severity, fieldName, requiredField
	ExtendedErrorCodeSETTINGFIELDMAKESOTHERFIELDREQUIRED ExtendedErrorCode = "SETTINGFIELDMAKESOTHERFIELDREQUIRED"

	// Errors with this extended error code have the following properties: severity, fieldName, otherFieldName
	ExtendedErrorCodeSETTINGFIELDMAKESOTHERFIELDUNSUPPORTED ExtendedErrorCode = "SETTINGFIELDMAKESOTHERFIELDUNSUPPORTED"

	// Errors with this extended error code have the following properties: severity, fieldName, fieldValue
	ExtendedErrorCodeSOBJECTELEMENTINCOMPATIBLEDATATYPE ExtendedErrorCode = "SOBJECTELEMENTINCOMPATIBLEDATATYPE"

	// Errors with this extended error code have the following properties: severity, objectType, sobjectName
	ExtendedErrorCodeSOBJECTELEMENTMISMATCHEDOBJECTTYPE ExtendedErrorCode = "SOBJECTELEMENTMISMATCHEDOBJECTTYPE"

	// Errors with this extended error code have the following properties: severity, fieldName, objectType
	ExtendedErrorCodeSORTENCRYPTEDFIELDSNOTSUPPORTED ExtendedErrorCode = "SORTENCRYPTEDFIELDSNOTSUPPORTED"

	// Errors with this extended error code have the following properties: severity, sortOrder
	ExtendedErrorCodeSORTFIELDMISSING ExtendedErrorCode = "SORTFIELDMISSING"

	// Errors with this extended error code have the following properties: severity, fieldName, objectName
	ExtendedErrorCodeSORTFIELDNOTSUPPORTED ExtendedErrorCode = "SORTFIELDNOTSUPPORTED"

	// Errors with this extended error code have the following properties: severity, fieldName, objectName
	ExtendedErrorCodeSORTGEOLOCATIONFIELDSNOTSUPPORTED ExtendedErrorCode = "SORTGEOLOCATIONFIELDSNOTSUPPORTED"

	// Errors with this extended error code have the following properties: severity, maxLimit
	ExtendedErrorCodeSORTLIMITINVALID ExtendedErrorCode = "SORTLIMITINVALID"

	// Errors with this extended error code have the following properties: severity, fieldName
	ExtendedErrorCodeSORTORDERMISSING ExtendedErrorCode = "SORTORDERMISSING"

	// Errors with this extended error code have the following properties: severity, fieldName, fieldType, requiedField
	ExtendedErrorCodeSPECIFICFIELDVALUEMAKESOTHERFIELDREQUIRED ExtendedErrorCode = "SPECIFICFIELDVALUEMAKESOTHERFIELDREQUIRED"

	// Errors with this extended error code have the following properties: severity
	ExtendedErrorCodeSTARTELEMENTMISSING ExtendedErrorCode = "STARTELEMENTMISSING"

	// Errors with this extended error code have the following properties: severity, flowName
	ExtendedErrorCodeSUBFLOWDESKTOPDESIGNERFLOWSNOTSUPPORTED ExtendedErrorCode = "SUBFLOWDESKTOPDESIGNERFLOWSNOTSUPPORTED"

	// Errors with this extended error code have the following properties: severity, subflowName, inputAssignmentNames
	ExtendedErrorCodeSUBFLOWINPUTELEMENTINCOMPATIBLEDATATYPES ExtendedErrorCode = "SUBFLOWINPUTELEMENTINCOMPATIBLEDATATYPES"

	// Errors with this extended error code have the following properties: severity, subflowName, inputAssignmentNames
	ExtendedErrorCodeSUBFLOWINPUTINVALIDVALUE ExtendedErrorCode = "SUBFLOWINPUTINVALIDVALUE"

	// Errors with this extended error code have the following properties: severity, subflowName, inputParameterNames
	ExtendedErrorCodeSUBFLOWINPUTMISMATCHEDCOLLECTIONTYPES ExtendedErrorCode = "SUBFLOWINPUTMISMATCHEDCOLLECTIONTYPES"

	// Errors with this extended error code have the following properties: severity, subflowName, inputParameterNames
	ExtendedErrorCodeSUBFLOWINPUTMISMATCHEDOBJECTS ExtendedErrorCode = "SUBFLOWINPUTMISMATCHEDOBJECTS"

	// Errors with this extended error code have the following properties: severity, subflowName
	ExtendedErrorCodeSUBFLOWINPUTMISSINGNAME ExtendedErrorCode = "SUBFLOWINPUTMISSINGNAME"

	// Errors with this extended error code have the following properties: severity, inputVariableName
	ExtendedErrorCodeSUBFLOWINPUTMULTIPLEASSIGNMENTSTOONEVARIABLE ExtendedErrorCode = "SUBFLOWINPUTMULTIPLEASSIGNMENTSTOONEVARIABLE"

	// Errors with this extended error code have the following properties: severity, inputVariableName
	ExtendedErrorCodeSUBFLOWINPUTREFERENCESFIELDONSOBJECTVARIABLE ExtendedErrorCode = "SUBFLOWINPUTREFERENCESFIELDONSOBJECTVARIABLE"

	// Errors with this extended error code have the following properties: severity, subflowName, inputAssignmentNames
	ExtendedErrorCodeSUBFLOWINPUTVALUEINCOMPATIBLEDATATYPES ExtendedErrorCode = "SUBFLOWINPUTVALUEINCOMPATIBLEDATATYPES"

	// Errors with this extended error code have the following properties: severity, subflowName, inputAssignmentNames
	ExtendedErrorCodeSUBFLOWINPUTVARIABLENOTFOUNDINMASTERFLOW ExtendedErrorCode = "SUBFLOWINPUTVARIABLENOTFOUNDINMASTERFLOW"

	// Errors with this extended error code have the following properties: severity, subflowName, inputAssignmentNames
	ExtendedErrorCodeSUBFLOWINPUTVARIABLENOTFOUNDINREFERENCEDFLOW ExtendedErrorCode = "SUBFLOWINPUTVARIABLENOTFOUNDINREFERENCEDFLOW"

	// Errors with this extended error code have the following properties: severity, subflowName, inputAssignmentNames
	ExtendedErrorCodeSUBFLOWINPUTVARIABLENOINPUTACCESS ExtendedErrorCode = "SUBFLOWINPUTVARIABLENOINPUTACCESS"

	// Errors with this extended error code have the following properties: severity
	ExtendedErrorCodeSUBFLOWINVALIDNAME ExtendedErrorCode = "SUBFLOWINVALIDNAME"

	// Errors with this extended error code have the following properties: severity, flowName
	ExtendedErrorCodeSUBFLOWINVALIDREFERENCE ExtendedErrorCode = "SUBFLOWINVALIDREFERENCE"

	// Errors with this extended error code have the following properties: severity, parentFlowName
	ExtendedErrorCodeSUBFLOWMASTERFLOWTYPENOTAUTOLAUNCHED ExtendedErrorCode = "SUBFLOWMASTERFLOWTYPENOTAUTOLAUNCHED"

	// Errors with this extended error code have the following properties: severity
	ExtendedErrorCodeSUBFLOWMISSINGNAME ExtendedErrorCode = "SUBFLOWMISSINGNAME"

	// Errors with this extended error code have the following properties: severity, subflowName, flowName
	ExtendedErrorCodeSUBFLOWNOACTIVEVERSION ExtendedErrorCode = "SUBFLOWNOACTIVEVERSION"

	// Errors with this extended error code have the following properties: severity, subflowName, flowVersion, outputParameterNames
	ExtendedErrorCodeSUBFLOWOUTPUTINCOMPATIBLEDATATYPES ExtendedErrorCode = "SUBFLOWOUTPUTINCOMPATIBLEDATATYPES"

	// Errors with this extended error code have the following properties: severity, subflowName, flowVersion, outputParameterNames
	ExtendedErrorCodeSUBFLOWOUTPUTMISMATCHEDCOLLECTIONTYPES ExtendedErrorCode = "SUBFLOWOUTPUTMISMATCHEDCOLLECTIONTYPES"

	// Errors with this extended error code have the following properties: severity, subflowName, flowVersion, outputParameterNames
	ExtendedErrorCodeSUBFLOWOUTPUTMISMATCHEDOBJECTS ExtendedErrorCode = "SUBFLOWOUTPUTMISMATCHEDOBJECTS"

	// Errors with this extended error code have the following properties: severity, outputAssignment
	ExtendedErrorCodeSUBFLOWOUTPUTMISSINGASSIGNTOREFERENCE ExtendedErrorCode = "SUBFLOWOUTPUTMISSINGASSIGNTOREFERENCE"

	// Errors with this extended error code have the following properties: severity, subflowName
	ExtendedErrorCodeSUBFLOWOUTPUTMISSINGNAME ExtendedErrorCode = "SUBFLOWOUTPUTMISSINGNAME"

	// Errors with this extended error code have the following properties: severity, outputVariableName
	ExtendedErrorCodeSUBFLOWOUTPUTMULTIPLEASSIGNMENTSTOONEVARIABLE ExtendedErrorCode = "SUBFLOWOUTPUTMULTIPLEASSIGNMENTSTOONEVARIABLE"

	// Errors with this extended error code have the following properties: severity, outputAssignment
	ExtendedErrorCodeSUBFLOWOUTPUTREFERENCESFIELDONSOBJECTVARIABLE ExtendedErrorCode = "SUBFLOWOUTPUTREFERENCESFIELDONSOBJECTVARIABLE"

	// Errors with this extended error code have the following properties: severity, subflowName, outputAssignmentName
	ExtendedErrorCodeSUBFLOWOUTPUTTARGETDOESNOTEXISTINMASTERFLOW ExtendedErrorCode = "SUBFLOWOUTPUTTARGETDOESNOTEXISTINMASTERFLOW"

	// Errors with this extended error code have the following properties: severity, subflowName, variableName
	ExtendedErrorCodeSUBFLOWOUTPUTVARIABLENOTFOUNDINMASTERFLOW ExtendedErrorCode = "SUBFLOWOUTPUTVARIABLENOTFOUNDINMASTERFLOW"

	// Errors with this extended error code have the following properties: severity, subflowName, flowVersion, outputParameterNames
	ExtendedErrorCodeSUBFLOWOUTPUTVARIABLENOTFOUNDINREFERENCEDFLOW ExtendedErrorCode = "SUBFLOWOUTPUTVARIABLENOTFOUNDINREFERENCEDFLOW"

	// Errors with this extended error code have the following properties: severity, subflowName, variableName
	ExtendedErrorCodeSUBFLOWOUTPUTVARIABLENOOUTPUTACCESS ExtendedErrorCode = "SUBFLOWOUTPUTVARIABLENOOUTPUTACCESS"

	// Errors with this extended error code have the following properties: severity, subflowName, flowName, processType
	ExtendedErrorCodeSUBFLOWPROCESSTYPEINCOMPATIBLE ExtendedErrorCode = "SUBFLOWPROCESSTYPEINCOMPATIBLE"

	// Errors with this extended error code have the following properties: severity
	ExtendedErrorCodeSUBFLOWREFERENCESMASTERFLOW ExtendedErrorCode = "SUBFLOWREFERENCESMASTERFLOW"

	// Errors with this extended error code have the following properties: severity, elementName
	ExtendedErrorCodeSURVEYADVANCEDCONDITIONLOGICNOTSUPPORTED ExtendedErrorCode = "SURVEYADVANCEDCONDITIONLOGICNOTSUPPORTED"

	// Errors with this extended error code have the following properties: severity, choiceName
	ExtendedErrorCodeSURVEYCHOICENOTREFERENCEDBYAQUESTION ExtendedErrorCode = "SURVEYCHOICENOTREFERENCEDBYAQUESTION"

	// Errors with this extended error code have the following properties: severity, choiceName
	ExtendedErrorCodeSURVEYCHOICEREFERENCEDBYMULTIPLEQUESTIONS ExtendedErrorCode = "SURVEYCHOICEREFERENCEDBYMULTIPLEQUESTIONS"

	// Errors with this extended error code have the following properties: severity, elementName
	ExtendedErrorCodeSURVEYELEMENTNEVERREACHED ExtendedErrorCode = "SURVEYELEMENTNEVERREACHED"

	// Errors with this extended error code have the following properties: severity, subflowName
	ExtendedErrorCodeSURVEYINACTIVESUBFLOWS ExtendedErrorCode = "SURVEYINACTIVESUBFLOWS"

	// Errors with this extended error code have the following properties: severity, surveyName
	ExtendedErrorCodeSURVEYMISSINGQUESTIONORSUBFLOW ExtendedErrorCode = "SURVEYMISSINGQUESTIONORSUBFLOW"

	// Errors with this extended error code have the following properties: severity, surveyName
	ExtendedErrorCodeSURVEYMISSINGREQUIREDVARIABLES ExtendedErrorCode = "SURVEYMISSINGREQUIREDVARIABLES"

	// Errors with this extended error code have the following properties: severity, flowDecision
	ExtendedErrorCodeSURVEYMULTIPLESCREENSCANNOTCONNECTTOSAMEDECISION ExtendedErrorCode = "SURVEYMULTIPLESCREENSCANNOTCONNECTTOSAMEDECISION"

	// Errors with this extended error code have the following properties: severity, subflowName
	ExtendedErrorCodeSURVEYNESTEDSUBFLOWS ExtendedErrorCode = "SURVEYNESTEDSUBFLOWS"

	// Errors with this extended error code have the following properties: severity, subflowName
	ExtendedErrorCodeSURVEYNONSURVEYSUBFLOWS ExtendedErrorCode = "SURVEYNONSURVEYSUBFLOWS"

	// Errors with this extended error code have the following properties: severity, elementName
	ExtendedErrorCodeSURVEYRULEINVALIDRIGHTOPERAND ExtendedErrorCode = "SURVEYRULEINVALIDRIGHTOPERAND"

	// Errors with this extended error code have the following properties: severity, elementName
	ExtendedErrorCodeSURVEYSCREENFIELDTYPENOTSUPPORTEDFORQUESTION ExtendedErrorCode = "SURVEYSCREENFIELDTYPENOTSUPPORTEDFORQUESTION"

	// Errors with this extended error code have the following properties: severity
	ExtendedErrorCodeSURVEYSTARTELEMENTINVALID ExtendedErrorCode = "SURVEYSTARTELEMENTINVALID"

	// Errors with this extended error code have the following properties: severity, surveyName
	ExtendedErrorCodeSURVEYVARIABLEACCESSINVALID ExtendedErrorCode = "SURVEYVARIABLEACCESSINVALID"

	// Errors with this extended error code have the following properties: severity
	ExtendedErrorCodeUNEXPECTEDERROR ExtendedErrorCode = "UNEXPECTEDERROR"

	// Errors with this extended error code have the following properties: severity, elementName, characterLimit
	ExtendedErrorCodeVALUECHARLIMITEXCEEDED ExtendedErrorCode = "VALUECHARLIMITEXCEEDED"

	// Errors with this extended error code have the following properties: severity, fieldName, datatype
	ExtendedErrorCodeVARIABLEFIELDNOTSUPPORTEDFORDATATYPE ExtendedErrorCode = "VARIABLEFIELDNOTSUPPORTEDFORDATATYPE"

	// Errors with this extended error code have the following properties: severity, fieldName, datatype
	ExtendedErrorCodeVARIABLEFIELDNOTSUPPORTEDFORDATATYPEANDCOLLECTION ExtendedErrorCode = "VARIABLEFIELDNOTSUPPORTEDFORDATATYPEANDCOLLECTION"

	// Errors with this extended error code have the following properties: severity, datatype, fieldName
	ExtendedErrorCodeVARIABLEFIELDREQUIREDFORDATATYPE ExtendedErrorCode = "VARIABLEFIELDREQUIREDFORDATATYPE"

	// Errors with this extended error code have the following properties: severity, elementName
	ExtendedErrorCodeVARIABLESCALEEXCEEDSLIMIT ExtendedErrorCode = "VARIABLESCALEEXCEEDSLIMIT"

	// Errors with this extended error code have the following properties: severity, elementName
	ExtendedErrorCodeVARIABLESCALENEGATIVEINTEGER ExtendedErrorCode = "VARIABLESCALENEGATIVEINTEGER"

	// Errors with this extended error code have the following properties: severity, elementName
	ExtendedErrorCodeVARIABLESCALENULL ExtendedErrorCode = "VARIABLESCALENULL"

	// Errors with this extended error code have the following properties: severity, waitEventName
	ExtendedErrorCodeWAITEVENTDEFAULTCONNECTORMISSINGLABEL ExtendedErrorCode = "WAITEVENTDEFAULTCONNECTORMISSINGLABEL"

	// Errors with this extended error code have the following properties: severity, parameterName
	ExtendedErrorCodeWAITEVENTDUPLICATEINPUTPARAM ExtendedErrorCode = "WAITEVENTDUPLICATEINPUTPARAM"

	// Errors with this extended error code have the following properties: severity, waitEventName, inputParameterName
	ExtendedErrorCodeWAITEVENTINPUTNOTSUPPORTEDFOREVENTTYPE ExtendedErrorCode = "WAITEVENTINPUTNOTSUPPORTEDFOREVENTTYPE"

	// Errors with this extended error code have the following properties: severity, waitEventName, parameterName
	ExtendedErrorCodeWAITEVENTINPUTREQUIRESLITERALVALUE ExtendedErrorCode = "WAITEVENTINPUTREQUIRESLITERALVALUE"

	// Errors with this extended error code have the following properties: severity, waitEventName
	ExtendedErrorCodeWAITEVENTINVALIDCONDITIONLOGIC ExtendedErrorCode = "WAITEVENTINVALIDCONDITIONLOGIC"

	// Errors with this extended error code have the following properties: severity
	ExtendedErrorCodeWAITEVENTMISSING ExtendedErrorCode = "WAITEVENTMISSING"

	// Errors with this extended error code have the following properties: severity, waitEventName
	ExtendedErrorCodeWAITEVENTMISSINGCONNECTOR ExtendedErrorCode = "WAITEVENTMISSINGCONNECTOR"

	// Errors with this extended error code have the following properties: severity, waitEventName
	ExtendedErrorCodeWAITEVENTMISSINGEVENTTYPE ExtendedErrorCode = "WAITEVENTMISSINGEVENTTYPE"

	// Errors with this extended error code have the following properties: severity, waitEventName
	ExtendedErrorCodeWAITEVENTOBJECTNOTSUPPORTEDFOREVENTTYPE ExtendedErrorCode = "WAITEVENTOBJECTNOTSUPPORTEDFOREVENTTYPE"

	// Errors with this extended error code have the following properties: severity, waitEventName, outputParameter
	ExtendedErrorCodeWAITEVENTOUTPUTNOTSUPPORTEDFOREVENTTYPE ExtendedErrorCode = "WAITEVENTOUTPUTNOTSUPPORTEDFOREVENTTYPE"

	// Errors with this extended error code have the following properties: severity, waitEventName, eventParameterName, incompatibleValue
	ExtendedErrorCodeWAITEVENTRELATIVEALARMINVALIDDATETIMEFIELD ExtendedErrorCode = "WAITEVENTRELATIVEALARMINVALIDDATETIMEFIELD"

	// Errors with this extended error code have the following properties: severity, waitEventName, eventParameterName, incompatibleValue
	ExtendedErrorCodeWAITEVENTRELATIVEALARMINVALIDFIELD ExtendedErrorCode = "WAITEVENTRELATIVEALARMINVALIDFIELD"

	// Errors with this extended error code have the following properties: severity, waitEventName, inputParameterName
	ExtendedErrorCodeWAITEVENTRELATIVEALARMINVALIDOBJECTTYPE ExtendedErrorCode = "WAITEVENTRELATIVEALARMINVALIDOBJECTTYPE"

	// Errors with this extended error code have the following properties: severity, waitEventName, eventParameterName, incompatibleValue
	ExtendedErrorCodeWAITEVENTRELATIVEALARMINVALIDOFFSETNUMBER ExtendedErrorCode = "WAITEVENTRELATIVEALARMINVALIDOFFSETNUMBER"

	// Errors with this extended error code have the following properties: severity, waitEventName, eventParameterName, incompatibleValue
	ExtendedErrorCodeWAITEVENTRELATIVEALARMINVALIDOFFSETUNIT ExtendedErrorCode = "WAITEVENTRELATIVEALARMINVALIDOFFSETUNIT"

	// Errors with this extended error code have the following properties: severity, waitEventName, parameterName
	ExtendedErrorCodeWAITEVENTREQUIREDINPUTMISSING ExtendedErrorCode = "WAITEVENTREQUIREDINPUTMISSING"

	// Errors with this extended error code have the following properties: severity, waitEventName
	ExtendedErrorCodeWAITEVENTTYPEINVALIDORNOTSUPPORTED ExtendedErrorCode = "WAITEVENTTYPEINVALIDORNOTSUPPORTED"

	// Errors with this extended error code have the following properties: severity, flowName
	ExtendedErrorCodeWORKFLOWMISSINGPROCESSMETADATAVALUES ExtendedErrorCode = "WORKFLOWMISSINGPROCESSMETADATAVALUES"

	// Errors with this extended error code have the following properties: severity, objectType
	ExtendedErrorCodeWORKFLOWOBJECTTYPENOTFOUND ExtendedErrorCode = "WORKFLOWOBJECTTYPENOTFOUND"

	// Errors with this extended error code have the following properties: severity, objectType
	ExtendedErrorCodeWORKFLOWOBJECTTYPENOTSUPPORTED ExtendedErrorCode = "WORKFLOWOBJECTTYPENOTSUPPORTED"

	// Errors with this extended error code have the following properties: severity, objectVariableName, oldObjectVariableName
	ExtendedErrorCodeWORKFLOWOBJECTVARIABLEANDOLDOBJECTVARIABLEREFERENCESAMESOBJECTVARIABLE ExtendedErrorCode = "WORKFLOWOBJECTVARIABLEANDOLDOBJECTVARIABLEREFERENCESAMESOBJECTVARIABLE"

	// Errors with this extended error code have the following properties: severity, objectType, objectVariableName
	ExtendedErrorCodeWORKFLOWOBJECTVARIABLEDOESNTSUPPORTINPUT ExtendedErrorCode = "WORKFLOWOBJECTVARIABLEDOESNTSUPPORTINPUT"

	// Errors with this extended error code have the following properties: severity, objectType, oldObjectVariableName
	ExtendedErrorCodeWORKFLOWOLDOBJECTVARIABLEDOESNTSUPPORTINPUT ExtendedErrorCode = "WORKFLOWOLDOBJECTVARIABLEDOESNTSUPPORTINPUT"

	// Errors with this extended error code have the following properties: severity, metadataValue
	ExtendedErrorCodeWORKFLOWPROCESSMETADATAVALUESMORETHANONENAME ExtendedErrorCode = "WORKFLOWPROCESSMETADATAVALUESMORETHANONENAME"

	// Errors with this extended error code have the following properties: severity, metadataValue
	ExtendedErrorCodeWORKFLOWPROCESSMETADATAVALUESMISSINGNAME ExtendedErrorCode = "WORKFLOWPROCESSMETADATAVALUESMISSINGNAME"

	// Errors with this extended error code have the following properties: severity, elementName
	ExtendedErrorCodeWORKFLOWRECURSIVECOUNTVARIABLEDOESNTSUPPORTINPUT ExtendedErrorCode = "WORKFLOWRECURSIVECOUNTVARIABLEDOESNTSUPPORTINPUT"

	// Errors with this extended error code have the following properties: severity
	ExtendedErrorCodeWORKFLOWTRIGGERTYPEINVALIDVALUE ExtendedErrorCode = "WORKFLOWTRIGGERTYPEINVALIDVALUE"
)

type ShareAccessLevel string

const (
	ShareAccessLevelRead ShareAccessLevel = "Read"

	ShareAccessLevelEdit ShareAccessLevel = "Edit"

	ShareAccessLevelAll ShareAccessLevel = "All"
)

type FieldType string

const (
	FieldTypeString FieldType = "string"

	FieldTypePicklist FieldType = "picklist"

	FieldTypeMultipicklist FieldType = "multipicklist"

	FieldTypeCombobox FieldType = "combobox"

	FieldTypeReference FieldType = "reference"

	FieldTypeBase64 FieldType = "base64"

	FieldTypeBoolean FieldType = "boolean"

	FieldTypeCurrency FieldType = "currency"

	FieldTypeTextarea FieldType = "textarea"

	FieldTypeInt FieldType = "int"

	FieldTypeDouble FieldType = "double"

	FieldTypePercent FieldType = "percent"

	FieldTypePhone FieldType = "phone"

	FieldTypeId FieldType = "id"

	FieldTypeDate FieldType = "date"

	FieldTypeDatetime FieldType = "datetime"

	FieldTypeTime FieldType = "time"

	FieldTypeUrl FieldType = "url"

	FieldTypeEmail FieldType = "email"

	FieldTypeEncryptedstring FieldType = "encryptedstring"

	FieldTypeDatacategorygroupreference FieldType = "datacategorygroupreference"

	FieldTypeLocation FieldType = "location"

	FieldTypeAddress FieldType = "address"

	FieldTypeAnyType FieldType = "anyType"

	FieldTypeComplexvalue FieldType = "complexvalue"

	FieldTypeLong FieldType = "long"
)

type SoapType string

const (
	SoapTypeTnsID SoapType = "tnsID"

	SoapTypeXsdbase64Binary SoapType = "xsdbase64Binary"

	SoapTypeXsdboolean SoapType = "xsdboolean"

	SoapTypeXsddouble SoapType = "xsddouble"

	SoapTypeXsdint SoapType = "xsdint"

	SoapTypeXsdstring SoapType = "xsdstring"

	SoapTypeXsddate SoapType = "xsddate"

	SoapTypeXsddateTime SoapType = "xsddateTime"

	SoapTypeXsdtime SoapType = "xsdtime"

	SoapTypeTnslocation SoapType = "tnslocation"

	SoapTypeTnsaddress SoapType = "tnsaddress"

	SoapTypeXsdanyType SoapType = "xsdanyType"

	SoapTypeUrnRelationshipReferenceTo SoapType = "urnRelationshipReferenceTo"

	SoapTypeUrnJunctionIdListNames SoapType = "urnJunctionIdListNames"

	SoapTypeUrnSearchLayoutFieldsDisplayed SoapType = "urnSearchLayoutFieldsDisplayed"

	SoapTypeUrnSearchLayoutField SoapType = "urnSearchLayoutField"

	SoapTypeUrnSearchLayoutButtonsDisplayed SoapType = "urnSearchLayoutButtonsDisplayed"

	SoapTypeUrnSearchLayoutButton SoapType = "urnSearchLayoutButton"

	SoapTypeUrnRecordTypesSupported SoapType = "urnRecordTypesSupported"

	SoapTypeTnsStringList SoapType = "tnsStringList"
)

type DifferenceType string

const (
	DifferenceTypeDIFFERENT DifferenceType = "DIFFERENT"

	DifferenceTypeNULL DifferenceType = "NULL"

	DifferenceTypeSAME DifferenceType = "SAME"

	DifferenceTypeSIMILAR DifferenceType = "SIMILAR"
)

type Article string

const (
	ArticleNone Article = "None"

	ArticleIndefinite Article = "Indefinite"

	ArticleDefinite Article = "Definite"
)

type CaseType string

const (
	CaseTypeNominative CaseType = "Nominative"

	CaseTypeAccusative CaseType = "Accusative"

	CaseTypeGenitive CaseType = "Genitive"

	CaseTypeDative CaseType = "Dative"

	CaseTypeInessive CaseType = "Inessive"

	CaseTypeElative CaseType = "Elative"

	CaseTypeIllative CaseType = "Illative"

	CaseTypeAdessive CaseType = "Adessive"

	CaseTypeAblative CaseType = "Ablative"

	CaseTypeAllative CaseType = "Allative"

	CaseTypeEssive CaseType = "Essive"

	CaseTypeTranslative CaseType = "Translative"

	CaseTypePartitive CaseType = "Partitive"

	CaseTypeObjective CaseType = "Objective"

	CaseTypeSubjective CaseType = "Subjective"

	CaseTypeInstrumental CaseType = "Instrumental"

	CaseTypePrepositional CaseType = "Prepositional"

	CaseTypeLocative CaseType = "Locative"

	CaseTypeVocative CaseType = "Vocative"

	CaseTypeSublative CaseType = "Sublative"

	CaseTypeSuperessive CaseType = "Superessive"

	CaseTypeDelative CaseType = "Delative"

	CaseTypeCausalfinal CaseType = "Causalfinal"

	CaseTypeEssiveformal CaseType = "Essiveformal"

	CaseTypeTermanative CaseType = "Termanative"

	CaseTypeDistributive CaseType = "Distributive"

	CaseTypeErgative CaseType = "Ergative"

	CaseTypeAdverbial CaseType = "Adverbial"

	CaseTypeAbessive CaseType = "Abessive"

	CaseTypeComitative CaseType = "Comitative"
)

type Gender string

const (
	GenderNeuter Gender = "Neuter"

	GenderMasculine Gender = "Masculine"

	GenderFeminine Gender = "Feminine"

	GenderAnimateMasculine Gender = "AnimateMasculine"
)

type GrammaticalNumber string

const (
	GrammaticalNumberSingular GrammaticalNumber = "Singular"

	GrammaticalNumberPlural GrammaticalNumber = "Plural"
)

type Possessive string

const (
	PossessiveNone Possessive = "None"

	PossessiveFirst Possessive = "First"

	PossessiveSecond Possessive = "Second"
)

type StartsWith string

const (
	StartsWithConsonant StartsWith = "Consonant"

	StartsWithVowel StartsWith = "Vowel"

	StartsWithSpecial StartsWith = "Special"
)

type FeedLayoutFilterType string

const (
	FeedLayoutFilterTypeAllUpdates FeedLayoutFilterType = "AllUpdates"

	FeedLayoutFilterTypeFeedItemType FeedLayoutFilterType = "FeedItemType"

	FeedLayoutFilterTypeCustom FeedLayoutFilterType = "Custom"
)

type TabOrderType string

const (
	TabOrderTypeLeftToRight TabOrderType = "LeftToRight"

	TabOrderTypeTopToBottom TabOrderType = "TopToBottom"
)

type WebLinkWindowType string

const (
	WebLinkWindowTypeNewWindow WebLinkWindowType = "newWindow"

	WebLinkWindowTypeSidebar WebLinkWindowType = "sidebar"

	WebLinkWindowTypeNoSidebar WebLinkWindowType = "noSidebar"

	WebLinkWindowTypeReplace WebLinkWindowType = "replace"

	WebLinkWindowTypeOnClickJavaScript WebLinkWindowType = "onClickJavaScript"
)

type WebLinkPosition string

const (
	WebLinkPositionFullScreen WebLinkPosition = "fullScreen"

	WebLinkPositionNone WebLinkPosition = "none"

	WebLinkPositionTopLeft WebLinkPosition = "topLeft"
)

type WebLinkType string

const (
	WebLinkTypeUrl WebLinkType = "url"

	WebLinkTypeSControl WebLinkType = "sControl"

	WebLinkTypeJavascript WebLinkType = "javascript"

	WebLinkTypePage WebLinkType = "page"

	WebLinkTypeFlow WebLinkType = "flow"
)

type LayoutComponentType string

const (
	LayoutComponentTypeReportChart LayoutComponentType = "ReportChart"

	LayoutComponentTypeField LayoutComponentType = "Field"

	LayoutComponentTypeSeparator LayoutComponentType = "Separator"

	LayoutComponentTypeSControl LayoutComponentType = "SControl"

	LayoutComponentTypeEmptySpace LayoutComponentType = "EmptySpace"

	LayoutComponentTypeVisualforcePage LayoutComponentType = "VisualforcePage"

	LayoutComponentTypeExpandedLookup LayoutComponentType = "ExpandedLookup"

	LayoutComponentTypeAuraComponent LayoutComponentType = "AuraComponent"

	LayoutComponentTypeCanvas LayoutComponentType = "Canvas"

	LayoutComponentTypeCustomLink LayoutComponentType = "CustomLink"

	LayoutComponentTypeAnalyticsCloud LayoutComponentType = "AnalyticsCloud"
)

type ReportChartSize string

const (
	ReportChartSizeSMALL ReportChartSize = "SMALL"

	ReportChartSizeMEDIUM ReportChartSize = "MEDIUM"

	ReportChartSizeLARGE ReportChartSize = "LARGE"
)

type EmailPriority string

const (
	EmailPriorityHighest EmailPriority = "Highest"

	EmailPriorityHigh EmailPriority = "High"

	EmailPriorityNormal EmailPriority = "Normal"

	EmailPriorityLow EmailPriority = "Low"

	EmailPriorityLowest EmailPriority = "Lowest"
)

type SendEmailOptOutPolicy string

const (
	SendEmailOptOutPolicySEND SendEmailOptOutPolicy = "SEND"

	SendEmailOptOutPolicyFILTER SendEmailOptOutPolicy = "FILTER"

	SendEmailOptOutPolicyREJECT SendEmailOptOutPolicy = "REJECT"
)

type AttachmentRetrievalOption string

const (
	AttachmentRetrievalOptionNone AttachmentRetrievalOption = "None"

	AttachmentRetrievalOptionMetadataOnly AttachmentRetrievalOption = "MetadataOnly"

	AttachmentRetrievalOptionMetadataWithBody AttachmentRetrievalOption = "MetadataWithBody"
)

type OrderByDirection string

const (
	OrderByDirectionAscending OrderByDirection = "ascending"

	OrderByDirectionDescending OrderByDirection = "descending"
)

type OrderByNullsPosition string

const (
	OrderByNullsPositionFirst OrderByNullsPosition = "first"

	OrderByNullsPositionLast OrderByNullsPosition = "last"
)

type SoqlOperator string

const (
	SoqlOperatorEquals SoqlOperator = "equals"

	SoqlOperatorExcludes SoqlOperator = "excludes"

	SoqlOperatorGreaterThan SoqlOperator = "greaterThan"

	SoqlOperatorGreaterThanOrEqualTo SoqlOperator = "greaterThanOrEqualTo"

	SoqlOperatorIn SoqlOperator = "in"

	SoqlOperatorIncludes SoqlOperator = "includes"

	SoqlOperatorLessThan SoqlOperator = "lessThan"

	SoqlOperatorLessThanOrEqualTo SoqlOperator = "lessThanOrEqualTo"

	SoqlOperatorLike SoqlOperator = "like"

	SoqlOperatorNotEquals SoqlOperator = "notEquals"

	SoqlOperatorNotIn SoqlOperator = "notIn"

	SoqlOperatorWithin SoqlOperator = "within"

	SoqlOperatorNotLike SoqlOperator = "notLike"
)

type SoqlConjunction string

const (
	SoqlConjunctionAnd SoqlConjunction = "and"

	SoqlConjunctionOr SoqlConjunction = "or"
)

type AppMenuType string

const (
	AppMenuTypeAppSwitcher AppMenuType = "AppSwitcher"

	AppMenuTypeSalesforce1 AppMenuType = "Salesforce1"

	AppMenuTypeNetworkTabs AppMenuType = "NetworkTabs"
)

type ListViewIsSoqlCompatible string

const (
	ListViewIsSoqlCompatibleTRUE ListViewIsSoqlCompatible = "TRUE"

	ListViewIsSoqlCompatibleFALSE ListViewIsSoqlCompatible = "FALSE"

	ListViewIsSoqlCompatibleALL ListViewIsSoqlCompatible = "ALL"
)

type DebugLevel string

const (
	DebugLevelNone DebugLevel = "None"

	DebugLevelDebugOnly DebugLevel = "DebugOnly"

	DebugLevelDb DebugLevel = "Db"

	DebugLevelProfiling DebugLevel = "Profiling"

	DebugLevelCallout DebugLevel = "Callout"

	DebugLevelDetail DebugLevel = "Detail"
)

type LogCategory string

const (
	LogCategoryDb LogCategory = "Db"

	LogCategoryWorkflow LogCategory = "Workflow"

	LogCategoryValidation LogCategory = "Validation"

	LogCategoryCallout LogCategory = "Callout"

	LogCategoryApexcode LogCategory = "Apexcode"

	LogCategoryApexprofiling LogCategory = "Apexprofiling"

	LogCategoryVisualforce LogCategory = "Visualforce"

	LogCategorySystem LogCategory = "System"

	LogCategoryAll LogCategory = "All"
)

type LogCategoryLevel string

const (
	LogCategoryLevelNone LogCategoryLevel = "None"

	LogCategoryLevelFinest LogCategoryLevel = "Finest"

	LogCategoryLevelFiner LogCategoryLevel = "Finer"

	LogCategoryLevelFine LogCategoryLevel = "Fine"

	LogCategoryLevelDebug LogCategoryLevel = "Debug"

	LogCategoryLevelInfo LogCategoryLevel = "Info"

	LogCategoryLevelWarn LogCategoryLevel = "Warn"

	LogCategoryLevelError LogCategoryLevel = "Error"
)

type OwnerChangeOptionType string

const (
	OwnerChangeOptionTypeEnforceNewOwnerHasReadAccess OwnerChangeOptionType = "EnforceNewOwnerHasReadAccess"

	OwnerChangeOptionTypeTransferOpenActivities OwnerChangeOptionType = "TransferOpenActivities"

	OwnerChangeOptionTypeTransferNotesAndAttachments OwnerChangeOptionType = "TransferNotesAndAttachments"

	OwnerChangeOptionTypeTransferOthersOpenOpportunities OwnerChangeOptionType = "TransferOthersOpenOpportunities"

	OwnerChangeOptionTypeTransferOwnedOpenOpportunities OwnerChangeOptionType = "TransferOwnedOpenOpportunities"

	OwnerChangeOptionTypeTransferContracts OwnerChangeOptionType = "TransferContracts"

	OwnerChangeOptionTypeTransferOrders OwnerChangeOptionType = "TransferOrders"

	OwnerChangeOptionTypeTransferContacts OwnerChangeOptionType = "TransferContacts"

	OwnerChangeOptionTypeKeepSalesTeam OwnerChangeOptionType = "KeepSalesTeam"

	OwnerChangeOptionTypeKeepSalesTeamGrantCurrentOwnerReadWriteAccess OwnerChangeOptionType = "KeepSalesTeamGrantCurrentOwnerReadWriteAccess"
)

type FindDuplicates struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com findDuplicates"`

	SObjects []*SObject `xml:"sObjects,omitempty"`
}

type FindDuplicatesByIds struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com findDuplicatesByIds"`

	Ids []*ID `xml:"ids,omitempty"`
}

type FindDuplicatesByIdsResponse struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com findDuplicatesByIdsResponse"`

	Result []*FindDuplicatesResult `xml:"result,omitempty"`
}

type FindDuplicatesResponse struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com findDuplicatesResponse"`

	Result []*FindDuplicatesResult `xml:"result,omitempty"`
}

type Login struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com login"`

	Username string `xml:"username,omitempty"`
	Password string `xml:"password,omitempty"`
}

type LoginResponse struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com loginResponse"`

	Result *LoginResult `xml:"result,omitempty"`
}

type DescribeSObject struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com describeSObject"`

	SObjectType string `xml:"sObjectType,omitempty"`
}

type DescribeSObjectResponse struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com describeSObjectResponse"`

	Result *DescribeSObjectResult `xml:"result,omitempty"`
}

type DescribeSObjects struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com describeSObjects"`

	SObjectType string `xml:"sObjectType,omitempty"`
}

type DescribeSObjectsResponse struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com describeSObjectsResponse"`

	Result *DescribeSObjectResult `xml:"result,omitempty"`
}

type DescribeGlobal struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com describeGlobal"`
}

type DescribeGlobalResponse struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com describeGlobalResponse"`

	Result *DescribeGlobalResult `xml:"result,omitempty"`
}

type DescribeGlobalTheme struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com describeGlobalTheme"`
}

type DescribeGlobalThemeResponse struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com describeGlobalThemeResponse"`

	Result *DescribeGlobalTheme `xml:"result,omitempty"`
}

type DescribeTheme struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com describeTheme"`

	SobjectType []string `xml:"sobjectType,omitempty"`
}

type DescribeThemeResponse struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com describeThemeResponse"`

	Result *DescribeThemeResult `xml:"result,omitempty"`
}

type DescribeDataCategoryGroups struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com describeDataCategoryGroups"`

	SObjectType string `xml:"sObjectType,omitempty"`
}

type DescribeDataCategoryGroupsResponse struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com describeDataCategoryGroupsResponse"`

	Result *DescribeDataCategoryGroupResult `xml:"result,omitempty"`
}

type DescribeDataCategoryGroupStructures struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com describeDataCategoryGroupStructures"`

	Pairs             *DataCategoryGroupSobjectTypePair `xml:"pairs,omitempty"`
	TopCategoriesOnly bool                              `xml:"topCategoriesOnly,omitempty"`
}

type DescribeDataCategoryGroupStructuresResponse struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com describeDataCategoryGroupStructuresResponse"`

	Result *DescribeDataCategoryGroupStructureResult `xml:"result,omitempty"`
}

type DescribeDataCategoryMappings struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com describeDataCategoryMappings"`
}

type DescribeDataCategoryMappingsResponse struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com describeDataCategoryMappingsResponse"`

	Result *DescribeDataCategoryMappingResult `xml:"result,omitempty"`
}

type DescribeKnowledgeSettings struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com describeKnowledgeSettings"`
}

type DescribeKnowledgeSettingsResponse struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com describeKnowledgeSettingsResponse"`

	Result *KnowledgeSettings `xml:"result,omitempty"`
}

type DescribeAppMenu struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com describeAppMenu"`

	AppMenuType *AppMenuType `xml:"appMenuType,omitempty"`
	NetworkId   *ID          `xml:"networkId,omitempty"`
}

type DescribeAppMenuResponse struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com describeAppMenuResponse"`

	Result *DescribeAppMenuResult `xml:"result,omitempty"`
}

// type DescribeLayout struct {
// 	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com describeLayout"`

// 	SObjectType   string `xml:"sObjectType,omitempty"`
// 	LayoutName    string `xml:"layoutName,omitempty"`
// 	RecordTypeIds []*ID  `xml:"recordTypeIds,omitempty"`
// }

type DescribeLayoutResponse struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com describeLayoutResponse"`

	Result *DescribeLayoutResult `xml:"result,omitempty"`
}

type DescribeCompactLayouts struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com describeCompactLayouts"`

	SObjectType   string `xml:"sObjectType,omitempty"`
	RecordTypeIds []*ID  `xml:"recordTypeIds,omitempty"`
}

type DescribeCompactLayoutsResponse struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com describeCompactLayoutsResponse"`

	Result *DescribeCompactLayoutsResult `xml:"result,omitempty"`
}

type DescribePrimaryCompactLayouts struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com describePrimaryCompactLayouts"`

	SObjectTypes []string `xml:"sObjectTypes,omitempty"`
}

type DescribePrimaryCompactLayoutsResponse struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com describePrimaryCompactLayoutsResponse"`

	Result []*DescribeCompactLayout `xml:"result,omitempty"`
}

type DescribePathAssistants struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com describePathAssistants"`

	SObjectType   string `xml:"sObjectType,omitempty"`
	PicklistValue string `xml:"picklistValue,omitempty"`
	RecordTypeIds []*ID  `xml:"recordTypeIds,omitempty"`
}

type DescribePathAssistantsResponse struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com describePathAssistantsResponse"`

	Result *DescribePathAssistantsResult `xml:"result,omitempty"`
}

// type DescribeApprovalLayout struct {
// 	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com describeApprovalLayout"`

// 	SObjectType          string   `xml:"sObjectType,omitempty"`
// 	ApprovalProcessNames []string `xml:"approvalProcessNames,omitempty"`
// }

type DescribeApprovalLayoutResponse struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com describeApprovalLayoutResponse"`

	Result *DescribeApprovalLayoutResult `xml:"result,omitempty"`
}

type DescribeSoftphoneLayout struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com describeSoftphoneLayout"`
}

type DescribeSoftphoneLayoutResponse struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com describeSoftphoneLayoutResponse"`

	Result *DescribeSoftphoneLayoutResult `xml:"result,omitempty"`
}

type DescribeSoqlListViews struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com describeSoqlListViews"`

	Request *DescribeSoqlListViewsRequest `xml:"request,omitempty"`
}

type DescribeSoqlListViewsResponse struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com describeSoqlListViewsResponse"`

	Result *DescribeSoqlListViewResult `xml:"result,omitempty"`
}

type ExecuteListView struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com executeListView"`

	Request *ExecuteListViewRequest `xml:"request,omitempty"`
}

type ExecuteListViewResponse struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com executeListViewResponse"`

	Result *ExecuteListViewResult `xml:"result,omitempty"`
}

type DescribeSObjectListViews struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com describeSObjectListViews"`

	SObjectType      string                    `xml:"sObjectType,omitempty"`
	RecentsOnly      bool                      `xml:"recentsOnly,omitempty"`
	IsSoqlCompatible *ListViewIsSoqlCompatible `xml:"isSoqlCompatible,omitempty"`
	Limit            int32                     `xml:"limit,omitempty"`
	Offset           int32                     `xml:"offset,omitempty"`
}

type DescribeSObjectListViewsResponse struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com describeSObjectListViewsResponse"`

	Result *DescribeSoqlListViewResult `xml:"result,omitempty"`
}

type DescribeSearchLayouts struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com describeSearchLayouts"`

	SObjectType []string `xml:"sObjectType,omitempty"`
}

type DescribeSearchLayoutsResponse struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com describeSearchLayoutsResponse"`

	Result []*DescribeSearchLayoutResult `xml:"result,omitempty"`
}

type DescribeSearchScopeOrder struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com describeSearchScopeOrder"`

	IncludeRealTimeEntities bool `xml:"includeRealTimeEntities,omitempty"`
}

type DescribeSearchScopeOrderResponse struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com describeSearchScopeOrderResponse"`

	Result []*DescribeSearchScopeOrderResult `xml:"result,omitempty"`
}

type DescribeSearchableEntities struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com describeSearchableEntities"`

	IncludeOnlyEntitiesWithTabs bool `xml:"includeOnlyEntitiesWithTabs,omitempty"`
}

type DescribeSearchableEntitiesResponse struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com describeSearchableEntitiesResponse"`

	Result []*DescribeSearchableEntityResult `xml:"result,omitempty"`
}

type DescribeTabs struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com describeTabs"`
}

type DescribeTabsResponse struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com describeTabsResponse"`

	Result []*DescribeTabSetResult `xml:"result,omitempty"`
}

type DescribeAllTabs struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com describeAllTabs"`
}

type DescribeAllTabsResponse struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com describeAllTabsResponse"`

	Result []*DescribeTab `xml:"result,omitempty"`
}

type DescribeNouns struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com describeNouns"`

	Nouns         string `xml:"nouns,omitempty"`
	OnlyRenamed   bool   `xml:"onlyRenamed,omitempty"`
	IncludeFields bool   `xml:"includeFields,omitempty"`
}

type DescribeNounsResponse struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com describeNounsResponse"`

	Result []*DescribeNounResult `xml:"result,omitempty"`
}

type Create struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com create"`

	SObjects []*SObject `xml:"sObjects,omitempty"`
}

type CreateResponse struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com createResponse"`

	Result []*SaveResult `xml:"result,omitempty"`
}

type SendEmail struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com sendEmail"`

	Messages *Email `xml:"messages,omitempty"`
}

type SendEmailResponse struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com sendEmailResponse"`

	Result *SendEmailResult `xml:"result,omitempty"`
}

type RenderEmailTemplate struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com renderEmailTemplate"`

	RenderRequests *RenderEmailTemplateRequest `xml:"renderRequests,omitempty"`
}

type RenderEmailTemplateResponse struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com renderEmailTemplateResponse"`

	Result *RenderEmailTemplateResult `xml:"result,omitempty"`
}

type RenderStoredEmailTemplate struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com renderStoredEmailTemplate"`

	Request *RenderStoredEmailTemplateRequest `xml:"request,omitempty"`
}

type RenderStoredEmailTemplateResponse struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com renderStoredEmailTemplateResponse"`

	Result *RenderStoredEmailTemplateResult `xml:"result,omitempty"`
}

type SendEmailMessage struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com sendEmailMessage"`

	Ids *ID `xml:"ids,omitempty"`
}

type SendEmailMessageResponse struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com sendEmailMessageResponse"`

	Result *SendEmailResult `xml:"result,omitempty"`
}

type Update struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com update"`

	SObjects []*SObject `xml:"sObjects,omitempty"`
}

type UpdateResponse struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com updateResponse"`

	Result []*SaveResult `xml:"result,omitempty"`
}

type Upsert struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com upsert"`

	ExternalIDFieldName string     `xml:"externalIDFieldName,omitempty"`
	SObjects            []*SObject `xml:"sObjects,omitempty"`
}

type UpsertResponse struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com upsertResponse"`

	Result []*UpsertResult `xml:"result,omitempty"`
}

type Merge struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com merge"`

	Request []*MergeRequest `xml:"request,omitempty"`
}

type MergeResponse struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com mergeResponse"`

	Result []*MergeResult `xml:"result,omitempty"`
}

type Delete struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com delete"`

	Ids []*ID `xml:"ids,omitempty"`
}

type DeleteResponse struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com deleteResponse"`

	Result []*DeleteResult `xml:"result,omitempty"`
}

type DeleteByExample struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com deleteByExample"`

	SObjects []*SObject `xml:"sObjects,omitempty"`
}

type DeleteByExampleResponse struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com deleteByExampleResponse"`

	Result []*DeleteByExampleResult `xml:"result,omitempty"`
}

type Undelete struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com undelete"`

	Ids []*ID `xml:"ids,omitempty"`
}

type UndeleteResponse struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com undeleteResponse"`

	Result []*UndeleteResult `xml:"result,omitempty"`
}

type EmptyRecycleBin struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com emptyRecycleBin"`

	Ids []*ID `xml:"ids,omitempty"`
}

type EmptyRecycleBinResponse struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com emptyRecycleBinResponse"`

	Result []*EmptyRecycleBinResult `xml:"result,omitempty"`
}

type Process struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com process"`

	Actions []*ProcessRequest `xml:"actions,omitempty"`
}

type ProcessResponse struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com processResponse"`

	Result []*ProcessResult `xml:"result,omitempty"`
}

type PerformQuickActions struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com performQuickActions"`

	QuickActions []*PerformQuickActionRequest `xml:"quickActions,omitempty"`
}

type PerformQuickActionsResponse struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com performQuickActionsResponse"`

	Result []*PerformQuickActionResult `xml:"result,omitempty"`
}

type RetrieveMassQuickActionTemplates struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com retrieveMassQuickActionTemplates"`

	QuickActionName string `xml:"quickActionName,omitempty"`
	ContextIds      []*ID  `xml:"contextIds,omitempty"`
}

type RetrieveMassQuickActionTemplatesResponse struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com retrieveMassQuickActionTemplatesResponse"`

	Result []*QuickActionTemplateResult `xml:"result,omitempty"`
}

type RetrieveQuickActionTemplates struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com retrieveQuickActionTemplates"`

	QuickActionNames []string `xml:"quickActionNames,omitempty"`
	ContextId        *ID      `xml:"contextId,omitempty"`
}

type RetrieveQuickActionTemplatesResponse struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com retrieveQuickActionTemplatesResponse"`

	Result []*QuickActionTemplateResult `xml:"result,omitempty"`
}

type DescribeQuickActions struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com describeQuickActions"`

	QuickActions []string `xml:"quickActions,omitempty"`
}

type DescribeQuickActionsResponse struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com describeQuickActionsResponse"`

	Result []*DescribeQuickActionResult `xml:"result,omitempty"`
}

type DescribeQuickActionsForRecordType struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com describeQuickActionsForRecordType"`

	QuickActions []string `xml:"quickActions,omitempty"`
	RecordTypeId string   `xml:"recordTypeId,omitempty"`
}

type DescribeQuickActionsForRecordTypeResponse struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com describeQuickActionsForRecordTypeResponse"`

	Result []*DescribeQuickActionResult `xml:"result,omitempty"`
}

type DescribeAvailableQuickActions struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com describeAvailableQuickActions"`

	ContextType string `xml:"contextType,omitempty"`
}

type DescribeAvailableQuickActionsResponse struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com describeAvailableQuickActionsResponse"`

	Result []*DescribeAvailableQuickActionResult `xml:"result,omitempty"`
}

type DescribeVisualForce struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com describeVisualForce"`

	IncludeAllDetails bool   `xml:"includeAllDetails,omitempty"`
	NamespacePrefix   string `xml:"namespacePrefix,omitempty"`
}

type DescribeVisualForceResponse struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com describeVisualForceResponse"`

	Result *DescribeVisualForceResult `xml:"result,omitempty"`
}

type Retrieve struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com retrieve"`

	FieldList   string `xml:"fieldList,omitempty"`
	SObjectType string `xml:"sObjectType,omitempty"`
	Ids         []*ID  `xml:"ids,omitempty"`
}

type RetrieveResponse struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com retrieveResponse"`

	Result []*SObject `xml:"result,omitempty"`
}

type ConvertLead struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com convertLead"`

	LeadConverts []*LeadConvert `xml:"leadConverts,omitempty"`
}

type ConvertLeadResponse struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com convertLeadResponse"`

	Result []*LeadConvertResult `xml:"result,omitempty"`
}

type GetUpdated struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com getUpdated"`

	SObjectType string    `xml:"sObjectType,omitempty"`
	StartDate   time.Time `xml:"startDate,omitempty"`
	EndDate     time.Time `xml:"endDate,omitempty"`
}

type GetUpdatedResponse struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com getUpdatedResponse"`

	Result *GetUpdatedResult `xml:"result,omitempty"`
}

type GetDeleted struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com getDeleted"`

	SObjectType string    `xml:"sObjectType,omitempty"`
	StartDate   time.Time `xml:"startDate,omitempty"`
	EndDate     time.Time `xml:"endDate,omitempty"`
}

type GetDeletedResponse struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com getDeletedResponse"`

	Result *GetDeletedResult `xml:"result,omitempty"`
}

type Logout struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com logout"`
}

type LogoutResponse struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com logoutResponse"`
}

type InvalidateSessions struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com invalidateSessions"`

	SessionIds []string `xml:"sessionIds,omitempty"`
}

type InvalidateSessionsResponse struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com invalidateSessionsResponse"`

	Result []*InvalidateSessionsResult `xml:"result,omitempty"`
}

type Query struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com query"`

	QueryString string `xml:"queryString,omitempty"`
}

type QueryResponse struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com queryResponse"`

	Result *QueryResult `xml:"result,omitempty"`
}

type QueryAll struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com queryAll"`

	QueryString string `xml:"queryString,omitempty"`
}

type QueryAllResponse struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com queryAllResponse"`

	Result *QueryResult `xml:"result,omitempty"`
}

type QueryMore struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com queryMore"`

	QueryLocator *QueryLocator `xml:"queryLocator,omitempty"`
}

type QueryMoreResponse struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com queryMoreResponse"`

	Result *QueryResult `xml:"result,omitempty"`
}

type Search struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com search"`

	SearchString string `xml:"searchString,omitempty"`
}

type SearchResponse struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com searchResponse"`

	Result *SearchResult `xml:"result,omitempty"`
}

type GetServerTimestamp struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com getServerTimestamp"`
}

type GetServerTimestampResponse struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com getServerTimestampResponse"`

	Result *GetServerTimestampResult `xml:"result,omitempty"`
}

type SetPassword struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com setPassword"`

	UserId   *ID    `xml:"userId,omitempty"`
	Password string `xml:"password,omitempty"`
}

type SetPasswordResponse struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com setPasswordResponse"`

	Result *SetPasswordResult `xml:"result,omitempty"`
}

type ChangeOwnPassword struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com changeOwnPassword"`

	OldPassword string `xml:"oldPassword,omitempty"`
	NewPassword string `xml:"newPassword,omitempty"`
}

type ChangeOwnPasswordResponse struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com changeOwnPasswordResponse"`

	Result *ChangeOwnPasswordResult `xml:"result,omitempty"`
}

type ResetPassword struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com resetPassword"`

	UserId *ID `xml:"userId,omitempty"`
}

type ResetPasswordResponse struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com resetPasswordResponse"`

	Result *ResetPasswordResult `xml:"result,omitempty"`
}

type GetUserInfo struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com getUserInfo"`
}

type GetUserInfoResponse struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com getUserInfoResponse"`

	Result *GetUserInfoResult `xml:"result,omitempty"`
}

type SessionHeader struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com SessionHeader"`

	SessionId string `xml:"sessionId,omitempty"`
}

type LoginScopeHeader struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com LoginScopeHeader"`

	OrganizationId *ID `xml:"organizationId,omitempty"`
	PortalId       *ID `xml:"portalId,omitempty"`
}

type QueryOptions struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com QueryOptions"`

	BatchSize int32 `xml:"batchSize,omitempty"`
}

type DebuggingHeader struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DebuggingHeader"`

	Categories []*LogInfo  `xml:"categories,omitempty"`
	DebugLevel *DebugLevel `xml:"debugLevel,omitempty"`
}

type DebuggingInfo struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DebuggingInfo"`

	DebugLog string `xml:"debugLog,omitempty"`
}

type PackageVersionHeader struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com PackageVersionHeader"`

	PackageVersions []*PackageVersion `xml:"packageVersions,omitempty"`
}

type AllowFieldTruncationHeader struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com AllowFieldTruncationHeader"`

	AllowFieldTruncation bool `xml:"allowFieldTruncation,omitempty"`
}

type DisableFeedTrackingHeader struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DisableFeedTrackingHeader"`

	DisableFeedTracking bool `xml:"disableFeedTracking,omitempty"`
}

type StreamingEnabledHeader struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com StreamingEnabledHeader"`

	StreamingEnabled bool `xml:"streamingEnabled,omitempty"`
}

type AllOrNoneHeader struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com AllOrNoneHeader"`

	AllOrNone bool `xml:"allOrNone,omitempty"`
}

type DuplicateRuleHeader struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DuplicateRuleHeader"`

	AllowSave            bool `xml:"allowSave,omitempty"`
	IncludeRecordDetails bool `xml:"includeRecordDetails,omitempty"`
	RunAsCurrentUser     bool `xml:"runAsCurrentUser,omitempty"`
}

type LimitInfoHeader struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com LimitInfoHeader"`

	LimitInfo []*LimitInfo `xml:"limitInfo,omitempty"`
}

type MruHeader struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com MruHeader"`

	UpdateMru bool `xml:"updateMru,omitempty"`
}

type EmailHeader struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com EmailHeader"`

	TriggerAutoResponseEmail bool `xml:"triggerAutoResponseEmail,omitempty"`
	TriggerOtherEmail        bool `xml:"triggerOtherEmail,omitempty"`
	TriggerUserEmail         bool `xml:"triggerUserEmail,omitempty"`
}

type AssignmentRuleHeader struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com AssignmentRuleHeader"`

	AssignmentRuleId *ID  `xml:"assignmentRuleId,omitempty"`
	UseDefaultRule   bool `xml:"useDefaultRule,omitempty"`
}

type UserTerritoryDeleteHeader struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com UserTerritoryDeleteHeader"`

	TransferToUserId *ID `xml:"transferToUserId,omitempty"`
}

type LocaleOptions struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com LocaleOptions"`

	Language       string `xml:"language,omitempty"`
	LocalizeErrors bool   `xml:"localizeErrors,omitempty"`
}

type OwnerChangeOptions struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com OwnerChangeOptions"`

	Options []*OwnerChangeOption `xml:"options,omitempty"`
}

type Address struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com address"`

	*Location

	City            string `xml:"city,omitempty"`
	Country         string `xml:"country,omitempty"`
	CountryCode     string `xml:"countryCode,omitempty"`
	GeocodeAccuracy string `xml:"geocodeAccuracy,omitempty"`
	PostalCode      string `xml:"postalCode,omitempty"`
	State           string `xml:"state,omitempty"`
	StateCode       string `xml:"stateCode,omitempty"`
	Street          string `xml:"street,omitempty"`
}

type Location struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com location"`

	Latitude  float64 `xml:"latitude,omitempty"`
	Longitude float64 `xml:"longitude,omitempty"`
}

type QueryResult struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com QueryResult"`

	Done         bool          `xml:"done,omitempty"`
	QueryLocator *QueryLocator `xml:"queryLocator,omitempty"`
	Records      []*SObject    `xml:"records,omitempty"`
	Size         int32         `xml:"size,omitempty"`
}

type SearchResult struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com SearchResult"`

	QueryId               string                 `xml:"queryId,omitempty"`
	SearchRecords         []*SearchRecord        `xml:"searchRecords,omitempty"`
	SearchResultsMetadata *SearchResultsMetadata `xml:"searchResultsMetadata,omitempty"`
}

type SearchRecord struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com SearchRecord"`

	Record               *SObject              `xml:"record,omitempty"`
	SearchRecordMetadata *SearchRecordMetadata `xml:"searchRecordMetadata,omitempty"`
	Snippet              *SearchSnippet        `xml:"snippet,omitempty"`
}

type SearchRecordMetadata struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com SearchRecordMetadata"`

	SearchPromoted bool `xml:"searchPromoted,omitempty"`
	SpellCorrected bool `xml:"spellCorrected,omitempty"`
}

type SearchSnippet struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com SearchSnippet"`

	Text        string           `xml:"text,omitempty"`
	WholeFields []*NameValuePair `xml:"wholeFields,omitempty"`
}

type SearchResultsMetadata struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com SearchResultsMetadata"`

	EntityLabelMetadata []*LabelsSearchMetadata `xml:"entityLabelMetadata,omitempty"`
	EntityMetadata      []*EntitySearchMetadata `xml:"entityMetadata,omitempty"`
}

type LabelsSearchMetadata struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com LabelsSearchMetadata"`

	EntityFieldLabels []*NameValuePair `xml:"entityFieldLabels,omitempty"`
	EntityName        string           `xml:"entityName,omitempty"`
}

type EntitySearchMetadata struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com EntitySearchMetadata"`

	EntityName              string                         `xml:"entityName,omitempty"`
	FieldMetadata           []*FieldLevelSearchMetadata    `xml:"fieldMetadata,omitempty"`
	IntentQueryMetadata     *EntityIntentQueryMetadata     `xml:"intentQueryMetadata,omitempty"`
	SearchPromotionMetadata *EntitySearchPromotionMetadata `xml:"searchPromotionMetadata,omitempty"`
	SpellCorrectionMetadata *EntitySpellCorrectionMetadata `xml:"spellCorrectionMetadata,omitempty"`
}

type FieldLevelSearchMetadata struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com FieldLevelSearchMetadata"`

	Label string `xml:"label,omitempty"`
	Name  string `xml:"name,omitempty"`
	Type_ string `xml:"type,omitempty"`
}

type EntitySpellCorrectionMetadata struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com EntitySpellCorrectionMetadata"`

	CorrectedQuery         string `xml:"correctedQuery,omitempty"`
	HasNonCorrectedResults bool   `xml:"hasNonCorrectedResults,omitempty"`
}

type EntitySearchPromotionMetadata struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com EntitySearchPromotionMetadata"`

	PromotedResultCount int32 `xml:"promotedResultCount,omitempty"`
}

type EntityIntentQueryMetadata struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com EntityIntentQueryMetadata"`

	IntentQuery bool   `xml:"intentQuery,omitempty"`
	Message     string `xml:"message,omitempty"`
}

type RelationshipReferenceTo struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com RelationshipReferenceTo"`

	ReferenceTo []string `xml:"referenceTo,omitempty"`
}

type RecordTypesSupported struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com RecordTypesSupported"`

	RecordTypeInfos []*RecordTypeInfo `xml:"recordTypeInfos,omitempty"`
}

type JunctionIdListNames struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com JunctionIdListNames"`

	Names []string `xml:"names,omitempty"`
}

type SearchLayoutButtonsDisplayed struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com SearchLayoutButtonsDisplayed"`

	Applicable bool                  `xml:"applicable,omitempty"`
	Buttons    []*SearchLayoutButton `xml:"buttons,omitempty"`
}

type SearchLayoutButton struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com SearchLayoutButton"`

	ApiName string `xml:"apiName,omitempty"`
	Label   string `xml:"label,omitempty"`
}

type SearchLayoutFieldsDisplayed struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com SearchLayoutFieldsDisplayed"`

	Applicable bool                 `xml:"applicable,omitempty"`
	Fields     []*SearchLayoutField `xml:"fields,omitempty"`
}

type SearchLayoutField struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com SearchLayoutField"`

	ApiName  string `xml:"apiName,omitempty"`
	Label    string `xml:"label,omitempty"`
	Sortable bool   `xml:"sortable,omitempty"`
}

type NameValuePair struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com NameValuePair"`

	Name  string `xml:"name,omitempty"`
	Value string `xml:"value,omitempty"`
}

type NameObjectValuePair struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com NameObjectValuePair"`

	Name  string        `xml:"name,omitempty"`
	Value []interface{} `xml:"value,omitempty"`
}

type GetUpdatedResult struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com GetUpdatedResult"`

	Ids               []*ID     `xml:"ids,omitempty"`
	LatestDateCovered time.Time `xml:"latestDateCovered,omitempty"`
}

type GetDeletedResult struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com GetDeletedResult"`

	DeletedRecords        []*DeletedRecord `xml:"deletedRecords,omitempty"`
	EarliestDateAvailable time.Time        `xml:"earliestDateAvailable,omitempty"`
	LatestDateCovered     time.Time        `xml:"latestDateCovered,omitempty"`
}

type DeletedRecord struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DeletedRecord"`

	DeletedDate time.Time `xml:"deletedDate,omitempty"`
	Id          *ID       `xml:"id,omitempty"`
}

type GetServerTimestampResult struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com GetServerTimestampResult"`

	Timestamp time.Time `xml:"timestamp,omitempty"`
}

type InvalidateSessionsResult struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com InvalidateSessionsResult"`

	Errors  []*Error `xml:"errors,omitempty"`
	Success bool     `xml:"success,omitempty"`
}

type SetPasswordResult struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com SetPasswordResult"`
}

type ChangeOwnPasswordResult struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com ChangeOwnPasswordResult"`
}

type ResetPasswordResult struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com ResetPasswordResult"`

	Password string `xml:"password,omitempty"`
}

type GetUserInfoResult struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com userInfo"`

	AccessibilityMode          bool   `xml:"accessibilityMode,omitempty"`
	ChatterExternal            bool   `xml:"chatterExternal,omitempty"`
	CurrencySymbol             string `xml:"currencySymbol,omitempty"`
	OrgAttachmentFileSizeLimit int32  `xml:"orgAttachmentFileSizeLimit,omitempty"`
	OrgDefaultCurrencyIsoCode  string `xml:"orgDefaultCurrencyIsoCode,omitempty"`
	OrgDefaultCurrencyLocale   string `xml:"orgDefaultCurrencyLocale,omitempty"`
	OrgDisallowHtmlAttachments bool   `xml:"orgDisallowHtmlAttachments,omitempty"`
	OrgHasPersonAccounts       bool   `xml:"orgHasPersonAccounts,omitempty"`
	OrganizationId             *ID    `xml:"organizationId,omitempty"`
	OrganizationMultiCurrency  bool   `xml:"organizationMultiCurrency,omitempty"`
	OrganizationName           string `xml:"organizationName,omitempty"`
	ProfileId                  *ID    `xml:"profileId,omitempty"`
	RoleId                     *ID    `xml:"roleId,omitempty"`
	SessionSecondsValid        int32  `xml:"sessionSecondsValid,omitempty"`
	UserDefaultCurrencyIsoCode string `xml:"userDefaultCurrencyIsoCode,omitempty"`
	UserEmail                  string `xml:"userEmail,omitempty"`
	UserFullName               string `xml:"userFullName,omitempty"`
	UserId                     *ID    `xml:"userId,omitempty"`
	UserLanguage               string `xml:"userLanguage,omitempty"`
	UserLocale                 string `xml:"userLocale,omitempty"`
	UserName                   string `xml:"userName,omitempty"`
	UserTimeZone               string `xml:"userTimeZone,omitempty"`
	UserType                   string `xml:"userType,omitempty"`
	UserUiSkin                 string `xml:"userUiSkin,omitempty"`
}

type LoginResult struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com result"`

	MetadataServerUrl string             `xml:"metadataServerUrl,omitempty"`
	PasswordExpired   bool               `xml:"passwordExpired,omitempty"`
	Sandbox           bool               `xml:"sandbox,omitempty"`
	ServerUrl         string             `xml:"serverUrl,omitempty"`
	SessionId         string             `xml:"sessionId,omitempty"`
	UserId            *ID                `xml:"userId,omitempty"`
	UserInfo          *GetUserInfoResult `xml:"userInfo,omitempty"`
}

type ExtendedErrorDetails struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com ExtendedErrorDetails"`

	ExtendedErrorCode *ExtendedErrorCode `xml:"extendedErrorCode,omitempty"`
}

type Error struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com Error"`

	ExtendedErrorDetails []*ExtendedErrorDetails `xml:"extendedErrorDetails,omitempty"`
	Fields               []string                `xml:"fields,omitempty"`
	Message              string                  `xml:"message,omitempty"`
	StatusCode           *StatusCode             `xml:"statusCode,omitempty"`
}

type SendEmailError struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com SendEmailError"`

	Fields         []string    `xml:"fields,omitempty"`
	Message        string      `xml:"message,omitempty"`
	StatusCode     *StatusCode `xml:"statusCode,omitempty"`
	TargetObjectId *ID         `xml:"targetObjectId,omitempty"`
}

type SaveResult struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com SaveResult"`

	Errors  []*Error `xml:"errors,omitempty"`
	Id      *ID      `xml:"id,omitempty"`
	Success bool     `xml:"success,omitempty"`
}

type RenderEmailTemplateError struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com RenderEmailTemplateError"`

	FieldName  string      `xml:"fieldName,omitempty"`
	Message    string      `xml:"message,omitempty"`
	Offset     int32       `xml:"offset,omitempty"`
	StatusCode *StatusCode `xml:"statusCode,omitempty"`
}

type UpsertResult struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com UpsertResult"`

	Created bool     `xml:"created,omitempty"`
	Errors  []*Error `xml:"errors,omitempty"`
	Id      *ID      `xml:"id,omitempty"`
	Success bool     `xml:"success,omitempty"`
}

type PerformQuickActionResult struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com PerformQuickActionResult"`

	ContextId      *ID      `xml:"contextId,omitempty"`
	Created        bool     `xml:"created,omitempty"`
	Errors         []*Error `xml:"errors,omitempty"`
	FeedItemIds    []*ID    `xml:"feedItemIds,omitempty"`
	Ids            []*ID    `xml:"ids,omitempty"`
	Success        bool     `xml:"success,omitempty"`
	SuccessMessage string   `xml:"successMessage,omitempty"`
}

type QuickActionTemplateResult struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com QuickActionTemplateResult"`

	ContextId            string   `xml:"contextId,omitempty"`
	DefaultValueFormulas *SObject `xml:"defaultValueFormulas,omitempty"`
	DefaultValues        *SObject `xml:"defaultValues,omitempty"`
	Errors               []*Error `xml:"errors,omitempty"`
	Success              bool     `xml:"success,omitempty"`
}

type MergeRequest struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com MergeRequest"`

	AdditionalInformationMap []*AdditionalInformationMap `xml:"additionalInformationMap,omitempty"`
	MasterRecord             *SObject                    `xml:"masterRecord,omitempty"`
	RecordToMergeIds         []*ID                       `xml:"recordToMergeIds,omitempty"`
}

type MergeResult struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com MergeResult"`

	Errors            []*Error `xml:"errors,omitempty"`
	Id                *ID      `xml:"id,omitempty"`
	MergedRecordIds   []*ID    `xml:"mergedRecordIds,omitempty"`
	Success           bool     `xml:"success,omitempty"`
	UpdatedRelatedIds []*ID    `xml:"updatedRelatedIds,omitempty"`
}

type ProcessRequest struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com ProcessRequest"`

	Comments        string `xml:"comments,omitempty"`
	NextApproverIds []*ID  `xml:"nextApproverIds,omitempty"`
}

type ProcessSubmitRequest struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com ProcessSubmitRequest"`

	*ProcessRequest

	ObjectId                  *ID    `xml:"objectId,omitempty"`
	SubmitterId               *ID    `xml:"submitterId,omitempty"`
	ProcessDefinitionNameOrId string `xml:"processDefinitionNameOrId,omitempty"`
	SkipEntryCriteria         bool   `xml:"skipEntryCriteria,omitempty"`
}

type ProcessWorkitemRequest struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com ProcessWorkitemRequest"`

	*ProcessRequest

	Action     string `xml:"action,omitempty"`
	WorkitemId *ID    `xml:"workitemId,omitempty"`
}

type PerformQuickActionRequest struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com PerformQuickActionRequest"`

	ContextId       *ID        `xml:"contextId,omitempty"`
	QuickActionName string     `xml:"quickActionName,omitempty"`
	Records         []*SObject `xml:"records,omitempty"`
}

type DescribeAvailableQuickActionResult struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DescribeAvailableQuickActionResult"`

	ActionEnumOrId string `xml:"actionEnumOrId,omitempty"`
	Label          string `xml:"label,omitempty"`
	Name           string `xml:"name,omitempty"`
	Type_          string `xml:"type,omitempty"`
}

type DescribeQuickActionResult struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DescribeQuickActionResult"`

	AccessLevelRequired             *ShareAccessLevel                  `xml:"accessLevelRequired,omitempty"`
	ActionEnumOrId                  string                             `xml:"actionEnumOrId,omitempty"`
	CanvasApplicationId             *ID                                `xml:"canvasApplicationId,omitempty"`
	CanvasApplicationName           string                             `xml:"canvasApplicationName,omitempty"`
	Colors                          []*DescribeColor                   `xml:"colors,omitempty"`
	ContextSobjectType              string                             `xml:"contextSobjectType,omitempty"`
	DefaultValues                   []*DescribeQuickActionDefaultValue `xml:"defaultValues,omitempty"`
	FlowDevName                     string                             `xml:"flowDevName,omitempty"`
	FlowRecordIdVar                 string                             `xml:"flowRecordIdVar,omitempty"`
	Height                          int32                              `xml:"height,omitempty"`
	IconName                        string                             `xml:"iconName,omitempty"`
	IconUrl                         string                             `xml:"iconUrl,omitempty"`
	Icons                           []*DescribeIcon                    `xml:"icons,omitempty"`
	Label                           string                             `xml:"label,omitempty"`
	Layout                          *DescribeLayoutSection             `xml:"layout,omitempty"`
	LightningComponentBundleId      *ID                                `xml:"lightningComponentBundleId,omitempty"`
	LightningComponentBundleName    string                             `xml:"lightningComponentBundleName,omitempty"`
	LightningComponentQualifiedName string                             `xml:"lightningComponentQualifiedName,omitempty"`
	MiniIconUrl                     string                             `xml:"miniIconUrl,omitempty"`
	Name                            string                             `xml:"name,omitempty"`
	ShowQuickActionLcHeader         bool                               `xml:"showQuickActionLcHeader,omitempty"`
	ShowQuickActionVfHeader         bool                               `xml:"showQuickActionVfHeader,omitempty"`
	TargetParentField               string                             `xml:"targetParentField,omitempty"`
	TargetRecordTypeId              *ID                                `xml:"targetRecordTypeId,omitempty"`
	TargetSobjectType               string                             `xml:"targetSobjectType,omitempty"`
	Type_                           string                             `xml:"type,omitempty"`
	VisualforcePageName             string                             `xml:"visualforcePageName,omitempty"`
	VisualforcePageUrl              string                             `xml:"visualforcePageUrl,omitempty"`
	Width                           int32                              `xml:"width,omitempty"`
}

type DescribeQuickActionDefaultValue struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DescribeQuickActionDefaultValue"`

	DefaultValue string `xml:"defaultValue,omitempty"`
	Field        string `xml:"field,omitempty"`
}

type DescribeVisualForceResult struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DescribeVisualForceResult"`

	Domain string `xml:"domain,omitempty"`
}

type ProcessResult struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com ProcessResult"`

	ActorIds       []*ID    `xml:"actorIds,omitempty"`
	EntityId       *ID      `xml:"entityId,omitempty"`
	Errors         []*Error `xml:"errors,omitempty"`
	InstanceId     *ID      `xml:"instanceId,omitempty"`
	InstanceStatus string   `xml:"instanceStatus,omitempty"`
	NewWorkitemIds []*ID    `xml:"newWorkitemIds,omitempty"`
	Success        bool     `xml:"success,omitempty"`
}

type DeleteResult struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DeleteResult"`

	Errors  []*Error `xml:"errors,omitempty"`
	Id      *ID      `xml:"id,omitempty"`
	Success bool     `xml:"success,omitempty"`
}

type UndeleteResult struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com UndeleteResult"`

	Errors  []*Error `xml:"errors,omitempty"`
	Id      *ID      `xml:"id,omitempty"`
	Success bool     `xml:"success,omitempty"`
}

type DeleteByExampleResult struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DeleteByExampleResult"`

	Entity   *SObject `xml:"entity,omitempty"`
	Errors   []*Error `xml:"errors,omitempty"`
	RowCount int64    `xml:"rowCount,omitempty"`
	Success  bool     `xml:"success,omitempty"`
}

type EmptyRecycleBinResult struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com EmptyRecycleBinResult"`

	Errors  []*Error `xml:"errors,omitempty"`
	Id      *ID      `xml:"id,omitempty"`
	Success bool     `xml:"success,omitempty"`
}

type LeadConvert struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com LeadConvert"`

	AccountId                *ID      `xml:"accountId,omitempty"`
	AccountRecord            *SObject `xml:"accountRecord,omitempty"`
	BypassAccountDedupeCheck bool     `xml:"bypassAccountDedupeCheck,omitempty"`
	BypassContactDedupeCheck bool     `xml:"bypassContactDedupeCheck,omitempty"`
	ContactId                *ID      `xml:"contactId,omitempty"`
	ContactRecord            *SObject `xml:"contactRecord,omitempty"`
	ConvertedStatus          string   `xml:"convertedStatus,omitempty"`
	DoNotCreateOpportunity   bool     `xml:"doNotCreateOpportunity,omitempty"`
	LeadId                   *ID      `xml:"leadId,omitempty"`
	OpportunityId            *ID      `xml:"opportunityId,omitempty"`
	OpportunityName          string   `xml:"opportunityName,omitempty"`
	OpportunityRecord        *SObject `xml:"opportunityRecord,omitempty"`
	OverwriteLeadSource      bool     `xml:"overwriteLeadSource,omitempty"`
	OwnerId                  *ID      `xml:"ownerId,omitempty"`
	SendNotificationEmail    bool     `xml:"sendNotificationEmail,omitempty"`
}

type LeadConvertResult struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com LeadConvertResult"`

	AccountId     *ID      `xml:"accountId,omitempty"`
	ContactId     *ID      `xml:"contactId,omitempty"`
	Errors        []*Error `xml:"errors,omitempty"`
	LeadId        *ID      `xml:"leadId,omitempty"`
	OpportunityId *ID      `xml:"opportunityId,omitempty"`
	Success       bool     `xml:"success,omitempty"`
}

type DescribeSObjectResult struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DescribeSObjectResult"`

	ActionOverrides       []*ActionOverride    `xml:"actionOverrides,omitempty"`
	Activateable          bool                 `xml:"activateable,omitempty"`
	ChildRelationships    []*ChildRelationship `xml:"childRelationships,omitempty"`
	CompactLayoutable     bool                 `xml:"compactLayoutable,omitempty"`
	Createable            bool                 `xml:"createable,omitempty"`
	Custom                bool                 `xml:"custom,omitempty"`
	CustomSetting         bool                 `xml:"customSetting,omitempty"`
	Deletable             bool                 `xml:"deletable,omitempty"`
	DeprecatedAndHidden   bool                 `xml:"deprecatedAndHidden,omitempty"`
	FeedEnabled           bool                 `xml:"feedEnabled,omitempty"`
	Fields                []*Field             `xml:"fields,omitempty"`
	HasSubtypes           bool                 `xml:"hasSubtypes,omitempty"`
	IdEnabled             bool                 `xml:"idEnabled,omitempty"`
	IsSubtype             bool                 `xml:"isSubtype,omitempty"`
	KeyPrefix             string               `xml:"keyPrefix,omitempty"`
	Label                 string               `xml:"label,omitempty"`
	LabelPlural           string               `xml:"labelPlural,omitempty"`
	Layoutable            bool                 `xml:"layoutable,omitempty"`
	Mergeable             bool                 `xml:"mergeable,omitempty"`
	MruEnabled            bool                 `xml:"mruEnabled,omitempty"`
	Name                  string               `xml:"name,omitempty"`
	NamedLayoutInfos      []*NamedLayoutInfo   `xml:"namedLayoutInfos,omitempty"`
	NetworkScopeFieldName string               `xml:"networkScopeFieldName,omitempty"`
	Queryable             bool                 `xml:"queryable,omitempty"`
	RecordTypeInfos       []*RecordTypeInfo    `xml:"recordTypeInfos,omitempty"`
	Replicateable         bool                 `xml:"replicateable,omitempty"`
	Retrieveable          bool                 `xml:"retrieveable,omitempty"`
	SearchLayoutable      bool                 `xml:"searchLayoutable,omitempty"`
	Searchable            bool                 `xml:"searchable,omitempty"`
	SupportedScopes       []*ScopeInfo         `xml:"supportedScopes,omitempty"`
	Triggerable           bool                 `xml:"triggerable,omitempty"`
	Undeletable           bool                 `xml:"undeletable,omitempty"`
	Updateable            bool                 `xml:"updateable,omitempty"`
	UrlDetail             string               `xml:"urlDetail,omitempty"`
	UrlEdit               string               `xml:"urlEdit,omitempty"`
	UrlNew                string               `xml:"urlNew,omitempty"`
}

type DescribeGlobalSObjectResult struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DescribeGlobalSObjectResult"`

	Activateable        bool   `xml:"activateable,omitempty"`
	Createable          bool   `xml:"createable,omitempty"`
	Custom              bool   `xml:"custom,omitempty"`
	CustomSetting       bool   `xml:"customSetting,omitempty"`
	Deletable           bool   `xml:"deletable,omitempty"`
	DeprecatedAndHidden bool   `xml:"deprecatedAndHidden,omitempty"`
	FeedEnabled         bool   `xml:"feedEnabled,omitempty"`
	HasSubtypes         bool   `xml:"hasSubtypes,omitempty"`
	IdEnabled           bool   `xml:"idEnabled,omitempty"`
	IsSubtype           bool   `xml:"isSubtype,omitempty"`
	KeyPrefix           string `xml:"keyPrefix,omitempty"`
	Label               string `xml:"label,omitempty"`
	LabelPlural         string `xml:"labelPlural,omitempty"`
	Layoutable          bool   `xml:"layoutable,omitempty"`
	Mergeable           bool   `xml:"mergeable,omitempty"`
	MruEnabled          bool   `xml:"mruEnabled,omitempty"`
	Name                string `xml:"name,omitempty"`
	Queryable           bool   `xml:"queryable,omitempty"`
	Replicateable       bool   `xml:"replicateable,omitempty"`
	Retrieveable        bool   `xml:"retrieveable,omitempty"`
	Searchable          bool   `xml:"searchable,omitempty"`
	Triggerable         bool   `xml:"triggerable,omitempty"`
	Undeletable         bool   `xml:"undeletable,omitempty"`
	Updateable          bool   `xml:"updateable,omitempty"`
}

type ChildRelationship struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com ChildRelationship"`

	CascadeDelete       bool     `xml:"cascadeDelete,omitempty"`
	ChildSObject        string   `xml:"childSObject,omitempty"`
	DeprecatedAndHidden bool     `xml:"deprecatedAndHidden,omitempty"`
	Field               string   `xml:"field,omitempty"`
	JunctionIdListNames []string `xml:"junctionIdListNames,omitempty"`
	JunctionReferenceTo []string `xml:"junctionReferenceTo,omitempty"`
	RelationshipName    string   `xml:"relationshipName,omitempty"`
	RestrictedDelete    bool     `xml:"restrictedDelete,omitempty"`
}

type DescribeGlobalResult struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DescribeGlobalResult"`

	Encoding     string                         `xml:"encoding,omitempty"`
	MaxBatchSize int32                          `xml:"maxBatchSize,omitempty"`
	Sobjects     []*DescribeGlobalSObjectResult `xml:"sobjects,omitempty"`
}

type DescribeGlobalThemes struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DescribeGlobalTheme"`

	Global *DescribeGlobalResult `xml:"global,omitempty"`
	Theme  *DescribeThemeResult  `xml:"theme,omitempty"`
}

type ScopeInfo struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com ScopeInfo"`

	Label string `xml:"label,omitempty"`
	Name  string `xml:"name,omitempty"`
}

type StringList struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com StringList"`

	Values []string `xml:"values,omitempty"`
}

type FilteredLookupInfo struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com FilteredLookupInfo"`

	ControllingFields []string `xml:"controllingFields,omitempty"`
	Dependent         bool     `xml:"dependent,omitempty"`
	OptionalFilter    bool     `xml:"optionalFilter,omitempty"`
}

type Field struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com Field"`

	Aggregatable                 bool                `xml:"aggregatable,omitempty"`
	AiPredictionField            bool                `xml:"aiPredictionField,omitempty"`
	AutoNumber                   bool                `xml:"autoNumber,omitempty"`
	ByteLength                   int32               `xml:"byteLength,omitempty"`
	Calculated                   bool                `xml:"calculated,omitempty"`
	CalculatedFormula            string              `xml:"calculatedFormula,omitempty"`
	CascadeDelete                bool                `xml:"cascadeDelete,omitempty"`
	CaseSensitive                bool                `xml:"caseSensitive,omitempty"`
	CompoundFieldName            string              `xml:"compoundFieldName,omitempty"`
	ControllerName               string              `xml:"controllerName,omitempty"`
	Createable                   bool                `xml:"createable,omitempty"`
	Custom                       bool                `xml:"custom,omitempty"`
	DefaultValue                 interface{}         `xml:"defaultValue,omitempty"`
	DefaultValueFormula          string              `xml:"defaultValueFormula,omitempty"`
	DefaultedOnCreate            bool                `xml:"defaultedOnCreate,omitempty"`
	DependentPicklist            bool                `xml:"dependentPicklist,omitempty"`
	DeprecatedAndHidden          bool                `xml:"deprecatedAndHidden,omitempty"`
	Digits                       int32               `xml:"digits,omitempty"`
	DisplayLocationInDecimal     bool                `xml:"displayLocationInDecimal,omitempty"`
	Encrypted                    bool                `xml:"encrypted,omitempty"`
	ExternalId                   bool                `xml:"externalId,omitempty"`
	ExtraTypeInfo                string              `xml:"extraTypeInfo,omitempty"`
	Filterable                   bool                `xml:"filterable,omitempty"`
	FilteredLookupInfo           *FilteredLookupInfo `xml:"filteredLookupInfo,omitempty"`
	FormulaTreatNullNumberAsZero bool                `xml:"formulaTreatNullNumberAsZero,omitempty"`
	Groupable                    bool                `xml:"groupable,omitempty"`
	HighScaleNumber              bool                `xml:"highScaleNumber,omitempty"`
	HtmlFormatted                bool                `xml:"htmlFormatted,omitempty"`
	IdLookup                     bool                `xml:"idLookup,omitempty"`
	InlineHelpText               string              `xml:"inlineHelpText,omitempty"`
	Label                        string              `xml:"label,omitempty"`
	Length                       int32               `xml:"length,omitempty"`
	Mask                         string              `xml:"mask,omitempty"`
	MaskType                     string              `xml:"maskType,omitempty"`
	Name                         string              `xml:"name,omitempty"`
	NameField                    bool                `xml:"nameField,omitempty"`
	NamePointing                 bool                `xml:"namePointing,omitempty"`
	Nillable                     bool                `xml:"nillable,omitempty"`
	Permissionable               bool                `xml:"permissionable,omitempty"`
	PicklistValues               []*PicklistEntry    `xml:"picklistValues,omitempty"`
	PolymorphicForeignKey        bool                `xml:"polymorphicForeignKey,omitempty"`
	Precision                    int32               `xml:"precision,omitempty"`
	QueryByDistance              bool                `xml:"queryByDistance,omitempty"`
	ReferenceTargetField         string              `xml:"referenceTargetField,omitempty"`
	ReferenceTo                  []string            `xml:"referenceTo,omitempty"`
	RelationshipName             string              `xml:"relationshipName,omitempty"`
	RelationshipOrder            int32               `xml:"relationshipOrder,omitempty"`
	RestrictedDelete             bool                `xml:"restrictedDelete,omitempty"`
	RestrictedPicklist           bool                `xml:"restrictedPicklist,omitempty"`
	Scale                        int32               `xml:"scale,omitempty"`
	SearchPrefilterable          bool                `xml:"searchPrefilterable,omitempty"`
	SoapType                     *SoapType           `xml:"soapType,omitempty"`
	Sortable                     bool                `xml:"sortable,omitempty"`
	Type_                        *FieldType          `xml:"type,omitempty"`
	Unique                       bool                `xml:"unique,omitempty"`
	Updateable                   bool                `xml:"updateable,omitempty"`
	WriteRequiresMasterRead      bool                `xml:"writeRequiresMasterRead,omitempty"`
}

type PicklistEntry struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com PicklistEntry"`

	Active       bool   `xml:"active,omitempty"`
	DefaultValue bool   `xml:"defaultValue,omitempty"`
	Label        string `xml:"label,omitempty"`
	ValidFor     []byte `xml:"validFor,omitempty"`
	Value        string `xml:"value,omitempty"`
}

type DescribeDataCategoryGroupResult struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DescribeDataCategoryGroupResult"`

	CategoryCount int32  `xml:"categoryCount,omitempty"`
	Description   string `xml:"description,omitempty"`
	Label         string `xml:"label,omitempty"`
	Name          string `xml:"name,omitempty"`
	Sobject       string `xml:"sobject,omitempty"`
}

type DescribeDataCategoryGroupStructureResult struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DescribeDataCategoryGroupStructureResult"`

	Description   string          `xml:"description,omitempty"`
	Label         string          `xml:"label,omitempty"`
	Name          string          `xml:"name,omitempty"`
	Sobject       string          `xml:"sobject,omitempty"`
	TopCategories []*DataCategory `xml:"topCategories,omitempty"`
}

type DataCategoryGroupSobjectTypePair struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DataCategoryGroupSobjectTypePair"`

	DataCategoryGroupName string `xml:"dataCategoryGroupName,omitempty"`
	Sobject               string `xml:"sobject,omitempty"`
}

type DataCategory struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DataCategory"`

	ChildCategories []*DataCategory `xml:"childCategories,omitempty"`
	Label           string          `xml:"label,omitempty"`
	Name            string          `xml:"name,omitempty"`
}

type DescribeDataCategoryMappingResult struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DescribeDataCategoryMappingResult"`

	DataCategoryGroupId    string `xml:"dataCategoryGroupId,omitempty"`
	DataCategoryGroupLabel string `xml:"dataCategoryGroupLabel,omitempty"`
	DataCategoryGroupName  string `xml:"dataCategoryGroupName,omitempty"`
	DataCategoryId         string `xml:"dataCategoryId,omitempty"`
	DataCategoryLabel      string `xml:"dataCategoryLabel,omitempty"`
	DataCategoryName       string `xml:"dataCategoryName,omitempty"`
	Id                     string `xml:"id,omitempty"`
	MappedEntity           string `xml:"mappedEntity,omitempty"`
	MappedField            string `xml:"mappedField,omitempty"`
}

type KnowledgeSettings struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com KnowledgeSettings"`

	DefaultLanguage  string                   `xml:"defaultLanguage,omitempty"`
	KnowledgeEnabled bool                     `xml:"knowledgeEnabled,omitempty"`
	Languages        []*KnowledgeLanguageItem `xml:"languages,omitempty"`
}

type KnowledgeLanguageItem struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com KnowledgeLanguageItem"`

	Active     bool   `xml:"active,omitempty"`
	AssigneeId string `xml:"assigneeId,omitempty"`
	Name       string `xml:"name,omitempty"`
}

type FieldDiff struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com FieldDiff"`

	Difference *DifferenceType `xml:"difference,omitempty"`
	Name       string          `xml:"name,omitempty"`
}

type AdditionalInformationMap struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com AdditionalInformationMap"`

	Name  string `xml:"name,omitempty"`
	Value string `xml:"value,omitempty"`
}

type MatchRecord struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com MatchRecord"`

	AdditionalInformation []*AdditionalInformationMap `xml:"additionalInformation,omitempty"`
	FieldDiffs            []*FieldDiff                `xml:"fieldDiffs,omitempty"`
	MatchConfidence       float64                     `xml:"matchConfidence,omitempty"`
	Record                *SObject                    `xml:"record,omitempty"`
}

type MatchResult struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com MatchResult"`

	EntityType   string         `xml:"entityType,omitempty"`
	Errors       []*Error       `xml:"errors,omitempty"`
	MatchEngine  string         `xml:"matchEngine,omitempty"`
	MatchRecords []*MatchRecord `xml:"matchRecords,omitempty"`
	Rule         string         `xml:"rule,omitempty"`
	Size         int32          `xml:"size,omitempty"`
	Success      bool           `xml:"success,omitempty"`
}

type DuplicateResult struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DuplicateResult"`

	AllowSave               bool           `xml:"allowSave,omitempty"`
	DuplicateRule           string         `xml:"duplicateRule,omitempty"`
	DuplicateRuleEntityType string         `xml:"duplicateRuleEntityType,omitempty"`
	ErrorMessage            string         `xml:"errorMessage,omitempty"`
	MatchResults            []*MatchResult `xml:"matchResults,omitempty"`
}

type DuplicateError struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DuplicateError"`

	*Error

	DuplicateResult *DuplicateResult `xml:"duplicateResult,omitempty"`
}

type DescribeNounResult struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DescribeNounResult"`

	CaseValues    []*NameCaseValue `xml:"caseValues,omitempty"`
	DeveloperName string           `xml:"developerName,omitempty"`
	Gender        *Gender          `xml:"gender,omitempty"`
	Name          string           `xml:"name,omitempty"`
	PluralAlias   string           `xml:"pluralAlias,omitempty"`
	StartsWith    *StartsWith      `xml:"startsWith,omitempty"`
}

type NameCaseValue struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com NameCaseValue"`

	Article    *Article           `xml:"article,omitempty"`
	CaseType   *CaseType          `xml:"caseType,omitempty"`
	Number     *GrammaticalNumber `xml:"number,omitempty"`
	Possessive *Possessive        `xml:"possessive,omitempty"`
	Value      string             `xml:"value,omitempty"`
}

type FindDuplicatesResult struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com FindDuplicatesResult"`

	DuplicateResults []*DuplicateResult `xml:"duplicateResults,omitempty"`
	Errors           []*Error           `xml:"errors,omitempty"`
	Success          bool               `xml:"success,omitempty"`
}

type DescribeAppMenuResult struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DescribeAppMenuResult"`

	AppMenuItems []*DescribeAppMenuItem `xml:"appMenuItems,omitempty"`
}

type DescribeAppMenuItem struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DescribeAppMenuItem"`

	Colors  []*DescribeColor `xml:"colors,omitempty"`
	Content string           `xml:"content,omitempty"`
	Icons   []*DescribeIcon  `xml:"icons,omitempty"`
	Label   string           `xml:"label,omitempty"`
	Name    string           `xml:"name,omitempty"`
	Type_   string           `xml:"type,omitempty"`
	Url     string           `xml:"url,omitempty"`
}

type DescribeThemeResult struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DescribeThemeResult"`

	ThemeItems []*DescribeThemeItem `xml:"themeItems,omitempty"`
}

type DescribeThemeItem struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DescribeThemeItem"`

	Colors []*DescribeColor `xml:"colors,omitempty"`
	Icons  []*DescribeIcon  `xml:"icons,omitempty"`
	Name   string           `xml:"name,omitempty"`
}

type DescribeSoftphoneLayoutResult struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DescribeSoftphoneLayoutResult"`

	CallTypes []*DescribeSoftphoneLayoutCallType `xml:"callTypes,omitempty"`
	Id        *ID                                `xml:"id,omitempty"`
	Name      string                             `xml:"name,omitempty"`
}

type DescribeSoftphoneLayoutCallType struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DescribeSoftphoneLayoutCallType"`

	InfoFields           []*DescribeSoftphoneLayoutInfoField `xml:"infoFields,omitempty"`
	Name                 string                              `xml:"name,omitempty"`
	ScreenPopOptions     []*DescribeSoftphoneScreenPopOption `xml:"screenPopOptions,omitempty"`
	ScreenPopsOpenWithin string                              `xml:"screenPopsOpenWithin,omitempty"`
	Sections             []*DescribeSoftphoneLayoutSection   `xml:"sections,omitempty"`
}

type DescribeSoftphoneScreenPopOption struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DescribeSoftphoneScreenPopOption"`

	MatchType     string `xml:"matchType,omitempty"`
	ScreenPopData string `xml:"screenPopData,omitempty"`
	ScreenPopType string `xml:"screenPopType,omitempty"`
}

type DescribeSoftphoneLayoutInfoField struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DescribeSoftphoneLayoutInfoField"`

	Name string `xml:"name,omitempty"`
}

type DescribeSoftphoneLayoutSection struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DescribeSoftphoneLayoutSection"`

	EntityApiName string                         `xml:"entityApiName,omitempty"`
	Items         []*DescribeSoftphoneLayoutItem `xml:"items,omitempty"`
}

type DescribeSoftphoneLayoutItem struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DescribeSoftphoneLayoutItem"`

	ItemApiName string `xml:"itemApiName,omitempty"`
}

type DescribeCompactLayoutsResult struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DescribeCompactLayoutsResult"`

	CompactLayouts                  []*DescribeCompactLayout          `xml:"compactLayouts,omitempty"`
	DefaultCompactLayoutId          *ID                               `xml:"defaultCompactLayoutId,omitempty"`
	RecordTypeCompactLayoutMappings []*RecordTypeCompactLayoutMapping `xml:"recordTypeCompactLayoutMappings,omitempty"`
}

type DescribeCompactLayout struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DescribeCompactLayout"`

	Actions    []*DescribeLayoutButton `xml:"actions,omitempty"`
	FieldItems []*DescribeLayoutItem   `xml:"fieldItems,omitempty"`
	Id         *ID                     `xml:"id,omitempty"`
	ImageItems []*DescribeLayoutItem   `xml:"imageItems,omitempty"`
	Label      string                  `xml:"label,omitempty"`
	Name       string                  `xml:"name,omitempty"`
	ObjectType string                  `xml:"objectType,omitempty"`
}

type RecordTypeCompactLayoutMapping struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com RecordTypeCompactLayoutMapping"`

	Available         bool   `xml:"available,omitempty"`
	CompactLayoutId   *ID    `xml:"compactLayoutId,omitempty"`
	CompactLayoutName string `xml:"compactLayoutName,omitempty"`
	RecordTypeId      *ID    `xml:"recordTypeId,omitempty"`
	RecordTypeName    string `xml:"recordTypeName,omitempty"`
}

type DescribePathAssistantsResult struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DescribePathAssistantsResult"`

	PathAssistants []*DescribePathAssistant `xml:"pathAssistants,omitempty"`
}

type DescribePathAssistant struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DescribePathAssistant"`

	Active                 bool                         `xml:"active,omitempty"`
	ApiName                string                       `xml:"apiName,omitempty"`
	Label                  string                       `xml:"label,omitempty"`
	PathPicklistField      string                       `xml:"pathPicklistField,omitempty"`
	PicklistsForRecordType []*PicklistForRecordType     `xml:"picklistsForRecordType,omitempty"`
	RecordTypeId           *ID                          `xml:"recordTypeId,omitempty"`
	Steps                  []*DescribePathAssistantStep `xml:"steps,omitempty"`
}

type DescribePathAssistantStep struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DescribePathAssistantStep"`

	Closed        bool                          `xml:"closed,omitempty"`
	Converted     bool                          `xml:"converted,omitempty"`
	Fields        []*DescribePathAssistantField `xml:"fields,omitempty"`
	Info          string                        `xml:"info,omitempty"`
	LayoutSection *DescribeLayoutSection        `xml:"layoutSection,omitempty"`
	PicklistLabel string                        `xml:"picklistLabel,omitempty"`
	PicklistValue string                        `xml:"picklistValue,omitempty"`
	Won           bool                          `xml:"won,omitempty"`
}

type DescribePathAssistantField struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DescribePathAssistantField"`

	ApiName  string `xml:"apiName,omitempty"`
	Label    string `xml:"label,omitempty"`
	ReadOnly bool   `xml:"readOnly,omitempty"`
	Required bool   `xml:"required,omitempty"`
}

type DescribeApprovalLayoutResult struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DescribeApprovalLayoutResult"`

	ApprovalLayouts []*DescribeApprovalLayout `xml:"approvalLayouts,omitempty"`
}

type DescribeApprovalLayout struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DescribeApprovalLayout"`

	Id          *ID                   `xml:"id,omitempty"`
	Label       string                `xml:"label,omitempty"`
	LayoutItems []*DescribeLayoutItem `xml:"layoutItems,omitempty"`
	Name        string                `xml:"name,omitempty"`
}

type DescribeLayoutResult struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DescribeLayoutResult"`

	Layouts                    []*DescribeLayout    `xml:"layouts,omitempty"`
	RecordTypeMappings         []*RecordTypeMapping `xml:"recordTypeMappings,omitempty"`
	RecordTypeSelectorRequired bool                 `xml:"recordTypeSelectorRequired,omitempty"`
}

type DescribeLayout struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DescribeLayout"`

	ButtonLayoutSection          *DescribeLayoutButtonSection   `xml:"buttonLayoutSection,omitempty"`
	DetailLayoutSections         []*DescribeLayoutSection       `xml:"detailLayoutSections,omitempty"`
	EditLayoutSections           []*DescribeLayoutSection       `xml:"editLayoutSections,omitempty"`
	FeedView                     *DescribeLayoutFeedView        `xml:"feedView,omitempty"`
	HighlightsPanelLayoutSection *DescribeLayoutSection         `xml:"highlightsPanelLayoutSection,omitempty"`
	Id                           *ID                            `xml:"id,omitempty"`
	QuickActionList              *DescribeQuickActionListResult `xml:"quickActionList,omitempty"`
	RelatedContent               *RelatedContent                `xml:"relatedContent,omitempty"`
	RelatedLists                 []*RelatedList                 `xml:"relatedLists,omitempty"`
	SaveOptions                  []*DescribeLayoutSaveOption    `xml:"saveOptions,omitempty"`
}

type DescribeQuickActionListResult struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DescribeQuickActionListResult"`

	QuickActionListItems []*DescribeQuickActionListItemResult `xml:"quickActionListItems,omitempty"`
}

type DescribeQuickActionListItemResult struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DescribeQuickActionListItemResult"`

	AccessLevelRequired *ShareAccessLevel `xml:"accessLevelRequired,omitempty"`
	Colors              []*DescribeColor  `xml:"colors,omitempty"`
	IconUrl             string            `xml:"iconUrl,omitempty"`
	Icons               []*DescribeIcon   `xml:"icons,omitempty"`
	Label               string            `xml:"label,omitempty"`
	MiniIconUrl         string            `xml:"miniIconUrl,omitempty"`
	QuickActionName     string            `xml:"quickActionName,omitempty"`
	TargetSobjectType   string            `xml:"targetSobjectType,omitempty"`
	Type_               string            `xml:"type,omitempty"`
}

type DescribeLayoutFeedView struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DescribeLayoutFeedView"`

	FeedFilters []*DescribeLayoutFeedFilter `xml:"feedFilters,omitempty"`
}

type DescribeLayoutFeedFilter struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DescribeLayoutFeedFilter"`

	Label string                `xml:"label,omitempty"`
	Name  string                `xml:"name,omitempty"`
	Type_ *FeedLayoutFilterType `xml:"type,omitempty"`
}

type DescribeLayoutSaveOption struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DescribeLayoutSaveOption"`

	DefaultValue   bool   `xml:"defaultValue,omitempty"`
	IsDisplayed    bool   `xml:"isDisplayed,omitempty"`
	Label          string `xml:"label,omitempty"`
	Name           string `xml:"name,omitempty"`
	RestHeaderName string `xml:"restHeaderName,omitempty"`
	SoapHeaderName string `xml:"soapHeaderName,omitempty"`
}

type DescribeLayoutSection struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DescribeLayoutSection"`

	Collapsed             bool                 `xml:"collapsed,omitempty"`
	Columns               int32                `xml:"columns,omitempty"`
	Heading               string               `xml:"heading,omitempty"`
	LayoutRows            []*DescribeLayoutRow `xml:"layoutRows,omitempty"`
	LayoutSectionId       *ID                  `xml:"layoutSectionId,omitempty"`
	ParentLayoutId        *ID                  `xml:"parentLayoutId,omitempty"`
	Rows                  int32                `xml:"rows,omitempty"`
	TabOrder              *TabOrderType        `xml:"tabOrder,omitempty"`
	UseCollapsibleSection bool                 `xml:"useCollapsibleSection,omitempty"`
	UseHeading            bool                 `xml:"useHeading,omitempty"`
}

type DescribeLayoutButtonSection struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DescribeLayoutButtonSection"`

	DetailButtons []*DescribeLayoutButton `xml:"detailButtons,omitempty"`
}

type DescribeLayoutRow struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DescribeLayoutRow"`

	LayoutItems []*DescribeLayoutItem `xml:"layoutItems,omitempty"`
	NumItems    int32                 `xml:"numItems,omitempty"`
}

type DescribeLayoutItem struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DescribeLayoutItem"`

	EditableForNew    bool                       `xml:"editableForNew,omitempty"`
	EditableForUpdate bool                       `xml:"editableForUpdate,omitempty"`
	Label             string                     `xml:"label,omitempty"`
	LayoutComponents  []*DescribeLayoutComponent `xml:"layoutComponents,omitempty"`
	Placeholder       bool                       `xml:"placeholder,omitempty"`
	Required          bool                       `xml:"required,omitempty"`
}

type DescribeLayoutButton struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DescribeLayoutButton"`

	Behavior       *WebLinkWindowType `xml:"behavior,omitempty"`
	Colors         []*DescribeColor   `xml:"colors,omitempty"`
	Content        string             `xml:"content,omitempty"`
	ContentSource  *WebLinkType       `xml:"contentSource,omitempty"`
	Custom         bool               `xml:"custom,omitempty"`
	Encoding       string             `xml:"encoding,omitempty"`
	Height         int32              `xml:"height,omitempty"`
	Icons          []*DescribeIcon    `xml:"icons,omitempty"`
	Label          string             `xml:"label,omitempty"`
	Menubar        bool               `xml:"menubar,omitempty"`
	Name           string             `xml:"name,omitempty"`
	Overridden     bool               `xml:"overridden,omitempty"`
	Resizeable     bool               `xml:"resizeable,omitempty"`
	Scrollbars     bool               `xml:"scrollbars,omitempty"`
	ShowsLocation  bool               `xml:"showsLocation,omitempty"`
	ShowsStatus    bool               `xml:"showsStatus,omitempty"`
	Toolbar        bool               `xml:"toolbar,omitempty"`
	Url            string             `xml:"url,omitempty"`
	Width          int32              `xml:"width,omitempty"`
	WindowPosition *WebLinkPosition   `xml:"windowPosition,omitempty"`
}

type DescribeLayoutComponent struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DescribeLayoutComponent"`

	DisplayLines int32                `xml:"displayLines,omitempty"`
	TabOrder     int32                `xml:"tabOrder,omitempty"`
	Type_        *LayoutComponentType `xml:"type,omitempty"`
	Value        string               `xml:"value,omitempty"`
}

type FieldComponent struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com FieldComponent"`

	*DescribeLayoutComponent

	Field *Field `xml:"field,omitempty"`
}

type FieldLayoutComponent struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com FieldLayoutComponent"`

	*DescribeLayoutComponent

	Components []*DescribeLayoutComponent `xml:"components,omitempty"`
	FieldType  *FieldType                 `xml:"fieldType,omitempty"`
}

type VisualforcePage struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com VisualforcePage"`

	*DescribeLayoutComponent

	ShowLabel       bool   `xml:"showLabel,omitempty"`
	ShowScrollbars  bool   `xml:"showScrollbars,omitempty"`
	SuggestedHeight string `xml:"suggestedHeight,omitempty"`
	SuggestedWidth  string `xml:"suggestedWidth,omitempty"`
	Url             string `xml:"url,omitempty"`
}

type Canvas struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com Canvas"`

	*DescribeLayoutComponent

	DisplayLocation string `xml:"displayLocation,omitempty"`
	ReferenceId     string `xml:"referenceId,omitempty"`
	ShowLabel       bool   `xml:"showLabel,omitempty"`
	ShowScrollbars  bool   `xml:"showScrollbars,omitempty"`
	SuggestedHeight string `xml:"suggestedHeight,omitempty"`
	SuggestedWidth  string `xml:"suggestedWidth,omitempty"`
}

type ReportChartComponent struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com ReportChartComponent"`

	*DescribeLayoutComponent

	CacheData              bool             `xml:"cacheData,omitempty"`
	ContextFilterableField string           `xml:"contextFilterableField,omitempty"`
	Error                  string           `xml:"error,omitempty"`
	HideOnError            bool             `xml:"hideOnError,omitempty"`
	IncludeContext         bool             `xml:"includeContext,omitempty"`
	ShowTitle              bool             `xml:"showTitle,omitempty"`
	Size                   *ReportChartSize `xml:"size,omitempty"`
}

type AnalyticsCloudComponent struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com AnalyticsCloudComponent"`

	*DescribeLayoutComponent

	Error       string `xml:"error,omitempty"`
	Filter      string `xml:"filter,omitempty"`
	Height      string `xml:"height,omitempty"`
	HideOnError bool   `xml:"hideOnError,omitempty"`
	ShowSharing bool   `xml:"showSharing,omitempty"`
	ShowTitle   bool   `xml:"showTitle,omitempty"`
	Width       string `xml:"width,omitempty"`
}

type CustomLinkComponent struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com CustomLinkComponent"`

	*DescribeLayoutComponent

	CustomLink *DescribeLayoutButton `xml:"customLink,omitempty"`
}

type NamedLayoutInfo struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com NamedLayoutInfo"`

	Name string `xml:"name,omitempty"`
}

type RecordTypeInfo struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com RecordTypeInfo"`

	Active                   bool   `xml:"active,omitempty"`
	Available                bool   `xml:"available,omitempty"`
	DefaultRecordTypeMapping bool   `xml:"defaultRecordTypeMapping,omitempty"`
	DeveloperName            string `xml:"developerName,omitempty"`
	Master                   bool   `xml:"master,omitempty"`
	Name                     string `xml:"name,omitempty"`
	RecordTypeId             *ID    `xml:"recordTypeId,omitempty"`
}

type RecordTypeMapping struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com RecordTypeMapping"`

	Active                   bool                     `xml:"active,omitempty"`
	Available                bool                     `xml:"available,omitempty"`
	DefaultRecordTypeMapping bool                     `xml:"defaultRecordTypeMapping,omitempty"`
	DeveloperName            string                   `xml:"developerName,omitempty"`
	LayoutId                 *ID                      `xml:"layoutId,omitempty"`
	Master                   bool                     `xml:"master,omitempty"`
	Name                     string                   `xml:"name,omitempty"`
	PicklistsForRecordType   []*PicklistForRecordType `xml:"picklistsForRecordType,omitempty"`
	RecordTypeId             *ID                      `xml:"recordTypeId,omitempty"`
}

type PicklistForRecordType struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com PicklistForRecordType"`

	PicklistName   string           `xml:"picklistName,omitempty"`
	PicklistValues []*PicklistEntry `xml:"picklistValues,omitempty"`
}

type RelatedContent struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com RelatedContent"`

	RelatedContentItems []*DescribeRelatedContentItem `xml:"relatedContentItems,omitempty"`
}

type DescribeRelatedContentItem struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DescribeRelatedContentItem"`

	DescribeLayoutItem *DescribeLayoutItem `xml:"describeLayoutItem,omitempty"`
}

type RelatedList struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com RelatedList"`

	AccessLevelRequiredForCreate *ShareAccessLevel       `xml:"accessLevelRequiredForCreate,omitempty"`
	Buttons                      []*DescribeLayoutButton `xml:"buttons,omitempty"`
	Columns                      []*RelatedListColumn    `xml:"columns,omitempty"`
	Custom                       bool                    `xml:"custom,omitempty"`
	Field                        string                  `xml:"field,omitempty"`
	Label                        string                  `xml:"label,omitempty"`
	LimitRows                    int32                   `xml:"limitRows,omitempty"`
	Name                         string                  `xml:"name,omitempty"`
	Sobject                      string                  `xml:"sobject,omitempty"`
	Sort                         []*RelatedListSort      `xml:"sort,omitempty"`
}

type RelatedListColumn struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com RelatedListColumn"`

	Field        string `xml:"field,omitempty"`
	FieldApiName string `xml:"fieldApiName,omitempty"`
	Format       string `xml:"format,omitempty"`
	Label        string `xml:"label,omitempty"`
	LookupId     string `xml:"lookupId,omitempty"`
	Name         string `xml:"name,omitempty"`
	Sortable     bool   `xml:"sortable,omitempty"`
}

type RelatedListSort struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com RelatedListSort"`

	Ascending bool   `xml:"ascending,omitempty"`
	Column    string `xml:"column,omitempty"`
}

type EmailFileAttachment struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com EmailFileAttachment"`

	Body        []byte `xml:"body,omitempty"`
	ContentType string `xml:"contentType,omitempty"`
	FileName    string `xml:"fileName,omitempty"`
	Id          *ID    `xml:"id,omitempty"`
	Inline      bool   `xml:"inline,omitempty"`
}

type Email struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com Email"`

	BccSender         bool           `xml:"bccSender,omitempty"`
	EmailPriority     *EmailPriority `xml:"emailPriority,omitempty"`
	ReplyTo           string         `xml:"replyTo,omitempty"`
	SaveAsActivity    bool           `xml:"saveAsActivity,omitempty"`
	SenderDisplayName string         `xml:"senderDisplayName,omitempty"`
	Subject           string         `xml:"subject,omitempty"`
	UseSignature      bool           `xml:"useSignature,omitempty"`
}

type MassEmailMessage struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com MassEmailMessage"`

	*Email

	Description     string `xml:"description,omitempty"`
	TargetObjectIds *ID    `xml:"targetObjectIds,omitempty"`
	TemplateId      *ID    `xml:"templateId,omitempty"`
	WhatIds         *ID    `xml:"whatIds,omitempty"`
}

type SingleEmailMessage struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com SingleEmailMessage"`

	*Email

	BccAddresses                 string                 `xml:"bccAddresses,omitempty"`
	CcAddresses                  string                 `xml:"ccAddresses,omitempty"`
	Charset                      string                 `xml:"charset,omitempty"`
	DocumentAttachments          []*ID                  `xml:"documentAttachments,omitempty"`
	EntityAttachments            []*ID                  `xml:"entityAttachments,omitempty"`
	FileAttachments              []*EmailFileAttachment `xml:"fileAttachments,omitempty"`
	HtmlBody                     string                 `xml:"htmlBody,omitempty"`
	InReplyTo                    string                 `xml:"inReplyTo,omitempty"`
	OptOutPolicy                 *SendEmailOptOutPolicy `xml:"optOutPolicy,omitempty"`
	OrgWideEmailAddressId        *ID                    `xml:"orgWideEmailAddressId,omitempty"`
	PlainTextBody                string                 `xml:"plainTextBody,omitempty"`
	References                   string                 `xml:"references,omitempty"`
	TargetObjectId               *ID                    `xml:"targetObjectId,omitempty"`
	TemplateId                   *ID                    `xml:"templateId,omitempty"`
	TemplateName                 string                 `xml:"templateName,omitempty"`
	ToAddresses                  string                 `xml:"toAddresses,omitempty"`
	TreatBodiesAsTemplate        bool                   `xml:"treatBodiesAsTemplate,omitempty"`
	TreatTargetObjectAsRecipient bool                   `xml:"treatTargetObjectAsRecipient,omitempty"`
	WhatId                       *ID                    `xml:"whatId,omitempty"`
}

type SendEmailResult struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com SendEmailResult"`

	Errors  []*SendEmailError `xml:"errors,omitempty"`
	Success bool              `xml:"success,omitempty"`
}

type ListViewColumn struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com ListViewColumn"`

	AscendingLabel  string            `xml:"ascendingLabel,omitempty"`
	DescendingLabel string            `xml:"descendingLabel,omitempty"`
	FieldNameOrPath string            `xml:"fieldNameOrPath,omitempty"`
	Hidden          bool              `xml:"hidden,omitempty"`
	Label           string            `xml:"label,omitempty"`
	Searchable      bool              `xml:"searchable,omitempty"`
	SelectListItem  string            `xml:"selectListItem,omitempty"`
	SortDirection   *OrderByDirection `xml:"sortDirection,omitempty"`
	SortIndex       int32             `xml:"sortIndex,omitempty"`
	Sortable        bool              `xml:"sortable,omitempty"`
	Type_           *FieldType        `xml:"type,omitempty"`
}

type ListViewOrderBy struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com ListViewOrderBy"`

	FieldNameOrPath string                `xml:"fieldNameOrPath,omitempty"`
	NullsPosition   *OrderByNullsPosition `xml:"nullsPosition,omitempty"`
	SortDirection   *OrderByDirection     `xml:"sortDirection,omitempty"`
}

type DescribeSoqlListView struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DescribeSoqlListView"`

	Columns         []*ListViewColumn   `xml:"columns,omitempty"`
	Id              *ID                 `xml:"id,omitempty"`
	OrderBy         []*ListViewOrderBy  `xml:"orderBy,omitempty"`
	Query           string              `xml:"query,omitempty"`
	RelatedEntityId *ID                 `xml:"relatedEntityId,omitempty"`
	Scope           string              `xml:"scope,omitempty"`
	ScopeEntityId   *ID                 `xml:"scopeEntityId,omitempty"`
	SobjectType     string              `xml:"sobjectType,omitempty"`
	WhereCondition  *SoqlWhereCondition `xml:"whereCondition,omitempty"`
}

type DescribeSoqlListViewsRequest struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DescribeSoqlListViewsRequest"`

	ListViewParams []*DescribeSoqlListViewParams `xml:"listViewParams,omitempty"`
}

type DescribeSoqlListViewParams struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DescribeSoqlListViewParams"`

	DeveloperNameOrId string `xml:"developerNameOrId,omitempty"`
	SobjectType       string `xml:"sobjectType,omitempty"`
}

type DescribeSoqlListViewResult struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DescribeSoqlListViewResult"`

	DescribeSoqlListViews []*DescribeSoqlListView `xml:"describeSoqlListViews,omitempty"`
}

type ExecuteListViewRequest struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com ExecuteListViewRequest"`

	DeveloperNameOrId string             `xml:"developerNameOrId,omitempty"`
	Limit             int32              `xml:"limit,omitempty"`
	Offset            int32              `xml:"offset,omitempty"`
	OrderBy           []*ListViewOrderBy `xml:"orderBy,omitempty"`
	SobjectType       string             `xml:"sobjectType,omitempty"`
}

type ExecuteListViewResult struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com ExecuteListViewResult"`

	Columns       []*ListViewColumn `xml:"columns,omitempty"`
	DeveloperName string            `xml:"developerName,omitempty"`
	Done          bool              `xml:"done,omitempty"`
	Id            *ID               `xml:"id,omitempty"`
	Label         string            `xml:"label,omitempty"`
	Records       []*ListViewRecord `xml:"records,omitempty"`
	Size          int32             `xml:"size,omitempty"`
}

type ListViewRecord struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com ListViewRecord"`

	Columns []*ListViewRecordColumn `xml:"columns,omitempty"`
}

type ListViewRecordColumn struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com ListViewRecordColumn"`

	FieldNameOrPath string `xml:"fieldNameOrPath,omitempty"`
	Value           string `xml:"value,omitempty"`
}

type SoqlWhereCondition struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com SoqlWhereCondition"`
}

type SoqlCondition struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com SoqlCondition"`

	*SoqlWhereCondition

	Field    string        `xml:"field,omitempty"`
	Operator *SoqlOperator `xml:"operator,omitempty"`
	Values   []string      `xml:"values,omitempty"`
}

type SoqlNotCondition struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com SoqlNotCondition"`

	*SoqlWhereCondition

	Condition *SoqlWhereCondition `xml:"condition,omitempty"`
}

type SoqlConditionGroup struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com SoqlConditionGroup"`

	*SoqlWhereCondition

	Conditions  []*SoqlWhereCondition `xml:"conditions,omitempty"`
	Conjunction *SoqlConjunction      `xml:"conjunction,omitempty"`
}

type SoqlSubQueryCondition struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com SoqlSubQueryCondition"`

	*SoqlWhereCondition

	Field    string        `xml:"field,omitempty"`
	Operator *SoqlOperator `xml:"operator,omitempty"`
	SubQuery string        `xml:"subQuery,omitempty"`
}

type DescribeSearchLayoutResult struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DescribeSearchLayoutResult"`

	ErrorMsg      string            `xml:"errorMsg,omitempty"`
	Label         string            `xml:"label,omitempty"`
	LimitRows     int32             `xml:"limitRows,omitempty"`
	ObjectType    string            `xml:"objectType,omitempty"`
	SearchColumns []*DescribeColumn `xml:"searchColumns,omitempty"`
}

type DescribeColumn struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DescribeColumn"`

	Field  string `xml:"field,omitempty"`
	Format string `xml:"format,omitempty"`
	Label  string `xml:"label,omitempty"`
	Name   string `xml:"name,omitempty"`
}

type DescribeSearchScopeOrderResult struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DescribeSearchScopeOrderResult"`

	KeyPrefix string `xml:"keyPrefix,omitempty"`
	Name      string `xml:"name,omitempty"`
}

type DescribeSearchableEntityResult struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DescribeSearchableEntityResult"`

	Label       string `xml:"label,omitempty"`
	Name        string `xml:"name,omitempty"`
	PluralLabel string `xml:"pluralLabel,omitempty"`
}

type DescribeTabSetResult struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DescribeTabSetResult"`

	Description string         `xml:"description,omitempty"`
	Label       string         `xml:"label,omitempty"`
	LogoUrl     string         `xml:"logoUrl,omitempty"`
	Namespace   string         `xml:"namespace,omitempty"`
	Selected    bool           `xml:"selected,omitempty"`
	TabSetId    string         `xml:"tabSetId,omitempty"`
	Tabs        []*DescribeTab `xml:"tabs,omitempty"`
}

type DescribeTab struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DescribeTab"`

	Colors      []*DescribeColor `xml:"colors,omitempty"`
	Custom      bool             `xml:"custom,omitempty"`
	IconUrl     string           `xml:"iconUrl,omitempty"`
	Icons       []*DescribeIcon  `xml:"icons,omitempty"`
	Label       string           `xml:"label,omitempty"`
	MiniIconUrl string           `xml:"miniIconUrl,omitempty"`
	Name        string           `xml:"name,omitempty"`
	SobjectName string           `xml:"sobjectName,omitempty"`
	Url         string           `xml:"url,omitempty"`
}

type DescribeColor struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DescribeColor"`

	Color   string `xml:"color,omitempty"`
	Context string `xml:"context,omitempty"`
	Theme   string `xml:"theme,omitempty"`
}

type DescribeIcon struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com DescribeIcon"`

	ContentType string `xml:"contentType,omitempty"`
	Height      int32  `xml:"height,omitempty"`
	Theme       string `xml:"theme,omitempty"`
	Url         string `xml:"url,omitempty"`
	Width       int32  `xml:"width,omitempty"`
}

type ActionOverride struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com ActionOverride"`

	FormFactor         string `xml:"formFactor,omitempty"`
	IsAvailableInTouch bool   `xml:"isAvailableInTouch,omitempty"`
	Name               string `xml:"name,omitempty"`
	PageId             *ID    `xml:"pageId,omitempty"`
	Url                string `xml:"url,omitempty"`
}

type RenderEmailTemplateRequest struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com RenderEmailTemplateRequest"`

	EscapeHtmlInMergeFields bool   `xml:"escapeHtmlInMergeFields,omitempty"`
	TemplateBodies          string `xml:"templateBodies,omitempty"`
	WhatId                  *ID    `xml:"whatId,omitempty"`
	WhoId                   *ID    `xml:"whoId,omitempty"`
}

type RenderEmailTemplateBodyResult struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com RenderEmailTemplateBodyResult"`

	Errors     []*RenderEmailTemplateError `xml:"errors,omitempty"`
	MergedBody string                      `xml:"mergedBody,omitempty"`
	Success    bool                        `xml:"success,omitempty"`
}

type RenderEmailTemplateResult struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com RenderEmailTemplateResult"`

	BodyResults *RenderEmailTemplateBodyResult `xml:"bodyResults,omitempty"`
	Errors      []*Error                       `xml:"errors,omitempty"`
	Success     bool                           `xml:"success,omitempty"`
}

type RenderStoredEmailTemplateRequest struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com RenderStoredEmailTemplateRequest"`

	AttachmentRetrievalOption *AttachmentRetrievalOption `xml:"attachmentRetrievalOption,omitempty"`
	TemplateId                *ID                        `xml:"templateId,omitempty"`
	UpdateTemplateUsage       bool                       `xml:"updateTemplateUsage,omitempty"`
	WhatId                    *ID                        `xml:"whatId,omitempty"`
	WhoId                     *ID                        `xml:"whoId,omitempty"`
}

type RenderStoredEmailTemplateResult struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com RenderStoredEmailTemplateResult"`

	Errors        []*Error            `xml:"errors,omitempty"`
	RenderedEmail *SingleEmailMessage `xml:"renderedEmail,omitempty"`
	Success       bool                `xml:"success,omitempty"`
}

type LogInfo struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com LogInfo"`

	Category *LogCategory      `xml:"category,omitempty"`
	Level    *LogCategoryLevel `xml:"level,omitempty"`
}

type PackageVersion struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com PackageVersion"`

	MajorNumber int32  `xml:"majorNumber,omitempty"`
	MinorNumber int32  `xml:"minorNumber,omitempty"`
	Namespace   string `xml:"namespace,omitempty"`
}

type LimitInfo struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com LimitInfo"`

	Current int32  `xml:"current,omitempty"`
	Limit   int32  `xml:"limit,omitempty"`
	Type_   string `xml:"type,omitempty"`
}

type OwnerChangeOption struct {
	XMLName xml.Name `xml:"urn:enterprise.soap.sforce.com OwnerChangeOption"`

	Type_   *OwnerChangeOptionType `xml:"type,omitempty"`
	Execute bool                   `xml:"execute,omitempty"`
}

type ExceptionCode string

const (
	ExceptionCodeAPEXRESTSERVICESDISABLED ExceptionCode = "APEXRESTSERVICESDISABLED"

	ExceptionCodeAPEXTRIGGERCOUPLINGLIMIT ExceptionCode = "APEXTRIGGERCOUPLINGLIMIT"

	ExceptionCodeAPICURRENTLYDISABLED ExceptionCode = "APICURRENTLYDISABLED"

	ExceptionCodeAPIDISABLEDFORORG ExceptionCode = "APIDISABLEDFORORG"

	ExceptionCodeARGUMENTOBJECTPARSEERROR ExceptionCode = "ARGUMENTOBJECTPARSEERROR"

	ExceptionCodeASYNCOPERATIONLOCATOR ExceptionCode = "ASYNCOPERATIONLOCATOR"

	ExceptionCodeASYNCQUERYUNSUPPORTEDQUERY ExceptionCode = "ASYNCQUERYUNSUPPORTEDQUERY"

	ExceptionCodeBATCHPROCESSINGHALTED ExceptionCode = "BATCHPROCESSINGHALTED"

	ExceptionCodeBIGOBJECTUNSUPPORTEDOPERATION ExceptionCode = "BIGOBJECTUNSUPPORTEDOPERATION"

	ExceptionCodeCANNOTDELETEENTITY ExceptionCode = "CANNOTDELETEENTITY"

	ExceptionCodeCANNOTDELETEOWNER ExceptionCode = "CANNOTDELETEOWNER"

	ExceptionCodeCANTADDSTANDADRDPORTALUSERTOTERRITORY ExceptionCode = "CANTADDSTANDADRDPORTALUSERTOTERRITORY"

	ExceptionCodeCANTADDSTANDARDPORTALUSERTOTERRITORY ExceptionCode = "CANTADDSTANDARDPORTALUSERTOTERRITORY"

	ExceptionCodeCIRCULAROBJECTGRAPH ExceptionCode = "CIRCULAROBJECTGRAPH"

	ExceptionCodeCLIENTNOTACCESSIBLEFORUSER ExceptionCode = "CLIENTNOTACCESSIBLEFORUSER"

	ExceptionCodeCLIENTREQUIREUPDATEFORUSER ExceptionCode = "CLIENTREQUIREUPDATEFORUSER"

	ExceptionCodeCONTENTALREADYANASSETEXCEPTION ExceptionCode = "CONTENTALREADYANASSETEXCEPTION"

	ExceptionCodeCONTENTCUSTOMDOWNLOADEXCEPTION ExceptionCode = "CONTENTCUSTOMDOWNLOADEXCEPTION"

	ExceptionCodeCONTENTHUBAUTHENTICATIONEXCEPTION ExceptionCode = "CONTENTHUBAUTHENTICATIONEXCEPTION"

	ExceptionCodeCONTENTHUBFILEDOWNLOADEXCEPTION ExceptionCode = "CONTENTHUBFILEDOWNLOADEXCEPTION"

	ExceptionCodeCONTENTHUBFILEHASNOURLEXCEPTION ExceptionCode = "CONTENTHUBFILEHASNOURLEXCEPTION"

	ExceptionCodeCONTENTHUBFILENOTFOUNDEXCEPTION ExceptionCode = "CONTENTHUBFILENOTFOUNDEXCEPTION"

	ExceptionCodeCONTENTHUBINVALIDOBJECTTYPEEXCEPTION ExceptionCode = "CONTENTHUBINVALIDOBJECTTYPEEXCEPTION"

	ExceptionCodeCONTENTHUBINVALIDPAGENUMBEREXCEPTION ExceptionCode = "CONTENTHUBINVALIDPAGENUMBEREXCEPTION"

	ExceptionCodeCONTENTHUBINVALIDPAYLOAD ExceptionCode = "CONTENTHUBINVALIDPAYLOAD"

	ExceptionCodeCONTENTHUBINVALIDRENDITIONPAGENUMBEREXCEPTION ExceptionCode = "CONTENTHUBINVALIDRENDITIONPAGENUMBEREXCEPTION"

	ExceptionCodeCONTENTHUBITEMTYPENOTFOUNDEXCEPTION ExceptionCode = "CONTENTHUBITEMTYPENOTFOUNDEXCEPTION"

	ExceptionCodeCONTENTHUBOBJECTNOTFOUNDEXCEPTION ExceptionCode = "CONTENTHUBOBJECTNOTFOUNDEXCEPTION"

	ExceptionCodeCONTENTHUBOPERATIONNOTSUPPORTEDEXCEPTION ExceptionCode = "CONTENTHUBOPERATIONNOTSUPPORTEDEXCEPTION"

	ExceptionCodeCONTENTHUBSECURITYEXCEPTION ExceptionCode = "CONTENTHUBSECURITYEXCEPTION"

	ExceptionCodeCONTENTHUBTIMEOUTEXCEPTION ExceptionCode = "CONTENTHUBTIMEOUTEXCEPTION"

	ExceptionCodeCONTENTHUBUNEXPECTEDEXCEPTION ExceptionCode = "CONTENTHUBUNEXPECTEDEXCEPTION"

	ExceptionCodeCONTENTIMAGESCALINGINVALIDARGUMENTSEXCEPTION ExceptionCode = "CONTENTIMAGESCALINGINVALIDARGUMENTSEXCEPTION"

	ExceptionCodeCONTENTIMAGESCALINGINVALIDIMAGEEXCEPTION ExceptionCode = "CONTENTIMAGESCALINGINVALIDIMAGEEXCEPTION"

	ExceptionCodeCONTENTIMAGESCALINGMAXRENDITIONSEXCEPTION ExceptionCode = "CONTENTIMAGESCALINGMAXRENDITIONSEXCEPTION"

	ExceptionCodeCONTENTIMAGESCALINGTIMEOUTEXCEPTION ExceptionCode = "CONTENTIMAGESCALINGTIMEOUTEXCEPTION"

	ExceptionCodeCONTENTIMAGESCALINGUNKNOWNEXCEPTION ExceptionCode = "CONTENTIMAGESCALINGUNKNOWNEXCEPTION"

	ExceptionCodeCUSTOMMETADATALIMITEXCEEDED ExceptionCode = "CUSTOMMETADATALIMITEXCEEDED"

	ExceptionCodeCUSTOMSETTINGSLIMITEXCEEDED ExceptionCode = "CUSTOMSETTINGSLIMITEXCEEDED"

	ExceptionCodeDATACLOUDAPICLIENTEXCEPTION ExceptionCode = "DATACLOUDAPICLIENTEXCEPTION"

	ExceptionCodeDATACLOUDAPIDISABLEDEXCEPTION ExceptionCode = "DATACLOUDAPIDISABLEDEXCEPTION"

	ExceptionCodeDATACLOUDAPIINVALIDQUERYEXCEPTION ExceptionCode = "DATACLOUDAPIINVALIDQUERYEXCEPTION"

	ExceptionCodeDATACLOUDAPISERVERBUSYEXCEPTION ExceptionCode = "DATACLOUDAPISERVERBUSYEXCEPTION"

	ExceptionCodeDATACLOUDAPISERVEREXCEPTION ExceptionCode = "DATACLOUDAPISERVEREXCEPTION"

	ExceptionCodeDATACLOUDAPITIMEOUTEXCEPTION ExceptionCode = "DATACLOUDAPITIMEOUTEXCEPTION"

	ExceptionCodeDATACLOUDAPIUNAVAILABLE ExceptionCode = "DATACLOUDAPIUNAVAILABLE"

	ExceptionCodeDATAINTEGRATIONINPUTERROR ExceptionCode = "DATAINTEGRATIONINPUTERROR"

	ExceptionCodeDATAINTEGRATIONNOTFOUND ExceptionCode = "DATAINTEGRATIONNOTFOUND"

	ExceptionCodeDATAINTEGRATIONNOACCESS ExceptionCode = "DATAINTEGRATIONNOACCESS"

	ExceptionCodeDATAINTEGRATIONPACKAGEERROR ExceptionCode = "DATAINTEGRATIONPACKAGEERROR"

	ExceptionCodeDATAINTEGRATIONSERVERERROR ExceptionCode = "DATAINTEGRATIONSERVERERROR"

	ExceptionCodeDATAINTEGRATIONVENDORSETUPERROR ExceptionCode = "DATAINTEGRATIONVENDORSETUPERROR"

	ExceptionCodeDUPLICATEARGUMENTVALUE ExceptionCode = "DUPLICATEARGUMENTVALUE"

	ExceptionCodeDUPLICATEVALUE ExceptionCode = "DUPLICATEVALUE"

	ExceptionCodeEMAILBATCHSIZELIMITEXCEEDED ExceptionCode = "EMAILBATCHSIZELIMITEXCEEDED"

	ExceptionCodeEMAILTOCASEINVALIDROUTING ExceptionCode = "EMAILTOCASEINVALIDROUTING"

	ExceptionCodeEMAILTOCASELIMITEXCEEDED ExceptionCode = "EMAILTOCASELIMITEXCEEDED"

	ExceptionCodeEMAILTOCASENOTENABLED ExceptionCode = "EMAILTOCASENOTENABLED"

	ExceptionCodeENTITYNOTQUERYABLE ExceptionCode = "ENTITYNOTQUERYABLE"

	ExceptionCodeENVIRONMENTHUBMEMBERSHIPCONFLICT ExceptionCode = "ENVIRONMENTHUBMEMBERSHIPCONFLICT"

	ExceptionCodeEXCEEDEDIDLIMIT ExceptionCode = "EXCEEDEDIDLIMIT"

	ExceptionCodeEXCEEDEDLEADCONVERTLIMIT ExceptionCode = "EXCEEDEDLEADCONVERTLIMIT"

	ExceptionCodeEXCEEDEDMAXFILTERENTITIES ExceptionCode = "EXCEEDEDMAXFILTERENTITIES"

	ExceptionCodeEXCEEDEDMAXSIZEREQUEST ExceptionCode = "EXCEEDEDMAXSIZEREQUEST"

	ExceptionCodeEXCEEDEDMAXSOBJECTS ExceptionCode = "EXCEEDEDMAXSOBJECTS"

	ExceptionCodeEXCEEDEDMAXTYPESLIMIT ExceptionCode = "EXCEEDEDMAXTYPESLIMIT"

	ExceptionCodeEXCEEDEDQUOTA ExceptionCode = "EXCEEDEDQUOTA"

	ExceptionCodeEXTERNALOBJECTAUTHENTICATIONEXCEPTION ExceptionCode = "EXTERNALOBJECTAUTHENTICATIONEXCEPTION"

	ExceptionCodeEXTERNALOBJECTCONNECTIONEXCEPTION ExceptionCode = "EXTERNALOBJECTCONNECTIONEXCEPTION"

	ExceptionCodeEXTERNALOBJECTEXCEPTION ExceptionCode = "EXTERNALOBJECTEXCEPTION"

	ExceptionCodeEXTERNALOBJECTUNSUPPORTEDEXCEPTION ExceptionCode = "EXTERNALOBJECTUNSUPPORTEDEXCEPTION"

	ExceptionCodeFEDERATEDSEARCHERROR ExceptionCode = "FEDERATEDSEARCHERROR"

	ExceptionCodeFEEDNOTENABLEDFOROBJECT ExceptionCode = "FEEDNOTENABLEDFOROBJECT"

	ExceptionCodeFUNCTIONALITYNOTENABLED ExceptionCode = "FUNCTIONALITYNOTENABLED"

	ExceptionCodeFUNCTIONALITYTEMPORARILYUNAVAILABLE ExceptionCode = "FUNCTIONALITYTEMPORARILYUNAVAILABLE"

	ExceptionCodeIDREQUIRED ExceptionCode = "IDREQUIRED"

	ExceptionCodeILLEGALQUERYPARAMETERVALUE ExceptionCode = "ILLEGALQUERYPARAMETERVALUE"

	ExceptionCodeINACTIVEOWNERORUSER ExceptionCode = "INACTIVEOWNERORUSER"

	ExceptionCodeINACTIVEPORTAL ExceptionCode = "INACTIVEPORTAL"

	ExceptionCodeINSERTUPDATEDELETENOTALLOWEDDURINGMAINTENANCE ExceptionCode = "INSERTUPDATEDELETENOTALLOWEDDURINGMAINTENANCE"

	ExceptionCodeINSTALLKEYINVALID ExceptionCode = "INSTALLKEYINVALID"

	ExceptionCodeINSTALLKEYREQUIRED ExceptionCode = "INSTALLKEYREQUIRED"

	ExceptionCodeINSUFFICIENTACCESS ExceptionCode = "INSUFFICIENTACCESS"

	ExceptionCodeINSUFFICIENTACCESSAPEXMETADATADEPLOY ExceptionCode = "INSUFFICIENTACCESSAPEXMETADATADEPLOY"

	ExceptionCodeINTERNALCANVASERROR ExceptionCode = "INTERNALCANVASERROR"

	ExceptionCodeINVALIDASSIGNMENTRULE ExceptionCode = "INVALIDASSIGNMENTRULE"

	ExceptionCodeINVALIDAUTHHEADER ExceptionCode = "INVALIDAUTHHEADER"

	ExceptionCodeINVALIDBATCHREQUEST ExceptionCode = "INVALIDBATCHREQUEST"

	ExceptionCodeINVALIDBATCHSIZE ExceptionCode = "INVALIDBATCHSIZE"

	ExceptionCodeINVALIDCLIENT ExceptionCode = "INVALIDCLIENT"

	ExceptionCodeINVALIDCROSSREFERENCEKEY ExceptionCode = "INVALIDCROSSREFERENCEKEY"

	ExceptionCodeINVALIDDATEFORMAT ExceptionCode = "INVALIDDATEFORMAT"

	ExceptionCodeINVALIDFIELD ExceptionCode = "INVALIDFIELD"

	ExceptionCodeINVALIDFILTERLANGUAGE ExceptionCode = "INVALIDFILTERLANGUAGE"

	ExceptionCodeINVALIDFILTERVALUE ExceptionCode = "INVALIDFILTERVALUE"

	ExceptionCodeINVALIDIDFIELD ExceptionCode = "INVALIDIDFIELD"

	ExceptionCodeINVALIDINPUTCOMBINATION ExceptionCode = "INVALIDINPUTCOMBINATION"

	ExceptionCodeINVALIDLOCALELANGUAGE ExceptionCode = "INVALIDLOCALELANGUAGE"

	ExceptionCodeINVALIDLOCATOR ExceptionCode = "INVALIDLOCATOR"

	ExceptionCodeINVALIDLOGIN ExceptionCode = "INVALIDLOGIN"

	ExceptionCodeINVALIDMULTIPARTREQUEST ExceptionCode = "INVALIDMULTIPARTREQUEST"

	ExceptionCodeINVALIDNEWPASSWORD ExceptionCode = "INVALIDNEWPASSWORD"

	ExceptionCodeINVALIDOLDPASSWORD ExceptionCode = "INVALIDOLDPASSWORD"

	ExceptionCodeINVALIDOPERATION ExceptionCode = "INVALIDOPERATION"

	ExceptionCodeINVALIDOPERATIONWITHEXPIREDPASSWORD ExceptionCode = "INVALIDOPERATIONWITHEXPIREDPASSWORD"

	ExceptionCodeINVALIDPACKAGEVERSION ExceptionCode = "INVALIDPACKAGEVERSION"

	ExceptionCodeINVALIDPAGINGOPTION ExceptionCode = "INVALIDPAGINGOPTION"

	ExceptionCodeINVALIDQUERYFILTEROPERATOR ExceptionCode = "INVALIDQUERYFILTEROPERATOR"

	ExceptionCodeINVALIDQUERYLOCATOR ExceptionCode = "INVALIDQUERYLOCATOR"

	ExceptionCodeINVALIDQUERYSCOPE ExceptionCode = "INVALIDQUERYSCOPE"

	ExceptionCodeINVALIDREPLICATIONDATE ExceptionCode = "INVALIDREPLICATIONDATE"

	ExceptionCodeINVALIDSEARCH ExceptionCode = "INVALIDSEARCH"

	ExceptionCodeINVALIDSEARCHSCOPE ExceptionCode = "INVALIDSEARCHSCOPE"

	ExceptionCodeINVALIDSESSIONID ExceptionCode = "INVALIDSESSIONID"

	ExceptionCodeINVALIDSOAPHEADER ExceptionCode = "INVALIDSOAPHEADER"

	ExceptionCodeINVALIDSORTOPTION ExceptionCode = "INVALIDSORTOPTION"

	ExceptionCodeINVALIDSSOGATEWAYURL ExceptionCode = "INVALIDSSOGATEWAYURL"

	ExceptionCodeINVALIDTYPE ExceptionCode = "INVALIDTYPE"

	ExceptionCodeINVALIDTYPEFOROPERATION ExceptionCode = "INVALIDTYPEFOROPERATION"

	ExceptionCodeJIGSAWACTIONDISABLED ExceptionCode = "JIGSAWACTIONDISABLED"

	ExceptionCodeJIGSAWIMPORTLIMITEXCEEDED ExceptionCode = "JIGSAWIMPORTLIMITEXCEEDED"

	ExceptionCodeJIGSAWREQUESTNOTSUPPORTED ExceptionCode = "JIGSAWREQUESTNOTSUPPORTED"

	ExceptionCodeJSONPARSERERROR ExceptionCode = "JSONPARSERERROR"

	ExceptionCodeKEYHASBEENDESTROYED ExceptionCode = "KEYHASBEENDESTROYED"

	ExceptionCodeLICENSINGDATAERROR ExceptionCode = "LICENSINGDATAERROR"

	ExceptionCodeLICENSINGUNKNOWNERROR ExceptionCode = "LICENSINGUNKNOWNERROR"

	ExceptionCodeLIMITEXCEEDED ExceptionCode = "LIMITEXCEEDED"

	ExceptionCodeLOGINCHALLENGEISSUED ExceptionCode = "LOGINCHALLENGEISSUED"

	ExceptionCodeLOGINCHALLENGEPENDING ExceptionCode = "LOGINCHALLENGEPENDING"

	ExceptionCodeLOGINDURINGRESTRICTEDDOMAIN ExceptionCode = "LOGINDURINGRESTRICTEDDOMAIN"

	ExceptionCodeLOGINDURINGRESTRICTEDTIME ExceptionCode = "LOGINDURINGRESTRICTEDTIME"

	ExceptionCodeLOGINMUSTUSESECURITYTOKEN ExceptionCode = "LOGINMUSTUSESECURITYTOKEN"

	ExceptionCodeMALFORMEDID ExceptionCode = "MALFORMEDID"

	ExceptionCodeMALFORMEDQUERY ExceptionCode = "MALFORMEDQUERY"

	ExceptionCodeMALFORMEDSEARCH ExceptionCode = "MALFORMEDSEARCH"

	ExceptionCodeMISMATCHINGVERSIONS ExceptionCode = "MISMATCHINGVERSIONS"

	ExceptionCodeMISSINGARGUMENT ExceptionCode = "MISSINGARGUMENT"

	ExceptionCodeMISSINGRECORD ExceptionCode = "MISSINGRECORD"

	ExceptionCodeMODIFIED ExceptionCode = "MODIFIED"

	ExceptionCodeMUTUALAUTHENTICATIONFAILED ExceptionCode = "MUTUALAUTHENTICATIONFAILED"

	ExceptionCodeNOTACCEPTABLE ExceptionCode = "NOTACCEPTABLE"

	ExceptionCodeNOTMODIFIED ExceptionCode = "NOTMODIFIED"

	ExceptionCodeNOACTIVEDUPLICATERULE ExceptionCode = "NOACTIVEDUPLICATERULE"

	ExceptionCodeNORECIPIENTS ExceptionCode = "NORECIPIENTS"

	ExceptionCodeNOSOFTPHONELAYOUT ExceptionCode = "NOSOFTPHONELAYOUT"

	ExceptionCodeNUMBEROUTSIDEVALIDRANGE ExceptionCode = "NUMBEROUTSIDEVALIDRANGE"

	ExceptionCodeOPERATIONTOOLARGE ExceptionCode = "OPERATIONTOOLARGE"

	ExceptionCodeORGINMAINTENANCE ExceptionCode = "ORGINMAINTENANCE"

	ExceptionCodeORGISDOTORG ExceptionCode = "ORGISDOTORG"

	ExceptionCodeORGISSIGNINGUP ExceptionCode = "ORGISSIGNINGUP"

	ExceptionCodeORGLOCKED ExceptionCode = "ORGLOCKED"

	ExceptionCodeORGNOTOWNEDBYINSTANCE ExceptionCode = "ORGNOTOWNEDBYINSTANCE"

	ExceptionCodePASSWORDLOCKOUT ExceptionCode = "PASSWORDLOCKOUT"

	ExceptionCodePORTALNOACCESS ExceptionCode = "PORTALNOACCESS"

	ExceptionCodePOSTBODYPARSEERROR ExceptionCode = "POSTBODYPARSEERROR"

	ExceptionCodeQATHETEROGENOUSCONTEXTIDS ExceptionCode = "QATHETEROGENOUSCONTEXTIDS"

	ExceptionCodeQATINVALIDCONTEXTID ExceptionCode = "QATINVALIDCONTEXTID"

	ExceptionCodeQATINVALIDQUICKACTION ExceptionCode = "QATINVALIDQUICKACTION"

	ExceptionCodeQUERYTIMEOUT ExceptionCode = "QUERYTIMEOUT"

	ExceptionCodeQUERYTOOCOMPLICATED ExceptionCode = "QUERYTOOCOMPLICATED"

	ExceptionCodeREALTIMEPROCESSINGTIMEEXCEEDEDLIMIT ExceptionCode = "REALTIMEPROCESSINGTIMEEXCEEDEDLIMIT"

	ExceptionCodeREQUESTLIMITEXCEEDED ExceptionCode = "REQUESTLIMITEXCEEDED"

	ExceptionCodeREQUESTRUNNINGTOOLONG ExceptionCode = "REQUESTRUNNINGTOOLONG"

	ExceptionCodeSERVERUNAVAILABLE ExceptionCode = "SERVERUNAVAILABLE"

	ExceptionCodeSERVICEDESKNOTENABLED ExceptionCode = "SERVICEDESKNOTENABLED"

	ExceptionCodeSOCIALCRMFEEDSERVICEAPICLIENTEXCEPTION ExceptionCode = "SOCIALCRMFEEDSERVICEAPICLIENTEXCEPTION"

	ExceptionCodeSOCIALCRMFEEDSERVICEAPISERVEREXCEPTION ExceptionCode = "SOCIALCRMFEEDSERVICEAPISERVEREXCEPTION"

	ExceptionCodeSOCIALCRMFEEDSERVICEAPIUNAVAILABLE ExceptionCode = "SOCIALCRMFEEDSERVICEAPIUNAVAILABLE"

	ExceptionCodeSSOSERVICEDOWN ExceptionCode = "SSOSERVICEDOWN"

	ExceptionCodeSSTADMINFILEDOWNLOADEXCEPTION ExceptionCode = "SSTADMINFILEDOWNLOADEXCEPTION"

	ExceptionCodeSTATETRANSITIONNOTALLOWED ExceptionCode = "STATETRANSITIONNOTALLOWED"

	ExceptionCodeTOOMANYAPEXREQUESTS ExceptionCode = "TOOMANYAPEXREQUESTS"

	ExceptionCodeTOOMANYRECIPIENTS ExceptionCode = "TOOMANYRECIPIENTS"

	ExceptionCodeTOOMANYRECORDS ExceptionCode = "TOOMANYRECORDS"

	ExceptionCodeTRIALEXPIRED ExceptionCode = "TRIALEXPIRED"

	ExceptionCodeTXNSECURITYENDASESSION ExceptionCode = "TXNSECURITYENDASESSION"

	ExceptionCodeTXNSECURITYFAILCLOSE ExceptionCode = "TXNSECURITYFAILCLOSE"

	ExceptionCodeTXNSECURITYNOACCESS ExceptionCode = "TXNSECURITYNOACCESS"

	ExceptionCodeTXNSECURITYTWOFAREQUIRED ExceptionCode = "TXNSECURITYTWOFAREQUIRED"

	ExceptionCodeUNABLETOLOCKROW ExceptionCode = "UNABLETOLOCKROW"

	ExceptionCodeUNKNOWNATTACHMENTEXCEPTION ExceptionCode = "UNKNOWNATTACHMENTEXCEPTION"

	ExceptionCodeUNKNOWNEXCEPTION ExceptionCode = "UNKNOWNEXCEPTION"

	ExceptionCodeUNKNOWNORGSETTING ExceptionCode = "UNKNOWNORGSETTING"

	ExceptionCodeUNSUPPORTEDAPIVERSION ExceptionCode = "UNSUPPORTEDAPIVERSION"

	ExceptionCodeUNSUPPORTEDATTACHMENTENCODING ExceptionCode = "UNSUPPORTEDATTACHMENTENCODING"

	ExceptionCodeUNSUPPORTEDCLIENT ExceptionCode = "UNSUPPORTEDCLIENT"

	ExceptionCodeUNSUPPORTEDMEDIATYPE ExceptionCode = "UNSUPPORTEDMEDIATYPE"

	ExceptionCodeXMLPARSERERROR ExceptionCode = "XMLPARSERERROR"
)

type FaultCode string

const (
	FaultCodeFnsAPEXRESTSERVICESDISABLED FaultCode = "fnsAPEXRESTSERVICESDISABLED"

	FaultCodeFnsAPEXTRIGGERCOUPLINGLIMIT FaultCode = "fnsAPEXTRIGGERCOUPLINGLIMIT"

	FaultCodeFnsAPICURRENTLYDISABLED FaultCode = "fnsAPICURRENTLYDISABLED"

	FaultCodeFnsAPIDISABLEDFORORG FaultCode = "fnsAPIDISABLEDFORORG"

	FaultCodeFnsARGUMENTOBJECTPARSEERROR FaultCode = "fnsARGUMENTOBJECTPARSEERROR"

	FaultCodeFnsASYNCOPERATIONLOCATOR FaultCode = "fnsASYNCOPERATIONLOCATOR"

	FaultCodeFnsASYNCQUERYUNSUPPORTEDQUERY FaultCode = "fnsASYNCQUERYUNSUPPORTEDQUERY"

	FaultCodeFnsBATCHPROCESSINGHALTED FaultCode = "fnsBATCHPROCESSINGHALTED"

	FaultCodeFnsBIGOBJECTUNSUPPORTEDOPERATION FaultCode = "fnsBIGOBJECTUNSUPPORTEDOPERATION"

	FaultCodeFnsCANNOTDELETEENTITY FaultCode = "fnsCANNOTDELETEENTITY"

	FaultCodeFnsCANNOTDELETEOWNER FaultCode = "fnsCANNOTDELETEOWNER"

	FaultCodeFnsCANTADDSTANDADRDPORTALUSERTOTERRITORY FaultCode = "fnsCANTADDSTANDADRDPORTALUSERTOTERRITORY"

	FaultCodeFnsCANTADDSTANDARDPORTALUSERTOTERRITORY FaultCode = "fnsCANTADDSTANDARDPORTALUSERTOTERRITORY"

	FaultCodeFnsCIRCULAROBJECTGRAPH FaultCode = "fnsCIRCULAROBJECTGRAPH"

	FaultCodeFnsCLIENTNOTACCESSIBLEFORUSER FaultCode = "fnsCLIENTNOTACCESSIBLEFORUSER"

	FaultCodeFnsCLIENTREQUIREUPDATEFORUSER FaultCode = "fnsCLIENTREQUIREUPDATEFORUSER"

	FaultCodeFnsCONTENTALREADYANASSETEXCEPTION FaultCode = "fnsCONTENTALREADYANASSETEXCEPTION"

	FaultCodeFnsCONTENTCUSTOMDOWNLOADEXCEPTION FaultCode = "fnsCONTENTCUSTOMDOWNLOADEXCEPTION"

	FaultCodeFnsCONTENTHUBAUTHENTICATIONEXCEPTION FaultCode = "fnsCONTENTHUBAUTHENTICATIONEXCEPTION"

	FaultCodeFnsCONTENTHUBFILEDOWNLOADEXCEPTION FaultCode = "fnsCONTENTHUBFILEDOWNLOADEXCEPTION"

	FaultCodeFnsCONTENTHUBFILEHASNOURLEXCEPTION FaultCode = "fnsCONTENTHUBFILEHASNOURLEXCEPTION"

	FaultCodeFnsCONTENTHUBFILENOTFOUNDEXCEPTION FaultCode = "fnsCONTENTHUBFILENOTFOUNDEXCEPTION"

	FaultCodeFnsCONTENTHUBINVALIDOBJECTTYPEEXCEPTION FaultCode = "fnsCONTENTHUBINVALIDOBJECTTYPEEXCEPTION"

	FaultCodeFnsCONTENTHUBINVALIDPAGENUMBEREXCEPTION FaultCode = "fnsCONTENTHUBINVALIDPAGENUMBEREXCEPTION"

	FaultCodeFnsCONTENTHUBINVALIDPAYLOAD FaultCode = "fnsCONTENTHUBINVALIDPAYLOAD"

	FaultCodeFnsCONTENTHUBINVALIDRENDITIONPAGENUMBEREXCEPTION FaultCode = "fnsCONTENTHUBINVALIDRENDITIONPAGENUMBEREXCEPTION"

	FaultCodeFnsCONTENTHUBITEMTYPENOTFOUNDEXCEPTION FaultCode = "fnsCONTENTHUBITEMTYPENOTFOUNDEXCEPTION"

	FaultCodeFnsCONTENTHUBOBJECTNOTFOUNDEXCEPTION FaultCode = "fnsCONTENTHUBOBJECTNOTFOUNDEXCEPTION"

	FaultCodeFnsCONTENTHUBOPERATIONNOTSUPPORTEDEXCEPTION FaultCode = "fnsCONTENTHUBOPERATIONNOTSUPPORTEDEXCEPTION"

	FaultCodeFnsCONTENTHUBSECURITYEXCEPTION FaultCode = "fnsCONTENTHUBSECURITYEXCEPTION"

	FaultCodeFnsCONTENTHUBTIMEOUTEXCEPTION FaultCode = "fnsCONTENTHUBTIMEOUTEXCEPTION"

	FaultCodeFnsCONTENTHUBUNEXPECTEDEXCEPTION FaultCode = "fnsCONTENTHUBUNEXPECTEDEXCEPTION"

	FaultCodeFnsCONTENTIMAGESCALINGINVALIDARGUMENTSEXCEPTION FaultCode = "fnsCONTENTIMAGESCALINGINVALIDARGUMENTSEXCEPTION"

	FaultCodeFnsCONTENTIMAGESCALINGINVALIDIMAGEEXCEPTION FaultCode = "fnsCONTENTIMAGESCALINGINVALIDIMAGEEXCEPTION"

	FaultCodeFnsCONTENTIMAGESCALINGMAXRENDITIONSEXCEPTION FaultCode = "fnsCONTENTIMAGESCALINGMAXRENDITIONSEXCEPTION"

	FaultCodeFnsCONTENTIMAGESCALINGTIMEOUTEXCEPTION FaultCode = "fnsCONTENTIMAGESCALINGTIMEOUTEXCEPTION"

	FaultCodeFnsCONTENTIMAGESCALINGUNKNOWNEXCEPTION FaultCode = "fnsCONTENTIMAGESCALINGUNKNOWNEXCEPTION"

	FaultCodeFnsCUSTOMMETADATALIMITEXCEEDED FaultCode = "fnsCUSTOMMETADATALIMITEXCEEDED"

	FaultCodeFnsCUSTOMSETTINGSLIMITEXCEEDED FaultCode = "fnsCUSTOMSETTINGSLIMITEXCEEDED"

	FaultCodeFnsDATACLOUDAPICLIENTEXCEPTION FaultCode = "fnsDATACLOUDAPICLIENTEXCEPTION"

	FaultCodeFnsDATACLOUDAPIDISABLEDEXCEPTION FaultCode = "fnsDATACLOUDAPIDISABLEDEXCEPTION"

	FaultCodeFnsDATACLOUDAPIINVALIDQUERYEXCEPTION FaultCode = "fnsDATACLOUDAPIINVALIDQUERYEXCEPTION"

	FaultCodeFnsDATACLOUDAPISERVERBUSYEXCEPTION FaultCode = "fnsDATACLOUDAPISERVERBUSYEXCEPTION"

	FaultCodeFnsDATACLOUDAPISERVEREXCEPTION FaultCode = "fnsDATACLOUDAPISERVEREXCEPTION"

	FaultCodeFnsDATACLOUDAPITIMEOUTEXCEPTION FaultCode = "fnsDATACLOUDAPITIMEOUTEXCEPTION"

	FaultCodeFnsDATACLOUDAPIUNAVAILABLE FaultCode = "fnsDATACLOUDAPIUNAVAILABLE"

	FaultCodeFnsDATAINTEGRATIONINPUTERROR FaultCode = "fnsDATAINTEGRATIONINPUTERROR"

	FaultCodeFnsDATAINTEGRATIONNOTFOUND FaultCode = "fnsDATAINTEGRATIONNOTFOUND"

	FaultCodeFnsDATAINTEGRATIONNOACCESS FaultCode = "fnsDATAINTEGRATIONNOACCESS"

	FaultCodeFnsDATAINTEGRATIONPACKAGEERROR FaultCode = "fnsDATAINTEGRATIONPACKAGEERROR"

	FaultCodeFnsDATAINTEGRATIONSERVERERROR FaultCode = "fnsDATAINTEGRATIONSERVERERROR"

	FaultCodeFnsDATAINTEGRATIONVENDORSETUPERROR FaultCode = "fnsDATAINTEGRATIONVENDORSETUPERROR"

	FaultCodeFnsDUPLICATEARGUMENTVALUE FaultCode = "fnsDUPLICATEARGUMENTVALUE"

	FaultCodeFnsDUPLICATEVALUE FaultCode = "fnsDUPLICATEVALUE"

	FaultCodeFnsEMAILBATCHSIZELIMITEXCEEDED FaultCode = "fnsEMAILBATCHSIZELIMITEXCEEDED"

	FaultCodeFnsEMAILTOCASEINVALIDROUTING FaultCode = "fnsEMAILTOCASEINVALIDROUTING"

	FaultCodeFnsEMAILTOCASELIMITEXCEEDED FaultCode = "fnsEMAILTOCASELIMITEXCEEDED"

	FaultCodeFnsEMAILTOCASENOTENABLED FaultCode = "fnsEMAILTOCASENOTENABLED"

	FaultCodeFnsENTITYNOTQUERYABLE FaultCode = "fnsENTITYNOTQUERYABLE"

	FaultCodeFnsENVIRONMENTHUBMEMBERSHIPCONFLICT FaultCode = "fnsENVIRONMENTHUBMEMBERSHIPCONFLICT"

	FaultCodeFnsEXCEEDEDIDLIMIT FaultCode = "fnsEXCEEDEDIDLIMIT"

	FaultCodeFnsEXCEEDEDLEADCONVERTLIMIT FaultCode = "fnsEXCEEDEDLEADCONVERTLIMIT"

	FaultCodeFnsEXCEEDEDMAXFILTERENTITIES FaultCode = "fnsEXCEEDEDMAXFILTERENTITIES"

	FaultCodeFnsEXCEEDEDMAXSIZEREQUEST FaultCode = "fnsEXCEEDEDMAXSIZEREQUEST"

	FaultCodeFnsEXCEEDEDMAXSOBJECTS FaultCode = "fnsEXCEEDEDMAXSOBJECTS"

	FaultCodeFnsEXCEEDEDMAXTYPESLIMIT FaultCode = "fnsEXCEEDEDMAXTYPESLIMIT"

	FaultCodeFnsEXCEEDEDQUOTA FaultCode = "fnsEXCEEDEDQUOTA"

	FaultCodeFnsEXTERNALOBJECTAUTHENTICATIONEXCEPTION FaultCode = "fnsEXTERNALOBJECTAUTHENTICATIONEXCEPTION"

	FaultCodeFnsEXTERNALOBJECTCONNECTIONEXCEPTION FaultCode = "fnsEXTERNALOBJECTCONNECTIONEXCEPTION"

	FaultCodeFnsEXTERNALOBJECTEXCEPTION FaultCode = "fnsEXTERNALOBJECTEXCEPTION"

	FaultCodeFnsEXTERNALOBJECTUNSUPPORTEDEXCEPTION FaultCode = "fnsEXTERNALOBJECTUNSUPPORTEDEXCEPTION"

	FaultCodeFnsFEDERATEDSEARCHERROR FaultCode = "fnsFEDERATEDSEARCHERROR"

	FaultCodeFnsFEEDNOTENABLEDFOROBJECT FaultCode = "fnsFEEDNOTENABLEDFOROBJECT"

	FaultCodeFnsFUNCTIONALITYNOTENABLED FaultCode = "fnsFUNCTIONALITYNOTENABLED"

	FaultCodeFnsFUNCTIONALITYTEMPORARILYUNAVAILABLE FaultCode = "fnsFUNCTIONALITYTEMPORARILYUNAVAILABLE"

	FaultCodeFnsIDREQUIRED FaultCode = "fnsIDREQUIRED"

	FaultCodeFnsILLEGALQUERYPARAMETERVALUE FaultCode = "fnsILLEGALQUERYPARAMETERVALUE"

	FaultCodeFnsINACTIVEOWNERORUSER FaultCode = "fnsINACTIVEOWNERORUSER"

	FaultCodeFnsINACTIVEPORTAL FaultCode = "fnsINACTIVEPORTAL"

	FaultCodeFnsINSERTUPDATEDELETENOTALLOWEDDURINGMAINTENANCE FaultCode = "fnsINSERTUPDATEDELETENOTALLOWEDDURINGMAINTENANCE"

	FaultCodeFnsINSTALLKEYINVALID FaultCode = "fnsINSTALLKEYINVALID"

	FaultCodeFnsINSTALLKEYREQUIRED FaultCode = "fnsINSTALLKEYREQUIRED"

	FaultCodeFnsINSUFFICIENTACCESS FaultCode = "fnsINSUFFICIENTACCESS"

	FaultCodeFnsINSUFFICIENTACCESSAPEXMETADATADEPLOY FaultCode = "fnsINSUFFICIENTACCESSAPEXMETADATADEPLOY"

	FaultCodeFnsINTERNALCANVASERROR FaultCode = "fnsINTERNALCANVASERROR"

	FaultCodeFnsINVALIDASSIGNMENTRULE FaultCode = "fnsINVALIDASSIGNMENTRULE"

	FaultCodeFnsINVALIDAUTHHEADER FaultCode = "fnsINVALIDAUTHHEADER"

	FaultCodeFnsINVALIDBATCHREQUEST FaultCode = "fnsINVALIDBATCHREQUEST"

	FaultCodeFnsINVALIDBATCHSIZE FaultCode = "fnsINVALIDBATCHSIZE"

	FaultCodeFnsINVALIDCLIENT FaultCode = "fnsINVALIDCLIENT"

	FaultCodeFnsINVALIDCROSSREFERENCEKEY FaultCode = "fnsINVALIDCROSSREFERENCEKEY"

	FaultCodeFnsINVALIDDATEFORMAT FaultCode = "fnsINVALIDDATEFORMAT"

	FaultCodeFnsINVALIDFIELD FaultCode = "fnsINVALIDFIELD"

	FaultCodeFnsINVALIDFILTERLANGUAGE FaultCode = "fnsINVALIDFILTERLANGUAGE"

	FaultCodeFnsINVALIDFILTERVALUE FaultCode = "fnsINVALIDFILTERVALUE"

	FaultCodeFnsINVALIDIDFIELD FaultCode = "fnsINVALIDIDFIELD"

	FaultCodeFnsINVALIDINPUTCOMBINATION FaultCode = "fnsINVALIDINPUTCOMBINATION"

	FaultCodeFnsINVALIDLOCALELANGUAGE FaultCode = "fnsINVALIDLOCALELANGUAGE"

	FaultCodeFnsINVALIDLOCATOR FaultCode = "fnsINVALIDLOCATOR"

	FaultCodeFnsINVALIDLOGIN FaultCode = "fnsINVALIDLOGIN"

	FaultCodeFnsINVALIDMULTIPARTREQUEST FaultCode = "fnsINVALIDMULTIPARTREQUEST"

	FaultCodeFnsINVALIDNEWPASSWORD FaultCode = "fnsINVALIDNEWPASSWORD"

	FaultCodeFnsINVALIDOLDPASSWORD FaultCode = "fnsINVALIDOLDPASSWORD"

	FaultCodeFnsINVALIDOPERATION FaultCode = "fnsINVALIDOPERATION"

	FaultCodeFnsINVALIDOPERATIONWITHEXPIREDPASSWORD FaultCode = "fnsINVALIDOPERATIONWITHEXPIREDPASSWORD"

	FaultCodeFnsINVALIDPACKAGEVERSION FaultCode = "fnsINVALIDPACKAGEVERSION"

	FaultCodeFnsINVALIDPAGINGOPTION FaultCode = "fnsINVALIDPAGINGOPTION"

	FaultCodeFnsINVALIDQUERYFILTEROPERATOR FaultCode = "fnsINVALIDQUERYFILTEROPERATOR"

	FaultCodeFnsINVALIDQUERYLOCATOR FaultCode = "fnsINVALIDQUERYLOCATOR"

	FaultCodeFnsINVALIDQUERYSCOPE FaultCode = "fnsINVALIDQUERYSCOPE"

	FaultCodeFnsINVALIDREPLICATIONDATE FaultCode = "fnsINVALIDREPLICATIONDATE"

	FaultCodeFnsINVALIDSEARCH FaultCode = "fnsINVALIDSEARCH"

	FaultCodeFnsINVALIDSEARCHSCOPE FaultCode = "fnsINVALIDSEARCHSCOPE"

	FaultCodeFnsINVALIDSESSIONID FaultCode = "fnsINVALIDSESSIONID"

	FaultCodeFnsINVALIDSOAPHEADER FaultCode = "fnsINVALIDSOAPHEADER"

	FaultCodeFnsINVALIDSORTOPTION FaultCode = "fnsINVALIDSORTOPTION"

	FaultCodeFnsINVALIDSSOGATEWAYURL FaultCode = "fnsINVALIDSSOGATEWAYURL"

	FaultCodeFnsINVALIDTYPE FaultCode = "fnsINVALIDTYPE"

	FaultCodeFnsINVALIDTYPEFOROPERATION FaultCode = "fnsINVALIDTYPEFOROPERATION"

	FaultCodeFnsJIGSAWACTIONDISABLED FaultCode = "fnsJIGSAWACTIONDISABLED"

	FaultCodeFnsJIGSAWIMPORTLIMITEXCEEDED FaultCode = "fnsJIGSAWIMPORTLIMITEXCEEDED"

	FaultCodeFnsJIGSAWREQUESTNOTSUPPORTED FaultCode = "fnsJIGSAWREQUESTNOTSUPPORTED"

	FaultCodeFnsJSONPARSERERROR FaultCode = "fnsJSONPARSERERROR"

	FaultCodeFnsKEYHASBEENDESTROYED FaultCode = "fnsKEYHASBEENDESTROYED"

	FaultCodeFnsLICENSINGDATAERROR FaultCode = "fnsLICENSINGDATAERROR"

	FaultCodeFnsLICENSINGUNKNOWNERROR FaultCode = "fnsLICENSINGUNKNOWNERROR"

	FaultCodeFnsLIMITEXCEEDED FaultCode = "fnsLIMITEXCEEDED"

	FaultCodeFnsLOGINCHALLENGEISSUED FaultCode = "fnsLOGINCHALLENGEISSUED"

	FaultCodeFnsLOGINCHALLENGEPENDING FaultCode = "fnsLOGINCHALLENGEPENDING"

	FaultCodeFnsLOGINDURINGRESTRICTEDDOMAIN FaultCode = "fnsLOGINDURINGRESTRICTEDDOMAIN"

	FaultCodeFnsLOGINDURINGRESTRICTEDTIME FaultCode = "fnsLOGINDURINGRESTRICTEDTIME"

	FaultCodeFnsLOGINMUSTUSESECURITYTOKEN FaultCode = "fnsLOGINMUSTUSESECURITYTOKEN"

	FaultCodeFnsMALFORMEDID FaultCode = "fnsMALFORMEDID"

	FaultCodeFnsMALFORMEDQUERY FaultCode = "fnsMALFORMEDQUERY"

	FaultCodeFnsMALFORMEDSEARCH FaultCode = "fnsMALFORMEDSEARCH"

	FaultCodeFnsMISMATCHINGVERSIONS FaultCode = "fnsMISMATCHINGVERSIONS"

	FaultCodeFnsMISSINGARGUMENT FaultCode = "fnsMISSINGARGUMENT"

	FaultCodeFnsMISSINGRECORD FaultCode = "fnsMISSINGRECORD"

	FaultCodeFnsMODIFIED FaultCode = "fnsMODIFIED"

	FaultCodeFnsMUTUALAUTHENTICATIONFAILED FaultCode = "fnsMUTUALAUTHENTICATIONFAILED"

	FaultCodeFnsNOTACCEPTABLE FaultCode = "fnsNOTACCEPTABLE"

	FaultCodeFnsNOTMODIFIED FaultCode = "fnsNOTMODIFIED"

	FaultCodeFnsNOACTIVEDUPLICATERULE FaultCode = "fnsNOACTIVEDUPLICATERULE"

	FaultCodeFnsNORECIPIENTS FaultCode = "fnsNORECIPIENTS"

	FaultCodeFnsNOSOFTPHONELAYOUT FaultCode = "fnsNOSOFTPHONELAYOUT"

	FaultCodeFnsNUMBEROUTSIDEVALIDRANGE FaultCode = "fnsNUMBEROUTSIDEVALIDRANGE"

	FaultCodeFnsOPERATIONTOOLARGE FaultCode = "fnsOPERATIONTOOLARGE"

	FaultCodeFnsORGINMAINTENANCE FaultCode = "fnsORGINMAINTENANCE"

	FaultCodeFnsORGISDOTORG FaultCode = "fnsORGISDOTORG"

	FaultCodeFnsORGISSIGNINGUP FaultCode = "fnsORGISSIGNINGUP"

	FaultCodeFnsORGLOCKED FaultCode = "fnsORGLOCKED"

	FaultCodeFnsORGNOTOWNEDBYINSTANCE FaultCode = "fnsORGNOTOWNEDBYINSTANCE"

	FaultCodeFnsPASSWORDLOCKOUT FaultCode = "fnsPASSWORDLOCKOUT"

	FaultCodeFnsPORTALNOACCESS FaultCode = "fnsPORTALNOACCESS"

	FaultCodeFnsPOSTBODYPARSEERROR FaultCode = "fnsPOSTBODYPARSEERROR"

	FaultCodeFnsQATHETEROGENOUSCONTEXTIDS FaultCode = "fnsQATHETEROGENOUSCONTEXTIDS"

	FaultCodeFnsQATINVALIDCONTEXTID FaultCode = "fnsQATINVALIDCONTEXTID"

	FaultCodeFnsQATINVALIDQUICKACTION FaultCode = "fnsQATINVALIDQUICKACTION"

	FaultCodeFnsQUERYTIMEOUT FaultCode = "fnsQUERYTIMEOUT"

	FaultCodeFnsQUERYTOOCOMPLICATED FaultCode = "fnsQUERYTOOCOMPLICATED"

	FaultCodeFnsREALTIMEPROCESSINGTIMEEXCEEDEDLIMIT FaultCode = "fnsREALTIMEPROCESSINGTIMEEXCEEDEDLIMIT"

	FaultCodeFnsREQUESTLIMITEXCEEDED FaultCode = "fnsREQUESTLIMITEXCEEDED"

	FaultCodeFnsREQUESTRUNNINGTOOLONG FaultCode = "fnsREQUESTRUNNINGTOOLONG"

	FaultCodeFnsSERVERUNAVAILABLE FaultCode = "fnsSERVERUNAVAILABLE"

	FaultCodeFnsSERVICEDESKNOTENABLED FaultCode = "fnsSERVICEDESKNOTENABLED"

	FaultCodeFnsSOCIALCRMFEEDSERVICEAPICLIENTEXCEPTION FaultCode = "fnsSOCIALCRMFEEDSERVICEAPICLIENTEXCEPTION"

	FaultCodeFnsSOCIALCRMFEEDSERVICEAPISERVEREXCEPTION FaultCode = "fnsSOCIALCRMFEEDSERVICEAPISERVEREXCEPTION"

	FaultCodeFnsSOCIALCRMFEEDSERVICEAPIUNAVAILABLE FaultCode = "fnsSOCIALCRMFEEDSERVICEAPIUNAVAILABLE"

	FaultCodeFnsSSOSERVICEDOWN FaultCode = "fnsSSOSERVICEDOWN"

	FaultCodeFnsSSTADMINFILEDOWNLOADEXCEPTION FaultCode = "fnsSSTADMINFILEDOWNLOADEXCEPTION"

	FaultCodeFnsSTATETRANSITIONNOTALLOWED FaultCode = "fnsSTATETRANSITIONNOTALLOWED"

	FaultCodeFnsTOOMANYAPEXREQUESTS FaultCode = "fnsTOOMANYAPEXREQUESTS"

	FaultCodeFnsTOOMANYRECIPIENTS FaultCode = "fnsTOOMANYRECIPIENTS"

	FaultCodeFnsTOOMANYRECORDS FaultCode = "fnsTOOMANYRECORDS"

	FaultCodeFnsTRIALEXPIRED FaultCode = "fnsTRIALEXPIRED"

	FaultCodeFnsTXNSECURITYENDASESSION FaultCode = "fnsTXNSECURITYENDASESSION"

	FaultCodeFnsTXNSECURITYFAILCLOSE FaultCode = "fnsTXNSECURITYFAILCLOSE"

	FaultCodeFnsTXNSECURITYNOACCESS FaultCode = "fnsTXNSECURITYNOACCESS"

	FaultCodeFnsTXNSECURITYTWOFAREQUIRED FaultCode = "fnsTXNSECURITYTWOFAREQUIRED"

	FaultCodeFnsUNABLETOLOCKROW FaultCode = "fnsUNABLETOLOCKROW"

	FaultCodeFnsUNKNOWNATTACHMENTEXCEPTION FaultCode = "fnsUNKNOWNATTACHMENTEXCEPTION"

	FaultCodeFnsUNKNOWNEXCEPTION FaultCode = "fnsUNKNOWNEXCEPTION"

	FaultCodeFnsUNKNOWNORGSETTING FaultCode = "fnsUNKNOWNORGSETTING"

	FaultCodeFnsUNSUPPORTEDAPIVERSION FaultCode = "fnsUNSUPPORTEDAPIVERSION"

	FaultCodeFnsUNSUPPORTEDATTACHMENTENCODING FaultCode = "fnsUNSUPPORTEDATTACHMENTENCODING"

	FaultCodeFnsUNSUPPORTEDCLIENT FaultCode = "fnsUNSUPPORTEDCLIENT"

	FaultCodeFnsUNSUPPORTEDMEDIATYPE FaultCode = "fnsUNSUPPORTEDMEDIATYPE"

	FaultCodeFnsXMLPARSERERROR FaultCode = "fnsXMLPARSERERROR"
)

type ApiFault struct {
	XMLName xml.Name `xml:"urn:fault.enterprise.soap.sforce.com ApiFault"`

	ExceptionCode        *ExceptionCode          `xml:"exceptionCode,omitempty"`
	ExceptionMessage     string                  `xml:"exceptionMessage,omitempty"`
	ExtendedErrorDetails []*ExtendedErrorDetails `xml:"extendedErrorDetails,omitempty"`
}

type ApiQueryFault struct {
	XMLName xml.Name `xml:"urn:fault.enterprise.soap.sforce.com ApiQueryFault"`

	*ApiFault

	Row    int32 `xml:"row,omitempty"`
	Column int32 `xml:"column,omitempty"`
}

type LoginFault struct {
	XMLName xml.Name `xml:"urn:fault.enterprise.soap.sforce.com LoginFault"`

	*ApiFault
}

type InvalidQueryLocatorFault struct {
	XMLName xml.Name `xml:"urn:fault.enterprise.soap.sforce.com InvalidQueryLocatorFault"`

	*ApiFault
}

type InvalidNewPasswordFault struct {
	XMLName xml.Name `xml:"urn:fault.enterprise.soap.sforce.com InvalidNewPasswordFault"`

	*ApiFault
}

type InvalidOldPasswordFault struct {
	XMLName xml.Name `xml:"urn:fault.enterprise.soap.sforce.com InvalidOldPasswordFault"`

	*ApiFault
}

type InvalidIdFault struct {
	XMLName xml.Name `xml:"urn:fault.enterprise.soap.sforce.com InvalidIdFault"`

	*ApiFault
}

type UnexpectedErrorFault struct {
	XMLName xml.Name `xml:"urn:fault.enterprise.soap.sforce.com UnexpectedErrorFault"`

	*ApiFault
}

type InvalidFieldFault struct {
	XMLName xml.Name `xml:"urn:fault.enterprise.soap.sforce.com InvalidFieldFault"`

	*ApiQueryFault
}

type InvalidSObjectFault struct {
	XMLName xml.Name `xml:"urn:fault.enterprise.soap.sforce.com InvalidSObjectFault"`

	*ApiQueryFault
}

type MalformedQueryFault struct {
	XMLName xml.Name `xml:"urn:fault.enterprise.soap.sforce.com MalformedQueryFault"`

	*ApiQueryFault
}

type MalformedSearchFault struct {
	XMLName xml.Name `xml:"urn:fault.enterprise.soap.sforce.com MalformedSearchFault"`

	*ApiQueryFault
}

type Soap struct {
	client *SOAPClient
}

func NewSoap(url string, tls bool, auth *BasicAuth) *Soap {
	if url == "" {
		url = "https://login.salesforce.com/services/Soap/c/44.0"
	}
	client := NewSOAPClient(url, tls, auth)

	return &Soap{
		client: client,
	}
}

// Error can be either of the following types:
//
//   - LoginFault
//   - UnexpectedErrorFault
//   - InvalidIdFault
/* Login to the Salesforce.com SOAP Api */
func (service *Soap) Login(request *Login) (*LoginResponse, error) {
	response := new(LoginResponse)
	err := service.client.Call("Login", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidSObjectFault
//   - UnexpectedErrorFault
/* Describe an sObject */
func (service *Soap) DescribeSObject(request *DescribeSObject) (*DescribeSObjectResponse, error) {
	response := new(DescribeSObjectResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidSObjectFault
//   - UnexpectedErrorFault
/* Describe multiple sObjects (upto 100) */
func (service *Soap) DescribeSObjects(request *DescribeSObjects) (*DescribeSObjectsResponse, error) {
	response := new(DescribeSObjectsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - UnexpectedErrorFault
/* Describe the Global state */
func (service *Soap) DescribeGlobal(request *DescribeGlobal) (*DescribeGlobalResponse, error) {
	response := new(DescribeGlobalResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidSObjectFault
//   - UnexpectedErrorFault
/* Describe all the data category groups available for a given set of types */
func (service *Soap) DescribeDataCategoryGroups(request *DescribeDataCategoryGroups) (*DescribeDataCategoryGroupsResponse, error) {
	response := new(DescribeDataCategoryGroupsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidSObjectFault
//   - UnexpectedErrorFault
/* Describe the data category group structures for a given set of pair of types and data category group name */
func (service *Soap) DescribeDataCategoryGroupStructures(request *DescribeDataCategoryGroupStructures) (*DescribeDataCategoryGroupStructuresResponse, error) {
	response := new(DescribeDataCategoryGroupStructuresResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - UnexpectedErrorFault
/* Describe your Data Category Mappings. */
func (service *Soap) DescribeDataCategoryMappings(request *DescribeDataCategoryMappings) (*DescribeDataCategoryMappingsResponse, error) {
	response := new(DescribeDataCategoryMappingsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - UnexpectedErrorFault
/* Describes your Knowledge settings, such as if knowledgeEnabled is on or off, its default language and supported languages */
func (service *Soap) DescribeKnowledgeSettings(request *DescribeKnowledgeSettings) (*DescribeKnowledgeSettingsResponse, error) {
	response := new(DescribeKnowledgeSettingsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - UnexpectedErrorFault
/* Describe the items in an AppMenu */
func (service *Soap) DescribeAppMenu(request *DescribeAppMenu) (*DescribeAppMenuResponse, error) {
	response := new(DescribeAppMenuResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - UnexpectedErrorFault
/* Describe Gloal and Themes */
func (service *Soap) DescribeGlobalTheme(request *DescribeGlobalTheme) (*DescribeGlobalThemeResponse, error) {
	response := new(DescribeGlobalThemeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - UnexpectedErrorFault
/* Describe Themes */
func (service *Soap) DescribeTheme(request *DescribeTheme) (*DescribeThemeResponse, error) {
	response := new(DescribeThemeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidSObjectFault
//   - UnexpectedErrorFault
//   - InvalidIdFault
/* Describe the layout of the given sObject or the given actionable global page. */
func (service *Soap) DescribeLayout(request *DescribeLayout) (*DescribeLayoutResponse, error) {
	response := new(DescribeLayoutResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - UnexpectedErrorFault
/* Describe the layout of the SoftPhone */
func (service *Soap) DescribeSoftphoneLayout(request *DescribeSoftphoneLayout) (*DescribeSoftphoneLayoutResponse, error) {
	response := new(DescribeSoftphoneLayoutResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidSObjectFault
//   - UnexpectedErrorFault
/* Describe the search view of an sObject */
func (service *Soap) DescribeSearchLayouts(request *DescribeSearchLayouts) (*DescribeSearchLayoutsResponse, error) {
	response := new(DescribeSearchLayoutsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Describe a list of entity names that reflects the current user's searchable entities */
func (service *Soap) DescribeSearchableEntities(request *DescribeSearchableEntities) (*DescribeSearchableEntitiesResponse, error) {
	response := new(DescribeSearchableEntitiesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Describe a list of objects representing the order and scope of objects on a users search result page */
func (service *Soap) DescribeSearchScopeOrder(request *DescribeSearchScopeOrder) (*DescribeSearchScopeOrderResponse, error) {
	response := new(DescribeSearchScopeOrderResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Describe the compact layouts of the given sObject */
func (service *Soap) DescribeCompactLayouts(request *DescribeCompactLayouts) (*DescribeCompactLayoutsResponse, error) {
	response := new(DescribeCompactLayoutsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Describe the Path Assistants for the given sObject and optionally RecordTypes */
func (service *Soap) DescribePathAssistants(request *DescribePathAssistants) (*DescribePathAssistantsResponse, error) {
	response := new(DescribePathAssistantsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Describe the approval layouts of the given sObject */
func (service *Soap) DescribeApprovalLayout(request *DescribeApprovalLayout) (*DescribeApprovalLayoutResponse, error) {
	response := new(DescribeApprovalLayoutResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidSObjectFault
//   - UnexpectedErrorFault
/* Describe the ListViews as SOQL metadata for the generation of SOQL. */
func (service *Soap) DescribeSoqlListViews(request *DescribeSoqlListViews) (*DescribeSoqlListViewsResponse, error) {
	response := new(DescribeSoqlListViewsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Execute the specified list view and return the presentation-ready results. */
func (service *Soap) ExecuteListView(request *ExecuteListView) (*ExecuteListViewResponse, error) {
	response := new(ExecuteListViewResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidSObjectFault
//   - UnexpectedErrorFault
/* Describe the ListViews of a SObject as SOQL metadata for the generation of SOQL. */
func (service *Soap) DescribeSObjectListViews(request *DescribeSObjectListViews) (*DescribeSObjectListViewsResponse, error) {
	response := new(DescribeSObjectListViewsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - UnexpectedErrorFault
/* Describe the tabs that appear on a users page */
func (service *Soap) DescribeTabs(request *DescribeTabs) (*DescribeTabsResponse, error) {
	response := new(DescribeTabsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - UnexpectedErrorFault
/* Describe all tabs available to a user */
func (service *Soap) DescribeAllTabs(request *DescribeAllTabs) (*DescribeAllTabsResponse, error) {
	response := new(DescribeAllTabsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Describe the primary compact layouts for the sObjects requested */
func (service *Soap) DescribePrimaryCompactLayouts(request *DescribePrimaryCompactLayouts) (*DescribePrimaryCompactLayoutsResponse, error) {
	response := new(DescribePrimaryCompactLayoutsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidSObjectFault
//   - UnexpectedErrorFault
//   - InvalidIdFault
//   - InvalidFieldFault
/* Create a set of new sObjects */
func (service *Soap) Create(request *Create) (*CreateResponse, error) {
	response := new(CreateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidSObjectFault
//   - UnexpectedErrorFault
//   - InvalidIdFault
//   - InvalidFieldFault
/* Update a set of sObjects */
func (service *Soap) Update(request *Update) (*UpdateResponse, error) {
	response := new(UpdateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidSObjectFault
//   - UnexpectedErrorFault
//   - InvalidIdFault
//   - InvalidFieldFault
/* Update or insert a set of sObjects based on object id */
func (service *Soap) Upsert(request *Upsert) (*UpsertResponse, error) {
	response := new(UpsertResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidSObjectFault
//   - UnexpectedErrorFault
//   - InvalidIdFault
//   - InvalidFieldFault
/* Merge and update a set of sObjects based on object id */
func (service *Soap) Merge(request *Merge) (*MergeResponse, error) {
	response := new(MergeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - UnexpectedErrorFault
/* Delete a set of sObjects */
func (service *Soap) Delete(request *Delete) (*DeleteResponse, error) {
	response := new(DeleteResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - UnexpectedErrorFault
/* Undelete a set of sObjects */
func (service *Soap) Undelete(request *Undelete) (*UndeleteResponse, error) {
	response := new(UndeleteResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - UnexpectedErrorFault
/* Empty a set of sObjects from the recycle bin */
func (service *Soap) EmptyRecycleBin(request *EmptyRecycleBin) (*EmptyRecycleBinResponse, error) {
	response := new(EmptyRecycleBinResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidSObjectFault
//   - InvalidFieldFault
//   - MalformedQueryFault
//   - UnexpectedErrorFault
//   - InvalidIdFault
/* Get a set of sObjects */
func (service *Soap) Retrieve(request *Retrieve) (*RetrieveResponse, error) {
	response := new(RetrieveResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - UnexpectedErrorFault
//   - InvalidIdFault
/* Submit an entity to a workflow process or process a workitem */
func (service *Soap) Process(request *Process) (*ProcessResponse, error) {
	response := new(ProcessResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - UnexpectedErrorFault
/* convert a set of leads */
func (service *Soap) ConvertLead(request *ConvertLead) (*ConvertLeadResponse, error) {
	response := new(ConvertLeadResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - UnexpectedErrorFault
/* Logout the current user, invalidating the current session. */
func (service *Soap) Logout(request *Logout) (*LogoutResponse, error) {
	response := new(LogoutResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - UnexpectedErrorFault
/* Logs out and invalidates session ids */
func (service *Soap) InvalidateSessions(request *InvalidateSessions) (*InvalidateSessionsResponse, error) {
	response := new(InvalidateSessionsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidSObjectFault
//   - UnexpectedErrorFault
/* Get the IDs for deleted sObjects */
func (service *Soap) GetDeleted(request *GetDeleted) (*GetDeletedResponse, error) {
	response := new(GetDeletedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidSObjectFault
//   - UnexpectedErrorFault
/* Get the IDs for updated sObjects */
func (service *Soap) GetUpdated(request *GetUpdated) (*GetUpdatedResponse, error) {
	response := new(GetUpdatedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidSObjectFault
//   - InvalidFieldFault
//   - MalformedQueryFault
//   - InvalidIdFault
//   - UnexpectedErrorFault
//   - InvalidQueryLocatorFault
/* Create a Query Cursor */
func (service *Soap) Query(request *Query) (*QueryResponse, error) {
	response := new(QueryResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidSObjectFault
//   - InvalidFieldFault
//   - MalformedQueryFault
//   - InvalidIdFault
//   - UnexpectedErrorFault
//   - InvalidQueryLocatorFault
/* Create a Query Cursor, including deleted sObjects */
func (service *Soap) QueryAll(request *QueryAll) (*QueryAllResponse, error) {
	response := new(QueryAllResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidQueryLocatorFault
//   - UnexpectedErrorFault
//   - InvalidFieldFault
//   - MalformedQueryFault
/* Gets the next batch of sObjects from a query */
func (service *Soap) QueryMore(request *QueryMore) (*QueryMoreResponse, error) {
	response := new(QueryMoreResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidSObjectFault
//   - InvalidFieldFault
//   - MalformedSearchFault
//   - UnexpectedErrorFault
/* Search for sObjects */
func (service *Soap) Search(request *Search) (*SearchResponse, error) {
	response := new(SearchResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - UnexpectedErrorFault
/* Gets server timestamp */
func (service *Soap) GetServerTimestamp(request *GetServerTimestamp) (*GetServerTimestampResponse, error) {
	response := new(GetServerTimestampResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidIdFault
//   - InvalidNewPasswordFault
//   - UnexpectedErrorFault
/* Set a user's password */
func (service *Soap) SetPassword(request *SetPassword) (*SetPasswordResponse, error) {
	response := new(SetPasswordResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidNewPasswordFault
//   - InvalidOldPasswordFault
//   - UnexpectedErrorFault
/* Change the current user's password */
func (service *Soap) ChangeOwnPassword(request *ChangeOwnPassword) (*ChangeOwnPasswordResponse, error) {
	response := new(ChangeOwnPasswordResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidIdFault
//   - UnexpectedErrorFault
/* Reset a user's password */
func (service *Soap) ResetPassword(request *ResetPassword) (*ResetPasswordResponse, error) {
	response := new(ResetPasswordResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - UnexpectedErrorFault
/* Returns standard information relevant to the current user */
func (service *Soap) GetUserInfo(request *GetUserInfo) (*GetUserInfoResponse, error) {
	response := new(GetUserInfoResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - UnexpectedErrorFault
/* Delete a set of sObjects by example. The passed SOBject is a template for the object to delete */
func (service *Soap) DeleteByExample(request *DeleteByExample) (*DeleteByExampleResponse, error) {
	response := new(DeleteByExampleResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - UnexpectedErrorFault
/* Send existing draft EmailMessage */
func (service *Soap) SendEmailMessage(request *SendEmailMessage) (*SendEmailMessageResponse, error) {
	response := new(SendEmailMessageResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// Error can be either of the following types:
//
//   - UnexpectedErrorFault
/* Send outbound email */
func (service *Soap) SendEmail(request *SendEmail) (*SendEmailResponse, error) {
	response := new(SendEmailResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - UnexpectedErrorFault
/* Perform a template merge on one or more blocks of text. */
func (service *Soap) RenderEmailTemplate(request *RenderEmailTemplate) (*RenderEmailTemplateResponse, error) {
	response := new(RenderEmailTemplateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - UnexpectedErrorFault
/* Perform a template merge using an email template stored in the database. */
func (service *Soap) RenderStoredEmailTemplate(request *RenderStoredEmailTemplate) (*RenderStoredEmailTemplateResponse, error) {
	response := new(RenderStoredEmailTemplateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Perform a series of predefined actions such as quick create or log a task */
func (service *Soap) PerformQuickActions(request *PerformQuickActions) (*PerformQuickActionsResponse, error) {
	response := new(PerformQuickActionsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Describe the details of a series of quick actions */
func (service *Soap) DescribeQuickActions(request *DescribeQuickActions) (*DescribeQuickActionsResponse, error) {
	response := new(DescribeQuickActionsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Describe the details of a series of quick actions in context of requested recordType id for Update actions */
func (service *Soap) DescribeQuickActionsForRecordType(request *DescribeQuickActionsForRecordType) (*DescribeQuickActionsForRecordTypeResponse, error) {
	response := new(DescribeQuickActionsForRecordTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Describe the details of a series of quick actions available for the given contextType */
func (service *Soap) DescribeAvailableQuickActions(request *DescribeAvailableQuickActions) (*DescribeAvailableQuickActionsResponse, error) {
	response := new(DescribeAvailableQuickActionsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Retrieve the template sobjects, if appropriate, for the given quick action names in a given context */
func (service *Soap) RetrieveQuickActionTemplates(request *RetrieveQuickActionTemplates) (*RetrieveQuickActionTemplatesResponse, error) {
	response := new(RetrieveQuickActionTemplatesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Retrieve the template sobjects, if appropriate, for the given quick action names in a given contexts when used a mass quick action */
func (service *Soap) RetrieveMassQuickActionTemplates(request *RetrieveMassQuickActionTemplates) (*RetrieveMassQuickActionTemplatesResponse, error) {
	response := new(RetrieveMassQuickActionTemplatesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Describe visualforce for an org */
func (service *Soap) DescribeVisualForce(request *DescribeVisualForce) (*DescribeVisualForceResponse, error) {
	response := new(DescribeVisualForceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidSObjectFault
//   - UnexpectedErrorFault
//   - InvalidFieldFault
/* Find duplicates for a set of sObjects */
func (service *Soap) FindDuplicates(request *FindDuplicates) (*FindDuplicatesResponse, error) {
	response := new(FindDuplicatesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidIdFault
//   - UnexpectedErrorFault
/* Find duplicates for a set of ids */
func (service *Soap) FindDuplicatesByIds(request *FindDuplicatesByIds) (*FindDuplicatesByIdsResponse, error) {
	response := new(FindDuplicatesByIdsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Return the renameable nouns from the server for use in presentation using the salesforce grammar engine */
func (service *Soap) DescribeNouns(request *DescribeNouns) (*DescribeNounsResponse, error) {
	response := new(DescribeNounsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

var timeout = time.Duration(30 * time.Second)

func dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, timeout)
}

type SOAPEnvelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`

	Body SOAPBody
}

type SOAPHeader struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header"`

	Header interface{}
}

type SOAPBody struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`

	Fault   *SOAPFault  `xml:",omitempty"`
	Content interface{} `xml:",omitempty"`
}

type SOAPFault struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault"`

	Code   string `xml:"faultcode,omitempty"`
	String string `xml:"faultstring,omitempty"`
	Actor  string `xml:"faultactor,omitempty"`
	Detail string `xml:"detail,omitempty"`
}

type BasicAuth struct {
	Login    string
	Password string
}

type SOAPClient struct {
	url  string
	tls  bool
	auth *BasicAuth
}

func (b *SOAPBody) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if b.Content == nil {
		return xml.UnmarshalError("Content must be a pointer to a struct")
	}

	var (
		token    xml.Token
		err      error
		consumed bool
	)

Loop:
	for {
		if token, err = d.Token(); err != nil {
			return err
		}

		if token == nil {
			break
		}

		switch se := token.(type) {
		case xml.StartElement:
			if consumed {
				return xml.UnmarshalError("Found multiple elements inside SOAP body; not wrapped-document/literal WS-I compliant")
			} else if se.Name.Space == "http://schemas.xmlsoap.org/soap/envelope/" && se.Name.Local == "Fault" {
				b.Fault = &SOAPFault{}
				b.Content = nil

				err = d.DecodeElement(b.Fault, &se)
				if err != nil {
					return err
				}

				consumed = true
			} else {
				if err = d.DecodeElement(b.Content, &se); err != nil {
					return err
				}

				consumed = true
			}
		case xml.EndElement:
			break Loop
		}
	}

	return nil
}

func (f *SOAPFault) Error() string {
	return f.String
}

func NewSOAPClient(url string, tls bool, auth *BasicAuth) *SOAPClient {
	return &SOAPClient{
		url:  url,
		tls:  tls,
		auth: auth,
	}
}

func (s *SOAPClient) Call(soapAction string, request, response interface{}) error {
	envelope := SOAPEnvelope{
	//Header:        SoapHeader{},
	}

	envelope.Body.Content = request
	buffer := new(bytes.Buffer)

	encoder := xml.NewEncoder(buffer)
	//encoder.Indent("  ", "    ")

	if err := encoder.Encode(envelope); err != nil {
		return err
	}

	if err := encoder.Flush(); err != nil {
		return err
	}

	log.Println(buffer.String())

	req, err := http.NewRequest("POST", s.url, buffer)
	if err != nil {
		return err
	}
	if s.auth != nil {
		req.SetBasicAuth(s.auth.Login, s.auth.Password)
	}

	req.Header.Add("Content-Type", "text/xml; charset=\"utf-8\"")
	if soapAction != "" {
		req.Header.Add("SOAPAction", soapAction)
	}

	req.Header.Set("User-Agent", "gowsdl/0.1")
	req.Close = true

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: s.tls,
		},
		Dial: dialTimeout,
	}

	client := &http.Client{Transport: tr}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	rawbody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if len(rawbody) == 0 {
		log.Println("empty response")
		return nil
	}

	log.Println(string(rawbody))
	respEnvelope := new(SOAPEnvelope)
	respEnvelope.Body = SOAPBody{Content: response}
	err = xml.Unmarshal(rawbody, respEnvelope)
	if err != nil {
		return err
	}

	fault := respEnvelope.Body.Fault
	if fault != nil {
		return fault
	}

	return nil
}
