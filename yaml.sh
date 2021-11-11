
kubectl create ns squidgame -o yaml --dry-run > squidgame.yaml

kubectl create deployment myapp -n squidgame --image=alexixva/python-flask-distroless --replicas=2 --dry-run -o yaml >> squidgame.yaml
kubectl expose deployment myapp -n squidgame --type=ClusterIP --port=5000 --dry-run -o yaml >> squidgame.yaml

kubectl create deployment kafka -n squidgame --image=alexixva/python-2 --replicas=2 --dry-run -o yaml >> squidgame.yaml
kubectl expose deployment kafka -n squidgame --type=ClusterIP --port=5000 --dry-run -o yaml >> squidgame.yaml

kubectl create deployment rabbit -n squidgame --image=alexixva/python-flask-distroless --replicas=2 --dry-run -o yaml >> squidgame.yaml
kubectl expose deployment rabbit -n squidgame --type=ClusterIP --port=5000 --dry-run -o yaml >> squidgame.yaml

