# 01_Web communication

# **1) 웹 통신 & Protocol**

- 인터넷 상에서의 통신을 말한다.
- 많은 정보들이 주고받기에 인터넷에 엄격한 규약이 존재한다. 이 것을 **Protocol**이다.

# **2) Protocol 종류**

- 일반적인 프로토콜**Http : Hyper Text Transer ProtocolHttps : secure Hyper Text Transer Protocol**
- TCP/IP 프로토콜을 가지고 서버와 클라이언트 사이의 파일 전송을 하기 위한 프로토콜**FTP : File Transfer Protocol**
- 파일 전송 프로토콜**Telnet : Terminal NetworkSSH : Secure Shell**
- 보안된 소켓 통신ㅃ을 위한 프로토콜을**SMTP : Simple Mail Transfer Protocol**
- 기타**TCP/UDP : Transmission Control Protocol/User Datagram ProtocolIP : Internet Protocol**

# **1. Http 프로토콜**

# **1) Hyper Text Transfer Protocol**

- Hyper Text를 전송하기 위한 프로토콜
- Hyper Textf란, 웹 문서를 구성하고 있는 언어. 즉, HTML을 의미한다.

# **2) HTML**

- Hyper Text Markup Language
- Hyper Text : text를 넘어서 링크, 이미지 등 다양한 것들을 표현할 수 있는 것
- HTML : 웹 문서의 뼈대를 구성하는 언어. 브라우저를 통해서 웹 문서를 읽을 수 있다.

# **3) Http 통신 - Requst & Response**

- 요청 Request, 응답 Response으로 이루어짐.
- 클라이언트가 서버에게 요청을 보냄
- 서버는 요청에 대한 응답 결과를 줌
- 클라이언트 사용자에게 응답 받은 결과를 보여줌 (랜더링이라고 한다.)

# **4) Http 통신 - stateless**

- Http 통신은 state 개념이 존재하지 않는다.
- 통신을 주고 받아도 클라이언트와 서버는 연결되어 있는 것이 아니라 각가가의 통신은 독립적인다.
- 상태를 저장하지 않는다는 의미다. (서로 요청한 것들을 기억하지 못한다.)
- 그래서 로그인 같은 경우 세션/저장소 같은 방식으로 이용하여 기억하는것 처럼 보이게한다.

# 5**) Http Request 구조**

### **1) Start Line(요청 내용) :**

- **Http메소드** :
- Request target : 요청의 의도를 담고 있음 (GET, POST, DELETE, UPDATE)
- Http ver : 버전에 따라 요청 메시지 구조나 데이터가 다를 수 있어서, version을 명시

### **2) Header : HTTP 요청이 전송되는 목표 주소**

### **3) Body**