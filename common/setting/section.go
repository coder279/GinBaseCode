package setting

var sections = make(map[string]interface{})

/**
 * @summary 根据传递过来的字符串赋予配置属性
 * @param k string
 * @param v interface
 * @return error
 */
func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.viper.UnmarshalKey(k,v)
	if err != nil {
		return err
	}
	if _,ok := sections[k];!ok {
		sections[k] = v
	}
	return nil
}
/**
 * @summary 读取所有配置
 * @return error
 */
func (s *Setting) ReadAllSection() error {
	for k,v := range sections {
		err := s.ReadSection(k,v)
		if err != nil {
			return err
		}
	}
	return nil
}


