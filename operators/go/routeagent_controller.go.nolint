package routeagent

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"

	submarinerv1alpha1 "github.com/submariner-operator/submariner-operator/pkg/apis/submariner/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_routeagent")

// Add creates a new Routeagent Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileRouteagent{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("routeagent-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource Routeagent
	err = c.Watch(&source.Kind{Type: &submarinerv1alpha1.Routeagent{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// Watch for changes to secondary resource Pods and requeue the owner Routeagent
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &submarinerv1alpha1.Routeagent{},
	})
	if err != nil {
		return err
	}

	return nil
}

// blank assignment to verify that ReconcileRouteagent implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcileRouteagent{}

// ReconcileRouteagent reconciles a Routeagent object
type ReconcileRouteagent struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a Routeagent object and makes changes based on the state read
// and what is in the Routeagent.Spec
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileRouteagent) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling Routeagent")

	// Fetch the Routeagent instance
	instance := &submarinerv1alpha1.Routeagent{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	daemonSet := newRouteAgentDaemonSet(instance)

	// Set Routeagent instance as the owner and controller
	if err := controllerutil.SetControllerReference(instance, daemonSet, r.scheme); err != nil {
		return reconcile.Result{}, err
	}

	found := &appsv1.DaemonSet{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: daemonSet.Name, Namespace: daemonSet.Namespace}, found)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("Creating a new DaemonSet", "DaemonSet.Namespace", daemonSet.Namespace, "DaemonSet.Name", daemonSet.Name)
		err = r.client.Create(context.TODO(), daemonSet)
		if err != nil {
			return reconcile.Result{}, err
		}

		return reconcile.Result{}, nil
	} else if err != nil {
		return reconcile.Result{}, err
	}

	reqLogger.Info("Skip reconcile: DaemonSet already exists", "DaemonSet.Namespace", found.Namespace, "DaemonSet.Name", found.Name)
	return reconcile.Result{}, nil
}

func newRouteAgentDaemonSet(cr *submarinerv1alpha1.Routeagent) *appsv1.DaemonSet {
	labels := map[string]string{
		"app":       "submariner-routeagent",
		"component": "routeagent",
	}

	matchLabels := map[string]string{
		"app": "submariner-routeagent",
	}

	allow_privilege_escalation := true
	privileged := true
	security_context_all_cap_allow_escal := corev1.SecurityContext{
		Capabilities:             &corev1.Capabilities{Add: []corev1.Capability{"ALL"}},
		AllowPrivilegeEscalation: &allow_privilege_escalation,
		Privileged:               &privileged}

	routeAgentDaemonSet := &appsv1.DaemonSet{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: cr.Namespace,
			Name:      "submariner-routeagent",
			Labels:    labels,
		},
		Spec: appsv1.DaemonSetSpec{
			Selector: &metav1.LabelSelector{MatchLabels: matchLabels},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "submariner-routeagent",
							Image: "submariner-route-agent:local",
							// FIXME: Should be entrypoint script, find/use correct file for routeagent
							Command:         []string{"submariner-route-agent.sh"},
							SecurityContext: &security_context_all_cap_allow_escal,
							VolumeMounts: []corev1.VolumeMount{
								{Name: "host-slash", MountPath: "/host", ReadOnly: true},
							},
							Env: []corev1.EnvVar{
								{Name: "SUBMARINER_NAMESPACE", Value: cr.Spec.Namespace},
								{Name: "SUBMARINER_CLUSTERID", Value: cr.Spec.ClusterID},
								{Name: "SUBMARINER_DEBUG", Value: cr.Spec.Debug},
								{Name: "SUBMARINER_CLUSTERCIDR", Value: cr.Spec.ClusterCIDR},
								{Name: "SUBMARINER_SERVICECIDR", Value: cr.Spec.ServiceCIDR},
							},
						},
					},
					// TODO: Use SA submariner-routeagent or submariner?
					ServiceAccountName: "submariner-operator",
					HostNetwork:        true,
					Volumes: []corev1.Volume{
						{Name: "host-slash", VolumeSource: corev1.VolumeSource{HostPath: &corev1.HostPathVolumeSource{Path: "/"}}},
					},
				},
			},
		},
	}

	return routeAgentDaemonSet
}
