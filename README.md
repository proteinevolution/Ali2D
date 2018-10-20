# Ali2D

The following packages are currently included:

## prepare

Build with:
```
cd prepare
go build
```
which will build the `prepare` executable. It reflects the functionality of the former `prepareAli2D.jar`.

### Example call
The call
``
./prepare tcoffee_alignment_7956270.fasta 50 output.json
``
has produced the following JSON:
```
[
  {
    "sequenceFile": "/home/lukas/repos/ali2d/prepare/tcoffee_alignment_7956270.fasta.0.fas",
    "coveredBy": "/home/lukas/repos/ali2d/prepare/tcoffee_alignment_7956270.fasta.4.fas",
    "sequenceIdentity": 78.1021897810219
  },
  {
    "sequenceFile": "/home/lukas/repos/ali2d/prepare/tcoffee_alignment_7956270.fasta.1.fas",
    "coveredBy": "/home/lukas/repos/ali2d/prepare/tcoffee_alignment_7956270.fasta.4.fas",
    "sequenceIdentity": 80.2919708029197
  },
  {
    "sequenceFile": "/home/lukas/repos/ali2d/prepare/tcoffee_alignment_7956270.fasta.2.fas",
    "coveredBy": "/home/lukas/repos/ali2d/prepare/tcoffee_alignment_7956270.fasta.4.fas",
    "sequenceIdentity": 80.88235294117648
  },
  {
    "sequenceFile": "/home/lukas/repos/ali2d/prepare/tcoffee_alignment_7956270.fasta.3.fas",
    "coveredBy": "/home/lukas/repos/ali2d/prepare/tcoffee_alignment_7956270.fasta.4.fas",
    "sequenceIdentity": 70.8029197080292
  },
  {
    "sequenceFile": "/home/lukas/repos/ali2d/prepare/tcoffee_alignment_7956270.fasta.4.fas",
    "coveredBy": null,
    "sequenceIdentity": null
  },
  {
    "sequenceFile": "/home/lukas/repos/ali2d/prepare/tcoffee_alignment_7956270.fasta.5.fas",
    "coveredBy": "/home/lukas/repos/ali2d/prepare/tcoffee_alignment_7956270.fasta.4.fas",
    "sequenceIdentity": 67.62589928057554
  },
  {
    "sequenceFile": "/home/lukas/repos/ali2d/prepare/tcoffee_alignment_7956270.fasta.6.fas",
    "coveredBy": null,
    "sequenceIdentity": null
  }
]
```

