package main

import (
	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"golang.org/x/exp/shiny/materialdesign/icons"
	"image/color"
)

var (
	list = &layout.List{
		Axis:      layout.Vertical,
		Alignment: layout.Middle,
	}
)

func main() {

	gofont.Register()
	th := material.NewTheme()

	action := make(map[string]*material.Icon)
	action["Action3DRotation"] = mustIcon(material.NewIcon(icons.Action3DRotation))
	action["ActionAccessibility"] = mustIcon(material.NewIcon(icons.ActionAccessibility))
	action["ActionAccessible"] = mustIcon(material.NewIcon(icons.ActionAccessible))
	action["ActionAccountBalance"] = mustIcon(material.NewIcon(icons.ActionAccountBalance))
	action["ActionAccountBalanceWallet"] = mustIcon(material.NewIcon(icons.ActionAccountBalanceWallet))
	action["ActionAccountBox"] = mustIcon(material.NewIcon(icons.ActionAccountBox))
	action["ActionAccountCircle"] = mustIcon(material.NewIcon(icons.ActionAccountCircle))
	action["ActionAddShoppingCart"] = mustIcon(material.NewIcon(icons.ActionAddShoppingCart))
	action["ActionAlarm"] = mustIcon(material.NewIcon(icons.ActionAlarm))
	action["ActionAlarmAdd"] = mustIcon(material.NewIcon(icons.ActionAlarmAdd))
	action["ActionAlarmOff"] = mustIcon(material.NewIcon(icons.ActionAlarmOff))
	action["ActionAlarmOn"] = mustIcon(material.NewIcon(icons.ActionAlarmOn))
	action["ActionAllOut"] = mustIcon(material.NewIcon(icons.ActionAllOut))
	action["ActionAndroid"] = mustIcon(material.NewIcon(icons.ActionAndroid))
	action["ActionAnnouncement"] = mustIcon(material.NewIcon(icons.ActionAnnouncement))
	action["ActionAspectRatio"] = mustIcon(material.NewIcon(icons.ActionAspectRatio))
	action["ActionAssessment"] = mustIcon(material.NewIcon(icons.ActionAssessment))
	action["ActionAssignment"] = mustIcon(material.NewIcon(icons.ActionAssignment))
	action["ActionAssignmentInd"] = mustIcon(material.NewIcon(icons.ActionAssignmentInd))
	action["ActionAssignmentLate"] = mustIcon(material.NewIcon(icons.ActionAssignmentLate))
	action["ActionAssignmentReturn"] = mustIcon(material.NewIcon(icons.ActionAssignmentReturn))
	action["ActionAssignmentReturned"] = mustIcon(material.NewIcon(icons.ActionAssignmentReturned))
	action["ActionAssignmentTurnedIn"] = mustIcon(material.NewIcon(icons.ActionAssignmentTurnedIn))
	action["ActionAutorenew"] = mustIcon(material.NewIcon(icons.ActionAutorenew))
	action["ActionBackup"] = mustIcon(material.NewIcon(icons.ActionBackup))
	action["ActionBook"] = mustIcon(material.NewIcon(icons.ActionBook))
	action["ActionBookmark"] = mustIcon(material.NewIcon(icons.ActionBookmark))
	action["ActionBookmarkBorder"] = mustIcon(material.NewIcon(icons.ActionBookmarkBorder))
	action["ActionBugReport"] = mustIcon(material.NewIcon(icons.ActionBugReport))
	action["ActionBuild"] = mustIcon(material.NewIcon(icons.ActionBuild))
	action["ActionCached"] = mustIcon(material.NewIcon(icons.ActionCached))
	action["ActionCameraEnhance"] = mustIcon(material.NewIcon(icons.ActionCameraEnhance))
	action["ActionCardGiftcard"] = mustIcon(material.NewIcon(icons.ActionCardGiftcard))
	action["ActionCardMembership"] = mustIcon(material.NewIcon(icons.ActionCardMembership))
	action["ActionCardTravel"] = mustIcon(material.NewIcon(icons.ActionCardTravel))
	action["ActionChangeHistory"] = mustIcon(material.NewIcon(icons.ActionChangeHistory))
	action["ActionCheckCircle"] = mustIcon(material.NewIcon(icons.ActionCheckCircle))
	action["ActionChromeReaderMode"] = mustIcon(material.NewIcon(icons.ActionChromeReaderMode))
	action["ActionClass"] = mustIcon(material.NewIcon(icons.ActionClass))
	action["ActionCode"] = mustIcon(material.NewIcon(icons.ActionCode))
	action["ActionCompareArrows"] = mustIcon(material.NewIcon(icons.ActionCompareArrows))
	action["ActionCopyright"] = mustIcon(material.NewIcon(icons.ActionCopyright))
	action["ActionCreditCard"] = mustIcon(material.NewIcon(icons.ActionCreditCard))
	action["ActionDashboard"] = mustIcon(material.NewIcon(icons.ActionDashboard))
	action["ActionDateRange"] = mustIcon(material.NewIcon(icons.ActionDateRange))
	action["ActionDelete"] = mustIcon(material.NewIcon(icons.ActionDelete))
	action["ActionDeleteForever"] = mustIcon(material.NewIcon(icons.ActionDeleteForever))
	action["ActionDescription"] = mustIcon(material.NewIcon(icons.ActionDescription))
	action["ActionDNS"] = mustIcon(material.NewIcon(icons.ActionDNS))
	action["ActionDone"] = mustIcon(material.NewIcon(icons.ActionDone))
	action["ActionDoneAll"] = mustIcon(material.NewIcon(icons.ActionDoneAll))
	action["ActionDonutLarge"] = mustIcon(material.NewIcon(icons.ActionDonutLarge))
	action["ActionDonutSmall"] = mustIcon(material.NewIcon(icons.ActionDonutSmall))
	action["ActionEject"] = mustIcon(material.NewIcon(icons.ActionEject))
	action["ActionEuroSymbol"] = mustIcon(material.NewIcon(icons.ActionEuroSymbol))
	action["ActionEvent"] = mustIcon(material.NewIcon(icons.ActionEvent))
	action["ActionEventSeat"] = mustIcon(material.NewIcon(icons.ActionEventSeat))
	action["ActionExitToApp"] = mustIcon(material.NewIcon(icons.ActionExitToApp))
	action["ActionExplore"] = mustIcon(material.NewIcon(icons.ActionExplore))
	action["ActionExtension"] = mustIcon(material.NewIcon(icons.ActionExtension))
	action["ActionFace"] = mustIcon(material.NewIcon(icons.ActionFace))
	action["ActionFavorite"] = mustIcon(material.NewIcon(icons.ActionFavorite))
	action["ActionFavoriteBorder"] = mustIcon(material.NewIcon(icons.ActionFavoriteBorder))
	action["ActionFeedback"] = mustIcon(material.NewIcon(icons.ActionFeedback))
	action["ActionFindInPage"] = mustIcon(material.NewIcon(icons.ActionFindInPage))
	action["ActionFindReplace"] = mustIcon(material.NewIcon(icons.ActionFindReplace))
	action["ActionFingerprint"] = mustIcon(material.NewIcon(icons.ActionFingerprint))
	action["ActionFlightLand"] = mustIcon(material.NewIcon(icons.ActionFlightLand))
	action["ActionFlightTakeoff"] = mustIcon(material.NewIcon(icons.ActionFlightTakeoff))
	action["ActionFlipToBack"] = mustIcon(material.NewIcon(icons.ActionFlipToBack))
	action["ActionFlipToFront"] = mustIcon(material.NewIcon(icons.ActionFlipToFront))
	action["ActionGTranslate"] = mustIcon(material.NewIcon(icons.ActionGTranslate))
	action["ActionGavel"] = mustIcon(material.NewIcon(icons.ActionGavel))
	action["ActionGetApp"] = mustIcon(material.NewIcon(icons.ActionGetApp))
	action["ActionGIF"] = mustIcon(material.NewIcon(icons.ActionGIF))
	action["ActionGrade"] = mustIcon(material.NewIcon(icons.ActionGrade))
	action["ActionGroupWork"] = mustIcon(material.NewIcon(icons.ActionGroupWork))
	action["ActionHelp"] = mustIcon(material.NewIcon(icons.ActionHelp))
	action["ActionHelpOutline"] = mustIcon(material.NewIcon(icons.ActionHelpOutline))
	action["ActionHighlightOff"] = mustIcon(material.NewIcon(icons.ActionHighlightOff))
	action["ActionHistory"] = mustIcon(material.NewIcon(icons.ActionHistory))
	action["ActionHome"] = mustIcon(material.NewIcon(icons.ActionHome))
	action["ActionHourglassEmpty"] = mustIcon(material.NewIcon(icons.ActionHourglassEmpty))
	action["ActionHourglassFull"] = mustIcon(material.NewIcon(icons.ActionHourglassFull))
	action["ActionHTTP"] = mustIcon(material.NewIcon(icons.ActionHTTP))
	action["ActionHTTPS"] = mustIcon(material.NewIcon(icons.ActionHTTPS))
	action["ActionImportantDevices"] = mustIcon(material.NewIcon(icons.ActionImportantDevices))
	action["ActionInfo"] = mustIcon(material.NewIcon(icons.ActionInfo))
	action["ActionInfoOutline"] = mustIcon(material.NewIcon(icons.ActionInfoOutline))
	action["ActionInput"] = mustIcon(material.NewIcon(icons.ActionInput))
	action["ActionInvertColors"] = mustIcon(material.NewIcon(icons.ActionInvertColors))
	action["ActionLabel"] = mustIcon(material.NewIcon(icons.ActionLabel))
	action["ActionLabelOutline"] = mustIcon(material.NewIcon(icons.ActionLabelOutline))
	action["ActionLanguage"] = mustIcon(material.NewIcon(icons.ActionLanguage))
	action["ActionLaunch"] = mustIcon(material.NewIcon(icons.ActionLaunch))
	action["ActionLightbulbOutline"] = mustIcon(material.NewIcon(icons.ActionLightbulbOutline))
	action["ActionLineStyle"] = mustIcon(material.NewIcon(icons.ActionLineStyle))
	action["ActionLineWeight"] = mustIcon(material.NewIcon(icons.ActionLineWeight))
	action["ActionList"] = mustIcon(material.NewIcon(icons.ActionList))
	action["ActionLock"] = mustIcon(material.NewIcon(icons.ActionLock))
	action["ActionLockOpen"] = mustIcon(material.NewIcon(icons.ActionLockOpen))
	action["ActionLockOutline"] = mustIcon(material.NewIcon(icons.ActionLockOutline))
	action["ActionLoyalty"] = mustIcon(material.NewIcon(icons.ActionLoyalty))
	action["ActionMarkUnreadMailbox"] = mustIcon(material.NewIcon(icons.ActionMarkUnreadMailbox))
	action["ActionMotorcycle"] = mustIcon(material.NewIcon(icons.ActionMotorcycle))
	action["ActionNoteAdd"] = mustIcon(material.NewIcon(icons.ActionNoteAdd))
	action["ActionOfflinePin"] = mustIcon(material.NewIcon(icons.ActionOfflinePin))
	action["ActionOpacity"] = mustIcon(material.NewIcon(icons.ActionOpacity))
	action["ActionOpenInBrowser"] = mustIcon(material.NewIcon(icons.ActionOpenInBrowser))
	action["ActionOpenInNew"] = mustIcon(material.NewIcon(icons.ActionOpenInNew))
	action["ActionOpenWith"] = mustIcon(material.NewIcon(icons.ActionOpenWith))
	action["ActionPageview"] = mustIcon(material.NewIcon(icons.ActionPageview))
	action["ActionPanTool"] = mustIcon(material.NewIcon(icons.ActionPanTool))
	action["ActionPayment"] = mustIcon(material.NewIcon(icons.ActionPayment))
	action["ActionPermCameraMic"] = mustIcon(material.NewIcon(icons.ActionPermCameraMic))
	action["ActionPermContactCalendar"] = mustIcon(material.NewIcon(icons.ActionPermContactCalendar))
	action["ActionPermDataSetting"] = mustIcon(material.NewIcon(icons.ActionPermDataSetting))
	action["ActionPermDeviceInformation"] = mustIcon(material.NewIcon(icons.ActionPermDeviceInformation))
	action["ActionPermIdentity"] = mustIcon(material.NewIcon(icons.ActionPermIdentity))
	action["ActionPermMedia"] = mustIcon(material.NewIcon(icons.ActionPermMedia))
	action["ActionPermPhoneMsg"] = mustIcon(material.NewIcon(icons.ActionPermPhoneMsg))
	action["ActionPermScanWiFi"] = mustIcon(material.NewIcon(icons.ActionPermScanWiFi))
	action["ActionPets"] = mustIcon(material.NewIcon(icons.ActionPets))
	action["ActionPictureInPicture"] = mustIcon(material.NewIcon(icons.ActionPictureInPicture))
	action["ActionPictureInPictureAlt"] = mustIcon(material.NewIcon(icons.ActionPictureInPictureAlt))
	action["ActionPlayForWork"] = mustIcon(material.NewIcon(icons.ActionPlayForWork))
	action["ActionPolymer"] = mustIcon(material.NewIcon(icons.ActionPolymer))
	action["ActionPowerSettingsNew"] = mustIcon(material.NewIcon(icons.ActionPowerSettingsNew))
	action["ActionPregnantWoman"] = mustIcon(material.NewIcon(icons.ActionPregnantWoman))
	action["ActionPrint"] = mustIcon(material.NewIcon(icons.ActionPrint))
	action["ActionQueryBuilder"] = mustIcon(material.NewIcon(icons.ActionQueryBuilder))
	action["ActionQuestionAnswer"] = mustIcon(material.NewIcon(icons.ActionQuestionAnswer))
	action["ActionReceipt"] = mustIcon(material.NewIcon(icons.ActionReceipt))
	action["ActionRecordVoiceOver"] = mustIcon(material.NewIcon(icons.ActionRecordVoiceOver))
	action["ActionRedeem"] = mustIcon(material.NewIcon(icons.ActionRedeem))
	action["ActionRemoveShoppingCart"] = mustIcon(material.NewIcon(icons.ActionRemoveShoppingCart))
	action["ActionReorder"] = mustIcon(material.NewIcon(icons.ActionReorder))
	action["ActionReportProblem"] = mustIcon(material.NewIcon(icons.ActionReportProblem))
	action["ActionRestore"] = mustIcon(material.NewIcon(icons.ActionRestore))
	action["ActionRestorePage"] = mustIcon(material.NewIcon(icons.ActionRestorePage))
	action["ActionRoom"] = mustIcon(material.NewIcon(icons.ActionRoom))
	action["ActionRoundedCorner"] = mustIcon(material.NewIcon(icons.ActionRoundedCorner))
	action["ActionRowing"] = mustIcon(material.NewIcon(icons.ActionRowing))
	action["ActionSchedule"] = mustIcon(material.NewIcon(icons.ActionSchedule))
	action["ActionSearch"] = mustIcon(material.NewIcon(icons.ActionSearch))
	action["ActionSettings"] = mustIcon(material.NewIcon(icons.ActionSettings))
	action["ActionSettingsApplications"] = mustIcon(material.NewIcon(icons.ActionSettingsApplications))
	action["ActionSettingsBackupRestore"] = mustIcon(material.NewIcon(icons.ActionSettingsBackupRestore))
	action["ActionSettingsBluetooth"] = mustIcon(material.NewIcon(icons.ActionSettingsBluetooth))
	action["ActionSettingsBrightness"] = mustIcon(material.NewIcon(icons.ActionSettingsBrightness))
	action["ActionSettingsCell"] = mustIcon(material.NewIcon(icons.ActionSettingsCell))
	action["ActionSettingsEthernet"] = mustIcon(material.NewIcon(icons.ActionSettingsEthernet))
	action["ActionSettingsInputAntenna"] = mustIcon(material.NewIcon(icons.ActionSettingsInputAntenna))
	action["ActionSettingsInputComponent"] = mustIcon(material.NewIcon(icons.ActionSettingsInputComponent))
	action["ActionSettingsInputComposite"] = mustIcon(material.NewIcon(icons.ActionSettingsInputComposite))
	action["ActionSettingsInputHDMI"] = mustIcon(material.NewIcon(icons.ActionSettingsInputHDMI))
	action["ActionSettingsInputSVideo"] = mustIcon(material.NewIcon(icons.ActionSettingsInputSVideo))
	action["ActionSettingsOverscan"] = mustIcon(material.NewIcon(icons.ActionSettingsOverscan))
	action["ActionSettingsPhone"] = mustIcon(material.NewIcon(icons.ActionSettingsPhone))
	action["ActionSettingsPower"] = mustIcon(material.NewIcon(icons.ActionSettingsPower))
	action["ActionSettingsRemote"] = mustIcon(material.NewIcon(icons.ActionSettingsRemote))
	action["ActionSettingsVoice"] = mustIcon(material.NewIcon(icons.ActionSettingsVoice))
	action["ActionShop"] = mustIcon(material.NewIcon(icons.ActionShop))
	action["ActionShopTwo"] = mustIcon(material.NewIcon(icons.ActionShopTwo))
	action["ActionShoppingBasket"] = mustIcon(material.NewIcon(icons.ActionShoppingBasket))
	action["ActionShoppingCart"] = mustIcon(material.NewIcon(icons.ActionShoppingCart))
	action["ActionSpeakerNotes"] = mustIcon(material.NewIcon(icons.ActionSpeakerNotes))
	action["ActionSpeakerNotesOff"] = mustIcon(material.NewIcon(icons.ActionSpeakerNotesOff))
	action["ActionSpellcheck"] = mustIcon(material.NewIcon(icons.ActionSpellcheck))
	action["ActionStarRate"] = mustIcon(material.NewIcon(icons.ActionStarRate))
	action["ActionStars"] = mustIcon(material.NewIcon(icons.ActionStars))
	action["ActionStore"] = mustIcon(material.NewIcon(icons.ActionStore))
	action["ActionSubject"] = mustIcon(material.NewIcon(icons.ActionSubject))
	action["ActionSupervisorAccount"] = mustIcon(material.NewIcon(icons.ActionSupervisorAccount))
	action["ActionSwapHoriz"] = mustIcon(material.NewIcon(icons.ActionSwapHoriz))
	action["ActionSwapVert"] = mustIcon(material.NewIcon(icons.ActionSwapVert))
	action["ActionSwapVerticalCircle"] = mustIcon(material.NewIcon(icons.ActionSwapVerticalCircle))
	action["ActionSystemUpdateAlt"] = mustIcon(material.NewIcon(icons.ActionSystemUpdateAlt))
	action["ActionTab"] = mustIcon(material.NewIcon(icons.ActionTab))
	action["ActionTabUnselected"] = mustIcon(material.NewIcon(icons.ActionTabUnselected))
	action["ActionTheaters"] = mustIcon(material.NewIcon(icons.ActionTheaters))
	action["ActionThumbDown"] = mustIcon(material.NewIcon(icons.ActionThumbDown))
	action["ActionThumbUp"] = mustIcon(material.NewIcon(icons.ActionThumbUp))
	action["ActionThumbsUpDown"] = mustIcon(material.NewIcon(icons.ActionThumbsUpDown))
	action["ActionTimeline"] = mustIcon(material.NewIcon(icons.ActionTimeline))
	action["ActionTOC"] = mustIcon(material.NewIcon(icons.ActionTOC))
	action["ActionToday"] = mustIcon(material.NewIcon(icons.ActionToday))
	action["ActionToll"] = mustIcon(material.NewIcon(icons.ActionToll))
	action["ActionTouchApp"] = mustIcon(material.NewIcon(icons.ActionTouchApp))
	action["ActionTrackChanges"] = mustIcon(material.NewIcon(icons.ActionTrackChanges))
	action["ActionTranslate"] = mustIcon(material.NewIcon(icons.ActionTranslate))
	action["ActionTrendingDown"] = mustIcon(material.NewIcon(icons.ActionTrendingDown))
	action["ActionTrendingFlat"] = mustIcon(material.NewIcon(icons.ActionTrendingFlat))
	action["ActionTrendingUp"] = mustIcon(material.NewIcon(icons.ActionTrendingUp))
	action["ActionTurnedIn"] = mustIcon(material.NewIcon(icons.ActionTurnedIn))
	action["ActionTurnedInNot"] = mustIcon(material.NewIcon(icons.ActionTurnedInNot))
	action["ActionUpdate"] = mustIcon(material.NewIcon(icons.ActionUpdate))
	action["ActionVerifiedUser"] = mustIcon(material.NewIcon(icons.ActionVerifiedUser))
	action["ActionViewAgenda"] = mustIcon(material.NewIcon(icons.ActionViewAgenda))
	action["ActionViewArray"] = mustIcon(material.NewIcon(icons.ActionViewArray))
	action["ActionViewCarousel"] = mustIcon(material.NewIcon(icons.ActionViewCarousel))
	action["ActionViewColumn"] = mustIcon(material.NewIcon(icons.ActionViewColumn))
	action["ActionViewDay"] = mustIcon(material.NewIcon(icons.ActionViewDay))
	action["ActionViewHeadline"] = mustIcon(material.NewIcon(icons.ActionViewHeadline))
	action["ActionViewList"] = mustIcon(material.NewIcon(icons.ActionViewList))
	action["ActionViewModule"] = mustIcon(material.NewIcon(icons.ActionViewModule))
	action["ActionViewQuilt"] = mustIcon(material.NewIcon(icons.ActionViewQuilt))
	action["ActionViewStream"] = mustIcon(material.NewIcon(icons.ActionViewStream))
	action["ActionViewWeek"] = mustIcon(material.NewIcon(icons.ActionViewWeek))
	action["ActionVisibility"] = mustIcon(material.NewIcon(icons.ActionVisibility))
	action["ActionVisibilityOff"] = mustIcon(material.NewIcon(icons.ActionVisibilityOff))
	action["ActionWatchLater"] = mustIcon(material.NewIcon(icons.ActionWatchLater))
	action["ActionWork"] = mustIcon(material.NewIcon(icons.ActionWork))
	action["ActionYoutubeSearchedFor"] = mustIcon(material.NewIcon(icons.ActionYoutubeSearchedFor))
	action["ActionZoomIn"] = mustIcon(material.NewIcon(icons.ActionZoomIn))
	action["ActionZoomOut"] = mustIcon(material.NewIcon(icons.ActionZoomOut))

	av := make(map[string]*material.Icon)
	av["AlertAddAlert"] = mustIcon(material.NewIcon(icons.AlertAddAlert))
	av["AlertError"] = mustIcon(material.NewIcon(icons.AlertError))
	av["AlertErrorOutline"] = mustIcon(material.NewIcon(icons.AlertErrorOutline))
	av["AlertWarning"] = mustIcon(material.NewIcon(icons.AlertWarning))
	av["AVAddToQueue"] = mustIcon(material.NewIcon(icons.AVAddToQueue))
	av["AVAirplay"] = mustIcon(material.NewIcon(icons.AVAirplay))
	av["AVAlbum"] = mustIcon(material.NewIcon(icons.AVAlbum))
	av["AVArtTrack"] = mustIcon(material.NewIcon(icons.AVArtTrack))
	av["AVAVTimer"] = mustIcon(material.NewIcon(icons.AVAVTimer))
	av["AVBrandingWatermark"] = mustIcon(material.NewIcon(icons.AVBrandingWatermark))
	av["AVCallToAction"] = mustIcon(material.NewIcon(icons.AVCallToAction))
	av["AVClosedCaption"] = mustIcon(material.NewIcon(icons.AVClosedCaption))
	av["AVEqualizer"] = mustIcon(material.NewIcon(icons.AVEqualizer))
	av["AVExplicit"] = mustIcon(material.NewIcon(icons.AVExplicit))
	av["AVFastForward"] = mustIcon(material.NewIcon(icons.AVFastForward))
	av["AVFastRewind"] = mustIcon(material.NewIcon(icons.AVFastRewind))
	av["AVFeaturedPlayList"] = mustIcon(material.NewIcon(icons.AVFeaturedPlayList))
	av["AVFeaturedVideo"] = mustIcon(material.NewIcon(icons.AVFeaturedVideo))
	av["AVFiberDVR"] = mustIcon(material.NewIcon(icons.AVFiberDVR))
	av["AVFiberManualRecord"] = mustIcon(material.NewIcon(icons.AVFiberManualRecord))
	av["AVFiberNew"] = mustIcon(material.NewIcon(icons.AVFiberNew))
	av["AVFiberPin"] = mustIcon(material.NewIcon(icons.AVFiberPin))
	av["AVFiberSmartRecord"] = mustIcon(material.NewIcon(icons.AVFiberSmartRecord))
	av["AVForward10"] = mustIcon(material.NewIcon(icons.AVForward10))
	av["AVForward30"] = mustIcon(material.NewIcon(icons.AVForward30))
	av["AVForward5"] = mustIcon(material.NewIcon(icons.AVForward5))
	av["AVGames"] = mustIcon(material.NewIcon(icons.AVGames))
	av["AVHD"] = mustIcon(material.NewIcon(icons.AVHD))
	av["AVHearing"] = mustIcon(material.NewIcon(icons.AVHearing))
	av["AVHighQuality"] = mustIcon(material.NewIcon(icons.AVHighQuality))
	av["AVLibraryAdd"] = mustIcon(material.NewIcon(icons.AVLibraryAdd))
	av["AVLibraryBooks"] = mustIcon(material.NewIcon(icons.AVLibraryBooks))
	av["AVLibraryMusic"] = mustIcon(material.NewIcon(icons.AVLibraryMusic))
	av["AVLoop"] = mustIcon(material.NewIcon(icons.AVLoop))
	av["AVMic"] = mustIcon(material.NewIcon(icons.AVMic))
	av["AVMicNone"] = mustIcon(material.NewIcon(icons.AVMicNone))
	av["AVMicOff"] = mustIcon(material.NewIcon(icons.AVMicOff))
	av["AVMovie"] = mustIcon(material.NewIcon(icons.AVMovie))
	av["AVMusicVideo"] = mustIcon(material.NewIcon(icons.AVMusicVideo))
	av["AVNewReleases"] = mustIcon(material.NewIcon(icons.AVNewReleases))
	av["AVNotInterested"] = mustIcon(material.NewIcon(icons.AVNotInterested))
	av["AVNote"] = mustIcon(material.NewIcon(icons.AVNote))
	av["AVPause"] = mustIcon(material.NewIcon(icons.AVPause))
	av["AVPauseCircleFilled"] = mustIcon(material.NewIcon(icons.AVPauseCircleFilled))
	av["AVPauseCircleOutline"] = mustIcon(material.NewIcon(icons.AVPauseCircleOutline))
	av["AVPlayArrow"] = mustIcon(material.NewIcon(icons.AVPlayArrow))
	av["AVPlayCircleFilled"] = mustIcon(material.NewIcon(icons.AVPlayCircleFilled))
	av["AVPlayCircleOutline"] = mustIcon(material.NewIcon(icons.AVPlayCircleOutline))
	av["AVPlaylistAdd"] = mustIcon(material.NewIcon(icons.AVPlaylistAdd))
	av["AVPlaylistAddCheck"] = mustIcon(material.NewIcon(icons.AVPlaylistAddCheck))
	av["AVPlaylistPlay"] = mustIcon(material.NewIcon(icons.AVPlaylistPlay))
	av["AVQueue"] = mustIcon(material.NewIcon(icons.AVQueue))
	av["AVQueueMusic"] = mustIcon(material.NewIcon(icons.AVQueueMusic))
	av["AVQueuePlayNext"] = mustIcon(material.NewIcon(icons.AVQueuePlayNext))
	av["AVRadio"] = mustIcon(material.NewIcon(icons.AVRadio))
	av["AVRecentActors"] = mustIcon(material.NewIcon(icons.AVRecentActors))
	av["AVRemoveFromQueue"] = mustIcon(material.NewIcon(icons.AVRemoveFromQueue))
	av["AVRepeat"] = mustIcon(material.NewIcon(icons.AVRepeat))
	av["AVRepeatOne"] = mustIcon(material.NewIcon(icons.AVRepeatOne))
	av["AVReplay"] = mustIcon(material.NewIcon(icons.AVReplay))
	av["AVReplay10"] = mustIcon(material.NewIcon(icons.AVReplay10))
	av["AVReplay30"] = mustIcon(material.NewIcon(icons.AVReplay30))
	av["AVReplay5"] = mustIcon(material.NewIcon(icons.AVReplay5))
	av["AVShuffle"] = mustIcon(material.NewIcon(icons.AVShuffle))
	av["AVSkipNext"] = mustIcon(material.NewIcon(icons.AVSkipNext))
	av["AVSkipPrevious"] = mustIcon(material.NewIcon(icons.AVSkipPrevious))
	av["AVSlowMotionVideo"] = mustIcon(material.NewIcon(icons.AVSlowMotionVideo))
	av["AVSnooze"] = mustIcon(material.NewIcon(icons.AVSnooze))
	av["AVSortByAlpha"] = mustIcon(material.NewIcon(icons.AVSortByAlpha))
	av["AVStop"] = mustIcon(material.NewIcon(icons.AVStop))
	av["AVSubscriptions"] = mustIcon(material.NewIcon(icons.AVSubscriptions))
	av["AVSubtitles"] = mustIcon(material.NewIcon(icons.AVSubtitles))
	av["AVSurroundSound"] = mustIcon(material.NewIcon(icons.AVSurroundSound))
	av["AVVideoCall"] = mustIcon(material.NewIcon(icons.AVVideoCall))
	av["AVVideoLabel"] = mustIcon(material.NewIcon(icons.AVVideoLabel))
	av["AVVideoLibrary"] = mustIcon(material.NewIcon(icons.AVVideoLibrary))
	av["AVVideocam"] = mustIcon(material.NewIcon(icons.AVVideocam))
	av["AVVideocamOff"] = mustIcon(material.NewIcon(icons.AVVideocamOff))
	av["AVVolumeDown"] = mustIcon(material.NewIcon(icons.AVVolumeDown))
	av["AVVolumeMute"] = mustIcon(material.NewIcon(icons.AVVolumeMute))
	av["AVVolumeOff"] = mustIcon(material.NewIcon(icons.AVVolumeOff))
	av["AVVolumeUp"] = mustIcon(material.NewIcon(icons.AVVolumeUp))
	av["AVWeb"] = mustIcon(material.NewIcon(icons.AVWeb))
	av["AVWebAsset"] = mustIcon(material.NewIcon(icons.AVWebAsset))

	communication := make(map[string]*material.Icon)
	communication["CommunicationBusiness"] = mustIcon(material.NewIcon(icons.CommunicationBusiness))
	communication["CommunicationCall"] = mustIcon(material.NewIcon(icons.CommunicationCall))
	communication["CommunicationCallEnd"] = mustIcon(material.NewIcon(icons.CommunicationCallEnd))
	communication["CommunicationCallMade"] = mustIcon(material.NewIcon(icons.CommunicationCallMade))
	communication["CommunicationCallMerge"] = mustIcon(material.NewIcon(icons.CommunicationCallMerge))
	communication["CommunicationCallMissed"] = mustIcon(material.NewIcon(icons.CommunicationCallMissed))
	communication["CommunicationCallMissedOutgoing"] = mustIcon(material.NewIcon(icons.CommunicationCallMissedOutgoing))
	communication["CommunicationCallReceived"] = mustIcon(material.NewIcon(icons.CommunicationCallReceived))
	communication["CommunicationCallSplit"] = mustIcon(material.NewIcon(icons.CommunicationCallSplit))
	communication["CommunicationChat"] = mustIcon(material.NewIcon(icons.CommunicationChat))
	communication["CommunicationChatBubble"] = mustIcon(material.NewIcon(icons.CommunicationChatBubble))
	communication["CommunicationChatBubbleOutline"] = mustIcon(material.NewIcon(icons.CommunicationChatBubbleOutline))
	communication["CommunicationClearAll"] = mustIcon(material.NewIcon(icons.CommunicationClearAll))
	communication["CommunicationComment"] = mustIcon(material.NewIcon(icons.CommunicationComment))
	communication["CommunicationContactMail"] = mustIcon(material.NewIcon(icons.CommunicationContactMail))
	communication["CommunicationContactPhone"] = mustIcon(material.NewIcon(icons.CommunicationContactPhone))
	communication["CommunicationContacts"] = mustIcon(material.NewIcon(icons.CommunicationContacts))
	communication["CommunicationDialerSIP"] = mustIcon(material.NewIcon(icons.CommunicationDialerSIP))
	communication["CommunicationDialpad"] = mustIcon(material.NewIcon(icons.CommunicationDialpad))
	communication["CommunicationEmail"] = mustIcon(material.NewIcon(icons.CommunicationEmail))
	communication["CommunicationForum"] = mustIcon(material.NewIcon(icons.CommunicationForum))
	communication["CommunicationImportContacts"] = mustIcon(material.NewIcon(icons.CommunicationImportContacts))
	communication["CommunicationImportExport"] = mustIcon(material.NewIcon(icons.CommunicationImportExport))
	communication["CommunicationInvertColorsOff"] = mustIcon(material.NewIcon(icons.CommunicationInvertColorsOff))
	communication["CommunicationLiveHelp"] = mustIcon(material.NewIcon(icons.CommunicationLiveHelp))
	communication["CommunicationLocationOff"] = mustIcon(material.NewIcon(icons.CommunicationLocationOff))
	communication["CommunicationLocationOn"] = mustIcon(material.NewIcon(icons.CommunicationLocationOn))
	communication["CommunicationMailOutline"] = mustIcon(material.NewIcon(icons.CommunicationMailOutline))
	communication["CommunicationMessage"] = mustIcon(material.NewIcon(icons.CommunicationMessage))
	communication["CommunicationNoSIM"] = mustIcon(material.NewIcon(icons.CommunicationNoSIM))
	communication["CommunicationPhone"] = mustIcon(material.NewIcon(icons.CommunicationPhone))
	communication["CommunicationPhoneLinkErase"] = mustIcon(material.NewIcon(icons.CommunicationPhoneLinkErase))
	communication["CommunicationPhoneLinkLock"] = mustIcon(material.NewIcon(icons.CommunicationPhoneLinkLock))
	communication["CommunicationPhoneLinkRing"] = mustIcon(material.NewIcon(icons.CommunicationPhoneLinkRing))
	communication["CommunicationPhoneLinkSetup"] = mustIcon(material.NewIcon(icons.CommunicationPhoneLinkSetup))
	communication["CommunicationPortableWiFiOff"] = mustIcon(material.NewIcon(icons.CommunicationPortableWiFiOff))
	communication["CommunicationPresentToAll"] = mustIcon(material.NewIcon(icons.CommunicationPresentToAll))
	communication["CommunicationRingVolume"] = mustIcon(material.NewIcon(icons.CommunicationRingVolume))
	communication["CommunicationRSSFeed"] = mustIcon(material.NewIcon(icons.CommunicationRSSFeed))
	communication["CommunicationScreenShare"] = mustIcon(material.NewIcon(icons.CommunicationScreenShare))
	communication["CommunicationSpeakerPhone"] = mustIcon(material.NewIcon(icons.CommunicationSpeakerPhone))
	communication["CommunicationStayCurrentLandscape"] = mustIcon(material.NewIcon(icons.CommunicationStayCurrentLandscape))
	communication["CommunicationStayCurrentPortrait"] = mustIcon(material.NewIcon(icons.CommunicationStayCurrentPortrait))
	communication["CommunicationStayPrimaryLandscape"] = mustIcon(material.NewIcon(icons.CommunicationStayPrimaryLandscape))
	communication["CommunicationStayPrimaryPortrait"] = mustIcon(material.NewIcon(icons.CommunicationStayPrimaryPortrait))
	communication["CommunicationStopScreenShare"] = mustIcon(material.NewIcon(icons.CommunicationStopScreenShare))
	communication["CommunicationSwapCalls"] = mustIcon(material.NewIcon(icons.CommunicationSwapCalls))
	communication["CommunicationTextSMS"] = mustIcon(material.NewIcon(icons.CommunicationTextSMS))
	communication["CommunicationVoicemail"] = mustIcon(material.NewIcon(icons.CommunicationVoicemail))
	communication["CommunicationVPNKey"] = mustIcon(material.NewIcon(icons.CommunicationVPNKey))

	content := make(map[string]*material.Icon)
	content["ContentAdd"] = mustIcon(material.NewIcon(icons.ContentAdd))
	content["ContentAddBox"] = mustIcon(material.NewIcon(icons.ContentAddBox))
	content["ContentAddCircle"] = mustIcon(material.NewIcon(icons.ContentAddCircle))
	content["ContentAddCircleOutline"] = mustIcon(material.NewIcon(icons.ContentAddCircleOutline))
	content["ContentArchive"] = mustIcon(material.NewIcon(icons.ContentArchive))
	content["ContentBackspace"] = mustIcon(material.NewIcon(icons.ContentBackspace))
	content["ContentBlock"] = mustIcon(material.NewIcon(icons.ContentBlock))
	content["ContentClear"] = mustIcon(material.NewIcon(icons.ContentClear))
	content["ContentContentCopy"] = mustIcon(material.NewIcon(icons.ContentContentCopy))
	content["ContentContentCut"] = mustIcon(material.NewIcon(icons.ContentContentCut))
	content["ContentContentPaste"] = mustIcon(material.NewIcon(icons.ContentContentPaste))
	content["ContentCreate"] = mustIcon(material.NewIcon(icons.ContentCreate))
	content["ContentDeleteSweep"] = mustIcon(material.NewIcon(icons.ContentDeleteSweep))
	content["ContentDrafts"] = mustIcon(material.NewIcon(icons.ContentDrafts))
	content["ContentFilterList"] = mustIcon(material.NewIcon(icons.ContentFilterList))
	content["ContentFlag"] = mustIcon(material.NewIcon(icons.ContentFlag))
	content["ContentFontDownload"] = mustIcon(material.NewIcon(icons.ContentFontDownload))
	content["ContentForward"] = mustIcon(material.NewIcon(icons.ContentForward))
	content["ContentGesture"] = mustIcon(material.NewIcon(icons.ContentGesture))
	content["ContentInbox"] = mustIcon(material.NewIcon(icons.ContentInbox))
	content["ContentLink"] = mustIcon(material.NewIcon(icons.ContentLink))
	content["ContentLowPriority"] = mustIcon(material.NewIcon(icons.ContentLowPriority))
	content["ContentMail"] = mustIcon(material.NewIcon(icons.ContentMail))
	content["ContentMarkUnread"] = mustIcon(material.NewIcon(icons.ContentMarkUnread))
	content["ContentMoveToInbox"] = mustIcon(material.NewIcon(icons.ContentMoveToInbox))
	content["ContentNextWeek"] = mustIcon(material.NewIcon(icons.ContentNextWeek))
	content["ContentRedo"] = mustIcon(material.NewIcon(icons.ContentRedo))
	content["ContentRemove"] = mustIcon(material.NewIcon(icons.ContentRemove))
	content["ContentRemoveCircle"] = mustIcon(material.NewIcon(icons.ContentRemoveCircle))
	content["ContentRemoveCircleOutline"] = mustIcon(material.NewIcon(icons.ContentRemoveCircleOutline))
	content["ContentReply"] = mustIcon(material.NewIcon(icons.ContentReply))
	content["ContentReplyAll"] = mustIcon(material.NewIcon(icons.ContentReplyAll))
	content["ContentReport"] = mustIcon(material.NewIcon(icons.ContentReport))
	content["ContentSave"] = mustIcon(material.NewIcon(icons.ContentSave))
	content["ContentSelectAll"] = mustIcon(material.NewIcon(icons.ContentSelectAll))
	content["ContentSend"] = mustIcon(material.NewIcon(icons.ContentSend))
	content["ContentSort"] = mustIcon(material.NewIcon(icons.ContentSort))
	content["ContentTextFormat"] = mustIcon(material.NewIcon(icons.ContentTextFormat))
	content["ContentUnarchive"] = mustIcon(material.NewIcon(icons.ContentUnarchive))
	content["ContentUndo"] = mustIcon(material.NewIcon(icons.ContentUndo))
	content["ContentWeekend"] = mustIcon(material.NewIcon(icons.ContentWeekend))

	device := make(map[string]*material.Icon)
	device["DeviceAccessAlarm"] = mustIcon(material.NewIcon(icons.DeviceAccessAlarm))
	device["DeviceAccessAlarms"] = mustIcon(material.NewIcon(icons.DeviceAccessAlarms))
	device["DeviceAccessTime"] = mustIcon(material.NewIcon(icons.DeviceAccessTime))
	device["DeviceAddAlarm"] = mustIcon(material.NewIcon(icons.DeviceAddAlarm))
	device["DeviceAirplaneModeActive"] = mustIcon(material.NewIcon(icons.DeviceAirplaneModeActive))
	device["DeviceAirplaneModeInactive"] = mustIcon(material.NewIcon(icons.DeviceAirplaneModeInactive))
	device["DeviceBattery20"] = mustIcon(material.NewIcon(icons.DeviceBattery20))
	device["DeviceBattery30"] = mustIcon(material.NewIcon(icons.DeviceBattery30))
	device["DeviceBattery50"] = mustIcon(material.NewIcon(icons.DeviceBattery50))
	device["DeviceBattery60"] = mustIcon(material.NewIcon(icons.DeviceBattery60))
	device["DeviceBattery80"] = mustIcon(material.NewIcon(icons.DeviceBattery80))
	device["DeviceBattery90"] = mustIcon(material.NewIcon(icons.DeviceBattery90))
	device["DeviceBatteryAlert"] = mustIcon(material.NewIcon(icons.DeviceBatteryAlert))
	device["DeviceBatteryCharging20"] = mustIcon(material.NewIcon(icons.DeviceBatteryCharging20))
	device["DeviceBatteryCharging30"] = mustIcon(material.NewIcon(icons.DeviceBatteryCharging30))
	device["DeviceBatteryCharging50"] = mustIcon(material.NewIcon(icons.DeviceBatteryCharging50))
	device["DeviceBatteryCharging60"] = mustIcon(material.NewIcon(icons.DeviceBatteryCharging60))
	device["DeviceBatteryCharging80"] = mustIcon(material.NewIcon(icons.DeviceBatteryCharging80))
	device["DeviceBatteryCharging90"] = mustIcon(material.NewIcon(icons.DeviceBatteryCharging90))
	device["DeviceBatteryChargingFull"] = mustIcon(material.NewIcon(icons.DeviceBatteryChargingFull))
	device["DeviceBatteryFull"] = mustIcon(material.NewIcon(icons.DeviceBatteryFull))
	device["DeviceBatteryStd"] = mustIcon(material.NewIcon(icons.DeviceBatteryStd))
	device["DeviceBatteryUnknown"] = mustIcon(material.NewIcon(icons.DeviceBatteryUnknown))
	device["DeviceBluetooth"] = mustIcon(material.NewIcon(icons.DeviceBluetooth))
	device["DeviceBluetoothConnected"] = mustIcon(material.NewIcon(icons.DeviceBluetoothConnected))
	device["DeviceBluetoothDisabled"] = mustIcon(material.NewIcon(icons.DeviceBluetoothDisabled))
	device["DeviceBluetoothSearching"] = mustIcon(material.NewIcon(icons.DeviceBluetoothSearching))
	device["DeviceBrightnessAuto"] = mustIcon(material.NewIcon(icons.DeviceBrightnessAuto))
	device["DeviceBrightnessHigh"] = mustIcon(material.NewIcon(icons.DeviceBrightnessHigh))
	device["DeviceBrightnessLow"] = mustIcon(material.NewIcon(icons.DeviceBrightnessLow))
	device["DeviceBrightnessMedium"] = mustIcon(material.NewIcon(icons.DeviceBrightnessMedium))
	device["DeviceDataUsage"] = mustIcon(material.NewIcon(icons.DeviceDataUsage))
	device["DeviceDeveloperMode"] = mustIcon(material.NewIcon(icons.DeviceDeveloperMode))
	device["DeviceDevices"] = mustIcon(material.NewIcon(icons.DeviceDevices))
	device["DeviceDVR"] = mustIcon(material.NewIcon(icons.DeviceDVR))
	device["DeviceGPSFixed"] = mustIcon(material.NewIcon(icons.DeviceGPSFixed))
	device["DeviceGPSNotFixed"] = mustIcon(material.NewIcon(icons.DeviceGPSNotFixed))
	device["DeviceGPSOff"] = mustIcon(material.NewIcon(icons.DeviceGPSOff))
	device["DeviceGraphicEq"] = mustIcon(material.NewIcon(icons.DeviceGraphicEq))
	device["DeviceLocationDisabled"] = mustIcon(material.NewIcon(icons.DeviceLocationDisabled))
	device["DeviceLocationSearching"] = mustIcon(material.NewIcon(icons.DeviceLocationSearching))
	device["DeviceNetworkCell"] = mustIcon(material.NewIcon(icons.DeviceNetworkCell))
	device["DeviceNetworkWiFi"] = mustIcon(material.NewIcon(icons.DeviceNetworkWiFi))
	device["DeviceNFC"] = mustIcon(material.NewIcon(icons.DeviceNFC))
	device["DeviceScreenLockLandscape"] = mustIcon(material.NewIcon(icons.DeviceScreenLockLandscape))
	device["DeviceScreenLockPortrait"] = mustIcon(material.NewIcon(icons.DeviceScreenLockPortrait))
	device["DeviceScreenLockRotation"] = mustIcon(material.NewIcon(icons.DeviceScreenLockRotation))
	device["DeviceScreenRotation"] = mustIcon(material.NewIcon(icons.DeviceScreenRotation))
	device["DeviceSDStorage"] = mustIcon(material.NewIcon(icons.DeviceSDStorage))
	device["DeviceSettingsSystemDaydream"] = mustIcon(material.NewIcon(icons.DeviceSettingsSystemDaydream))
	device["DeviceSignalCellular0Bar"] = mustIcon(material.NewIcon(icons.DeviceSignalCellular0Bar))
	device["DeviceSignalCellular1Bar"] = mustIcon(material.NewIcon(icons.DeviceSignalCellular1Bar))
	device["DeviceSignalCellular2Bar"] = mustIcon(material.NewIcon(icons.DeviceSignalCellular2Bar))
	device["DeviceSignalCellular3Bar"] = mustIcon(material.NewIcon(icons.DeviceSignalCellular3Bar))
	device["DeviceSignalCellular4Bar"] = mustIcon(material.NewIcon(icons.DeviceSignalCellular4Bar))
	device["DeviceSignalCellularConnectedNoInternet0Bar"] = mustIcon(material.NewIcon(icons.DeviceSignalCellularConnectedNoInternet0Bar))
	device["DeviceSignalCellularConnectedNoInternet1Bar"] = mustIcon(material.NewIcon(icons.DeviceSignalCellularConnectedNoInternet1Bar))
	device["DeviceSignalCellularConnectedNoInternet2Bar"] = mustIcon(material.NewIcon(icons.DeviceSignalCellularConnectedNoInternet2Bar))
	device["DeviceSignalCellularConnectedNoInternet3Bar"] = mustIcon(material.NewIcon(icons.DeviceSignalCellularConnectedNoInternet3Bar))
	device["DeviceSignalCellularConnectedNoInternet4Bar"] = mustIcon(material.NewIcon(icons.DeviceSignalCellularConnectedNoInternet4Bar))
	device["DeviceSignalCellularNoSIM"] = mustIcon(material.NewIcon(icons.DeviceSignalCellularNoSIM))
	device["DeviceSignalCellularNull"] = mustIcon(material.NewIcon(icons.DeviceSignalCellularNull))
	device["DeviceSignalCellularOff"] = mustIcon(material.NewIcon(icons.DeviceSignalCellularOff))
	device["DeviceSignalWiFi0Bar"] = mustIcon(material.NewIcon(icons.DeviceSignalWiFi0Bar))
	device["DeviceSignalWiFi1Bar"] = mustIcon(material.NewIcon(icons.DeviceSignalWiFi1Bar))
	device["DeviceSignalWiFi1BarLock"] = mustIcon(material.NewIcon(icons.DeviceSignalWiFi1BarLock))
	device["DeviceSignalWiFi2Bar"] = mustIcon(material.NewIcon(icons.DeviceSignalWiFi2Bar))
	device["DeviceSignalWiFi2BarLock"] = mustIcon(material.NewIcon(icons.DeviceSignalWiFi2BarLock))
	device["DeviceSignalWiFi3Bar"] = mustIcon(material.NewIcon(icons.DeviceSignalWiFi3Bar))
	device["DeviceSignalWiFi3BarLock"] = mustIcon(material.NewIcon(icons.DeviceSignalWiFi3BarLock))
	device["DeviceSignalWiFi4Bar"] = mustIcon(material.NewIcon(icons.DeviceSignalWiFi4Bar))
	device["DeviceSignalWiFi4BarLock"] = mustIcon(material.NewIcon(icons.DeviceSignalWiFi4BarLock))
	device["DeviceSignalWiFiOff"] = mustIcon(material.NewIcon(icons.DeviceSignalWiFiOff))
	device["DeviceStorage"] = mustIcon(material.NewIcon(icons.DeviceStorage))
	device["DeviceUSB"] = mustIcon(material.NewIcon(icons.DeviceUSB))
	device["DeviceWallpaper"] = mustIcon(material.NewIcon(icons.DeviceWallpaper))
	device["DeviceWidgets"] = mustIcon(material.NewIcon(icons.DeviceWidgets))
	device["DeviceWiFiLock"] = mustIcon(material.NewIcon(icons.DeviceWiFiLock))
	device["DeviceWiFiTethering"] = mustIcon(material.NewIcon(icons.DeviceWiFiTethering))

	editor := make(map[string]*material.Icon)
	editor["EditorAttachFile"] = mustIcon(material.NewIcon(icons.EditorAttachFile))
	editor["EditorAttachMoney"] = mustIcon(material.NewIcon(icons.EditorAttachMoney))
	editor["EditorBorderAll"] = mustIcon(material.NewIcon(icons.EditorBorderAll))
	editor["EditorBorderBottom"] = mustIcon(material.NewIcon(icons.EditorBorderBottom))
	editor["EditorBorderClear"] = mustIcon(material.NewIcon(icons.EditorBorderClear))
	editor["EditorBorderColor"] = mustIcon(material.NewIcon(icons.EditorBorderColor))
	editor["EditorBorderHorizontal"] = mustIcon(material.NewIcon(icons.EditorBorderHorizontal))
	editor["EditorBorderInner"] = mustIcon(material.NewIcon(icons.EditorBorderInner))
	editor["EditorBorderLeft"] = mustIcon(material.NewIcon(icons.EditorBorderLeft))
	editor["EditorBorderOuter"] = mustIcon(material.NewIcon(icons.EditorBorderOuter))
	editor["EditorBorderRight"] = mustIcon(material.NewIcon(icons.EditorBorderRight))
	editor["EditorBorderStyle"] = mustIcon(material.NewIcon(icons.EditorBorderStyle))
	editor["EditorBorderTop"] = mustIcon(material.NewIcon(icons.EditorBorderTop))
	editor["EditorBorderVertical"] = mustIcon(material.NewIcon(icons.EditorBorderVertical))
	editor["EditorBubbleChart"] = mustIcon(material.NewIcon(icons.EditorBubbleChart))
	editor["EditorDragHandle"] = mustIcon(material.NewIcon(icons.EditorDragHandle))
	editor["EditorFormatAlignCenter"] = mustIcon(material.NewIcon(icons.EditorFormatAlignCenter))
	editor["EditorFormatAlignJustify"] = mustIcon(material.NewIcon(icons.EditorFormatAlignJustify))
	editor["EditorFormatAlignLeft"] = mustIcon(material.NewIcon(icons.EditorFormatAlignLeft))
	editor["EditorFormatAlignRight"] = mustIcon(material.NewIcon(icons.EditorFormatAlignRight))
	editor["EditorFormatBold"] = mustIcon(material.NewIcon(icons.EditorFormatBold))
	editor["EditorFormatClear"] = mustIcon(material.NewIcon(icons.EditorFormatClear))
	editor["EditorFormatColorFill"] = mustIcon(material.NewIcon(icons.EditorFormatColorFill))
	editor["EditorFormatColorReset"] = mustIcon(material.NewIcon(icons.EditorFormatColorReset))
	editor["EditorFormatColorText"] = mustIcon(material.NewIcon(icons.EditorFormatColorText))
	editor["EditorFormatIndentDecrease"] = mustIcon(material.NewIcon(icons.EditorFormatIndentDecrease))
	editor["EditorFormatIndentIncrease"] = mustIcon(material.NewIcon(icons.EditorFormatIndentIncrease))
	editor["EditorFormatItalic"] = mustIcon(material.NewIcon(icons.EditorFormatItalic))
	editor["EditorFormatLineSpacing"] = mustIcon(material.NewIcon(icons.EditorFormatLineSpacing))
	editor["EditorFormatListBulleted"] = mustIcon(material.NewIcon(icons.EditorFormatListBulleted))
	editor["EditorFormatListNumbered"] = mustIcon(material.NewIcon(icons.EditorFormatListNumbered))
	editor["EditorFormatPaint"] = mustIcon(material.NewIcon(icons.EditorFormatPaint))
	editor["EditorFormatQuote"] = mustIcon(material.NewIcon(icons.EditorFormatQuote))
	editor["EditorFormatShapes"] = mustIcon(material.NewIcon(icons.EditorFormatShapes))
	editor["EditorFormatSize"] = mustIcon(material.NewIcon(icons.EditorFormatSize))
	editor["EditorFormatStrikethrough"] = mustIcon(material.NewIcon(icons.EditorFormatStrikethrough))
	editor["EditorFormatTextDirectionLToR"] = mustIcon(material.NewIcon(icons.EditorFormatTextDirectionLToR))
	editor["EditorFormatTextDirectionRToL"] = mustIcon(material.NewIcon(icons.EditorFormatTextDirectionRToL))
	editor["EditorFormatUnderlined"] = mustIcon(material.NewIcon(icons.EditorFormatUnderlined))
	editor["EditorFunctions"] = mustIcon(material.NewIcon(icons.EditorFunctions))
	editor["EditorHighlight"] = mustIcon(material.NewIcon(icons.EditorHighlight))
	editor["EditorInsertChart"] = mustIcon(material.NewIcon(icons.EditorInsertChart))
	editor["EditorInsertComment"] = mustIcon(material.NewIcon(icons.EditorInsertComment))
	editor["EditorInsertDriveFile"] = mustIcon(material.NewIcon(icons.EditorInsertDriveFile))
	editor["EditorInsertEmoticon"] = mustIcon(material.NewIcon(icons.EditorInsertEmoticon))
	editor["EditorInsertInvitation"] = mustIcon(material.NewIcon(icons.EditorInsertInvitation))
	editor["EditorInsertLink"] = mustIcon(material.NewIcon(icons.EditorInsertLink))
	editor["EditorInsertPhoto"] = mustIcon(material.NewIcon(icons.EditorInsertPhoto))
	editor["EditorLinearScale"] = mustIcon(material.NewIcon(icons.EditorLinearScale))
	editor["EditorMergeType"] = mustIcon(material.NewIcon(icons.EditorMergeType))
	editor["EditorModeComment"] = mustIcon(material.NewIcon(icons.EditorModeComment))
	editor["EditorModeEdit"] = mustIcon(material.NewIcon(icons.EditorModeEdit))
	editor["EditorMonetizationOn"] = mustIcon(material.NewIcon(icons.EditorMonetizationOn))
	editor["EditorMoneyOff"] = mustIcon(material.NewIcon(icons.EditorMoneyOff))
	editor["EditorMultilineChart"] = mustIcon(material.NewIcon(icons.EditorMultilineChart))
	editor["EditorPieChart"] = mustIcon(material.NewIcon(icons.EditorPieChart))
	editor["EditorPieChartOutlined"] = mustIcon(material.NewIcon(icons.EditorPieChartOutlined))
	editor["EditorPublish"] = mustIcon(material.NewIcon(icons.EditorPublish))
	editor["EditorShortText"] = mustIcon(material.NewIcon(icons.EditorShortText))
	editor["EditorShowChart"] = mustIcon(material.NewIcon(icons.EditorShowChart))
	editor["EditorSpaceBar"] = mustIcon(material.NewIcon(icons.EditorSpaceBar))
	editor["EditorStrikethroughS"] = mustIcon(material.NewIcon(icons.EditorStrikethroughS))
	editor["EditorTextFields"] = mustIcon(material.NewIcon(icons.EditorTextFields))
	editor["EditorTitle"] = mustIcon(material.NewIcon(icons.EditorTitle))
	editor["EditorVerticalAlignBottom"] = mustIcon(material.NewIcon(icons.EditorVerticalAlignBottom))
	editor["EditorVerticalAlignCenter"] = mustIcon(material.NewIcon(icons.EditorVerticalAlignCenter))
	editor["EditorVerticalAlignTop"] = mustIcon(material.NewIcon(icons.EditorVerticalAlignTop))
	editor["EditorWrapText"] = mustIcon(material.NewIcon(icons.EditorWrapText))

	file := make(map[string]*material.Icon)
	file["FileAttachment"] = mustIcon(material.NewIcon(icons.FileAttachment))
	file["FileCloud"] = mustIcon(material.NewIcon(icons.FileCloud))
	file["FileCloudCircle"] = mustIcon(material.NewIcon(icons.FileCloudCircle))
	file["FileCloudDone"] = mustIcon(material.NewIcon(icons.FileCloudDone))
	file["FileCloudDownload"] = mustIcon(material.NewIcon(icons.FileCloudDownload))
	file["FileCloudOff"] = mustIcon(material.NewIcon(icons.FileCloudOff))
	file["FileCloudQueue"] = mustIcon(material.NewIcon(icons.FileCloudQueue))
	file["FileCloudUpload"] = mustIcon(material.NewIcon(icons.FileCloudUpload))
	file["FileCreateNewFolder"] = mustIcon(material.NewIcon(icons.FileCreateNewFolder))
	file["FileFileDownload"] = mustIcon(material.NewIcon(icons.FileFileDownload))
	file["FileFileUpload"] = mustIcon(material.NewIcon(icons.FileFileUpload))
	file["FileFolder"] = mustIcon(material.NewIcon(icons.FileFolder))
	file["FileFolderOpen"] = mustIcon(material.NewIcon(icons.FileFolderOpen))
	file["FileFolderShared"] = mustIcon(material.NewIcon(icons.FileFolderShared))

	hardware := make(map[string]*material.Icon)
	hardware["HardwareCast"] = mustIcon(material.NewIcon(icons.HardwareCast))
	hardware["HardwareCastConnected"] = mustIcon(material.NewIcon(icons.HardwareCastConnected))
	hardware["HardwareComputer"] = mustIcon(material.NewIcon(icons.HardwareComputer))
	hardware["HardwareDesktopMac"] = mustIcon(material.NewIcon(icons.HardwareDesktopMac))
	hardware["HardwareDesktopWindows"] = mustIcon(material.NewIcon(icons.HardwareDesktopWindows))
	hardware["HardwareDeveloperBoard"] = mustIcon(material.NewIcon(icons.HardwareDeveloperBoard))
	hardware["HardwareDeviceHub"] = mustIcon(material.NewIcon(icons.HardwareDeviceHub))
	hardware["HardwareDevicesOther"] = mustIcon(material.NewIcon(icons.HardwareDevicesOther))
	hardware["HardwareDock"] = mustIcon(material.NewIcon(icons.HardwareDock))
	hardware["HardwareGamepad"] = mustIcon(material.NewIcon(icons.HardwareGamepad))
	hardware["HardwareHeadset"] = mustIcon(material.NewIcon(icons.HardwareHeadset))
	hardware["HardwareHeadsetMic"] = mustIcon(material.NewIcon(icons.HardwareHeadsetMic))
	hardware["HardwareKeyboard"] = mustIcon(material.NewIcon(icons.HardwareKeyboard))
	hardware["HardwareKeyboardArrowDown"] = mustIcon(material.NewIcon(icons.HardwareKeyboardArrowDown))
	hardware["HardwareKeyboardArrowLeft"] = mustIcon(material.NewIcon(icons.HardwareKeyboardArrowLeft))
	hardware["HardwareKeyboardArrowRight"] = mustIcon(material.NewIcon(icons.HardwareKeyboardArrowRight))
	hardware["HardwareKeyboardArrowUp"] = mustIcon(material.NewIcon(icons.HardwareKeyboardArrowUp))
	hardware["HardwareKeyboardBackspace"] = mustIcon(material.NewIcon(icons.HardwareKeyboardBackspace))
	hardware["HardwareKeyboardCapslock"] = mustIcon(material.NewIcon(icons.HardwareKeyboardCapslock))
	hardware["HardwareKeyboardHide"] = mustIcon(material.NewIcon(icons.HardwareKeyboardHide))
	hardware["HardwareKeyboardReturn"] = mustIcon(material.NewIcon(icons.HardwareKeyboardReturn))
	hardware["HardwareKeyboardTab"] = mustIcon(material.NewIcon(icons.HardwareKeyboardTab))
	hardware["HardwareKeyboardVoice"] = mustIcon(material.NewIcon(icons.HardwareKeyboardVoice))
	hardware["HardwareLaptop"] = mustIcon(material.NewIcon(icons.HardwareLaptop))
	hardware["HardwareLaptopChromebook"] = mustIcon(material.NewIcon(icons.HardwareLaptopChromebook))
	hardware["HardwareLaptopMac"] = mustIcon(material.NewIcon(icons.HardwareLaptopMac))
	hardware["HardwareLaptopWindows"] = mustIcon(material.NewIcon(icons.HardwareLaptopWindows))
	hardware["HardwareMemory"] = mustIcon(material.NewIcon(icons.HardwareMemory))
	hardware["HardwareMouse"] = mustIcon(material.NewIcon(icons.HardwareMouse))
	hardware["HardwarePhoneAndroid"] = mustIcon(material.NewIcon(icons.HardwarePhoneAndroid))
	hardware["HardwarePhoneIPhone"] = mustIcon(material.NewIcon(icons.HardwarePhoneIPhone))
	hardware["HardwarePhoneLink"] = mustIcon(material.NewIcon(icons.HardwarePhoneLink))
	hardware["HardwarePhoneLinkOff"] = mustIcon(material.NewIcon(icons.HardwarePhoneLinkOff))
	hardware["HardwarePowerInput"] = mustIcon(material.NewIcon(icons.HardwarePowerInput))
	hardware["HardwareRouter"] = mustIcon(material.NewIcon(icons.HardwareRouter))
	hardware["HardwareScanner"] = mustIcon(material.NewIcon(icons.HardwareScanner))
	hardware["HardwareSecurity"] = mustIcon(material.NewIcon(icons.HardwareSecurity))
	hardware["HardwareSIMCard"] = mustIcon(material.NewIcon(icons.HardwareSIMCard))
	hardware["HardwareSmartphone"] = mustIcon(material.NewIcon(icons.HardwareSmartphone))
	hardware["HardwareSpeaker"] = mustIcon(material.NewIcon(icons.HardwareSpeaker))
	hardware["HardwareSpeakerGroup"] = mustIcon(material.NewIcon(icons.HardwareSpeakerGroup))
	hardware["HardwareTablet"] = mustIcon(material.NewIcon(icons.HardwareTablet))
	hardware["HardwareTabletAndroid"] = mustIcon(material.NewIcon(icons.HardwareTabletAndroid))
	hardware["HardwareTabletMac"] = mustIcon(material.NewIcon(icons.HardwareTabletMac))
	hardware["HardwareToys"] = mustIcon(material.NewIcon(icons.HardwareToys))
	hardware["HardwareTV"] = mustIcon(material.NewIcon(icons.HardwareTV))
	hardware["HardwareVideogameAsset"] = mustIcon(material.NewIcon(icons.HardwareVideogameAsset))
	hardware["HardwareWatch"] = mustIcon(material.NewIcon(icons.HardwareWatch))

	image := make(map[string]*material.Icon)
	image["ImageAddAPhoto"] = mustIcon(material.NewIcon(icons.ImageAddAPhoto))
	image["ImageAddToPhotos"] = mustIcon(material.NewIcon(icons.ImageAddToPhotos))
	image["ImageAdjust"] = mustIcon(material.NewIcon(icons.ImageAdjust))
	image["ImageAssistant"] = mustIcon(material.NewIcon(icons.ImageAssistant))
	image["ImageAssistantPhoto"] = mustIcon(material.NewIcon(icons.ImageAssistantPhoto))
	image["ImageAudiotrack"] = mustIcon(material.NewIcon(icons.ImageAudiotrack))
	image["ImageBlurCircular"] = mustIcon(material.NewIcon(icons.ImageBlurCircular))
	image["ImageBlurLinear"] = mustIcon(material.NewIcon(icons.ImageBlurLinear))
	image["ImageBlurOff"] = mustIcon(material.NewIcon(icons.ImageBlurOff))
	image["ImageBlurOn"] = mustIcon(material.NewIcon(icons.ImageBlurOn))
	image["ImageBrightness1"] = mustIcon(material.NewIcon(icons.ImageBrightness1))
	image["ImageBrightness2"] = mustIcon(material.NewIcon(icons.ImageBrightness2))
	image["ImageBrightness3"] = mustIcon(material.NewIcon(icons.ImageBrightness3))
	image["ImageBrightness4"] = mustIcon(material.NewIcon(icons.ImageBrightness4))
	image["ImageBrightness5"] = mustIcon(material.NewIcon(icons.ImageBrightness5))
	image["ImageBrightness6"] = mustIcon(material.NewIcon(icons.ImageBrightness6))
	image["ImageBrightness7"] = mustIcon(material.NewIcon(icons.ImageBrightness7))
	image["ImageBrokenImage"] = mustIcon(material.NewIcon(icons.ImageBrokenImage))
	image["ImageBrush"] = mustIcon(material.NewIcon(icons.ImageBrush))
	image["ImageBurstMode"] = mustIcon(material.NewIcon(icons.ImageBurstMode))
	image["ImageCamera"] = mustIcon(material.NewIcon(icons.ImageCamera))
	image["ImageCameraAlt"] = mustIcon(material.NewIcon(icons.ImageCameraAlt))
	image["ImageCameraFront"] = mustIcon(material.NewIcon(icons.ImageCameraFront))
	image["ImageCameraRear"] = mustIcon(material.NewIcon(icons.ImageCameraRear))
	image["ImageCameraRoll"] = mustIcon(material.NewIcon(icons.ImageCameraRoll))
	image["ImageCenterFocusStrong"] = mustIcon(material.NewIcon(icons.ImageCenterFocusStrong))
	image["ImageCenterFocusWeak"] = mustIcon(material.NewIcon(icons.ImageCenterFocusWeak))
	image["ImageCollections"] = mustIcon(material.NewIcon(icons.ImageCollections))
	image["ImageCollectionsBookmark"] = mustIcon(material.NewIcon(icons.ImageCollectionsBookmark))
	image["ImageColorLens"] = mustIcon(material.NewIcon(icons.ImageColorLens))
	image["ImageColorize"] = mustIcon(material.NewIcon(icons.ImageColorize))
	image["ImageCompare"] = mustIcon(material.NewIcon(icons.ImageCompare))
	image["ImageControlPoint"] = mustIcon(material.NewIcon(icons.ImageControlPoint))
	image["ImageControlPointDuplicate"] = mustIcon(material.NewIcon(icons.ImageControlPointDuplicate))
	image["ImageCrop"] = mustIcon(material.NewIcon(icons.ImageCrop))
	image["ImageCrop169"] = mustIcon(material.NewIcon(icons.ImageCrop169))
	image["ImageCrop32"] = mustIcon(material.NewIcon(icons.ImageCrop32))
	image["ImageCrop54"] = mustIcon(material.NewIcon(icons.ImageCrop54))
	image["ImageCrop75"] = mustIcon(material.NewIcon(icons.ImageCrop75))
	image["ImageCropDIN"] = mustIcon(material.NewIcon(icons.ImageCropDIN))
	image["ImageCropFree"] = mustIcon(material.NewIcon(icons.ImageCropFree))
	image["ImageCropLandscape"] = mustIcon(material.NewIcon(icons.ImageCropLandscape))
	image["ImageCropOriginal"] = mustIcon(material.NewIcon(icons.ImageCropOriginal))
	image["ImageCropPortrait"] = mustIcon(material.NewIcon(icons.ImageCropPortrait))
	image["ImageCropRotate"] = mustIcon(material.NewIcon(icons.ImageCropRotate))
	image["ImageCropSquare"] = mustIcon(material.NewIcon(icons.ImageCropSquare))
	image["ImageDehaze"] = mustIcon(material.NewIcon(icons.ImageDehaze))
	image["ImageDetails"] = mustIcon(material.NewIcon(icons.ImageDetails))
	image["ImageEdit"] = mustIcon(material.NewIcon(icons.ImageEdit))
	image["ImageExposure"] = mustIcon(material.NewIcon(icons.ImageExposure))
	image["ImageExposureNeg1"] = mustIcon(material.NewIcon(icons.ImageExposureNeg1))
	image["ImageExposureNeg2"] = mustIcon(material.NewIcon(icons.ImageExposureNeg2))
	image["ImageExposurePlus1"] = mustIcon(material.NewIcon(icons.ImageExposurePlus1))
	image["ImageExposurePlus2"] = mustIcon(material.NewIcon(icons.ImageExposurePlus2))
	image["ImageExposureZero"] = mustIcon(material.NewIcon(icons.ImageExposureZero))
	image["ImageFilter"] = mustIcon(material.NewIcon(icons.ImageFilter))
	image["ImageFilter1"] = mustIcon(material.NewIcon(icons.ImageFilter1))
	image["ImageFilter2"] = mustIcon(material.NewIcon(icons.ImageFilter2))
	image["ImageFilter3"] = mustIcon(material.NewIcon(icons.ImageFilter3))
	image["ImageFilter4"] = mustIcon(material.NewIcon(icons.ImageFilter4))
	image["ImageFilter5"] = mustIcon(material.NewIcon(icons.ImageFilter5))
	image["ImageFilter6"] = mustIcon(material.NewIcon(icons.ImageFilter6))
	image["ImageFilter7"] = mustIcon(material.NewIcon(icons.ImageFilter7))
	image["ImageFilter8"] = mustIcon(material.NewIcon(icons.ImageFilter8))
	image["ImageFilter9"] = mustIcon(material.NewIcon(icons.ImageFilter9))
	image["ImageFilter9Plus"] = mustIcon(material.NewIcon(icons.ImageFilter9Plus))
	image["ImageFilterBAndW"] = mustIcon(material.NewIcon(icons.ImageFilterBAndW))
	image["ImageFilterCenterFocus"] = mustIcon(material.NewIcon(icons.ImageFilterCenterFocus))
	image["ImageFilterDrama"] = mustIcon(material.NewIcon(icons.ImageFilterDrama))
	image["ImageFilterFrames"] = mustIcon(material.NewIcon(icons.ImageFilterFrames))
	image["ImageFilterHDR"] = mustIcon(material.NewIcon(icons.ImageFilterHDR))
	image["ImageFilterNone"] = mustIcon(material.NewIcon(icons.ImageFilterNone))
	image["ImageFilterTiltShift"] = mustIcon(material.NewIcon(icons.ImageFilterTiltShift))
	image["ImageFilterVintage"] = mustIcon(material.NewIcon(icons.ImageFilterVintage))
	image["ImageFlare"] = mustIcon(material.NewIcon(icons.ImageFlare))
	image["ImageFlashAuto"] = mustIcon(material.NewIcon(icons.ImageFlashAuto))
	image["ImageFlashOff"] = mustIcon(material.NewIcon(icons.ImageFlashOff))
	image["ImageFlashOn"] = mustIcon(material.NewIcon(icons.ImageFlashOn))
	image["ImageFlip"] = mustIcon(material.NewIcon(icons.ImageFlip))
	image["ImageGradient"] = mustIcon(material.NewIcon(icons.ImageGradient))
	image["ImageGrain"] = mustIcon(material.NewIcon(icons.ImageGrain))
	image["ImageGridOff"] = mustIcon(material.NewIcon(icons.ImageGridOff))
	image["ImageGridOn"] = mustIcon(material.NewIcon(icons.ImageGridOn))
	image["ImageHDROff"] = mustIcon(material.NewIcon(icons.ImageHDROff))
	image["ImageHDROn"] = mustIcon(material.NewIcon(icons.ImageHDROn))
	image["ImageHDRStrong"] = mustIcon(material.NewIcon(icons.ImageHDRStrong))
	image["ImageHDRWeak"] = mustIcon(material.NewIcon(icons.ImageHDRWeak))
	image["ImageHealing"] = mustIcon(material.NewIcon(icons.ImageHealing))
	image["ImageImage"] = mustIcon(material.NewIcon(icons.ImageImage))
	image["ImageImageAspectRatio"] = mustIcon(material.NewIcon(icons.ImageImageAspectRatio))
	image["ImageISO"] = mustIcon(material.NewIcon(icons.ImageISO))
	image["ImageLandscape"] = mustIcon(material.NewIcon(icons.ImageLandscape))
	image["ImageLeakAdd"] = mustIcon(material.NewIcon(icons.ImageLeakAdd))
	image["ImageLeakRemove"] = mustIcon(material.NewIcon(icons.ImageLeakRemove))
	image["ImageLens"] = mustIcon(material.NewIcon(icons.ImageLens))
	image["ImageLinkedCamera"] = mustIcon(material.NewIcon(icons.ImageLinkedCamera))
	image["ImageLooks"] = mustIcon(material.NewIcon(icons.ImageLooks))
	image["ImageLooks3"] = mustIcon(material.NewIcon(icons.ImageLooks3))
	image["ImageLooks4"] = mustIcon(material.NewIcon(icons.ImageLooks4))
	image["ImageLooks5"] = mustIcon(material.NewIcon(icons.ImageLooks5))
	image["ImageLooks6"] = mustIcon(material.NewIcon(icons.ImageLooks6))
	image["ImageLooksOne"] = mustIcon(material.NewIcon(icons.ImageLooksOne))
	image["ImageLooksTwo"] = mustIcon(material.NewIcon(icons.ImageLooksTwo))
	image["ImageLoupe"] = mustIcon(material.NewIcon(icons.ImageLoupe))
	image["ImageMonochromePhotos"] = mustIcon(material.NewIcon(icons.ImageMonochromePhotos))
	image["ImageMovieCreation"] = mustIcon(material.NewIcon(icons.ImageMovieCreation))
	image["ImageMovieFilter"] = mustIcon(material.NewIcon(icons.ImageMovieFilter))
	image["ImageMusicNote"] = mustIcon(material.NewIcon(icons.ImageMusicNote))
	image["ImageNature"] = mustIcon(material.NewIcon(icons.ImageNature))
	image["ImageNaturePeople"] = mustIcon(material.NewIcon(icons.ImageNaturePeople))
	image["ImageNavigateBefore"] = mustIcon(material.NewIcon(icons.ImageNavigateBefore))
	image["ImageNavigateNext"] = mustIcon(material.NewIcon(icons.ImageNavigateNext))
	image["ImagePalette"] = mustIcon(material.NewIcon(icons.ImagePalette))
	image["ImagePanorama"] = mustIcon(material.NewIcon(icons.ImagePanorama))
	image["ImagePanoramaFishEye"] = mustIcon(material.NewIcon(icons.ImagePanoramaFishEye))
	image["ImagePanoramaHorizontal"] = mustIcon(material.NewIcon(icons.ImagePanoramaHorizontal))
	image["ImagePanoramaVertical"] = mustIcon(material.NewIcon(icons.ImagePanoramaVertical))
	image["ImagePanoramaWideAngle"] = mustIcon(material.NewIcon(icons.ImagePanoramaWideAngle))
	image["ImagePhoto"] = mustIcon(material.NewIcon(icons.ImagePhoto))
	image["ImagePhotoAlbum"] = mustIcon(material.NewIcon(icons.ImagePhotoAlbum))
	image["ImagePhotoCamera"] = mustIcon(material.NewIcon(icons.ImagePhotoCamera))
	image["ImagePhotoFilter"] = mustIcon(material.NewIcon(icons.ImagePhotoFilter))
	image["ImagePhotoLibrary"] = mustIcon(material.NewIcon(icons.ImagePhotoLibrary))
	image["ImagePhotoSizeSelectActual"] = mustIcon(material.NewIcon(icons.ImagePhotoSizeSelectActual))
	image["ImagePhotoSizeSelectLarge"] = mustIcon(material.NewIcon(icons.ImagePhotoSizeSelectLarge))
	image["ImagePhotoSizeSelectSmall"] = mustIcon(material.NewIcon(icons.ImagePhotoSizeSelectSmall))
	image["ImagePictureAsPDF"] = mustIcon(material.NewIcon(icons.ImagePictureAsPDF))
	image["ImagePortrait"] = mustIcon(material.NewIcon(icons.ImagePortrait))
	image["ImageRemoveRedEye"] = mustIcon(material.NewIcon(icons.ImageRemoveRedEye))
	image["ImageRotate90DegreesCCW"] = mustIcon(material.NewIcon(icons.ImageRotate90DegreesCCW))
	image["ImageRotateLeft"] = mustIcon(material.NewIcon(icons.ImageRotateLeft))
	image["ImageRotateRight"] = mustIcon(material.NewIcon(icons.ImageRotateRight))
	image["ImageSlideshow"] = mustIcon(material.NewIcon(icons.ImageSlideshow))
	image["ImageStraighten"] = mustIcon(material.NewIcon(icons.ImageStraighten))
	image["ImageStyle"] = mustIcon(material.NewIcon(icons.ImageStyle))
	image["ImageSwitchCamera"] = mustIcon(material.NewIcon(icons.ImageSwitchCamera))
	image["ImageSwitchVideo"] = mustIcon(material.NewIcon(icons.ImageSwitchVideo))
	image["ImageTagFaces"] = mustIcon(material.NewIcon(icons.ImageTagFaces))
	image["ImageTexture"] = mustIcon(material.NewIcon(icons.ImageTexture))
	image["ImageTimeLapse"] = mustIcon(material.NewIcon(icons.ImageTimeLapse))
	image["ImageTimer"] = mustIcon(material.NewIcon(icons.ImageTimer))
	image["ImageTimer10"] = mustIcon(material.NewIcon(icons.ImageTimer10))
	image["ImageTimer3"] = mustIcon(material.NewIcon(icons.ImageTimer3))
	image["ImageTimerOff"] = mustIcon(material.NewIcon(icons.ImageTimerOff))
	image["ImageTonality"] = mustIcon(material.NewIcon(icons.ImageTonality))
	image["ImageTransform"] = mustIcon(material.NewIcon(icons.ImageTransform))
	image["ImageTune"] = mustIcon(material.NewIcon(icons.ImageTune))
	image["ImageViewComfy"] = mustIcon(material.NewIcon(icons.ImageViewComfy))
	image["ImageViewCompact"] = mustIcon(material.NewIcon(icons.ImageViewCompact))
	image["ImageVignette"] = mustIcon(material.NewIcon(icons.ImageVignette))
	image["ImageWBAuto"] = mustIcon(material.NewIcon(icons.ImageWBAuto))
	image["ImageWBCloudy"] = mustIcon(material.NewIcon(icons.ImageWBCloudy))
	image["ImageWBIncandescent"] = mustIcon(material.NewIcon(icons.ImageWBIncandescent))
	image["ImageWBIridescent"] = mustIcon(material.NewIcon(icons.ImageWBIridescent))
	image["ImageWBSunny"] = mustIcon(material.NewIcon(icons.ImageWBSunny))

	maps := make(map[string]*material.Icon)
	maps["MapsAddLocation"] = mustIcon(material.NewIcon(icons.MapsAddLocation))
	maps["MapsBeenhere"] = mustIcon(material.NewIcon(icons.MapsBeenhere))
	maps["MapsDirections"] = mustIcon(material.NewIcon(icons.MapsDirections))
	maps["MapsDirectionsBike"] = mustIcon(material.NewIcon(icons.MapsDirectionsBike))
	maps["MapsDirectionsBoat"] = mustIcon(material.NewIcon(icons.MapsDirectionsBoat))
	maps["MapsDirectionsBus"] = mustIcon(material.NewIcon(icons.MapsDirectionsBus))
	maps["MapsDirectionsCar"] = mustIcon(material.NewIcon(icons.MapsDirectionsCar))
	maps["MapsDirectionsRailway"] = mustIcon(material.NewIcon(icons.MapsDirectionsRailway))
	maps["MapsDirectionsRun"] = mustIcon(material.NewIcon(icons.MapsDirectionsRun))
	maps["MapsDirectionsSubway"] = mustIcon(material.NewIcon(icons.MapsDirectionsSubway))
	maps["MapsDirectionsTransit"] = mustIcon(material.NewIcon(icons.MapsDirectionsTransit))
	maps["MapsDirectionsWalk"] = mustIcon(material.NewIcon(icons.MapsDirectionsWalk))
	maps["MapsEditLocation"] = mustIcon(material.NewIcon(icons.MapsEditLocation))
	maps["MapsEVStation"] = mustIcon(material.NewIcon(icons.MapsEVStation))
	maps["MapsFlight"] = mustIcon(material.NewIcon(icons.MapsFlight))
	maps["MapsHotel"] = mustIcon(material.NewIcon(icons.MapsHotel))
	maps["MapsLayers"] = mustIcon(material.NewIcon(icons.MapsLayers))
	maps["MapsLayersClear"] = mustIcon(material.NewIcon(icons.MapsLayersClear))
	maps["MapsLocalActivity"] = mustIcon(material.NewIcon(icons.MapsLocalActivity))
	maps["MapsLocalAirport"] = mustIcon(material.NewIcon(icons.MapsLocalAirport))
	maps["MapsLocalATM"] = mustIcon(material.NewIcon(icons.MapsLocalATM))
	maps["MapsLocalBar"] = mustIcon(material.NewIcon(icons.MapsLocalBar))
	maps["MapsLocalCafe"] = mustIcon(material.NewIcon(icons.MapsLocalCafe))
	maps["MapsLocalCarWash"] = mustIcon(material.NewIcon(icons.MapsLocalCarWash))
	maps["MapsLocalConvenienceStore"] = mustIcon(material.NewIcon(icons.MapsLocalConvenienceStore))
	maps["MapsLocalDining"] = mustIcon(material.NewIcon(icons.MapsLocalDining))
	maps["MapsLocalDrink"] = mustIcon(material.NewIcon(icons.MapsLocalDrink))
	maps["MapsLocalFlorist"] = mustIcon(material.NewIcon(icons.MapsLocalFlorist))
	maps["MapsLocalGasStation"] = mustIcon(material.NewIcon(icons.MapsLocalGasStation))
	maps["MapsLocalGroceryStore"] = mustIcon(material.NewIcon(icons.MapsLocalGroceryStore))
	maps["MapsLocalHospital"] = mustIcon(material.NewIcon(icons.MapsLocalHospital))
	maps["MapsLocalHotel"] = mustIcon(material.NewIcon(icons.MapsLocalHotel))
	maps["MapsLocalLaundryService"] = mustIcon(material.NewIcon(icons.MapsLocalLaundryService))
	maps["MapsLocalLibrary"] = mustIcon(material.NewIcon(icons.MapsLocalLibrary))
	maps["MapsLocalMall"] = mustIcon(material.NewIcon(icons.MapsLocalMall))
	maps["MapsLocalMovies"] = mustIcon(material.NewIcon(icons.MapsLocalMovies))
	maps["MapsLocalOffer"] = mustIcon(material.NewIcon(icons.MapsLocalOffer))
	maps["MapsLocalParking"] = mustIcon(material.NewIcon(icons.MapsLocalParking))
	maps["MapsLocalPharmacy"] = mustIcon(material.NewIcon(icons.MapsLocalPharmacy))
	maps["MapsLocalPhone"] = mustIcon(material.NewIcon(icons.MapsLocalPhone))
	maps["MapsLocalPizza"] = mustIcon(material.NewIcon(icons.MapsLocalPizza))
	maps["MapsLocalPlay"] = mustIcon(material.NewIcon(icons.MapsLocalPlay))
	maps["MapsLocalPostOffice"] = mustIcon(material.NewIcon(icons.MapsLocalPostOffice))
	maps["MapsLocalPrintshop"] = mustIcon(material.NewIcon(icons.MapsLocalPrintshop))
	maps["MapsLocalSee"] = mustIcon(material.NewIcon(icons.MapsLocalSee))
	maps["MapsLocalShipping"] = mustIcon(material.NewIcon(icons.MapsLocalShipping))
	maps["MapsLocalTaxi"] = mustIcon(material.NewIcon(icons.MapsLocalTaxi))
	maps["MapsMap"] = mustIcon(material.NewIcon(icons.MapsMap))
	maps["MapsMyLocation"] = mustIcon(material.NewIcon(icons.MapsMyLocation))
	maps["MapsNavigation"] = mustIcon(material.NewIcon(icons.MapsNavigation))
	maps["MapsNearMe"] = mustIcon(material.NewIcon(icons.MapsNearMe))
	maps["MapsPersonPin"] = mustIcon(material.NewIcon(icons.MapsPersonPin))
	maps["MapsPersonPinCircle"] = mustIcon(material.NewIcon(icons.MapsPersonPinCircle))
	maps["MapsPinDrop"] = mustIcon(material.NewIcon(icons.MapsPinDrop))
	maps["MapsPlace"] = mustIcon(material.NewIcon(icons.MapsPlace))
	maps["MapsRateReview"] = mustIcon(material.NewIcon(icons.MapsRateReview))
	maps["MapsRestaurant"] = mustIcon(material.NewIcon(icons.MapsRestaurant))
	maps["MapsRestaurantMenu"] = mustIcon(material.NewIcon(icons.MapsRestaurantMenu))
	maps["MapsSatellite"] = mustIcon(material.NewIcon(icons.MapsSatellite))
	maps["MapsStoreMallDirectory"] = mustIcon(material.NewIcon(icons.MapsStoreMallDirectory))
	maps["MapsStreetView"] = mustIcon(material.NewIcon(icons.MapsStreetView))
	maps["MapsSubway"] = mustIcon(material.NewIcon(icons.MapsSubway))
	maps["MapsTerrain"] = mustIcon(material.NewIcon(icons.MapsTerrain))
	maps["MapsTraffic"] = mustIcon(material.NewIcon(icons.MapsTraffic))
	maps["MapsTrain"] = mustIcon(material.NewIcon(icons.MapsTrain))
	maps["MapsTram"] = mustIcon(material.NewIcon(icons.MapsTram))
	maps["MapsTransferWithinAStation"] = mustIcon(material.NewIcon(icons.MapsTransferWithinAStation))
	maps["MapsZoomOutMap"] = mustIcon(material.NewIcon(icons.MapsZoomOutMap))

	nav := make(map[string]*material.Icon)
	nav["NavigationApps"] = mustIcon(material.NewIcon(icons.NavigationApps))
	nav["NavigationArrowBack"] = mustIcon(material.NewIcon(icons.NavigationArrowBack))
	nav["NavigationArrowDownward"] = mustIcon(material.NewIcon(icons.NavigationArrowDownward))
	nav["NavigationArrowDropDown"] = mustIcon(material.NewIcon(icons.NavigationArrowDropDown))
	nav["NavigationArrowDropDownCircle"] = mustIcon(material.NewIcon(icons.NavigationArrowDropDownCircle))
	nav["NavigationArrowDropUp"] = mustIcon(material.NewIcon(icons.NavigationArrowDropUp))
	nav["NavigationArrowForward"] = mustIcon(material.NewIcon(icons.NavigationArrowForward))
	nav["NavigationArrowUpward"] = mustIcon(material.NewIcon(icons.NavigationArrowUpward))
	nav["NavigationCancel"] = mustIcon(material.NewIcon(icons.NavigationCancel))
	nav["NavigationCheck"] = mustIcon(material.NewIcon(icons.NavigationCheck))
	nav["NavigationChevronLeft"] = mustIcon(material.NewIcon(icons.NavigationChevronLeft))
	nav["NavigationChevronRight"] = mustIcon(material.NewIcon(icons.NavigationChevronRight))
	nav["NavigationClose"] = mustIcon(material.NewIcon(icons.NavigationClose))
	nav["NavigationExpandLess"] = mustIcon(material.NewIcon(icons.NavigationExpandLess))
	nav["NavigationExpandMore"] = mustIcon(material.NewIcon(icons.NavigationExpandMore))
	nav["NavigationFirstPage"] = mustIcon(material.NewIcon(icons.NavigationFirstPage))
	nav["NavigationFullscreen"] = mustIcon(material.NewIcon(icons.NavigationFullscreen))
	nav["NavigationFullscreenExit"] = mustIcon(material.NewIcon(icons.NavigationFullscreenExit))
	nav["NavigationLastPage"] = mustIcon(material.NewIcon(icons.NavigationLastPage))
	nav["NavigationMenu"] = mustIcon(material.NewIcon(icons.NavigationMenu))
	nav["NavigationMoreHoriz"] = mustIcon(material.NewIcon(icons.NavigationMoreHoriz))
	nav["NavigationMoreVert"] = mustIcon(material.NewIcon(icons.NavigationMoreVert))
	nav["NavigationRefresh"] = mustIcon(material.NewIcon(icons.NavigationRefresh))
	nav["NavigationSubdirectoryArrowLeft"] = mustIcon(material.NewIcon(icons.NavigationSubdirectoryArrowLeft))
	nav["NavigationSubdirectoryArrowRight"] = mustIcon(material.NewIcon(icons.NavigationSubdirectoryArrowRight))
	nav["NavigationUnfoldLess"] = mustIcon(material.NewIcon(icons.NavigationUnfoldLess))
	nav["NavigationUnfoldMore"] = mustIcon(material.NewIcon(icons.NavigationUnfoldMore))

	notification := make(map[string]*material.Icon)
	notification["NotificationADB"] = mustIcon(material.NewIcon(icons.NotificationADB))
	notification["NotificationAirlineSeatFlat"] = mustIcon(material.NewIcon(icons.NotificationAirlineSeatFlat))
	notification["NotificationAirlineSeatFlatAngled"] = mustIcon(material.NewIcon(icons.NotificationAirlineSeatFlatAngled))
	notification["NotificationAirlineSeatIndividualSuite"] = mustIcon(material.NewIcon(icons.NotificationAirlineSeatIndividualSuite))
	notification["NotificationAirlineSeatLegroomExtra"] = mustIcon(material.NewIcon(icons.NotificationAirlineSeatLegroomExtra))
	notification["NotificationAirlineSeatLegroomNormal"] = mustIcon(material.NewIcon(icons.NotificationAirlineSeatLegroomNormal))
	notification["NotificationAirlineSeatLegroomReduced"] = mustIcon(material.NewIcon(icons.NotificationAirlineSeatLegroomReduced))
	notification["NotificationAirlineSeatReclineExtra"] = mustIcon(material.NewIcon(icons.NotificationAirlineSeatReclineExtra))
	notification["NotificationAirlineSeatReclineNormal"] = mustIcon(material.NewIcon(icons.NotificationAirlineSeatReclineNormal))
	notification["NotificationBluetoothAudio"] = mustIcon(material.NewIcon(icons.NotificationBluetoothAudio))
	notification["NotificationConfirmationNumber"] = mustIcon(material.NewIcon(icons.NotificationConfirmationNumber))
	notification["NotificationDiscFull"] = mustIcon(material.NewIcon(icons.NotificationDiscFull))
	notification["NotificationDoNotDisturb"] = mustIcon(material.NewIcon(icons.NotificationDoNotDisturb))
	notification["NotificationDoNotDisturbAlt"] = mustIcon(material.NewIcon(icons.NotificationDoNotDisturbAlt))
	notification["NotificationDoNotDisturbOff"] = mustIcon(material.NewIcon(icons.NotificationDoNotDisturbOff))
	notification["NotificationDoNotDisturbOn"] = mustIcon(material.NewIcon(icons.NotificationDoNotDisturbOn))
	notification["NotificationDriveETA"] = mustIcon(material.NewIcon(icons.NotificationDriveETA))
	notification["NotificationEnhancedEncryption"] = mustIcon(material.NewIcon(icons.NotificationEnhancedEncryption))
	notification["NotificationEventAvailable"] = mustIcon(material.NewIcon(icons.NotificationEventAvailable))
	notification["NotificationEventBusy"] = mustIcon(material.NewIcon(icons.NotificationEventBusy))
	notification["NotificationEventNote"] = mustIcon(material.NewIcon(icons.NotificationEventNote))
	notification["NotificationFolderSpecial"] = mustIcon(material.NewIcon(icons.NotificationFolderSpecial))
	notification["NotificationLiveTV"] = mustIcon(material.NewIcon(icons.NotificationLiveTV))
	notification["NotificationMMS"] = mustIcon(material.NewIcon(icons.NotificationMMS))
	notification["NotificationMore"] = mustIcon(material.NewIcon(icons.NotificationMore))
	notification["NotificationNetworkCheck"] = mustIcon(material.NewIcon(icons.NotificationNetworkCheck))
	notification["NotificationNetworkLocked"] = mustIcon(material.NewIcon(icons.NotificationNetworkLocked))
	notification["NotificationNoEncryption"] = mustIcon(material.NewIcon(icons.NotificationNoEncryption))
	notification["NotificationOnDemandVideo"] = mustIcon(material.NewIcon(icons.NotificationOnDemandVideo))
	notification["NotificationPersonalVideo"] = mustIcon(material.NewIcon(icons.NotificationPersonalVideo))
	notification["NotificationPhoneBluetoothSpeaker"] = mustIcon(material.NewIcon(icons.NotificationPhoneBluetoothSpeaker))
	notification["NotificationPhoneForwarded"] = mustIcon(material.NewIcon(icons.NotificationPhoneForwarded))
	notification["NotificationPhoneInTalk"] = mustIcon(material.NewIcon(icons.NotificationPhoneInTalk))
	notification["NotificationPhoneLocked"] = mustIcon(material.NewIcon(icons.NotificationPhoneLocked))
	notification["NotificationPhoneMissed"] = mustIcon(material.NewIcon(icons.NotificationPhoneMissed))
	notification["NotificationPhonePaused"] = mustIcon(material.NewIcon(icons.NotificationPhonePaused))
	notification["NotificationPower"] = mustIcon(material.NewIcon(icons.NotificationPower))
	notification["NotificationPriorityHigh"] = mustIcon(material.NewIcon(icons.NotificationPriorityHigh))
	notification["NotificationRVHookup"] = mustIcon(material.NewIcon(icons.NotificationRVHookup))
	notification["NotificationSDCard"] = mustIcon(material.NewIcon(icons.NotificationSDCard))
	notification["NotificationSIMCardAlert"] = mustIcon(material.NewIcon(icons.NotificationSIMCardAlert))
	notification["NotificationSMS"] = mustIcon(material.NewIcon(icons.NotificationSMS))
	notification["NotificationSMSFailed"] = mustIcon(material.NewIcon(icons.NotificationSMSFailed))
	notification["NotificationSync"] = mustIcon(material.NewIcon(icons.NotificationSync))
	notification["NotificationSyncDisabled"] = mustIcon(material.NewIcon(icons.NotificationSyncDisabled))
	notification["NotificationSyncProblem"] = mustIcon(material.NewIcon(icons.NotificationSyncProblem))
	notification["NotificationSystemUpdate"] = mustIcon(material.NewIcon(icons.NotificationSystemUpdate))
	notification["NotificationTapAndPlay"] = mustIcon(material.NewIcon(icons.NotificationTapAndPlay))
	notification["NotificationTimeToLeave"] = mustIcon(material.NewIcon(icons.NotificationTimeToLeave))
	notification["NotificationVibration"] = mustIcon(material.NewIcon(icons.NotificationVibration))
	notification["NotificationVoiceChat"] = mustIcon(material.NewIcon(icons.NotificationVoiceChat))
	notification["NotificationVPNLock"] = mustIcon(material.NewIcon(icons.NotificationVPNLock))
	notification["NotificationWC"] = mustIcon(material.NewIcon(icons.NotificationWC))
	notification["NotificationWiFi"] = mustIcon(material.NewIcon(icons.NotificationWiFi))

	places := make(map[string]*material.Icon)
	places["PlacesACUnit"] = mustIcon(material.NewIcon(icons.PlacesACUnit))
	places["PlacesAirportShuttle"] = mustIcon(material.NewIcon(icons.PlacesAirportShuttle))
	places["PlacesAllInclusive"] = mustIcon(material.NewIcon(icons.PlacesAllInclusive))
	places["PlacesBeachAccess"] = mustIcon(material.NewIcon(icons.PlacesBeachAccess))
	places["PlacesBusinessCenter"] = mustIcon(material.NewIcon(icons.PlacesBusinessCenter))
	places["PlacesCasino"] = mustIcon(material.NewIcon(icons.PlacesCasino))
	places["PlacesChildCare"] = mustIcon(material.NewIcon(icons.PlacesChildCare))
	places["PlacesChildFriendly"] = mustIcon(material.NewIcon(icons.PlacesChildFriendly))
	places["PlacesFitnessCenter"] = mustIcon(material.NewIcon(icons.PlacesFitnessCenter))
	places["PlacesFreeBreakfast"] = mustIcon(material.NewIcon(icons.PlacesFreeBreakfast))
	places["PlacesGolfCourse"] = mustIcon(material.NewIcon(icons.PlacesGolfCourse))
	places["PlacesHotTub"] = mustIcon(material.NewIcon(icons.PlacesHotTub))
	places["PlacesKitchen"] = mustIcon(material.NewIcon(icons.PlacesKitchen))
	places["PlacesPool"] = mustIcon(material.NewIcon(icons.PlacesPool))
	places["PlacesRoomService"] = mustIcon(material.NewIcon(icons.PlacesRoomService))
	places["PlacesRVHookup"] = mustIcon(material.NewIcon(icons.PlacesRVHookup))
	places["PlacesSmokeFree"] = mustIcon(material.NewIcon(icons.PlacesSmokeFree))
	places["PlacesSmokingRooms"] = mustIcon(material.NewIcon(icons.PlacesSmokingRooms))
	places["PlacesSpa"] = mustIcon(material.NewIcon(icons.PlacesSpa))

	social := make(map[string]*material.Icon)
	social["SocialCake"] = mustIcon(material.NewIcon(icons.SocialCake))
	social["SocialDomain"] = mustIcon(material.NewIcon(icons.SocialDomain))
	social["SocialGroup"] = mustIcon(material.NewIcon(icons.SocialGroup))
	social["SocialGroupAdd"] = mustIcon(material.NewIcon(icons.SocialGroupAdd))
	social["SocialLocationCity"] = mustIcon(material.NewIcon(icons.SocialLocationCity))
	social["SocialMood"] = mustIcon(material.NewIcon(icons.SocialMood))
	social["SocialMoodBad"] = mustIcon(material.NewIcon(icons.SocialMoodBad))
	social["SocialNotifications"] = mustIcon(material.NewIcon(icons.SocialNotifications))
	social["SocialNotificationsActive"] = mustIcon(material.NewIcon(icons.SocialNotificationsActive))
	social["SocialNotificationsNone"] = mustIcon(material.NewIcon(icons.SocialNotificationsNone))
	social["SocialNotificationsOff"] = mustIcon(material.NewIcon(icons.SocialNotificationsOff))
	social["SocialNotificationsPaused"] = mustIcon(material.NewIcon(icons.SocialNotificationsPaused))
	social["SocialPages"] = mustIcon(material.NewIcon(icons.SocialPages))
	social["SocialPartyMode"] = mustIcon(material.NewIcon(icons.SocialPartyMode))
	social["SocialPeople"] = mustIcon(material.NewIcon(icons.SocialPeople))
	social["SocialPeopleOutline"] = mustIcon(material.NewIcon(icons.SocialPeopleOutline))
	social["SocialPerson"] = mustIcon(material.NewIcon(icons.SocialPerson))
	social["SocialPersonAdd"] = mustIcon(material.NewIcon(icons.SocialPersonAdd))
	social["SocialPersonOutline"] = mustIcon(material.NewIcon(icons.SocialPersonOutline))
	social["SocialPlusOne"] = mustIcon(material.NewIcon(icons.SocialPlusOne))
	social["SocialPoll"] = mustIcon(material.NewIcon(icons.SocialPoll))
	social["SocialPublic"] = mustIcon(material.NewIcon(icons.SocialPublic))
	social["SocialSchool"] = mustIcon(material.NewIcon(icons.SocialSchool))
	social["SocialSentimentDissatisfied"] = mustIcon(material.NewIcon(icons.SocialSentimentDissatisfied))
	social["SocialSentimentNeutral"] = mustIcon(material.NewIcon(icons.SocialSentimentNeutral))
	social["SocialSentimentSatisfied"] = mustIcon(material.NewIcon(icons.SocialSentimentSatisfied))
	social["SocialSentimentVeryDissatisfied"] = mustIcon(material.NewIcon(icons.SocialSentimentVeryDissatisfied))
	social["SocialSentimentVerySatisfied"] = mustIcon(material.NewIcon(icons.SocialSentimentVerySatisfied))
	social["SocialShare"] = mustIcon(material.NewIcon(icons.SocialShare))
	social["SocialWhatsHot"] = mustIcon(material.NewIcon(icons.SocialWhatsHot))

	toggle := make(map[string]*material.Icon)
	toggle["ToggleCheckBox"] = mustIcon(material.NewIcon(icons.ToggleCheckBox))
	toggle["ToggleCheckBoxOutlineBlank"] = mustIcon(material.NewIcon(icons.ToggleCheckBoxOutlineBlank))
	toggle["ToggleIndeterminateCheckBox"] = mustIcon(material.NewIcon(icons.ToggleIndeterminateCheckBox))
	toggle["ToggleRadioButtonChecked"] = mustIcon(material.NewIcon(icons.ToggleRadioButtonChecked))
	toggle["ToggleRadioButtonUnchecked"] = mustIcon(material.NewIcon(icons.ToggleRadioButtonUnchecked))
	toggle["ToggleStar"] = mustIcon(material.NewIcon(icons.ToggleStar))
	toggle["ToggleStarBorder"] = mustIcon(material.NewIcon(icons.ToggleStarBorder))
	toggle["ToggleStarHalf"] = mustIcon(material.NewIcon(icons.ToggleStarHalf))

	go func() {
		w := app.NewWindow(
			app.Size(unit.Dp(400), unit.Dp(800)),
			app.Title("ParallelCoin"),
		)
		gtx := layout.NewContext(w.Queue())
		for e := range w.Events() {
			if e, ok := e.(system.FrameEvent); ok {
				gtx.Reset(e.Config, e.Size)
				widgets := []func(){
					renderIcon(gtx, th, i["Action3DRotation"], "Action3DRotation"),
					renderIcon(gtx, th, i["ActionAccessibility"], "ActionAccessibility"),
					renderIcon(gtx, th, i["ActionAccessible"], "ActionAccessible"),
					renderIcon(gtx, th, i["ActionAccountBalance"], "ActionAccountBalance"),
					renderIcon(gtx, th, i["ActionAccountBalanceWallet"], "ActionAccountBalanceWallet"),
					renderIcon(gtx, th, i["ActionAccountBox"], "ActionAccountBox"),
					renderIcon(gtx, th, i["ActionAccountCircle"], "ActionAccountCircle"),
					renderIcon(gtx, th, i["ActionAddShoppingCart"], "ActionAddShoppingCart"),
					renderIcon(gtx, th, i["ActionAlarm"], "ActionAlarm"),
					renderIcon(gtx, th, i["ActionAlarmAdd"], "ActionAlarmAdd"),
					renderIcon(gtx, th, i["ActionAlarmOff"], "ActionAlarmOff"),
					renderIcon(gtx, th, i["ActionAlarmOn"], "ActionAlarmOn"),
					renderIcon(gtx, th, i["ActionAllOut"], "ActionAllOut"),
					renderIcon(gtx, th, i["ActionAndroid"], "ActionAndroid"),
					renderIcon(gtx, th, i["ActionAnnouncement"], "ActionAnnouncement"),
					renderIcon(gtx, th, i["ActionAspectRatio"], "ActionAspectRatio"),
					renderIcon(gtx, th, i["ActionAssessment"], "ActionAssessment"),
					renderIcon(gtx, th, i["ActionAssignment"], "ActionAssignment"),
					renderIcon(gtx, th, i["ActionAssignmentInd"], "ActionAssignmentInd"),
					renderIcon(gtx, th, i["ActionAssignmentLate"], "ActionAssignmentLate"),
					renderIcon(gtx, th, i["ActionAssignmentReturn"], "ActionAssignmentReturn"),
					renderIcon(gtx, th, i["ActionAssignmentReturned"], "ActionAssignmentReturned"),
					renderIcon(gtx, th, i["ActionAssignmentTurnedIn"], "ActionAssignmentTurnedIn"),
					renderIcon(gtx, th, i["ActionAutorenew"], "ActionAutorenew"),
					renderIcon(gtx, th, i["ActionBackup"], "ActionBackup"),
					renderIcon(gtx, th, i["ActionBook"], "ActionBook"),
					renderIcon(gtx, th, i["ActionBookmark"], "ActionBookmark"),
					renderIcon(gtx, th, i["ActionBookmarkBorder"], "ActionBookmarkBorder"),
					renderIcon(gtx, th, i["ActionBugReport"], "ActionBugReport"),
					renderIcon(gtx, th, i["ActionBuild"], "ActionBuild"),
					renderIcon(gtx, th, i["ActionCached"], "ActionCached"),
					renderIcon(gtx, th, i["ActionCameraEnhance"], "ActionCameraEnhance"),
					renderIcon(gtx, th, i["ActionCardGiftcard"], "ActionCardGiftcard"),
					renderIcon(gtx, th, i["ActionCardMembership"], "ActionCardMembership"),
					renderIcon(gtx, th, i["ActionCardTravel"], "ActionCardTravel"),
					renderIcon(gtx, th, i["ActionChangeHistory"], "ActionChangeHistory"),
					renderIcon(gtx, th, i["ActionCheckCircle"], "ActionCheckCircle"),
					renderIcon(gtx, th, i["ActionChromeReaderMode"], "ActionChromeReaderMode"),
					renderIcon(gtx, th, i["ActionClass"], "ActionClass"),
					renderIcon(gtx, th, i["ActionCode"], "ActionCode"),
					renderIcon(gtx, th, i["ActionCompareArrows"], "ActionCompareArrows"),
					renderIcon(gtx, th, i["ActionCopyright"], "ActionCopyright"),
					renderIcon(gtx, th, i["ActionCreditCard"], "ActionCreditCard"),
					renderIcon(gtx, th, i["ActionDashboard"], "ActionDashboard"),
					renderIcon(gtx, th, i["ActionDateRange"], "ActionDateRange"),
					renderIcon(gtx, th, i["ActionDelete"], "ActionDelete"),
					renderIcon(gtx, th, i["ActionDeleteForever"], "ActionDeleteForever"),
					renderIcon(gtx, th, i["ActionDescription"], "ActionDescription"),
					renderIcon(gtx, th, i["ActionDNS"], "ActionDNS"),
					renderIcon(gtx, th, i["ActionDone"], "ActionDone"),
					renderIcon(gtx, th, i["ActionDoneAll"], "ActionDoneAll"),
					renderIcon(gtx, th, i["ActionDonutLarge"], "ActionDonutLarge"),
					renderIcon(gtx, th, i["ActionDonutSmall"], "ActionDonutSmall"),
					renderIcon(gtx, th, i["ActionEject"], "ActionEject"),
					renderIcon(gtx, th, i["ActionEuroSymbol"], "ActionEuroSymbol"),
					renderIcon(gtx, th, i["ActionEvent"], "ActionEvent"),
					renderIcon(gtx, th, i["ActionEventSeat"], "ActionEventSeat"),
					renderIcon(gtx, th, i["ActionExitToApp"], "ActionExitToApp"),
					renderIcon(gtx, th, i["ActionExplore"], "ActionExplore"),
					renderIcon(gtx, th, i["ActionExtension"], "ActionExtension"),
					renderIcon(gtx, th, i["ActionFace"], "ActionFace"),
					renderIcon(gtx, th, i["ActionFavorite"], "ActionFavorite"),
					renderIcon(gtx, th, i["ActionFavoriteBorder"], "ActionFavoriteBorder"),
					renderIcon(gtx, th, i["ActionFeedback"], "ActionFeedback"),
					renderIcon(gtx, th, i["ActionFindInPage"], "ActionFindInPage"),
					renderIcon(gtx, th, i["ActionFindReplace"], "ActionFindReplace"),
					renderIcon(gtx, th, i["ActionFingerprint"], "ActionFingerprint"),
					renderIcon(gtx, th, i["ActionFlightLand"], "ActionFlightLand"),
					renderIcon(gtx, th, i["ActionFlightTakeoff"], "ActionFlightTakeoff"),
					renderIcon(gtx, th, i["ActionFlipToBack"], "ActionFlipToBack"),
					renderIcon(gtx, th, i["ActionFlipToFront"], "ActionFlipToFront"),
					renderIcon(gtx, th, i["ActionGTranslate"], "ActionGTranslate"),
					renderIcon(gtx, th, i["ActionGavel"], "ActionGavel"),
					renderIcon(gtx, th, i["ActionGetApp"], "ActionGetApp"),
					renderIcon(gtx, th, i["ActionGIF"], "ActionGIF"),
					renderIcon(gtx, th, i["ActionGrade"], "ActionGrade"),
					renderIcon(gtx, th, i["ActionGroupWork"], "ActionGroupWork"),
					renderIcon(gtx, th, i["ActionHelp"], "ActionHelp"),
					renderIcon(gtx, th, i["ActionHelpOutline"], "ActionHelpOutline"),
					renderIcon(gtx, th, i["ActionHighlightOff"], "ActionHighlightOff"),
					renderIcon(gtx, th, i["ActionHistory"], "ActionHistory"),
					renderIcon(gtx, th, i["ActionHome"], "ActionHome"),
					renderIcon(gtx, th, i["ActionHourglassEmpty"], "ActionHourglassEmpty"),
					renderIcon(gtx, th, i["ActionHourglassFull"], "ActionHourglassFull"),
					renderIcon(gtx, th, i["ActionHTTP"], "ActionHTTP"),
					renderIcon(gtx, th, i["ActionHTTPS"], "ActionHTTPS"),
					renderIcon(gtx, th, i["ActionImportantDevices"], "ActionImportantDevices"),
					renderIcon(gtx, th, i["ActionInfo"], "ActionInfo"),
					renderIcon(gtx, th, i["ActionInfoOutline"], "ActionInfoOutline"),
					renderIcon(gtx, th, i["ActionInput"], "ActionInput"),
					renderIcon(gtx, th, i["ActionInvertColors"], "ActionInvertColors"),
					renderIcon(gtx, th, i["ActionLabel"], "ActionLabel"),
					renderIcon(gtx, th, i["ActionLabelOutline"], "ActionLabelOutline"),
					renderIcon(gtx, th, i["ActionLanguage"], "ActionLanguage"),
					renderIcon(gtx, th, i["ActionLaunch"], "ActionLaunch"),
					renderIcon(gtx, th, i["ActionLightbulbOutline"], "ActionLightbulbOutline"),
					renderIcon(gtx, th, i["ActionLineStyle"], "ActionLineStyle"),
					renderIcon(gtx, th, i["ActionLineWeight"], "ActionLineWeight"),
					renderIcon(gtx, th, i["ActionList"], "ActionList"),
					renderIcon(gtx, th, i["ActionLock"], "ActionLock"),
					renderIcon(gtx, th, i["ActionLockOpen"], "ActionLockOpen"),
					renderIcon(gtx, th, i["ActionLockOutline"], "ActionLockOutline"),
					renderIcon(gtx, th, i["ActionLoyalty"], "ActionLoyalty"),
					renderIcon(gtx, th, i["ActionMarkUnreadMailbox"], "ActionMarkUnreadMailbox"),
					renderIcon(gtx, th, i["ActionMotorcycle"], "ActionMotorcycle"),
					renderIcon(gtx, th, i["ActionNoteAdd"], "ActionNoteAdd"),
					renderIcon(gtx, th, i["ActionOfflinePin"], "ActionOfflinePin"),
					renderIcon(gtx, th, i["ActionOpacity"], "ActionOpacity"),
					renderIcon(gtx, th, i["ActionOpenInBrowser"], "ActionOpenInBrowser"),
					renderIcon(gtx, th, i["ActionOpenInNew"], "ActionOpenInNew"),
					renderIcon(gtx, th, i["ActionOpenWith"], "ActionOpenWith"),
					renderIcon(gtx, th, i["ActionPageview"], "ActionPageview"),
					renderIcon(gtx, th, i["ActionPanTool"], "ActionPanTool"),
					renderIcon(gtx, th, i["ActionPayment"], "ActionPayment"),
					renderIcon(gtx, th, i["ActionPermCameraMic"], "ActionPermCameraMic"),
					renderIcon(gtx, th, i["ActionPermContactCalendar"], "ActionPermContactCalendar"),
					renderIcon(gtx, th, i["ActionPermDataSetting"], "ActionPermDataSetting"),
					renderIcon(gtx, th, i["ActionPermDeviceInformation"], "ActionPermDeviceInformation"),
					renderIcon(gtx, th, i["ActionPermIdentity"], "ActionPermIdentity"),
					renderIcon(gtx, th, i["ActionPermMedia"], "ActionPermMedia"),
					renderIcon(gtx, th, i["ActionPermPhoneMsg"], "ActionPermPhoneMsg"),
					renderIcon(gtx, th, i["ActionPermScanWiFi"], "ActionPermScanWiFi"),
					renderIcon(gtx, th, i["ActionPets"], "ActionPets"),
					renderIcon(gtx, th, i["ActionPictureInPicture"], "ActionPictureInPicture"),
					renderIcon(gtx, th, i["ActionPictureInPictureAlt"], "ActionPictureInPictureAlt"),
					renderIcon(gtx, th, i["ActionPlayForWork"], "ActionPlayForWork"),
					renderIcon(gtx, th, i["ActionPolymer"], "ActionPolymer"),
					renderIcon(gtx, th, i["ActionPowerSettingsNew"], "ActionPowerSettingsNew"),
					renderIcon(gtx, th, i["ActionPregnantWoman"], "ActionPregnantWoman"),
					renderIcon(gtx, th, i["ActionPrint"], "ActionPrint"),
					renderIcon(gtx, th, i["ActionQueryBuilder"], "ActionQueryBuilder"),
					renderIcon(gtx, th, i["ActionQuestionAnswer"], "ActionQuestionAnswer"),
					renderIcon(gtx, th, i["ActionReceipt"], "ActionReceipt"),
					renderIcon(gtx, th, i["ActionRecordVoiceOver"], "ActionRecordVoiceOver"),
					renderIcon(gtx, th, i["ActionRedeem"], "ActionRedeem"),
					renderIcon(gtx, th, i["ActionRemoveShoppingCart"], "ActionRemoveShoppingCart"),
					renderIcon(gtx, th, i["ActionReorder"], "ActionReorder"),
					renderIcon(gtx, th, i["ActionReportProblem"], "ActionReportProblem"),
					renderIcon(gtx, th, i["ActionRestore"], "ActionRestore"),
					renderIcon(gtx, th, i["ActionRestorePage"], "ActionRestorePage"),
					renderIcon(gtx, th, i["ActionRoom"], "ActionRoom"),
					renderIcon(gtx, th, i["ActionRoundedCorner"], "ActionRoundedCorner"),
					renderIcon(gtx, th, i["ActionRowing"], "ActionRowing"),
					renderIcon(gtx, th, i["ActionSchedule"], "ActionSchedule"),
					renderIcon(gtx, th, i["ActionSearch"], "ActionSearch"),
					renderIcon(gtx, th, i["ActionSettings"], "ActionSettings"),
					renderIcon(gtx, th, i["ActionSettingsApplications"], "ActionSettingsApplications"),
					renderIcon(gtx, th, i["ActionSettingsBackupRestore"], "ActionSettingsBackupRestore"),
					renderIcon(gtx, th, i["ActionSettingsBluetooth"], "ActionSettingsBluetooth"),
					renderIcon(gtx, th, i["ActionSettingsBrightness"], "ActionSettingsBrightness"),
					renderIcon(gtx, th, i["ActionSettingsCell"], "ActionSettingsCell"),
					renderIcon(gtx, th, i["ActionSettingsEthernet"], "ActionSettingsEthernet"),
					renderIcon(gtx, th, i["ActionSettingsInputAntenna"], "ActionSettingsInputAntenna"),
					renderIcon(gtx, th, i["ActionSettingsInputComponent"], "ActionSettingsInputComponent"),
					renderIcon(gtx, th, i["ActionSettingsInputComposite"], "ActionSettingsInputComposite"),
					renderIcon(gtx, th, i["ActionSettingsInputHDMI"], "ActionSettingsInputHDMI"),
					renderIcon(gtx, th, i["ActionSettingsInputSVideo"], "ActionSettingsInputSVideo"),
					renderIcon(gtx, th, i["ActionSettingsOverscan"], "ActionSettingsOverscan"),
					renderIcon(gtx, th, i["ActionSettingsPhone"], "ActionSettingsPhone"),
					renderIcon(gtx, th, i["ActionSettingsPower"], "ActionSettingsPower"),
					renderIcon(gtx, th, i["ActionSettingsRemote"], "ActionSettingsRemote"),
					renderIcon(gtx, th, i["ActionSettingsVoice"], "ActionSettingsVoice"),
					renderIcon(gtx, th, i["ActionShop"], "ActionShop"),
					renderIcon(gtx, th, i["ActionShopTwo"], "ActionShopTwo"),
					renderIcon(gtx, th, i["ActionShoppingBasket"], "ActionShoppingBasket"),
					renderIcon(gtx, th, i["ActionShoppingCart"], "ActionShoppingCart"),
					renderIcon(gtx, th, i["ActionSpeakerNotes"], "ActionSpeakerNotes"),
					renderIcon(gtx, th, i["ActionSpeakerNotesOff"], "ActionSpeakerNotesOff"),
					renderIcon(gtx, th, i["ActionSpellcheck"], "ActionSpellcheck"),
					renderIcon(gtx, th, i["ActionStarRate"], "ActionStarRate"),
					renderIcon(gtx, th, i["ActionStars"], "ActionStars"),
					renderIcon(gtx, th, i["ActionStore"], "ActionStore"),
					renderIcon(gtx, th, i["ActionSubject"], "ActionSubject"),
					renderIcon(gtx, th, i["ActionSupervisorAccount"], "ActionSupervisorAccount"),
					renderIcon(gtx, th, i["ActionSwapHoriz"], "ActionSwapHoriz"),
					renderIcon(gtx, th, i["ActionSwapVert"], "ActionSwapVert"),
					renderIcon(gtx, th, i["ActionSwapVerticalCircle"], "ActionSwapVerticalCircle"),
					renderIcon(gtx, th, i["ActionSystemUpdateAlt"], "ActionSystemUpdateAlt"),
					renderIcon(gtx, th, i["ActionTab"], "ActionTab"),
					renderIcon(gtx, th, i["ActionTabUnselected"], "ActionTabUnselected"),
					renderIcon(gtx, th, i["ActionTheaters"], "ActionTheaters"),
					renderIcon(gtx, th, i["ActionThumbDown"], "ActionThumbDown"),
					renderIcon(gtx, th, i["ActionThumbUp"], "ActionThumbUp"),
					renderIcon(gtx, th, i["ActionThumbsUpDown"], "ActionThumbsUpDown"),
					renderIcon(gtx, th, i["ActionTimeline"], "ActionTimeline"),
					renderIcon(gtx, th, i["ActionTOC"], "ActionTOC"),
					renderIcon(gtx, th, i["ActionToday"], "ActionToday"),
					renderIcon(gtx, th, i["ActionToll"], "ActionToll"),
					renderIcon(gtx, th, i["ActionTouchApp"], "ActionTouchApp"),
					renderIcon(gtx, th, i["ActionTrackChanges"], "ActionTrackChanges"),
					renderIcon(gtx, th, i["ActionTranslate"], "ActionTranslate"),
					renderIcon(gtx, th, i["ActionTrendingDown"], "ActionTrendingDown"),
					renderIcon(gtx, th, i["ActionTrendingFlat"], "ActionTrendingFlat"),
					renderIcon(gtx, th, i["ActionTrendingUp"], "ActionTrendingUp"),
					renderIcon(gtx, th, i["ActionTurnedIn"], "ActionTurnedIn"),
					renderIcon(gtx, th, i["ActionTurnedInNot"], "ActionTurnedInNot"),
					renderIcon(gtx, th, i["ActionUpdate"], "ActionUpdate"),
					renderIcon(gtx, th, i["ActionVerifiedUser"], "ActionVerifiedUser"),
					renderIcon(gtx, th, i["ActionViewAgenda"], "ActionViewAgenda"),
					renderIcon(gtx, th, i["ActionViewArray"], "ActionViewArray"),
					renderIcon(gtx, th, i["ActionViewCarousel"], "ActionViewCarousel"),
					renderIcon(gtx, th, i["ActionViewColumn"], "ActionViewColumn"),
					renderIcon(gtx, th, i["ActionViewDay"], "ActionViewDay"),
					renderIcon(gtx, th, i["ActionViewHeadline"], "ActionViewHeadline"),
					renderIcon(gtx, th, i["ActionViewList"], "ActionViewList"),
					renderIcon(gtx, th, i["ActionViewModule"], "ActionViewModule"),
					renderIcon(gtx, th, i["ActionViewQuilt"], "ActionViewQuilt"),
					renderIcon(gtx, th, i["ActionViewStream"], "ActionViewStream"),
					renderIcon(gtx, th, i["ActionViewWeek"], "ActionViewWeek"),
					renderIcon(gtx, th, i["ActionVisibility"], "ActionVisibility"),
					renderIcon(gtx, th, i["ActionVisibilityOff"], "ActionVisibilityOff"),
					renderIcon(gtx, th, i["ActionWatchLater"], "ActionWatchLater"),
					renderIcon(gtx, th, i["ActionWork"], "ActionWork"),
					renderIcon(gtx, th, i["ActionYoutubeSearchedFor"], "ActionYoutubeSearchedFor"),
					renderIcon(gtx, th, i["ActionZoomIn"], "ActionZoomIn"),
					renderIcon(gtx, th, i["ActionZoomOut"], "ActionZoomOut"),
					renderIcon(gtx, th, i["AlertAddAlert"], "AlertAddAlert"),
					renderIcon(gtx, th, i["AlertError"], "AlertError"),
					renderIcon(gtx, th, i["AlertErrorOutline"], "AlertErrorOutline"),
					renderIcon(gtx, th, i["AlertWarning"], "AlertWarning"),
					renderIcon(gtx, th, i["AVAddToQueue"], "AVAddToQueue"),
					renderIcon(gtx, th, i["AVAirplay"], "AVAirplay"),
					renderIcon(gtx, th, i["AVAlbum"], "AVAlbum"),
					renderIcon(gtx, th, i["AVArtTrack"], "AVArtTrack"),
					renderIcon(gtx, th, i["AVAVTimer"], "AVAVTimer"),
					renderIcon(gtx, th, i["AVBrandingWatermark"], "AVBrandingWatermark"),
					renderIcon(gtx, th, i["AVCallToAction"], "AVCallToAction"),
					renderIcon(gtx, th, i["AVClosedCaption"], "AVClosedCaption"),
					renderIcon(gtx, th, i["AVEqualizer"], "AVEqualizer"),
					renderIcon(gtx, th, i["AVExplicit"], "AVExplicit"),
					renderIcon(gtx, th, i["AVFastForward"], "AVFastForward"),
					renderIcon(gtx, th, i["AVFastRewind"], "AVFastRewind"),
					renderIcon(gtx, th, i["AVFeaturedPlayList"], "AVFeaturedPlayList"),
					renderIcon(gtx, th, i["AVFeaturedVideo"], "AVFeaturedVideo"),
					renderIcon(gtx, th, i["AVFiberDVR"], "AVFiberDVR"),
					renderIcon(gtx, th, i["AVFiberManualRecord"], "AVFiberManualRecord"),
					renderIcon(gtx, th, i["AVFiberNew"], "AVFiberNew"),
					renderIcon(gtx, th, i["AVFiberPin"], "AVFiberPin"),
					renderIcon(gtx, th, i["AVFiberSmartRecord"], "AVFiberSmartRecord"),
					renderIcon(gtx, th, i["AVForward10"], "AVForward10"),
					renderIcon(gtx, th, i["AVForward30"], "AVForward30"),
					renderIcon(gtx, th, i["AVForward5"], "AVForward5"),
					renderIcon(gtx, th, i["AVGames"], "AVGames"),
					renderIcon(gtx, th, i["AVHD"], "AVHD"),
					renderIcon(gtx, th, i["AVHearing"], "AVHearing"),
					renderIcon(gtx, th, i["AVHighQuality"], "AVHighQuality"),
					renderIcon(gtx, th, i["AVLibraryAdd"], "AVLibraryAdd"),
					renderIcon(gtx, th, i["AVLibraryBooks"], "AVLibraryBooks"),
					renderIcon(gtx, th, i["AVLibraryMusic"], "AVLibraryMusic"),
					renderIcon(gtx, th, i["AVLoop"], "AVLoop"),
					renderIcon(gtx, th, i["AVMic"], "AVMic"),
					renderIcon(gtx, th, i["AVMicNone"], "AVMicNone"),
					renderIcon(gtx, th, i["AVMicOff"], "AVMicOff"),
					renderIcon(gtx, th, i["AVMovie"], "AVMovie"),
					renderIcon(gtx, th, i["AVMusicVideo"], "AVMusicVideo"),
					renderIcon(gtx, th, i["AVNewReleases"], "AVNewReleases"),
					renderIcon(gtx, th, i["AVNotInterested"], "AVNotInterested"),
					renderIcon(gtx, th, i["AVNote"], "AVNote"),
					renderIcon(gtx, th, i["AVPause"], "AVPause"),
					renderIcon(gtx, th, i["AVPauseCircleFilled"], "AVPauseCircleFilled"),
					renderIcon(gtx, th, i["AVPauseCircleOutline"], "AVPauseCircleOutline"),
					renderIcon(gtx, th, i["AVPlayArrow"], "AVPlayArrow"),
					renderIcon(gtx, th, i["AVPlayCircleFilled"], "AVPlayCircleFilled"),
					renderIcon(gtx, th, i["AVPlayCircleOutline"], "AVPlayCircleOutline"),
					renderIcon(gtx, th, i["AVPlaylistAdd"], "AVPlaylistAdd"),
					renderIcon(gtx, th, i["AVPlaylistAddCheck"], "AVPlaylistAddCheck"),
					renderIcon(gtx, th, i["AVPlaylistPlay"], "AVPlaylistPlay"),
					renderIcon(gtx, th, i["AVQueue"], "AVQueue"),
					renderIcon(gtx, th, i["AVQueueMusic"], "AVQueueMusic"),
					renderIcon(gtx, th, i["AVQueuePlayNext"], "AVQueuePlayNext"),
					renderIcon(gtx, th, i["AVRadio"], "AVRadio"),
					renderIcon(gtx, th, i["AVRecentActors"], "AVRecentActors"),
					renderIcon(gtx, th, i["AVRemoveFromQueue"], "AVRemoveFromQueue"),
					renderIcon(gtx, th, i["AVRepeat"], "AVRepeat"),
					renderIcon(gtx, th, i["AVRepeatOne"], "AVRepeatOne"),
					renderIcon(gtx, th, i["AVReplay"], "AVReplay"),
					renderIcon(gtx, th, i["AVReplay10"], "AVReplay10"),
					renderIcon(gtx, th, i["AVReplay30"], "AVReplay30"),
					renderIcon(gtx, th, i["AVReplay5"], "AVReplay5"),
					renderIcon(gtx, th, i["AVShuffle"], "AVShuffle"),
					renderIcon(gtx, th, i["AVSkipNext"], "AVSkipNext"),
					renderIcon(gtx, th, i["AVSkipPrevious"], "AVSkipPrevious"),
					renderIcon(gtx, th, i["AVSlowMotionVideo"], "AVSlowMotionVideo"),
					renderIcon(gtx, th, i["AVSnooze"], "AVSnooze"),
					renderIcon(gtx, th, i["AVSortByAlpha"], "AVSortByAlpha"),
					renderIcon(gtx, th, i["AVStop"], "AVStop"),
					renderIcon(gtx, th, i["AVSubscriptions"], "AVSubscriptions"),
					renderIcon(gtx, th, i["AVSubtitles"], "AVSubtitles"),
					renderIcon(gtx, th, i["AVSurroundSound"], "AVSurroundSound"),
					renderIcon(gtx, th, i["AVVideoCall"], "AVVideoCall"),
					renderIcon(gtx, th, i["AVVideoLabel"], "AVVideoLabel"),
					renderIcon(gtx, th, i["AVVideoLibrary"], "AVVideoLibrary"),
					renderIcon(gtx, th, i["AVVideocam"], "AVVideocam"),
					renderIcon(gtx, th, i["AVVideocamOff"], "AVVideocamOff"),
					renderIcon(gtx, th, i["AVVolumeDown"], "AVVolumeDown"),
					renderIcon(gtx, th, i["AVVolumeMute"], "AVVolumeMute"),
					renderIcon(gtx, th, i["AVVolumeOff"], "AVVolumeOff"),
					renderIcon(gtx, th, i["AVVolumeUp"], "AVVolumeUp"),
					renderIcon(gtx, th, i["AVWeb"], "AVWeb"),
					renderIcon(gtx, th, i["AVWebAsset"], "AVWebAsset"),
					renderIcon(gtx, th, i["CommunicationBusiness"], "CommunicationBusiness"),
					renderIcon(gtx, th, i["CommunicationCall"], "CommunicationCall"),
					renderIcon(gtx, th, i["CommunicationCallEnd"], "CommunicationCallEnd"),
					renderIcon(gtx, th, i["CommunicationCallMade"], "CommunicationCallMade"),
					renderIcon(gtx, th, i["CommunicationCallMerge"], "CommunicationCallMerge"),
					renderIcon(gtx, th, i["CommunicationCallMissed"], "CommunicationCallMissed"),
					renderIcon(gtx, th, i["CommunicationCallMissedOutgoing"], "CommunicationCallMissedOutgoing"),
					renderIcon(gtx, th, i["CommunicationCallReceived"], "CommunicationCallReceived"),
					renderIcon(gtx, th, i["CommunicationCallSplit"], "CommunicationCallSplit"),
					renderIcon(gtx, th, i["CommunicationChat"], "CommunicationChat"),
					renderIcon(gtx, th, i["CommunicationChatBubble"], "CommunicationChatBubble"),
					renderIcon(gtx, th, i["CommunicationChatBubbleOutline"], "CommunicationChatBubbleOutline"),
					renderIcon(gtx, th, i["CommunicationClearAll"], "CommunicationClearAll"),
					renderIcon(gtx, th, i["CommunicationComment"], "CommunicationComment"),
					renderIcon(gtx, th, i["CommunicationContactMail"], "CommunicationContactMail"),
					renderIcon(gtx, th, i["CommunicationContactPhone"], "CommunicationContactPhone"),
					renderIcon(gtx, th, i["CommunicationContacts"], "CommunicationContacts"),
					renderIcon(gtx, th, i["CommunicationDialerSIP"], "CommunicationDialerSIP"),
					renderIcon(gtx, th, i["CommunicationDialpad"], "CommunicationDialpad"),
					renderIcon(gtx, th, i["CommunicationEmail"], "CommunicationEmail"),
					renderIcon(gtx, th, i["CommunicationForum"], "CommunicationForum"),
					renderIcon(gtx, th, i["CommunicationImportContacts"], "CommunicationImportContacts"),
					renderIcon(gtx, th, i["CommunicationImportExport"], "CommunicationImportExport"),
					renderIcon(gtx, th, i["CommunicationInvertColorsOff"], "CommunicationInvertColorsOff"),
					renderIcon(gtx, th, i["CommunicationLiveHelp"], "CommunicationLiveHelp"),
					renderIcon(gtx, th, i["CommunicationLocationOff"], "CommunicationLocationOff"),
					renderIcon(gtx, th, i["CommunicationLocationOn"], "CommunicationLocationOn"),
					renderIcon(gtx, th, i["CommunicationMailOutline"], "CommunicationMailOutline"),
					renderIcon(gtx, th, i["CommunicationMessage"], "CommunicationMessage"),
					renderIcon(gtx, th, i["CommunicationNoSIM"], "CommunicationNoSIM"),
					renderIcon(gtx, th, i["CommunicationPhone"], "CommunicationPhone"),
					renderIcon(gtx, th, i["CommunicationPhoneLinkErase"], "CommunicationPhoneLinkErase"),
					renderIcon(gtx, th, i["CommunicationPhoneLinkLock"], "CommunicationPhoneLinkLock"),
					renderIcon(gtx, th, i["CommunicationPhoneLinkRing"], "CommunicationPhoneLinkRing"),
					renderIcon(gtx, th, i["CommunicationPhoneLinkSetup"], "CommunicationPhoneLinkSetup"),
					renderIcon(gtx, th, i["CommunicationPortableWiFiOff"], "CommunicationPortableWiFiOff"),
					renderIcon(gtx, th, i["CommunicationPresentToAll"], "CommunicationPresentToAll"),
					renderIcon(gtx, th, i["CommunicationRingVolume"], "CommunicationRingVolume"),
					renderIcon(gtx, th, i["CommunicationRSSFeed"], "CommunicationRSSFeed"),
					renderIcon(gtx, th, i["CommunicationScreenShare"], "CommunicationScreenShare"),
					renderIcon(gtx, th, i["CommunicationSpeakerPhone"], "CommunicationSpeakerPhone"),
					renderIcon(gtx, th, i["CommunicationStayCurrentLandscape"], "CommunicationStayCurrentLandscape"),
					renderIcon(gtx, th, i["CommunicationStayCurrentPortrait"], "CommunicationStayCurrentPortrait"),
					renderIcon(gtx, th, i["CommunicationStayPrimaryLandscape"], "CommunicationStayPrimaryLandscape"),
					renderIcon(gtx, th, i["CommunicationStayPrimaryPortrait"], "CommunicationStayPrimaryPortrait"),
					renderIcon(gtx, th, i["CommunicationStopScreenShare"], "CommunicationStopScreenShare"),
					renderIcon(gtx, th, i["CommunicationSwapCalls"], "CommunicationSwapCalls"),
					renderIcon(gtx, th, i["CommunicationTextSMS"], "CommunicationTextSMS"),
					renderIcon(gtx, th, i["CommunicationVoicemail"], "CommunicationVoicemail"),
					renderIcon(gtx, th, i["CommunicationVPNKey"], "CommunicationVPNKey"),
					renderIcon(gtx, th, i["ContentAdd"], "ContentAdd"),
					renderIcon(gtx, th, i["ContentAddBox"], "ContentAddBox"),
					renderIcon(gtx, th, i["ContentAddCircle"], "ContentAddCircle"),
					renderIcon(gtx, th, i["ContentAddCircleOutline"], "ContentAddCircleOutline"),
					renderIcon(gtx, th, i["ContentArchive"], "ContentArchive"),
					renderIcon(gtx, th, i["ContentBackspace"], "ContentBackspace"),
					renderIcon(gtx, th, i["ContentBlock"], "ContentBlock"),
					renderIcon(gtx, th, i["ContentClear"], "ContentClear"),
					renderIcon(gtx, th, i["ContentContentCopy"], "ContentContentCopy"),
					renderIcon(gtx, th, i["ContentContentCut"], "ContentContentCut"),
					renderIcon(gtx, th, i["ContentContentPaste"], "ContentContentPaste"),
					renderIcon(gtx, th, i["ContentCreate"], "ContentCreate"),
					renderIcon(gtx, th, i["ContentDeleteSweep"], "ContentDeleteSweep"),
					renderIcon(gtx, th, i["ContentDrafts"], "ContentDrafts"),
					renderIcon(gtx, th, i["ContentFilterList"], "ContentFilterList"),
					renderIcon(gtx, th, i["ContentFlag"], "ContentFlag"),
					renderIcon(gtx, th, i["ContentFontDownload"], "ContentFontDownload"),
					renderIcon(gtx, th, i["ContentForward"], "ContentForward"),
					renderIcon(gtx, th, i["ContentGesture"], "ContentGesture"),
					renderIcon(gtx, th, i["ContentInbox"], "ContentInbox"),
					renderIcon(gtx, th, i["ContentLink"], "ContentLink"),
					renderIcon(gtx, th, i["ContentLowPriority"], "ContentLowPriority"),
					renderIcon(gtx, th, i["ContentMail"], "ContentMail"),
					renderIcon(gtx, th, i["ContentMarkUnread"], "ContentMarkUnread"),
					renderIcon(gtx, th, i["ContentMoveToInbox"], "ContentMoveToInbox"),
					renderIcon(gtx, th, i["ContentNextWeek"], "ContentNextWeek"),
					renderIcon(gtx, th, i["ContentRedo"], "ContentRedo"),
					renderIcon(gtx, th, i["ContentRemove"], "ContentRemove"),
					renderIcon(gtx, th, i["ContentRemoveCircle"], "ContentRemoveCircle"),
					renderIcon(gtx, th, i["ContentRemoveCircleOutline"], "ContentRemoveCircleOutline"),
					renderIcon(gtx, th, i["ContentReply"], "ContentReply"),
					renderIcon(gtx, th, i["ContentReplyAll"], "ContentReplyAll"),
					renderIcon(gtx, th, i["ContentReport"], "ContentReport"),
					renderIcon(gtx, th, i["ContentSave"], "ContentSave"),
					renderIcon(gtx, th, i["ContentSelectAll"], "ContentSelectAll"),
					renderIcon(gtx, th, i["ContentSend"], "ContentSend"),
					renderIcon(gtx, th, i["ContentSort"], "ContentSort"),
					renderIcon(gtx, th, i["ContentTextFormat"], "ContentTextFormat"),
					renderIcon(gtx, th, i["ContentUnarchive"], "ContentUnarchive"),
					renderIcon(gtx, th, i["ContentUndo"], "ContentUndo"),
					renderIcon(gtx, th, i["ContentWeekend"], "ContentWeekend"),
					renderIcon(gtx, th, i["DeviceAccessAlarm"], "DeviceAccessAlarm"),
					renderIcon(gtx, th, i["DeviceAccessAlarms"], "DeviceAccessAlarms"),
					renderIcon(gtx, th, i["DeviceAccessTime"], "DeviceAccessTime"),
					renderIcon(gtx, th, i["DeviceAddAlarm"], "DeviceAddAlarm"),
					renderIcon(gtx, th, i["DeviceAirplaneModeActive"], "DeviceAirplaneModeActive"),
					renderIcon(gtx, th, i["DeviceAirplaneModeInactive"], "DeviceAirplaneModeInactive"),
					renderIcon(gtx, th, i["DeviceBattery20"], "DeviceBattery20"),
					renderIcon(gtx, th, i["DeviceBattery30"], "DeviceBattery30"),
					renderIcon(gtx, th, i["DeviceBattery50"], "DeviceBattery50"),
					renderIcon(gtx, th, i["DeviceBattery60"], "DeviceBattery60"),
					renderIcon(gtx, th, i["DeviceBattery80"], "DeviceBattery80"),
					renderIcon(gtx, th, i["DeviceBattery90"], "DeviceBattery90"),
					renderIcon(gtx, th, i["DeviceBatteryAlert"], "DeviceBatteryAlert"),
					renderIcon(gtx, th, i["DeviceBatteryCharging20"], "DeviceBatteryCharging20"),
					renderIcon(gtx, th, i["DeviceBatteryCharging30"], "DeviceBatteryCharging30"),
					renderIcon(gtx, th, i["DeviceBatteryCharging50"], "DeviceBatteryCharging50"),
					renderIcon(gtx, th, i["DeviceBatteryCharging60"], "DeviceBatteryCharging60"),
					renderIcon(gtx, th, i["DeviceBatteryCharging80"], "DeviceBatteryCharging80"),
					renderIcon(gtx, th, i["DeviceBatteryCharging90"], "DeviceBatteryCharging90"),
					renderIcon(gtx, th, i["DeviceBatteryChargingFull"], "DeviceBatteryChargingFull"),
					renderIcon(gtx, th, i["DeviceBatteryFull"], "DeviceBatteryFull"),
					renderIcon(gtx, th, i["DeviceBatteryStd"], "DeviceBatteryStd"),
					renderIcon(gtx, th, i["DeviceBatteryUnknown"], "DeviceBatteryUnknown"),
					renderIcon(gtx, th, i["DeviceBluetooth"], "DeviceBluetooth"),
					renderIcon(gtx, th, i["DeviceBluetoothConnected"], "DeviceBluetoothConnected"),
					renderIcon(gtx, th, i["DeviceBluetoothDisabled"], "DeviceBluetoothDisabled"),
					renderIcon(gtx, th, i["DeviceBluetoothSearching"], "DeviceBluetoothSearching"),
					renderIcon(gtx, th, i["DeviceBrightnessAuto"], "DeviceBrightnessAuto"),
					renderIcon(gtx, th, i["DeviceBrightnessHigh"], "DeviceBrightnessHigh"),
					renderIcon(gtx, th, i["DeviceBrightnessLow"], "DeviceBrightnessLow"),
					renderIcon(gtx, th, i["DeviceBrightnessMedium"], "DeviceBrightnessMedium"),
					renderIcon(gtx, th, i["DeviceDataUsage"], "DeviceDataUsage"),
					renderIcon(gtx, th, i["DeviceDeveloperMode"], "DeviceDeveloperMode"),
					renderIcon(gtx, th, i["DeviceDevices"], "DeviceDevices"),
					renderIcon(gtx, th, i["DeviceDVR"], "DeviceDVR"),
					renderIcon(gtx, th, i["DeviceGPSFixed"], "DeviceGPSFixed"),
					renderIcon(gtx, th, i["DeviceGPSNotFixed"], "DeviceGPSNotFixed"),
					renderIcon(gtx, th, i["DeviceGPSOff"], "DeviceGPSOff"),
					renderIcon(gtx, th, i["DeviceGraphicEq"], "DeviceGraphicEq"),
					renderIcon(gtx, th, i["DeviceLocationDisabled"], "DeviceLocationDisabled"),
					renderIcon(gtx, th, i["DeviceLocationSearching"], "DeviceLocationSearching"),
					renderIcon(gtx, th, i["DeviceNetworkCell"], "DeviceNetworkCell"),
					renderIcon(gtx, th, i["DeviceNetworkWiFi"], "DeviceNetworkWiFi"),
					renderIcon(gtx, th, i["DeviceNFC"], "DeviceNFC"),
					renderIcon(gtx, th, i["DeviceScreenLockLandscape"], "DeviceScreenLockLandscape"),
					renderIcon(gtx, th, i["DeviceScreenLockPortrait"], "DeviceScreenLockPortrait"),
					renderIcon(gtx, th, i["DeviceScreenLockRotation"], "DeviceScreenLockRotation"),
					renderIcon(gtx, th, i["DeviceScreenRotation"], "DeviceScreenRotation"),
					renderIcon(gtx, th, i["DeviceSDStorage"], "DeviceSDStorage"),
					renderIcon(gtx, th, i["DeviceSettingsSystemDaydream"], "DeviceSettingsSystemDaydream"),
					renderIcon(gtx, th, i["DeviceSignalCellular0Bar"], "DeviceSignalCellular0Bar"),
					renderIcon(gtx, th, i["DeviceSignalCellular1Bar"], "DeviceSignalCellular1Bar"),
					renderIcon(gtx, th, i["DeviceSignalCellular2Bar"], "DeviceSignalCellular2Bar"),
					renderIcon(gtx, th, i["DeviceSignalCellular3Bar"], "DeviceSignalCellular3Bar"),
					renderIcon(gtx, th, i["DeviceSignalCellular4Bar"], "DeviceSignalCellular4Bar"),
					renderIcon(gtx, th, i["DeviceSignalCellularConnectedNoInternet0Bar"], "DeviceSignalCellularConnectedNoInternet0Bar"),
					renderIcon(gtx, th, i["DeviceSignalCellularConnectedNoInternet1Bar"], "DeviceSignalCellularConnectedNoInternet1Bar"),
					renderIcon(gtx, th, i["DeviceSignalCellularConnectedNoInternet2Bar"], "DeviceSignalCellularConnectedNoInternet2Bar"),
					renderIcon(gtx, th, i["DeviceSignalCellularConnectedNoInternet3Bar"], "DeviceSignalCellularConnectedNoInternet3Bar"),
					renderIcon(gtx, th, i["DeviceSignalCellularConnectedNoInternet4Bar"], "DeviceSignalCellularConnectedNoInternet4Bar"),
					renderIcon(gtx, th, i["DeviceSignalCellularNoSIM"], "DeviceSignalCellularNoSIM"),
					renderIcon(gtx, th, i["DeviceSignalCellularNull"], "DeviceSignalCellularNull"),
					renderIcon(gtx, th, i["DeviceSignalCellularOff"], "DeviceSignalCellularOff"),
					renderIcon(gtx, th, i["DeviceSignalWiFi0Bar"], "DeviceSignalWiFi0Bar"),
					renderIcon(gtx, th, i["DeviceSignalWiFi1Bar"], "DeviceSignalWiFi1Bar"),
					renderIcon(gtx, th, i["DeviceSignalWiFi1BarLock"], "DeviceSignalWiFi1BarLock"),
					renderIcon(gtx, th, i["DeviceSignalWiFi2Bar"], "DeviceSignalWiFi2Bar"),
					renderIcon(gtx, th, i["DeviceSignalWiFi2BarLock"], "DeviceSignalWiFi2BarLock"),
					renderIcon(gtx, th, i["DeviceSignalWiFi3Bar"], "DeviceSignalWiFi3Bar"),
					renderIcon(gtx, th, i["DeviceSignalWiFi3BarLock"], "DeviceSignalWiFi3BarLock"),
					renderIcon(gtx, th, i["DeviceSignalWiFi4Bar"], "DeviceSignalWiFi4Bar"),
					renderIcon(gtx, th, i["DeviceSignalWiFi4BarLock"], "DeviceSignalWiFi4BarLock"),
					renderIcon(gtx, th, i["DeviceSignalWiFiOff"], "DeviceSignalWiFiOff"),
					renderIcon(gtx, th, i["DeviceStorage"], "DeviceStorage"),
					renderIcon(gtx, th, i["DeviceUSB"], "DeviceUSB"),
					renderIcon(gtx, th, i["DeviceWallpaper"], "DeviceWallpaper"),
					renderIcon(gtx, th, i["DeviceWidgets"], "DeviceWidgets"),
					renderIcon(gtx, th, i["DeviceWiFiLock"], "DeviceWiFiLock"),
					renderIcon(gtx, th, i["DeviceWiFiTethering"], "DeviceWiFiTethering"),
					renderIcon(gtx, th, i["EditorAttachFile"], "EditorAttachFile"),
					renderIcon(gtx, th, i["EditorAttachMoney"], "EditorAttachMoney"),
					renderIcon(gtx, th, i["EditorBorderAll"], "EditorBorderAll"),
					renderIcon(gtx, th, i["EditorBorderBottom"], "EditorBorderBottom"),
					renderIcon(gtx, th, i["EditorBorderClear"], "EditorBorderClear"),
					renderIcon(gtx, th, i["EditorBorderColor"], "EditorBorderColor"),
					renderIcon(gtx, th, i["EditorBorderHorizontal"], "EditorBorderHorizontal"),
					renderIcon(gtx, th, i["EditorBorderInner"], "EditorBorderInner"),
					renderIcon(gtx, th, i["EditorBorderLeft"], "EditorBorderLeft"),
					renderIcon(gtx, th, i["EditorBorderOuter"], "EditorBorderOuter"),
					renderIcon(gtx, th, i["EditorBorderRight"], "EditorBorderRight"),
					renderIcon(gtx, th, i["EditorBorderStyle"], "EditorBorderStyle"),
					renderIcon(gtx, th, i["EditorBorderTop"], "EditorBorderTop"),
					renderIcon(gtx, th, i["EditorBorderVertical"], "EditorBorderVertical"),
					renderIcon(gtx, th, i["EditorBubbleChart"], "EditorBubbleChart"),
					renderIcon(gtx, th, i["EditorDragHandle"], "EditorDragHandle"),
					renderIcon(gtx, th, i["EditorFormatAlignCenter"], "EditorFormatAlignCenter"),
					renderIcon(gtx, th, i["EditorFormatAlignJustify"], "EditorFormatAlignJustify"),
					renderIcon(gtx, th, i["EditorFormatAlignLeft"], "EditorFormatAlignLeft"),
					renderIcon(gtx, th, i["EditorFormatAlignRight"], "EditorFormatAlignRight"),
					renderIcon(gtx, th, i["EditorFormatBold"], "EditorFormatBold"),
					renderIcon(gtx, th, i["EditorFormatClear"], "EditorFormatClear"),
					renderIcon(gtx, th, i["EditorFormatColorFill"], "EditorFormatColorFill"),
					renderIcon(gtx, th, i["EditorFormatColorReset"], "EditorFormatColorReset"),
					renderIcon(gtx, th, i["EditorFormatColorText"], "EditorFormatColorText"),
					renderIcon(gtx, th, i["EditorFormatIndentDecrease"], "EditorFormatIndentDecrease"),
					renderIcon(gtx, th, i["EditorFormatIndentIncrease"], "EditorFormatIndentIncrease"),
					renderIcon(gtx, th, i["EditorFormatItalic"], "EditorFormatItalic"),
					renderIcon(gtx, th, i["EditorFormatLineSpacing"], "EditorFormatLineSpacing"),
					renderIcon(gtx, th, i["EditorFormatListBulleted"], "EditorFormatListBulleted"),
					renderIcon(gtx, th, i["EditorFormatListNumbered"], "EditorFormatListNumbered"),
					renderIcon(gtx, th, i["EditorFormatPaint"], "EditorFormatPaint"),
					renderIcon(gtx, th, i["EditorFormatQuote"], "EditorFormatQuote"),
					renderIcon(gtx, th, i["EditorFormatShapes"], "EditorFormatShapes"),
					renderIcon(gtx, th, i["EditorFormatSize"], "EditorFormatSize"),
					renderIcon(gtx, th, i["EditorFormatStrikethrough"], "EditorFormatStrikethrough"),
					renderIcon(gtx, th, i["EditorFormatTextDirectionLToR"], "EditorFormatTextDirectionLToR"),
					renderIcon(gtx, th, i["EditorFormatTextDirectionRToL"], "EditorFormatTextDirectionRToL"),
					renderIcon(gtx, th, i["EditorFormatUnderlined"], "EditorFormatUnderlined"),
					renderIcon(gtx, th, i["EditorFunctions"], "EditorFunctions"),
					renderIcon(gtx, th, i["EditorHighlight"], "EditorHighlight"),
					renderIcon(gtx, th, i["EditorInsertChart"], "EditorInsertChart"),
					renderIcon(gtx, th, i["EditorInsertComment"], "EditorInsertComment"),
					renderIcon(gtx, th, i["EditorInsertDriveFile"], "EditorInsertDriveFile"),
					renderIcon(gtx, th, i["EditorInsertEmoticon"], "EditorInsertEmoticon"),
					renderIcon(gtx, th, i["EditorInsertInvitation"], "EditorInsertInvitation"),
					renderIcon(gtx, th, i["EditorInsertLink"], "EditorInsertLink"),
					renderIcon(gtx, th, i["EditorInsertPhoto"], "EditorInsertPhoto"),
					renderIcon(gtx, th, i["EditorLinearScale"], "EditorLinearScale"),
					renderIcon(gtx, th, i["EditorMergeType"], "EditorMergeType"),
					renderIcon(gtx, th, i["EditorModeComment"], "EditorModeComment"),
					renderIcon(gtx, th, i["EditorModeEdit"], "EditorModeEdit"),
					renderIcon(gtx, th, i["EditorMonetizationOn"], "EditorMonetizationOn"),
					renderIcon(gtx, th, i["EditorMoneyOff"], "EditorMoneyOff"),
					renderIcon(gtx, th, i["EditorMultilineChart"], "EditorMultilineChart"),
					renderIcon(gtx, th, i["EditorPieChart"], "EditorPieChart"),
					renderIcon(gtx, th, i["EditorPieChartOutlined"], "EditorPieChartOutlined"),
					renderIcon(gtx, th, i["EditorPublish"], "EditorPublish"),
					renderIcon(gtx, th, i["EditorShortText"], "EditorShortText"),
					renderIcon(gtx, th, i["EditorShowChart"], "EditorShowChart"),
					renderIcon(gtx, th, i["EditorSpaceBar"], "EditorSpaceBar"),
					renderIcon(gtx, th, i["EditorStrikethroughS"], "EditorStrikethroughS"),
					renderIcon(gtx, th, i["EditorTextFields"], "EditorTextFields"),
					renderIcon(gtx, th, i["EditorTitle"], "EditorTitle"),
					renderIcon(gtx, th, i["EditorVerticalAlignBottom"], "EditorVerticalAlignBottom"),
					renderIcon(gtx, th, i["EditorVerticalAlignCenter"], "EditorVerticalAlignCenter"),
					renderIcon(gtx, th, i["EditorVerticalAlignTop"], "EditorVerticalAlignTop"),
					renderIcon(gtx, th, i["EditorWrapText"], "EditorWrapText"),
					renderIcon(gtx, th, i["FileAttachment"], "FileAttachment"),
					renderIcon(gtx, th, i["FileCloud"], "FileCloud"),
					renderIcon(gtx, th, i["FileCloudCircle"], "FileCloudCircle"),
					renderIcon(gtx, th, i["FileCloudDone"], "FileCloudDone"),
					renderIcon(gtx, th, i["FileCloudDownload"], "FileCloudDownload"),
					renderIcon(gtx, th, i["FileCloudOff"], "FileCloudOff"),
					renderIcon(gtx, th, i["FileCloudQueue"], "FileCloudQueue"),
					renderIcon(gtx, th, i["FileCloudUpload"], "FileCloudUpload"),
					renderIcon(gtx, th, i["FileCreateNewFolder"], "FileCreateNewFolder"),
					renderIcon(gtx, th, i["FileFileDownload"], "FileFileDownload"),
					renderIcon(gtx, th, i["FileFileUpload"], "FileFileUpload"),
					renderIcon(gtx, th, i["FileFolder"], "FileFolder"),
					renderIcon(gtx, th, i["FileFolderOpen"], "FileFolderOpen"),
					renderIcon(gtx, th, i["FileFolderShared"], "FileFolderShared"),
					renderIcon(gtx, th, i["HardwareCast"], "HardwareCast"),
					renderIcon(gtx, th, i["HardwareCastConnected"], "HardwareCastConnected"),
					renderIcon(gtx, th, i["HardwareComputer"], "HardwareComputer"),
					renderIcon(gtx, th, i["HardwareDesktopMac"], "HardwareDesktopMac"),
					renderIcon(gtx, th, i["HardwareDesktopWindows"], "HardwareDesktopWindows"),
					renderIcon(gtx, th, i["HardwareDeveloperBoard"], "HardwareDeveloperBoard"),
					renderIcon(gtx, th, i["HardwareDeviceHub"], "HardwareDeviceHub"),
					renderIcon(gtx, th, i["HardwareDevicesOther"], "HardwareDevicesOther"),
					renderIcon(gtx, th, i["HardwareDock"], "HardwareDock"),
					renderIcon(gtx, th, i["HardwareGamepad"], "HardwareGamepad"),
					renderIcon(gtx, th, i["HardwareHeadset"], "HardwareHeadset"),
					renderIcon(gtx, th, i["HardwareHeadsetMic"], "HardwareHeadsetMic"),
					renderIcon(gtx, th, i["HardwareKeyboard"], "HardwareKeyboard"),
					renderIcon(gtx, th, i["HardwareKeyboardArrowDown"], "HardwareKeyboardArrowDown"),
					renderIcon(gtx, th, i["HardwareKeyboardArrowLeft"], "HardwareKeyboardArrowLeft"),
					renderIcon(gtx, th, i["HardwareKeyboardArrowRight"], "HardwareKeyboardArrowRight"),
					renderIcon(gtx, th, i["HardwareKeyboardArrowUp"], "HardwareKeyboardArrowUp"),
					renderIcon(gtx, th, i["HardwareKeyboardBackspace"], "HardwareKeyboardBackspace"),
					renderIcon(gtx, th, i["HardwareKeyboardCapslock"], "HardwareKeyboardCapslock"),
					renderIcon(gtx, th, i["HardwareKeyboardHide"], "HardwareKeyboardHide"),
					renderIcon(gtx, th, i["HardwareKeyboardReturn"], "HardwareKeyboardReturn"),
					renderIcon(gtx, th, i["HardwareKeyboardTab"], "HardwareKeyboardTab"),
					renderIcon(gtx, th, i["HardwareKeyboardVoice"], "HardwareKeyboardVoice"),
					renderIcon(gtx, th, i["HardwareLaptop"], "HardwareLaptop"),
					renderIcon(gtx, th, i["HardwareLaptopChromebook"], "HardwareLaptopChromebook"),
					renderIcon(gtx, th, i["HardwareLaptopMac"], "HardwareLaptopMac"),
					renderIcon(gtx, th, i["HardwareLaptopWindows"], "HardwareLaptopWindows"),
					renderIcon(gtx, th, i["HardwareMemory"], "HardwareMemory"),
					renderIcon(gtx, th, i["HardwareMouse"], "HardwareMouse"),
					renderIcon(gtx, th, i["HardwarePhoneAndroid"], "HardwarePhoneAndroid"),
					renderIcon(gtx, th, i["HardwarePhoneIPhone"], "HardwarePhoneIPhone"),
					renderIcon(gtx, th, i["HardwarePhoneLink"], "HardwarePhoneLink"),
					renderIcon(gtx, th, i["HardwarePhoneLinkOff"], "HardwarePhoneLinkOff"),
					renderIcon(gtx, th, i["HardwarePowerInput"], "HardwarePowerInput"),
					renderIcon(gtx, th, i["HardwareRouter"], "HardwareRouter"),
					renderIcon(gtx, th, i["HardwareScanner"], "HardwareScanner"),
					renderIcon(gtx, th, i["HardwareSecurity"], "HardwareSecurity"),
					renderIcon(gtx, th, i["HardwareSIMCard"], "HardwareSIMCard"),
					renderIcon(gtx, th, i["HardwareSmartphone"], "HardwareSmartphone"),
					renderIcon(gtx, th, i["HardwareSpeaker"], "HardwareSpeaker"),
					renderIcon(gtx, th, i["HardwareSpeakerGroup"], "HardwareSpeakerGroup"),
					renderIcon(gtx, th, i["HardwareTablet"], "HardwareTablet"),
					renderIcon(gtx, th, i["HardwareTabletAndroid"], "HardwareTabletAndroid"),
					renderIcon(gtx, th, i["HardwareTabletMac"], "HardwareTabletMac"),
					renderIcon(gtx, th, i["HardwareToys"], "HardwareToys"),
					renderIcon(gtx, th, i["HardwareTV"], "HardwareTV"),
					renderIcon(gtx, th, i["HardwareVideogameAsset"], "HardwareVideogameAsset"),
					renderIcon(gtx, th, i["HardwareWatch"], "HardwareWatch"),
					renderIcon(gtx, th, i["ImageAddAPhoto"], "ImageAddAPhoto"),
					renderIcon(gtx, th, i["ImageAddToPhotos"], "ImageAddToPhotos"),
					renderIcon(gtx, th, i["ImageAdjust"], "ImageAdjust"),
					renderIcon(gtx, th, i["ImageAssistant"], "ImageAssistant"),
					renderIcon(gtx, th, i["ImageAssistantPhoto"], "ImageAssistantPhoto"),
					renderIcon(gtx, th, i["ImageAudiotrack"], "ImageAudiotrack"),
					renderIcon(gtx, th, i["ImageBlurCircular"], "ImageBlurCircular"),
					renderIcon(gtx, th, i["ImageBlurLinear"], "ImageBlurLinear"),
					renderIcon(gtx, th, i["ImageBlurOff"], "ImageBlurOff"),
					renderIcon(gtx, th, i["ImageBlurOn"], "ImageBlurOn"),
					renderIcon(gtx, th, i["ImageBrightness1"], "ImageBrightness1"),
					renderIcon(gtx, th, i["ImageBrightness2"], "ImageBrightness2"),
					renderIcon(gtx, th, i["ImageBrightness3"], "ImageBrightness3"),
					renderIcon(gtx, th, i["ImageBrightness4"], "ImageBrightness4"),
					renderIcon(gtx, th, i["ImageBrightness5"], "ImageBrightness5"),
					renderIcon(gtx, th, i["ImageBrightness6"], "ImageBrightness6"),
					renderIcon(gtx, th, i["ImageBrightness7"], "ImageBrightness7"),
					renderIcon(gtx, th, i["ImageBrokenImage"], "ImageBrokenImage"),
					renderIcon(gtx, th, i["ImageBrush"], "ImageBrush"),
					renderIcon(gtx, th, i["ImageBurstMode"], "ImageBurstMode"),
					renderIcon(gtx, th, i["ImageCamera"], "ImageCamera"),
					renderIcon(gtx, th, i["ImageCameraAlt"], "ImageCameraAlt"),
					renderIcon(gtx, th, i["ImageCameraFront"], "ImageCameraFront"),
					renderIcon(gtx, th, i["ImageCameraRear"], "ImageCameraRear"),
					renderIcon(gtx, th, i["ImageCameraRoll"], "ImageCameraRoll"),
					renderIcon(gtx, th, i["ImageCenterFocusStrong"], "ImageCenterFocusStrong"),
					renderIcon(gtx, th, i["ImageCenterFocusWeak"], "ImageCenterFocusWeak"),
					renderIcon(gtx, th, i["ImageCollections"], "ImageCollections"),
					renderIcon(gtx, th, i["ImageCollectionsBookmark"], "ImageCollectionsBookmark"),
					renderIcon(gtx, th, i["ImageColorLens"], "ImageColorLens"),
					renderIcon(gtx, th, i["ImageColorize"], "ImageColorize"),
					renderIcon(gtx, th, i["ImageCompare"], "ImageCompare"),
					renderIcon(gtx, th, i["ImageControlPoint"], "ImageControlPoint"),
					renderIcon(gtx, th, i["ImageControlPointDuplicate"], "ImageControlPointDuplicate"),
					renderIcon(gtx, th, i["ImageCrop"], "ImageCrop"),
					renderIcon(gtx, th, i["ImageCrop169"], "ImageCrop169"),
					renderIcon(gtx, th, i["ImageCrop32"], "ImageCrop32"),
					renderIcon(gtx, th, i["ImageCrop54"], "ImageCrop54"),
					renderIcon(gtx, th, i["ImageCrop75"], "ImageCrop75"),
					renderIcon(gtx, th, i["ImageCropDIN"], "ImageCropDIN"),
					renderIcon(gtx, th, i["ImageCropFree"], "ImageCropFree"),
					renderIcon(gtx, th, i["ImageCropLandscape"], "ImageCropLandscape"),
					renderIcon(gtx, th, i["ImageCropOriginal"], "ImageCropOriginal"),
					renderIcon(gtx, th, i["ImageCropPortrait"], "ImageCropPortrait"),
					renderIcon(gtx, th, i["ImageCropRotate"], "ImageCropRotate"),
					renderIcon(gtx, th, i["ImageCropSquare"], "ImageCropSquare"),
					renderIcon(gtx, th, i["ImageDehaze"], "ImageDehaze"),
					renderIcon(gtx, th, i["ImageDetails"], "ImageDetails"),
					renderIcon(gtx, th, i["ImageEdit"], "ImageEdit"),
					renderIcon(gtx, th, i["ImageExposure"], "ImageExposure"),
					renderIcon(gtx, th, i["ImageExposureNeg1"], "ImageExposureNeg1"),
					renderIcon(gtx, th, i["ImageExposureNeg2"], "ImageExposureNeg2"),
					renderIcon(gtx, th, i["ImageExposurePlus1"], "ImageExposurePlus1"),
					renderIcon(gtx, th, i["ImageExposurePlus2"], "ImageExposurePlus2"),
					renderIcon(gtx, th, i["ImageExposureZero"], "ImageExposureZero"),
					renderIcon(gtx, th, i["ImageFilter"], "ImageFilter"),
					renderIcon(gtx, th, i["ImageFilter1"], "ImageFilter1"),
					renderIcon(gtx, th, i["ImageFilter2"], "ImageFilter2"),
					renderIcon(gtx, th, i["ImageFilter3"], "ImageFilter3"),
					renderIcon(gtx, th, i["ImageFilter4"], "ImageFilter4"),
					renderIcon(gtx, th, i["ImageFilter5"], "ImageFilter5"),
					renderIcon(gtx, th, i["ImageFilter6"], "ImageFilter6"),
					renderIcon(gtx, th, i["ImageFilter7"], "ImageFilter7"),
					renderIcon(gtx, th, i["ImageFilter8"], "ImageFilter8"),
					renderIcon(gtx, th, i["ImageFilter9"], "ImageFilter9"),
					renderIcon(gtx, th, i["ImageFilter9Plus"], "ImageFilter9Plus"),
					renderIcon(gtx, th, i["ImageFilterBAndW"], "ImageFilterBAndW"),
					renderIcon(gtx, th, i["ImageFilterCenterFocus"], "ImageFilterCenterFocus"),
					renderIcon(gtx, th, i["ImageFilterDrama"], "ImageFilterDrama"),
					renderIcon(gtx, th, i["ImageFilterFrames"], "ImageFilterFrames"),
					renderIcon(gtx, th, i["ImageFilterHDR"], "ImageFilterHDR"),
					renderIcon(gtx, th, i["ImageFilterNone"], "ImageFilterNone"),
					renderIcon(gtx, th, i["ImageFilterTiltShift"], "ImageFilterTiltShift"),
					renderIcon(gtx, th, i["ImageFilterVintage"], "ImageFilterVintage"),
					renderIcon(gtx, th, i["ImageFlare"], "ImageFlare"),
					renderIcon(gtx, th, i["ImageFlashAuto"], "ImageFlashAuto"),
					renderIcon(gtx, th, i["ImageFlashOff"], "ImageFlashOff"),
					renderIcon(gtx, th, i["ImageFlashOn"], "ImageFlashOn"),
					renderIcon(gtx, th, i["ImageFlip"], "ImageFlip"),
					renderIcon(gtx, th, i["ImageGradient"], "ImageGradient"),
					renderIcon(gtx, th, i["ImageGrain"], "ImageGrain"),
					renderIcon(gtx, th, i["ImageGridOff"], "ImageGridOff"),
					renderIcon(gtx, th, i["ImageGridOn"], "ImageGridOn"),
					renderIcon(gtx, th, i["ImageHDROff"], "ImageHDROff"),
					renderIcon(gtx, th, i["ImageHDROn"], "ImageHDROn"),
					renderIcon(gtx, th, i["ImageHDRStrong"], "ImageHDRStrong"),
					renderIcon(gtx, th, i["ImageHDRWeak"], "ImageHDRWeak"),
					renderIcon(gtx, th, i["ImageHealing"], "ImageHealing"),
					renderIcon(gtx, th, i["ImageImage"], "ImageImage"),
					renderIcon(gtx, th, i["ImageImageAspectRatio"], "ImageImageAspectRatio"),
					renderIcon(gtx, th, i["ImageISO"], "ImageISO"),
					renderIcon(gtx, th, i["ImageLandscape"], "ImageLandscape"),
					renderIcon(gtx, th, i["ImageLeakAdd"], "ImageLeakAdd"),
					renderIcon(gtx, th, i["ImageLeakRemove"], "ImageLeakRemove"),
					renderIcon(gtx, th, i["ImageLens"], "ImageLens"),
					renderIcon(gtx, th, i["ImageLinkedCamera"], "ImageLinkedCamera"),
					renderIcon(gtx, th, i["ImageLooks"], "ImageLooks"),
					renderIcon(gtx, th, i["ImageLooks3"], "ImageLooks3"),
					renderIcon(gtx, th, i["ImageLooks4"], "ImageLooks4"),
					renderIcon(gtx, th, i["ImageLooks5"], "ImageLooks5"),
					renderIcon(gtx, th, i["ImageLooks6"], "ImageLooks6"),
					renderIcon(gtx, th, i["ImageLooksOne"], "ImageLooksOne"),
					renderIcon(gtx, th, i["ImageLooksTwo"], "ImageLooksTwo"),
					renderIcon(gtx, th, i["ImageLoupe"], "ImageLoupe"),
					renderIcon(gtx, th, i["ImageMonochromePhotos"], "ImageMonochromePhotos"),
					renderIcon(gtx, th, i["ImageMovieCreation"], "ImageMovieCreation"),
					renderIcon(gtx, th, i["ImageMovieFilter"], "ImageMovieFilter"),
					renderIcon(gtx, th, i["ImageMusicNote"], "ImageMusicNote"),
					renderIcon(gtx, th, i["ImageNature"], "ImageNature"),
					renderIcon(gtx, th, i["ImageNaturePeople"], "ImageNaturePeople"),
					renderIcon(gtx, th, i["ImageNavigateBefore"], "ImageNavigateBefore"),
					renderIcon(gtx, th, i["ImageNavigateNext"], "ImageNavigateNext"),
					renderIcon(gtx, th, i["ImagePalette"], "ImagePalette"),
					renderIcon(gtx, th, i["ImagePanorama"], "ImagePanorama"),
					renderIcon(gtx, th, i["ImagePanoramaFishEye"], "ImagePanoramaFishEye"),
					renderIcon(gtx, th, i["ImagePanoramaHorizontal"], "ImagePanoramaHorizontal"),
					renderIcon(gtx, th, i["ImagePanoramaVertical"], "ImagePanoramaVertical"),
					renderIcon(gtx, th, i["ImagePanoramaWideAngle"], "ImagePanoramaWideAngle"),
					renderIcon(gtx, th, i["ImagePhoto"], "ImagePhoto"),
					renderIcon(gtx, th, i["ImagePhotoAlbum"], "ImagePhotoAlbum"),
					renderIcon(gtx, th, i["ImagePhotoCamera"], "ImagePhotoCamera"),
					renderIcon(gtx, th, i["ImagePhotoFilter"], "ImagePhotoFilter"),
					renderIcon(gtx, th, i["ImagePhotoLibrary"], "ImagePhotoLibrary"),
					renderIcon(gtx, th, i["ImagePhotoSizeSelectActual"], "ImagePhotoSizeSelectActual"),
					renderIcon(gtx, th, i["ImagePhotoSizeSelectLarge"], "ImagePhotoSizeSelectLarge"),
					renderIcon(gtx, th, i["ImagePhotoSizeSelectSmall"], "ImagePhotoSizeSelectSmall"),
					renderIcon(gtx, th, i["ImagePictureAsPDF"], "ImagePictureAsPDF"),
					renderIcon(gtx, th, i["ImagePortrait"], "ImagePortrait"),
					renderIcon(gtx, th, i["ImageRemoveRedEye"], "ImageRemoveRedEye"),
					renderIcon(gtx, th, i["ImageRotate90DegreesCCW"], "ImageRotate90DegreesCCW"),
					renderIcon(gtx, th, i["ImageRotateLeft"], "ImageRotateLeft"),
					renderIcon(gtx, th, i["ImageRotateRight"], "ImageRotateRight"),
					renderIcon(gtx, th, i["ImageSlideshow"], "ImageSlideshow"),
					renderIcon(gtx, th, i["ImageStraighten"], "ImageStraighten"),
					renderIcon(gtx, th, i["ImageStyle"], "ImageStyle"),
					renderIcon(gtx, th, i["ImageSwitchCamera"], "ImageSwitchCamera"),
					renderIcon(gtx, th, i["ImageSwitchVideo"], "ImageSwitchVideo"),
					renderIcon(gtx, th, i["ImageTagFaces"], "ImageTagFaces"),
					renderIcon(gtx, th, i["ImageTexture"], "ImageTexture"),
					renderIcon(gtx, th, i["ImageTimeLapse"], "ImageTimeLapse"),
					renderIcon(gtx, th, i["ImageTimer"], "ImageTimer"),
					renderIcon(gtx, th, i["ImageTimer10"], "ImageTimer10"),
					renderIcon(gtx, th, i["ImageTimer3"], "ImageTimer3"),
					renderIcon(gtx, th, i["ImageTimerOff"], "ImageTimerOff"),
					renderIcon(gtx, th, i["ImageTonality"], "ImageTonality"),
					renderIcon(gtx, th, i["ImageTransform"], "ImageTransform"),
					renderIcon(gtx, th, i["ImageTune"], "ImageTune"),
					renderIcon(gtx, th, i["ImageViewComfy"], "ImageViewComfy"),
					renderIcon(gtx, th, i["ImageViewCompact"], "ImageViewCompact"),
					renderIcon(gtx, th, i["ImageVignette"], "ImageVignette"),
					renderIcon(gtx, th, i["ImageWBAuto"], "ImageWBAuto"),
					renderIcon(gtx, th, i["ImageWBCloudy"], "ImageWBCloudy"),
					renderIcon(gtx, th, i["ImageWBIncandescent"], "ImageWBIncandescent"),
					renderIcon(gtx, th, i["ImageWBIridescent"], "ImageWBIridescent"),
					renderIcon(gtx, th, i["ImageWBSunny"], "ImageWBSunny"),
					renderIcon(gtx, th, i["MapsAddLocation"], "MapsAddLocation"),
					renderIcon(gtx, th, i["MapsBeenhere"], "MapsBeenhere"),
					renderIcon(gtx, th, i["MapsDirections"], "MapsDirections"),
					renderIcon(gtx, th, i["MapsDirectionsBike"], "MapsDirectionsBike"),
					renderIcon(gtx, th, i["MapsDirectionsBoat"], "MapsDirectionsBoat"),
					renderIcon(gtx, th, i["MapsDirectionsBus"], "MapsDirectionsBus"),
					renderIcon(gtx, th, i["MapsDirectionsCar"], "MapsDirectionsCar"),
					renderIcon(gtx, th, i["MapsDirectionsRailway"], "MapsDirectionsRailway"),
					renderIcon(gtx, th, i["MapsDirectionsRun"], "MapsDirectionsRun"),
					renderIcon(gtx, th, i["MapsDirectionsSubway"], "MapsDirectionsSubway"),
					renderIcon(gtx, th, i["MapsDirectionsTransit"], "MapsDirectionsTransit"),
					renderIcon(gtx, th, i["MapsDirectionsWalk"], "MapsDirectionsWalk"),
					renderIcon(gtx, th, i["MapsEditLocation"], "MapsEditLocation"),
					renderIcon(gtx, th, i["MapsEVStation"], "MapsEVStation"),
					renderIcon(gtx, th, i["MapsFlight"], "MapsFlight"),
					renderIcon(gtx, th, i["MapsHotel"], "MapsHotel"),
					renderIcon(gtx, th, i["MapsLayers"], "MapsLayers"),
					renderIcon(gtx, th, i["MapsLayersClear"], "MapsLayersClear"),
					renderIcon(gtx, th, i["MapsLocalActivity"], "MapsLocalActivity"),
					renderIcon(gtx, th, i["MapsLocalAirport"], "MapsLocalAirport"),
					renderIcon(gtx, th, i["MapsLocalATM"], "MapsLocalATM"),
					renderIcon(gtx, th, i["MapsLocalBar"], "MapsLocalBar"),
					renderIcon(gtx, th, i["MapsLocalCafe"], "MapsLocalCafe"),
					renderIcon(gtx, th, i["MapsLocalCarWash"], "MapsLocalCarWash"),
					renderIcon(gtx, th, i["MapsLocalConvenienceStore"], "MapsLocalConvenienceStore"),
					renderIcon(gtx, th, i["MapsLocalDining"], "MapsLocalDining"),
					renderIcon(gtx, th, i["MapsLocalDrink"], "MapsLocalDrink"),
					renderIcon(gtx, th, i["MapsLocalFlorist"], "MapsLocalFlorist"),
					renderIcon(gtx, th, i["MapsLocalGasStation"], "MapsLocalGasStation"),
					renderIcon(gtx, th, i["MapsLocalGroceryStore"], "MapsLocalGroceryStore"),
					renderIcon(gtx, th, i["MapsLocalHospital"], "MapsLocalHospital"),
					renderIcon(gtx, th, i["MapsLocalHotel"], "MapsLocalHotel"),
					renderIcon(gtx, th, i["MapsLocalLaundryService"], "MapsLocalLaundryService"),
					renderIcon(gtx, th, i["MapsLocalLibrary"], "MapsLocalLibrary"),
					renderIcon(gtx, th, i["MapsLocalMall"], "MapsLocalMall"),
					renderIcon(gtx, th, i["MapsLocalMovies"], "MapsLocalMovies"),
					renderIcon(gtx, th, i["MapsLocalOffer"], "MapsLocalOffer"),
					renderIcon(gtx, th, i["MapsLocalParking"], "MapsLocalParking"),
					renderIcon(gtx, th, i["MapsLocalPharmacy"], "MapsLocalPharmacy"),
					renderIcon(gtx, th, i["MapsLocalPhone"], "MapsLocalPhone"),
					renderIcon(gtx, th, i["MapsLocalPizza"], "MapsLocalPizza"),
					renderIcon(gtx, th, i["MapsLocalPlay"], "MapsLocalPlay"),
					renderIcon(gtx, th, i["MapsLocalPostOffice"], "MapsLocalPostOffice"),
					renderIcon(gtx, th, i["MapsLocalPrintshop"], "MapsLocalPrintshop"),
					renderIcon(gtx, th, i["MapsLocalSee"], "MapsLocalSee"),
					renderIcon(gtx, th, i["MapsLocalShipping"], "MapsLocalShipping"),
					renderIcon(gtx, th, i["MapsLocalTaxi"], "MapsLocalTaxi"),
					renderIcon(gtx, th, i["MapsMap"], "MapsMap"),
					renderIcon(gtx, th, i["MapsMyLocation"], "MapsMyLocation"),
					renderIcon(gtx, th, i["MapsNavigation"], "MapsNavigation"),
					renderIcon(gtx, th, i["MapsNearMe"], "MapsNearMe"),
					renderIcon(gtx, th, i["MapsPersonPin"], "MapsPersonPin"),
					renderIcon(gtx, th, i["MapsPersonPinCircle"], "MapsPersonPinCircle"),
					renderIcon(gtx, th, i["MapsPinDrop"], "MapsPinDrop"),
					renderIcon(gtx, th, i["MapsPlace"], "MapsPlace"),
					renderIcon(gtx, th, i["MapsRateReview"], "MapsRateReview"),
					renderIcon(gtx, th, i["MapsRestaurant"], "MapsRestaurant"),
					renderIcon(gtx, th, i["MapsRestaurantMenu"], "MapsRestaurantMenu"),
					renderIcon(gtx, th, i["MapsSatellite"], "MapsSatellite"),
					renderIcon(gtx, th, i["MapsStoreMallDirectory"], "MapsStoreMallDirectory"),
					renderIcon(gtx, th, i["MapsStreetView"], "MapsStreetView"),
					renderIcon(gtx, th, i["MapsSubway"], "MapsSubway"),
					renderIcon(gtx, th, i["MapsTerrain"], "MapsTerrain"),
					renderIcon(gtx, th, i["MapsTraffic"], "MapsTraffic"),
					renderIcon(gtx, th, i["MapsTrain"], "MapsTrain"),
					renderIcon(gtx, th, i["MapsTram"], "MapsTram"),
					renderIcon(gtx, th, i["MapsTransferWithinAStation"], "MapsTransferWithinAStation"),
					renderIcon(gtx, th, i["MapsZoomOutMap"], "MapsZoomOutMap"),
					renderIcon(gtx, th, i["NavigationApps"], "NavigationApps"),
					renderIcon(gtx, th, i["NavigationArrowBack"], "NavigationArrowBack"),
					renderIcon(gtx, th, i["NavigationArrowDownward"], "NavigationArrowDownward"),
					renderIcon(gtx, th, i["NavigationArrowDropDown"], "NavigationArrowDropDown"),
					renderIcon(gtx, th, i["NavigationArrowDropDownCircle"], "NavigationArrowDropDownCircle"),
					renderIcon(gtx, th, i["NavigationArrowDropUp"], "NavigationArrowDropUp"),
					renderIcon(gtx, th, i["NavigationArrowForward"], "NavigationArrowForward"),
					renderIcon(gtx, th, i["NavigationArrowUpward"], "NavigationArrowUpward"),
					renderIcon(gtx, th, i["NavigationCancel"], "NavigationCancel"),
					renderIcon(gtx, th, i["NavigationCheck"], "NavigationCheck"),
					renderIcon(gtx, th, i["NavigationChevronLeft"], "NavigationChevronLeft"),
					renderIcon(gtx, th, i["NavigationChevronRight"], "NavigationChevronRight"),
					renderIcon(gtx, th, i["NavigationClose"], "NavigationClose"),
					renderIcon(gtx, th, i["NavigationExpandLess"], "NavigationExpandLess"),
					renderIcon(gtx, th, i["NavigationExpandMore"], "NavigationExpandMore"),
					renderIcon(gtx, th, i["NavigationFirstPage"], "NavigationFirstPage"),
					renderIcon(gtx, th, i["NavigationFullscreen"], "NavigationFullscreen"),
					renderIcon(gtx, th, i["NavigationFullscreenExit"], "NavigationFullscreenExit"),
					renderIcon(gtx, th, i["NavigationLastPage"], "NavigationLastPage"),
					renderIcon(gtx, th, i["NavigationMenu"], "NavigationMenu"),
					renderIcon(gtx, th, i["NavigationMoreHoriz"], "NavigationMoreHoriz"),
					renderIcon(gtx, th, i["NavigationMoreVert"], "NavigationMoreVert"),
					renderIcon(gtx, th, i["NavigationRefresh"], "NavigationRefresh"),
					renderIcon(gtx, th, i["NavigationSubdirectoryArrowLeft"], "NavigationSubdirectoryArrowLeft"),
					renderIcon(gtx, th, i["NavigationSubdirectoryArrowRight"], "NavigationSubdirectoryArrowRight"),
					renderIcon(gtx, th, i["NavigationUnfoldLess"], "NavigationUnfoldLess"),
					renderIcon(gtx, th, i["NavigationUnfoldMore"], "NavigationUnfoldMore"),
					renderIcon(gtx, th, i["NotificationADB"], "NotificationADB"),
					renderIcon(gtx, th, i["NotificationAirlineSeatFlat"], "NotificationAirlineSeatFlat"),
					renderIcon(gtx, th, i["NotificationAirlineSeatFlatAngled"], "NotificationAirlineSeatFlatAngled"),
					renderIcon(gtx, th, i["NotificationAirlineSeatIndividualSuite"], "NotificationAirlineSeatIndividualSuite"),
					renderIcon(gtx, th, i["NotificationAirlineSeatLegroomExtra"], "NotificationAirlineSeatLegroomExtra"),
					renderIcon(gtx, th, i["NotificationAirlineSeatLegroomNormal"], "NotificationAirlineSeatLegroomNormal"),
					renderIcon(gtx, th, i["NotificationAirlineSeatLegroomReduced"], "NotificationAirlineSeatLegroomReduced"),
					renderIcon(gtx, th, i["NotificationAirlineSeatReclineExtra"], "NotificationAirlineSeatReclineExtra"),
					renderIcon(gtx, th, i["NotificationAirlineSeatReclineNormal"], "NotificationAirlineSeatReclineNormal"),
					renderIcon(gtx, th, i["NotificationBluetoothAudio"], "NotificationBluetoothAudio"),
					renderIcon(gtx, th, i["NotificationConfirmationNumber"], "NotificationConfirmationNumber"),
					renderIcon(gtx, th, i["NotificationDiscFull"], "NotificationDiscFull"),
					renderIcon(gtx, th, i["NotificationDoNotDisturb"], "NotificationDoNotDisturb"),
					renderIcon(gtx, th, i["NotificationDoNotDisturbAlt"], "NotificationDoNotDisturbAlt"),
					renderIcon(gtx, th, i["NotificationDoNotDisturbOff"], "NotificationDoNotDisturbOff"),
					renderIcon(gtx, th, i["NotificationDoNotDisturbOn"], "NotificationDoNotDisturbOn"),
					renderIcon(gtx, th, i["NotificationDriveETA"], "NotificationDriveETA"),
					renderIcon(gtx, th, i["NotificationEnhancedEncryption"], "NotificationEnhancedEncryption"),
					renderIcon(gtx, th, i["NotificationEventAvailable"], "NotificationEventAvailable"),
					renderIcon(gtx, th, i["NotificationEventBusy"], "NotificationEventBusy"),
					renderIcon(gtx, th, i["NotificationEventNote"], "NotificationEventNote"),
					renderIcon(gtx, th, i["NotificationFolderSpecial"], "NotificationFolderSpecial"),
					renderIcon(gtx, th, i["NotificationLiveTV"], "NotificationLiveTV"),
					renderIcon(gtx, th, i["NotificationMMS"], "NotificationMMS"),
					renderIcon(gtx, th, i["NotificationMore"], "NotificationMore"),
					renderIcon(gtx, th, i["NotificationNetworkCheck"], "NotificationNetworkCheck"),
					renderIcon(gtx, th, i["NotificationNetworkLocked"], "NotificationNetworkLocked"),
					renderIcon(gtx, th, i["NotificationNoEncryption"], "NotificationNoEncryption"),
					renderIcon(gtx, th, i["NotificationOnDemandVideo"], "NotificationOnDemandVideo"),
					renderIcon(gtx, th, i["NotificationPersonalVideo"], "NotificationPersonalVideo"),
					renderIcon(gtx, th, i["NotificationPhoneBluetoothSpeaker"], "NotificationPhoneBluetoothSpeaker"),
					renderIcon(gtx, th, i["NotificationPhoneForwarded"], "NotificationPhoneForwarded"),
					renderIcon(gtx, th, i["NotificationPhoneInTalk"], "NotificationPhoneInTalk"),
					renderIcon(gtx, th, i["NotificationPhoneLocked"], "NotificationPhoneLocked"),
					renderIcon(gtx, th, i["NotificationPhoneMissed"], "NotificationPhoneMissed"),
					renderIcon(gtx, th, i["NotificationPhonePaused"], "NotificationPhonePaused"),
					renderIcon(gtx, th, i["NotificationPower"], "NotificationPower"),
					renderIcon(gtx, th, i["NotificationPriorityHigh"], "NotificationPriorityHigh"),
					renderIcon(gtx, th, i["NotificationRVHookup"], "NotificationRVHookup"),
					renderIcon(gtx, th, i["NotificationSDCard"], "NotificationSDCard"),
					renderIcon(gtx, th, i["NotificationSIMCardAlert"], "NotificationSIMCardAlert"),
					renderIcon(gtx, th, i["NotificationSMS"], "NotificationSMS"),
					renderIcon(gtx, th, i["NotificationSMSFailed"], "NotificationSMSFailed"),
					renderIcon(gtx, th, i["NotificationSync"], "NotificationSync"),
					renderIcon(gtx, th, i["NotificationSyncDisabled"], "NotificationSyncDisabled"),
					renderIcon(gtx, th, i["NotificationSyncProblem"], "NotificationSyncProblem"),
					renderIcon(gtx, th, i["NotificationSystemUpdate"], "NotificationSystemUpdate"),
					renderIcon(gtx, th, i["NotificationTapAndPlay"], "NotificationTapAndPlay"),
					renderIcon(gtx, th, i["NotificationTimeToLeave"], "NotificationTimeToLeave"),
					renderIcon(gtx, th, i["NotificationVibration"], "NotificationVibration"),
					renderIcon(gtx, th, i["NotificationVoiceChat"], "NotificationVoiceChat"),
					renderIcon(gtx, th, i["NotificationVPNLock"], "NotificationVPNLock"),
					renderIcon(gtx, th, i["NotificationWC"], "NotificationWC"),
					renderIcon(gtx, th, i["NotificationWiFi"], "NotificationWiFi"),
					renderIcon(gtx, th, i["PlacesACUnit"], "PlacesACUnit"),
					renderIcon(gtx, th, i["PlacesAirportShuttle"], "PlacesAirportShuttle"),
					renderIcon(gtx, th, i["PlacesAllInclusive"], "PlacesAllInclusive"),
					renderIcon(gtx, th, i["PlacesBeachAccess"], "PlacesBeachAccess"),
					renderIcon(gtx, th, i["PlacesBusinessCenter"], "PlacesBusinessCenter"),
					renderIcon(gtx, th, i["PlacesCasino"], "PlacesCasino"),
					renderIcon(gtx, th, i["PlacesChildCare"], "PlacesChildCare"),
					renderIcon(gtx, th, i["PlacesChildFriendly"], "PlacesChildFriendly"),
					renderIcon(gtx, th, i["PlacesFitnessCenter"], "PlacesFitnessCenter"),
					renderIcon(gtx, th, i["PlacesFreeBreakfast"], "PlacesFreeBreakfast"),
					renderIcon(gtx, th, i["PlacesGolfCourse"], "PlacesGolfCourse"),
					renderIcon(gtx, th, i["PlacesHotTub"], "PlacesHotTub"),
					renderIcon(gtx, th, i["PlacesKitchen"], "PlacesKitchen"),
					renderIcon(gtx, th, i["PlacesPool"], "PlacesPool"),
					renderIcon(gtx, th, i["PlacesRoomService"], "PlacesRoomService"),
					renderIcon(gtx, th, i["PlacesRVHookup"], "PlacesRVHookup"),
					renderIcon(gtx, th, i["PlacesSmokeFree"], "PlacesSmokeFree"),
					renderIcon(gtx, th, i["PlacesSmokingRooms"], "PlacesSmokingRooms"),
					renderIcon(gtx, th, i["PlacesSpa"], "PlacesSpa"),
					renderIcon(gtx, th, i["SocialCake"], "SocialCake"),
					renderIcon(gtx, th, i["SocialDomain"], "SocialDomain"),
					renderIcon(gtx, th, i["SocialGroup"], "SocialGroup"),
					renderIcon(gtx, th, i["SocialGroupAdd"], "SocialGroupAdd"),
					renderIcon(gtx, th, i["SocialLocationCity"], "SocialLocationCity"),
					renderIcon(gtx, th, i["SocialMood"], "SocialMood"),
					renderIcon(gtx, th, i["SocialMoodBad"], "SocialMoodBad"),
					renderIcon(gtx, th, i["SocialNotifications"], "SocialNotifications"),
					renderIcon(gtx, th, i["SocialNotificationsActive"], "SocialNotificationsActive"),
					renderIcon(gtx, th, i["SocialNotificationsNone"], "SocialNotificationsNone"),
					renderIcon(gtx, th, i["SocialNotificationsOff"], "SocialNotificationsOff"),
					renderIcon(gtx, th, i["SocialNotificationsPaused"], "SocialNotificationsPaused"),
					renderIcon(gtx, th, i["SocialPages"], "SocialPages"),
					renderIcon(gtx, th, i["SocialPartyMode"], "SocialPartyMode"),
					renderIcon(gtx, th, i["SocialPeople"], "SocialPeople"),
					renderIcon(gtx, th, i["SocialPeopleOutline"], "SocialPeopleOutline"),
					renderIcon(gtx, th, i["SocialPerson"], "SocialPerson"),
					renderIcon(gtx, th, i["SocialPersonAdd"], "SocialPersonAdd"),
					renderIcon(gtx, th, i["SocialPersonOutline"], "SocialPersonOutline"),
					renderIcon(gtx, th, i["SocialPlusOne"], "SocialPlusOne"),
					renderIcon(gtx, th, i["SocialPoll"], "SocialPoll"),
					renderIcon(gtx, th, i["SocialPublic"], "SocialPublic"),
					renderIcon(gtx, th, i["SocialSchool"], "SocialSchool"),
					renderIcon(gtx, th, i["SocialSentimentDissatisfied"], "SocialSentimentDissatisfied"),
					renderIcon(gtx, th, i["SocialSentimentNeutral"], "SocialSentimentNeutral"),
					renderIcon(gtx, th, i["SocialSentimentSatisfied"], "SocialSentimentSatisfied"),
					renderIcon(gtx, th, i["SocialSentimentVeryDissatisfied"], "SocialSentimentVeryDissatisfied"),
					renderIcon(gtx, th, i["SocialSentimentVerySatisfied"], "SocialSentimentVerySatisfied"),
					renderIcon(gtx, th, i["SocialShare"], "SocialShare"),
					renderIcon(gtx, th, i["SocialWhatsHot"], "SocialWhatsHot"),
					renderIcon(gtx, th, i["ToggleCheckBox"], "ToggleCheckBox"),
					renderIcon(gtx, th, i["ToggleCheckBoxOutlineBlank"], "ToggleCheckBoxOutlineBlank"),
					renderIcon(gtx, th, i["ToggleIndeterminateCheckBox"], "ToggleIndeterminateCheckBox"),
					renderIcon(gtx, th, i["ToggleRadioButtonChecked"], "ToggleRadioButtonChecked"),
					renderIcon(gtx, th, i["ToggleRadioButtonUnchecked"], "ToggleRadioButtonUnchecked"),
					renderIcon(gtx, th, i["ToggleStar"], "ToggleStar"),
					renderIcon(gtx, th, i["ToggleStarBorder"], "ToggleStarBorder"),
					renderIcon(gtx, th, i["ToggleStarHalf"], "ToggleStarHalf"),
				}
				list.Layout(gtx, len(widgets), func(i int) {
					layout.UniformInset(unit.Dp(16)).Layout(gtx, widgets[i])
				})

				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}

func renderIcon(gtx *layout.Context, th *material.Theme, icon *material.Icon, iconLabel string) func() {
	return func() {
		layout.Flex{
			Spacing: layout.SpaceBetween,
		}.Layout(gtx,
			layout.Rigid(func() {
				icon.Color = color.RGBA{A: 0xff, R: 0xcf, G: 0x30, B: 0x30}
				icon.Layout(gtx, unit.Dp(float32(32)))
			}),
			layout.Rigid(func() {
				th.H6(iconLabel).Layout(gtx)
			}),
		)
	}
}

func renderIconFlex(gtx *layout.Context, th *material.Theme, icon *material.Icon, iconLabel string) layout.FlexChild {
	return layout.Rigid(func() {
		layout.Flex{
			Spacing: layout.SpaceBetween,
		}.Layout(gtx,
			layout.Rigid(func() {
				icon.Color = color.RGBA{A: 0xff, R: 0xcf, G: 0x30, B: 0x30}
				icon.Layout(gtx, unit.Dp(float32(32)))
			}),
			layout.Rigid(func() {
				th.H6(iconLabel).Layout(gtx)
			}),
		)
	})
}

func mustIcon(ic *material.Icon, err error) *material.Icon {
	if err != nil {
		panic(err)
	}
	return ic
}
