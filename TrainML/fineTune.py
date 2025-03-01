from openai import OpenAI

client = OpenAI(api_key="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")

job = client.fine_tuning.jobs.create(
    training_file="file-Cgeij5npe7EV9UauMkbzDx",
    model="gpt-4o-2024-08-06",
    method={
        "type": "dpo",
        "dpo": {
            "hyperparameters": {"beta": 0.1},
        },
    },
)

# {
#   "id": "file-Cgeij5npe7EV9UauMkbzDx",
#   "bytes": 2326,
#   "created_at": 1740692492,
#   "filename": "/Users/hanxu/Desktop/ETHDenver/TrainML/input.jsonl",
#   "object": "file",
#   "purpose": "fine-tune",
#   "status": "processed",
#   "status_details": null,
#   "expires_at": null
# }