Originaly this project ment to be a distributed key-value storage. Origignal design was taken from this website 
https://www.educative.io/courses/grokking-modern-system-design-interview-for-engineers-managers/system-design-the-key-value-store  and this medium article https://medium.com/@mehar.chand.cloud/design-key-value-store-for-distributed-system-83ff11594b3e 
Curently It is rest api with 2 services communicating via http. The first one balances workloads and ensures that data is distributed and replicated.
Seccond is node api that stores data in json document using hashing techniques. 
Project is unfinihsed and needs to be changed in many ways.
