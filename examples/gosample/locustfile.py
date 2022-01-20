from locust import HttpUser, task, between

class SimpleTest(HttpUser):
    wait_time = between(0.5, 2.5)

    @task
    def test(self):
      self.client.get("http://localhost:3000/test")
