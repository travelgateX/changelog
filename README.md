# API Change Log. Análisis
## Alcance proyecto
Desarrollar un API en el que los distintos recursos puedan “alimentar” con los cambios que se van realizando. Se pretende agrupar y estandarizar una práctica común a la vez de comprobar si es posible “automatizar” las entradas con los commits de cada recurso en particular.

## Requerimientos técnicos
El “end point” del API ha de estar desarrollado en GoLang + Graphql.

## Requerimientos funcionales
-Los cambios se agruparán por un código de versionado incremental y en formato X.Y.Z.
-Los cambios irán ligados a un recurso
-Los cambios podrán ser de los siguientes tipos:
--Unreleased: funcionalidades todavía sin añadir
--Added: nuevas funcionalidades
--Changed: cambios en funcionalidades existentes
--Deprecated: funcionalidades que desaparecen pronto
--Removed: funcionalidades eliminadas
--Fixed: errores corregidos
--Security: vulnerabilidades corregidas
