package md2

import (
	"encoding/json"
	"errors"
)

type MediumPost struct {
	B       string  `json:"b"`
	Payload Payload `json:"payload"`
	Success bool    `json:"success"`
	V       int64   `json:"v"`
}

type Payload struct {
	References              References              `json:"references"`
	Mode                    string                  `json:"mode"`
	ShareKey                string                  `json:"shareKey"`
	Collaborators           []any                   `json:"collaborators"`
	CollectionUserRelations CollectionUserRelations `json:"collectionUserRelations"`
	MentionedUsers          MentionedUsers          `json:"mentionedUsers"`
	Value                   Value                   `json:"value"`
	HideMeter               bool                    `json:"hideMeter"`
}

type CollectionUserRelation struct {
	CollectionID string `json:"collectionId"`
	Role         string `json:"role"`
	UserID       string `json:"userId"`
}

type CollectionUserRelations []CollectionUserRelation

type MentionedUser struct {
	BackgroundImageID                    string  `json:"backgroundImageId"`
	Bio                                  string  `json:"bio"`
	FacebookDisplayName                  string  `json:"facebookDisplayName"`
	ImageID                              string  `json:"imageId"`
	LanguageCode                         string  `json:"languageCode"`
	Name                                 string  `json:"name"`
	TwitterScreenName                    string  `json:"twitterScreenName"`
	Type                                 string  `json:"type"`
	UserID                               string  `json:"userId"`
	Username                             string  `json:"username"`
	UserDismissableFlags                 []int64 `json:"userDismissableFlags"`
	AllowNotes                           int64   `json:"allowNotes"`
	CreatedAt                            int64   `json:"createdAt"`
	MediumMemberAt                       int64   `json:"mediumMemberAt"`
	PostSubscribeMembershipUpsellShownAt int64   `json:"postSubscribeMembershipUpsellShownAt"`
	SubdomainCreatedAt                   int64   `json:"subdomainCreatedAt"`
	HasCompletedProfile                  bool    `json:"hasCompletedProfile"`
	HasSeenIcelandOnboarding             bool    `json:"hasSeenIcelandOnboarding"`
	IsCreatorPartnerProgramEnrolled      bool    `json:"isCreatorPartnerProgramEnrolled"`
	IsMembershipTrialEligible            bool    `json:"isMembershipTrialEligible"`
	IsSuspended                          bool    `json:"isSuspended"`
	IsWriterProgramEnrolled              bool    `json:"isWriterProgramEnrolled"`
	OptInToIceland                       bool    `json:"optInToIceland"`
}

type MentionedUsers []MentionedUser

type References struct {
	Social      any `json:"social"`
	SocialStats any `json:"socialStats"`
	User        any `json:"user"`
}

type Value struct {
	Content                           Content  `json:"content"`
	ApprovedHomeCollectionID          string   `json:"approvedHomeCollectionId"`
	CanonicalURL                      string   `json:"canonicalUrl"`
	CreatorID                         string   `json:"creatorId"`
	DetectedLanguage                  string   `json:"detectedLanguage"`
	DisplayAuthor                     string   `json:"displayAuthor"`
	EditorialPreviewDek               string   `json:"editorialPreviewDek"`
	EditorialPreviewTitle             string   `json:"editorialPreviewTitle"`
	ExperimentalCSS                   string   `json:"experimentalCss"`
	HomeCollectionID                  string   `json:"homeCollectionId"`
	ID                                string   `json:"id"`
	ImportedURL                       string   `json:"importedUrl"`
	InResponseToMediaResourceID       string   `json:"inResponseToMediaResourceId"`
	InResponseToPostID                string   `json:"inResponseToPostId"`
	LatestPublishedVersion            string   `json:"latestPublishedVersion"`
	LatestVersion                     string   `json:"latestVersion"`
	MediumURL                         string   `json:"mediumUrl"`
	MigrationID                       string   `json:"migrationId"`
	NewsletterID                      string   `json:"newsletterId"`
	ProxyPostFaviconURL               string   `json:"proxyPostFaviconUrl"`
	ProxyPostProviderName             string   `json:"proxyPostProviderName"`
	SeoTitle                          string   `json:"seoTitle"`
	SequenceID                        string   `json:"sequenceId"`
	Slug                              string   `json:"slug"`
	SocialDek                         string   `json:"socialDek"`
	SocialTitle                       string   `json:"socialTitle"`
	Title                             string   `json:"title"`
	TranslationSourceCreatorID        string   `json:"translationSourceCreatorId"`
	TranslationSourcePostID           string   `json:"translationSourcePostId"`
	Type                              string   `json:"type"`
	UniqueSlug                        string   `json:"uniqueSlug"`
	VersionID                         string   `json:"versionId"`
	WebCanonicalURL                   string   `json:"webCanonicalUrl"`
	Virtuals                          Virtuals `json:"virtuals"`
	AcceptedAt                        int64    `json:"acceptedAt"`
	AudioVersionDurationSec           int64    `json:"audioVersionDurationSec"`
	CardType                          int64    `json:"cardType"`
	CreatedAt                         int64    `json:"createdAt"`
	CurationEligibleAt                int64    `json:"curationEligibleAt"`
	DeletedAt                         int64    `json:"deletedAt"`
	FeatureLockRequestAcceptedAt      int64    `json:"featureLockRequestAcceptedAt"`
	FirstPublishedAt                  int64    `json:"firstPublishedAt"`
	HightowerMinimumGuaranteeEndsAt   int64    `json:"hightowerMinimumGuaranteeEndsAt"`
	HightowerMinimumGuaranteeStartsAt int64    `json:"hightowerMinimumGuaranteeStartsAt"`
	ImportedPublishedAt               int64    `json:"importedPublishedAt"`
	InResponseToRemovedAt             int64    `json:"inResponseToRemovedAt"`
	LatestPublishedAt                 int64    `json:"latestPublishedAt"`
	LatestRev                         int64    `json:"latestRev"`
	LayerCake                         int64    `json:"layerCake"`
	License                           int64    `json:"license"`
	LockedPostSource                  int64    `json:"lockedPostSource"`
	MongerRequestType                 int64    `json:"mongerRequestType"`
	ProxyPostType                     int64    `json:"proxyPostType"`
	ResponseDistribution              int64    `json:"responseDistribution"`
	ResponseHiddenOnParentPostAt      int64    `json:"responseHiddenOnParentPostAt"`
	SeriesLastAppendedAt              int64    `json:"seriesLastAppendedAt"`
	ShortformType                     int64    `json:"shortformType"`
	UpdatedAt                         int64    `json:"updatedAt"`
	Visibility                        int64    `json:"visibility"`
	AllowResponses                    bool     `json:"allowResponses"`
	Coverless                         bool     `json:"coverless"`
	HasUnpublishedEdits               bool     `json:"hasUnpublishedEdits"`
	IsApprovedTranslation             bool     `json:"isApprovedTranslation"`
	IsBlockedFromHightower            bool     `json:"isBlockedFromHightower"`
	IsDistributionAlertDismissed      bool     `json:"isDistributionAlertDismissed"`
	IsEligibleForRevenue              bool     `json:"isEligibleForRevenue"`
	IsLimitedState                    bool     `json:"isLimitedState"`
	IsLockedResponse                  bool     `json:"isLockedResponse"`
	IsMarkedPaywallOnly               bool     `json:"isMarkedPaywallOnly"`
	IsNewsletter                      bool     `json:"isNewsletter"`
	IsProxyPost                       bool     `json:"isProxyPost"`
	IsPublishToEmail                  bool     `json:"isPublishToEmail"`
	IsSeries                          bool     `json:"isSeries"`
	IsShortform                       bool     `json:"isShortform"`
	IsSubscriptionLocked              bool     `json:"isSubscriptionLocked"`
	IsSuspended                       bool     `json:"isSuspended"`
	IsTitleSynthesized                bool     `json:"isTitleSynthesized"`
	NotifyFacebook                    bool     `json:"notifyFacebook"`
	NotifyFollowers                   bool     `json:"notifyFollowers"`
	NotifyTwitter                     bool     `json:"notifyTwitter"`
	ResponsesLocked                   bool     `json:"responsesLocked"`
	Vote                              bool     `json:"vote"`
}

type Content struct {
	Subtitle    string      `json:"subtitle"`
	BodyModel   BodyModel   `json:"bodyModel"`
	PostDisplay PostDisplay `json:"postDisplay"`
}

type BodyModel struct {
	Paragraphs Paragraphs `json:"paragraphs"`
	Sections   Sections   `json:"sections"`
}

type ParagraphType int

const (
	Basic         ParagraphType = 1
	BigT          ParagraphType = 3
	Image         ParagraphType = 4
	Quote         ParagraphType = 6
	CodeBlock     ParagraphType = 8
	UnOrderedList ParagraphType = 9
	OrderedList   ParagraphType = 10
	Embed         ParagraphType = 11
	SmallT        ParagraphType = 13
	EmbeddedLink  ParagraphType = 14
)

type Paragraph struct {
	MixtapeMetadata   MixtapeMetadata   `json:"mixtapeMetadata"`
	Name              string            `json:"name"`
	Text              string            `json:"text"`
	CodeBlockMetadata CodeBlockMetadata `json:"codeBlockMetadata"`
	Markups           Markups           `json:"markups"`
	Iframe            Iframe            `json:"iframe"`
	Metadata          Metadata          `json:"metadata"`
	Type              ParagraphType     `json:"type"`
	Layout            int               `json:"layout"`
	HasDropCap        bool              `json:"hasDropCap,omitempty"`
}

type Paragraphs []Paragraph

type MarkupType int

const (
	Bold          MarkupType = 1
	Italic        MarkupType = 2
	LinkOrMention MarkupType = 3
	Highlight     MarkupType = 10
)

type Markup struct {
	Href       string     `json:"href"`
	Title      string     `json:"title"`
	Rel        string     `json:"rel"`
	UserId     string     `json:"userId"`
	Type       MarkupType `json:"type"`
	Start      int        `json:"start"`
	End        int        `json:"end"`
	AnchorType int        `json:"anchorType"`
}

type Markups []Markup

type Metadata struct {
	Id              string `json:"id"`
	Alt             string `json:"alt"`
	UnsplashPhotoId string `json:"unsplashPhotoId"`
	OriginalWidth   int    `json:"originalWidth"`
	OriginalHeight  int    `json:"originalHeight"`
	IsFeatured      bool   `json:"isFeatured"`
}

type Iframe struct {
	MediaResourceId string `json:"mediaResourceId"`
	ThumbnailUrl    string `json:"thumbnailUrl"`
	IframeWidth     int    `json:"iframeWidth"`
	IframeHeight    int    `json:"iframeHeight"`
}

type CodeBlockMetadata struct {
	Lang string `json:"lang"`
	Mode int    `json:"mode"`
}

type Sections []Section

type Section struct {
	Name       string `json:"name"`
	StartIndex int    `json:"startIndex"`
}

type PostDisplay struct {
	Coverless bool `json:"coverless"`
}

type Virtuals struct {
	Links                   Links            `json:"links"`
	MetaDescription         string           `json:"metaDescription"`
	Subtitle                string           `json:"subtitle"`
	PreviewImage            PreviewImage     `json:"previewImage"`
	Tags                    []interface{}    `json:"tags"`
	Topics                  Topics           `json:"topics"`
	UsersBySocialRecommends []interface{}    `json:"usersBySocialRecommends"`
	UserPostRelation        UserPostRelation `json:"userPostRelation"`
	ImageCount              int64            `json:"imageCount"`
	ReadingList             int64            `json:"readingList"`
	ReadingTime             float64          `json:"readingTime"`
	Recommends              int64            `json:"recommends"`
	ResponsesCreatedCount   int64            `json:"responsesCreatedCount"`
	SectionCount            int64            `json:"sectionCount"`
	SocialRecommendsCount   int64            `json:"socialRecommendsCount"`
	TotalClapCount          int64            `json:"totalClapCount"`
	WordCount               int64            `json:"wordCount"`
	AllowNotes              bool             `json:"allowNotes"`
	IsBookmarked            bool             `json:"isBookmarked"`
	IsLockedPreviewOnly     bool             `json:"isLockedPreviewOnly"`
	NoIndex                 bool             `json:"noIndex"`
}

type Links struct {
	Version     string  `json:"version"`
	Entries     Entries `json:"entries"`
	GeneratedAt int64   `json:"generatedAt"`
}

type Entry struct {
	URL        string        `json:"url"`
	Alts       []interface{} `json:"alts"`
	HTTPStatus int64         `json:"httpStatus"`
}

type Entries []Entry

type PreviewImage struct {
	BackgroundSize string `json:"backgroundSize"`
	Filter         string `json:"filter"`
	ImageID        string `json:"imageId"`
	Strategy       string `json:"strategy"`
	Height         int64  `json:"height"`
	OriginalHeight int64  `json:"originalHeight"`
	OriginalWidth  int64  `json:"originalWidth"`
	Width          int64  `json:"width"`
}

type Topic struct {
	Description     string        `json:"description"`
	Name            string        `json:"name"`
	SeoTitle        string        `json:"seoTitle"`
	Slug            string        `json:"slug"`
	TopicID         string        `json:"topicId"`
	Type            string        `json:"type"`
	RelatedTags     []interface{} `json:"relatedTags"`
	RelatedTopicIds []interface{} `json:"relatedTopicIds"`
	RelatedTopics   []interface{} `json:"relatedTopics"`
	Image           TopicImage    `json:"image"`
	CreatedAt       int64         `json:"createdAt"`
	DeletedAt       int64         `json:"deletedAt"`
	Visibility      int64         `json:"visibility"`
}

type TopicImage struct {
	ID             string `json:"id"`
	OriginalHeight int64  `json:"originalHeight"`
	OriginalWidth  int64  `json:"originalWidth"`
}

type Topics []Topic

type UserPostRelation struct {
	LastReadParagraphName              string `json:"lastReadParagraphName"`
	LastReadSectionName                string `json:"lastReadSectionName"`
	LastReadVersionID                  string `json:"lastReadVersionId"`
	PostID                             string `json:"postId"`
	UserID                             string `json:"userId"`
	AudioProgressSec                   int64  `json:"audioProgressSec"`
	ClapCount                          int64  `json:"clapCount"`
	CollaboratorAddedAt                int64  `json:"collaboratorAddedAt"`
	LastReadAt                         int64  `json:"lastReadAt"`
	LastReadPercentage                 int64  `json:"lastReadPercentage"`
	NotesAddedAt                       int64  `json:"notesAddedAt"`
	PresentedCountInResponseManagement int64  `json:"presentedCountInResponseManagement"`
	PresentedCountInStream             int64  `json:"presentedCountInStream"`
	QueuedAt                           int64  `json:"queuedAt"`
	ReadAt                             int64  `json:"readAt"`
	ReadLaterAddedAt                   int64  `json:"readLaterAddedAt"`
	SeriesFirstViewedAt                int64  `json:"seriesFirstViewedAt"`
	SeriesLastViewedAt                 int64  `json:"seriesLastViewedAt"`
	SeriesUpdateNotifsOptedInAt        int64  `json:"seriesUpdateNotifsOptedInAt"`
	SubscribedAt                       int64  `json:"subscribedAt"`
	ViewedAt                           int64  `json:"viewedAt"`
	VotedAt                            int64  `json:"votedAt"`
}

type MixtapeMetadata struct {
	MediaResourceId  string `json:"mediaResourceId"`
	ThumbnailImageId string `json:"thumbnailImageId"`
	Href             string `json:"href"`
}

func toMediumPost(mp *MediumPost, jsonByte []byte) error {
	if err := json.Unmarshal(jsonByte, mp); err != nil {
		return errors.New("invalid media post response")
	}
	return nil
}
