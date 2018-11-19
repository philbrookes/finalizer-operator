## Issue with finalizers on CRs

This repo exists to show a potential problem with operators that make use of finalizers. The issue is that when an operator adds a finalizer to a CR in the same namespace as itself and then the namespace is deleted. The operator is deleted before it can clean up it's own finalizers, and so the namespace is never removed.

To see this error occur, I have made an ansible script to make this easier, I tested this myself on a minishift cluster.

To run the ansible script, you of course need ansible, and you also need oc installed locally pointed at an active cluster.

You can then run: `ansible-playbook ansible/playbooks/install.yml`.

After running this script your cluster will have a `finalizer-test` project hanging indefinitely at `terminating`. In this namespace you will see an item with a finalizer that will never be deleted:
```
oc get item example -n finalizer-test -o yaml
```

