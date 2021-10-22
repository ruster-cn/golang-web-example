package cluster

func (service *ClusterService) DeleteClusterByID(id int) error {
	if err := service.dao.DeleteClusterByID(id); err != nil {
		return err
	}
	return nil
}
