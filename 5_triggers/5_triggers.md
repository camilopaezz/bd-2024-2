# Taller triggers

```sql
create table Productos (
 nombre varchar(30) not null,
  precio float(2) not null,
  stock int default 0,
  producto_id char(36),
 primary key (producto_id)
);

create table Ventas (
 producto_id char(36) not null,
  cantidad int not null,
  fecha datetime default current_timestamp,
  venta_id int auto_increment,
  primary key(venta_id),
  foreign key (producto_id) references Productos(producto_id)
);

create table Auditoria (
 tipo enum('insert', 'update', 'delete') not null,
  fecha datetime default current_timestamp,
  info varchar(100) not null,
  auditoria_id int auto_increment,
  primary key (auditoria_id)
);
```

```sql
create function calcular_total(precio float(2), cantidad int)
returns float(2) deterministic
begin
 return precio * cantidad;
end;

create function calcular_total(precio float(2), cantidad int)
returns FLOAT(2) deterministic
begin
  return precio * cantidad;
end;

create procedure registrar_venta(
    IN in_producto_id CHAR(36),
    IN in_cantidad INT)

 begin
    start transaction;
    insert into Ventas (producto_id, cantidad) 
    values (in_producto_id, in_cantidad);
    
    update Productos 
    set stock = stock - in_cantidad 
    where producto_id = in_producto_id;
    
    commit;
end;

create trigger after_venta
after insert on Ventas
for each row
begin
    insert into Auditoria (tipo, info)
    values ('insert', 
        CONCAT('Venta ID ', NEW.venta_id,
            ', Producto ', NEW.producto_id,
            ', Cantidad ', NEW.cantidad));
end;

create view resumen_ventas as
select 
    p.nombre as producto,
    v.cantidad,
    calcular_total(p.precio, v.cantidad) as total
from Ventas v
join Productos p on v.producto_id = p.producto_id;

create trigger validar_stock
before insert on Ventas
for each row
begin
    declare current_stock INT;
    
    select stock into current_stock 
    from Productos 
    where producto_id = NEW.producto_id;
    
    if current_stock < NEW.cantidad then
        signal sqlstate '45000'
        set message_text = 'Stock insuficiente para realizar la venta';
    end if;
end;
```
