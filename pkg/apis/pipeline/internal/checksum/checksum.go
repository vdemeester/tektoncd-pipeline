package checksum

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"maps"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// SignatureAnnotation is the key of signature in annotation map
	SignatureAnnotation = "tekton.dev/signature"
)

// PrepareObjectMeta will remove annotations not configured from user side -- "kubectl-client-side-apply" and "kubectl.kubernetes.io/last-applied-configuration"
// (added when an object is created with `kubectl apply`) to avoid verification failure and extract the signature.
// Returns a copy of the input object metadata with the annotations removed and the object's signature,
// if it is present in the metadata.
func PrepareObjectMeta(in metav1.Object) metav1.ObjectMeta {
	outMeta := metav1.ObjectMeta{}

	// exclude the fields populated by system.
	outMeta.Name = in.GetName()
	outMeta.GenerateName = in.GetGenerateName()
	outMeta.Namespace = in.GetNamespace()

	if in.GetLabels() != nil {
		outMeta.Labels = make(map[string]string)
		maps.Copy(outMeta.Labels, in.GetLabels())
	}

	outMeta.Annotations = make(map[string]string)
	maps.Copy(outMeta.Annotations, in.GetAnnotations())

	// exclude the annotations added by other components
	delete(outMeta.Annotations, "kubectl-client-side-apply")
	delete(outMeta.Annotations, "kubectl.kubernetes.io/last-applied-configuration")
	delete(outMeta.Annotations, SignatureAnnotation)

	return outMeta
}

// ComputeSha256Checksum computes the sha256 checksum of the tekton object.
func ComputeSha256Checksum(obj any) ([]byte, error) {
	ts, err := json.Marshal(obj)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal the object: %w", err)
	}
	h := sha256.New()
	h.Write(ts)
	return h.Sum(nil), nil
}
