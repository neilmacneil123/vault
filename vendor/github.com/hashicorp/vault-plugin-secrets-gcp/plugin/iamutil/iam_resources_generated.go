// THIS FILE IS AUTOGENERATED USING go generate. DO NOT EDIT.
package iamutil

func GetEnabledIamResources() EnabledResources {
	return &generatedIamResources{
		resources: generatedResources,
	}
}

var generatedServices = map[string]map[string]*ServiceConfig{
	"cloudiot": {
		"v1": &ServiceConfig{
			Name:               "cloudiot",
			Version:            "v1",
			IsPreferredVersion: true,
			RootUrl:            "https://cloudiot.googleapis.com/",
			ServicePath:        "",
		},
	},
	"cloudkms": {
		"v1": &ServiceConfig{
			Name:               "cloudkms",
			Version:            "v1",
			IsPreferredVersion: true,
			RootUrl:            "https://cloudkms.googleapis.com/",
			ServicePath:        "",
		},
	},
	"cloudresourcemanager": {
		"v1": &ServiceConfig{
			Name:               "cloudresourcemanager",
			Version:            "v1",
			IsPreferredVersion: true,
			RootUrl:            "https://cloudresourcemanager.googleapis.com/",
			ServicePath:        "",
		},
		"v1beta1": &ServiceConfig{
			Name:               "cloudresourcemanager",
			Version:            "v1beta1",
			IsPreferredVersion: false,
			RootUrl:            "https://cloudresourcemanager.googleapis.com/",
			ServicePath:        "",
		},
		"v2beta1": &ServiceConfig{
			Name:               "cloudresourcemanager",
			Version:            "v2beta1",
			IsPreferredVersion: false,
			RootUrl:            "https://cloudresourcemanager.googleapis.com/",
			ServicePath:        "",
		},
	},
	"cloudtasks": {
		"v2beta2": &ServiceConfig{
			Name:               "cloudtasks",
			Version:            "v2beta2",
			IsPreferredVersion: true,
			RootUrl:            "https://cloudtasks.googleapis.com/",
			ServicePath:        "",
		},
	},
	"clouduseraccounts": {
		"alpha": &ServiceConfig{
			Name:               "clouduseraccounts",
			Version:            "alpha",
			IsPreferredVersion: false,
			RootUrl:            "https://www.googleapis.com/",
			ServicePath:        "clouduseraccounts/alpha/projects/",
		},
		"vm_alpha": &ServiceConfig{
			Name:               "clouduseraccounts",
			Version:            "vm_alpha",
			IsPreferredVersion: true,
			RootUrl:            "https://www.googleapis.com/",
			ServicePath:        "clouduseraccounts/vm_alpha/projects/",
		},
	},
	"compute": {
		"alpha": &ServiceConfig{
			Name:               "compute",
			Version:            "alpha",
			IsPreferredVersion: false,
			RootUrl:            "https://www.googleapis.com/",
			ServicePath:        "compute/alpha/projects/",
		},
		"beta": &ServiceConfig{
			Name:               "compute",
			Version:            "beta",
			IsPreferredVersion: false,
			RootUrl:            "https://www.googleapis.com/",
			ServicePath:        "compute/beta/projects/",
		},
	},
	"dataproc": {
		"v1beta2": &ServiceConfig{
			Name:               "dataproc",
			Version:            "v1beta2",
			IsPreferredVersion: false,
			RootUrl:            "https://dataproc.googleapis.com/",
			ServicePath:        "",
		},
	},
	"deploymentmanager": {
		"alpha": &ServiceConfig{
			Name:               "deploymentmanager",
			Version:            "alpha",
			IsPreferredVersion: false,
			RootUrl:            "https://www.googleapis.com/",
			ServicePath:        "deploymentmanager/alpha/projects/",
		},
		"v2": &ServiceConfig{
			Name:               "deploymentmanager",
			Version:            "v2",
			IsPreferredVersion: true,
			RootUrl:            "https://www.googleapis.com/",
			ServicePath:        "deploymentmanager/v2/projects/",
		},
		"v2beta": &ServiceConfig{
			Name:               "deploymentmanager",
			Version:            "v2beta",
			IsPreferredVersion: false,
			RootUrl:            "https://www.googleapis.com/",
			ServicePath:        "deploymentmanager/v2beta/projects/",
		},
	},
	"genomics": {
		"v1": &ServiceConfig{
			Name:               "genomics",
			Version:            "v1",
			IsPreferredVersion: true,
			RootUrl:            "https://genomics.googleapis.com/",
			ServicePath:        "",
		},
	},
	"iam": {
		"v1": &ServiceConfig{
			Name:               "iam",
			Version:            "v1",
			IsPreferredVersion: true,
			RootUrl:            "https://iam.googleapis.com/",
			ServicePath:        "",
		},
	},
	"ml": {
		"v1": &ServiceConfig{
			Name:               "ml",
			Version:            "v1",
			IsPreferredVersion: true,
			RootUrl:            "https://ml.googleapis.com/",
			ServicePath:        "",
		},
	},
	"pubsub": {
		"v1": &ServiceConfig{
			Name:               "pubsub",
			Version:            "v1",
			IsPreferredVersion: true,
			RootUrl:            "https://pubsub.googleapis.com/",
			ServicePath:        "",
		},
		"v1beta2": &ServiceConfig{
			Name:               "pubsub",
			Version:            "v1beta2",
			IsPreferredVersion: false,
			RootUrl:            "https://pubsub.googleapis.com/",
			ServicePath:        "",
		},
	},
	"runtimeconfig": {
		"v1beta1": &ServiceConfig{
			Name:               "runtimeconfig",
			Version:            "v1beta1",
			IsPreferredVersion: false,
			RootUrl:            "https://runtimeconfig.googleapis.com/",
			ServicePath:        "",
		},
	},
	"servicemanagement": {
		"v1": &ServiceConfig{
			Name:               "servicemanagement",
			Version:            "v1",
			IsPreferredVersion: true,
			RootUrl:            "https://servicemanagement.googleapis.com/",
			ServicePath:        "",
		},
	},
	"sourcerepo": {
		"v1": &ServiceConfig{
			Name:               "sourcerepo",
			Version:            "v1",
			IsPreferredVersion: true,
			RootUrl:            "https://sourcerepo.googleapis.com/",
			ServicePath:        "",
		},
	},
	"spanner": {
		"v1": &ServiceConfig{
			Name:               "spanner",
			Version:            "v1",
			IsPreferredVersion: true,
			RootUrl:            "https://spanner.googleapis.com/",
			ServicePath:        "",
		},
	},
	"storage": {
		"v1": &ServiceConfig{
			Name:               "storage",
			Version:            "v1",
			IsPreferredVersion: true,
			RootUrl:            "https://www.googleapis.com/",
			ServicePath:        "storage/v1/",
		},
	},
}

var generatedResources = map[string]map[string]map[string]*IamResourceConfig{
	"b": {
		"storage": {
			"v1": &IamResourceConfig{
				Service: generatedServices["storage"]["v1"],
				SetIamPolicy: &HttpMethodCfg{
					HttpMethod: "PUT",
					Path:       "b/{bucket}/iam",
					ReplacementKeys: map[string]string{
						"b": "bucket",
					},
				},
				GetIamPolicy: &HttpMethodCfg{
					HttpMethod: "GET",
					Path:       "b/{bucket}/iam",
					ReplacementKeys: map[string]string{
						"b": "bucket",
					},
				},
			},
		},
	},
	"b/o": {
		"storage": {
			"v1": &IamResourceConfig{
				Service: generatedServices["storage"]["v1"],
				SetIamPolicy: &HttpMethodCfg{
					HttpMethod: "PUT",
					Path:       "b/{bucket}/o/{object}/iam",
					ReplacementKeys: map[string]string{
						"b": "bucket",
						"o": "object",
					},
				},
				GetIamPolicy: &HttpMethodCfg{
					HttpMethod: "GET",
					Path:       "b/{bucket}/o/{object}/iam",
					ReplacementKeys: map[string]string{
						"b": "bucket",
						"o": "object",
					},
				},
			},
		},
	},
	"buckets": {
		"storage": {
			"v1": &IamResourceConfig{
				Service: generatedServices["storage"]["v1"],
				SetIamPolicy: &HttpMethodCfg{
					HttpMethod: "PUT",
					Path:       "b/{bucket}/iam",
					ReplacementKeys: map[string]string{
						"buckets": "bucket",
					},
				},
				GetIamPolicy: &HttpMethodCfg{
					HttpMethod: "GET",
					Path:       "b/{bucket}/iam",
					ReplacementKeys: map[string]string{
						"buckets": "bucket",
					},
				},
			},
		},
	},
	"buckets/objects": {
		"storage": {
			"v1": &IamResourceConfig{
				Service: generatedServices["storage"]["v1"],
				SetIamPolicy: &HttpMethodCfg{
					HttpMethod: "PUT",
					Path:       "b/{bucket}/o/{object}/iam",
					ReplacementKeys: map[string]string{
						"buckets": "bucket",
						"objects": "object",
					},
				},
				GetIamPolicy: &HttpMethodCfg{
					HttpMethod: "GET",
					Path:       "b/{bucket}/o/{object}/iam",
					ReplacementKeys: map[string]string{
						"buckets": "bucket",
						"objects": "object",
					},
				},
			},
		},
	},
	"datasets": {
		"genomics": {
			"v1": &IamResourceConfig{
				Service: generatedServices["genomics"]["v1"],
				SetIamPolicy: &HttpMethodCfg{
					HttpMethod: "POST",
					Path:       "v1/datasets/{datasetsId}:setIamPolicy",
					ReplacementKeys: map[string]string{
						"datasets": "datasetsId",
					},
				},
				GetIamPolicy: &HttpMethodCfg{
					HttpMethod: "POST",
					Path:       "v1/datasets/{datasetsId}:getIamPolicy",
					ReplacementKeys: map[string]string{
						"datasets": "datasetsId",
					},
				},
			},
		},
	},
	"folders": {
		"cloudresourcemanager": {
			"v2beta1": &IamResourceConfig{
				Service: generatedServices["cloudresourcemanager"]["v2beta1"],
				SetIamPolicy: &HttpMethodCfg{
					HttpMethod: "POST",
					Path:       "v2beta1/folders/{foldersId}:setIamPolicy",
					ReplacementKeys: map[string]string{
						"folders": "foldersId",
					},
				},
				GetIamPolicy: &HttpMethodCfg{
					HttpMethod: "POST",
					Path:       "v2beta1/folders/{foldersId}:getIamPolicy",
					ReplacementKeys: map[string]string{
						"folders": "foldersId",
					},
				},
			},
		},
	},
	"organizations": {
		"cloudresourcemanager": {
			"v1": &IamResourceConfig{
				Service: generatedServices["cloudresourcemanager"]["v1"],
				SetIamPolicy: &HttpMethodCfg{
					HttpMethod: "POST",
					Path:       "v1/organizations/{organizationsId}:setIamPolicy",
					ReplacementKeys: map[string]string{
						"organizations": "organizationsId",
					},
				},
				GetIamPolicy: &HttpMethodCfg{
					HttpMethod: "POST",
					Path:       "v1/organizations/{organizationsId}:getIamPolicy",
					ReplacementKeys: map[string]string{
						"organizations": "organizationsId",
					},
				},
			},
			"v1beta1": &IamResourceConfig{
				Service: generatedServices["cloudresourcemanager"]["v1beta1"],
				SetIamPolicy: &HttpMethodCfg{
					HttpMethod: "POST",
					Path:       "v1beta1/organizations/{organizationsId}:setIamPolicy",
					ReplacementKeys: map[string]string{
						"organizations": "organizationsId",
					},
				},
				GetIamPolicy: &HttpMethodCfg{
					HttpMethod: "POST",
					Path:       "v1beta1/organizations/{organizationsId}:getIamPolicy",
					ReplacementKeys: map[string]string{
						"organizations": "organizationsId",
					},
				},
			},
		},
	},
	"projects": {
		"cloudresourcemanager": {
			"v1": &IamResourceConfig{
				Service: generatedServices["cloudresourcemanager"]["v1"],
				SetIamPolicy: &HttpMethodCfg{
					HttpMethod: "POST",
					Path:       "v1/projects/{resource}:setIamPolicy",
					ReplacementKeys: map[string]string{
						"projects": "resource",
					},
				},
				GetIamPolicy: &HttpMethodCfg{
					HttpMethod: "POST",
					Path:       "v1/projects/{resource}:getIamPolicy",
					ReplacementKeys: map[string]string{
						"projects": "resource",
					},
				},
			},
			"v1beta1": &IamResourceConfig{
				Service: generatedServices["cloudresourcemanager"]["v1beta1"],
				SetIamPolicy: &HttpMethodCfg{
					HttpMethod: "POST",
					Path:       "v1beta1/projects/{resource}:setIamPolicy",
					ReplacementKeys: map[string]string{
						"projects": "resource",
					},
				},
				GetIamPolicy: &HttpMethodCfg{
					HttpMethod: "POST",
					Path:       "v1beta1/projects/{resource}:getIamPolicy",
					ReplacementKeys: map[string]string{
						"projects": "resource",
					},
				},
			},
		},
	},
	"projects/backendBuckets": {
		"compute": {
			"alpha": &IamResourceConfig{
				Service: generatedServices["compute"]["alpha"],
				SetIamPolicy: &HttpMethodCfg{
					HttpMethod: "POST",
					Path:       "{project}/global/backendBuckets/{resource}/setIamPolicy",
					ReplacementKeys: map[string]string{
						"backendBuckets": "resource",
						"projects":       "project",
					},
				},
				GetIamPolicy: &HttpMethodCfg{
					HttpMethod: "GET",
					Path:       "{project}/global/backendBuckets/{resource}/getIamPolicy",
					ReplacementKeys: map[string]string{
						"backendBuckets": "resource",
						"projects":       "project",
					},
				},
			},
		},
	},
	"projects/configs": {
		"runtimeconfig": {
			"v1beta1": &IamResourceConfig{
				Service: generatedServices["runtimeconfig"]["v1beta1"],
				SetIamPolicy: &HttpMethodCfg{
					HttpMethod: "POST",
					Path:       "v1beta1/projects/{projectsId}/configs/{configsId}:setIamPolicy",
					ReplacementKeys: map[string]string{
						"configs":  "configsId",
						"projects": "projectsId",
					},
				},
				GetIamPolicy: &HttpMethodCfg{
					HttpMethod: "GET",
					Path:       "v1beta1/projects/{projectsId}/configs/{configsId}:getIamPolicy",
					ReplacementKeys: map[string]string{
						"configs":  "configsId",
						"projects": "projectsId",
					},
				},
			},
		},
	},
	"projects/deployments": {
		"deploymentmanager": {
			"alpha": &IamResourceConfig{
				Service: generatedServices["deploymentmanager"]["alpha"],
				SetIamPolicy: &HttpMethodCfg{
					HttpMethod: "POST",
					Path:       "{project}/global/deployments/{resource}/setIamPolicy",
					ReplacementKeys: map[string]string{
						"deployments": "resource",
						"projects":    "project",
					},
				},
				GetIamPolicy: &HttpMethodCfg{
					HttpMethod: "GET",
					Path:       "{project}/global/deployments/{resource}/getIamPolicy",
					ReplacementKeys: map[string]string{
						"deployments": "resource",
						"projects":    "project",
					},
				},
			},
			"v2": &IamResourceConfig{
				Service: generatedServices["deploymentmanager"]["v2"],
				SetIamPolicy: &HttpMethodCfg{
					HttpMethod: "POST",
					Path:       "{project}/global/deployments/{resource}/setIamPolicy",
					ReplacementKeys: map[string]string{
						"deployments": "resource",
						"projects":    "project",
					},
				},
				GetIamPolicy: &HttpMethodCfg{
					HttpMethod: "GET",
					Path:       "{project}/global/deployments/{resource}/getIamPolicy",
					ReplacementKeys: map[string]string{
						"deployments": "resource",
						"projects":    "project",
					},
				},
			},
			"v2beta": &IamResourceConfig{
				Service: generatedServices["deploymentmanager"]["v2beta"],
				SetIamPolicy: &HttpMethodCfg{
					HttpMethod: "POST",
					Path:       "{project}/global/deployments/{resource}/setIamPolicy",
					ReplacementKeys: map[string]string{
						"deployments": "resource",
						"projects":    "project",
					},
				},
				GetIamPolicy: &HttpMethodCfg{
					HttpMethod: "GET",
					Path:       "{project}/global/deployments/{resource}/getIamPolicy",
					ReplacementKeys: map[string]string{
						"deployments": "resource",
						"projects":    "project",
					},
				},
			},
		},
	},
	"projects/groups": {
		"clouduseraccounts": {
			"alpha": &IamResourceConfig{
				Service: generatedServices["clouduseraccounts"]["alpha"],
				SetIamPolicy: &HttpMethodCfg{
					HttpMethod: "POST",
					Path:       "{project}/global/groups/{resource}/setIamPolicy",
					ReplacementKeys: map[string]string{
						"groups":   "resource",
						"projects": "project",
					},
				},
				GetIamPolicy: &HttpMethodCfg{
					HttpMethod: "GET",
					Path:       "{project}/global/groups/{resource}/getIamPolicy",
					ReplacementKeys: map[string]string{
						"groups":   "resource",
						"projects": "project",
					},
				},
			},
			"vm_alpha": &IamResourceConfig{
				Service: generatedServices["clouduseraccounts"]["vm_alpha"],
				SetIamPolicy: &HttpMethodCfg{
					HttpMethod: "POST",
					Path:       "{project}/global/groups/{resource}/setIamPolicy",
					ReplacementKeys: map[string]string{
						"groups":   "resource",
						"projects": "project",
					},
				},
				GetIamPolicy: &HttpMethodCfg{
					HttpMethod: "GET",
					Path:       "{project}/global/groups/{resource}/getIamPolicy",
					ReplacementKeys: map[string]string{
						"groups":   "resource",
						"projects": "project",
					},
				},
			},
		},
	},
	"projects/images": {
		"compute": {
			"alpha": &IamResourceConfig{
				Service: generatedServices["compute"]["alpha"],
				SetIamPolicy: &HttpMethodCfg{
					HttpMethod: "POST",
					Path:       "{project}/global/images/{resource}/setIamPolicy",
					ReplacementKeys: map[string]string{
						"images":   "resource",
						"projects": "project",
					},
				},
				GetIamPolicy: &HttpMethodCfg{
					HttpMethod: "GET",
					Path:       "{project}/global/images/{resource}/getIamPolicy",
					ReplacementKeys: map[string]string{
						"images":   "resource",
						"projects": "project",
					},
				},
			},
		},
	},
	"projects/instances": {
		"spanner": {
			"v1": &IamResourceConfig{
				Service: generatedServices["spanner"]["v1"],
				SetIamPolicy: &HttpMethodCfg{
					HttpMethod: "POST",
					Path:       "v1/projects/{projectsId}/instances/{instancesId}:setIamPolicy",
					ReplacementKeys: map[string]string{
						"instances": "instancesId",
						"projects":  "projectsId",
					},
				},
				GetIamPolicy: &HttpMethodCfg{
					HttpMethod: "POST",
					Path:       "v1/projects/{projectsId}/instances/{instancesId}:getIamPolicy",
					ReplacementKeys: map[string]string{
						"instances": "instancesId",
						"projects":  "projectsId",
					},
				},
			},
		},
	},
	"projects/instances/databases": {
		"spanner": {
			"v1": &IamResourceConfig{
				Service: generatedServices["spanner"]["v1"],
				SetIamPolicy: &HttpMethodCfg{
					HttpMethod: "POST",
					Path:       "v1/projects/{projectsId}/instances/{instancesId}/databases/{databasesId}:setIamPolicy",
					ReplacementKeys: map[string]string{
						"databases": "databasesId",
						"instances": "instancesId",
						"projects":  "projectsId",
					},
				},
				GetIamPolicy: &HttpMethodCfg{
					HttpMethod: "POST",
					Path:       "v1/projects/{projectsId}/instances/{instancesId}/databases/{databasesId}:getIamPolicy",
					ReplacementKeys: map[string]string{
						"databases": "databasesId",
						"instances": "instancesId",
						"projects":  "projectsId",
					},
				},
			},
		},
	},
	"projects/interconnects": {
		"compute": {
			"alpha": &IamResourceConfig{
				Service: generatedServices["compute"]["alpha"],
				SetIamPolicy: &HttpMethodCfg{
					HttpMethod: "POST",
					Path:       "{project}/global/interconnects/{resource}/setIamPolicy",
					ReplacementKeys: map[string]string{
						"interconnects": "resource",
						"projects":      "project",
					},
				},
				GetIamPolicy: &HttpMethodCfg{
					HttpMethod: "GET",
					Path:       "{project}/global/interconnects/{resource}/getIamPolicy",
					ReplacementKeys: map[string]string{
						"interconnects": "resource",
						"projects":      "project",
					},
				},
			},
		},
	},
	"projects/jobs": {
		"ml": {
			"v1": &IamResourceConfig{
				Service: generatedServices["ml"]["v1"],
				SetIamPolicy: &HttpMethodCfg{
					HttpMethod: "POST",
					Path:       "v1/projects/{projectsId}/jobs/{jobsId}:setIamPolicy",
					ReplacementKeys: map[string]string{
						"jobs":     "jobsId",
						"projects": "projectsId",
					},
				},
				GetIamPolicy: &HttpMethodCfg{
					HttpMethod: "GET",
					Path:       "v1/projects/{projectsId}/jobs/{jobsId}:getIamPolicy",
					ReplacementKeys: map[string]string{
						"jobs":     "jobsId",
						"projects": "projectsId",
					},
				},
			},
		},
	},
	"projects/licenseCodes": {
		"compute": {
			"alpha": &IamResourceConfig{
				Service: generatedServices["compute"]["alpha"],
				SetIamPolicy: &HttpMethodCfg{
					HttpMethod: "POST",
					Path:       "{project}/global/licenseCodes/{resource}/setIamPolicy",
					ReplacementKeys: map[string]string{
						"licenseCodes": "resource",
						"projects":     "project",
					},
				},
				GetIamPolicy: &HttpMethodCfg{
					HttpMethod: "GET",
					Path:       "{project}/global/licenseCodes/{resource}/getIamPolicy",
					ReplacementKeys: map[string]string{
						"licenseCodes": "resource",
						"projects":     "project",
					},
				},
			},
		},
	},
	"projects/licenses": {
		"compute": {
			"alpha": &IamResourceConfig{
				Service: generatedServices["compute"]["alpha"],
				SetIamPolicy: &HttpMethodCfg{
					HttpMethod: "POST",
					Path:       "{project}/global/licenses/{resource}/setIamPolicy",
					ReplacementKeys: map[string]string{
						"licenses": "resource",
						"projects": "project",
					},
				},
				GetIamPolicy: &HttpMethodCfg{
					HttpMethod: "GET",
					Path:       "{project}/global/licenses/{resource}/getIamPolicy",
					ReplacementKeys: map[string]string{
						"licenses": "resource",
						"projects": "project",
					},
				},
			},
		},
	},
	"projects/locations/keyRings": {
		"cloudkms": {
			"v1": &IamResourceConfig{
				Service: generatedServices["cloudkms"]["v1"],
				SetIamPolicy: &HttpMethodCfg{
					HttpMethod: "POST",
					Path:       "v1/projects/{projectsId}/locations/{locationsId}/keyRings/{keyRingsId}:setIamPolicy",
					ReplacementKeys: map[string]string{
						"keyRings":  "keyRingsId",
						"locations": "locationsId",
						"projects":  "projectsId",
					},
				},
				GetIamPolicy: &HttpMethodCfg{
					HttpMethod: "GET",
					Path:       "v1/projects/{projectsId}/locations/{locationsId}/keyRings/{keyRingsId}:getIamPolicy",
					ReplacementKeys: map[string]string{
						"keyRings":  "keyRingsId",
						"locations": "locationsId",
						"projects":  "projectsId",
					},
				},
			},
		},
	},
	"projects/locations/keyRings/cryptoKeys": {
		"cloudkms": {
			"v1": &IamResourceConfig{
				Service: generatedServices["cloudkms"]["v1"],
				SetIamPolicy: &HttpMethodCfg{
					HttpMethod: "POST",
					Path:       "v1/projects/{projectsId}/locations/{locationsId}/keyRings/{keyRingsId}/cryptoKeys/{cryptoKeysId}:setIamPolicy",
					ReplacementKeys: map[string]string{
						"cryptoKeys": "cryptoKeysId",
						"keyRings":   "keyRingsId",
						"locations":  "locationsId",
						"projects":   "projectsId",
					},
				},
				GetIamPolicy: &HttpMethodCfg{
					HttpMethod: "GET",
					Path:       "v1/projects/{projectsId}/locations/{locationsId}/keyRings/{keyRingsId}/cryptoKeys/{cryptoKeysId}:getIamPolicy",
					ReplacementKeys: map[string]string{
						"cryptoKeys": "cryptoKeysId",
						"keyRings":   "keyRingsId",
						"locations":  "locationsId",
						"projects":   "projectsId",
					},
				},
			},
		},
	},
	"projects/locations/queues": {
		"cloudtasks": {
			"v2beta2": &IamResourceConfig{
				Service: generatedServices["cloudtasks"]["v2beta2"],
				SetIamPolicy: &HttpMethodCfg{
					HttpMethod: "POST",
					Path:       "v2beta2/projects/{projectsId}/locations/{locationsId}/queues/{queuesId}:setIamPolicy",
					ReplacementKeys: map[string]string{
						"locations": "locationsId",
						"projects":  "projectsId",
						"queues":    "queuesId",
					},
				},
				GetIamPolicy: &HttpMethodCfg{
					HttpMethod: "POST",
					Path:       "v2beta2/projects/{projectsId}/locations/{locationsId}/queues/{queuesId}:getIamPolicy",
					ReplacementKeys: map[string]string{
						"locations": "locationsId",
						"projects":  "projectsId",
						"queues":    "queuesId",
					},
				},
			},
		},
	},
	"projects/locations/registries": {
		"cloudiot": {
			"v1": &IamResourceConfig{
				Service: generatedServices["cloudiot"]["v1"],
				SetIamPolicy: &HttpMethodCfg{
					HttpMethod: "POST",
					Path:       "v1/projects/{projectsId}/locations/{locationsId}/registries/{registriesId}:setIamPolicy",
					ReplacementKeys: map[string]string{
						"locations":  "locationsId",
						"projects":   "projectsId",
						"registries": "registriesId",
					},
				},
				GetIamPolicy: &HttpMethodCfg{
					HttpMethod: "POST",
					Path:       "v1/projects/{projectsId}/locations/{locationsId}/registries/{registriesId}:getIamPolicy",
					ReplacementKeys: map[string]string{
						"locations":  "locationsId",
						"projects":   "projectsId",
						"registries": "registriesId",
					},
				},
			},
		},
	},
	"projects/models": {
		"ml": {
			"v1": &IamResourceConfig{
				Service: generatedServices["ml"]["v1"],
				SetIamPolicy: &HttpMethodCfg{
					HttpMethod: "POST",
					Path:       "v1/projects/{projectsId}/models/{modelsId}:setIamPolicy",
					ReplacementKeys: map[string]string{
						"models":   "modelsId",
						"projects": "projectsId",
					},
				},
				GetIamPolicy: &HttpMethodCfg{
					HttpMethod: "GET",
					Path:       "v1/projects/{projectsId}/models/{modelsId}:getIamPolicy",
					ReplacementKeys: map[string]string{
						"models":   "modelsId",
						"projects": "projectsId",
					},
				},
			},
		},
	},
	"projects/regions/clusters": {
		"dataproc": {
			"v1beta2": &IamResourceConfig{
				Service: generatedServices["dataproc"]["v1beta2"],
				SetIamPolicy: &HttpMethodCfg{
					HttpMethod: "POST",
					Path:       "v1beta2/projects/{projectsId}/regions/{regionsId}/clusters/{clustersId}:setIamPolicy",
					ReplacementKeys: map[string]string{
						"clusters": "clustersId",
						"projects": "projectsId",
						"regions":  "regionsId",
					},
				},
				GetIamPolicy: &HttpMethodCfg{
					HttpMethod: "GET",
					Path:       "v1beta2/projects/{projectsId}/regions/{regionsId}/clusters/{clustersId}:getIamPolicy",
					ReplacementKeys: map[string]string{
						"clusters": "clustersId",
						"projects": "projectsId",
						"regions":  "regionsId",
					},
				},
			},
		},
	},
	"projects/regions/interconnectAttachments": {
		"compute": {
			"alpha": &IamResourceConfig{
				Service: generatedServices["compute"]["alpha"],
				SetIamPolicy: &HttpMethodCfg{
					HttpMethod: "POST",
					Path:       "{project}/regions/{region}/interconnectAttachments/{resource}/setIamPolicy",
					ReplacementKeys: map[string]string{
						"interconnectAttachments": "resource",
						"projects":                "project",
						"regions":                 "region",
					},
				},
				GetIamPolicy: &HttpMethodCfg{
					HttpMethod: "GET",
					Path:       "{project}/regions/{region}/interconnectAttachments/{resource}/getIamPolicy",
					ReplacementKeys: map[string]string{
						"interconnectAttachments": "resource",
						"projects":                "project",
						"regions":                 "region",
					},
				},
			},
		},
	},
	"projects/regions/maintenancePolicies": {
		"compute": {
			"alpha": &IamResourceConfig{
				Service: generatedServices["compute"]["alpha"],
				SetIamPolicy: &HttpMethodCfg{
					HttpMethod: "POST",
					Path:       "{project}/regions/{region}/maintenancePolicies/{resource}/setIamPolicy",
					ReplacementKeys: map[string]string{
						"maintenancePolicies": "resource",
						"projects":            "project",
						"regions":             "region",
					},
				},
				GetIamPolicy: &HttpMethodCfg{
					HttpMethod: "GET",
					Path:       "{project}/regions/{region}/maintenancePolicies/{resource}/getIamPolicy",
					ReplacementKeys: map[string]string{
						"maintenancePolicies": "resource",
						"projects":            "project",
						"regions":             "region",
					},
				},
			},
		},
	},
	"projects/regions/nodeTemplates": {
		"compute": {
			"alpha": &IamResourceConfig{
				Service: generatedServices["compute"]["alpha"],
				SetIamPolicy: &HttpMethodCfg{
					HttpMethod: "POST",
					Path:       "{project}/regions/{region}/nodeTemplates/{resource}/setIamPolicy",
					ReplacementKeys: map[string]string{
						"nodeTemplates": "resource",
						"projects":      "project",
						"regions":       "region",
					},
				},
				GetIamPolicy: &HttpMethodCfg{
					HttpMethod: "GET",
					Path:       "{project}/regions/{region}/nodeTemplates/{resource}/getIamPolicy",
					ReplacementKeys: map[string]string{
						"nodeTemplates": "resource",
						"projects":      "project",
						"regions":       "region",
					},
				},
			},
		},
	},
	"projects/regions/subnetworks": {
		"compute": {
			"alpha": &IamResourceConfig{
				Service: generatedServices["compute"]["alpha"],
				SetIamPolicy: &HttpMethodCfg{
					HttpMethod: "POST",
					Path:       "{project}/regions/{region}/subnetworks/{resource}/setIamPolicy",
					ReplacementKeys: map[string]string{
						"projects":    "project",
						"regions":     "region",
						"subnetworks": "resource",
					},
				},
				GetIamPolicy: &HttpMethodCfg{
					HttpMethod: "GET",
					Path:       "{project}/regions/{region}/subnetworks/{resource}/getIamPolicy",
					ReplacementKeys: map[string]string{
						"projects":    "project",
						"regions":     "region",
						"subnetworks": "resource",
					},
				},
			},
			"beta": &IamResourceConfig{
				Service: generatedServices["compute"]["beta"],
				SetIamPolicy: &HttpMethodCfg{
					HttpMethod: "POST",
					Path:       "{project}/regions/{region}/subnetworks/{resource}/setIamPolicy",
					ReplacementKeys: map[string]string{
						"projects":    "project",
						"regions":     "region",
						"subnetworks": "resource",
					},
				},
				GetIamPolicy: &HttpMethodCfg{
					HttpMethod: "GET",
					Path:       "{project}/regions/{region}/subnetworks/{resource}/getIamPolicy",
					ReplacementKeys: map[string]string{
						"projects":    "project",
						"regions":     "region",
						"subnetworks": "resource",
					},
				},
			},
		},
	},
	"projects/repos": {
		"sourcerepo": {
			"v1": &IamResourceConfig{
				Service: generatedServices["sourcerepo"]["v1"],
				SetIamPolicy: &HttpMethodCfg{
					HttpMethod: "POST",
					Path:       "v1/projects/{projectsId}/repos/{reposId}:setIamPolicy",
					ReplacementKeys: map[string]string{
						"projects": "projectsId",
						"repos":    "reposId",
					},
				},
				GetIamPolicy: &HttpMethodCfg{
					HttpMethod: "GET",
					Path:       "v1/projects/{projectsId}/repos/{reposId}:getIamPolicy",
					ReplacementKeys: map[string]string{
						"projects": "projectsId",
						"repos":    "reposId",
					},
				},
			},
		},
	},
	"projects/serviceAccounts": {
		"iam": {
			"v1": &IamResourceConfig{
				Service: generatedServices["iam"]["v1"],
				SetIamPolicy: &HttpMethodCfg{
					HttpMethod: "POST",
					Path:       "v1/projects/{projectsId}/serviceAccounts/{serviceAccountsId}:setIamPolicy",
					ReplacementKeys: map[string]string{
						"projects":        "projectsId",
						"serviceAccounts": "serviceAccountsId",
					},
				},
				GetIamPolicy: &HttpMethodCfg{
					HttpMethod: "POST",
					Path:       "v1/projects/{projectsId}/serviceAccounts/{serviceAccountsId}:getIamPolicy",
					ReplacementKeys: map[string]string{
						"projects":        "projectsId",
						"serviceAccounts": "serviceAccountsId",
					},
				},
			},
		},
	},
	"projects/snapshots": {
		"compute": {
			"alpha": &IamResourceConfig{
				Service: generatedServices["compute"]["alpha"],
				SetIamPolicy: &HttpMethodCfg{
					HttpMethod: "POST",
					Path:       "{project}/global/snapshots/{resource}/setIamPolicy",
					ReplacementKeys: map[string]string{
						"projects":  "project",
						"snapshots": "resource",
					},
				},
				GetIamPolicy: &HttpMethodCfg{
					HttpMethod: "GET",
					Path:       "{project}/global/snapshots/{resource}/getIamPolicy",
					ReplacementKeys: map[string]string{
						"projects":  "project",
						"snapshots": "resource",
					},
				},
			},
		},
		"pubsub": {
			"v1": &IamResourceConfig{
				Service: generatedServices["pubsub"]["v1"],
				SetIamPolicy: &HttpMethodCfg{
					HttpMethod: "POST",
					Path:       "v1/projects/{projectsId}/snapshots/{snapshotsId}:setIamPolicy",
					ReplacementKeys: map[string]string{
						"projects":  "projectsId",
						"snapshots": "snapshotsId",
					},
				},
				GetIamPolicy: &HttpMethodCfg{
					HttpMethod: "GET",
					Path:       "v1/projects/{projectsId}/snapshots/{snapshotsId}:getIamPolicy",
					ReplacementKeys: map[string]string{
						"projects":  "projectsId",
						"snapshots": "snapshotsId",
					},
				},
			},
		},
	},
	"projects/subscriptions": {
		"pubsub": {
			"v1": &IamResourceConfig{
				Service: generatedServices["pubsub"]["v1"],
				SetIamPolicy: &HttpMethodCfg{
					HttpMethod: "POST",
					Path:       "v1/projects/{projectsId}/subscriptions/{subscriptionsId}:setIamPolicy",
					ReplacementKeys: map[string]string{
						"projects":      "projectsId",
						"subscriptions": "subscriptionsId",
					},
				},
				GetIamPolicy: &HttpMethodCfg{
					HttpMethod: "GET",
					Path:       "v1/projects/{projectsId}/subscriptions/{subscriptionsId}:getIamPolicy",
					ReplacementKeys: map[string]string{
						"projects":      "projectsId",
						"subscriptions": "subscriptionsId",
					},
				},
			},
			"v1beta2": &IamResourceConfig{
				Service: generatedServices["pubsub"]["v1beta2"],
				SetIamPolicy: &HttpMethodCfg{
					HttpMethod: "POST",
					Path:       "v1beta2/projects/{projectsId}/subscriptions/{subscriptionsId}:setIamPolicy",
					ReplacementKeys: map[string]string{
						"projects":      "projectsId",
						"subscriptions": "subscriptionsId",
					},
				},
				GetIamPolicy: &HttpMethodCfg{
					HttpMethod: "GET",
					Path:       "v1beta2/projects/{projectsId}/subscriptions/{subscriptionsId}:getIamPolicy",
					ReplacementKeys: map[string]string{
						"projects":      "projectsId",
						"subscriptions": "subscriptionsId",
					},
				},
			},
		},
	},
	"projects/topics": {
		"pubsub": {
			"v1": &IamResourceConfig{
				Service: generatedServices["pubsub"]["v1"],
				SetIamPolicy: &HttpMethodCfg{
					HttpMethod: "POST",
					Path:       "v1/projects/{projectsId}/topics/{topicsId}:setIamPolicy",
					ReplacementKeys: map[string]string{
						"projects": "projectsId",
						"topics":   "topicsId",
					},
				},
				GetIamPolicy: &HttpMethodCfg{
					HttpMethod: "GET",
					Path:       "v1/projects/{projectsId}/topics/{topicsId}:getIamPolicy",
					ReplacementKeys: map[string]string{
						"projects": "projectsId",
						"topics":   "topicsId",
					},
				},
			},
			"v1beta2": &IamResourceConfig{
				Service: generatedServices["pubsub"]["v1beta2"],
				SetIamPolicy: &HttpMethodCfg{
					HttpMethod: "POST",
					Path:       "v1beta2/projects/{projectsId}/topics/{topicsId}:setIamPolicy",
					ReplacementKeys: map[string]string{
						"projects": "projectsId",
						"topics":   "topicsId",
					},
				},
				GetIamPolicy: &HttpMethodCfg{
					HttpMethod: "GET",
					Path:       "v1beta2/projects/{projectsId}/topics/{topicsId}:getIamPolicy",
					ReplacementKeys: map[string]string{
						"projects": "projectsId",
						"topics":   "topicsId",
					},
				},
			},
		},
	},
	"projects/users": {
		"clouduseraccounts": {
			"alpha": &IamResourceConfig{
				Service: generatedServices["clouduseraccounts"]["alpha"],
				SetIamPolicy: &HttpMethodCfg{
					HttpMethod: "POST",
					Path:       "{project}/global/users/{resource}/setIamPolicy",
					ReplacementKeys: map[string]string{
						"projects": "project",
						"users":    "resource",
					},
				},
				GetIamPolicy: &HttpMethodCfg{
					HttpMethod: "GET",
					Path:       "{project}/global/users/{resource}/getIamPolicy",
					ReplacementKeys: map[string]string{
						"projects": "project",
						"users":    "resource",
					},
				},
			},
			"vm_alpha": &IamResourceConfig{
				Service: generatedServices["clouduseraccounts"]["vm_alpha"],
				SetIamPolicy: &HttpMethodCfg{
					HttpMethod: "POST",
					Path:       "{project}/global/users/{resource}/setIamPolicy",
					ReplacementKeys: map[string]string{
						"projects": "project",
						"users":    "resource",
					},
				},
				GetIamPolicy: &HttpMethodCfg{
					HttpMethod: "GET",
					Path:       "{project}/global/users/{resource}/getIamPolicy",
					ReplacementKeys: map[string]string{
						"projects": "project",
						"users":    "resource",
					},
				},
			},
		},
	},
	"projects/zones/disks": {
		"compute": {
			"alpha": &IamResourceConfig{
				Service: generatedServices["compute"]["alpha"],
				SetIamPolicy: &HttpMethodCfg{
					HttpMethod: "POST",
					Path:       "{project}/zones/{zone}/disks/{resource}/setIamPolicy",
					ReplacementKeys: map[string]string{
						"disks":    "resource",
						"projects": "project",
						"zones":    "zone",
					},
				},
				GetIamPolicy: &HttpMethodCfg{
					HttpMethod: "GET",
					Path:       "{project}/zones/{zone}/disks/{resource}/getIamPolicy",
					ReplacementKeys: map[string]string{
						"disks":    "resource",
						"projects": "project",
						"zones":    "zone",
					},
				},
			},
		},
	},
	"projects/zones/hosts": {
		"compute": {
			"alpha": &IamResourceConfig{
				Service: generatedServices["compute"]["alpha"],
				SetIamPolicy: &HttpMethodCfg{
					HttpMethod: "POST",
					Path:       "{project}/zones/{zone}/hosts/{resource}/setIamPolicy",
					ReplacementKeys: map[string]string{
						"hosts":    "resource",
						"projects": "project",
						"zones":    "zone",
					},
				},
				GetIamPolicy: &HttpMethodCfg{
					HttpMethod: "GET",
					Path:       "{project}/zones/{zone}/hosts/{resource}/getIamPolicy",
					ReplacementKeys: map[string]string{
						"hosts":    "resource",
						"projects": "project",
						"zones":    "zone",
					},
				},
			},
		},
	},
	"projects/zones/instances": {
		"compute": {
			"alpha": &IamResourceConfig{
				Service: generatedServices["compute"]["alpha"],
				SetIamPolicy: &HttpMethodCfg{
					HttpMethod: "POST",
					Path:       "{project}/zones/{zone}/instances/{resource}/setIamPolicy",
					ReplacementKeys: map[string]string{
						"instances": "resource",
						"projects":  "project",
						"zones":     "zone",
					},
				},
				GetIamPolicy: &HttpMethodCfg{
					HttpMethod: "GET",
					Path:       "{project}/zones/{zone}/instances/{resource}/getIamPolicy",
					ReplacementKeys: map[string]string{
						"instances": "resource",
						"projects":  "project",
						"zones":     "zone",
					},
				},
			},
		},
	},
	"projects/zones/nodeGroups": {
		"compute": {
			"alpha": &IamResourceConfig{
				Service: generatedServices["compute"]["alpha"],
				SetIamPolicy: &HttpMethodCfg{
					HttpMethod: "POST",
					Path:       "{project}/zones/{zone}/nodeGroups/{resource}/setIamPolicy",
					ReplacementKeys: map[string]string{
						"nodeGroups": "resource",
						"projects":   "project",
						"zones":      "zone",
					},
				},
				GetIamPolicy: &HttpMethodCfg{
					HttpMethod: "GET",
					Path:       "{project}/zones/{zone}/nodeGroups/{resource}/getIamPolicy",
					ReplacementKeys: map[string]string{
						"nodeGroups": "resource",
						"projects":   "project",
						"zones":      "zone",
					},
				},
			},
		},
	},
	"services": {
		"servicemanagement": {
			"v1": &IamResourceConfig{
				Service: generatedServices["servicemanagement"]["v1"],
				SetIamPolicy: &HttpMethodCfg{
					HttpMethod: "POST",
					Path:       "v1/services/{servicesId}:setIamPolicy",
					ReplacementKeys: map[string]string{
						"services": "servicesId",
					},
				},
				GetIamPolicy: &HttpMethodCfg{
					HttpMethod: "POST",
					Path:       "v1/services/{servicesId}:getIamPolicy",
					ReplacementKeys: map[string]string{
						"services": "servicesId",
					},
				},
			},
		},
	},
	"services/consumers": {
		"servicemanagement": {
			"v1": &IamResourceConfig{
				Service: generatedServices["servicemanagement"]["v1"],
				SetIamPolicy: &HttpMethodCfg{
					HttpMethod: "POST",
					Path:       "v1/services/{servicesId}/consumers/{consumersId}:setIamPolicy",
					ReplacementKeys: map[string]string{
						"consumers": "consumersId",
						"services":  "servicesId",
					},
				},
				GetIamPolicy: &HttpMethodCfg{
					HttpMethod: "POST",
					Path:       "v1/services/{servicesId}/consumers/{consumersId}:getIamPolicy",
					ReplacementKeys: map[string]string{
						"consumers": "consumersId",
						"services":  "servicesId",
					},
				},
			},
		},
	},
}
