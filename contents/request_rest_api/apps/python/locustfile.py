from locust import FastHttpUser, task


class HealthUser(FastHttpUser):
    @task
    def health(self):
        self.client.get("/health/")
