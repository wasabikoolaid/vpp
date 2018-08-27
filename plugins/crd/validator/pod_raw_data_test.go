// Code generated by 'create-test-data.sh' on Sun Aug 26 21:47:11 PDT 2018 DO NOT EDIT.

package validator

func getRawPodTestData() []string {
	return []string{

		`{
			"container": [
				{
					"name": "contiv-etcd"
				}
			],
			"host_ip_address": "10.20.0.2",
			"ip_address": "10.20.0.2",
			"label": [
				{
					"key": "controller-revision-hash",
					"value": "contiv-etcd-695db96cd9"
				},
				{
					"key": "k8s-app",
					"value": "contiv-etcd"
				},
				{
					"key": "statefulset.kubernetes.io/pod-name",
					"value": "contiv-etcd-0"
				}
			],
			"name": "contiv-etcd-0",
			"namespace": "kube-system"
		}`,
		`{
			"container": [
				{
					"name": "contiv-ksr"
				}
			],
			"host_ip_address": "10.20.0.2",
			"ip_address": "10.20.0.2",
			"label": [
				{
					"key": "controller-revision-hash",
					"value": "1646389250"
				},
				{
					"key": "k8s-app",
					"value": "contiv-ksr"
				},
				{
					"key": "pod-template-generation",
					"value": "1"
				}
			],
			"name": "contiv-ksr-wkqvk",
			"namespace": "kube-system"
		}`,
		`{
			"container": [
				{
					"name": "contiv-vswitch",
					"port": [
						{
							"container_port": 9999,
							"host_port": 9999
						}
					]
				}
			],
			"host_ip_address": "10.20.0.11",
			"ip_address": "10.20.0.11",
			"label": [
				{
					"key": "k8s-app",
					"value": "contiv-vswitch"
				},
				{
					"key": "pod-template-generation",
					"value": "1"
				},
				{
					"key": "controller-revision-hash",
					"value": "1265075853"
				}
			],
			"name": "contiv-vswitch-2ffpz",
			"namespace": "kube-system"
		}`,
		`{
			"container": [
				{
					"name": "contiv-vswitch",
					"port": [
						{
							"container_port": 9999,
							"host_port": 9999
						}
					]
				}
			],
			"host_ip_address": "10.20.0.10",
			"ip_address": "10.20.0.10",
			"label": [
				{
					"key": "pod-template-generation",
					"value": "1"
				},
				{
					"key": "controller-revision-hash",
					"value": "1265075853"
				},
				{
					"key": "k8s-app",
					"value": "contiv-vswitch"
				}
			],
			"name": "contiv-vswitch-pn6cl",
			"namespace": "kube-system"
		}`,
		`{
			"container": [
				{
					"name": "contiv-vswitch",
					"port": [
						{
							"container_port": 9999,
							"host_port": 9999
						}
					]
				}
			],
			"host_ip_address": "10.20.0.2",
			"ip_address": "10.20.0.2",
			"label": [
				{
					"key": "controller-revision-hash",
					"value": "1265075853"
				},
				{
					"key": "k8s-app",
					"value": "contiv-vswitch"
				},
				{
					"key": "pod-template-generation",
					"value": "1"
				}
			],
			"name": "contiv-vswitch-tvpcw",
			"namespace": "kube-system"
		}`,
		`{
			"container": [
				{
					"name": "etcd"
				}
			],
			"host_ip_address": "10.20.0.2",
			"ip_address": "10.20.0.2",
			"label": [
				{
					"key": "component",
					"value": "etcd"
				},
				{
					"key": "tier",
					"value": "control-plane"
				}
			],
			"name": "etcd-k8s-master",
			"namespace": "kube-system"
		}`,
		`{
			"container": [
				{
					"name": "kube-apiserver"
				}
			],
			"host_ip_address": "10.20.0.2",
			"ip_address": "10.20.0.2",
			"label": [
				{
					"key": "component",
					"value": "kube-apiserver"
				},
				{
					"key": "tier",
					"value": "control-plane"
				}
			],
			"name": "kube-apiserver-k8s-master",
			"namespace": "kube-system"
		}`,
		`{
			"container": [
				{
					"name": "kube-controller-manager"
				}
			],
			"host_ip_address": "10.20.0.2",
			"ip_address": "10.20.0.2",
			"label": [
				{
					"key": "component",
					"value": "kube-controller-manager"
				},
				{
					"key": "tier",
					"value": "control-plane"
				}
			],
			"name": "kube-controller-manager-k8s-master",
			"namespace": "kube-system"
		}`,
		`{
			"container": [
				{
					"name": "kubedns",
					"port": [
						{
							"container_port": 10053,
							"name": "dns-local",
							"protocol": 1
						},
						{
							"container_port": 10053,
							"name": "dns-tcp-local"
						},
						{
							"container_port": 10055,
							"name": "metrics"
						}
					]
				},
				{
					"name": "dnsmasq",
					"port": [
						{
							"container_port": 53,
							"name": "dns",
							"protocol": 1
						},
						{
							"container_port": 53,
							"name": "dns-tcp"
						}
					]
				},
				{
					"name": "sidecar",
					"port": [
						{
							"container_port": 10054,
							"name": "metrics"
						}
					]
				}
			],
			"host_ip_address": "10.20.0.2",
			"ip_address": "10.1.3.7",
			"label": [
				{
					"key": "k8s-app",
					"value": "kube-dns"
				},
				{
					"key": "pod-template-hash",
					"value": "4290830601"
				}
			],
			"name": "kube-dns-86f4d74b45-mvgbk",
			"namespace": "kube-system"
		}`,
		`{
			"container": [
				{
					"name": "kube-proxy"
				}
			],
			"host_ip_address": "10.20.0.11",
			"ip_address": "10.20.0.11",
			"label": [
				{
					"key": "pod-template-generation",
					"value": "1"
				},
				{
					"key": "controller-revision-hash",
					"value": "2270935963"
				},
				{
					"key": "k8s-app",
					"value": "kube-proxy"
				}
			],
			"name": "kube-proxy-725s2",
			"namespace": "kube-system"
		}`,
		`{
			"container": [
				{
					"name": "kube-proxy"
				}
			],
			"host_ip_address": "10.20.0.10",
			"ip_address": "10.20.0.10",
			"label": [
				{
					"key": "controller-revision-hash",
					"value": "2270935963"
				},
				{
					"key": "k8s-app",
					"value": "kube-proxy"
				},
				{
					"key": "pod-template-generation",
					"value": "1"
				}
			],
			"name": "kube-proxy-rmbnq",
			"namespace": "kube-system"
		}`,
		`{
			"container": [
				{
					"name": "kube-proxy"
				}
			],
			"host_ip_address": "10.20.0.2",
			"ip_address": "10.20.0.2",
			"label": [
				{
					"key": "k8s-app",
					"value": "kube-proxy"
				},
				{
					"key": "pod-template-generation",
					"value": "1"
				},
				{
					"key": "controller-revision-hash",
					"value": "2270935963"
				}
			],
			"name": "kube-proxy-tnpwl",
			"namespace": "kube-system"
		}`,
		`{
			"container": [
				{
					"name": "kube-scheduler"
				}
			],
			"host_ip_address": "10.20.0.2",
			"ip_address": "10.20.0.2",
			"label": [
				{
					"key": "component",
					"value": "kube-scheduler"
				},
				{
					"key": "tier",
					"value": "control-plane"
				}
			],
			"name": "kube-scheduler-k8s-master",
			"namespace": "kube-system"
		}`,
	}
}