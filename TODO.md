
- Devices should emit data via RMQ broker
  - e.g.: camera devices executables should send video frames when activated by roc calls
- Store RMQ user and pass in another way than env vars
- Ide, Fix missing RMQ user and pass: missing /sensors/modules.yaml - (modules.yaml is in circuit dir, update rpc_client to get the path of modules.yaml as param)
- Modules, Devices, Circuit -> should be created by Crd POST, not helm install
  - We could also use Fleet
- Implement OpenAPI definitions
- Golang process on master node to call the rabbitMQ endpoints, fetch data and update CRDS
- Create Rancher extensions to handle sensors CRDs
- Create a manifest.yaml file for each supported sensors (svg, description, link to homepage)
- Sensor's CRDS should be updated ONLY by controller pod
- Rancher UI should display
  - sensor CRD
  - raspberry CRD
- Define a builder kit for devices
  - The device <-> rpc-server interface should use OpenAPI definition to build a skeleton and call the executable built in devices docker images.
  - The device's API should be defined in the settings file by dev
- Create Circuits to connect devices
- Define job python code using code-server in rancher extension 
- Inject python script in Job's ConfigMap

- Go dependecy injection do define lmbda to apply to Circuits 
  https://medium.com/avenue-tech/dependency-injection-in-go-35293ef7b6
  Google Wire https://github.com/google/wire?tab=readme-ov-file

- Go Plugins to define Circuit workloads
- Circuit crd should have play, stop, pause fields
- Circuit crd should have a base64 to store code-source
- Add diagram to README.md file, mermaid: https://github.com/mermaid-js/mermaid