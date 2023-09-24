// Реализация проверки коллизий игровых объектов:
// Создадим структуру GameObject - которая представляет игровой объект (у объекта есть уникальный идентификатор и позиция на игровом поле)
// Команда CheckCollisionCommand - проверяет столкновения для определенного объекта со всеми другими объектами в его окрестности
// Структура Neighborhood - представляет окрестность и содержит объекты находящиеся в этой окрестности
// Метод CheckCollisions в Neighborhood - вызывает команду проверки столкновений для каждого объекта в окрестности

package main
import "fmt"
type GameObject struct {
  ID       int
  Position [2]float64 // x, y
}
type CollisionCommand interface {
  Execute(obj *GameObject, gameObjects []*GameObject)
}
type CheckCollisionCommand struct{}
func (c *CheckCollisionCommand) Execute(obj *GameObject, gameObjects []*GameObject) {
  // Функция проверки коллизий
  for _, gameObject := range gameObjects {
    if obj.ID != gameObject.ID { // Проверка, чтобы не сравнивать объект с самим собой
      if checkCollision(obj, gameObject) {
        fmt.Printf("Collision detected between object %d and object %d\n", obj.ID, gameObject.ID)
      }
    }
  }
}
func checkCollision(obj1 *GameObject, obj2 *GameObject) bool {
  // Заглушка для проверки столкновения между obj1 и obj2
  return false
}
type Neighborhood struct {
  Objects         []*GameObject
  CollisionSystem CollisionCommand
}
func NewNeighborhood() *Neighborhood {
  return &Neighborhood{
    Objects:         make([]*GameObject, 0),
    CollisionSystem: &CheckCollisionCommand{},
  }
}
func (n *Neighborhood) CheckCollisions() {
  for _, obj := range n.Objects {
    n.CollisionSystem.Execute(obj, n.Objects)
  }
}
func main() {
  // Создание двух объектов для демонстрации
  obj1 := &GameObject{ID: 1, Position: [2]float64{5, 5}}
  obj2 := &GameObject{ID: 2, Position: [2]float64{5.5, 5.5}}
  // Создание окрестности и добавление объектов
  neighborhood := NewNeighborhood()
  neighborhood.Objects = append(neighborhood.Objects, obj1, obj2)
  // Проверка столкновений
  neighborhood.CheckCollisions()
}
