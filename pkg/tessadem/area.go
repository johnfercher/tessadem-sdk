package tessadem

type AreaRequest struct {
	Units      Unit
	Northeast  *Vector2D
	Southwest  *Vector2D
	SquareSize int
}

type AreaResponse struct {
	Results [][]*Vector3D `json:"results"`
}
