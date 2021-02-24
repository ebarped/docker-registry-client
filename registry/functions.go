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

func (registry *Registry) DeleteTag(image, tag string, dryRunMode bool) error {
	digest, err := registry.ManifestDigest(image, tag)
	if err != nil {
		return err
	}
	if !dryRunMode {
		err = registry.DeleteManifest(image, digest)
	}
	if err != nil {
		return err
	}
	return nil
}