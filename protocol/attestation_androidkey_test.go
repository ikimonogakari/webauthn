package protocol

import (
	"crypto/sha256"
	"testing"

	"github.com/go-webauthn/webauthn/metadata"
)

func TestVerifyAndroidKeyFormat(t *testing.T) {
	type args struct {
		att            AttestationObject
		clientDataHash []byte
	}
	successAttResponse0 := attestationTestUnpackResponse(t, androidKeyTestResponse0["success"]).Response.AttestationObject
	successClientDataHash0 := sha256.Sum256(attestationTestUnpackResponse(t, androidKeyTestResponse0["success"]).Raw.AttestationResponse.ClientDataJSON)
	successAttResponse1 := attestationTestUnpackResponse(t, androidKeyTestResponse1["success"]).Response.AttestationObject
	successClientDataHash1 := sha256.Sum256(attestationTestUnpackResponse(t, androidKeyTestResponse1["success"]).Raw.AttestationResponse.ClientDataJSON)
	tests := []struct {
		name    string
		args    args
		want    string
		want1   []interface{}
		wantErr bool
	}{
		{
			"success",
			args{
				successAttResponse0,
				successClientDataHash0[:],
			},
			string(metadata.BasicFull),
			nil,
			false,
		},
		{
			"success",
			args{
				successAttResponse1,
				successClientDataHash1[:],
			},
			string(metadata.BasicFull),
			nil,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _, err := verifyAndroidKeyFormat(tt.args.att, tt.args.clientDataHash)
			if (err != nil) != tt.wantErr {
				t.Errorf("verifyAndroidKeyFormat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("verifyAndroidKeyFormat() got = %v, want %v", got, tt.want)
			}
			//if !reflect.DeepEqual(got1, tt.want1) {
			//	t.Errorf("verifySafetyNetFormat() got1 = %v, want %v", got1, tt.want1)
			//}
		})
	}
}

var androidKeyTestResponse0 = map[string]string{
	`success`: `{
		"rawId": "U5cxFNxLbU9-SAi1K7k9atYwXhghkAMbxpL__VPtBlw",
		"id": "U5cxFNxLbU9-SAi1K7k9atYwXhghkAMbxpL__VPtBlw",
		"response": {
		  "clientDataJSON": "eyJvcmlnaW4iOiJodHRwczovL2xvY2FsaG9zdDo0NDMyOSIsImNoYWxsZW5nZSI6IjlNNWY3bGp5MVl2UWNzOE9pV1FWQ3ciLCJ0eXBlIjoid2ViYXV0aG4uY3JlYXRlIn0",
		  "attestationObject": "o2NmbXRrYW5kcm9pZC1rZXlnYXR0U3RtdKNjYWxnJmNzaWdYSDBGAiEAlbQ-jtl8o9GtEstcEFH1Z_NlYsTYSn96lilEF17oEsMCIQDza5_axjn2jKZO63RlVf47DDFZbceW9b_tsh1nwOYQbmN4NWOCWQMFMIIDATCCAqegAwIBAgIBATAKBggqhkjOPQQDAjCBzjFFMEMGA1UEAww8RkFLRSBBbmRyb2lkIEtleXN0b3JlIFNvZnR3YXJlIEF0dGVzdGF0aW9uIEludGVybWVkaWF0ZSBGQUtFMTEwLwYJKoZIhvcNAQkBFiJjb25mb3JtYW5jZS10b29sc0BmaWRvYWxsaWFuY2Uub3JnMRYwFAYDVQQKDA1GSURPIEFsbGlhbmNlMQwwCgYDVQQLDANDV0cxCzAJBgNVBAYTAlVTMQswCQYDVQQIDAJNWTESMBAGA1UEBwwJV2FrZWZpZWxkMCAXDTcwMDIwMTAwMDAwMFoYDzIwOTkwMTMxMjM1OTU5WjApMScwJQYDVQQDDB5GQUtFIEFuZHJvaWQgS2V5c3RvcmUgS2V5IEZBS0UwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQbh-BQBJz7JeQ27dVvu3tyRieiEeXyDYoaWatRdy_D7q3TK96jumKlwIl5ZA2zHmKNLz4K2zsANq1X4tHp8MNZo4IBFjCCARIwCwYDVR0PBAQDAgeAMIHhBgorBgEEAdZ5AgERBIHSMIHPAgECCgEAAgEBCgEABCDc0UoXtU1CwwItW3ne2faKDcFCabFI31BufXEFVK_ENwQAMGm_hT0IAgYBXtPjz6C_hUVZBFcwVTEvMC0EKGNvbS5hbmRyb2lkLmtleXN0b3JlLmFuZHJvaWRrZXlzdG9yZWRlbW8CAQExIgQgdM_LUHSI9SkQhZHHpQWRnzJ3MvvB2ANSauqYAAbS2JgwMqEFMQMCAQKiAwIBA6MEAgIBAKUFMQMCAQSqAwIBAb-DeAMCAQK_hT4DAgEAv4U_AgUAMB8GA1UdIwQYMBaAFFKaGzLgVqrNUQ_vX4A3BovykSMdMAoGCCqGSM49BAMCA0gAMEUCIQDAPV7eQIWfL5BCmj82NszDlQ2IJsOZq_WxidwxD7On_QIgFipplgUF6OHvmHiDdaHJfFweeo60OtCDGDftjQEmF7FZAu4wggLqMIICkaADAgECAgECMAoGCCqGSM49BAMCMIHGMT0wOwYDVQQDDDRGQUtFIEFuZHJvaWQgS2V5c3RvcmUgU29mdHdhcmUgQXR0ZXN0YXRpb24gUm9vdCBGQUtFMTEwLwYJKoZIhvcNAQkBFiJjb25mb3JtYW5jZS10b29sc0BmaWRvYWxsaWFuY2Uub3JnMRYwFAYDVQQKDA1GSURPIEFsbGlhbmNlMQwwCgYDVQQLDANDV0cxCzAJBgNVBAYTAlVTMQswCQYDVQQIDAJNWTESMBAGA1UEBwwJV2FrZWZpZWxkMB4XDTE4MDUwOTEyMzE0NFoXDTQ1MDkyNDEyMzE0NFowgc4xRTBDBgNVBAMMPEZBS0UgQW5kcm9pZCBLZXlzdG9yZSBTb2Z0d2FyZSBBdHRlc3RhdGlvbiBJbnRlcm1lZGlhdGUgRkFLRTExMC8GCSqGSIb3DQEJARYiY29uZm9ybWFuY2UtdG9vbHNAZmlkb2FsbGlhbmNlLm9yZzEWMBQGA1UECgwNRklETyBBbGxpYW5jZTEMMAoGA1UECwwDQ1dHMQswCQYDVQQGEwJVUzELMAkGA1UECAwCTVkxEjAQBgNVBAcMCVdha2VmaWVsZDBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABKtQYStiTRe7w7UbBEk7BUkLjB-LnbzzebLe3KB8UqHXtg3TIXXcK37dvCbbCNVfhvZxtpTcME2kooqMTgOm9cejZjBkMBIGA1UdEwEB_wQIMAYBAf8CAQAwDgYDVR0PAQH_BAQDAgKEMB0GA1UdDgQWBBSj0qos7w2M8iQC1Ry0YLy_alskFDAfBgNVHSMEGDAWgBRSmhsy4FaqzVEP71-ANwaL8pEjHTAKBggqhkjOPQQDAgNHADBEAiBp3Z6j8YH7Qko5rRoK37nS4zPXhv65RWBV-j3MmXi50gIgPtMPpvcGtVbpFCQqsGbyhxPdkji8ltcYXQVfMhdUpRZoYXV0aERhdGFYpEmWDeWIDoxodDQXD2R2YFuP5K65ooYyx5lc87qDHZdjQQAAAFpVDktUqkdAn5qVGrdsEwExACBTlzEU3EttT35ICLUruT1q1jBeGCGQAxvGkv_9U-0GXKUBAgMmIAEhWCAbh-BQBJz7JeQ27dVvu3tyRieiEeXyDYoaWatRdy_D7iJYIK3TK96jumKlwIl5ZA2zHmKNLz4K2zsANq1X4tHp8MNZ"
		},
		"type": "public-key"
	  }`,
}

var androidKeyTestResponse1 = map[string]string{
	`success`: `{
		"id": "V51GE29tGbhby7sbg1cZ_qL8V8njqEsXpAnwQBobvgw",
		"rawId": "V51GE29tGbhby7sbg1cZ_qL8V8njqEsXpAnwQBobvgw",
		"response": {
		  "attestationObject": "o2NmbXRrYW5kcm9pZC1rZXlnYXR0U3RtdKNjYWxnJmNzaWdYRzBFAiAbZhfcF0KSXj5rdEevvnBcC8ZfRQlNl9XYWRTiIGKSHwIhAIerc7jWjOF_lJ71n_GAcaHwDUtPxkjAAdYugnZ4QxkmY3g1Y4JZAxowggMWMIICvaADAgECAgEBMAoGCCqGSM49BAMCMIHkMUUwQwYDVQQDDDxGQUtFIEFuZHJvaWQgS2V5c3RvcmUgU29mdHdhcmUgQXR0ZXN0YXRpb24gSW50ZXJtZWRpYXRlIEZBS0UxMTAvBgkqhkiG9w0BCQEWImNvbmZvcm1hbmNlLXRvb2xzQGZpZG9hbGxpYW5jZS5vcmcxFjAUBgNVBAoMDUZJRE8gQWxsaWFuY2UxIjAgBgNVBAsMGUF1dGhlbnRpY2F0b3IgQXR0ZXN0YXRpb24xCzAJBgNVBAYTAlVTMQswCQYDVQQIDAJNWTESMBAGA1UEBwwJV2FrZWZpZWxkMCAXDTcwMDIwMTAwMDAwMFoYDzIwOTkwMTMxMjM1OTU5WjApMScwJQYDVQQDDB5GQUtFIEFuZHJvaWQgS2V5c3RvcmUgS2V5IEZBS0UwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARuowgSu5AoRj8Vi_ZNSFBbGUZJXFG9MkDT6jADlr7tOK9NEgjVX53-ergXpyPaFZrAR9py-xnzfjILn_Kzb8Iqo4IBFjCCARIwCwYDVR0PBAQDAgeAMIHhBgorBgEEAdZ5AgERBIHSMIHPAgECCgEAAgEBCgEABCCfVEl83pSDSerk9I3pcICNTdzc5N3u4jt21cXdzBuJjgQAMGm_hT0IAgYBXtPjz6C_hUVZBFcwVTEvMC0EKGNvbS5hbmRyb2lkLmtleXN0b3JlLmFuZHJvaWRrZXlzdG9yZWRlbW8CAQExIgQgdM_LUHSI9SkQhZHHpQWRnzJ3MvvB2ANSauqYAAbS2JgwMqEFMQMCAQKiAwIBA6MEAgIBAKUFMQMCAQSqAwIBAb-DeAMCAQK_hT4DAgEAv4U_AgUAMB8GA1UdIwQYMBaAFKPSqizvDYzyJALVHLRgvL9qWyQUMAoGCCqGSM49BAMCA0cAMEQCIC7WHb2PyULnjp1M1TVI3Wti_eDhe6sFweuQAdecXtHhAiAS_eZkFsx_VNsrTu3XfZ2D7wIt-vT6nTljfHZ4zqU5xlkDGDCCAxQwggK6oAMCAQICAQIwCgYIKoZIzj0EAwIwgdwxPTA7BgNVBAMMNEZBS0UgQW5kcm9pZCBLZXlzdG9yZSBTb2Z0d2FyZSBBdHRlc3RhdGlvbiBSb290IEZBS0UxMTAvBgkqhkiG9w0BCQEWImNvbmZvcm1hbmNlLXRvb2xzQGZpZG9hbGxpYW5jZS5vcmcxFjAUBgNVBAoMDUZJRE8gQWxsaWFuY2UxIjAgBgNVBAsMGUF1dGhlbnRpY2F0b3IgQXR0ZXN0YXRpb24xCzAJBgNVBAYTAlVTMQswCQYDVQQIDAJNWTESMBAGA1UEBwwJV2FrZWZpZWxkMB4XDTE5MDQyNTA1NDkzMloXDTQ2MDkxMDA1NDkzMlowgeQxRTBDBgNVBAMMPEZBS0UgQW5kcm9pZCBLZXlzdG9yZSBTb2Z0d2FyZSBBdHRlc3RhdGlvbiBJbnRlcm1lZGlhdGUgRkFLRTExMC8GCSqGSIb3DQEJARYiY29uZm9ybWFuY2UtdG9vbHNAZmlkb2FsbGlhbmNlLm9yZzEWMBQGA1UECgwNRklETyBBbGxpYW5jZTEiMCAGA1UECwwZQXV0aGVudGljYXRvciBBdHRlc3RhdGlvbjELMAkGA1UEBhMCVVMxCzAJBgNVBAgMAk1ZMRIwEAYDVQQHDAlXYWtlZmllbGQwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASrUGErYk0Xu8O1GwRJOwVJC4wfi52883my3tygfFKh17YN0yF13Ct-3bwm2wjVX4b2cbaU3DBNpKKKjE4DpvXHo2MwYTAPBgNVHRMBAf8EBTADAQH_MA4GA1UdDwEB_wQEAwIChDAdBgNVHQ4EFgQUo9KqLO8NjPIkAtUctGC8v2pbJBQwHwYDVR0jBBgwFoAUUpobMuBWqs1RD-9fgDcGi_KRIx0wCgYIKoZIzj0EAwIDSAAwRQIhALFvLkAvtHrObTmN8P0-yLIT496P_weSEEbB6vCJWSh9AiBu-UOorCeLcF4WixOG9E5Li2nXe4uM2q6mbKGkll8u-WhhdXRoRGF0YVikPdxHEOnAiLIp26idVjIguzn3Ipr_RlsKZWsa-5qK-KBBAAAAYFUOS1SqR0CfmpUat2wTATEAIFedRhNvbRm4W8u7G4NXGf6i_FfJ46hLF6QJ8EAaG74MpQECAyYgASFYIG6jCBK7kChGPxWL9k1IUFsZRklcUb0yQNPqMAOWvu04Ilggr00SCNVfnf56uBenI9oVmsBH2nL7GfN-Mguf8rNvwio",
		  "clientDataJSON": "eyJvcmlnaW4iOiJodHRwczovL2Rldi5kb250bmVlZGEucHciLCJjaGFsbGVuZ2UiOiI0YWI3ZGZkMS1hNjk1LTQ3NzctOTg1Zi1hZDI5OTM4MjhlOTkiLCJ0eXBlIjoid2ViYXV0aG4uY3JlYXRlIn0"
		},
		"type": "public-key"
	  }`,
}
