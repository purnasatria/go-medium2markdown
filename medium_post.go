package main

type MediumPost struct {
	B       string  `json:"b"`
	Payload Payload `json:"payload"`
	Success bool    `json:"success"`
	V       int64   `json:"v"`
}

type Payload struct {
	Collaborators           []any                   `json:"collaborators"`
	CollectionUserRelations CollectionUserRelations `json:"collectionUserRelations"`
	HideMeter               bool                    `json:"hideMeter"`
	MentionedUsers          MentionedUsers          `json:"mentionedUsers"`
	Mode                    string                  `json:"mode"`
	References              References              `json:"references"`
	ShareKey                string                  `json:"shareKey"`
	Value                   Value                   `json:"value"`
}

type CollectionUserRelation struct {
	CollectionID string `json:"collectionId"`
	Role         string `json:"role"`
	UserID       string `json:"userId"`
}

type CollectionUserRelations []CollectionUserRelation

type MentionedUser struct {
	AllowNotes                           int64   `json:"allowNotes"`
	BackgroundImageID                    string  `json:"backgroundImageId"`
	Bio                                  string  `json:"bio"`
	CreatedAt                            int64   `json:"createdAt"`
	FacebookDisplayName                  string  `json:"facebookDisplayName"`
	HasCompletedProfile                  bool    `json:"hasCompletedProfile"`
	HasSeenIcelandOnboarding             bool    `json:"hasSeenIcelandOnboarding"`
	ImageID                              string  `json:"imageId"`
	IsCreatorPartnerProgramEnrolled      bool    `json:"isCreatorPartnerProgramEnrolled"`
	IsMembershipTrialEligible            bool    `json:"isMembershipTrialEligible"`
	IsSuspended                          bool    `json:"isSuspended"`
	IsWriterProgramEnrolled              bool    `json:"isWriterProgramEnrolled"`
	LanguageCode                         string  `json:"languageCode"`
	MediumMemberAt                       int64   `json:"mediumMemberAt"`
	Name                                 string  `json:"name"`
	OptInToIceland                       bool    `json:"optInToIceland"`
	PostSubscribeMembershipUpsellShownAt int64   `json:"postSubscribeMembershipUpsellShownAt"`
	SubdomainCreatedAt                   int64   `json:"subdomainCreatedAt"`
	TwitterScreenName                    string  `json:"twitterScreenName"`
	Type                                 string  `json:"type"`
	UserDismissableFlags                 []int64 `json:"userDismissableFlags"`
	UserID                               string  `json:"userId"`
	Username                             string  `json:"username"`
}

type MentionedUsers []MentionedUser

type References struct {
	Social      any `json:"social"`
	SocialStats any `json:"socialStats"`
	User        any `json:"user"`
}

type Value struct {
	AcceptedAt                        int64    `json:"acceptedAt"`
	AllowResponses                    bool     `json:"allowResponses"`
	ApprovedHomeCollectionID          string   `json:"approvedHomeCollectionId"`
	AudioVersionDurationSec           int64    `json:"audioVersionDurationSec"`
	CanonicalURL                      string   `json:"canonicalUrl"`
	CardType                          int64    `json:"cardType"`
	Content                           Content  `json:"content"`
	Coverless                         bool     `json:"coverless"`
	CreatedAt                         int64    `json:"createdAt"`
	CreatorID                         string   `json:"creatorId"`
	CurationEligibleAt                int64    `json:"curationEligibleAt"`
	DeletedAt                         int64    `json:"deletedAt"`
	DetectedLanguage                  string   `json:"detectedLanguage"`
	DisplayAuthor                     string   `json:"displayAuthor"`
	EditorialPreviewDek               string   `json:"editorialPreviewDek"`
	EditorialPreviewTitle             string   `json:"editorialPreviewTitle"`
	ExperimentalCSS                   string   `json:"experimentalCss"`
	FeatureLockRequestAcceptedAt      int64    `json:"featureLockRequestAcceptedAt"`
	FirstPublishedAt                  int64    `json:"firstPublishedAt"`
	HasUnpublishedEdits               bool     `json:"hasUnpublishedEdits"`
	HightowerMinimumGuaranteeEndsAt   int64    `json:"hightowerMinimumGuaranteeEndsAt"`
	HightowerMinimumGuaranteeStartsAt int64    `json:"hightowerMinimumGuaranteeStartsAt"`
	HomeCollectionID                  string   `json:"homeCollectionId"`
	ID                                string   `json:"id"`
	ImportedPublishedAt               int64    `json:"importedPublishedAt"`
	ImportedURL                       string   `json:"importedUrl"`
	InResponseToMediaResourceID       string   `json:"inResponseToMediaResourceId"`
	InResponseToPostID                string   `json:"inResponseToPostId"`
	InResponseToRemovedAt             int64    `json:"inResponseToRemovedAt"`
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
	LatestPublishedAt                 int64    `json:"latestPublishedAt"`
	LatestPublishedVersion            string   `json:"latestPublishedVersion"`
	LatestRev                         int64    `json:"latestRev"`
	LatestVersion                     string   `json:"latestVersion"`
	LayerCake                         int64    `json:"layerCake"`
	License                           int64    `json:"license"`
	LockedPostSource                  int64    `json:"lockedPostSource"`
	MediumURL                         string   `json:"mediumUrl"`
	MigrationID                       string   `json:"migrationId"`
	MongerRequestType                 int64    `json:"mongerRequestType"`
	NewsletterID                      string   `json:"newsletterId"`
	NotifyFacebook                    bool     `json:"notifyFacebook"`
	NotifyFollowers                   bool     `json:"notifyFollowers"`
	NotifyTwitter                     bool     `json:"notifyTwitter"`
	ProxyPostFaviconURL               string   `json:"proxyPostFaviconUrl"`
	ProxyPostProviderName             string   `json:"proxyPostProviderName"`
	ProxyPostType                     int64    `json:"proxyPostType"`
	ResponseDistribution              int64    `json:"responseDistribution"`
	ResponseHiddenOnParentPostAt      int64    `json:"responseHiddenOnParentPostAt"`
	ResponsesLocked                   bool     `json:"responsesLocked"`
	SeoTitle                          string   `json:"seoTitle"`
	SequenceID                        string   `json:"sequenceId"`
	SeriesLastAppendedAt              int64    `json:"seriesLastAppendedAt"`
	ShortformType                     int64    `json:"shortformType"`
	Slug                              string   `json:"slug"`
	SocialDek                         string   `json:"socialDek"`
	SocialTitle                       string   `json:"socialTitle"`
	Title                             string   `json:"title"`
	TranslationSourceCreatorID        string   `json:"translationSourceCreatorId"`
	TranslationSourcePostID           string   `json:"translationSourcePostId"`
	Type                              string   `json:"type"`
	UniqueSlug                        string   `json:"uniqueSlug"`
	UpdatedAt                         int64    `json:"updatedAt"`
	VersionID                         string   `json:"versionId"`
	Virtuals                          Virtuals `json:"virtuals"`
	Visibility                        int64    `json:"visibility"`
	Vote                              bool     `json:"vote"`
	WebCanonicalURL                   string   `json:"webCanonicalUrl"`
}

type Content struct {
	BodyModel   BodyModel   `json:"bodyModel"`
	PostDisplay PostDisplay `json:"postDisplay"`
	Subtitle    string      `json:"subtitle"`
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
	Name              string            `json:"name"`
	Type              ParagraphType     `json:"type"`
	Text              string            `json:"text"`
	Markups           Markups           `json:"markups"`
	HasDropCap        bool              `json:"hasDropCap,omitempty"`
	Layout            int               `json:"layout"`
	Metadata          Metadata          `json:"metadata"`
	Iframe            Iframe            `json:"iframe"`
	CodeBlockMetadata CodeBlockMetadata `json:"codeBlockMetadata"`
	MixtapeMetadata   MixtapeMetadata   `json:"mixtapeMetadata"`
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
	Type       MarkupType `json:"type"`
	Start      int        `json:"start"`
	End        int        `json:"end"`
	Href       string     `json:"href"`
	Title      string     `json:"title"`
	Rel        string     `json:"rel"`
	AnchorType int        `json:"anchorType"`
	UserId     string     `json:"userId"`
}

type Markups []Markup

type Metadata struct {
	Id              string `json:"id"`
	OriginalWidth   int    `json:"originalWidth"`
	OriginalHeight  int    `json:"originalHeight"`
	IsFeatured      bool   `json:"isFeatured"`
	Alt             string `json:"alt"`
	UnsplashPhotoId string `json:"unsplashPhotoId"`
}

type Iframe struct {
	MediaResourceId string `json:"mediaResourceId"`
	IframeWidth     int    `json:"iframeWidth"`
	IframeHeight    int    `json:"iframeHeight"`
	ThumbnailUrl    string `json:"thumbnailUrl"`
}

type CodeBlockMetadata struct {
	Mode int    `json:"mode"`
	Lang string `json:"lang"`
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
	AllowNotes              bool             `json:"allowNotes"`
	ImageCount              int64            `json:"imageCount"`
	IsBookmarked            bool             `json:"isBookmarked"`
	IsLockedPreviewOnly     bool             `json:"isLockedPreviewOnly"`
	Links                   Links            `json:"links"`
	MetaDescription         string           `json:"metaDescription"`
	NoIndex                 bool             `json:"noIndex"`
	PreviewImage            PreviewImage     `json:"previewImage"`
	ReadingList             int64            `json:"readingList"`
	ReadingTime             float64          `json:"readingTime"`
	Recommends              int64            `json:"recommends"`
	ResponsesCreatedCount   int64            `json:"responsesCreatedCount"`
	SectionCount            int64            `json:"sectionCount"`
	SocialRecommendsCount   int64            `json:"socialRecommendsCount"`
	Subtitle                string           `json:"subtitle"`
	Tags                    []interface{}    `json:"tags"`
	Topics                  Topics           `json:"topics"`
	TotalClapCount          int64            `json:"totalClapCount"`
	UserPostRelation        UserPostRelation `json:"userPostRelation"`
	UsersBySocialRecommends []interface{}    `json:"usersBySocialRecommends"`
	WordCount               int64            `json:"wordCount"`
}

type Links struct {
	Entries     Entries `json:"entries"`
	GeneratedAt int64   `json:"generatedAt"`
	Version     string  `json:"version"`
}

type Entry struct {
	Alts       []interface{} `json:"alts"`
	HTTPStatus int64         `json:"httpStatus"`
	URL        string        `json:"url"`
}

type Entries []Entry

type PreviewImage struct {
	BackgroundSize string `json:"backgroundSize"`
	Filter         string `json:"filter"`
	Height         int64  `json:"height"`
	ImageID        string `json:"imageId"`
	OriginalHeight int64  `json:"originalHeight"`
	OriginalWidth  int64  `json:"originalWidth"`
	Strategy       string `json:"strategy"`
	Width          int64  `json:"width"`
}

type Topic struct {
	CreatedAt       int64         `json:"createdAt"`
	DeletedAt       int64         `json:"deletedAt"`
	Description     string        `json:"description"`
	Image           TopicImage    `json:"image"`
	Name            string        `json:"name"`
	RelatedTags     []interface{} `json:"relatedTags"`
	RelatedTopicIds []interface{} `json:"relatedTopicIds"`
	RelatedTopics   []interface{} `json:"relatedTopics"`
	SeoTitle        string        `json:"seoTitle"`
	Slug            string        `json:"slug"`
	TopicID         string        `json:"topicId"`
	Type            string        `json:"type"`
	Visibility      int64         `json:"visibility"`
}

type TopicImage struct {
	ID             string `json:"id"`
	OriginalHeight int64  `json:"originalHeight"`
	OriginalWidth  int64  `json:"originalWidth"`
}

type Topics []Topic

type UserPostRelation struct {
	AudioProgressSec                   int64  `json:"audioProgressSec"`
	ClapCount                          int64  `json:"clapCount"`
	CollaboratorAddedAt                int64  `json:"collaboratorAddedAt"`
	LastReadAt                         int64  `json:"lastReadAt"`
	LastReadParagraphName              string `json:"lastReadParagraphName"`
	LastReadPercentage                 int64  `json:"lastReadPercentage"`
	LastReadSectionName                string `json:"lastReadSectionName"`
	LastReadVersionID                  string `json:"lastReadVersionId"`
	NotesAddedAt                       int64  `json:"notesAddedAt"`
	PostID                             string `json:"postId"`
	PresentedCountInResponseManagement int64  `json:"presentedCountInResponseManagement"`
	PresentedCountInStream             int64  `json:"presentedCountInStream"`
	QueuedAt                           int64  `json:"queuedAt"`
	ReadAt                             int64  `json:"readAt"`
	ReadLaterAddedAt                   int64  `json:"readLaterAddedAt"`
	SeriesFirstViewedAt                int64  `json:"seriesFirstViewedAt"`
	SeriesLastViewedAt                 int64  `json:"seriesLastViewedAt"`
	SeriesUpdateNotifsOptedInAt        int64  `json:"seriesUpdateNotifsOptedInAt"`
	SubscribedAt                       int64  `json:"subscribedAt"`
	UserID                             string `json:"userId"`
	ViewedAt                           int64  `json:"viewedAt"`
	VotedAt                            int64  `json:"votedAt"`
}

type MixtapeMetadata struct {
	MediaResourceId  string `json:"mediaResourceId"`
	ThumbnailImageId string `json:"thumbnailImageId"`
	Href             string `json:"href"`
}
