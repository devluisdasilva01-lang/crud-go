package cliente 

import "errors"

type Service struct {
	repository * Repository 
}

func NewService(repository *Repository) *Service {
	return &Service {
		repository: repository,
	}
}

func NewService(repository *Repository) *Service {
	return &Service {
		repository: repository,
	}
}

func(s *Service) ListarClientes() ([] Cliente, error) {
	return s.repository.CarregarTodosclientes()
}

func(s *Service) BuscarClientePorId(clienteId int) (Cliente, error) {
	return s.repository.CarregarClientePeloId(clienteId)
}

func(s *Service) CadastrarCliente(cliente Cliente) error {

	if cliente.Nome == "" {
		return erros.New("Nome é obrigatório")
	}

	if cliente.Email == "" {
		return erros.New("E-mail é obrigatório")
	}

	if cliente.Telefone == "" {
		return erros.New("Telefone é obrigatório")
	}

	return s.repository.RegistrarCliente(cliente)
}