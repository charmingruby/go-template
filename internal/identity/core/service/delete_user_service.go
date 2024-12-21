package service

type DeleteUserInput struct {
	ID string `json:"id"`
}

func (s *Service) DeleteUser(input DeleteUserInput) error {
	err := s.userRepository.Delete(input.ID)

	if err != nil {
		return err
	}

	return nil
}
