package registry


// GetImages returns a map[string][]string which keys are the image name
// and the values are a []string with the tags of each image
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
	return imageList, nil
}

// DeleteTag deletes an image tag from the registry
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