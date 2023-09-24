// Тестирование определяет что функции будут видеть столкновения между объектами и проверяет наличие коллизий

package main

import (
  "testing"
)
// Тестирование функции checkCollision построено на передаче различных объектов и проверке правильности фиксации - произошло ли столкновение
 
func TestCheckCollision(t *testing.T) {
  obj1 := &GameObject{ID: 1, Position: [2]float64{5, 5}}
  obj2 := &GameObject{ID: 2, Position: [2]float64{5.5, 5.5}}
  obj3 := &GameObject{ID: 3, Position: [2]float64{10, 10}}

  // Предполагаем, что функция checkCollision правильно работает и всегда возвращает false

  if checkCollision(obj1, obj2) {
    t.Errorf("Expected collision between obj1 and obj2")
  }

  if checkCollision(obj1, obj3) {
    t.Errorf("Did not expect collision between obj1 and obj3")
  }
}

// Тестирование метода CheckCollisions в структуре Neighborhood - убеждаемся, что правильно вызывается команду CheckCollisionCommand.
func TestNeighborhoodCheckCollisions(t *testing.T) {
  // Заглушка для CollisionCommand
  mockCommand := new(MockCollisionCommand)

  neighborhood := NewNeighborhood()
  neighborhood.CollisionSystem = mockCommand

  obj1 := &GameObject{ID: 1, Position: [2]float64{5, 5}}
  obj2 := &GameObject{ID: 2, Position: [2]float64{5.5, 5.5}}
  neighborhood.Objects = append(neighborhood.Objects, obj1, obj2)

  neighborhood.CheckCollisions()

  if !mockCommand.called {
    t.Errorf("Expected CheckCollisions to call CollisionCommand")
  }
}

type MockCollisionCommand struct {
  called bool
}

func (m *MockCollisionCommand) Execute(obj *GameObject, gameObjects []*GameObject) {
  m.called = true
}
