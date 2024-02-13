# Five9 SCIM Bridge
This takes advantage of the Five9 VCC Admin API in order to bridge IDPs with Five9. This enables automated provisioning of users. This application takes advantage of the leg work done by the following projects:
- https://github.com/elimity/scim  
Base SCIM Server functionality
- https://github.com/scim2/filter-parser  
Parsing SCIM filter expressions

# Features
- Creation of users
- User permission management via "roles" attribute
- User deactivation
- VCC Federation ID sync
- Built in Filter evaluator that takes advantage of go slices

## Notable Behavior
- DELETE requests do not delete users, only deactivate
- Does not update the VCC userName attribute since that is read only. Service tries to key off the IDP provided userName as the VCC federationId value
- Will try to set the userName attribute only on Create/POST requests 
- Filter evaluator is custom built so problems are likely however testing so far proves it is sufficient for the needs this fulfills
- List of users is stored in memory, pre mapped into SCIM resources and is refreshed every 15 minutes


# Getting started
1. Clone repo
2. Configure env variables based on .env.example
3. `go get ./cmd/five9-scim`
4. `go run ./cmd/five9-scim`

# Startup Logic
1. Initialize the Five9 API client using env variables
2. Initialize the data Sync go routine
3. Start the HTTP server utilizing the SCIM handler
