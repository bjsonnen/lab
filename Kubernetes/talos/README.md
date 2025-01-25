# Talos Linux

## Setup

1. Install talosctl: `brew install siderolabs/tap/talosctl`
2. Access your future control-plane node and generate configs: `talosctl gen config <cluster-name> https://<control-plane ip>`
3. Apply the config to your future control-plane node: `talosctl apply-config --insecure --nodes 192.168.0.2 --file controlplane.yaml`
4. Wait a few minutes
5. Bootstrap Kubernetes on the control-plane node: `talosctl bootstrap --nodes <control-plane ip> --endpoints <control-plane ip> --talosconfig=./talosconfig`

Important: Make sure you bootstrap before you add worker nodes!

6. Download the kubeconfig file: talosctl kubeconfig --nodes <control-plane ip> --endpoints <control-plane ip> --talosconfig=./talosconfig
7. Add a worker node: `talosctl apply-config --insecure --nodes <worker-ip> --file controlplane.yaml`
