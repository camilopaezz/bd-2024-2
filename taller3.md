# Taller 3: Normalización de Bases de datos - Forma Normal
Camilo Esteban Paez

## Análisis inicial
La tabla `productos` tiene las siguiente columnas: `PedidoID | ClienteID | ProductoID | Cantidad | Fecha | MetodoPago` , `PedidoID` es su PRIMARY KEY, `ClientID | ProductoID` son sus FOREIGN KEYS que permiten la no repetición de la información tanto de los clientes como del producto, `Cantidad` es un INT lo cual es tal-vez algo excesivo, `Fecha` es un DATE y `MetodoPago` es un VARCHAR(50) aquí también hay un problema.

### Problemas Identificados
#### No relacionados a la normalización
- `Cantidad` tiene tipo INT firmado, al ser una tienda minorista no es necesario una cantidad posible tan grande, ademas no es posible pedir una cantidad negativa

#### Relacionados con la normalizacion
- Incapacidad de determinar informacion relacionada con fechas, al usar el tipo DATE se omite el tiempo que para determinar hechos sobre la compra puede ser necesario, se podria solucionar usando el tipo DATETIME en vez
- 
- `MetodoPago` se comporta como un Enum pero esta definido como string, se puede extraer a otra tabla, compuesta por las columnas `MetodoPagoId | descripcion`.
