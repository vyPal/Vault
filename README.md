# Mimlex Vault
The storage backend for Mimlex.

## Routes
### POST /health
Check if the service is up and running.

#### Response
```json
{
    "status": "ok"
}
```
### GET /files/list/:username/:appname/:filepath
List all files in a directory.

#### Parameters
- recursive: boolean (optional)
- version: string (optional)

#### Required Scopes
- files:list
- files:app OR files:all

#### Response
<details>
<summary>Response</summary>
```json
[
    {
        "etag": "id",
        "name": "full file path",
        "lastModified": "2024-09-28T17:11:38.789Z",
        "size": 15714,
        "contentType": "",
        "expires": "0001-01-01T00:00:00Z",
        "metadata": null,
        "UserTagCount": 0,
        "Owner": {
            "owner": {
                "Space": "http://s3.amazonaws.com/doc/2006-03-01/",
                "Local": "Owner"
            },
            "name": "02d6176db174dc93cb1b899f7c6078f08654445fe8cf1b6ce98d8855f66bdbf4",
            "id": "minio"
        },
        "Grant": null,
        "storageClass": "STANDARD",
        "IsLatest": false,
        "IsDeleteMarker": false,
        "VersionID": "",
        "ReplicationStatus": "",
        "ReplicationReady": false,
        "Expiration": "0001-01-01T00:00:00Z",
        "ExpirationRuleID": "",
        "Restore": null,
        "ChecksumCRC32": "",
        "ChecksumCRC32C": "",
        "ChecksumSHA1": "",
        "ChecksumSHA256": "",
        "Internal": null
    }
]
```
</details>

### GET /files/download/:username/:appname/:filepath
Download a file.

#### Required Scopes
- files:read
- files:app OR files:all

#### Response
File content.

### POST /files/upload/:username/:appname/:filepath
Upload a file.

#### Request
File content.

#### Required Scopes
- files:upload
- files:app OR files:all

### DELETE /files/delete/:username/:appname/:filepath
Delete a file.

#### Required Scopes
- files:write
- files:app OR files:all

#### Response
File deleted.

### GET /files/metadata/:username/:appname/:filepath
Get file metadata.

#### Required Scopes
- files:read
- files:app OR files:all

#### Response
```json
{
    "etag": "id",
    "name": "full file path",
    "lastModified": "2024-09-28T17:11:38.789Z",
    "size": 15714,
    "contentType": "",
    "expires": "0001-01-01T00:00:00Z",
    "metadata": null,
    "UserTagCount": 0,
    "Owner": {
        "owner": {
            "Space": "http://s3.amazonaws.com/doc/2006-03-01/",
            "Local": "Owner"
        },
        "name": "02d6176db174dc93cb1b899f7c6078f08654445fe8cf1b6ce98d8855f66bdbf4",
        "id": "minio"
    },
    "Grant": null,
    "storageClass": "STANDARD",
    "IsLatest": false,
    "IsDeleteMarker": false,
    "VersionID": "",
    "ReplicationStatus": "",
    "ReplicationReady": false,
    "Expiration": "0001-01-01T00:00:00Z",
    "ExpirationRuleID": "",
    "Restore": null,
    "ChecksumCRC32": "",
    "ChecksumCRC32C": "",
    "ChecksumSHA1": "",
    "ChecksumSHA256": "",
    "Internal": null
}
```
