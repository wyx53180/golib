	data := `{"code":200,"data":{"ext":{"utm_medium":"distribute.pc_search_hot_word.none-task-hot_word-alirecmd-1.nonecase"},"size":1,"items":[{"ext":null,"resourceId":"","mediaAssetInfo":null,"productId":"golang make","reportData":{"eventClick":true,"data":{"mod":"popu_895","extra":"{\"utm_medium\":\"distribute.pc_search_hot_word.none-task-hot_word-alirecmd-1.nonecase\",\"dist_request_id\":\"1646296542475_57712\",\"hotword\":\"golang make\"}","dist_request_id":"1646296542475_57712","ab_strategy":"default","index":"1","strategy":"alirecmd"},"urlParams":{"utm_medium":"distribute.pc_search_hot_word.none-task-hot_word-alirecmd-1.nonecase","depth_1-utm_source":"distribute.pc_search_hot_word.none-task-hot_word-alirecmd-1.nonecase"},"eventView":true},"recommendType":"ali","index":1,"style":"word_1","strategyId":"alirecmd","productType":"hot_word"}]},"message":"success"}`

# Loads
	d := gojson.Loads(&data)
	m := d.Get("data", "items").Index(0).Get("recommendType")
	fmt.Println(m.Data)
	fmt.Println(m.Default(1))

	

# Dumps
	m_data := map[int]int{1: 1, 2: 211}
	data = gojson.Dumps(m_data)
	fmt.Println(data)
