package testdata

// Sun events for Los Angeles assuming observer in sea level, ignoring DST.
// Taken from AccurateTimes.
var SunLosAngeles = TestData{
	Name:     "Los Angeles",
	Z:        PST,
	Location: losAngeles,
	Times: []TestTime{
		{"01/01/2010", "06:59:49", "11:57:33", "16:55:26"},
		{"02/01/2010", "06:59:59", "11:58:01", "16:56:12"},
		{"03/01/2010", "07:00:08", "11:58:29", "16:56:59"},
		{"04/01/2010", "07:00:15", "11:58:56", "16:57:47"},
		{"05/01/2010", "07:00:20", "11:59:23", "16:58:36"},
		{"06/01/2010", "07:00:23", "11:59:49", "16:59:26"},
		{"07/01/2010", "07:00:24", "12:00:15", "17:00:17"},
		{"08/01/2010", "07:00:24", "12:00:40", "17:01:09"},
		{"09/01/2010", "07:00:21", "12:01:05", "17:02:01"},
		{"10/01/2010", "07:00:17", "12:01:29", "17:02:55"},
		{"11/01/2010", "07:00:11", "12:01:53", "17:03:49"},
		{"12/01/2010", "07:00:03", "12:02:16", "17:04:44"},
		{"13/01/2010", "06:59:53", "12:02:39", "17:05:39"},
		{"14/01/2010", "06:59:42", "12:03:01", "17:06:35"},
		{"15/01/2010", "06:59:28", "12:03:22", "17:07:32"},
		{"16/01/2010", "06:59:13", "12:03:43", "17:08:29"},
		{"17/01/2010", "06:58:56", "12:04:03", "17:09:27"},
		{"18/01/2010", "06:58:36", "12:04:22", "17:10:24"},
		{"19/01/2010", "06:58:16", "12:04:40", "17:11:23"},
		{"20/01/2010", "06:57:53", "12:04:58", "17:12:21"},
		{"21/01/2010", "06:57:28", "12:05:15", "17:13:20"},
		{"22/01/2010", "06:57:02", "12:05:31", "17:14:19"},
		{"23/01/2010", "06:56:34", "12:05:46", "17:15:18"},
		{"24/01/2010", "06:56:04", "12:06:01", "17:16:17"},
		{"25/01/2010", "06:55:33", "12:06:15", "17:17:17"},
		{"26/01/2010", "06:55:00", "12:06:27", "17:18:16"},
		{"27/01/2010", "06:54:25", "12:06:40", "17:19:16"},
		{"28/01/2010", "06:53:48", "12:06:51", "17:20:15"},
		{"29/01/2010", "06:53:10", "12:07:01", "17:21:15"},
		{"30/01/2010", "06:52:30", "12:07:11", "17:22:14"},
		{"31/01/2010", "06:51:49", "12:07:20", "17:23:13"},
		{"01/02/2010", "06:51:06", "12:07:28", "17:24:13"},
		{"02/02/2010", "06:50:22", "12:07:35", "17:25:12"},
		{"03/02/2010", "06:49:36", "12:07:41", "17:26:11"},
		{"04/02/2010", "06:48:48", "12:07:47", "17:27:10"},
		{"05/02/2010", "06:48:00", "12:07:52", "17:28:09"},
		{"06/02/2010", "06:47:09", "12:07:56", "17:29:07"},
		{"07/02/2010", "06:46:18", "12:07:59", "17:30:06"},
		{"08/02/2010", "06:45:25", "12:08:01", "17:31:04"},
		{"09/02/2010", "06:44:31", "12:08:03", "17:32:02"},
		{"10/02/2010", "06:43:35", "12:08:04", "17:32:59"},
		{"11/02/2010", "06:42:38", "12:08:04", "17:33:57"},
		{"12/02/2010", "06:41:40", "12:08:04", "17:34:54"},
		{"13/02/2010", "06:40:41", "12:08:02", "17:35:51"},
		{"14/02/2010", "06:39:41", "12:08:00", "17:36:47"},
		{"15/02/2010", "06:38:39", "12:07:57", "17:37:43"},
		{"16/02/2010", "06:37:37", "12:07:54", "17:38:39"},
		{"17/02/2010", "06:36:33", "12:07:50", "17:39:35"},
		{"18/02/2010", "06:35:28", "12:07:45", "17:40:30"},
		{"19/02/2010", "06:34:23", "12:07:39", "17:41:24"},
		{"20/02/2010", "06:33:16", "12:07:33", "17:42:19"},
		{"21/02/2010", "06:32:08", "12:07:26", "17:43:13"},
		{"22/02/2010", "06:30:59", "12:07:18", "17:44:07"},
		{"23/02/2010", "06:29:50", "12:07:10", "17:45:00"},
		{"24/02/2010", "06:28:39", "12:07:01", "17:45:53"},
		{"25/02/2010", "06:27:28", "12:06:52", "17:46:45"},
		{"26/02/2010", "06:26:16", "12:06:42", "17:47:37"},
		{"27/02/2010", "06:25:03", "12:06:31", "17:48:29"},
		{"28/02/2010", "06:23:49", "12:06:20", "17:49:21"},
		{"01/03/2010", "06:22:35", "12:06:08", "17:50:12"},
		{"02/03/2010", "06:21:20", "12:05:56", "17:51:03"},
		{"03/03/2010", "06:20:04", "12:05:43", "17:51:53"},
		{"04/03/2010", "06:18:48", "12:05:30", "17:52:44"},
		{"05/03/2010", "06:17:31", "12:05:17", "17:53:34"},
		{"06/03/2010", "06:16:14", "12:05:03", "17:54:23"},
		{"07/03/2010", "06:14:56", "12:04:48", "17:55:13"},
		{"08/03/2010", "06:13:37", "12:04:34", "17:56:02"},
		{"09/03/2010", "06:12:18", "12:04:19", "17:56:51"},
		{"10/03/2010", "06:10:59", "12:04:03", "17:57:40"},
		{"11/03/2010", "06:09:40", "12:03:48", "17:58:28"},
		{"12/03/2010", "06:08:19", "12:03:32", "17:59:16"},
		{"13/03/2010", "06:06:59", "12:03:16", "18:00:04"},
		{"14/03/2010", "06:05:38", "12:02:59", "18:00:52"},
		{"15/03/2010", "06:04:17", "12:02:42", "18:01:40"},
		{"16/03/2010", "06:02:56", "12:02:25", "18:02:27"},
		{"17/03/2010", "06:01:35", "12:02:08", "18:03:14"},
		{"18/03/2010", "06:00:13", "12:01:51", "18:04:01"},
		{"19/03/2010", "05:58:51", "12:01:33", "18:04:48"},
		{"20/03/2010", "05:57:29", "12:01:16", "18:05:35"},
		{"21/03/2010", "05:56:07", "12:00:58", "18:06:22"},
		{"22/03/2010", "05:54:45", "12:00:40", "18:07:08"},
		{"23/03/2010", "05:53:23", "12:00:22", "18:07:54"},
		{"24/03/2010", "05:52:01", "12:00:04", "18:08:40"},
		{"25/03/2010", "05:50:39", "11:59:46", "18:09:26"},
		{"26/03/2010", "05:49:16", "11:59:27", "18:10:12"},
		{"27/03/2010", "05:47:54", "11:59:09", "18:10:58"},
		{"28/03/2010", "05:46:32", "11:58:51", "18:11:44"},
		{"29/03/2010", "05:45:10", "11:58:33", "18:12:30"},
		{"30/03/2010", "05:43:48", "11:58:15", "18:13:15"},
		{"31/03/2010", "05:42:27", "11:57:57", "18:14:01"},
		{"01/04/2010", "05:41:05", "11:57:39", "18:14:46"},
		{"02/04/2010", "05:39:44", "11:57:21", "18:15:32"},
		{"03/04/2010", "05:38:23", "11:57:03", "18:16:18"},
		{"04/04/2010", "05:37:03", "11:56:46", "18:17:03"},
		{"05/04/2010", "05:35:42", "11:56:29", "18:17:49"},
		{"06/04/2010", "05:34:22", "11:56:12", "18:18:35"},
		{"07/04/2010", "05:33:03", "11:55:55", "18:19:21"},
		{"08/04/2010", "05:31:44", "11:55:38", "18:20:07"},
		{"09/04/2010", "05:30:25", "11:55:22", "18:20:52"},
		{"10/04/2010", "05:29:07", "11:55:06", "18:21:38"},
		{"11/04/2010", "05:27:49", "11:54:50", "18:22:24"},
		{"12/04/2010", "05:26:32", "11:54:35", "18:23:11"},
		{"13/04/2010", "05:25:16", "11:54:19", "18:23:57"},
		{"14/04/2010", "05:24:00", "11:54:05", "18:24:43"},
		{"15/04/2010", "05:22:44", "11:53:50", "18:25:29"},
		{"16/04/2010", "05:21:30", "11:53:36", "18:26:16"},
		{"17/04/2010", "05:20:16", "11:53:22", "18:27:02"},
		{"18/04/2010", "05:19:02", "11:53:09", "18:27:49"},
		{"19/04/2010", "05:17:50", "11:52:56", "18:28:35"},
		{"20/04/2010", "05:16:38", "11:52:43", "18:29:22"},
		{"21/04/2010", "05:15:27", "11:52:31", "18:30:09"},
		{"22/04/2010", "05:14:16", "11:52:20", "18:30:55"},
		{"23/04/2010", "05:13:07", "11:52:08", "18:31:42"},
		{"24/04/2010", "05:11:58", "11:51:57", "18:32:29"},
		{"25/04/2010", "05:10:50", "11:51:47", "18:33:16"},
		{"26/04/2010", "05:09:44", "11:51:37", "18:34:03"},
		{"27/04/2010", "05:08:38", "11:51:28", "18:34:50"},
		{"28/04/2010", "05:07:33", "11:51:19", "18:35:36"},
		{"29/04/2010", "05:06:29", "11:51:10", "18:36:23"},
		{"30/04/2010", "05:05:26", "11:51:02", "18:37:10"},
		{"01/05/2010", "05:04:24", "11:50:55", "18:37:57"},
		{"02/05/2010", "05:03:23", "11:50:48", "18:38:44"},
		{"03/05/2010", "05:02:23", "11:50:42", "18:39:31"},
		{"04/05/2010", "05:01:25", "11:50:36", "18:40:18"},
		{"05/05/2010", "05:00:27", "11:50:31", "18:41:05"},
		{"06/05/2010", "04:59:31", "11:50:27", "18:41:52"},
		{"07/05/2010", "04:58:36", "11:50:23", "18:42:39"},
		{"08/05/2010", "04:57:43", "11:50:19", "18:43:25"},
		{"09/05/2010", "04:56:50", "11:50:17", "18:44:12"},
		{"10/05/2010", "04:55:59", "11:50:15", "18:44:59"},
		{"11/05/2010", "04:55:09", "11:50:13", "18:45:45"},
		{"12/05/2010", "04:54:21", "11:50:12", "18:46:31"},
		{"13/05/2010", "04:53:33", "11:50:12", "18:47:17"},
		{"14/05/2010", "04:52:48", "11:50:12", "18:48:02"},
		{"15/05/2010", "04:52:03", "11:50:12", "18:48:48"},
		{"16/05/2010", "04:51:20", "11:50:14", "18:49:33"},
		{"17/05/2010", "04:50:39", "11:50:15", "18:50:18"},
		{"18/05/2010", "04:49:59", "11:50:18", "18:51:02"},
		{"19/05/2010", "04:49:20", "11:50:21", "18:51:46"},
		{"20/05/2010", "04:48:43", "11:50:24", "18:52:29"},
		{"21/05/2010", "04:48:07", "11:50:28", "18:53:13"},
		{"22/05/2010", "04:47:33", "11:50:33", "18:53:55"},
		{"23/05/2010", "04:47:00", "11:50:37", "18:54:37"},
		{"24/05/2010", "04:46:29", "11:50:43", "18:55:19"},
		{"25/05/2010", "04:45:59", "11:50:49", "18:56:00"},
		{"26/05/2010", "04:45:31", "11:50:55", "18:56:40"},
		{"27/05/2010", "04:45:04", "11:51:02", "18:57:20"},
		{"28/05/2010", "04:44:39", "11:51:09", "18:57:59"},
		{"29/05/2010", "04:44:16", "11:51:17", "18:58:37"},
		{"30/05/2010", "04:43:54", "11:51:25", "18:59:14"},
		{"31/05/2010", "04:43:34", "11:51:34", "18:59:51"},
		{"01/06/2010", "04:43:15", "11:51:43", "19:00:27"},
		{"02/06/2010", "04:42:58", "11:51:53", "19:01:02"},
		{"03/06/2010", "04:42:43", "11:52:02", "19:01:37"},
		{"04/06/2010", "04:42:29", "11:52:13", "19:02:10"},
		{"05/06/2010", "04:42:17", "11:52:23", "19:02:43"},
		{"06/06/2010", "04:42:07", "11:52:34", "19:03:14"},
		{"07/06/2010", "04:41:58", "11:52:45", "19:03:45"},
		{"08/06/2010", "04:41:51", "11:52:57", "19:04:14"},
		{"09/06/2010", "04:41:45", "11:53:09", "19:04:43"},
		{"10/06/2010", "04:41:41", "11:53:21", "19:05:10"},
		{"11/06/2010", "04:41:39", "11:53:33", "19:05:36"},
		{"12/06/2010", "04:41:38", "11:53:46", "19:06:01"},
		{"13/06/2010", "04:41:39", "11:53:58", "19:06:25"},
		{"14/06/2010", "04:41:41", "11:54:11", "19:06:47"},
		{"15/06/2010", "04:41:45", "11:54:24", "19:07:09"},
		{"16/06/2010", "04:41:50", "11:54:37", "19:07:28"},
		{"17/06/2010", "04:41:57", "11:54:50", "19:07:47"},
		{"18/06/2010", "04:42:05", "11:55:03", "19:08:04"},
		{"19/06/2010", "04:42:15", "11:55:16", "19:08:20"},
		{"20/06/2010", "04:42:26", "11:55:29", "19:08:34"},
		{"21/06/2010", "04:42:38", "11:55:42", "19:08:47"},
		{"22/06/2010", "04:42:52", "11:55:55", "19:08:58"},
		{"23/06/2010", "04:43:07", "11:56:08", "19:09:08"},
		{"24/06/2010", "04:43:23", "11:56:21", "19:09:16"},
		{"25/06/2010", "04:43:41", "11:56:33", "19:09:23"},
		{"26/06/2010", "04:44:00", "11:56:46", "19:09:28"},
		{"27/06/2010", "04:44:20", "11:56:58", "19:09:32"},
		{"28/06/2010", "04:44:42", "11:57:11", "19:09:34"},
		{"29/06/2010", "04:45:04", "11:57:23", "19:09:34"},
		{"30/06/2010", "04:45:28", "11:57:34", "19:09:33"},
		{"01/07/2010", "04:45:53", "11:57:46", "19:09:31"},
		{"02/07/2010", "04:46:19", "11:57:57", "19:09:26"},
		{"03/07/2010", "04:46:47", "11:58:08", "19:09:20"},
		{"04/07/2010", "04:47:15", "11:58:19", "19:09:13"},
		{"05/07/2010", "04:47:44", "11:58:30", "19:09:04"},
		{"06/07/2010", "04:48:14", "11:58:40", "19:08:53"},
		{"07/07/2010", "04:48:45", "11:58:49", "19:08:41"},
		{"08/07/2010", "04:49:18", "11:58:59", "19:08:26"},
		{"09/07/2010", "04:49:50", "11:59:08", "19:08:11"},
		{"10/07/2010", "04:50:24", "11:59:16", "19:07:53"},
		{"11/07/2010", "04:50:59", "11:59:24", "19:07:34"},
		{"12/07/2010", "04:51:34", "11:59:32", "19:07:14"},
		{"13/07/2010", "04:52:10", "11:59:39", "19:06:52"},
		{"14/07/2010", "04:52:47", "11:59:46", "19:06:28"},
		{"15/07/2010", "04:53:24", "11:59:52", "19:06:02"},
		{"16/07/2010", "04:54:02", "11:59:58", "19:05:35"},
		{"17/07/2010", "04:54:41", "12:00:03", "19:05:06"},
		{"18/07/2010", "04:55:20", "12:00:08", "19:04:36"},
		{"19/07/2010", "04:55:59", "12:00:12", "19:04:04"},
		{"20/07/2010", "04:56:39", "12:00:15", "19:03:30"},
		{"21/07/2010", "04:57:19", "12:00:18", "19:02:55"},
		{"22/07/2010", "04:58:00", "12:00:20", "19:02:18"},
		{"23/07/2010", "04:58:41", "12:00:22", "19:01:40"},
		{"24/07/2010", "04:59:23", "12:00:23", "19:01:00"},
		{"25/07/2010", "05:00:04", "12:00:24", "19:00:19"},
		{"26/07/2010", "05:00:46", "12:00:24", "18:59:36"},
		{"27/07/2010", "05:01:29", "12:00:23", "18:58:52"},
		{"28/07/2010", "05:02:11", "12:00:22", "18:58:07"},
		{"29/07/2010", "05:02:54", "12:00:20", "18:57:20"},
		{"30/07/2010", "05:03:37", "12:00:17", "18:56:31"},
		{"31/07/2010", "05:04:20", "12:00:14", "18:55:42"},
		{"01/08/2010", "05:05:04", "12:00:10", "18:54:50"},
		{"02/08/2010", "05:05:47", "12:00:06", "18:53:58"},
		{"03/08/2010", "05:06:31", "12:00:01", "18:53:04"},
		{"04/08/2010", "05:07:15", "11:59:56", "18:52:10"},
		{"05/08/2010", "05:07:58", "11:59:50", "18:51:13"},
		{"06/08/2010", "05:08:42", "11:59:43", "18:50:16"},
		{"07/08/2010", "05:09:26", "11:59:36", "18:49:17"},
		{"08/08/2010", "05:10:10", "11:59:28", "18:48:18"},
		{"09/08/2010", "05:10:54", "11:59:20", "18:47:17"},
		{"10/08/2010", "05:11:38", "11:59:11", "18:46:15"},
		{"11/08/2010", "05:12:22", "11:59:02", "18:45:12"},
		{"12/08/2010", "05:13:06", "11:58:52", "18:44:07"},
		{"13/08/2010", "05:13:50", "11:58:41", "18:43:02"},
		{"14/08/2010", "05:14:34", "11:58:30", "18:41:56"},
		{"15/08/2010", "05:15:17", "11:58:18", "18:40:49"},
		{"16/08/2010", "05:16:01", "11:58:06", "18:39:40"},
		{"17/08/2010", "05:16:44", "11:57:53", "18:38:31"},
		{"18/08/2010", "05:17:28", "11:57:40", "18:37:21"},
		{"19/08/2010", "05:18:11", "11:57:26", "18:36:10"},
		{"20/08/2010", "05:18:54", "11:57:12", "18:34:58"},
		{"21/08/2010", "05:19:37", "11:56:57", "18:33:45"},
		{"22/08/2010", "05:20:20", "11:56:42", "18:32:32"},
		{"23/08/2010", "05:21:03", "11:56:26", "18:31:17"},
		{"24/08/2010", "05:21:45", "11:56:10", "18:30:02"},
		{"25/08/2010", "05:22:28", "11:55:53", "18:28:47"},
		{"26/08/2010", "05:23:10", "11:55:36", "18:27:30"},
		{"27/08/2010", "05:23:53", "11:55:19", "18:26:13"},
		{"28/08/2010", "05:24:35", "11:55:01", "18:24:55"},
		{"29/08/2010", "05:25:17", "11:54:43", "18:23:37"},
		{"30/08/2010", "05:25:59", "11:54:25", "18:22:18"},
		{"31/08/2010", "05:26:41", "11:54:06", "18:20:58"},
		{"01/09/2010", "05:27:23", "11:53:47", "18:19:38"},
		{"02/09/2010", "05:28:05", "11:53:28", "18:18:18"},
		{"03/09/2010", "05:28:47", "11:53:08", "18:16:57"},
		{"04/09/2010", "05:29:29", "11:52:48", "18:15:36"},
		{"05/09/2010", "05:30:11", "11:52:28", "18:14:14"},
		{"06/09/2010", "05:30:52", "11:52:08", "18:12:52"},
		{"07/09/2010", "05:31:34", "11:51:48", "18:11:29"},
		{"08/09/2010", "05:32:16", "11:51:27", "18:10:07"},
		{"09/09/2010", "05:32:57", "11:51:07", "18:08:43"},
		{"10/09/2010", "05:33:39", "11:50:46", "18:07:20"},
		{"11/09/2010", "05:34:21", "11:50:25", "18:05:56"},
		{"12/09/2010", "05:35:02", "11:50:04", "18:04:33"},
		{"13/09/2010", "05:35:44", "11:49:42", "18:03:09"},
		{"14/09/2010", "05:36:26", "11:49:21", "18:01:44"},
		{"15/09/2010", "05:37:07", "11:49:00", "18:00:20"},
		{"16/09/2010", "05:37:49", "11:48:38", "17:58:56"},
		{"17/09/2010", "05:38:30", "11:48:17", "17:57:31"},
		{"18/09/2010", "05:39:12", "11:47:55", "17:56:07"},
		{"19/09/2010", "05:39:54", "11:47:34", "17:54:42"},
		{"20/09/2010", "05:40:36", "11:47:12", "17:53:17"},
		{"21/09/2010", "05:41:18", "11:46:51", "17:51:53"},
		{"22/09/2010", "05:42:00", "11:46:30", "17:50:28"},
		{"23/09/2010", "05:42:42", "11:46:09", "17:49:04"},
		{"24/09/2010", "05:43:25", "11:45:48", "17:47:40"},
		{"25/09/2010", "05:44:07", "11:45:27", "17:46:16"},
		{"26/09/2010", "05:44:50", "11:45:06", "17:44:52"},
		{"27/09/2010", "05:45:33", "11:44:46", "17:43:28"},
		{"28/09/2010", "05:46:16", "11:44:26", "17:42:05"},
		{"29/09/2010", "05:46:59", "11:44:06", "17:40:41"},
		{"30/09/2010", "05:47:43", "11:43:46", "17:39:19"},
		{"01/10/2010", "05:48:27", "11:43:27", "17:37:56"},
		{"02/10/2010", "05:49:11", "11:43:08", "17:36:34"},
		{"03/10/2010", "05:49:55", "11:42:49", "17:35:12"},
		{"04/10/2010", "05:50:40", "11:42:30", "17:33:51"},
		{"05/10/2010", "05:51:25", "11:42:12", "17:32:30"},
		{"06/10/2010", "05:52:10", "11:41:55", "17:31:10"},
		{"07/10/2010", "05:52:56", "11:41:38", "17:29:50"},
		{"08/10/2010", "05:53:41", "11:41:21", "17:28:31"},
		{"09/10/2010", "05:54:28", "11:41:04", "17:27:12"},
		{"10/10/2010", "05:55:14", "11:40:49", "17:25:54"},
		{"11/10/2010", "05:56:01", "11:40:33", "17:24:36"},
		{"12/10/2010", "05:56:48", "11:40:18", "17:23:19"},
		{"13/10/2010", "05:57:35", "11:40:04", "17:22:03"},
		{"14/10/2010", "05:58:23", "11:39:50", "17:20:48"},
		{"15/10/2010", "05:59:11", "11:39:36", "17:19:33"},
		{"16/10/2010", "06:00:00", "11:39:23", "17:18:19"},
		{"17/10/2010", "06:00:48", "11:39:11", "17:17:06"},
		{"18/10/2010", "06:01:37", "11:38:59", "17:15:53"},
		{"19/10/2010", "06:02:27", "11:38:48", "17:14:42"},
		{"20/10/2010", "06:03:17", "11:38:38", "17:13:31"},
		{"21/10/2010", "06:04:07", "11:38:28", "17:12:21"},
		{"22/10/2010", "06:04:57", "11:38:19", "17:11:13"},
		{"23/10/2010", "06:05:48", "11:38:10", "17:10:05"},
		{"24/10/2010", "06:06:40", "11:38:02", "17:08:58"},
		{"25/10/2010", "06:07:31", "11:37:55", "17:07:53"},
		{"26/10/2010", "06:08:23", "11:37:49", "17:06:48"},
		{"27/10/2010", "06:09:16", "11:37:43", "17:05:45"},
		{"28/10/2010", "06:10:09", "11:37:38", "17:04:42"},
		{"29/10/2010", "06:11:02", "11:37:34", "17:03:41"},
		{"30/10/2010", "06:11:55", "11:37:31", "17:02:41"},
		{"31/10/2010", "06:12:49", "11:37:28", "17:01:43"},
		{"01/11/2010", "06:13:43", "11:37:27", "17:00:45"},
		{"02/11/2010", "06:14:38", "11:37:26", "16:59:49"},
		{"03/11/2010", "06:15:33", "11:37:26", "16:58:55"},
		{"04/11/2010", "06:16:28", "11:37:27", "16:58:02"},
		{"05/11/2010", "06:17:24", "11:37:28", "16:57:10"},
		{"06/11/2010", "06:18:19", "11:37:31", "16:56:19"},
		{"07/11/2010", "06:19:15", "11:37:34", "16:55:30"},
		{"08/11/2010", "06:20:11", "11:37:38", "16:54:43"},
		{"09/11/2010", "06:21:08", "11:37:43", "16:53:57"},
		{"10/11/2010", "06:22:04", "11:37:49", "16:53:12"},
		{"11/11/2010", "06:23:01", "11:37:56", "16:52:29"},
		{"12/11/2010", "06:23:57", "11:38:03", "16:51:48"},
		{"13/11/2010", "06:24:54", "11:38:11", "16:51:08"},
		{"14/11/2010", "06:25:51", "11:38:20", "16:50:30"},
		{"15/11/2010", "06:26:48", "11:38:30", "16:49:54"},
		{"16/11/2010", "06:27:45", "11:38:41", "16:49:19"},
		{"17/11/2010", "06:28:41", "11:38:53", "16:48:46"},
		{"18/11/2010", "06:29:38", "11:39:05", "16:48:14"},
		{"19/11/2010", "06:30:35", "11:39:19", "16:47:45"},
		{"20/11/2010", "06:31:31", "11:39:33", "16:47:17"},
		{"21/11/2010", "06:32:27", "11:39:47", "16:46:51"},
		{"22/11/2010", "06:33:24", "11:40:03", "16:46:26"},
		{"23/11/2010", "06:34:19", "11:40:20", "16:46:04"},
		{"24/11/2010", "06:35:15", "11:40:37", "16:45:43"},
		{"25/11/2010", "06:36:10", "11:40:55", "16:45:24"},
		{"26/11/2010", "06:37:05", "11:41:14", "16:45:07"},
		{"27/11/2010", "06:38:00", "11:41:33", "16:44:52"},
		{"28/11/2010", "06:38:54", "11:41:53", "16:44:39"},
		{"29/11/2010", "06:39:48", "11:42:14", "16:44:28"},
		{"30/11/2010", "06:40:41", "11:42:36", "16:44:19"},
		{"01/12/2010", "06:41:33", "11:42:58", "16:44:12"},
		{"02/12/2010", "06:42:25", "11:43:21", "16:44:06"},
		{"03/12/2010", "06:43:17", "11:43:45", "16:44:03"},
		{"04/12/2010", "06:44:07", "11:44:09", "16:44:01"},
		{"05/12/2010", "06:44:57", "11:44:34", "16:44:02"},
		{"06/12/2010", "06:45:46", "11:44:59", "16:44:04"},
		{"07/12/2010", "06:46:34", "11:45:25", "16:44:08"},
		{"08/12/2010", "06:47:22", "11:45:52", "16:44:14"},
		{"09/12/2010", "06:48:08", "11:46:18", "16:44:22"},
		{"10/12/2010", "06:48:53", "11:46:46", "16:44:32"},
		{"11/12/2010", "06:49:37", "11:47:13", "16:44:43"},
		{"12/12/2010", "06:50:20", "11:47:41", "16:44:57"},
		{"13/12/2010", "06:51:02", "11:48:09", "16:45:12"},
		{"14/12/2010", "06:51:43", "11:48:38", "16:45:29"},
		{"15/12/2010", "06:52:22", "11:49:07", "16:45:48"},
		{"16/12/2010", "06:53:00", "11:49:36", "16:46:09"},
		{"17/12/2010", "06:53:37", "11:50:05", "16:46:31"},
		{"18/12/2010", "06:54:12", "11:50:34", "16:46:55"},
		{"19/12/2010", "06:54:46", "11:51:04", "16:47:21"},
		{"20/12/2010", "06:55:19", "11:51:34", "16:47:48"},
		{"21/12/2010", "06:55:50", "11:52:03", "16:48:17"},
		{"22/12/2010", "06:56:19", "11:52:33", "16:48:48"},
		{"23/12/2010", "06:56:47", "11:53:03", "16:49:20"},
		{"24/12/2010", "06:57:14", "11:53:33", "16:49:54"},
		{"25/12/2010", "06:57:39", "11:54:02", "16:50:29"},
		{"26/12/2010", "06:58:02", "11:54:32", "16:51:06"},
		{"27/12/2010", "06:58:23", "11:55:01", "16:51:44"},
		{"28/12/2010", "06:58:43", "11:55:31", "16:52:23"},
		{"29/12/2010", "06:59:01", "11:56:00", "16:53:04"},
		{"30/12/2010", "06:59:18", "11:56:29", "16:53:47"},
		{"31/12/2010", "06:59:32", "11:56:58", "16:54:30"},
	},
}