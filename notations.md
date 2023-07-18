- SSL (Secure Sockets Layer)
é um protocolo de segurança que criptografa a comunicação entre um navegador e um servidor web. Isso protege os dados do usuário, como senhas, números de cartão de crédito e outras informações confidenciais, de serem interceptados por terceiros, para usar o SSL, um site precisa ter um certificado SSL, um certificado SSL é um arquivo digital que identifica o site e garante que ele seja seguro.

comando que gera a chave com ssl com nome e tamanho
`openssl genrsa -out server.key 2048`

comando que gera o certificado com encriptação e expiração
`openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650`

- API stateless
não guarda informações e sessão, cada request é independente.

- JWT (JSON Web Token)
é um padrão de token aberto que define um método compacto e autocontido para transmitir informações entre partes. JWTs são frequentemente usados como tokens de acesso, que são usados para autenticar usuários em APIs, e é composto de 3 partes.

  Header: O cabeçalho contém o tipo de token (JWT) e o algoritmo de hashing usado para assinar o token.

  Payload: O payload contém as claims, que são informações sobre o token, como o nome do usuário, o tempo de expiração e outros dados.

  Signature: A assinatura é um hash do header e do payload, assinado com uma chave secreta.

criando a chave privada com openssl
`openssl genrsa -out ./key 4096`

criando a chave publica com openssl
`openssl rsa -in ./key -pubout -out ./key.pub`

- GraphQL
é uma linguagem de consulta de API de código aberto que permite que os clientes especifiquem exatamente os dados que desejam receber de um servidor

- WebSocket
é uma tecnologia de comunicação de canal bidirecional que mantém uma conexão aberta entre um cliente e um servidor. Isso permite que o cliente e o servidor troquem mensagens a qualquer momento, em vez de precisar esperar por um pedido HTTP do cliente.

- GRPC
é um framework de RPC de código aberto, rápido e eficiente desenvolvido pelo Google. Ele é construído em cima do protocolo HTTP/2 e usa o modelo de dados Protobuf para serializar mensagens.