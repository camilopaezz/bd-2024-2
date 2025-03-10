# Taller 6 - Indices

- Camilo Esteban Paez Neuta
- cc. 1110450891
- cpaezn@unal.edu.co

## 1. Creacion de la base de datos

```sql
CREATE DATABASE universidad;

CREATE TABLE estudiantes (
    id INT PRIMARY KEY AUTO_INCREMENT,
    nombre VARCHAR(100),
    edad INT,
    carrera VARCHAR(50),
    promedio DECIMAL(3, 2)
) ENGINE=MEMORY;
```

## 2. Insercion de datos

Usando faker y un script hecho en go, insertamos los datos, si desea ver el codigo, esta diponible [aquí](https://github.com/camilopaezz/bd-2024-2/blob/main/scripts/insert6.go).

## 3. Consultas sin indices

Usando profiling se nos entregan los siguientes resultados:

```sql
SELECT * FROM estudiantes WHERE nombre = 'Juan';
// Sin indices: 0.00058325

SELECT * FROM estudiantes WHERE edad BETWEEN 20 AND 25;
// Sin indices: 0.00066525

SELECT * FROM estudiantes ORDER BY promedio DESC;
// Sin indices: 0.00147375
```

## 4. Creacion de indices

```sql
CREATE INDEX idx_nombre ON estudiantes(nombre);

CREATE INDEX idx_edad_carrera ON estudiantes(edad, carrera);

CREATE INDEX idx_promedio_hash ON estudiantes(promedio) USING HASH;
```

## 5. Consultas con indices

```sql
SELECT * FROM estudiantes WHERE nombre = 'Juan';
// Con indices: 0.00041075

SELECT * FROM estudiantes WHERE edad BETWEEN 20 AND 25;
// Con indices: 0.000634

SELECT * FROM estudiantes ORDER BY promedio DESC;
// Con indices: 0.000874
```

## 6. Analisis de resultados

### Resultados

```sql
SELECT * FROM estudiantes WHERE nombre = 'Juan';
// Sin indices: 0.00058325, Con indices: 0.00041075, diferencia: -5.35%

SELECT * FROM estudiantes WHERE edad BETWEEN 20 AND 25;
// Sin indices: 0.00066525, Con indices: 0.000634, diferencia: -4.69%

SELECT * FROM estudiantes ORDER BY promedio DESC;
// Sin indices: 0.00147375, Con indices: 0.000874, diferencia: -40.69%
```

- Es muy problable que al buscar a "Juan" si no existe tenga que recorrer gran parte del arbol en el peor de los casos, por eso no vemos una diferencia sustancial. Mientras que si buscamos "Rosemary esta diferencia si es notable"

```sql
SELECT * FROM estudiantes WHERE nombre = 'Rosemary';
// Sin indices: 0.00033475, Con indices: 0.0002515, diferencia: -24.86%
```

- Las limitaciones de las tablas hash es clara, para que sean usadas la información debe estar en memoria, en tablas con millones de filas se hace inviable, es esta la razón por la que las bases de datos en memoria como redis se usen en casos donde la velocidad de procesamiento es mas critica que la conservación de la información que se puede ir  eliminando o escribiéndola en disco después de cierto tiempo.

## Reflexiones

- Se observa una mejora significativa cuando los datos discriminantes son los suficiente dispersos y diferentes entre llaves primaria, ya que ordena los datos y acceder y filtar se hace menos complejo.

- No es tan util en consultas de rangos y ordenacion debido al funcionamiento de la funcion hash que no mantiene ningun orden.

- El crear un indice en una tabla implica mantener una estructura de datos, que tiene que ser rebalanceada cada insercion, los indices deben unicamente existir de acuerdo a como se consulta la base de datos.
