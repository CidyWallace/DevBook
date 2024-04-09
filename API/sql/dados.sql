insert into usuarios (nome, nick, email, senha)
values
("usuário 1", "usuario_1", "usuario1@gmail.com", "$2a$10$iv5wOic4FkacdWratFnLn.E.y8UCCVUYrFbbAftC8N729xcO6HmAy"),
("usuário 2", "usuario_2", "usuario2@gmail.com", "$2a$10$iv5wOic4FkacdWratFnLn.E.y8UCCVUYrFbbAftC8N729xcO6HmAy"),
("usuário 3", "usuario_3", "usuario3@gmail.com", "$2a$10$iv5wOic4FkacdWratFnLn.E.y8UCCVUYrFbbAftC8N729xcO6HmAy");

insert into seguidores(usuario_id, seguidor_id)
values
(1, 2),
(3, 1),
(1, 3);

insert into publicacoes(titulo, conteudo, autor_id)
values
("Publicação do usuário 1", "Essa é a publicação do usuário 1", 1),
("Publicação do usuário 1", "Essa é a publicação do usuário 2", 2),
("Publicação do usuário 1", "Essa é a publicação do usuário 3", 3);