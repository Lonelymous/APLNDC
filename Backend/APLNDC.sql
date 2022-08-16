CREATE DATABASE `APLNDC`;

CREATE TABLE `Usuarios`(

`UsuarioIdentificaciónNúmero`       INT UNSIGNED NOT NULL COMMENT "UserID" AUTO_INCREMENT,
`Nombre`            VARCHAR(128)    NOT NULL    COMMENT "Name",
`Contraseña hash`   VARCHAR(64)     NOT NULL    COMMENT "Password Hash",
`Sal`               VARCHAR(8)      NOT NULL    COMMENT "Salt",

PRIMARY KEY(`UsuarioIdentificaciónNúmero`)
);

CREATE TABLE `Tareas`(

`TareaIdentificaciónNúmero`         INT UNSIGNED        NOT NULL    COMMENT      "TaskID" AUTO_INCREMENT,
`UsuarioIdentificaciónNúmero`       INT UNSIGNED        NOT NULL    COMMENT      "UserID",
`Nombre`                            VARCHAR(64)         NOT NULL    COMMENT      "Name",
`Descripción`                       TEXT                NOT NULL    COMMENT      "Description",
`Prioridad`                         ENUM("critical","importante","sérieux","annyira nem qva fontos","unwichtig")NOT NULL    COMMENT      "Priority",
`Hecho`                             BOOLEAN             NOT NULL    COMMENT      "Done" DEFAULT FALSE,
`Plazo`                             DATETIME            NOT NULL    COMMENT      "Deadline",
`CreadoEn`                          DATETIME            NOT NULL    COMMENT      "CreatedAt" DEFAULT NOW(),

PRIMARY KEY (`IdentificaciónNúmero`),
FOREIGN KEY (`UsuarioIdentificaciónNúmero`) REFERENCES `Usuarios`(`UsuarioIdentificaciónNúmero`)
);