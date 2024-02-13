package models

// User Info Data
type Five9UserInfo struct {
	GeneralInfo Five9UserGeneralInfo `xml:"generalInfo"`

	Roles Five9UserRoles `xml:"roles"`

	Skills []Five9UserSkill `xml:"skills"`
}

func NewFive9UserInfo() *Five9UserInfo {
	userInfo := new(Five9UserInfo)

	userInfo.Roles.Reporting.Permissions = []Five9UserPermission{
		{Type: "CanScheduleReportsViaFtp", Value: false},
		{Type: "CanAccessRecordingsColumn", Value: false},
		{Type: "NICEEnabled", Value: false},
		{Type: "CanViewStandardReports", Value: false},
		{Type: "CanViewCustomReports", Value: false},
		{Type: "CanViewScheduledReports", Value: false},
		{Type: "CanViewRecentReports", Value: false},
		{Type: "CanViewRelease7Reports", Value: false},
		{Type: "CanViewCannedReports", Value: false},
		{Type: "CanViewDashboards", Value: false},
		{Type: "CanViewAllSkills", Value: false},
		{Type: "CanViewAllGroups", Value: false},
		{Type: "CanViewSocialReports", Value: false},
	}

	userInfo.Roles.Supervisor.Permissions = []Five9UserPermission{
		{Type: "CampaignManagementStart", Value: false},
		{Type: "CampaignManagementStop", Value: false},
		{Type: "CampaignManagementReset", Value: false},
		{Type: "CampaignManagementResetDispositions", Value: false},
		{Type: "CampaignManagementResetListPositions", Value: false},
		{Type: "CampaignManagementResetAbandonCallRate", Value: false},
		{Type: "CanViewTextDetailsTab", Value: false},
		{Type: "Users", Value: false},
		{Type: "Agents", Value: false},
		{Type: "Stations", Value: false},
		{Type: "ChatSessions", Value: false},
		{Type: "Campaigns", Value: false},
		{Type: "CanAccessDashboardMenu", Value: false},
		{Type: "CallMonitoring", Value: false},
		{Type: "CampaignManagement", Value: false},
		{Type: "CanChangeDisplayLanguage", Value: false},
		{Type: "CanMonitorIdleAgents", Value: false},
		{Type: "CanSilentMonitorChats", Value: false},
		{Type: "AllSkills", Value: false},
		{Type: "CanManageComplianceData", Value: false},
		{Type: "CanMonitorEmails", Value: false},
		{Type: "CanTransferEmails", Value: false},
		{Type: "BillingInfo", Value: false},
		{Type: "BargeInMonitor", Value: false},
		{Type: "WhisperMonitor", Value: false},
		{Type: "ViewDataForAllAgentGroups", Value: false},
		{Type: "ReviewVoiceRecordings", Value: false},
		{Type: "EditAgentSkills", Value: false},
		{Type: "CanUseSupervisorSoapApi", Value: false},
		{Type: "CanAccessShowFields", Value: false},
		{Type: "NICEEnabled", Value: false},
	}

	userInfo.Roles.Agent.Permissions = []Five9UserPermission{
		{Type: "MustPickSalesforceObjectForInteractionLog"},
		{Type: "ReceiveTransfer", Value: false},
		{Type: "MakeRecordings", Value: false},
		{Type: "SendMessages", Value: false},
		{Type: "CreateChatSessions", Value: false},
		{Type: "TrainingMode", Value: false},
		{Type: "CanSelectDisplayLanguage", Value: false},
		{Type: "CanViewMissedCalls", Value: false},
		{Type: "CanViewWebAnalytics", Value: false},
		{Type: "CanTransferChatsToAgents", Value: false},
		{Type: "CanTransferChatsToSkills", Value: false},
		{Type: "CanTransferEmailsToAgents", Value: false},
		{Type: "CanTransferEmailsToSkills", Value: false},
		{Type: "CannotRemoveCRM", Value: false},
		{Type: "CanCreateChatConferenceWithAgents", Value: false},
		{Type: "CanCreateChatConferenceWithSkills", Value: false},
		{Type: "CanTransferSocialsToAgents", Value: false},
		{Type: "CanTransferSocialsToSkills", Value: false},
		{Type: "ProcessVoiceMail", Value: false},
		{Type: "CallForwarding", Value: false},
		{Type: "CannotEditSession", Value: false},
		{Type: "TransferVoiceMail", Value: false},
		{Type: "DeleteVoiceMail", Value: false},
		{Type: "AddingToDNC", Value: false},
		{Type: "DialManuallyDNC", Value: false},
		{Type: "CreateCallbacks", Value: false},
		{Type: "PlayAudioFiles", Value: false},
		{Type: "CanWrapCall", Value: false},
		{Type: "CanPlaceCallOnHold", Value: false},
		{Type: "CanParkCall", Value: false},
		{Type: "SkipCrmInPreviewDialMode", Value: false},
		{Type: "ManageAvailabilityBySkill", Value: false},
		{Type: "BrowseWebInEmbeddedBrowser", Value: false},
		{Type: "ChangePreviewPreferences", Value: false},
		{Type: "CanRejectCalls", Value: false},
		{Type: "CanConfigureAutoAnswer", Value: false},
		{Type: "MakeTransferToAgents", Value: false},
		{Type: "MakeTransferToSkills", Value: false},
		{Type: "CreateConferenceWithAgents", Value: false},
		{Type: "CreateConferenceWithSkills", Value: false},
		{Type: "RecycleDispositionAllowed", Value: false},
		{Type: "MakeTransferToInboundCampaigns", Value: false},
		{Type: "MakeTransferToExternalCalls", Value: false},
		{Type: "CreateConferenceWithInboundCampaigns", Value: false},
		{Type: "CreateConferenceWithExternalCalls", Value: false},
		{Type: "MakeCallToSkills", Value: false},
		{Type: "MakeCallToAgents", Value: false},
		{Type: "MakeCallToExternalCalls", Value: false},
		{Type: "MakeCallToSpeedDialNumber", Value: false},
		{Type: "MakeTransferToSpeedDialNumber", Value: false},
		{Type: "CreateConferenceWithSpeedDialNumber", Value: false},
		{Type: "NICEEnabled", Value: false},
		{Type: "ScreenRecording", Value: false},
	}

	return userInfo
}

// Sub Data Types
type Five9UserGeneralInfo struct {
	Id                 int    `xml:"id"`
	Active             bool   `xml:"active"`
	CanChangePassword  bool   `xml:"canChangePassword"`
	MustChangePassword bool   `xml:"mustChangePassword"`
	Email              string `xml:"EMail"`
	// Extension         int    `xml:"extension"`
	FederationId    string `xml:"federationId"`
	UserName        string `xml:"userName"`
	FirstName       string `xml:"firstName"`
	LastName        string `xml:"lastName"`
	FullName        string `xml:"fullName"`
	Password        string `xml:"password,omitempty"`
	MediaTypeConfig struct {
		MediaTypes []Five9UserMediaTypeConfig `xml:"mediaTypes"`
	} `xml:"mediaTypeConfig,omitempty"`
}

type Five9UserRoles struct {
	Admin      Five9UserAdminRole      `xml:"admin"`
	Agent      Five9UserAgentRole      `xml:"agent"`
	Supervisor Five9UserSupervisorRole `xml:"supervisor"`
	Reporting  Five9UserReportingRole  `xml:"reporting"`
}

type Five9UserSkill struct {
	Id        int    `xml:"id"`
	Level     int    `xml:"level"`
	SkillName string `xml:"skillName"`
	UserName  string `xml:"userName"`
}

type Five9UserMediaTypeConfig struct {
	Enabled            bool   `xml:"enabled"`
	IntelligentRouting bool   `xml:"intlligentRouting"`
	MaxAllowed         int    `xml:"maxAlowed"`
	Type               string `xml:"type"`
}

type Five9UserPermission struct {
	Type  string `xml:"type"`
	Value bool   `xml:"value"`
}

type Five9UserAdminRole struct {
	Permissions []Five9UserPermission `xml:"permissions,omitempty"`
}

type Five9UserAgentRole struct {
	Permissions []Five9UserPermission `xml:"permissions,omitempty"`
}

type Five9UserSupervisorRole struct {
	Permissions []Five9UserPermission `xml:"permissions,omitempty"`
}

type Five9UserReportingRole struct {
	Permissions []Five9UserPermission `xml:"permissions,omitempty"`
}

func GetSupportedPermissions() []string {
	return []string{}
}
