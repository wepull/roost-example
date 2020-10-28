#! /bin/bash

###################################################################
# This script creates a new user in your unix environment and
# creates a new kubeconfig file in their home directory.
# Switch to the new user's account using su - <USER_NAME>
# thier password is the same as their <USER_NAME>
# Now you can try kubectl commands in this new context
###################################################################

# user and group names to be used in the client certificate
NEW_USER=bob  # the USER_NAME for the new user
PASS=$NEW_USER
group1=frontend
group2=developers

# Create a user and cd into its home directory
sudo useradd -p $(openssl passwd -crypt $PASS) -m -s /bin/bash $NEW_USER && cd /home/$NEW_USER

# Create a private key
sudo openssl genrsa -out $NEW_USER.key 2048

#If the user has multiple groups, this sets the user(/CN) and groups(/O) to be used for RBAC policies
sudo openssl req -new -key $NEW_USER.key \
-out $NEW_USER.csr \
-subj "/CN=$NEW_USER/O=$group1/O=$group2"


# Create CSR resource
cat <<EOF | kubectl apply -f -
apiVersion: certificates.k8s.io/v1beta1
kind: CertificateSigningRequest
metadata:
  name: $NEW_USER
spec:
  username: $NEW_USER
  groups:
  - $group1
  - $group2
  request: $(cat $NEW_USER.csr | base64 | tr -d "\n")
  usages:
  - client auth
EOF

# Approve the CSR
kubectl certificate approve $NEW_USER

# Get the certificate
kubectl get csr/$NEW_USER -o jsonpath={.status.certificate}  \
| base64 --decode > /tmp/$NEW_USER.crt

# copy it to /home/$NEW_USER dir
sudo cp /tmp/$NEW_USER.crt $NEW_USER.crt && rm /tmp/$NEW_USER.crt

# copy kubeconfig template
sudo cp -R /home/$USER/.kube /home/$NEW_USER/.kube

# delete admin context(Roost Desktop) and unset admin user(kubernetes-admin) from the copied file
sudo kubectl --kubeconfig=/home/$NEW_USER/.kube/config config delete-context 'Roost Desktop'
sudo kubectl --kubeconfig=/home/$NEW_USER/.kube/config config 'unset' users.kubernetes-admin 

# Set new user in kubeconfig
sudo kubectl --kubeconfig=/home/$NEW_USER/.kube/config config set-credentials $NEW_USER --client-key=$NEW_USER.key --client-certificate=$NEW_USER.crt --embed-certs=true

# Create new context in kubeconfig
sudo kubectl --kubeconfig=/home/$NEW_USER/.kube/config config set-context $NEW_USER-context --cluster=kubernetes --user=$NEW_USER

# Set appropriate current context in the new kubeconfig file
sudo kubectl --kubeconfig=/home/$NEW_USER/.kube/config config use-context $NEW_USER-context

# Give ownership of the .kube directory to the new user
sudo chown -R $NEW_USER: /home/$NEW_USER/

# Switch to the newly created context
# kubectl config use-context $NEW_USER-context

# an example for role type resource
# kubectl create role developer --verb=create --verb=get --verb=list --verb=update --verb=delete --resource=pods

# binding the above created role with our new user
# kubectl create rolebinding developer-binding-$NEW_USER --role=developer --user=$NEW_USER