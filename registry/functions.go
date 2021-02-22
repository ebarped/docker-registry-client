package registry

func (registry *Registry) GetImages() (map[string][]string, error) {
	repositories, err := registry.Repositories()
	if err != nil {
		return nil, err
	}
	imageList := make(map[string][]string)
	for _, image := range repositories {
		tags, err := registry.Tags(image)
		if err != nil {
			return nil, err
		}
		imageList[image] = tags
	}
	return imageList, err
}