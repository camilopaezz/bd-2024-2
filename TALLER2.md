# Taller 2 Punto 2

Bases de datos 2024-2

Camilo Esteban Paez Neuta - cpaezn@unal.edu.co

## Instruciones

2. Creacion de usuarios

```sql
CREATE USER 'admin_tienda'@'localhost'
	IDENTIFIED BY 'admin123';

CREATE USER 'empleado_tienda'@'localhost'
	IDENTIFIED BY 'empleado123';

CREATE USER 'supervisor_tienda'@'localhost'
	IDENTIFIED BY 'supervisor123';
```

3. Asignacion de permisos

```sql
-- IMPORTANTE!!: "tienda_online.*" permite todas las tablas

-- Control total
GRANT ALL PRIVILEGES ON tienda_online.* TO 'admin_tienda'@'localhost';

-- Insertar y consultar
GRANT INSERT, SELECT ON tienda_online.* TO 'empleado_tienda'@'localhost';

-- Unicamente consultar
GRANT SELECT ON tienda_online.* TO 'supervisor_tienda'@'localhost';
```

4. Pruebas de permisos

- Insertar registro con el usuario `empleado_tienda`

  ```sql
  INSERT INTO
    `productos` (`categoria`, `nombre`, `precio`, `stock`)
  VALUES
    ('Tecnologia', 'Mouse', 25, 15);
  ```

RESULTADO: _1 rows affected_

- Consultar tabla con el usuario `supervisor_tienda`

  ```sql
  SELECT * from productos
  ```

RESULTADO: No hay errores y se muestran los datos

- Intentar acciones no permitidas con cada uno

  - Eliminar registro con `empleado_tienda`

    ```sql
    DELETE FROM productos where precio=25
    ```

    **OUTPUT:** `DELETE command denied to user 'empleado_tienda'@'localhost' for table 'productos'`

  - Insertar registro con `supervisor_tienda`

        ```sql
        INSERT INTO
          `productos` (`categoria`, `nombre`, `precio`, `stock`)
        VALUES ('Tecnologia', 'Pc', 1000, 2)
        ```

        **OUTPUT:** `DELETE command denied to user 'empleado_tienda'@'localhost' for table 'productos'

    `

## Preguntas

- ¿Que ocurre al intentar realizar una accion no permitida?

  **RESPUESTA:** La base de datos devuelve un error que describe donde, quien y que, hizo una accion no permitida, nunca se llega a ejecutar el comando.

- ¿Que ocurre si intentas acceder con un usuario que no inexistente?

  **RESPUESTA:** La base de datos devuelve un error donde se le niega el acceso al usuario.
