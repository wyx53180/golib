	l := golog.NewLogger("debug", "my.log")
	l.Info("info")
	l.Debug("debug")
	l.Error("adf", 123, "eror")
  
  
/*  
2022-06-06 10:32:33 [INFO][test.go, main.logTest, 94]:[info]
2022-06-06 10:32:33 [DEBUG][test.go, main.logTest, 95]:[debug]
2022-06-06 10:32:33 [ERROR][test.go, main.logTest, 96]:[adf 123 eror]
*/
