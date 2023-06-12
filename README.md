# TiktokTechImmersion-2023

Assignment
Design and implement a backend Instant Messaging system.

Overview:
In this assignment, you will design and develop an IM system implementing a set of specific APIs using Golang. You need only develop the backend side of the system, focusing on core message features without the front-end part and the account/authentication part.

Requirements:
1. Architecture: The system should contain two services: one HTTP server and one RPC server. The IDL of HTTP API is provided below. The RPC IDL is not restricted, but you can also get an example from the demo repository.
2. Data storage: The system should store messages data. Receivers can access this data at any time. At least one database must be used (MySQL or Redis are recommended). There is no limitation on data schema design.
3. Message delivery: The system should be able to deliver messages to the intended recipients by PULL mode in a timely and consistent manner. Pull mode means there is no need to maintain the connection and push new messages to receivers in real-time. Only the pull API must be implemented, so the receiver can use pull API to fetch messages.
4. Performance and scalability: The system should be designed to handle a relatively large number of users and messages. (Support more than 20 concurrency in testing)
