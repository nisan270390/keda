// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/kedacore/keda/pkg/apis/keda/v1alpha1.AuthEnvironment":           schema_pkg_apis_keda_v1alpha1_AuthEnvironment(ref),
		"github.com/kedacore/keda/pkg/apis/keda/v1alpha1.AuthPodIdentity":           schema_pkg_apis_keda_v1alpha1_AuthPodIdentity(ref),
		"github.com/kedacore/keda/pkg/apis/keda/v1alpha1.AuthSecretTargetRef":       schema_pkg_apis_keda_v1alpha1_AuthSecretTargetRef(ref),
		"github.com/kedacore/keda/pkg/apis/keda/v1alpha1.Credential":                schema_pkg_apis_keda_v1alpha1_Credential(ref),
		"github.com/kedacore/keda/pkg/apis/keda/v1alpha1.HashiCorpVault":            schema_pkg_apis_keda_v1alpha1_HashiCorpVault(ref),
		"github.com/kedacore/keda/pkg/apis/keda/v1alpha1.ObjectReference":           schema_pkg_apis_keda_v1alpha1_ObjectReference(ref),
		"github.com/kedacore/keda/pkg/apis/keda/v1alpha1.ScaleTriggers":             schema_pkg_apis_keda_v1alpha1_ScaleTriggers(ref),
		"github.com/kedacore/keda/pkg/apis/keda/v1alpha1.ScaledObject":              schema_pkg_apis_keda_v1alpha1_ScaledObject(ref),
		"github.com/kedacore/keda/pkg/apis/keda/v1alpha1.ScaledObjectAuthRef":       schema_pkg_apis_keda_v1alpha1_ScaledObjectAuthRef(ref),
		"github.com/kedacore/keda/pkg/apis/keda/v1alpha1.ScaledObjectSpec":          schema_pkg_apis_keda_v1alpha1_ScaledObjectSpec(ref),
		"github.com/kedacore/keda/pkg/apis/keda/v1alpha1.ScaledObjectStatus":        schema_pkg_apis_keda_v1alpha1_ScaledObjectStatus(ref),
		"github.com/kedacore/keda/pkg/apis/keda/v1alpha1.TriggerAuthentication":     schema_pkg_apis_keda_v1alpha1_TriggerAuthentication(ref),
		"github.com/kedacore/keda/pkg/apis/keda/v1alpha1.TriggerAuthenticationSpec": schema_pkg_apis_keda_v1alpha1_TriggerAuthenticationSpec(ref),
		"github.com/kedacore/keda/pkg/apis/keda/v1alpha1.VaultSecret":               schema_pkg_apis_keda_v1alpha1_VaultSecret(ref),
	}
}

func schema_pkg_apis_keda_v1alpha1_AuthEnvironment(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "AuthEnvironment is used to authenticate using environment variables in the destination deployment spec",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"parameter": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"name": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"containerName": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"parameter", "name"},
			},
		},
	}
}

func schema_pkg_apis_keda_v1alpha1_AuthPodIdentity(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "AuthPodIdentity allows users to select the platform native identity mechanism",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"provider": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"provider"},
			},
		},
	}
}

func schema_pkg_apis_keda_v1alpha1_AuthSecretTargetRef(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "AuthSecretTargetRef is used to authenticate using a reference to a secret",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"parameter": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"name": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"key": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"parameter", "name", "key"},
			},
		},
	}
}

func schema_pkg_apis_keda_v1alpha1_Credential(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Credential defines the Hashicorp Vault credentials depending on the authentication method",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"token": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"serviceAccount": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
			},
		},
	}
}

func schema_pkg_apis_keda_v1alpha1_HashiCorpVault(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "HashiCorpVault is used to authenticate using Hashicorp Vault",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"address": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"authentication": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"secrets": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "",
							},
						},
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/kedacore/keda/pkg/apis/keda/v1alpha1.VaultSecret"),
									},
								},
							},
						},
					},
					"credetial": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kedacore/keda/pkg/apis/keda/v1alpha1.Credential"),
						},
					},
					"role": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"mount": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"address", "authentication", "secrets"},
			},
		},
		Dependencies: []string{
			"github.com/kedacore/keda/pkg/apis/keda/v1alpha1.Credential", "github.com/kedacore/keda/pkg/apis/keda/v1alpha1.VaultSecret"},
	}
}

func schema_pkg_apis_keda_v1alpha1_ObjectReference(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ObjectReference holds the a reference to the deployment this ScaledObject applies",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"deploymentName": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"containerName": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"deploymentName"},
			},
		},
	}
}

func schema_pkg_apis_keda_v1alpha1_ScaleTriggers(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ScaleTriggers reference the scaler that will be used",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"type": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"name": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"authenticationRef": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kedacore/keda/pkg/apis/keda/v1alpha1.ScaledObjectAuthRef"),
						},
					},
				},
				Required: []string{"type", "metadata"},
			},
		},
		Dependencies: []string{
			"github.com/kedacore/keda/pkg/apis/keda/v1alpha1.ScaledObjectAuthRef"},
	}
}

func schema_pkg_apis_keda_v1alpha1_ScaledObject(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ScaledObject is a specification for a ScaledObject resource",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kedacore/keda/pkg/apis/keda/v1alpha1.ScaledObjectSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kedacore/keda/pkg/apis/keda/v1alpha1.ScaledObjectStatus"),
						},
					},
				},
				Required: []string{"spec"},
			},
		},
		Dependencies: []string{
			"github.com/kedacore/keda/pkg/apis/keda/v1alpha1.ScaledObjectSpec", "github.com/kedacore/keda/pkg/apis/keda/v1alpha1.ScaledObjectStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_keda_v1alpha1_ScaledObjectAuthRef(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ScaledObjectAuthRef points to the TriggerAuthentication object that is used to authenticate the scaler with the environment",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"name"},
			},
		},
	}
}

func schema_pkg_apis_keda_v1alpha1_ScaledObjectSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ScaledObjectSpec is the spec for a ScaledObject resource",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"scaleType": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"scaleTargetRef": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kedacore/keda/pkg/apis/keda/v1alpha1.ObjectReference"),
						},
					},
					"jobTargetRef": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/api/batch/v1.JobSpec"),
						},
					},
					"pollingInterval": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"cooldownPeriod": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"minReplicaCount": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"maxReplicaCount": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"triggers": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "",
							},
						},
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/kedacore/keda/pkg/apis/keda/v1alpha1.ScaleTriggers"),
									},
								},
							},
						},
					},
				},
				Required: []string{"triggers"},
			},
		},
		Dependencies: []string{
			"github.com/kedacore/keda/pkg/apis/keda/v1alpha1.ObjectReference", "github.com/kedacore/keda/pkg/apis/keda/v1alpha1.ScaleTriggers", "k8s.io/api/batch/v1.JobSpec"},
	}
}

func schema_pkg_apis_keda_v1alpha1_ScaledObjectStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ScaledObjectStatus is the status for a ScaledObject resource",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"lastActiveTime": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
					"externalMetricNames": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "",
							},
						},
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.Time"},
	}
}

func schema_pkg_apis_keda_v1alpha1_TriggerAuthentication(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TriggerAuthentication defines how a trigger can authenticate",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kedacore/keda/pkg/apis/keda/v1alpha1.TriggerAuthenticationSpec"),
						},
					},
				},
				Required: []string{"spec"},
			},
		},
		Dependencies: []string{
			"github.com/kedacore/keda/pkg/apis/keda/v1alpha1.TriggerAuthenticationSpec", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_keda_v1alpha1_TriggerAuthenticationSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TriggerAuthenticationSpec defines the various ways to authenticate",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"podIdentity": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kedacore/keda/pkg/apis/keda/v1alpha1.AuthPodIdentity"),
						},
					},
					"secretTargetRef": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "",
							},
						},
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/kedacore/keda/pkg/apis/keda/v1alpha1.AuthSecretTargetRef"),
									},
								},
							},
						},
					},
					"env": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "",
							},
						},
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/kedacore/keda/pkg/apis/keda/v1alpha1.AuthEnvironment"),
									},
								},
							},
						},
					},
					"hashiCorpVault": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kedacore/keda/pkg/apis/keda/v1alpha1.HashiCorpVault"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/kedacore/keda/pkg/apis/keda/v1alpha1.AuthEnvironment", "github.com/kedacore/keda/pkg/apis/keda/v1alpha1.AuthPodIdentity", "github.com/kedacore/keda/pkg/apis/keda/v1alpha1.AuthSecretTargetRef", "github.com/kedacore/keda/pkg/apis/keda/v1alpha1.HashiCorpVault"},
	}
}

func schema_pkg_apis_keda_v1alpha1_VaultSecret(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "VaultSecret defines the mapping between the path of the secret in Vault to the parameter",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"parameter": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"path": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"key": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"parameter", "path", "key"},
			},
		},
	}
}
