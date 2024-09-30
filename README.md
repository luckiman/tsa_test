# tsa_test

## Contact Management API with Validation and Database Integration
This API allows for the creation of contact information using Go and the Gin framework, while integrating PostgreSQL for persistence. A key focus of the API is on validating Australian phone numbers in E.164 format, ensuring data integrity before it is stored in the database.

### What Was Solved and Implemented in the Project
#### 1. API Endpoint for Creating Contacts
Endpoint: POST /contacts
Functionality: Allows users to create a new contact by submitting a full name, an optional email, and one or more phone numbers. The phone numbers must be valid Australian numbers formatted according to the E.164 standard.
#### 2. Phone Number Validation
- ^\+61: The phone number must start with the +61 country code.
- [2-478]\d{8}: This validates landline and mobile numbers:
    Starts with 2, 3, 4, 7, or 8 (depending on the type of number).
    Followed by 8 digits for mobile numbers and 9 digits for landline numbers.
- 1800\d{6}: This validates toll-free numbers starting with 1800, followed by 6 digits.
- $: Ensures there are no extra characters after the valid phone number.
#### 3. Handling Optional Fields (Email)
- The email field is optional. The API allows contacts to be created without an email address.
- If the email is provided, it is stored in the database. If it is missing, the email field is set to NULL in the database.
#### 4. Database Integration with PostgreSQL
- The contact information is persisted in a PostgreSQL database.
- The contacts table stores:
    Full name (full_name: TEXT)
    Email (email: TEXT, which can be NULL)
    Phone numbers (phone_numbers: joined string with ',' in PostgreSQL, TEXT)
#### 5. Storing Multiple Phone Numbers
- The API allows each contact to have multiple phone numbers, stored as an array in PostgreSQL (TEXT[]).
- Phone numbers are validated and normalized (formatted according to E.164) before being stored.
#### 6. Error Handling
- If a phone number is invalid, the API returns a 400 Bad Request response with a descriptive error message.
{
    "error": "Invalid Australian phone number"
}
- General database or internal server errors are caught and returned with appropriate status codes.
#### 7. Test Coverage for API
- A series of unit tests were written to validate the core functionalities of the API, using Go's testing package and httptest.
- Tests include:
    Creating contacts with valid data.
    Handling invalid phone numbers correctly.
    Handling requests where the email is omitted.