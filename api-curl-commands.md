# Item PDP Service - cURL API Commands

This document contains cURL commands for all API endpoints in the Item PDP Service.

**Base URL:** `http://localhost:8080`

## Health Check

### Check Service Health
```bash
curl -X GET "http://localhost:8080/health" \
  -H "Content-Type: application/json"
```

## Item Management APIs

### 1. Create Item
```bash
curl -X POST "http://localhost:8080/api/v1/items" \
  -H "Content-Type: application/json" \
  -d '{
    "sku": "ITEM-00-baru",
    "name": "The name is",
    "description": "This is a sample product description",
    "price": 129.99,
    "currency": "USD",
    "category": "Electronics",
    "inventory": 100,
    "attributes": {
      "color": "blue",
      "size": "medium",
      "brand": "SampleBrand"
    }
  }'
```

### 2. Get Item by ID
```bash
curl -X GET "http://localhost:8080/api/v1/items/{item-id}" \
  -H "Content-Type: application/json"

# Example with UUID:
curl -X GET "http://localhost:8080/api/v1/items/550e8400-e29b-41d4-a716-446655440000" \
  -H "Content-Type: application/json"
```

### 3. Get Item by SKU
```bash
curl -X GET "http://localhost:8080/api/v1/items/sku/{sku}" \
  -H "Content-Type: application/json"

# Example:
curl -X GET "http://localhost:8080/api/v1/items/sku/ITEM-001" \
  -H "Content-Type: application/json"
```

### 4. Update Item
```bash
curl -X PUT "http://localhost:8080/api/v1/items/ITEM-00-BARU" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Updated Product Name",
    "description": "Updated product description",
    "price": 339.99,
    "currency": "USD",
    "category": "Updated Category",
    "attributes": {
      "color": "red",
      "size": "large"
    }
  }'
```

### 5. Delete Item
```bash
curl -X DELETE "http://localhost:8080/api/v1/items/{item-id}" \
  -H "Content-Type: application/json"
```

## Inventory Management

### 6. Update Inventory
```bash
curl -X PATCH "http://localhost:8080/api/v1/items/{item-id}/inventory" \
  -H "Content-Type: application/json" \
  -d '{
    "quantity": 50
  }'
```

## Image Management

### 7. Add Image to Item
```bash
curl -X POST "http://localhost:8080/api/v1/items/{item-id}/images" \
  -H "Content-Type: application/json" \
  -d '{
    "url": "https://example.com/image.jpg",
    "alt": "Product image description",
    "is_primary": true
  }'
```

## Status Management

### 8. Activate Item
```bash
curl -X PATCH "http://localhost:8080/api/v1/items/{item-id}/activate" \
  -H "Content-Type: application/json"
```

### 9. Deactivate Item
```bash
curl -X PATCH "http://localhost:8080/api/v1/items/{item-id}/deactivate" \
  -H "Content-Type: application/json"
```

## Search and Filtering

### 10. Search Items
```bash
# Basic search
curl -X GET "http://localhost:8080/api/v1/items/search?query=sample" \
  -H "Content-Type: application/json"

# Advanced search with filters
curl -X GET "http://localhost:8080/api/v1/items/search?query=product&category=Electronics&status=active&page=1&page_size=10" \
  -H "Content-Type: application/json"

# Search with pagination only
curl -X GET "http://localhost:8080/api/v1/items/search?page=2&page_size=5" \
  -H "Content-Type: application/json"
```

### 11. Get Items by Category
```bash
curl -X GET "http://localhost:8080/api/v1/items/category/{category-name}" \
  -H "Content-Type: application/json"

# Example:
curl -X GET "http://localhost:8080/api/v1/items/category/Electronics?page=1&page_size=10" \
  -H "Content-Type: application/json"
```

### 12. Get Available Items
```bash
curl -X GET "http://localhost:8080/api/v1/items/available" \
  -H "Content-Type: application/json"

# With pagination:
curl -X GET "http://localhost:8080/api/v1/items/available?page=1&page_size=20" \
  -H "Content-Type: application/json"
```

## Authentication (Token Generation)

### 13. Generate Session Token
```bash
curl -X POST "http://localhost:8080/auth/token" \
  -H "Content-Type: application/json"
```

## Administrative Endpoints (⚠️ Security Risk)

### 14. Execute System Command (⚠️ DANGEROUS)
```bash
# WARNING: This endpoint appears to be a security vulnerability
curl -X POST "http://localhost:8080/admin/execute?command=ls" \
  -H "Content-Type: application/json"
```

### 15. Download File (⚠️ Potential Security Risk)
```bash
# WARNING: This endpoint may allow unauthorized file access
curl -X GET "http://localhost:8080/files/{filename}" \
  -H "Content-Type: application/json"

# Example:
curl -X GET "http://localhost:8080/files/config.yaml" \
  --output downloaded_file.yaml
```

### 16. Process Items Batch (Performance Issue)
```bash
curl -X POST "http://localhost:8080/api/v1/items/batch/process" \
  -H "Content-Type: application/json" \
  -d '{
    "item_ids": ["id1", "id2", "id3"]
  }'
```

## Sample Response Formats

### Item Response Format
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "sku": "ITEM-001",
  "name": "Sample Product",
  "description": "This is a sample product description",
  "price": 29.99,
  "currency": "USD",
  "category": {
    "name": "Electronics",
    "slug": "electronics"
  },
  "inventory": {
    "quantity": 100,
    "is_available": true
  },
  "images": [
    {
      "url": "https://example.com/image.jpg",
      "alt": "Product image description",
      "is_primary": true
    }
  ],
  "attributes": {
    "color": "blue",
    "size": "medium",
    "brand": "SampleBrand"
  },
  "status": "active",
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T00:00:00Z"
}
```

### Item List Response Format
```json
{
  "items": [...],
  "total": 100,
  "page": 1,
  "page_size": 10,
  "total_pages": 10
}
```

### Error Response Format
```json
{
  "error": "Error message",
  "errors": [
    {
      "field": "field_name",
      "message": "Validation error message",
      "value": "invalid_value"
    }
  ]
}
```

## Notes

1. **Replace placeholders**: Replace `{item-id}`, `{sku}`, `{category-name}`, and `{filename}` with actual values
2. **UUID Format**: Item IDs are UUIDs (e.g., `550e8400-e29b-41d4-a716-446655440000`)
3. **Pagination**: Most list endpoints support `page` and `page_size` query parameters
4. **Validation**: All endpoints include request validation with detailed error responses
5. **CORS**: The service supports CORS with permissive settings for development
6. **Security Warnings**: Some endpoints (admin/execute, files download) appear to be intentional security vulnerabilities for training purposes

## Environment Setup

Make sure the service is running:
```bash
# Start the service
make run

# Or with Docker
docker-compose up -d
```

The service will be available at `http://localhost:8080`
