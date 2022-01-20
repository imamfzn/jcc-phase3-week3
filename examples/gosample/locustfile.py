from locust import HttpUser, task, between

class SimpleTest(HttpUser):
    wait_time = between(0.5, 2.5)

    @task
    def test(self):
      self.client.get("http://api.simple.local/test", timeout=30)
