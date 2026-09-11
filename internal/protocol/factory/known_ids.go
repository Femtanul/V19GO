package factory

// KnownClientIDs - список всех известных ID входящих пакетов (от клиента)
// Взято из PacketsHelper.py и JS-файлов проекта
var KnownClientIDs = []int32{
	// Auth
	10100, // ClientHelloMessage
	10101, // LoginMessage
	10107, // ClientCapabilitiesMessage
	10108, // KeepAliveMessage
	10110, // AnalyticEventMessage
	10177, // ClientInfoMessage

	// Home
	10212, // ChangeAvatarNameMessage
	14101, // GoHomeMessage
	14102, // EndClientTurnMessage
	14103, // StartGameMessage
	14106, // CancelMatchmakingMessage
	14109, // GoHomeFromOfflinePractiseMessage
	14110, // AskForBattleEndMessage
	14113, // GetPlayerProfileMessage

	// Battle
	10555, // ClientInputMessage

	// Team
	14350, // TeamCreateMessage
	14351, // TeamJoinMessage
	14353, // TeamLeaveMessage
	14366, // PlayerStatusMessage
}

// KnownServerIDs - список всех известных ID исходящих пакетов (от сервера)
var KnownServerIDs = []int32{
	// Auth
	20100, // ServerHelloMessage
	20103, // LoginFailedMessage
	20104, // LoginOKMessage

	// Battle
	20405, // MatchMakingStatusMessage
	20406, // MatchMakingCancelledMessage
	20559, // StartLoadingMessage
	23456, // BattleEndMessage
	24109, // VisionUpdateMessage
	24112, // UDPConnectionInfo

	// Home
	24101, // OwnHomeDataMessage
	24111, // AvailableServerCommandMessage (deliveryItems, LogicBoxCommand, SetNameCallback)
	24113, // PlayerProfileMessage
	24123, // SeasonRewardsMessage

	// Team
	24124, // TeamMessage
}