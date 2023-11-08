package goco2

/*
Package goco2 implements a microservice for monitoring CO2 emissions savings

Note that this implementation also allow for managing energy efficiency upgrade interventions.
The CO2 emissions saving calculation is based on the inserted upgrade interventions, in a different scenario the
interventions information could be retrieved from another microservice or else.

The intervention data is stored in memory for the sake of simplicity, a different storage backend could be used.

The service expose a REST API to:
 - GET CO2 saving information
 - POST new intervention data

Whatever messaging/queueing system may be implemented as a transport, such as grpc, kafka, nats etc.

Note that, for the sake of greater separation of concerns, the implementation could define interfaces also for the
handlers/controllers layer and the data access layer. This seems overkill for this basic example, where we only
separates service/business logic.
*/
