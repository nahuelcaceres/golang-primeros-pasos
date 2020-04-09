package tp1

import (
	"strconv"
)

// Producto contiene metodos que nos permiten acceder
// a atributos que esperamos de un Producto.
type Producto interface {
	ID() int
	Precio() int
}

// Productos es una lista de productos donde para cada producto
// se sabe el nombre del super mercado, el id y su precio.
// Esta estructura se puede cargar usando la funcion LeerProductos
// que carga informacion guardada en `productos.json`.
type Productos [][]string

// Carrito contiene el nombre de la tienda y el precio final luego
// de sumar todos los productos.
type Carrito struct {
	Tienda string
	Precio int
}

var mapSupermercados map[string]Supermercado

// CalcularPrecios recibe un arreglo de los IDs de productos y calcula,
// para cada super mercado, cuanto saldria comprar esos productos ahi.
// Retorna un slice de carritos, donde se tiene uno para cada super mercado.
func (p Productos) CalcularPrecios(ids ...int) []Carrito {

	if len(ids) == 0 {
		return nil
	}

	//Inicializar el map de Supermercados con ProductosItems
	mapSupermercados := GenerarMapSupermercados(p)

	carrito := []Carrito{}

	for _, supermercado := range mapSupermercados {

		var precioTotal int

		for _, idproducto := range ids {

			productoItem := supermercado.Productos[idproducto]
			precioTotal += productoItem.PrecioValue
		}

		carrito = append(carrito, Carrito{Tienda: supermercado.Nombre, Precio: precioTotal})

	}

	return carrito
}

// Promedio recibe el id de un producto y retorna el precio promedio
// de ese producto usando los precios de todos los supermercados.
func (p Productos) Promedio(idProducto int) float64 {

	mapSupermercados := GenerarMapSupermercados(p)

	var precioProducto int

	for _, supermercado := range mapSupermercados {

		precioProducto += supermercado.Productos[idProducto].PrecioValue

	}

	return float64(precioProducto) / float64(len(mapSupermercados))
}

// BuscarMasBarato recibe un id de producto a buscar y te retorna
// el producto mas barato que haya encontrado.
func (p Productos) BuscarMasBarato(idProducto int) (Producto, bool) {

	var tengoProducto bool

	mapSupermercados := GenerarMapSupermercados(p)

	productoItemMasBarato := ProductoItem{IdProducto: idProducto}

	for _, supermercado := range mapSupermercados {

		productoItem, existe := supermercado.Productos[idProducto]

		if existe {
			tengoProducto = true

			if (productoItemMasBarato.PrecioValue > productoItem.PrecioValue) || (productoItemMasBarato.PrecioValue == 0) {
				productoItemMasBarato.PrecioValue = productoItem.PrecioValue
			}

		}

	}

	return productoItemMasBarato, tengoProducto
}

func (productoItem ProductoItem) ID() int {
	return productoItem.IdProducto
}
func (productoItem ProductoItem) Precio() int {
	return productoItem.PrecioValue
}

//Por que no funciona en el archivo producto.go
func GenerarMapSupermercados(listadoProductos [][]string) map[string]Supermercado {

	if mapSupermercados == nil {

		mapSupermercados = make(map[string]Supermercado)

		for _, valor := range listadoProductos {

			supermercado, existe := mapSupermercados[valor[0]]

			if !existe {
				supermercado = Supermercado{Nombre: valor[0], Productos: make(map[int]ProductoItem)}
				mapSupermercados[supermercado.Nombre] = supermercado
			}

			precio, _ := strconv.Atoi(valor[2])
			idProducto, _ := strconv.Atoi(valor[1])

			productoItem := ProductoItem{
				IdProducto:  idProducto,
				Nombre:      "",
				PrecioValue: precio,
			}

			supermercado.Productos[productoItem.IdProducto] = productoItem

		}

	}

	return mapSupermercados
}
